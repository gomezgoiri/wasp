// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

#![allow(dead_code)]
#![allow(unused_imports)]

use crate::*;
use crate::coreaccounts::*;

#[derive(Clone)]
pub struct ImmutableFoundryCreateNewResults {
    pub proxy: Proxy,
}

impl ImmutableFoundryCreateNewResults {
    pub fn foundry_sn(&self) -> ScImmutableUint32 {
        ScImmutableUint32::new(self.proxy.root(RESULT_FOUNDRY_SN))
    }
}

#[derive(Clone)]
pub struct MutableFoundryCreateNewResults {
    pub proxy: Proxy,
}

impl MutableFoundryCreateNewResults {
    pub fn new() -> MutableFoundryCreateNewResults {
        MutableFoundryCreateNewResults {
            proxy: results_proxy(),
        }
    }

    pub fn foundry_sn(&self) -> ScMutableUint32 {
        ScMutableUint32::new(self.proxy.root(RESULT_FOUNDRY_SN))
    }
}

#[derive(Clone)]
pub struct ArrayOfImmutableNftID {
    pub(crate) proxy: Proxy,
}

impl ArrayOfImmutableNftID {
    pub fn length(&self) -> u32 {
        self.proxy.length()
    }

    pub fn get_nft_id(&self, index: u32) -> ScImmutableNftID {
        ScImmutableNftID::new(self.proxy.index(index))
    }
}

#[derive(Clone)]
pub struct ImmutableAccountNFTsResults {
    pub proxy: Proxy,
}

impl ImmutableAccountNFTsResults {
    pub fn nft_i_ds(&self) -> ArrayOfImmutableNftID {
        ArrayOfImmutableNftID { proxy: self.proxy.root(RESULT_NFT_I_DS) }
    }
}

#[derive(Clone)]
pub struct ArrayOfMutableNftID {
    pub(crate) proxy: Proxy,
}

impl ArrayOfMutableNftID {
    pub fn append_nft_id(&self) -> ScMutableNftID {
        ScMutableNftID::new(self.proxy.append())
    }

    pub fn clear(&self) {
        self.proxy.clear_array();
    }

    pub fn length(&self) -> u32 {
        self.proxy.length()
    }

    pub fn get_nft_id(&self, index: u32) -> ScMutableNftID {
        ScMutableNftID::new(self.proxy.index(index))
    }
}

#[derive(Clone)]
pub struct MutableAccountNFTsResults {
    pub proxy: Proxy,
}

impl MutableAccountNFTsResults {
    pub fn new() -> MutableAccountNFTsResults {
        MutableAccountNFTsResults {
            proxy: results_proxy(),
        }
    }

    pub fn nft_i_ds(&self) -> ArrayOfMutableNftID {
        ArrayOfMutableNftID { proxy: self.proxy.root(RESULT_NFT_I_DS) }
    }
}

#[derive(Clone)]
pub struct MapAgentIDToImmutableBytes {
    pub(crate) proxy: Proxy,
}

impl MapAgentIDToImmutableBytes {
    pub fn get_bytes(&self, key: &ScAgentID) -> ScImmutableBytes {
        ScImmutableBytes::new(self.proxy.key(&agent_id_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct ImmutableAccountsResults {
    pub proxy: Proxy,
}

impl ImmutableAccountsResults {
    pub fn all_accounts(&self) -> MapAgentIDToImmutableBytes {
        MapAgentIDToImmutableBytes { proxy: self.proxy.clone() }
    }
}

#[derive(Clone)]
pub struct MapAgentIDToMutableBytes {
    pub(crate) proxy: Proxy,
}

impl MapAgentIDToMutableBytes {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_bytes(&self, key: &ScAgentID) -> ScMutableBytes {
        ScMutableBytes::new(self.proxy.key(&agent_id_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct MutableAccountsResults {
    pub proxy: Proxy,
}

impl MutableAccountsResults {
    pub fn new() -> MutableAccountsResults {
        MutableAccountsResults {
            proxy: results_proxy(),
        }
    }

    pub fn all_accounts(&self) -> MapAgentIDToMutableBytes {
        MapAgentIDToMutableBytes { proxy: self.proxy.clone() }
    }
}

#[derive(Clone)]
pub struct MapTokenIDToImmutableBigInt {
    pub(crate) proxy: Proxy,
}

impl MapTokenIDToImmutableBigInt {
    pub fn get_big_int(&self, key: &ScTokenID) -> ScImmutableBigInt {
        ScImmutableBigInt::new(self.proxy.key(&token_id_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct ImmutableBalanceResults {
    pub proxy: Proxy,
}

impl ImmutableBalanceResults {
    pub fn balances(&self) -> MapTokenIDToImmutableBigInt {
        MapTokenIDToImmutableBigInt { proxy: self.proxy.clone() }
    }
}

#[derive(Clone)]
pub struct MapTokenIDToMutableBigInt {
    pub(crate) proxy: Proxy,
}

impl MapTokenIDToMutableBigInt {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_big_int(&self, key: &ScTokenID) -> ScMutableBigInt {
        ScMutableBigInt::new(self.proxy.key(&token_id_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct MutableBalanceResults {
    pub proxy: Proxy,
}

impl MutableBalanceResults {
    pub fn new() -> MutableBalanceResults {
        MutableBalanceResults {
            proxy: results_proxy(),
        }
    }

    pub fn balances(&self) -> MapTokenIDToMutableBigInt {
        MapTokenIDToMutableBigInt { proxy: self.proxy.clone() }
    }
}

#[derive(Clone)]
pub struct ImmutableBalanceBaseTokenResults {
    pub proxy: Proxy,
}

impl ImmutableBalanceBaseTokenResults {
    pub fn balance(&self) -> ScImmutableUint64 {
        ScImmutableUint64::new(self.proxy.root(RESULT_BALANCE))
    }
}

#[derive(Clone)]
pub struct MutableBalanceBaseTokenResults {
    pub proxy: Proxy,
}

impl MutableBalanceBaseTokenResults {
    pub fn new() -> MutableBalanceBaseTokenResults {
        MutableBalanceBaseTokenResults {
            proxy: results_proxy(),
        }
    }

    pub fn balance(&self) -> ScMutableUint64 {
        ScMutableUint64::new(self.proxy.root(RESULT_BALANCE))
    }
}

#[derive(Clone)]
pub struct ImmutableBalanceNativeTokenResults {
    pub proxy: Proxy,
}

impl ImmutableBalanceNativeTokenResults {
    pub fn tokens(&self) -> ScImmutableBigInt {
        ScImmutableBigInt::new(self.proxy.root(RESULT_TOKENS))
    }
}

#[derive(Clone)]
pub struct MutableBalanceNativeTokenResults {
    pub proxy: Proxy,
}

impl MutableBalanceNativeTokenResults {
    pub fn new() -> MutableBalanceNativeTokenResults {
        MutableBalanceNativeTokenResults {
            proxy: results_proxy(),
        }
    }

    pub fn tokens(&self) -> ScMutableBigInt {
        ScMutableBigInt::new(self.proxy.root(RESULT_TOKENS))
    }
}

#[derive(Clone)]
pub struct ImmutableFoundryOutputResults {
    pub proxy: Proxy,
}

impl ImmutableFoundryOutputResults {
    pub fn foundry_output_bin(&self) -> ScImmutableBytes {
        ScImmutableBytes::new(self.proxy.root(RESULT_FOUNDRY_OUTPUT_BIN))
    }
}

#[derive(Clone)]
pub struct MutableFoundryOutputResults {
    pub proxy: Proxy,
}

impl MutableFoundryOutputResults {
    pub fn new() -> MutableFoundryOutputResults {
        MutableFoundryOutputResults {
            proxy: results_proxy(),
        }
    }

    pub fn foundry_output_bin(&self) -> ScMutableBytes {
        ScMutableBytes::new(self.proxy.root(RESULT_FOUNDRY_OUTPUT_BIN))
    }
}

#[derive(Clone)]
pub struct ImmutableGetAccountNonceResults {
    pub proxy: Proxy,
}

impl ImmutableGetAccountNonceResults {
    pub fn account_nonce(&self) -> ScImmutableUint64 {
        ScImmutableUint64::new(self.proxy.root(RESULT_ACCOUNT_NONCE))
    }
}

#[derive(Clone)]
pub struct MutableGetAccountNonceResults {
    pub proxy: Proxy,
}

impl MutableGetAccountNonceResults {
    pub fn new() -> MutableGetAccountNonceResults {
        MutableGetAccountNonceResults {
            proxy: results_proxy(),
        }
    }

    pub fn account_nonce(&self) -> ScMutableUint64 {
        ScMutableUint64::new(self.proxy.root(RESULT_ACCOUNT_NONCE))
    }
}

#[derive(Clone)]
pub struct MapTokenIDToImmutableBytes {
    pub(crate) proxy: Proxy,
}

impl MapTokenIDToImmutableBytes {
    pub fn get_bytes(&self, key: &ScTokenID) -> ScImmutableBytes {
        ScImmutableBytes::new(self.proxy.key(&token_id_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct ImmutableGetNativeTokenIDRegistryResults {
    pub proxy: Proxy,
}

impl ImmutableGetNativeTokenIDRegistryResults {
    pub fn mapping(&self) -> MapTokenIDToImmutableBytes {
        MapTokenIDToImmutableBytes { proxy: self.proxy.clone() }
    }
}

#[derive(Clone)]
pub struct MapTokenIDToMutableBytes {
    pub(crate) proxy: Proxy,
}

impl MapTokenIDToMutableBytes {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_bytes(&self, key: &ScTokenID) -> ScMutableBytes {
        ScMutableBytes::new(self.proxy.key(&token_id_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct MutableGetNativeTokenIDRegistryResults {
    pub proxy: Proxy,
}

impl MutableGetNativeTokenIDRegistryResults {
    pub fn new() -> MutableGetNativeTokenIDRegistryResults {
        MutableGetNativeTokenIDRegistryResults {
            proxy: results_proxy(),
        }
    }

    pub fn mapping(&self) -> MapTokenIDToMutableBytes {
        MapTokenIDToMutableBytes { proxy: self.proxy.clone() }
    }
}

#[derive(Clone)]
pub struct ImmutableNftDataResults {
    pub proxy: Proxy,
}

impl ImmutableNftDataResults {
    pub fn nft_data(&self) -> ScImmutableBytes {
        ScImmutableBytes::new(self.proxy.root(RESULT_NFT_DATA))
    }
}

#[derive(Clone)]
pub struct MutableNftDataResults {
    pub proxy: Proxy,
}

impl MutableNftDataResults {
    pub fn new() -> MutableNftDataResults {
        MutableNftDataResults {
            proxy: results_proxy(),
        }
    }

    pub fn nft_data(&self) -> ScMutableBytes {
        ScMutableBytes::new(self.proxy.root(RESULT_NFT_DATA))
    }
}

#[derive(Clone)]
pub struct ImmutableTotalAssetsResults {
    pub proxy: Proxy,
}

impl ImmutableTotalAssetsResults {
    pub fn assets(&self) -> MapTokenIDToImmutableBigInt {
        MapTokenIDToImmutableBigInt { proxy: self.proxy.clone() }
    }
}

#[derive(Clone)]
pub struct MutableTotalAssetsResults {
    pub proxy: Proxy,
}

impl MutableTotalAssetsResults {
    pub fn new() -> MutableTotalAssetsResults {
        MutableTotalAssetsResults {
            proxy: results_proxy(),
        }
    }

    pub fn assets(&self) -> MapTokenIDToMutableBigInt {
        MapTokenIDToMutableBigInt { proxy: self.proxy.clone() }
    }
}
