package info

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"

	"github.com/iotaledger/wasp/packages/peering"
	"github.com/iotaledger/wasp/packages/webapi/v1/model"
	"github.com/iotaledger/wasp/packages/webapi/v1/routes"
)

type infoService struct {
	waspVersion   string
	network       peering.NetworkProvider
	publisherPort int
}

func AddEndpoints(server echoswagger.ApiRouter, waspVersion string, network peering.NetworkProvider, publisherPort int) {
	s := &infoService{
		waspVersion:   waspVersion,
		network:       network,
		publisherPort: publisherPort,
	}

	server.GET(routes.Info(), s.handleInfo).
		SetDeprecated().
		SetSummary("Get information about the node").
		AddResponse(http.StatusOK, "Node properties", model.InfoResponse{}, nil)
}

func (s *infoService) handleInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, model.InfoResponse{
		Version:       s.waspVersion,
		NetworkID:     s.network.Self().NetID(),
		PublisherPort: s.publisherPort,
	})
}
