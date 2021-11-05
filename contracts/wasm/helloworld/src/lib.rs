// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

// @formatter:off

#![allow(dead_code)]
#![allow(unused_imports)]

use helloworld::*;
use wasmlib::*;
use wasmlib::host::*;

use crate::consts::*;
use crate::keys::*;
use crate::results::*;
use crate::state::*;

mod consts;
mod contract;
mod keys;
mod results;
mod state;
mod helloworld;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_func(FUNC_HELLO_WORLD, func_hello_world_thunk);
    exports.add_view(VIEW_GET_HELLO_WORLD, view_get_hello_world_thunk);

    unsafe {
        for i in 0..KEY_MAP_LEN {
            IDX_MAP[i] = get_key_id_from_string(KEY_MAP[i]);
        }
    }
}

pub struct HelloWorldContext {
	state: MutableHelloWorldState,
}

fn func_hello_world_thunk(ctx: &ScFuncContext) {
	ctx.log("helloworld.funcHelloWorld");
	let f = HelloWorldContext {
		state: MutableHelloWorldState {
			id: OBJ_ID_STATE,
		},
	};
	func_hello_world(ctx, &f);
	ctx.log("helloworld.funcHelloWorld ok");
}

pub struct GetHelloWorldContext {
	results: MutableGetHelloWorldResults,
	state: ImmutableHelloWorldState,
}

fn view_get_hello_world_thunk(ctx: &ScViewContext) {
	ctx.log("helloworld.viewGetHelloWorld");
	let f = GetHelloWorldContext {
		results: MutableGetHelloWorldResults {
			id: OBJ_ID_RESULTS,
		},
		state: ImmutableHelloWorldState {
			id: OBJ_ID_STATE,
		},
	};
	view_get_hello_world(ctx, &f);
	ctx.log("helloworld.viewGetHelloWorld ok");
}

// @formatter:on
