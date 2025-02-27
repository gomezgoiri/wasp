package chain

import (
	"net/http"

	"github.com/pangpanglabs/echoswagger/v2"

	loggerpkg "github.com/iotaledger/hive.go/core/logger"
	"github.com/iotaledger/wasp/packages/authentication"
	"github.com/iotaledger/wasp/packages/authentication/shared/permissions"
	"github.com/iotaledger/wasp/packages/webapi/interfaces"
	"github.com/iotaledger/wasp/packages/webapi/models"
	"github.com/iotaledger/wasp/packages/webapi/params"
)

type Controller struct {
	log *loggerpkg.Logger

	chainService     interfaces.ChainService
	evmService       interfaces.EVMService
	nodeService      interfaces.NodeService
	committeeService interfaces.CommitteeService
	offLedgerService interfaces.OffLedgerService
	registryService  interfaces.RegistryService
	vmService        interfaces.VMService
}

func NewChainController(log *loggerpkg.Logger,
	chainService interfaces.ChainService,
	committeeService interfaces.CommitteeService,
	evmService interfaces.EVMService,
	nodeService interfaces.NodeService,
	offLedgerService interfaces.OffLedgerService,
	registryService interfaces.RegistryService,
	vmService interfaces.VMService,
) interfaces.APIController {
	return &Controller{
		log:              log,
		chainService:     chainService,
		evmService:       evmService,
		committeeService: committeeService,
		nodeService:      nodeService,
		offLedgerService: offLedgerService,
		registryService:  registryService,
		vmService:        vmService,
	}
}

func (c *Controller) Name() string {
	return "chains"
}

func (c *Controller) RegisterPublic(publicAPI echoswagger.ApiGroup, mocker interfaces.Mocker) {
	// Echoswagger does not support ANY, so create a fake route, and overwrite it with Echo ANY afterwords.
	evmURL := "chains/:chainID/evm"
	publicAPI.
		GET(evmURL, c.handleJSONRPC).
		AddParamPath("", params.ParamChainID, params.DescriptionChainID).
		AddResponse(http.StatusOK, "The evm json RPC", "", nil).
		AddResponse(http.StatusNotFound, "The evm json RPC failure", "", nil)

	publicAPI.
		EchoGroup().Any("chains/:chainID/evm", c.handleJSONRPC)

	publicAPI.GET("chains/:chainID/evm/tx/:txHash", c.getRequestID).
		AddParamPath("", params.ParamChainID, params.DescriptionChainID).
		AddParamPath("", params.ParamTxHash, params.DescriptionTxHash).
		AddResponse(http.StatusOK, "Request ID", mocker.Get(models.RequestIDResponse{}), nil).
		AddResponse(http.StatusNotFound, "Request ID not found", "", nil).
		SetSummary("Get the ISC request ID for the given Ethereum transaction hash").
		SetOperationId("getRequestIDFromEVMTransactionID")

	publicAPI.GET("chains/:chainID/state/:stateKey", c.getState).
		AddParamPath("", params.ParamChainID, params.DescriptionChainID).
		AddParamPath("", params.ParamStateKey, params.DescriptionStateKey).
		AddResponse(http.StatusOK, "Result", mocker.Get(models.StateResponse{}), nil).
		SetSummary("Fetch the raw value associated with the given key in the chain state").
		SetOperationId("getStateValue")
}

func (c *Controller) RegisterAdmin(adminAPI echoswagger.ApiGroup, mocker interfaces.Mocker) {
	adminAPI.GET("chains", c.getChainList, authentication.ValidatePermissions([]string{permissions.Read})).
		AddResponse(http.StatusOK, "A list of all available chains", mocker.Get([]models.ChainInfoResponse{}), nil).
		SetOperationId("getChains").
		SetSummary("Get a list of all chains")

	adminAPI.POST("chains/:chainID/activate", c.activateChain, authentication.ValidatePermissions([]string{permissions.Write})).
		AddParamPath("", params.ParamChainID, params.DescriptionChainID).
		AddResponse(http.StatusNotModified, "Chain was not activated", nil, nil).
		AddResponse(http.StatusOK, "Chain was successfully activated", nil, nil).
		SetOperationId("activateChain").
		SetSummary("Activate a chain")

	adminAPI.POST("chains/:chainID/deactivate", c.deactivateChain, authentication.ValidatePermissions([]string{permissions.Write})).
		AddParamPath("", params.ParamChainID, params.DescriptionChainID).
		AddResponse(http.StatusNotModified, "Chain was not deactivated", nil, nil).
		AddResponse(http.StatusOK, "Chain was successfully deactivated", nil, nil).
		SetOperationId("deactivateChain").
		SetSummary("Deactivate a chain")

	adminAPI.GET("chains/:chainID", c.getChainInfo, authentication.ValidatePermissions([]string{permissions.Read})).
		AddParamPath("", params.ParamChainID, params.DescriptionChainID).
		AddResponse(http.StatusOK, "Information about a specific chain", mocker.Get(models.ChainInfoResponse{}), nil).
		SetOperationId("getChainInfo").
		SetSummary("Get information about a specific chain")

	adminAPI.GET("chains/:chainID/committee", c.getCommitteeInfo, authentication.ValidatePermissions([]string{permissions.Read})).
		AddParamPath("", params.ParamChainID, params.DescriptionChainID).
		AddResponse(http.StatusOK, "A list of all nodes tied to the chain", mocker.Get(models.CommitteeInfoResponse{}), nil).
		SetOperationId("getCommitteeInfo").
		SetSummary("Get information about the deployed committee")

	adminAPI.GET("chains/:chainID/contracts", c.getContracts, authentication.ValidatePermissions([]string{permissions.Read})).
		AddParamPath("", params.ParamChainID, params.DescriptionChainID).
		AddResponse(http.StatusOK, "A list of all available contracts", mocker.Get([]models.ContractInfoResponse{}), nil).
		SetOperationId("getContracts").
		SetSummary("Get all available chain contracts")

	adminAPI.POST("chains/:chainID/chainrecord", c.setChainRecord, authentication.ValidatePermissions([]string{permissions.Write})).
		AddParamPath("", params.ParamChainID, params.DescriptionChainID).
		AddParamBody(mocker.Get(models.ChainRecord{}), "ChainRecord", "Chain Record", true).
		AddResponse(http.StatusCreated, "Chain record was saved", nil, nil).
		SetSummary("Sets the chain record.").
		SetOperationId("setChainRecord")

	adminAPI.PUT("chains/:chainID/access-node/:peer", c.addAccessNode, authentication.ValidatePermissions([]string{permissions.Write})).
		AddParamPath("", params.ParamChainID, params.DescriptionChainID).
		AddParamPath("", params.ParamPeer, params.DescriptionPeer).
		AddResponse(http.StatusCreated, "Access node was successfully added", nil, nil).
		SetSummary("Configure a trusted node to be an access node.").
		SetOperationId("addAccessNode")

	adminAPI.DELETE("chains/:chainID/access-node/:peer", c.removeAccessNode, authentication.ValidatePermissions([]string{permissions.Write})).
		AddParamPath("", params.ParamChainID, params.DescriptionChainID).
		AddParamPath("", params.ParamPeer, params.DescriptionPeer).
		AddResponse(http.StatusOK, "Access node was successfully removed", nil, nil).
		SetSummary("Remove an access node.").
		SetOperationId("removeAccessNode")
}
