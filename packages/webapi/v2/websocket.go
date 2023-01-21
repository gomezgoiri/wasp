package v2

import (
	_ "embed"

	"github.com/iotaledger/wasp/packages/publisher"

	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"

	"github.com/iotaledger/hive.go/core/logger"
	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/publisher/publisherws"
)

type webSocketAPI struct {
	pws *publisherws.PublisherWebSocket
}

func addWebSocketEndpoint(e echoswagger.ApiRoot, log *logger.Logger) *webSocketAPI {
	api := &webSocketAPI{
		pws: publisherws.New(log, []string{publisher.ISCEventKindNewBlock, publisher.ISCEventKindReceipt, publisher.ISCEventIssuerVM}),
	}

	e.Echo().GET("/ws", api.handleWebSocket)

	return api
}

func (w *webSocketAPI) handleWebSocket(c echo.Context) error {
	return w.pws.ServeHTTP(isc.EmptyChainID(), c.Response(), c.Request())
}
