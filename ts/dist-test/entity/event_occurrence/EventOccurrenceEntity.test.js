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
(0, node_test_1.describe)('EventOccurrenceEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('ART_INSTITUTE_OF_CHICAGO_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.ArtInstituteOfChicagoSDK.test();
        const ent = testsdk.EventOccurrence();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE;
        for (const op of ['list', 'load']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'event_occurrence.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "api_link": { "a": true, "h": "Api Link", "n": "api_link", "r": false, "sh": "REST API link for this resource", "t": "`$ANY`", "key$": "api_link", "index$": 0 }, "api_model": { "a": true, "h": "Api Model", "n": "api_model", "r": false, "sh": "REST API resource type or endpoint", "t": "`$ANY`", "key$": "api_model", "index$": 1 }, "button_caption": { "a": true, "h": "Button Caption", "n": "button_caption", "r": false, "sh": "Additional text below the ticket/registration button", "t": "`$ANY`", "key$": "button_caption", "index$": 2 }, "button_text": { "a": true, "h": "Button Text", "n": "button_text", "r": false, "sh": "The text used on the ticket/registration button", "t": "`$ANY`", "key$": "button_text", "index$": 3 }, "button_url": { "a": true, "h": "Button Url", "n": "button_url", "r": false, "sh": "The URL to the sales site or an RSVP link for this event", "t": "`$ANY`", "key$": "button_url", "index$": 4 }, "description": { "a": true, "h": "Description", "n": "description", "r": false, "sh": "Description of the event", "t": "`$STRING`", "key$": "description", "index$": 5 }, "end_at": { "a": true, "h": "End At", "n": "end_at", "r": false, "sh": "The date the event occurrence ends", "t": "`$ANY`", "key$": "end_at", "index$": 6 }, "event_id": { "a": true, "h": "Event Id", "n": "event_id", "r": false, "sh": "Identifier of the master event of which this is an occurrence", "t": "`$STRING`", "key$": "event_id", "index$": 7 }, "id": { "a": true, "h": "Id", "n": "id", "r": false, "sh": "Unique identifier of this resource.", "t": "`$STRING`", "key$": "id", "index$": 8 }, "image_url": { "a": true, "h": "Image Url", "n": "image_url", "r": false, "sh": "The URL of an image representing this page", "t": "`$ANY`", "key$": "image_url", "index$": 9 }, "is_private": { "a": true, "h": "Is Private", "n": "is_private", "r": false, "sh": "Whether the event is private.", "t": "`$BOOLEAN`", "key$": "is_private", "index$": 10 }, "is_sales_button_hidden": { "a": true, "h": "Is Sales Button Hidden", "n": "is_sales_button_hidden", "r": false, "sh": "Whether the buy tickets button is hidden on the website event page", "t": "`$BOOLEAN`", "key$": "is_sales_button_hidden", "index$": 11 }, "is_ticketed": { "a": true, "h": "Is Ticketed", "n": "is_ticketed", "r": false, "sh": "Whether a ticket is required to attend the event", "t": "`$BOOLEAN`", "key$": "is_ticketed", "index$": 12 }, "location": { "a": true, "h": "Location", "n": "location", "r": false, "sh": "Where the event takes place", "t": "`$ANY`", "key$": "location", "index$": 13 }, "off_sale_at": { "a": true, "h": "Off Sale At", "n": "off_sale_at", "r": false, "sh": "Date and time the event goes off sale", "t": "`$ANY`", "key$": "off_sale_at", "index$": 14 }, "on_sale_at": { "a": true, "h": "On Sale At", "n": "on_sale_at", "r": false, "sh": "Date and time the event goes on sale", "t": "`$ANY`", "key$": "on_sale_at", "index$": 15 }, "short_description": { "a": true, "h": "Short Description", "n": "short_description", "r": false, "sh": "Brief description of the event", "t": "`$ANY`", "key$": "short_description", "index$": 16 }, "source_updated_at": { "a": true, "h": "Source Updated At", "n": "source_updated_at", "r": false, "sh": "Date and time the resource was updated in the source system", "t": "`$ANY`", "key$": "source_updated_at", "index$": 17 }, "start_at": { "a": true, "h": "Start At", "n": "start_at", "r": false, "sh": "The date the event occurrence begins", "t": "`$ANY`", "key$": "start_at", "index$": 18 }, "suggest_autocomplete_all": { "a": true, "h": "Suggest Autocomplete All", "n": "suggest_autocomplete_all", "r": false, "sh": "Internal field to power the `/autosuggest` endpoint.", "t": "`$ANY`", "key$": "suggest_autocomplete_all", "index$": 19 }, "suggest_autocomplete_boosted": { "a": true, "h": "Suggest Autocomplete Boosted", "n": "suggest_autocomplete_boosted", "r": false, "sh": "Internal field to power the `/autocomplete` endpoint.", "t": "`$ANY`", "key$": "suggest_autocomplete_boosted", "index$": 20 }, "timestamp": { "a": true, "h": "Timestamp", "n": "timestamp", "r": false, "sh": "Date and time the record was updated in the aggregator search index", "t": "`$ANY`", "key$": "timestamp", "index$": 21 }, "title": { "a": true, "h": "Title", "n": "title", "r": false, "sh": "The name of this resource", "t": "`$STRING`", "key$": "title", "index$": 22 }, "title_display": { "a": true, "h": "Title Display", "n": "title_display", "r": false, "sh": "The name of this event formatted with HTML (optional)", "t": "`$ANY`", "key$": "title_display", "index$": 23 }, "updated_at": { "a": true, "h": "Updated At", "n": "updated_at", "r": false, "sh": "Date and time the record was updated in the aggregator database", "t": "`$ANY`", "key$": "updated_at", "index$": 24 } }, "id": { "field": "id", "name": "id" }, "name": "event_occurrence", "op": { "list": { "input": "data", "name": "list", "points": [{ "a": true, "co": { "id": "GET /event-occurrences", "source": "openapi3", "version": 2 }, "g": {}, "k": "http", "m": "GET", "o": "/event-occurrences", "q": {}, "r": {}, "s": [{ "lit": "event-occurrences" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "list" }, "load": { "input": "data", "name": "load", "points": [{ "a": true, "co": { "id": "GET /event-occurrences/{id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "GET", "o": "/event-occurrences/{id}", "q": { "exist": ["id"] }, "r": {}, "s": [{ "lit": "event-occurrences" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "load" } }, "relations": { "ancestors": [] }, "key$": "event_occurrence", "name__orig": "event_occurrence", "Name": "EventOccurrence", "name_": "event_occurrence", "name-": "event-occurrence", "NAME": "EVENT_OCCURRENCE", "index$": 13 }, { "active": true, "entity": "event_occurrence", "key$": "BasicEventOccurrenceFlow", "kind": "basic", "name": "BasicEventOccurrenceFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": {}, "m": {}, "o": "list", "s": [], "v": [{ "apply": "ItemExists", "def": { "ref": "event_occurrence_ref01" } }], "index$": 0 }, { "a": true, "d": {}, "i": { "ref": "event_occurrence_ref01", "srcdatavar": "event_occurrence_ref01_data", "suffix": "_dt0" }, "m": { "id": "event_occurrence01" }, "o": "load", "s": [], "v": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-event_occurrence_ref01" } }], "index$": 1 }] }, 'EventOccurrence', { "GET /event-occurrences": { "protocol": "http", "responses": { "200": { "description": "Successful operation", "schema": { "type": "array", "items": { "properties": { "id": { "description": "Unique identifier of this resource. Taken from the source system.", "key$": "id" }, "api_model": { "description": "REST API resource type or endpoint", "key$": "api_model" }, "api_link": { "description": "REST API link for this resource", "key$": "api_link" }, "title": { "description": "The name of this resource", "key$": "title" }, "title_display": { "description": "The name of this event formatted with HTML (optional)", "key$": "title_display" }, "event_id": { "description": "Identifier of the master event of which this is an occurrence", "key$": "event_id" }, "short_description": { "description": "Brief description of the event", "key$": "short_description" }, "description": { "description": "Description of the event", "key$": "description" }, "image_url": { "description": "The URL of an image representing this page", "key$": "image_url" }, "is_ticketed": { "description": "Whether a ticket is required to attend the event", "key$": "is_ticketed" }, "is_private": { "description": "Whether the event is private. Private events should be omitted from listings.", "key$": "is_private" }, "is_sales_button_hidden": { "description": "Whether the buy tickets button is hidden on the website event page", "key$": "is_sales_button_hidden" }, "start_at": { "description": "The date the event occurrence begins", "key$": "start_at" }, "end_at": { "description": "The date the event occurrence ends", "key$": "end_at" }, "on_sale_at": { "description": "Date and time the event goes on sale", "key$": "on_sale_at" }, "off_sale_at": { "description": "Date and time the event goes off sale", "key$": "off_sale_at" }, "location": { "description": "Where the event takes place", "key$": "location" }, "button_url": { "description": "The URL to the sales site or an RSVP link for this event", "key$": "button_url" }, "button_text": { "description": "The text used on the ticket/registration button", "key$": "button_text" }, "button_caption": { "description": "Additional text below the ticket/registration button", "key$": "button_caption" }, "suggest_autocomplete_boosted": { "description": "Internal field to power the `/autocomplete` endpoint. Do not use directly.", "key$": "suggest_autocomplete_boosted" }, "suggest_autocomplete_all": { "description": "Internal field to power the `/autosuggest` endpoint. Do not use directly.", "key$": "suggest_autocomplete_all" }, "source_updated_at": { "description": "Date and time the resource was updated in the source system", "key$": "source_updated_at" }, "updated_at": { "description": "Date and time the record was updated in the aggregator database", "key$": "updated_at" }, "timestamp": { "description": "Date and time the record was updated in the aggregator search index", "key$": "timestamp" } }, "type": "object", "x-ref": "#/definitions/EventOccurrence", "index$": 0 } } }, "default": { "description": "error", "schema": { "required": ["status", "error", "detail"], "properties": { "status": { "type": "integer" }, "error": { "type": "string" }, "detail": { "type": "string" } }, "x-ref": "#/definitions/Error" } } }, "produces": ["application/json"], "parameters": [], "securitySource": "unspecified" }, "GET /event-occurrences/{id}": { "protocol": "http", "responses": { "200": { "description": "Successful operation", "schema": { "type": "array", "items": { "properties": { "id": { "description": "Unique identifier of this resource. Taken from the source system.", "key$": "id" }, "api_model": { "description": "REST API resource type or endpoint", "key$": "api_model" }, "api_link": { "description": "REST API link for this resource", "key$": "api_link" }, "title": { "description": "The name of this resource", "key$": "title" }, "title_display": { "description": "The name of this event formatted with HTML (optional)", "key$": "title_display" }, "event_id": { "description": "Identifier of the master event of which this is an occurrence", "key$": "event_id" }, "short_description": { "description": "Brief description of the event", "key$": "short_description" }, "description": { "description": "Description of the event", "key$": "description" }, "image_url": { "description": "The URL of an image representing this page", "key$": "image_url" }, "is_ticketed": { "description": "Whether a ticket is required to attend the event", "key$": "is_ticketed" }, "is_private": { "description": "Whether the event is private. Private events should be omitted from listings.", "key$": "is_private" }, "is_sales_button_hidden": { "description": "Whether the buy tickets button is hidden on the website event page", "key$": "is_sales_button_hidden" }, "start_at": { "description": "The date the event occurrence begins", "key$": "start_at" }, "end_at": { "description": "The date the event occurrence ends", "key$": "end_at" }, "on_sale_at": { "description": "Date and time the event goes on sale", "key$": "on_sale_at" }, "off_sale_at": { "description": "Date and time the event goes off sale", "key$": "off_sale_at" }, "location": { "description": "Where the event takes place", "key$": "location" }, "button_url": { "description": "The URL to the sales site or an RSVP link for this event", "key$": "button_url" }, "button_text": { "description": "The text used on the ticket/registration button", "key$": "button_text" }, "button_caption": { "description": "Additional text below the ticket/registration button", "key$": "button_caption" }, "suggest_autocomplete_boosted": { "description": "Internal field to power the `/autocomplete` endpoint. Do not use directly.", "key$": "suggest_autocomplete_boosted" }, "suggest_autocomplete_all": { "description": "Internal field to power the `/autosuggest` endpoint. Do not use directly.", "key$": "suggest_autocomplete_all" }, "source_updated_at": { "description": "Date and time the resource was updated in the source system", "key$": "source_updated_at" }, "updated_at": { "description": "Date and time the record was updated in the aggregator database", "key$": "updated_at" }, "timestamp": { "description": "Date and time the record was updated in the aggregator search index", "key$": "timestamp" } }, "type": "object", "x-ref": "#/definitions/EventOccurrence", "key$": "items" } } }, "default": { "description": "error", "schema": { "required": ["status", "error", "detail"], "properties": { "status": { "type": "integer" }, "error": { "type": "string" }, "detail": { "type": "string" } }, "x-ref": "#/definitions/Error" } } }, "produces": ["application/json"], "parameters": [{ "name": "id", "in": "path", "type": "string", "required": true, "x-ref": "#/parameters/id", "index$": 0 }], "securitySource": "unspecified" } });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let event_occurrence_ref01_data = Object.values(setup.data.existing.event_occurrence)[0];
        // LIST
        const event_occurrence_ref01_ent = client.EventOccurrence();
        const event_occurrence_ref01_match = {};
        const event_occurrence_ref01_list = (await event_occurrence_ref01_ent.list(event_occurrence_ref01_match)).map((e) => e.data());
        // LOAD
        const event_occurrence_ref01_match_dt0 = {};
        event_occurrence_ref01_match_dt0.id = event_occurrence_ref01_data.id;
        const event_occurrence_ref01_data_dt0 = (await event_occurrence_ref01_ent.load(event_occurrence_ref01_match_dt0)).data();
        (0, node_assert_1.default)(event_occurrence_ref01_data_dt0.id === event_occurrence_ref01_data.id);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/event_occurrence/EventOccurrenceTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.ArtInstituteOfChicagoSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['event_occurrence01', 'event_occurrence02', 'event_occurrence03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_OCCURRENCE_ENTID': idmap,
        'ART_INSTITUTE_OF_CHICAGO_TEST_LIVE': 'FALSE',
        'ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN': 'FALSE',
    });
    idmap = env['ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_OCCURRENCE_ENTID'];
    const live = 'TRUE' === env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_OCCURRENCE_ENTID'];
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
//# sourceMappingURL=EventOccurrenceEntity.test.js.map