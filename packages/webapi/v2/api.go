package v2

import (
	"time"

	"github.com/pangpanglabs/echoswagger/v2"

	"github.com/iotaledger/hive.go/core/app/pkg/shutdown"
	"github.com/iotaledger/hive.go/core/configuration"
	loggerpkg "github.com/iotaledger/hive.go/core/logger"
	"github.com/iotaledger/wasp/packages/authentication"
	"github.com/iotaledger/wasp/packages/chains"
	"github.com/iotaledger/wasp/packages/dkg"
	"github.com/iotaledger/wasp/packages/metrics/nodeconnmetrics"
	"github.com/iotaledger/wasp/packages/peering"
	"github.com/iotaledger/wasp/packages/registry"
	userspkg "github.com/iotaledger/wasp/packages/users"
	"github.com/iotaledger/wasp/packages/webapi/v2/controllers/chain"
	"github.com/iotaledger/wasp/packages/webapi/v2/controllers/corecontracts"
	"github.com/iotaledger/wasp/packages/webapi/v2/controllers/metrics"
	"github.com/iotaledger/wasp/packages/webapi/v2/controllers/node"
	"github.com/iotaledger/wasp/packages/webapi/v2/controllers/requests"
	"github.com/iotaledger/wasp/packages/webapi/v2/controllers/users"
	"github.com/iotaledger/wasp/packages/webapi/v2/interfaces"
	"github.com/iotaledger/wasp/packages/webapi/v2/services"
)

func loadControllers(server echoswagger.ApiRoot, userManager *userspkg.UserManager, nodeIdentityProvider registry.NodeIdentityProvider, authConfig authentication.AuthConfiguration, mocker *Mocker, controllersToLoad []interfaces.APIController) {
	for _, controller := range controllersToLoad {
		publicGroup := server.Group(controller.Name(), "v2/")

		controller.RegisterPublic(publicGroup, mocker)

		claimValidator := func(claims *authentication.WaspClaims) bool {
			// The v2 api uses another way of permission handling, so we can always return true here.
			// Permissions are now validated at the route level. See the webapi/v2/controllers/*/controller.go routes.
			return true
		}

		adminGroup := server.Group(controller.Name(), "v2/").
			SetSecurity("Authorization")

		authentication.AddAuthentication(adminGroup.EchoGroup(), userManager, nodeIdentityProvider, authConfig, claimValidator)

		controller.RegisterAdmin(adminGroup, mocker)
	}
}

func Init(
	logger *loggerpkg.Logger,
	server echoswagger.ApiRoot,
	waspVersion string,
	config *configuration.Configuration,
	networkProvider peering.NetworkProvider,
	trustedNetworkManager peering.TrustedNetworkManager,
	userManager *userspkg.UserManager,
	chainRecordRegistryProvider registry.ChainRecordRegistryProvider,
	dkShareRegistryProvider registry.DKShareRegistryProvider,
	nodeIdentityProvider registry.NodeIdentityProvider,
	chainsProvider chains.Provider,
	dkgNodeProvider dkg.NodeProvider,
	shutdownHandler *shutdown.ShutdownHandler,
	nodeConnectionMetrics nodeconnmetrics.NodeConnectionMetrics,
	authConfig authentication.AuthConfiguration,
	nodeOwnerAddresses []string,
	requestCacheTTL time.Duration,
) {
	// load mock files to generate correct echo swagger documentation
	mocker := NewMocker()
	mocker.LoadMockFiles()

	vmService := services.NewVMService(chainsProvider)
	chainService := services.NewChainService(chainsProvider, nodeConnectionMetrics, chainRecordRegistryProvider, vmService)
	committeeService := services.NewCommitteeService(chainsProvider, networkProvider, dkShareRegistryProvider)
	registryService := services.NewRegistryService(chainsProvider, chainRecordRegistryProvider)
	offLedgerService := services.NewOffLedgerService(chainService, networkProvider, requestCacheTTL)
	metricsService := services.NewMetricsService(chainsProvider)
	peeringService := services.NewPeeringService(chainsProvider, networkProvider, trustedNetworkManager)
	evmService := services.NewEVMService(chainService, networkProvider)
	nodeService := services.NewNodeService(chainRecordRegistryProvider, nodeOwnerAddresses, nodeIdentityProvider, shutdownHandler, trustedNetworkManager)
	dkgService := services.NewDKGService(dkShareRegistryProvider, dkgNodeProvider)
	userService := services.NewUserService(userManager)
	// --

	controllersToLoad := []interfaces.APIController{
		chain.NewChainController(logger, chainService, committeeService, evmService, nodeService, offLedgerService, registryService, vmService),
		metrics.NewMetricsController(metricsService),
		node.NewNodeController(waspVersion, config, dkgService, nodeService, peeringService),
		requests.NewRequestsController(chainService, offLedgerService, peeringService, vmService),
		users.NewUsersController(userService),
		corecontracts.NewCoreContractsController(vmService),
	}

	addWebSocketEndpoint(server, logger)
	loadControllers(server, userManager, nodeIdentityProvider, authConfig, mocker, controllersToLoad)
}
