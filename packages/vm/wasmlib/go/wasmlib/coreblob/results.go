// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package coreblob

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type ImmutableStoreBlobResults struct {
	id int32
}

func (s ImmutableStoreBlobResults) Hash() wasmlib.ScImmutableHash {
	return wasmlib.NewScImmutableHash(s.id, wasmlib.KeyID(ResultHash))
}

type MutableStoreBlobResults struct {
	id int32
}

func (s MutableStoreBlobResults) Hash() wasmlib.ScMutableHash {
	return wasmlib.NewScMutableHash(s.id, wasmlib.KeyID(ResultHash))
}

type ImmutableGetBlobFieldResults struct {
	id int32
}

func (s ImmutableGetBlobFieldResults) Bytes() wasmlib.ScImmutableBytes {
	return wasmlib.NewScImmutableBytes(s.id, wasmlib.KeyID(ResultBytes))
}

type MutableGetBlobFieldResults struct {
	id int32
}

func (s MutableGetBlobFieldResults) Bytes() wasmlib.ScMutableBytes {
	return wasmlib.NewScMutableBytes(s.id, wasmlib.KeyID(ResultBytes))
}

type MapStringToImmutableInt32 struct {
	objID int32
}

func (m MapStringToImmutableInt32) GetInt32(key string) wasmlib.ScImmutableInt32 {
	return wasmlib.NewScImmutableInt32(m.objID, wasmlib.Key(key).KeyID())
}

type ImmutableGetBlobInfoResults struct {
	id int32
}

func (s ImmutableGetBlobInfoResults) BlobSizes() MapStringToImmutableInt32 {
	return MapStringToImmutableInt32{objID: s.id}
}

type MapStringToMutableInt32 struct {
	objID int32
}

func (m MapStringToMutableInt32) Clear() {
	wasmlib.Clear(m.objID)
}

func (m MapStringToMutableInt32) GetInt32(key string) wasmlib.ScMutableInt32 {
	return wasmlib.NewScMutableInt32(m.objID, wasmlib.Key(key).KeyID())
}

type MutableGetBlobInfoResults struct {
	id int32
}

func (s MutableGetBlobInfoResults) BlobSizes() MapStringToMutableInt32 {
	return MapStringToMutableInt32{objID: s.id}
}

type MapHashToImmutableInt32 struct {
	objID int32
}

func (m MapHashToImmutableInt32) GetInt32(key wasmlib.ScHash) wasmlib.ScImmutableInt32 {
	return wasmlib.NewScImmutableInt32(m.objID, key.KeyID())
}

type ImmutableListBlobsResults struct {
	id int32
}

func (s ImmutableListBlobsResults) BlobSizes() MapHashToImmutableInt32 {
	return MapHashToImmutableInt32{objID: s.id}
}

type MapHashToMutableInt32 struct {
	objID int32
}

func (m MapHashToMutableInt32) Clear() {
	wasmlib.Clear(m.objID)
}

func (m MapHashToMutableInt32) GetInt32(key wasmlib.ScHash) wasmlib.ScMutableInt32 {
	return wasmlib.NewScMutableInt32(m.objID, key.KeyID())
}

type MutableListBlobsResults struct {
	id int32
}

func (s MutableListBlobsResults) BlobSizes() MapHashToMutableInt32 {
	return MapHashToMutableInt32{objID: s.id}
}