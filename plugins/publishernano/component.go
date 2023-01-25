package publishernano

import (
	"context"
	"fmt"
	"go.nanomsg.org/mangos/v3"
	"go.nanomsg.org/mangos/v3/protocol/pub"
	_ "go.nanomsg.org/mangos/v3/transport/all"
	"go.uber.org/dig"

	"github.com/iotaledger/hive.go/core/app"
	"github.com/iotaledger/hive.go/core/events"
	"github.com/iotaledger/wasp/packages/daemon"
	"github.com/iotaledger/wasp/packages/publisher"
)

func init() {
	Plugin = &app.Plugin{
		Component: &app.Component{
			Name:           "PublisherNano",
			Params:         params,
			InitConfigPars: initConfigPars,
			Run:            run,
		},
		IsEnabled: func() bool {
			return ParamsPublisher.Enabled
		},
	}
}

var Plugin *app.Plugin

func initConfigPars(c *dig.Container) error {
	type cfgResult struct {
		dig.Out
		PublisherPort int `name:"publisherPort"`
	}

	if err := c.Provide(func() cfgResult {
		return cfgResult{
			PublisherPort: ParamsPublisher.Port,
		}
	}); err != nil {
		Plugin.LogPanic(err)
	}

	return nil
}

func run() error {
	messages := make(chan []byte, 100)

	port := ParamsPublisher.Port
	socket, err := openSocket(port)
	if err != nil {
		Plugin.LogErrorf("failed to initialize publisher: %w", err)
		return err
	}
	Plugin.LogInfof("nanomsg publisher is running on port %d", port)

	err = Plugin.Daemon().BackgroundWorker(Plugin.Name, func(ctx context.Context) {
		for {
			select {
			case msg := <-messages:
				if socket != nil {
					err := socket.Send(msg)
					if err != nil {
						Plugin.LogErrorf("failed to publish message: %w", err)
					}
				}
			case <-ctx.Done():
				if socket != nil {
					_ = socket.Close()
					socket = nil
				}
				return
			}
		}
	}, daemon.PriorityNanoMsg)
	if err != nil {
		panic(err)
	}

	publisher.Event.Hook(events.NewClosure(func(event *publisher.ISCEvent) {
		msg := event.Kind
		select {
		case messages <- []byte(msg):
		default:
			Plugin.LogWarnf("Failed to publish message: [%s]", msg)
		}
	}))

	return nil
}

func openSocket(port int) (mangos.Socket, error) {
	socket, err := pub.NewSocket()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("tcp://:%d", port)
	if err := socket.Listen(url); err != nil {
		return nil, err
	}
	return socket, nil
}
