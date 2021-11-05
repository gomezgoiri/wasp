// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;
use wasmlib::host::*;

use crate::*;
use crate::keys::*;

#[derive(Clone, Copy)]
pub struct ImmutableMintSupplyParams {
    pub(crate) id: i32,
}

impl ImmutableMintSupplyParams {

    pub fn description(&self) -> ScImmutableString {
		ScImmutableString::new(self.id, idx_map(IDX_PARAM_DESCRIPTION))
	}

    pub fn user_defined(&self) -> ScImmutableString {
		ScImmutableString::new(self.id, idx_map(IDX_PARAM_USER_DEFINED))
	}
}

#[derive(Clone, Copy)]
pub struct MutableMintSupplyParams {
    pub(crate) id: i32,
}

impl MutableMintSupplyParams {

    pub fn description(&self) -> ScMutableString {
		ScMutableString::new(self.id, idx_map(IDX_PARAM_DESCRIPTION))
	}

    pub fn user_defined(&self) -> ScMutableString {
		ScMutableString::new(self.id, idx_map(IDX_PARAM_USER_DEFINED))
	}
}

#[derive(Clone, Copy)]
pub struct ImmutableTransferOwnershipParams {
    pub(crate) id: i32,
}

impl ImmutableTransferOwnershipParams {

    pub fn color(&self) -> ScImmutableColor {
		ScImmutableColor::new(self.id, idx_map(IDX_PARAM_COLOR))
	}
}

#[derive(Clone, Copy)]
pub struct MutableTransferOwnershipParams {
    pub(crate) id: i32,
}

impl MutableTransferOwnershipParams {

    pub fn color(&self) -> ScMutableColor {
		ScMutableColor::new(self.id, idx_map(IDX_PARAM_COLOR))
	}
}

#[derive(Clone, Copy)]
pub struct ImmutableUpdateMetadataParams {
    pub(crate) id: i32,
}

impl ImmutableUpdateMetadataParams {

    pub fn color(&self) -> ScImmutableColor {
		ScImmutableColor::new(self.id, idx_map(IDX_PARAM_COLOR))
	}
}

#[derive(Clone, Copy)]
pub struct MutableUpdateMetadataParams {
    pub(crate) id: i32,
}

impl MutableUpdateMetadataParams {

    pub fn color(&self) -> ScMutableColor {
		ScMutableColor::new(self.id, idx_map(IDX_PARAM_COLOR))
	}
}

#[derive(Clone, Copy)]
pub struct ImmutableGetInfoParams {
    pub(crate) id: i32,
}

impl ImmutableGetInfoParams {

    pub fn color(&self) -> ScImmutableColor {
		ScImmutableColor::new(self.id, idx_map(IDX_PARAM_COLOR))
	}
}

#[derive(Clone, Copy)]
pub struct MutableGetInfoParams {
    pub(crate) id: i32,
}

impl MutableGetInfoParams {

    pub fn color(&self) -> ScMutableColor {
		ScMutableColor::new(self.id, idx_map(IDX_PARAM_COLOR))
	}
}
