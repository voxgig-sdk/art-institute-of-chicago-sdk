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
(0, node_test_1.describe)('ExhibitionEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('ART_INSTITUTE_OF_CHICAGO_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.ArtInstituteOfChicagoSDK.test();
        const ent = testsdk.Exhibition();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE;
        for (const op of ['list', 'load']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'exhibition.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [{ "active": true, "name": "aic_end_at", "req": false, "short": "Date the exhibition closed at the Art Institute of Chicago", "type": "`$ANY`", "index$": 0 }, { "active": true, "name": "aic_start_at", "req": false, "short": "Date the exhibition opened at the Art Institute of Chicago", "type": "`$ANY`", "index$": 1 }, { "active": true, "name": "alt_image_ids", "req": false, "short": "Unique identifiers of all non-preferred images of this exhibition.", "type": "`$ANY`", "index$": 2 }, { "active": true, "name": "api_link", "req": false, "short": "REST API link for this resource", "type": "`$ANY`", "index$": 3 }, { "active": true, "name": "api_model", "req": false, "short": "REST API resource type or endpoint", "type": "`$ANY`", "index$": 4 }, { "active": true, "name": "artist_ids", "req": false, "short": "Unique identifiers of the artist agent records representing who was shown in the exhibition", "type": "`$ANY`", "index$": 5 }, { "active": true, "name": "artwork_ids", "req": false, "short": "Unique identifiers of the artworks that were part of the exhibition", "type": "`$ANY`", "index$": 6 }, { "active": true, "name": "artwork_titles", "req": false, "short": "Names of the artworks that were part of the exhibition", "type": "`$ANY`", "index$": 7 }, { "active": true, "name": "document_ids", "req": false, "short": "Unique identifiers of assets that serve as documentation for this exhibition", "type": "`$ANY`", "index$": 8 }, { "active": true, "name": "gallery_id", "req": false, "short": "Unique identifier of the gallery that mainly housed the exhibition", "type": "`$STRING`", "index$": 9 }, { "active": true, "name": "gallery_title", "req": false, "short": "The name of the gallery that mainly housed the exhibition", "type": "`$ANY`", "index$": 10 }, { "active": true, "name": "id", "req": false, "short": "Unique identifier of this resource.", "type": "`$STRING`", "index$": 11 }, { "active": true, "name": "image_id", "req": false, "short": "Unique identifier of the preferred image to use to represent this exhibition", "type": "`$STRING`", "index$": 12 }, { "active": true, "name": "image_url", "req": false, "short": "URL to the hero image from the website", "type": "`$ANY`", "index$": 13 }, { "active": true, "name": "is_featured", "req": false, "short": "Is this exhibition currently featured on our website?", "type": "`$BOOLEAN`", "index$": 14 }, { "active": true, "name": "is_published", "req": false, "short": "Is this exhibition currently published on our website?", "type": "`$BOOLEAN`", "index$": 15 }, { "active": true, "name": "position", "req": false, "short": "Numering position represnting the order in which this exhibition is featured on the website", "type": "`$ANY`", "index$": 16 }, { "active": true, "name": "short_description", "req": false, "short": "Brief explanation of what this exhibition is", "type": "`$ANY`", "index$": 17 }, { "active": true, "name": "site_ids", "req": false, "short": "Unique identifiers of the microsites this exhibition is a part of", "type": "`$ANY`", "index$": 18 }, { "active": true, "name": "source_updated_at", "req": false, "short": "Date and time the resource was updated in the source system", "type": "`$ANY`", "index$": 19 }, { "active": true, "name": "status", "req": false, "short": "Whether the exhibition is open or closed", "type": "`$ANY`", "index$": 20 }, { "active": true, "name": "suggest_autocomplete_all", "req": false, "short": "Internal field to power the `/autosuggest` endpoint.", "type": "`$ANY`", "index$": 21 }, { "active": true, "name": "suggest_autocomplete_boosted", "req": false, "short": "Internal field to power the `/autocomplete` endpoint.", "type": "`$ANY`", "index$": 22 }, { "active": true, "name": "timestamp", "req": false, "short": "Date and time the record was updated in the aggregator search index", "type": "`$ANY`", "index$": 23 }, { "active": true, "name": "title", "req": false, "short": "The name of this resource", "type": "`$STRING`", "index$": 24 }, { "active": true, "name": "updated_at", "req": false, "short": "Date and time the record was updated in the aggregator database", "type": "`$ANY`", "index$": 25 }, { "active": true, "name": "web_url", "req": false, "short": "URL to this exhibition on our website", "type": "`$ANY`", "index$": 26 }], "id": { "field": "id", "name": "id" }, "name": "exhibition", "op": { "list": { "input": "data", "name": "list", "points": [{ "active": true, "args": {}, "contract": { "id": "GET /exhibitions", "json": "{\"parameters\":[],\"produces\":[\"application/json\"],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"Successful operation\",\"schema\":{\"items\":{\"properties\":{\"aic_end_at\":{\"description\":\"Date the exhibition closed at the Art Institute of Chicago\"},\"aic_start_at\":{\"description\":\"Date the exhibition opened at the Art Institute of Chicago\"},\"alt_image_ids\":{\"description\":\"Unique identifiers of all non-preferred images of this exhibition.\"},\"api_link\":{\"description\":\"REST API link for this resource\"},\"api_model\":{\"description\":\"REST API resource type or endpoint\"},\"artist_ids\":{\"description\":\"Unique identifiers of the artist agent records representing who was shown in the exhibition\"},\"artwork_ids\":{\"description\":\"Unique identifiers of the artworks that were part of the exhibition\"},\"artwork_titles\":{\"description\":\"Names of the artworks that were part of the exhibition\"},\"document_ids\":{\"description\":\"Unique identifiers of assets that serve as documentation for this exhibition\"},\"gallery_id\":{\"description\":\"Unique identifier of the gallery that mainly housed the exhibition\"},\"gallery_title\":{\"description\":\"The name of the gallery that mainly housed the exhibition\"},\"id\":{\"description\":\"Unique identifier of this resource. Taken from the source system.\"},\"image_id\":{\"description\":\"Unique identifier of the preferred image to use to represent this exhibition\"},\"image_url\":{\"description\":\"URL to the hero image from the website\"},\"is_featured\":{\"description\":\"Is this exhibition currently featured on our website?\"},\"is_published\":{\"description\":\"Is this exhibition currently published on our website? Only relevant for non-past exhibitions.\"},\"position\":{\"description\":\"Numering position represnting the order in which this exhibition is featured on the website\"},\"short_description\":{\"description\":\"Brief explanation of what this exhibition is\"},\"site_ids\":{\"description\":\"Unique identifiers of the microsites this exhibition is a part of\"},\"source_updated_at\":{\"description\":\"Date and time the resource was updated in the source system\"},\"status\":{\"description\":\"Whether the exhibition is open or closed\"},\"suggest_autocomplete_all\":{\"description\":\"Internal field to power the `/autosuggest` endpoint. Do not use directly.\"},\"suggest_autocomplete_boosted\":{\"description\":\"Internal field to power the `/autocomplete` endpoint. Do not use directly.\"},\"timestamp\":{\"description\":\"Date and time the record was updated in the aggregator search index\"},\"title\":{\"description\":\"The name of this resource\"},\"updated_at\":{\"description\":\"Date and time the record was updated in the aggregator database\"},\"web_url\":{\"description\":\"URL to this exhibition on our website\"}},\"type\":\"object\"},\"type\":\"array\"}},\"default\":{\"description\":\"error\",\"schema\":{\"properties\":{\"detail\":{\"type\":\"string\"},\"error\":{\"type\":\"string\"},\"status\":{\"type\":\"integer\"}},\"required\":[\"status\",\"error\",\"detail\"]}}},\"securitySource\":\"unspecified\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/exhibitions", "segments": [{ "lit": "exhibitions" }], "select": {}, "transform": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "list" }, "load": { "input": "data", "name": "load", "points": [{ "active": true, "args": { "params": [{ "active": true, "kind": "param", "name": "id", "orig": "id", "reqd": true, "type": "`$STRING`", "index$": 0 }] }, "contract": { "id": "GET /exhibitions/{id}", "json": "{\"parameters\":[{\"in\":\"path\",\"name\":\"id\",\"required\":true,\"type\":\"string\"}],\"produces\":[\"application/json\"],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"Successful operation\",\"schema\":{\"items\":{\"properties\":{\"aic_end_at\":{\"description\":\"Date the exhibition closed at the Art Institute of Chicago\"},\"aic_start_at\":{\"description\":\"Date the exhibition opened at the Art Institute of Chicago\"},\"alt_image_ids\":{\"description\":\"Unique identifiers of all non-preferred images of this exhibition.\"},\"api_link\":{\"description\":\"REST API link for this resource\"},\"api_model\":{\"description\":\"REST API resource type or endpoint\"},\"artist_ids\":{\"description\":\"Unique identifiers of the artist agent records representing who was shown in the exhibition\"},\"artwork_ids\":{\"description\":\"Unique identifiers of the artworks that were part of the exhibition\"},\"artwork_titles\":{\"description\":\"Names of the artworks that were part of the exhibition\"},\"document_ids\":{\"description\":\"Unique identifiers of assets that serve as documentation for this exhibition\"},\"gallery_id\":{\"description\":\"Unique identifier of the gallery that mainly housed the exhibition\"},\"gallery_title\":{\"description\":\"The name of the gallery that mainly housed the exhibition\"},\"id\":{\"description\":\"Unique identifier of this resource. Taken from the source system.\"},\"image_id\":{\"description\":\"Unique identifier of the preferred image to use to represent this exhibition\"},\"image_url\":{\"description\":\"URL to the hero image from the website\"},\"is_featured\":{\"description\":\"Is this exhibition currently featured on our website?\"},\"is_published\":{\"description\":\"Is this exhibition currently published on our website? Only relevant for non-past exhibitions.\"},\"position\":{\"description\":\"Numering position represnting the order in which this exhibition is featured on the website\"},\"short_description\":{\"description\":\"Brief explanation of what this exhibition is\"},\"site_ids\":{\"description\":\"Unique identifiers of the microsites this exhibition is a part of\"},\"source_updated_at\":{\"description\":\"Date and time the resource was updated in the source system\"},\"status\":{\"description\":\"Whether the exhibition is open or closed\"},\"suggest_autocomplete_all\":{\"description\":\"Internal field to power the `/autosuggest` endpoint. Do not use directly.\"},\"suggest_autocomplete_boosted\":{\"description\":\"Internal field to power the `/autocomplete` endpoint. Do not use directly.\"},\"timestamp\":{\"description\":\"Date and time the record was updated in the aggregator search index\"},\"title\":{\"description\":\"The name of this resource\"},\"updated_at\":{\"description\":\"Date and time the record was updated in the aggregator database\"},\"web_url\":{\"description\":\"URL to this exhibition on our website\"}},\"type\":\"object\"},\"type\":\"array\"}},\"default\":{\"description\":\"error\",\"schema\":{\"properties\":{\"detail\":{\"type\":\"string\"},\"error\":{\"type\":\"string\"},\"status\":{\"type\":\"integer\"}},\"required\":[\"status\",\"error\",\"detail\"]}}},\"securitySource\":\"unspecified\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/exhibitions/{id}", "segments": [{ "lit": "exhibitions" }, { "var": "id" }], "select": { "exist": ["id"] }, "transform": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "load" } }, "relations": { "ancestors": [] }, "key$": "exhibition", "name__orig": "exhibition", "Name": "Exhibition", "name_": "exhibition", "name-": "exhibition", "NAME": "EXHIBITION", "index$": 15 }, { "active": true, "entity": "exhibition", "key$": "BasicExhibitionFlow", "kind": "basic", "name": "BasicExhibitionFlow", "param": {}, "step": [{ "active": true, "data": {}, "input": {}, "match": {}, "op": "list", "spec": [], "valid": [{ "apply": "ItemExists", "def": { "ref": "exhibition_ref01" } }], "index$": 0 }, { "active": true, "data": {}, "input": { "ref": "exhibition_ref01", "srcdatavar": "exhibition_ref01_data", "suffix": "_dt0" }, "match": { "id": "exhibition01" }, "op": "load", "spec": [], "valid": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-exhibition_ref01" } }], "index$": 1 }] }, 'Exhibition');
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let exhibition_ref01_data = Object.values(setup.data.existing.exhibition)[0];
        // LIST
        const exhibition_ref01_ent = client.Exhibition();
        const exhibition_ref01_match = {};
        const exhibition_ref01_list = (await exhibition_ref01_ent.list(exhibition_ref01_match)).map((e) => e.data());
        // LOAD
        const exhibition_ref01_match_dt0 = {};
        exhibition_ref01_match_dt0.id = exhibition_ref01_data.id;
        const exhibition_ref01_data_dt0 = (await exhibition_ref01_ent.load(exhibition_ref01_match_dt0)).data();
        (0, node_assert_1.default)(exhibition_ref01_data_dt0.id === exhibition_ref01_data.id);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/exhibition/ExhibitionTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.ArtInstituteOfChicagoSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['exhibition01', 'exhibition02', 'exhibition03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'ART_INSTITUTE_OF_CHICAGO_TEST_EXHIBITION_ENTID': idmap,
        'ART_INSTITUTE_OF_CHICAGO_TEST_LIVE': 'FALSE',
        'ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN': 'FALSE',
    });
    idmap = env['ART_INSTITUTE_OF_CHICAGO_TEST_EXHIBITION_ENTID'];
    const live = 'TRUE' === env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['ART_INSTITUTE_OF_CHICAGO_TEST_EXHIBITION_ENTID'];
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
//# sourceMappingURL=ExhibitionEntity.test.js.map