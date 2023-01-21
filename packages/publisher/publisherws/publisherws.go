// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package publisherws

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/iotaledger/hive.go/core/events"
	"github.com/iotaledger/hive.go/core/logger"
	"github.com/iotaledger/hive.go/core/subscriptionmanager"
	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/publisher"
)

type PublisherWebSocket struct {
	log                 *logger.Logger
	msgTypes            map[string]bool
	subscriptionManager *subscriptionmanager.SubscriptionManager[string, string]
	sessions            map[string]*WebSocketSession
}

func New(log *logger.Logger, msgTypes []string) *PublisherWebSocket {
	msgTypesMap := make(map[string]bool)
	for _, t := range msgTypes {
		msgTypesMap[t] = true
	}

	subscriptionManager := subscriptionmanager.New(
		subscriptionmanager.WithMaxTopicSubscriptionsPerClient[string, string](5),
	)

	return &PublisherWebSocket{
		log:                 log.Named("PublisherWebSocket"),
		msgTypes:            msgTypesMap,
		subscriptionManager: subscriptionManager,
		sessions:            map[string]*WebSocketSession{},
	}
}

func (p *PublisherWebSocket) createEventWriter(session *WebSocketSession, chainID isc.ChainID, request *http.Request) *events.Closure {
	eventClosure := events.NewClosure(func(msgType string, parts []string) {
		if !p.msgTypes[msgType] {
			return
		}

		if len(parts) < 1 {
			return
		}

		if !chainID.Empty() {
			if parts[0] != chainID.String() {
				return
			}
		}

		if !p.subscriptionManager.HasSubscribers(msgType) {
			return
		}

		msg := msgType + " " + strings.Join(parts, " ")

		if err := session.WriteMessage([]byte(msg)); err != nil {
			p.log.Warnf("dropping websocket message for %s, reason: %v", request.RemoteAddr, err)
		}
	})

	return eventClosure
}

const (
	CommandSubscribe   = "subscribe"
	CommandUnsubscribe = "unsubscribe"
)

type ControlCommand struct {
	Command  string `json:"command"`
	Argument string `json:"argument"`
}

func (p *PublisherWebSocket) handleSubscriptionManagerCommands(connectionID string, message []byte) {
	var command ControlCommand

	if err := json.Unmarshal(message, &command); err != nil {
		p.log.Warnf("Could not deserialize message to type ControlCommand")
		return
	}

	switch command.Command {
	case CommandSubscribe:
		p.subscriptionManager.Subscribe(connectionID, command.Argument)
	case CommandUnsubscribe:
		p.subscriptionManager.Unsubscribe(connectionID, command.Argument)
	}
}

// ServeHTTP serves the websocket.
// Provide a chainID to filter for a certain chain, provide an empty chain id to get all chain events.
func (p *PublisherWebSocket) ServeHTTP(chainID isc.ChainID, responseWriter http.ResponseWriter, request *http.Request) error {
	session, err := AcceptWebSocketSession(responseWriter, request)
	if err != nil {
		return err
	}

	session.OnMessage = func(connectionID string, message []byte) {
		p.log.Infof("New message! [ConID: %v] [Message: '%v']", connectionID, message)
		p.handleSubscriptionManagerCommands(connectionID, message)
	}

	session.OnConnect = func(connectionID string) {
		p.log.Infof("Websocket connection created [ConID: %v]", connectionID)
		p.sessions[connectionID] = session
		p.subscriptionManager.Connect(connectionID)
	}

	session.OnDisconnect = func(connectionID string) {
		p.log.Infof("Websocket connection dropped [ConID: %v]", connectionID)
		if _, ok := p.sessions[connectionID]; ok {
			delete(p.sessions, connectionID)
			p.subscriptionManager.Disconnect(connectionID)
		}
	}

	session.OnError = func(connectionID string, err error) {
		p.log.Errorf("Websocket connection error [ConID: %v] [Error: %v]", connectionID, err)
	}

	p.log.Debugf("accepted websocket connection from %s", request.RemoteAddr)
	defer p.log.Debugf("closed websocket connection from %s", request.RemoteAddr)

	eventWriter := p.createEventWriter(session, chainID, request)
	publisher.Event.Hook(eventWriter)
	defer publisher.Event.Detach(eventWriter)

	return session.Listen()
}
