// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"time"

	"github.com/pangpanglabs/echoswagger/v2"

	loggerpkg "github.com/iotaledger/hive.go/core/logger"
	"github.com/iotaledger/wasp/packages/authentication"
	"github.com/iotaledger/wasp/packages/chains"
	"github.com/iotaledger/wasp/packages/dkg"
	"github.com/iotaledger/wasp/packages/metrics/nodeconnmetrics"
	"github.com/iotaledger/wasp/packages/peering"
	"github.com/iotaledger/wasp/packages/registry"
	"github.com/iotaledger/wasp/packages/users"
	"github.com/iotaledger/wasp/packages/webapi/v1/admapi"
	"github.com/iotaledger/wasp/packages/webapi/v1/evm"
	"github.com/iotaledger/wasp/packages/webapi/v1/info"
	"github.com/iotaledger/wasp/packages/webapi/v1/reqstatus"
	"github.com/iotaledger/wasp/packages/webapi/v1/request"
	"github.com/iotaledger/wasp/packages/webapi/v1/state"
)

func Init(
	logger *loggerpkg.Logger,
	server echoswagger.ApiRoot,
	waspVersion string,
	network peering.NetworkProvider,
	tnm peering.TrustedNetworkManager,
	userManager *users.UserManager,
	chainRecordRegistryProvider registry.ChainRecordRegistryProvider,
	dkShareRegistryProvider registry.DKShareRegistryProvider,
	nodeIdentityProvider registry.NodeIdentityProvider,
	chainsProvider chains.Provider,
	dkgNodeProvider dkg.NodeProvider,
	shutdownFunc admapi.ShutdownFunc,
	nodeConnectionMetrics nodeconnmetrics.NodeConnectionMetrics,
	authConfig authentication.AuthConfiguration,
	nodeOwnerAddresses []string,
	apiCacheTTL time.Duration,
	publisherPort int,
) {
	pub := server.Group("public", "").SetDescription("Public endpoints")
	addWebSocketEndpoint(pub, logger)

	info.AddEndpoints(pub, waspVersion, network, publisherPort)
	reqstatus.AddEndpoints(pub, chainsProvider.ChainProvider())
	state.AddEndpoints(pub, chainsProvider)
	evm.AddEndpoints(pub, chainsProvider, network.Self().PubKey)
	request.AddEndpoints(
		pub,
		chainsProvider.ChainProvider(),
		network.Self().PubKey(),
		apiCacheTTL,
	)

	adm := server.Group("admin", "").SetDescription("Admin endpoints")

	admapi.AddEndpoints(
		adm,
		network,
		tnm,
		userManager,
		chainRecordRegistryProvider,
		dkShareRegistryProvider,
		nodeIdentityProvider,
		chainsProvider,
		dkgNodeProvider,
		shutdownFunc,
		nodeConnectionMetrics,
		authConfig,
		nodeOwnerAddresses,
	)

	logger.Infof("added web api endpoints")
}
