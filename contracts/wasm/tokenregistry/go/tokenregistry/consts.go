// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:revive
package tokenregistry

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

const (
	ScName        = "tokenregistry"
	ScDescription = ""
	HScName       = wasmtypes.ScHname(0xe1ba0c78)
)

const (
	ParamDescription = "description"
	ParamToken       = "token"
	ParamUserDefined = "userDefined"
)

const (
	StateRegistry  = "registry"
	StateTokenList = "tokenList"
)

const (
	FuncMintSupply        = "mintSupply"
	FuncTransferOwnership = "transferOwnership"
	FuncUpdateMetadata    = "updateMetadata"
	ViewGetInfo           = "getInfo"
)

const (
	HFuncMintSupply        = wasmtypes.ScHname(0x564349a7)
	HFuncTransferOwnership = wasmtypes.ScHname(0xbb9eb5af)
	HFuncUpdateMetadata    = wasmtypes.ScHname(0xa26b23b6)
	HViewGetInfo           = wasmtypes.ScHname(0xcfedba5f)
)
