// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as client from "wasmlib/client"
import * as events from "./events"

const ArgNumber = "number";
const ArgPlayPeriod = "playPeriod";

const ResLastWinningNumber = "lastWinningNumber";
const ResRoundNumber = "roundNumber";
const ResRoundStartedAt = "roundStartedAt";
const ResRoundStatus = "roundStatus";

export class ForcePayoutFunc {
	
	post(): void {
	//TODO DoPost(null)
	}
}

export class ForceResetFunc {
	
	post(): void {
	//TODO DoPost(null)
	}
}

export class PayWinnersFunc {
	
	post(): void {
	//TODO DoPost(null)
	}
}

export class PlaceBetFunc {
	args: client.Arguments = new client.Arguments();
	
	number(v: client.Int64): void {
		this.args.setInt64(ArgNumber, v);
	}
	
	post(): void {
		this.args.mandatory(ArgNumber);
	//TODO DoPost(this.args)
	}
}

export class PlayPeriodFunc {
	args: client.Arguments = new client.Arguments();
	
	playPeriod(v: client.Int32): void {
		this.args.setInt32(ArgPlayPeriod, v);
	}
	
	post(): void {
		this.args.mandatory(ArgPlayPeriod);
	//TODO DoPost(this.args)
	}
}

export class LastWinningNumberView {

	call(): LastWinningNumberResults {
    	//TODO DoCall(null) instead of new client.Results()
		return new LastWinningNumberResults(new client.Results());
	}
}

export class LastWinningNumberResults {
	res: client.Results;

	constructor(res: client.Results) { this.res = res; }

	lastWinningNumber(): client.Int64 {
		return this.res.getInt64(ResLastWinningNumber);
	}
}

export class RoundNumberView {

	call(): RoundNumberResults {
    	//TODO DoCall(null) instead of new client.Results()
		return new RoundNumberResults(new client.Results());
	}
}

export class RoundNumberResults {
	res: client.Results;

	constructor(res: client.Results) { this.res = res; }

	roundNumber(): client.Int64 {
		return this.res.getInt64(ResRoundNumber);
	}
}

export class RoundStartedAtView {

	call(): RoundStartedAtResults {
    	//TODO DoCall(null) instead of new client.Results()
		return new RoundStartedAtResults(new client.Results());
	}
}

export class RoundStartedAtResults {
	res: client.Results;

	constructor(res: client.Results) { this.res = res; }

	roundStartedAt(): client.Int32 {
		return this.res.getInt32(ResRoundStartedAt);
	}
}

export class RoundStatusView {

	call(): RoundStatusResults {
    	//TODO DoCall(null) instead of new client.Results()
		return new RoundStatusResults(new client.Results());
	}
}

export class RoundStatusResults {
	res: client.Results;

	constructor(res: client.Results) { this.res = res; }

	roundStatus(): client.Int16 {
		return this.res.getInt16(ResRoundStatus);
	}
}

export class FairRouletteService extends client.Service {

	constructor(cl: client.ServiceClient, chainID: string) {
		super(cl, chainID, "df79d138", events.eventHandlers);
	}

	public forcePayout(): ForcePayoutFunc {
    	return new ForcePayoutFunc();
	}

	public forceReset(): ForceResetFunc {
    	return new ForceResetFunc();
	}

	public payWinners(): PayWinnersFunc {
    	return new PayWinnersFunc();
	}

	public placeBet(): PlaceBetFunc {
    	return new PlaceBetFunc();
	}

	public playPeriod(): PlayPeriodFunc {
    	return new PlayPeriodFunc();
	}

	public lastWinningNumber(): LastWinningNumberView {
    	return new LastWinningNumberView();
	}

	public roundNumber(): RoundNumberView {
    	return new RoundNumberView();
	}

	public roundStartedAt(): RoundStartedAtView {
    	return new RoundStartedAtView();
	}

	public roundStatus(): RoundStatusView {
    	return new RoundStatusView();
	}
}
