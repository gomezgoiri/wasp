// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

// @formatter:off

#![allow(dead_code)]

use std::ptr;

use wasmlib::*;

use crate::consts::*;
use crate::params::*;

pub struct MintSupplyCall {
	pub func: ScFunc,
	pub params: MutableMintSupplyParams,
}

pub struct TransferOwnershipCall {
	pub func: ScFunc,
	pub params: MutableTransferOwnershipParams,
}

pub struct UpdateMetadataCall {
	pub func: ScFunc,
	pub params: MutableUpdateMetadataParams,
}

pub struct GetInfoCall {
	pub func: ScView,
	pub params: MutableGetInfoParams,
}

pub struct ScFuncs {
}

impl ScFuncs {
    pub fn mint_supply(_ctx: & dyn ScFuncCallContext) -> MintSupplyCall {
        let mut f = MintSupplyCall {
            func: ScFunc::new(HSC_NAME, HFUNC_MINT_SUPPLY),
            params: MutableMintSupplyParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
    pub fn transfer_ownership(_ctx: & dyn ScFuncCallContext) -> TransferOwnershipCall {
        let mut f = TransferOwnershipCall {
            func: ScFunc::new(HSC_NAME, HFUNC_TRANSFER_OWNERSHIP),
            params: MutableTransferOwnershipParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
    pub fn update_metadata(_ctx: & dyn ScFuncCallContext) -> UpdateMetadataCall {
        let mut f = UpdateMetadataCall {
            func: ScFunc::new(HSC_NAME, HFUNC_UPDATE_METADATA),
            params: MutableUpdateMetadataParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
    pub fn get_info(_ctx: & dyn ScViewCallContext) -> GetInfoCall {
        let mut f = GetInfoCall {
            func: ScView::new(HSC_NAME, HVIEW_GET_INFO),
            params: MutableGetInfoParams { id: 0 },
        };
        f.func.set_ptrs(&mut f.params.id, ptr::null_mut());
        f
    }
}

// @formatter:on
