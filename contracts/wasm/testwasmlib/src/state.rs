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
use crate::typedefs::*;

pub struct MapStringToImmutableStringArray {
	pub(crate) obj_id: i32,
}

impl MapStringToImmutableStringArray {
    pub fn get_string_array(&self, key: &str) -> ImmutableStringArray {
        let sub_id = get_object_id(self.obj_id, key.get_key_id(), TYPE_ARRAY | TYPE_STRING);
        ImmutableStringArray { obj_id: sub_id }
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableTestWasmLibState {
    pub(crate) id: i32,
}

impl ImmutableTestWasmLibState {

    pub fn arrays(&self) -> MapStringToImmutableStringArray {
		let map_id = get_object_id(self.id, idx_map(IDX_STATE_ARRAYS), TYPE_MAP);
		MapStringToImmutableStringArray { obj_id: map_id }
	}
}

pub struct MapStringToMutableStringArray {
	pub(crate) obj_id: i32,
}

impl MapStringToMutableStringArray {
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    pub fn get_string_array(&self, key: &str) -> MutableStringArray {
        let sub_id = get_object_id(self.obj_id, key.get_key_id(), TYPE_ARRAY | TYPE_STRING);
        MutableStringArray { obj_id: sub_id }
    }
}

#[derive(Clone, Copy)]
pub struct MutableTestWasmLibState {
    pub(crate) id: i32,
}

impl MutableTestWasmLibState {

    pub fn arrays(&self) -> MapStringToMutableStringArray {
		let map_id = get_object_id(self.id, idx_map(IDX_STATE_ARRAYS), TYPE_MAP);
		MapStringToMutableStringArray { obj_id: map_id }
	}
}
