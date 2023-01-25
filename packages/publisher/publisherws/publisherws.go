// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package publisherws

import (
	"encoding/json"
	"github.com/iotaledger/hive.go/core/events"
	"github.com/iotaledger/hive.go/core/logger"
	"github.com/iotaledger/hive.go/core/subscriptionmanager"
	"github.com/iotaledger/hive.go/core/websockethub"
	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/publisher"
	"net/http"
)

type PublisherWebSocket struct {
	hub                 *websockethub.Hub
	log                 *logger.Logger
	msgTypes            map[string]bool
	sessions            map[string]*WebSocketSession
	subscriptionManager *subscriptionmanager.SubscriptionManager[websockethub.ClientID, string]
}

func New(log *logger.Logger, hub *websockethub.Hub, msgTypes []string) *PublisherWebSocket {
	msgTypesMap := make(map[string]bool)
	for _, t := range msgTypes {
		msgTypesMap[t] = true
	}

	subscriptionManager := subscriptionmanager.New(
		subscriptionmanager.WithMaxTopicSubscriptionsPerClient[websockethub.ClientID, string](5),
	)

	return &PublisherWebSocket{
		hub:                 hub,
		log:                 log.Named("PublisherWebSocket"),
		msgTypes:            msgTypesMap,
		subscriptionManager: subscriptionManager,
		sessions:            map[string]*WebSocketSession{},
	}
}

func (p *PublisherWebSocket) createEventWriter(session *websockethub.Client, chainID isc.ChainID) *events.Closure {
	eventClosure := events.NewClosure(func(event *publisher.ISCEvent) {
		if event == nil {
			return
		}

		if !p.msgTypes[event.Kind] {
			return
		}

		if !chainID.Empty() {
			if !event.ChainID.Equals(chainID) {
				return
			}
		}

		if !p.subscriptionManager.ClientSubscribedToTopic(session.ID(), event.Kind) {
			return
		}

		session.Send(event)
	})

	return eventClosure
}

type BaseCommand struct {
	Command string `json:"command"`
}

const (
	CommandSubscribe           = "subscribe"
	CommandClientWasSubscribed = "client_subscribed"

	CommandUnsubscribe           = "unsubscribe"
	CommandClientWasUnsubscribed = "client_unsubscribed"
)

type SubscriptionCommand struct {
	Command string `json:"command"`
	Topic   string `json:"topic"`
}

func (p *PublisherWebSocket) handleSubscriptionCommand(client *websockethub.Client, message []byte) {
	var command SubscriptionCommand
	if err := json.Unmarshal(message, &command); err != nil {
		p.log.Warnf("Could not deserialize message to type ControlCommand")
		return
	}

	switch command.Command {
	case CommandSubscribe:
		p.subscriptionManager.Subscribe(client.ID(), command.Topic)
		client.Send(SubscriptionCommand{
			Command: CommandClientWasSubscribed,
			Topic:   command.Topic,
		})
	case CommandUnsubscribe:
		p.subscriptionManager.Unsubscribe(client.ID(), command.Topic)
		client.Send(SubscriptionCommand{
			Command: CommandClientWasUnsubscribed,
			Topic:   command.Topic,
		})
	}
}

func (p *PublisherWebSocket) handleNodeCommands(client *websockethub.Client, message []byte) {
	p.log.Info(string(message))

	var baseCommand BaseCommand
	if err := json.Unmarshal(message, &baseCommand); err != nil {
		p.log.Warnf("Could not deserialize message to type BaseCommand")
		return
	}

	switch baseCommand.Command {
	case CommandSubscribe, CommandUnsubscribe:
		p.handleSubscriptionCommand(client, message)
	default:
		p.log.Warnf("Could not deserialize message")
	}
}

func (p *PublisherWebSocket) OnClientCreated(chainID isc.ChainID, client *websockethub.Client) {
	client.ReceiveChan = make(chan *websockethub.WebsocketMsg, 100)

	p.log.Info("client created!")

	eventWriter := p.createEventWriter(client, chainID)
	publisher.Event.Hook(eventWriter)
	defer publisher.Event.Detach(eventWriter)

	go func() {
		for {
			select {
			case <-client.ExitSignal:
				// client was disconnected
				p.log.Info("Connection ended")

				return

			case msg, ok := <-client.ReceiveChan:
				if !ok {
					// client was disconnected
					p.log.Info("Connection ended")

					return
				}

				p.log.Info("Message")
				p.handleNodeCommands(client, msg.Data)
			}
		}
	}()
}

func (p *PublisherWebSocket) OnConnect(client *websockethub.Client, request *http.Request) {
	p.log.Infof("accepted websocket connection from %s", request.RemoteAddr)
	p.subscriptionManager.Connect(client.ID())
}

func (p *PublisherWebSocket) OnDisconnect(client *websockethub.Client, request *http.Request) {
	p.subscriptionManager.Disconnect(client.ID())
	p.log.Infof("closed websocket connection from %s", request.RemoteAddr)
}

// ServeHTTP serves the websocket.
// Provide a chainID to filter for a certain chain, provide an empty chain id to get all chain events.
func (p *PublisherWebSocket) ServeHTTP(chainID isc.ChainID, responseWriter http.ResponseWriter, request *http.Request) error {
	p.hub.ServeWebsocket(responseWriter, request,
		func(client *websockethub.Client) {
			p.OnClientCreated(chainID, client)
		}, func(client *websockethub.Client) {
			p.OnConnect(client, request)
		}, func(client *websockethub.Client) {
			p.OnDisconnect(client, request)
		})

	return nil
}
