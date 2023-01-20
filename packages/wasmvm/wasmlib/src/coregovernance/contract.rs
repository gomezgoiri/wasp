// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

#![allow(dead_code)]

use crate::*;
use crate::coregovernance::*;

pub struct AddAllowedStateControllerAddressCall {
    pub func:   ScFunc,
    pub params: MutableAddAllowedStateControllerAddressParams,
}

pub struct AddCandidateNodeCall {
    pub func:   ScFunc,
    pub params: MutableAddCandidateNodeParams,
}

pub struct ChangeAccessNodesCall {
    pub func:   ScFunc,
    pub params: MutableChangeAccessNodesParams,
}

pub struct ClaimChainOwnershipCall {
    pub func: ScFunc,
}

pub struct DelegateChainOwnershipCall {
    pub func:   ScFunc,
    pub params: MutableDelegateChainOwnershipParams,
}

pub struct RemoveAllowedStateControllerAddressCall {
    pub func:   ScFunc,
    pub params: MutableRemoveAllowedStateControllerAddressParams,
}

pub struct RevokeAccessNodeCall {
    pub func:   ScFunc,
    pub params: MutableRevokeAccessNodeParams,
}

pub struct RotateStateControllerCall {
    pub func:   ScFunc,
    pub params: MutableRotateStateControllerParams,
}

pub struct SetChainInfoCall {
    pub func:   ScFunc,
    pub params: MutableSetChainInfoParams,
}

pub struct SetFeePolicyCall {
    pub func:   ScFunc,
    pub params: MutableSetFeePolicyParams,
}

pub struct GetAllowedStateControllerAddressesCall {
    pub func:    ScView,
    pub results: ImmutableGetAllowedStateControllerAddressesResults,
}

pub struct GetChainInfoCall {
    pub func:    ScView,
    pub results: ImmutableGetChainInfoResults,
}

pub struct GetChainNodesCall {
    pub func:    ScView,
    pub results: ImmutableGetChainNodesResults,
}

pub struct GetChainOwnerCall {
    pub func:    ScView,
    pub results: ImmutableGetChainOwnerResults,
}

pub struct GetFeePolicyCall {
    pub func:    ScView,
    pub results: ImmutableGetFeePolicyResults,
}

pub struct GetMaxBlobSizeCall {
    pub func:    ScView,
    pub results: ImmutableGetMaxBlobSizeResults,
}

pub struct ScFuncs {
}

impl ScFuncs {
    pub fn add_allowed_state_controller_address(ctx: &impl ScFuncCallContext) -> AddAllowedStateControllerAddressCall {
        let mut f = AddAllowedStateControllerAddressCall {
            func:    ScFunc::new(ctx, HSC_NAME, HFUNC_ADD_ALLOWED_STATE_CONTROLLER_ADDRESS),
            params:  MutableAddAllowedStateControllerAddressParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    // access nodes
    pub fn add_candidate_node(ctx: &impl ScFuncCallContext) -> AddCandidateNodeCall {
        let mut f = AddCandidateNodeCall {
            func:    ScFunc::new(ctx, HSC_NAME, HFUNC_ADD_CANDIDATE_NODE),
            params:  MutableAddCandidateNodeParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn change_access_nodes(ctx: &impl ScFuncCallContext) -> ChangeAccessNodesCall {
        let mut f = ChangeAccessNodesCall {
            func:    ScFunc::new(ctx, HSC_NAME, HFUNC_CHANGE_ACCESS_NODES),
            params:  MutableChangeAccessNodesParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    // chain owner
    pub fn claim_chain_ownership(ctx: &impl ScFuncCallContext) -> ClaimChainOwnershipCall {
        ClaimChainOwnershipCall {
            func: ScFunc::new(ctx, HSC_NAME, HFUNC_CLAIM_CHAIN_OWNERSHIP),
        }
    }

    pub fn delegate_chain_ownership(ctx: &impl ScFuncCallContext) -> DelegateChainOwnershipCall {
        let mut f = DelegateChainOwnershipCall {
            func:    ScFunc::new(ctx, HSC_NAME, HFUNC_DELEGATE_CHAIN_OWNERSHIP),
            params:  MutableDelegateChainOwnershipParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn remove_allowed_state_controller_address(ctx: &impl ScFuncCallContext) -> RemoveAllowedStateControllerAddressCall {
        let mut f = RemoveAllowedStateControllerAddressCall {
            func:    ScFunc::new(ctx, HSC_NAME, HFUNC_REMOVE_ALLOWED_STATE_CONTROLLER_ADDRESS),
            params:  MutableRemoveAllowedStateControllerAddressParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    pub fn revoke_access_node(ctx: &impl ScFuncCallContext) -> RevokeAccessNodeCall {
        let mut f = RevokeAccessNodeCall {
            func:    ScFunc::new(ctx, HSC_NAME, HFUNC_REVOKE_ACCESS_NODE),
            params:  MutableRevokeAccessNodeParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    // state controller
    pub fn rotate_state_controller(ctx: &impl ScFuncCallContext) -> RotateStateControllerCall {
        let mut f = RotateStateControllerCall {
            func:    ScFunc::new(ctx, HSC_NAME, HFUNC_ROTATE_STATE_CONTROLLER),
            params:  MutableRotateStateControllerParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    // chain info
    pub fn set_chain_info(ctx: &impl ScFuncCallContext) -> SetChainInfoCall {
        let mut f = SetChainInfoCall {
            func:    ScFunc::new(ctx, HSC_NAME, HFUNC_SET_CHAIN_INFO),
            params:  MutableSetChainInfoParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    // fees
    pub fn set_fee_policy(ctx: &impl ScFuncCallContext) -> SetFeePolicyCall {
        let mut f = SetFeePolicyCall {
            func:    ScFunc::new(ctx, HSC_NAME, HFUNC_SET_FEE_POLICY),
            params:  MutableSetFeePolicyParams { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        f
    }

    // state controller
    pub fn get_allowed_state_controller_addresses(ctx: &impl ScViewCallContext) -> GetAllowedStateControllerAddressesCall {
        let mut f = GetAllowedStateControllerAddressesCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_GET_ALLOWED_STATE_CONTROLLER_ADDRESSES),
            results: ImmutableGetAllowedStateControllerAddressesResults { proxy: Proxy::nil() },
        };
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    // chain info
    pub fn get_chain_info(ctx: &impl ScViewCallContext) -> GetChainInfoCall {
        let mut f = GetChainInfoCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_GET_CHAIN_INFO),
            results: ImmutableGetChainInfoResults { proxy: Proxy::nil() },
        };
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    // access nodes
    pub fn get_chain_nodes(ctx: &impl ScViewCallContext) -> GetChainNodesCall {
        let mut f = GetChainNodesCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_GET_CHAIN_NODES),
            results: ImmutableGetChainNodesResults { proxy: Proxy::nil() },
        };
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    // chain owner
    pub fn get_chain_owner(ctx: &impl ScViewCallContext) -> GetChainOwnerCall {
        let mut f = GetChainOwnerCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_GET_CHAIN_OWNER),
            results: ImmutableGetChainOwnerResults { proxy: Proxy::nil() },
        };
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    // fees
    pub fn get_fee_policy(ctx: &impl ScViewCallContext) -> GetFeePolicyCall {
        let mut f = GetFeePolicyCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_GET_FEE_POLICY),
            results: ImmutableGetFeePolicyResults { proxy: Proxy::nil() },
        };
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn get_max_blob_size(ctx: &impl ScViewCallContext) -> GetMaxBlobSizeCall {
        let mut f = GetMaxBlobSizeCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_GET_MAX_BLOB_SIZE),
            results: ImmutableGetMaxBlobSizeResults { proxy: Proxy::nil() },
        };
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }
}
