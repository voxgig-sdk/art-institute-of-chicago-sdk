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
// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
(0, utility_1.loadEnvLocal)(__dirname + '/../../../.env.local');
(0, node_test_1.describe)('HourEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('ART_INSTITUTE_OF_CHICAGO_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.ArtInstituteOfChicagoSDK.test();
        const ent = testsdk.Hour();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE;
        for (const op of ['list', 'load']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'hour.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [{ "active": true, "name": "additional_text", "req": false, "short": "Additional information about the hours", "type": "`$ANY`", "index$": 0 }, { "active": true, "name": "api_link", "req": false, "short": "REST API link for this resource", "type": "`$ANY`", "index$": 1 }, { "active": true, "name": "api_model", "req": false, "short": "REST API resource type or endpoint", "type": "`$ANY`", "index$": 2 }, { "active": true, "name": "friday_is_closed", "req": false, "short": "Whether the museum is closed on Fridays", "type": "`$ANY`", "index$": 3 }, { "active": true, "name": "friday_member_close", "req": false, "short": "The time member hours ends on Fridays", "type": "`$ANY`", "index$": 4 }, { "active": true, "name": "friday_member_open", "req": false, "short": "The time member hours starts on Fridays", "type": "`$ANY`", "index$": 5 }, { "active": true, "name": "friday_public_close", "req": false, "short": "The time public hours ends on Fridays", "type": "`$ANY`", "index$": 6 }, { "active": true, "name": "friday_public_open", "req": false, "short": "The time public hours starts on Fridays", "type": "`$ANY`", "index$": 7 }, { "active": true, "name": "id", "req": false, "short": "Unique identifier of this resource.", "type": "`$STRING`", "index$": 8 }, { "active": true, "name": "monday_is_closed", "req": false, "short": "Whether the museum is closed on Mondays", "type": "`$ANY`", "index$": 9 }, { "active": true, "name": "monday_member_close", "req": false, "short": "The time member hours ends on Mondays", "type": "`$ANY`", "index$": 10 }, { "active": true, "name": "monday_member_open", "req": false, "short": "The time member hours starts on Mondays", "type": "`$ANY`", "index$": 11 }, { "active": true, "name": "monday_public_close", "req": false, "short": "The time public hours ends on Mondays", "type": "`$ANY`", "index$": 12 }, { "active": true, "name": "monday_public_open", "req": false, "short": "The time public hours starts on Mondays", "type": "`$ANY`", "index$": 13 }, { "active": true, "name": "saturday_is_closed", "req": false, "short": "Whether the museum is closed on Saturdays", "type": "`$ANY`", "index$": 14 }, { "active": true, "name": "saturday_member_close", "req": false, "short": "The time member hours ends on Saturdays", "type": "`$ANY`", "index$": 15 }, { "active": true, "name": "saturday_member_open", "req": false, "short": "The time member hours starts on Saturdays", "type": "`$ANY`", "index$": 16 }, { "active": true, "name": "saturday_public_close", "req": false, "short": "The time public hours ends on Saturdays", "type": "`$ANY`", "index$": 17 }, { "active": true, "name": "saturday_public_open", "req": false, "short": "The time public hours starts on Saturdays", "type": "`$ANY`", "index$": 18 }, { "active": true, "name": "source_updated_at", "req": false, "short": "Date and time the resource was updated in the source system", "type": "`$ANY`", "index$": 19 }, { "active": true, "name": "suggest_autocomplete_all", "req": false, "short": "Internal field to power the `/autosuggest` endpoint.", "type": "`$ANY`", "index$": 20 }, { "active": true, "name": "suggest_autocomplete_boosted", "req": false, "short": "Internal field to power the `/autocomplete` endpoint.", "type": "`$ANY`", "index$": 21 }, { "active": true, "name": "summary", "req": false, "short": "Readable summary of the hours", "type": "`$ANY`", "index$": 22 }, { "active": true, "name": "sunday_is_closed", "req": false, "short": "Whether the museum is closed on Sundays", "type": "`$ANY`", "index$": 23 }, { "active": true, "name": "sunday_member_close", "req": false, "short": "The time member hours ends on Sundays", "type": "`$ANY`", "index$": 24 }, { "active": true, "name": "sunday_member_open", "req": false, "short": "The time member hours starts on Sundays", "type": "`$ANY`", "index$": 25 }, { "active": true, "name": "sunday_public_close", "req": false, "short": "The time public hours ends on Sundays", "type": "`$ANY`", "index$": 26 }, { "active": true, "name": "sunday_public_open", "req": false, "short": "The time public hours starts on Sundays", "type": "`$ANY`", "index$": 27 }, { "active": true, "name": "thursday_is_closed", "req": false, "short": "Whether the museum is closed on Thursdays", "type": "`$ANY`", "index$": 28 }, { "active": true, "name": "thursday_member_close", "req": false, "short": "The time member hours ends on Thursdays", "type": "`$ANY`", "index$": 29 }, { "active": true, "name": "thursday_member_open", "req": false, "short": "The time member hours starts on Thursdays", "type": "`$ANY`", "index$": 30 }, { "active": true, "name": "thursday_public_close", "req": false, "short": "The time public hours ends on Thursdays", "type": "`$ANY`", "index$": 31 }, { "active": true, "name": "thursday_public_open", "req": false, "short": "The time public hours starts on Thursdays", "type": "`$ANY`", "index$": 32 }, { "active": true, "name": "timestamp", "req": false, "short": "Date and time the record was updated in the aggregator search index", "type": "`$ANY`", "index$": 33 }, { "active": true, "name": "title", "req": false, "short": "The name of this resource", "type": "`$STRING`", "index$": 34 }, { "active": true, "name": "tuesday_is_closed", "req": false, "short": "Whether the museum is closed on Tuesdays", "type": "`$ANY`", "index$": 35 }, { "active": true, "name": "tuesday_member_close", "req": false, "short": "The time member hours ends on Tuesdays", "type": "`$ANY`", "index$": 36 }, { "active": true, "name": "tuesday_member_open", "req": false, "short": "The time member hours starts on Tuesdays", "type": "`$ANY`", "index$": 37 }, { "active": true, "name": "tuesday_public_close", "req": false, "short": "The time public hours ends on Tuesdays", "type": "`$ANY`", "index$": 38 }, { "active": true, "name": "tuesday_public_open", "req": false, "short": "The time public hours starts on Tuesdays", "type": "`$ANY`", "index$": 39 }, { "active": true, "name": "updated_at", "req": false, "short": "Date and time the record was updated in the aggregator database", "type": "`$ANY`", "index$": 40 }, { "active": true, "name": "wednesday_is_closed", "req": false, "short": "Whether the museum is closed on Wednesdays", "type": "`$ANY`", "index$": 41 }, { "active": true, "name": "wednesday_member_close", "req": false, "short": "The time member hours ends on Wednesdays", "type": "`$ANY`", "index$": 42 }, { "active": true, "name": "wednesday_member_open", "req": false, "short": "The time member hours starts on Wednesdays", "type": "`$ANY`", "index$": 43 }, { "active": true, "name": "wednesday_public_close", "req": false, "short": "The time public hours ends on Wednesdays", "type": "`$ANY`", "index$": 44 }, { "active": true, "name": "wednesday_public_open", "req": false, "short": "The time public hours starts on Wednesdays", "type": "`$ANY`", "index$": 45 }], "id": { "field": "id", "name": "id" }, "name": "hour", "op": { "list": { "input": "data", "name": "list", "points": [{ "active": true, "args": {}, "contract": { "id": "GET /hours", "json": "{\"parameters\":[],\"produces\":[\"application/json\"],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"Successful operation\",\"schema\":{\"items\":{\"properties\":{\"additional_text\":{\"description\":\"Additional information about the hours\"},\"api_link\":{\"description\":\"REST API link for this resource\"},\"api_model\":{\"description\":\"REST API resource type or endpoint\"},\"friday_is_closed\":{\"description\":\"Whether the museum is closed on Fridays\"},\"friday_member_close\":{\"description\":\"The time member hours ends on Fridays\"},\"friday_member_open\":{\"description\":\"The time member hours starts on Fridays\"},\"friday_public_close\":{\"description\":\"The time public hours ends on Fridays\"},\"friday_public_open\":{\"description\":\"The time public hours starts on Fridays\"},\"id\":{\"description\":\"Unique identifier of this resource. Taken from the source system.\"},\"monday_is_closed\":{\"description\":\"Whether the museum is closed on Mondays\"},\"monday_member_close\":{\"description\":\"The time member hours ends on Mondays\"},\"monday_member_open\":{\"description\":\"The time member hours starts on Mondays\"},\"monday_public_close\":{\"description\":\"The time public hours ends on Mondays\"},\"monday_public_open\":{\"description\":\"The time public hours starts on Mondays\"},\"saturday_is_closed\":{\"description\":\"Whether the museum is closed on Saturdays\"},\"saturday_member_close\":{\"description\":\"The time member hours ends on Saturdays\"},\"saturday_member_open\":{\"description\":\"The time member hours starts on Saturdays\"},\"saturday_public_close\":{\"description\":\"The time public hours ends on Saturdays\"},\"saturday_public_open\":{\"description\":\"The time public hours starts on Saturdays\"},\"source_updated_at\":{\"description\":\"Date and time the resource was updated in the source system\"},\"suggest_autocomplete_all\":{\"description\":\"Internal field to power the `/autosuggest` endpoint. Do not use directly.\"},\"suggest_autocomplete_boosted\":{\"description\":\"Internal field to power the `/autocomplete` endpoint. Do not use directly.\"},\"summary\":{\"description\":\"Readable summary of the hours\"},\"sunday_is_closed\":{\"description\":\"Whether the museum is closed on Sundays\"},\"sunday_member_close\":{\"description\":\"The time member hours ends on Sundays\"},\"sunday_member_open\":{\"description\":\"The time member hours starts on Sundays\"},\"sunday_public_close\":{\"description\":\"The time public hours ends on Sundays\"},\"sunday_public_open\":{\"description\":\"The time public hours starts on Sundays\"},\"thursday_is_closed\":{\"description\":\"Whether the museum is closed on Thursdays\"},\"thursday_member_close\":{\"description\":\"The time member hours ends on Thursdays\"},\"thursday_member_open\":{\"description\":\"The time member hours starts on Thursdays\"},\"thursday_public_close\":{\"description\":\"The time public hours ends on Thursdays\"},\"thursday_public_open\":{\"description\":\"The time public hours starts on Thursdays\"},\"timestamp\":{\"description\":\"Date and time the record was updated in the aggregator search index\"},\"title\":{\"description\":\"The name of this resource\"},\"tuesday_is_closed\":{\"description\":\"Whether the museum is closed on Tuesdays\"},\"tuesday_member_close\":{\"description\":\"The time member hours ends on Tuesdays\"},\"tuesday_member_open\":{\"description\":\"The time member hours starts on Tuesdays\"},\"tuesday_public_close\":{\"description\":\"The time public hours ends on Tuesdays\"},\"tuesday_public_open\":{\"description\":\"The time public hours starts on Tuesdays\"},\"updated_at\":{\"description\":\"Date and time the record was updated in the aggregator database\"},\"wednesday_is_closed\":{\"description\":\"Whether the museum is closed on Wednesdays\"},\"wednesday_member_close\":{\"description\":\"The time member hours ends on Wednesdays\"},\"wednesday_member_open\":{\"description\":\"The time member hours starts on Wednesdays\"},\"wednesday_public_close\":{\"description\":\"The time public hours ends on Wednesdays\"},\"wednesday_public_open\":{\"description\":\"The time public hours starts on Wednesdays\"}},\"type\":\"object\"},\"type\":\"array\"}},\"default\":{\"description\":\"error\",\"schema\":{\"properties\":{\"detail\":{\"type\":\"string\"},\"error\":{\"type\":\"string\"},\"status\":{\"type\":\"integer\"}},\"required\":[\"status\",\"error\",\"detail\"]}}},\"securitySource\":\"unspecified\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/hours", "segments": [{ "lit": "hours" }], "select": {}, "transform": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "list" }, "load": { "input": "data", "name": "load", "points": [{ "active": true, "args": { "params": [{ "active": true, "kind": "param", "name": "id", "orig": "id", "reqd": true, "type": "`$STRING`", "index$": 0 }] }, "contract": { "id": "GET /hours/{id}", "json": "{\"parameters\":[{\"in\":\"path\",\"name\":\"id\",\"required\":true,\"type\":\"string\"}],\"produces\":[\"application/json\"],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"Successful operation\",\"schema\":{\"items\":{\"properties\":{\"additional_text\":{\"description\":\"Additional information about the hours\"},\"api_link\":{\"description\":\"REST API link for this resource\"},\"api_model\":{\"description\":\"REST API resource type or endpoint\"},\"friday_is_closed\":{\"description\":\"Whether the museum is closed on Fridays\"},\"friday_member_close\":{\"description\":\"The time member hours ends on Fridays\"},\"friday_member_open\":{\"description\":\"The time member hours starts on Fridays\"},\"friday_public_close\":{\"description\":\"The time public hours ends on Fridays\"},\"friday_public_open\":{\"description\":\"The time public hours starts on Fridays\"},\"id\":{\"description\":\"Unique identifier of this resource. Taken from the source system.\"},\"monday_is_closed\":{\"description\":\"Whether the museum is closed on Mondays\"},\"monday_member_close\":{\"description\":\"The time member hours ends on Mondays\"},\"monday_member_open\":{\"description\":\"The time member hours starts on Mondays\"},\"monday_public_close\":{\"description\":\"The time public hours ends on Mondays\"},\"monday_public_open\":{\"description\":\"The time public hours starts on Mondays\"},\"saturday_is_closed\":{\"description\":\"Whether the museum is closed on Saturdays\"},\"saturday_member_close\":{\"description\":\"The time member hours ends on Saturdays\"},\"saturday_member_open\":{\"description\":\"The time member hours starts on Saturdays\"},\"saturday_public_close\":{\"description\":\"The time public hours ends on Saturdays\"},\"saturday_public_open\":{\"description\":\"The time public hours starts on Saturdays\"},\"source_updated_at\":{\"description\":\"Date and time the resource was updated in the source system\"},\"suggest_autocomplete_all\":{\"description\":\"Internal field to power the `/autosuggest` endpoint. Do not use directly.\"},\"suggest_autocomplete_boosted\":{\"description\":\"Internal field to power the `/autocomplete` endpoint. Do not use directly.\"},\"summary\":{\"description\":\"Readable summary of the hours\"},\"sunday_is_closed\":{\"description\":\"Whether the museum is closed on Sundays\"},\"sunday_member_close\":{\"description\":\"The time member hours ends on Sundays\"},\"sunday_member_open\":{\"description\":\"The time member hours starts on Sundays\"},\"sunday_public_close\":{\"description\":\"The time public hours ends on Sundays\"},\"sunday_public_open\":{\"description\":\"The time public hours starts on Sundays\"},\"thursday_is_closed\":{\"description\":\"Whether the museum is closed on Thursdays\"},\"thursday_member_close\":{\"description\":\"The time member hours ends on Thursdays\"},\"thursday_member_open\":{\"description\":\"The time member hours starts on Thursdays\"},\"thursday_public_close\":{\"description\":\"The time public hours ends on Thursdays\"},\"thursday_public_open\":{\"description\":\"The time public hours starts on Thursdays\"},\"timestamp\":{\"description\":\"Date and time the record was updated in the aggregator search index\"},\"title\":{\"description\":\"The name of this resource\"},\"tuesday_is_closed\":{\"description\":\"Whether the museum is closed on Tuesdays\"},\"tuesday_member_close\":{\"description\":\"The time member hours ends on Tuesdays\"},\"tuesday_member_open\":{\"description\":\"The time member hours starts on Tuesdays\"},\"tuesday_public_close\":{\"description\":\"The time public hours ends on Tuesdays\"},\"tuesday_public_open\":{\"description\":\"The time public hours starts on Tuesdays\"},\"updated_at\":{\"description\":\"Date and time the record was updated in the aggregator database\"},\"wednesday_is_closed\":{\"description\":\"Whether the museum is closed on Wednesdays\"},\"wednesday_member_close\":{\"description\":\"The time member hours ends on Wednesdays\"},\"wednesday_member_open\":{\"description\":\"The time member hours starts on Wednesdays\"},\"wednesday_public_close\":{\"description\":\"The time public hours ends on Wednesdays\"},\"wednesday_public_open\":{\"description\":\"The time public hours starts on Wednesdays\"}},\"type\":\"object\"},\"type\":\"array\"}},\"default\":{\"description\":\"error\",\"schema\":{\"properties\":{\"detail\":{\"type\":\"string\"},\"error\":{\"type\":\"string\"},\"status\":{\"type\":\"integer\"}},\"required\":[\"status\",\"error\",\"detail\"]}}},\"securitySource\":\"unspecified\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/hours/{id}", "segments": [{ "lit": "hours" }, { "var": "id" }], "select": { "exist": ["id"] }, "transform": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "load" } }, "relations": { "ancestors": [] }, "key$": "hour", "name__orig": "hour", "Name": "Hour", "name_": "hour", "name-": "hour", "NAME": "HOUR", "index$": 19 }, { "active": true, "entity": "hour", "key$": "BasicHourFlow", "kind": "basic", "name": "BasicHourFlow", "param": {}, "step": [{ "active": true, "data": {}, "input": {}, "match": {}, "op": "list", "spec": [], "valid": [{ "apply": "ItemExists", "def": { "ref": "hour_ref01" } }], "index$": 0 }, { "active": true, "data": {}, "input": { "ref": "hour_ref01", "srcdatavar": "hour_ref01_data", "suffix": "_dt0" }, "match": { "id": "hour01" }, "op": "load", "spec": [], "valid": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-hour_ref01" } }], "index$": 1 }] }, 'Hour');
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let hour_ref01_data = Object.values(setup.data.existing.hour)[0];
        // LIST
        const hour_ref01_ent = client.Hour();
        const hour_ref01_match = {};
        const hour_ref01_list = (await hour_ref01_ent.list(hour_ref01_match)).map((e) => e.data());
        // LOAD
        const hour_ref01_match_dt0 = {};
        hour_ref01_match_dt0.id = hour_ref01_data.id;
        const hour_ref01_data_dt0 = (await hour_ref01_ent.load(hour_ref01_match_dt0)).data();
        (0, node_assert_1.default)(hour_ref01_data_dt0.id === hour_ref01_data.id);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/hour/HourTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.ArtInstituteOfChicagoSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['hour01', 'hour02', 'hour03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'ART_INSTITUTE_OF_CHICAGO_TEST_HOUR_ENTID': idmap,
        'ART_INSTITUTE_OF_CHICAGO_TEST_LIVE': 'FALSE',
        'ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN': 'FALSE',
    });
    idmap = env['ART_INSTITUTE_OF_CHICAGO_TEST_HOUR_ENTID'];
    const live = 'TRUE' === env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['ART_INSTITUTE_OF_CHICAGO_TEST_HOUR_ENTID'];
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
//# sourceMappingURL=HourEntity.test.js.map