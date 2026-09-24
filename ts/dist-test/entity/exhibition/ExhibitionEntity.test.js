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
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "aic_end_at": { "a": true, "h": "Aic End At", "n": "aic_end_at", "r": false, "sh": "Date the exhibition closed at the Art Institute of Chicago", "t": "`$ANY`", "key$": "aic_end_at", "index$": 0 }, "aic_start_at": { "a": true, "h": "Aic Start At", "n": "aic_start_at", "r": false, "sh": "Date the exhibition opened at the Art Institute of Chicago", "t": "`$ANY`", "key$": "aic_start_at", "index$": 1 }, "alt_image_ids": { "a": true, "h": "Alt Image Ids", "n": "alt_image_ids", "r": false, "sh": "Unique identifiers of all non-preferred images of this exhibition.", "t": "`$ANY`", "key$": "alt_image_ids", "index$": 2 }, "api_link": { "a": true, "h": "Api Link", "n": "api_link", "r": false, "sh": "REST API link for this resource", "t": "`$ANY`", "key$": "api_link", "index$": 3 }, "api_model": { "a": true, "h": "Api Model", "n": "api_model", "r": false, "sh": "REST API resource type or endpoint", "t": "`$ANY`", "key$": "api_model", "index$": 4 }, "artist_ids": { "a": true, "h": "Artist Ids", "n": "artist_ids", "r": false, "sh": "Unique identifiers of the artist agent records representing who was shown in the exhibition", "t": "`$ANY`", "key$": "artist_ids", "index$": 5 }, "artwork_ids": { "a": true, "h": "Artwork Ids", "n": "artwork_ids", "r": false, "sh": "Unique identifiers of the artworks that were part of the exhibition", "t": "`$ANY`", "key$": "artwork_ids", "index$": 6 }, "artwork_titles": { "a": true, "h": "Artwork Titles", "n": "artwork_titles", "r": false, "sh": "Names of the artworks that were part of the exhibition", "t": "`$ANY`", "key$": "artwork_titles", "index$": 7 }, "document_ids": { "a": true, "h": "Document Ids", "n": "document_ids", "r": false, "sh": "Unique identifiers of assets that serve as documentation for this exhibition", "t": "`$ANY`", "key$": "document_ids", "index$": 8 }, "gallery_id": { "a": true, "h": "Gallery Id", "n": "gallery_id", "r": false, "sh": "Unique identifier of the gallery that mainly housed the exhibition", "t": "`$STRING`", "key$": "gallery_id", "index$": 9 }, "gallery_title": { "a": true, "h": "Gallery Title", "n": "gallery_title", "r": false, "sh": "The name of the gallery that mainly housed the exhibition", "t": "`$ANY`", "key$": "gallery_title", "index$": 10 }, "id": { "a": true, "h": "Id", "n": "id", "r": false, "sh": "Unique identifier of this resource.", "t": "`$STRING`", "key$": "id", "index$": 11 }, "image_id": { "a": true, "h": "Image Id", "n": "image_id", "r": false, "sh": "Unique identifier of the preferred image to use to represent this exhibition", "t": "`$STRING`", "key$": "image_id", "index$": 12 }, "image_url": { "a": true, "h": "Image Url", "n": "image_url", "r": false, "sh": "URL to the hero image from the website", "t": "`$ANY`", "key$": "image_url", "index$": 13 }, "is_featured": { "a": true, "h": "Is Featured", "n": "is_featured", "r": false, "sh": "Is this exhibition currently featured on our website?", "t": "`$BOOLEAN`", "key$": "is_featured", "index$": 14 }, "is_published": { "a": true, "h": "Is Published", "n": "is_published", "r": false, "sh": "Is this exhibition currently published on our website?", "t": "`$BOOLEAN`", "key$": "is_published", "index$": 15 }, "position": { "a": true, "h": "Position", "n": "position", "r": false, "sh": "Numering position represnting the order in which this exhibition is featured on the website", "t": "`$ANY`", "key$": "position", "index$": 16 }, "short_description": { "a": true, "h": "Short Description", "n": "short_description", "r": false, "sh": "Brief explanation of what this exhibition is", "t": "`$ANY`", "key$": "short_description", "index$": 17 }, "site_ids": { "a": true, "h": "Site Ids", "n": "site_ids", "r": false, "sh": "Unique identifiers of the microsites this exhibition is a part of", "t": "`$ANY`", "key$": "site_ids", "index$": 18 }, "source_updated_at": { "a": true, "h": "Source Updated At", "n": "source_updated_at", "r": false, "sh": "Date and time the resource was updated in the source system", "t": "`$ANY`", "key$": "source_updated_at", "index$": 19 }, "status": { "a": true, "h": "Status", "n": "status", "r": false, "sh": "Whether the exhibition is open or closed", "t": "`$ANY`", "key$": "status", "index$": 20 }, "suggest_autocomplete_all": { "a": true, "h": "Suggest Autocomplete All", "n": "suggest_autocomplete_all", "r": false, "sh": "Internal field to power the `/autosuggest` endpoint.", "t": "`$ANY`", "key$": "suggest_autocomplete_all", "index$": 21 }, "suggest_autocomplete_boosted": { "a": true, "h": "Suggest Autocomplete Boosted", "n": "suggest_autocomplete_boosted", "r": false, "sh": "Internal field to power the `/autocomplete` endpoint.", "t": "`$ANY`", "key$": "suggest_autocomplete_boosted", "index$": 22 }, "timestamp": { "a": true, "h": "Timestamp", "n": "timestamp", "r": false, "sh": "Date and time the record was updated in the aggregator search index", "t": "`$ANY`", "key$": "timestamp", "index$": 23 }, "title": { "a": true, "h": "Title", "n": "title", "r": false, "sh": "The name of this resource", "t": "`$STRING`", "key$": "title", "index$": 24 }, "updated_at": { "a": true, "h": "Updated At", "n": "updated_at", "r": false, "sh": "Date and time the record was updated in the aggregator database", "t": "`$ANY`", "key$": "updated_at", "index$": 25 }, "web_url": { "a": true, "h": "Web Url", "n": "web_url", "r": false, "sh": "URL to this exhibition on our website", "t": "`$ANY`", "key$": "web_url", "index$": 26 } }, "id": { "field": "id", "name": "id" }, "name": "exhibition", "op": { "list": { "input": "data", "name": "list", "points": [{ "a": true, "co": { "id": "GET /exhibitions", "source": "openapi3", "version": 2 }, "g": {}, "k": "http", "m": "GET", "o": "/exhibitions", "q": {}, "r": {}, "s": [{ "lit": "exhibitions" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "list" }, "load": { "input": "data", "name": "load", "points": [{ "a": true, "co": { "id": "GET /exhibitions/{id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "GET", "o": "/exhibitions/{id}", "q": { "exist": ["id"] }, "r": {}, "s": [{ "lit": "exhibitions" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "load" } }, "relations": { "ancestors": [] }, "key$": "exhibition", "name__orig": "exhibition", "Name": "Exhibition", "name_": "exhibition", "name-": "exhibition", "NAME": "EXHIBITION", "index$": 15 }, { "active": true, "entity": "exhibition", "key$": "BasicExhibitionFlow", "kind": "basic", "name": "BasicExhibitionFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": {}, "m": {}, "o": "list", "s": [], "v": [{ "apply": "ItemExists", "def": { "ref": "exhibition_ref01" } }], "index$": 0 }, { "a": true, "d": {}, "i": { "ref": "exhibition_ref01", "srcdatavar": "exhibition_ref01_data", "suffix": "_dt0" }, "m": { "id": "exhibition01" }, "o": "load", "s": [], "v": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-exhibition_ref01" } }], "index$": 1 }] }, 'Exhibition', { "GET /exhibitions": { "protocol": "http", "responses": { "200": { "description": "Successful operation", "schema": { "type": "array", "items": { "properties": { "id": { "description": "Unique identifier of this resource. Taken from the source system.", "key$": "id" }, "api_model": { "description": "REST API resource type or endpoint", "key$": "api_model" }, "api_link": { "description": "REST API link for this resource", "key$": "api_link" }, "title": { "description": "The name of this resource", "key$": "title" }, "is_featured": { "description": "Is this exhibition currently featured on our website?", "key$": "is_featured" }, "position": { "description": "Numering position represnting the order in which this exhibition is featured on the website", "key$": "position" }, "is_published": { "description": "Is this exhibition currently published on our website? Only relevant for non-past exhibitions.", "key$": "is_published" }, "short_description": { "description": "Brief explanation of what this exhibition is", "key$": "short_description" }, "web_url": { "description": "URL to this exhibition on our website", "key$": "web_url" }, "image_url": { "description": "URL to the hero image from the website", "key$": "image_url" }, "status": { "description": "Whether the exhibition is open or closed", "key$": "status" }, "aic_start_at": { "description": "Date the exhibition opened at the Art Institute of Chicago", "key$": "aic_start_at" }, "aic_end_at": { "description": "Date the exhibition closed at the Art Institute of Chicago", "key$": "aic_end_at" }, "gallery_id": { "description": "Unique identifier of the gallery that mainly housed the exhibition", "key$": "gallery_id" }, "gallery_title": { "description": "The name of the gallery that mainly housed the exhibition", "key$": "gallery_title" }, "artwork_ids": { "description": "Unique identifiers of the artworks that were part of the exhibition", "key$": "artwork_ids" }, "artwork_titles": { "description": "Names of the artworks that were part of the exhibition", "key$": "artwork_titles" }, "artist_ids": { "description": "Unique identifiers of the artist agent records representing who was shown in the exhibition", "key$": "artist_ids" }, "site_ids": { "description": "Unique identifiers of the microsites this exhibition is a part of", "key$": "site_ids" }, "image_id": { "description": "Unique identifier of the preferred image to use to represent this exhibition", "key$": "image_id" }, "alt_image_ids": { "description": "Unique identifiers of all non-preferred images of this exhibition.", "key$": "alt_image_ids" }, "document_ids": { "description": "Unique identifiers of assets that serve as documentation for this exhibition", "key$": "document_ids" }, "suggest_autocomplete_boosted": { "description": "Internal field to power the `/autocomplete` endpoint. Do not use directly.", "key$": "suggest_autocomplete_boosted" }, "suggest_autocomplete_all": { "description": "Internal field to power the `/autosuggest` endpoint. Do not use directly.", "key$": "suggest_autocomplete_all" }, "source_updated_at": { "description": "Date and time the resource was updated in the source system", "key$": "source_updated_at" }, "updated_at": { "description": "Date and time the record was updated in the aggregator database", "key$": "updated_at" }, "timestamp": { "description": "Date and time the record was updated in the aggregator search index", "key$": "timestamp" } }, "type": "object", "x-ref": "#/definitions/Exhibition", "index$": 0 } } }, "default": { "description": "error", "schema": { "required": ["status", "error", "detail"], "properties": { "status": { "type": "integer" }, "error": { "type": "string" }, "detail": { "type": "string" } }, "x-ref": "#/definitions/Error" } } }, "produces": ["application/json"], "parameters": [], "securitySource": "unspecified" }, "GET /exhibitions/{id}": { "protocol": "http", "responses": { "200": { "description": "Successful operation", "schema": { "type": "array", "items": { "properties": { "id": { "description": "Unique identifier of this resource. Taken from the source system.", "key$": "id" }, "api_model": { "description": "REST API resource type or endpoint", "key$": "api_model" }, "api_link": { "description": "REST API link for this resource", "key$": "api_link" }, "title": { "description": "The name of this resource", "key$": "title" }, "is_featured": { "description": "Is this exhibition currently featured on our website?", "key$": "is_featured" }, "position": { "description": "Numering position represnting the order in which this exhibition is featured on the website", "key$": "position" }, "is_published": { "description": "Is this exhibition currently published on our website? Only relevant for non-past exhibitions.", "key$": "is_published" }, "short_description": { "description": "Brief explanation of what this exhibition is", "key$": "short_description" }, "web_url": { "description": "URL to this exhibition on our website", "key$": "web_url" }, "image_url": { "description": "URL to the hero image from the website", "key$": "image_url" }, "status": { "description": "Whether the exhibition is open or closed", "key$": "status" }, "aic_start_at": { "description": "Date the exhibition opened at the Art Institute of Chicago", "key$": "aic_start_at" }, "aic_end_at": { "description": "Date the exhibition closed at the Art Institute of Chicago", "key$": "aic_end_at" }, "gallery_id": { "description": "Unique identifier of the gallery that mainly housed the exhibition", "key$": "gallery_id" }, "gallery_title": { "description": "The name of the gallery that mainly housed the exhibition", "key$": "gallery_title" }, "artwork_ids": { "description": "Unique identifiers of the artworks that were part of the exhibition", "key$": "artwork_ids" }, "artwork_titles": { "description": "Names of the artworks that were part of the exhibition", "key$": "artwork_titles" }, "artist_ids": { "description": "Unique identifiers of the artist agent records representing who was shown in the exhibition", "key$": "artist_ids" }, "site_ids": { "description": "Unique identifiers of the microsites this exhibition is a part of", "key$": "site_ids" }, "image_id": { "description": "Unique identifier of the preferred image to use to represent this exhibition", "key$": "image_id" }, "alt_image_ids": { "description": "Unique identifiers of all non-preferred images of this exhibition.", "key$": "alt_image_ids" }, "document_ids": { "description": "Unique identifiers of assets that serve as documentation for this exhibition", "key$": "document_ids" }, "suggest_autocomplete_boosted": { "description": "Internal field to power the `/autocomplete` endpoint. Do not use directly.", "key$": "suggest_autocomplete_boosted" }, "suggest_autocomplete_all": { "description": "Internal field to power the `/autosuggest` endpoint. Do not use directly.", "key$": "suggest_autocomplete_all" }, "source_updated_at": { "description": "Date and time the resource was updated in the source system", "key$": "source_updated_at" }, "updated_at": { "description": "Date and time the record was updated in the aggregator database", "key$": "updated_at" }, "timestamp": { "description": "Date and time the record was updated in the aggregator search index", "key$": "timestamp" } }, "type": "object", "x-ref": "#/definitions/Exhibition", "key$": "items" } } }, "default": { "description": "error", "schema": { "required": ["status", "error", "detail"], "properties": { "status": { "type": "integer" }, "error": { "type": "string" }, "detail": { "type": "string" } }, "x-ref": "#/definitions/Error" } } }, "produces": ["application/json"], "parameters": [{ "name": "id", "in": "path", "type": "string", "required": true, "x-ref": "#/parameters/id", "index$": 0 }], "securitySource": "unspecified" } });
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