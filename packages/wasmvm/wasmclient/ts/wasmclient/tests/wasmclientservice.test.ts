import {WasmClientContext, WasmClientService} from '../lib';
import * as testwasmlib from "testwasmlib";
import {bytesFromString} from "wasmlib";
import {KeyPair} from "../lib/isc";

const MYCHAIN = "tst1pql4kl8frx2nfe2cvmxshy8jsnq6crnh72jaywuxdrt0xnyylmh5ged3r92";
const MYSEED = "0x1fcf86092b3f7747335ea6838ed3445d98646b6f2285e7336fbf1a35563803dc";

function setupClient() {
    const svc = new WasmClientService('127.0.0.1:9090', '127.0.0.1:5550');
    const ctx = new WasmClientContext(svc, MYCHAIN, "testwasmlib");
    ctx.signRequests(KeyPair.fromSubSeed(bytesFromString(MYSEED), 0n));
    expect(ctx.Err == null).toBeTruthy();
    return ctx;
}

describe('wasmclient', function () {

    describe('Create service', function () {
        it('should create service', () => {
            const client = WasmClientService.DefaultWasmClientService();
            expect(client != null).toBeTruthy();
        });
    });

    describe('Create SC func', function () {
        it('should create SC func', () => {
            const n = testwasmlib.HScName;
            expect(n == testwasmlib.HScName).toBeTruthy();
        });
    });

    describe('Call web API', function () {
        it('should call web API', () => {
            const ctx = setupClient();

            const v = testwasmlib.ScFuncs.getRandom(ctx);
            v.func.call();
            expect(ctx.Err == null).toBeTruthy();
            const rnd = v.results.random().value();
            expect(rnd != 0n).toBeTruthy();
        });
    });

    describe('Post web API', function () {
        it('should post to web API', () => {
            const ctx = setupClient();

            const f = testwasmlib.ScFuncs.random(ctx);
            f.func.post();
            expect(ctx.Err == null).toBeTruthy();

            ctx.waitRequest();
            expect(ctx.Err == null).toBeTruthy();

            const v = testwasmlib.ScFuncs.getRandom(ctx);
            v.func.call();
            expect(ctx.Err == null).toBeTruthy();
            const rnd = v.results.random().value();
            expect(rnd != 0n).toBeTruthy();
        });
    });

    describe('Event handling', function () {
        it('should receive events', () => {
            const ctx = setupClient();

            const events = new testwasmlib.TestWasmLibEventHandlers();
            let name = "";
            events.onTestWasmLibTest((e) => {
                name = e.name;
            })
            ctx.register(events);

            // get new triggerEvent interface, pass params, and post the request
            const f = testwasmlib.ScFuncs.triggerEvent(ctx);
            f.params.name().setValue("Lala");
            f.params.address().setValue(ctx.currentChainID().address());
            f.func.post();
            expect(ctx.Err == null).toBeTruthy();

            ctx.waitRequest();
            expect(ctx.Err == null).toBeTruthy();

            // make sure we wait for the event to show up
            ctx.waitEvent();
            expect(ctx.Err == null).toBeTruthy();

            expect(name == "Lala").toBeTruthy();
        });
    });
});
