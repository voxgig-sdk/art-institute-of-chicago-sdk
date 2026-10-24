"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const node_path_1 = __importDefault(require("node:path"));
const Fs = __importStar(require("node:fs"));
const node_test_1 = require("node:test");
const node_assert_1 = __importDefault(require("node:assert"));
const live_runner_1 = require("../../live-runner");
const live_entity_1 = require("../../live-entity");
const __1 = require("../../..");
const utility_1 = require("../../utility");
(0, utility_1.loadEnvLocal)(__dirname + '/../../../.env.local');
(0, node_test_1.describe)('AgentRoleEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('ART_INSTITUTE_OF_CHICAGO_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.ArtInstituteOfChicagoSDK.test();
        const ent = testsdk.AgentRole();
        (0, node_assert_1.default)(null != ent);
    });
    class FailHook extends __1.BaseFeature {
        name = 'failhook';
        version = '0.0.1';
        active = true;
        unexpected = 0;
        init() { }
        PreSpec() { throw new Error('agent_role hook failed'); }
        PreUnexpected() { this.unexpected++; }
    }
    (0, node_test_1.test)('stream-error', async () => {
        const offline = { net: { offline: true } };
        await node_assert_1.default.rejects(async () => {
            for await (const _item of __1.ArtInstituteOfChicagoSDK.test(offline).AgentRole().stream('list')) { }
        }, /offline/);
        for await (const _item of __1.ArtInstituteOfChicagoSDK.test(offline).AgentRole()
            .stream('list', undefined, { ctrl: { throw: false } })) { }
        if (null != __1.config.feature?.rbac) {
            const denied = __1.ArtInstituteOfChicagoSDK.test(undefined, { feature: { rbac: { active: true, deny: true } } });
            await node_assert_1.default.rejects(async () => {
                for await (const _item of denied.AgentRole().stream('list')) { }
            }, (err) => 'rbac_denied' === err.code);
        }
    });
    (0, node_test_1.test)('stream-ctrl', async () => {
        const explain = {};
        const ctrl = { explain };
        for await (const _item of __1.ArtInstituteOfChicagoSDK.test().AgentRole().stream('list', undefined, { ctrl })) { }
        node_assert_1.default.deepStrictEqual(Object.keys(ctrl), ['explain']);
        (0, node_assert_1.default)(explain === ctrl.explain && 0 < Object.keys(explain).length);
    });
    (0, node_test_1.test)('unexpected', async () => {
        const hook = new FailHook();
        const client = new __1.ArtInstituteOfChicagoSDK({ feature: { test: { active: true } }, extend: [hook] });
        await node_assert_1.default.rejects(client.AgentRole().list(), /hook failed/);
        (0, node_assert_1.default)(0 < hook.unexpected);
        const fired = hook.unexpected;
        node_assert_1.default.strictEqual(await client.AgentRole().list(undefined, { throw: false }), undefined);
        (0, node_assert_1.default)(fired < hook.unexpected);
    });
    (0, node_test_1.test)('validate', async (t) => {
        if (null == __1.config.feature?.validate) {
            t.skip('feature not present in this SDK: validate');
            return;
        }
        const client = __1.ArtInstituteOfChicagoSDK.test(undefined, { feature: { validate: { active: true } } });
        await node_assert_1.default.rejects(client.AgentRole().list({ "id": 1 }), (err) => 'validate_failed' === err.code);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE;
        for (const op of ['list', 'load']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'agent_role.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "api_link": { "a": true, "h": "Api Link", "n": "api_link", "r": false, "sh": "REST API link for this resource", "t": "`$ANY`", "key$": "api_link", "index$": 0 }, "api_model": { "a": true, "h": "Api Model", "n": "api_model", "r": false, "sh": "REST API resource type or endpoint", "t": "`$ANY`", "key$": "api_model", "index$": 1 }, "id": { "a": true, "h": "Id", "n": "id", "r": false, "sh": "Unique identifier of this resource.", "t": "`$STRING`", "key$": "id", "index$": 2 }, "source_updated_at": { "a": true, "h": "Source Updated At", "n": "source_updated_at", "r": false, "sh": "Date and time the resource was updated in the source system", "t": "`$ANY`", "key$": "source_updated_at", "index$": 3 }, "suggest_autocomplete_all": { "a": true, "h": "Suggest Autocomplete All", "n": "suggest_autocomplete_all", "r": false, "sh": "Internal field to power the `/autosuggest` endpoint.", "t": "`$ANY`", "key$": "suggest_autocomplete_all", "index$": 4 }, "suggest_autocomplete_boosted": { "a": true, "h": "Suggest Autocomplete Boosted", "n": "suggest_autocomplete_boosted", "r": false, "sh": "Internal field to power the `/autocomplete` endpoint.", "t": "`$ANY`", "key$": "suggest_autocomplete_boosted", "index$": 5 }, "timestamp": { "a": true, "h": "Timestamp", "n": "timestamp", "r": false, "sh": "Date and time the record was updated in the aggregator search index", "t": "`$ANY`", "key$": "timestamp", "index$": 6 }, "title": { "a": true, "h": "Title", "n": "title", "r": false, "sh": "The name of this resource", "t": "`$STRING`", "key$": "title", "index$": 7 }, "updated_at": { "a": true, "h": "Updated At", "n": "updated_at", "r": false, "sh": "Date and time the record was updated in the aggregator database", "t": "`$ANY`", "key$": "updated_at", "index$": 8 } }, "id": { "field": "id", "name": "id" }, "name": "agent_role", "op": { "list": { "input": "data", "name": "list", "points": [{ "a": true, "co": { "id": "GET /agent-roles", "source": "openapi3", "version": 2 }, "g": {}, "k": "http", "m": "GET", "o": "/agent-roles", "q": {}, "r": {}, "s": [{ "lit": "agent-roles" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "list" }, "load": { "input": "data", "name": "load", "points": [{ "a": true, "co": { "id": "GET /agent-roles/{id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "GET", "o": "/agent-roles/{id}", "q": { "exist": ["id"] }, "r": {}, "s": [{ "lit": "agent-roles" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "load" } }, "relations": { "ancestors": [] }, "key$": "agent_role", "name__orig": "agent_role", "Name": "AgentRole", "name_": "agent_role", "name-": "agent-role", "NAME": "AGENT_ROLE", "index$": 1 }, { "active": true, "entity": "agent_role", "key$": "BasicAgentRoleFlow", "kind": "basic", "name": "BasicAgentRoleFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": {}, "m": {}, "o": "list", "s": [], "v": [{ "apply": "ItemExists", "def": { "ref": "agent_role_ref01" } }], "index$": 0 }, { "a": true, "d": {}, "i": { "ref": "agent_role_ref01", "srcdatavar": "agent_role_ref01_data", "suffix": "_dt0" }, "m": { "id": "agent_role01" }, "o": "load", "s": [], "v": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-agent_role_ref01" } }], "index$": 1 }] }, 'AgentRole', { "GET /agent-roles": { "protocol": "http", "parameters": [] }, "GET /agent-roles/{id}": { "protocol": "http", "parameters": [{ "name": "id", "in": "path", "type": "string", "required": true, "x-ref": "#/parameters/id", "index$": 0 }] } }, { strict: LIVE_STRICT, t });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let agent_role_ref01_data = Object.values(setup.data.existing.agent_role)[0];
        // LIST
        const agent_role_ref01_ent = client.AgentRole();
        const agent_role_ref01_match = {};
        const agent_role_ref01_list = (await agent_role_ref01_ent.list(agent_role_ref01_match)).map((e) => e.data());
        // LOAD
        const agent_role_ref01_match_dt0 = {};
        agent_role_ref01_match_dt0.id = agent_role_ref01_data.id;
        const agent_role_ref01_data_dt0 = (await agent_role_ref01_ent.load(agent_role_ref01_match_dt0)).data();
        (0, node_assert_1.default)(agent_role_ref01_data_dt0.id === agent_role_ref01_data.id);
    });
});
// main.kit.test.live.strict is true (the default is true): a live
// request that fails, or a live test missing an input it needs,
// fails the test.
// An account with no record for a test to read skips it either way.
const LIVE_STRICT = true;
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/agent_role/AgentRoleTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.ArtInstituteOfChicagoSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['agent_role01', 'agent_role02', 'agent_role03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'ART_INSTITUTE_OF_CHICAGO_TEST_AGENT_ROLE_ENTID': idmap,
        'ART_INSTITUTE_OF_CHICAGO_TEST_LIVE': 'FALSE',
        'ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN': 'FALSE',
    });
    idmap = env['ART_INSTITUTE_OF_CHICAGO_TEST_AGENT_ROLE_ENTID'];
    const live = 'TRUE' === env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['ART_INSTITUTE_OF_CHICAGO_TEST_AGENT_ROLE_ENTID'];
        idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {};
        if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
            throw new Error('Live ENTID must be a JSON object');
        }
        client = new __1.ArtInstituteOfChicagoSDK(merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            (0, utility_1.liveClientOptions)(),
            {},
            // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
            // last entry is undefined, and basicSetup is normally called with no
            // argument at all - so a bare 'extra' silently discarded the apikey
            // and server values above and handed the SDK undefined. Harmless
            // while there was nothing in that object; not harmless now.
            extra || {},
            { system: { fetch: transport.fetch } }
        ]));
    }
    const setup = {
        idmap,
        env,
        options,
        client,
        struct,
        data: entityData,
        explain: 'TRUE' === env.ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN,
        live,
        transport,
        now: Date.now(),
    };
    return setup;
}
//# sourceMappingURL=AgentRoleEntity.test.js.map