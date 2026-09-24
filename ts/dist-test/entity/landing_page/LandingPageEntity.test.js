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
(0, node_test_1.describe)('LandingPageEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('ART_INSTITUTE_OF_CHICAGO_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.ArtInstituteOfChicagoSDK.test();
        const ent = testsdk.LandingPage();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE;
        for (const op of ['list', 'load']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'landing_page.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "api_link": { "a": true, "h": "Api Link", "n": "api_link", "r": false, "sh": "REST API link for this resource", "t": "`$ANY`", "key$": "api_link", "index$": 0 }, "api_model": { "a": true, "h": "Api Model", "n": "api_model", "r": false, "sh": "REST API resource type or endpoint", "t": "`$ANY`", "key$": "api_model", "index$": 1 }, "copy": { "a": true, "h": "Copy", "n": "copy", "r": false, "sh": "The text of the page", "t": "`$ANY`", "key$": "copy", "index$": 2 }, "id": { "a": true, "h": "Id", "n": "id", "r": false, "sh": "Unique identifier of this resource.", "t": "`$STRING`", "key$": "id", "index$": 3 }, "search_tags": { "a": true, "h": "Search Tags", "n": "search_tags", "r": false, "sh": "Editor-specified list of tags to aid in internal search", "t": "`$ANY`", "key$": "search_tags", "index$": 4 }, "source_updated_at": { "a": true, "h": "Source Updated At", "n": "source_updated_at", "r": false, "sh": "Date and time the resource was updated in the source system", "t": "`$ANY`", "key$": "source_updated_at", "index$": 5 }, "suggest_autocomplete_all": { "a": true, "h": "Suggest Autocomplete All", "n": "suggest_autocomplete_all", "r": false, "sh": "Internal field to power the `/autosuggest` endpoint.", "t": "`$ANY`", "key$": "suggest_autocomplete_all", "index$": 6 }, "suggest_autocomplete_boosted": { "a": true, "h": "Suggest Autocomplete Boosted", "n": "suggest_autocomplete_boosted", "r": false, "sh": "Internal field to power the `/autocomplete` endpoint.", "t": "`$ANY`", "key$": "suggest_autocomplete_boosted", "index$": 7 }, "timestamp": { "a": true, "h": "Timestamp", "n": "timestamp", "r": false, "sh": "Date and time the record was updated in the aggregator search index", "t": "`$ANY`", "key$": "timestamp", "index$": 8 }, "title": { "a": true, "h": "Title", "n": "title", "r": false, "sh": "The name of this resource", "t": "`$STRING`", "key$": "title", "index$": 9 }, "updated_at": { "a": true, "h": "Updated At", "n": "updated_at", "r": false, "sh": "Date and time the record was updated in the aggregator database", "t": "`$ANY`", "key$": "updated_at", "index$": 10 }, "web_url": { "a": true, "h": "Web Url", "n": "web_url", "r": false, "sh": "The URL to this page on our website", "t": "`$ANY`", "key$": "web_url", "index$": 11 } }, "id": { "field": "id", "name": "id" }, "name": "landing_page", "op": { "list": { "input": "data", "name": "list", "points": [{ "a": true, "co": { "id": "GET /landing-pages", "source": "openapi3", "version": 2 }, "g": {}, "k": "http", "m": "GET", "o": "/landing-pages", "q": {}, "r": {}, "s": [{ "lit": "landing-pages" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "list" }, "load": { "input": "data", "name": "load", "points": [{ "a": true, "co": { "id": "GET /landing-pages/{id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "GET", "o": "/landing-pages/{id}", "q": { "exist": ["id"] }, "r": {}, "s": [{ "lit": "landing-pages" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "load" } }, "relations": { "ancestors": [] }, "key$": "landing_page", "name__orig": "landing_page", "Name": "LandingPage", "name_": "landing_page", "name-": "landing-page", "NAME": "LANDING_PAGE", "index$": 21 }, { "active": true, "entity": "landing_page", "key$": "BasicLandingPageFlow", "kind": "basic", "name": "BasicLandingPageFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": {}, "m": {}, "o": "list", "s": [], "v": [{ "apply": "ItemExists", "def": { "ref": "landing_page_ref01" } }], "index$": 0 }, { "a": true, "d": {}, "i": { "ref": "landing_page_ref01", "srcdatavar": "landing_page_ref01_data", "suffix": "_dt0" }, "m": { "id": "landing_page01" }, "o": "load", "s": [], "v": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-landing_page_ref01" } }], "index$": 1 }] }, 'LandingPage', { "GET /landing-pages": { "protocol": "http", "responses": { "200": { "description": "Successful operation", "schema": { "type": "array", "items": { "properties": { "id": { "description": "Unique identifier of this resource. Taken from the source system.", "key$": "id" }, "api_model": { "description": "REST API resource type or endpoint", "key$": "api_model" }, "api_link": { "description": "REST API link for this resource", "key$": "api_link" }, "title": { "description": "The name of this resource", "key$": "title" }, "web_url": { "description": "The URL to this page on our website", "key$": "web_url" }, "copy": { "description": "The text of the page", "key$": "copy" }, "search_tags": { "description": "Editor-specified list of tags to aid in internal search", "key$": "search_tags" }, "suggest_autocomplete_boosted": { "description": "Internal field to power the `/autocomplete` endpoint. Do not use directly.", "key$": "suggest_autocomplete_boosted" }, "suggest_autocomplete_all": { "description": "Internal field to power the `/autosuggest` endpoint. Do not use directly.", "key$": "suggest_autocomplete_all" }, "source_updated_at": { "description": "Date and time the resource was updated in the source system", "key$": "source_updated_at" }, "updated_at": { "description": "Date and time the record was updated in the aggregator database", "key$": "updated_at" }, "timestamp": { "description": "Date and time the record was updated in the aggregator search index", "key$": "timestamp" } }, "type": "object", "x-ref": "#/definitions/LandingPage", "index$": 0 } } }, "default": { "description": "error", "schema": { "required": ["status", "error", "detail"], "properties": { "status": { "type": "integer" }, "error": { "type": "string" }, "detail": { "type": "string" } }, "x-ref": "#/definitions/Error" } } }, "produces": ["application/json"], "parameters": [], "securitySource": "unspecified" }, "GET /landing-pages/{id}": { "protocol": "http", "responses": { "200": { "description": "Successful operation", "schema": { "type": "array", "items": { "properties": { "id": { "description": "Unique identifier of this resource. Taken from the source system.", "key$": "id" }, "api_model": { "description": "REST API resource type or endpoint", "key$": "api_model" }, "api_link": { "description": "REST API link for this resource", "key$": "api_link" }, "title": { "description": "The name of this resource", "key$": "title" }, "web_url": { "description": "The URL to this page on our website", "key$": "web_url" }, "copy": { "description": "The text of the page", "key$": "copy" }, "search_tags": { "description": "Editor-specified list of tags to aid in internal search", "key$": "search_tags" }, "suggest_autocomplete_boosted": { "description": "Internal field to power the `/autocomplete` endpoint. Do not use directly.", "key$": "suggest_autocomplete_boosted" }, "suggest_autocomplete_all": { "description": "Internal field to power the `/autosuggest` endpoint. Do not use directly.", "key$": "suggest_autocomplete_all" }, "source_updated_at": { "description": "Date and time the resource was updated in the source system", "key$": "source_updated_at" }, "updated_at": { "description": "Date and time the record was updated in the aggregator database", "key$": "updated_at" }, "timestamp": { "description": "Date and time the record was updated in the aggregator search index", "key$": "timestamp" } }, "type": "object", "x-ref": "#/definitions/LandingPage", "key$": "items" } } }, "default": { "description": "error", "schema": { "required": ["status", "error", "detail"], "properties": { "status": { "type": "integer" }, "error": { "type": "string" }, "detail": { "type": "string" } }, "x-ref": "#/definitions/Error" } } }, "produces": ["application/json"], "parameters": [{ "name": "id", "in": "path", "type": "string", "required": true, "x-ref": "#/parameters/id", "index$": 0 }], "securitySource": "unspecified" } });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let landing_page_ref01_data = Object.values(setup.data.existing.landing_page)[0];
        // LIST
        const landing_page_ref01_ent = client.LandingPage();
        const landing_page_ref01_match = {};
        const landing_page_ref01_list = (await landing_page_ref01_ent.list(landing_page_ref01_match)).map((e) => e.data());
        // LOAD
        const landing_page_ref01_match_dt0 = {};
        landing_page_ref01_match_dt0.id = landing_page_ref01_data.id;
        const landing_page_ref01_data_dt0 = (await landing_page_ref01_ent.load(landing_page_ref01_match_dt0)).data();
        (0, node_assert_1.default)(landing_page_ref01_data_dt0.id === landing_page_ref01_data.id);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/landing_page/LandingPageTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.ArtInstituteOfChicagoSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['landing_page01', 'landing_page02', 'landing_page03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'ART_INSTITUTE_OF_CHICAGO_TEST_LANDING_PAGE_ENTID': idmap,
        'ART_INSTITUTE_OF_CHICAGO_TEST_LIVE': 'FALSE',
        'ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN': 'FALSE',
    });
    idmap = env['ART_INSTITUTE_OF_CHICAGO_TEST_LANDING_PAGE_ENTID'];
    const live = 'TRUE' === env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['ART_INSTITUTE_OF_CHICAGO_TEST_LANDING_PAGE_ENTID'];
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
//# sourceMappingURL=LandingPageEntity.test.js.map