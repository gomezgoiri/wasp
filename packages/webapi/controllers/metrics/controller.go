package metrics

import (
	"net/http"

	"github.com/pangpanglabs/echoswagger/v2"

	"github.com/iotaledger/wasp/packages/authentication"
	"github.com/iotaledger/wasp/packages/authentication/shared/permissions"
	"github.com/iotaledger/wasp/packages/webapi/interfaces"
	"github.com/iotaledger/wasp/packages/webapi/models"
	"github.com/iotaledger/wasp/packages/webapi/params"
)

type Controller struct {
	chainService   interfaces.ChainService
	metricsService interfaces.MetricsService
}

func NewMetricsController(chainService interfaces.ChainService, metricsService interfaces.MetricsService) interfaces.APIController {
	return &Controller{
		chainService:   chainService,
		metricsService: metricsService,
	}
}

func (c *Controller) Name() string {
	return "metrics"
}

func (c *Controller) RegisterPublic(publicAPI echoswagger.ApiGroup, mocker interfaces.Mocker) {
}

func (c *Controller) RegisterAdmin(adminAPI echoswagger.ApiGroup, mocker interfaces.Mocker) {
	adminAPI.GET("metrics/l1", c.getL1Metrics, authentication.ValidatePermissions([]string{permissions.Read})).
		AddResponse(http.StatusOK, "A list of all available metrics.", models.ChainMetrics{}, nil).
		SetOperationId("getL1Metrics").
		SetSummary("Get accumulated metrics.")

	adminAPI.GET("metrics/chain/:chainID", c.getChainMetrics, authentication.ValidatePermissions([]string{permissions.Read})).
		AddParamPath("", params.ParamChainID, params.DescriptionChainID).
		AddResponse(http.StatusNotFound, "Chain not found", nil, nil).
		AddResponse(http.StatusOK, "A list of all available metrics.", models.ChainMetrics{}, nil).
		SetOperationId("getChainMetrics").
		SetSummary("Get chain specific metrics.")

	adminAPI.GET("metrics/chain/:chainID/workflow", c.getChainWorkflowMetrics, authentication.ValidatePermissions([]string{permissions.Read})).
		AddParamPath("", params.ParamChainID, params.DescriptionChainID).
		AddResponse(http.StatusNotFound, "Chain not found", nil, nil).
		AddResponse(http.StatusOK, "A list of all available metrics.", mocker.Get(models.ConsensusWorkflowMetrics{}), nil).
		SetOperationId("getChainWorkflowMetrics").
		SetSummary("Get chain workflow metrics.")

	adminAPI.GET("metrics/chain/:chainID/pipe", c.getChainPipeMetrics, authentication.ValidatePermissions([]string{permissions.Read})).
		AddParamPath("", params.ParamChainID, params.DescriptionChainID).
		AddResponse(http.StatusNotFound, "Chain not found", nil, nil).
		AddResponse(http.StatusOK, "A list of all available metrics.", mocker.Get(models.ConsensusPipeMetrics{}), nil).
		SetOperationId("getChainPipeMetrics").
		SetSummary("Get chain pipe event metrics.")
}
