// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

// @formatter:off

#![allow(dead_code)]

use wasmlib::*;

pub const SC_NAME:        &str = "donatewithfeedback";
pub const SC_DESCRIPTION: &str = "";
pub const HSC_NAME:       ScHname = ScHname(0x696d7f66);

pub const PARAM_AMOUNT: &str = "amount";
pub const PARAM_FEEDBACK: &str = "feedback";
pub const PARAM_NR: &str = "nr";

pub const RESULT_AMOUNT: &str = "amount";
pub const RESULT_COUNT: &str = "count";
pub const RESULT_DONATOR: &str = "donator";
pub const RESULT_ERROR: &str = "error";
pub const RESULT_FEEDBACK: &str = "feedback";
pub const RESULT_MAX_DONATION: &str = "maxDonation";
pub const RESULT_TIMESTAMP: &str = "timestamp";
pub const RESULT_TOTAL_DONATION: &str = "totalDonation";

pub const STATE_LOG: &str = "log";
pub const STATE_MAX_DONATION: &str = "maxDonation";
pub const STATE_TOTAL_DONATION: &str = "totalDonation";

pub const FUNC_DONATE:  &str = "donate";
pub const FUNC_WITHDRAW:  &str = "withdraw";
pub const VIEW_DONATION:  &str = "donation";
pub const VIEW_DONATION_INFO:  &str = "donationInfo";

pub const HFUNC_DONATE: ScHname = ScHname(0xdc9b133a);
pub const HFUNC_WITHDRAW: ScHname = ScHname(0x9dcc0f41);
pub const HVIEW_DONATION: ScHname = ScHname(0xbdb245ba);
pub const HVIEW_DONATION_INFO: ScHname = ScHname(0xc8f7c726);

// @formatter:on
