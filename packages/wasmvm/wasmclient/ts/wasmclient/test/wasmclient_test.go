// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/iotaledger/wasp/clients/chainclient"
	"github.com/iotaledger/wasp/packages/cryptolib"
	"github.com/iotaledger/wasp/tools/cluster/templates"
	clustertests "github.com/iotaledger/wasp/tools/cluster/tests"
)

func TestDeploy(t *testing.T) {
	t.SkipNow()
	templates.WaspConfig = strings.ReplaceAll(templates.WaspConfig, "rocksdb", "mapdb")
	e := clustertests.SetupWithChain(t)
	templates.WaspConfig = strings.ReplaceAll(templates.WaspConfig, "mapdb", "rocksdb")
	wallet := cryptolib.NewKeyPair()

	// request funds to the wallet that the wasmclient will use
	err := e.Clu.RequestFunds(wallet.Address())
	require.NoError(t, err)

	// deposit funds to the on-chain account
	chClient := chainclient.New(e.Clu.L1Client(), e.Clu.WaspClient(0), e.Chain.ChainID, wallet)
	reqTx, err := chClient.DepositFunds(10_000_000)
	require.NoError(t, err)
	_, err = e.Chain.CommitteeMultiClient().WaitUntilAllRequestsProcessedSuccessfully(e.Chain.ChainID, reqTx, 30*time.Second)
	require.NoError(t, err)

	time.Sleep(time.Hour)
}
