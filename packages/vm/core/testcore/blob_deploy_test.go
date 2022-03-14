// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package testcore

import (
	"testing"

	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core"
	"github.com/iotaledger/wasp/packages/vm/core/root"
	"github.com/stretchr/testify/require"
)

var wasmFile = "sbtests/sbtestsc/testcore_bg.wasm"

func TestDeploy(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")
	hwasm, err := chain.UploadWasmFromFile(nil, wasmFile)
	require.NoError(t, err)

	err = chain.DeployContract(nil, "testCore", hwasm)
	require.NoError(t, err)
}

func TestDeployWasm(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")
	err := chain.DeployWasmContract(nil, "testCore", wasmFile)
	require.NoError(t, err)
}

func TestDeployRubbish(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")
	name := "testCore"
	_, err := chain.FindContract(name)
	require.Error(t, err)
	err = chain.DeployWasmContract(nil, name, "blob_deploy_test.go")
	require.Error(t, err)

	_, err = chain.FindContract(name)
	require.Error(t, err)
}

func TestDeployNotAuthorized(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")
	user1, _ := env.NewKeyPairWithFunds()
	err := chain.DeployWasmContract(user1, "testCore", wasmFile)
	require.Error(t, err)
}

func TestDeployGrant(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")
	user1, addr1 := env.NewKeyPairWithFunds()
	user1AgentID := iscp.NewAgentID(addr1, 0)

	req := solo.NewCallParams(root.Contract.Name, root.FuncGrantDeployPermission.Name,
		root.ParamDeployer, user1AgentID,
	)
	_, err := chain.PostRequestSync(req.AddAssetsIotas(1), nil)
	require.NoError(t, err)

	err = chain.DeployWasmContract(user1, "testCore", wasmFile)
	require.NoError(t, err)

	_, _, contracts := chain.GetInfo()
	require.EqualValues(t, len(corecontracts.All)+1, len(contracts))

	err = chain.DeployWasmContract(user1, "testInccounter2", wasmFile)
	require.NoError(t, err)

	_, _, contracts = chain.GetInfo()
	require.EqualValues(t, len(corecontracts.All)+2, len(contracts))
}

func TestRevokeDeploy(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")
	user1, addr1 := env.NewKeyPairWithFunds()
	user1AgentID := iscp.NewAgentID(addr1, 0)

	req := solo.NewCallParams(root.Contract.Name, root.FuncGrantDeployPermission.Name,
		root.ParamDeployer, user1AgentID,
	)
	_, err := chain.PostRequestSync(req.AddAssetsIotas(1), nil)
	require.NoError(t, err)

	err = chain.DeployWasmContract(user1, "testCore", wasmFile)
	require.NoError(t, err)

	_, _, contracts := chain.GetInfo()
	require.EqualValues(t, len(corecontracts.All)+1, len(contracts))

	req = solo.NewCallParams(root.Contract.Name, root.FuncRevokeDeployPermission.Name,
		root.ParamDeployer, user1AgentID,
	).AddAssetsIotas(1)
	_, err = chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	err = chain.DeployWasmContract(user1, "testInccounter2", wasmFile)
	require.Error(t, err)

	_, _, contracts = chain.GetInfo()
	require.EqualValues(t, len(corecontracts.All)+1, len(contracts))
}

func TestDeployGrantFail(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")
	user1, addr1 := env.NewKeyPairWithFunds()
	user1AgentID := iscp.NewAgentID(addr1, 0)

	req := solo.NewCallParams(root.Contract.Name, root.FuncGrantDeployPermission.Name,
		root.ParamDeployer, user1AgentID,
	)
	_, err := chain.PostRequestSync(req.AddAssetsIotas(1), user1)
	require.Error(t, err)

	err = chain.DeployWasmContract(user1, "testCore", wasmFile)
	require.Error(t, err)
}

func TestOpenDeploymentToAnyone(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")

	userWallet, _ := env.NewKeyPairWithFunds()

	// deployment is closed to anyone by default
	err := chain.DeployWasmContract(userWallet, "testCore", wasmFile)
	require.Error(t, err)

	// enable open deployments
	req := solo.NewCallParams(root.Contract.Name, root.FuncRequireDeployPermissions.Name,
		root.ParamDeployPermissionsEnabled, codec.EncodeBool(false),
	)
	_, err = chain.PostRequestSync(req.AddAssetsIotas(1), nil)
	require.NoError(t, err)

	// deploy should now succeed
	err = chain.DeployWasmContract(userWallet, "testCore1", wasmFile)
	require.NoError(t, err)

	// disable open deployments
	req = solo.NewCallParams(root.Contract.Name, root.FuncRequireDeployPermissions.Name,
		root.ParamDeployPermissionsEnabled, codec.EncodeBool(true),
	)
	_, err = chain.PostRequestSync(req.AddAssetsIotas(1), nil)
	require.NoError(t, err)

	// deployment should fail after "open deployment" is disabled
	err = chain.DeployWasmContract(userWallet, "testCore3", wasmFile)
	require.Error(t, err)
}
