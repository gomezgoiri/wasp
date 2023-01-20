// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package publisherws

import (
	"github.com/iotaledger/hive.go/core/subscriptionmanager"
	"net/http"
	"strings"

	"github.com/iotaledger/hive.go/core/events"
	"github.com/iotaledger/hive.go/core/logger"
	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/publisher"
)

type PublisherWebSocket struct {
	log                 *logger.Logger
	msgTypes            map[string]bool
	subscriptionManager *subscriptionmanager.SubscriptionManager[string, string]
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

		msg := msgType + " " + strings.Join(parts, " ")

		if err := session.WriteMessage(msg); err != nil {
			p.log.Warnf("dropping websocket message for %s, reason: %v", request.RemoteAddr, err)
		}

	})

	return eventClosure
}

// ServeHTTP serves the websocket.
// Provide a chainID to filter for a certain chain, provide an empty chain id to get all chain events.
func (p *PublisherWebSocket) ServeHTTP(chainID isc.ChainID, responseWriter http.ResponseWriter, request *http.Request) error {
	session, err := InitializeWebSocketSessionFromRequest(responseWriter, request)

	if err != nil {
		return err
	}

	session.OnMessage = func(connectionId string, message string) {
		p.log.Infof("New message! [%v] %v", connectionId, message)
	}

	session.OnConnect = func(connectionId string) {
		p.log.Infof("Connection created! [%v]", connectionId)
	}

	session.OnDisconnect = func(connectionId string) {
		p.log.Infof("Connection died! [%v]", connectionId)
	}

	p.log.Debugf("accepted websocket connection from %s", request.RemoteAddr)
	defer p.log.Debugf("closed websocket connection from %s", request.RemoteAddr)

	eventWriter := p.createEventWriter(session, chainID, request)
	publisher.Event.Hook(eventWriter)
	defer publisher.Event.Detach(eventWriter)

	session.Listen()

	return nil
}
