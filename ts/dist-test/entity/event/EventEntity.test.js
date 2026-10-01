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
(0, node_test_1.describe)('EventEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('ART_INSTITUTE_OF_CHICAGO_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.ArtInstituteOfChicagoSDK.test();
        const ent = testsdk.Event();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE;
        for (const op of ['list', 'load']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'event.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "alt_audience_ids": { "a": true, "h": "Alt Audience Ids", "n": "alt_audience_ids", "r": false, "sh": "Unique identifiers indicating the alternate audiences for this event", "t": "`$ANY`", "key$": "alt_audience_ids", "index$": 0 }, "alt_event_type_ids": { "a": true, "h": "Alt Event Type Ids", "n": "alt_event_type_ids", "r": false, "sh": "Unique identifiers indicating the alternate types of this event", "t": "`$ANY`", "key$": "alt_event_type_ids", "index$": 1 }, "api_link": { "a": true, "h": "Api Link", "n": "api_link", "r": false, "sh": "REST API link for this resource", "t": "`$ANY`", "key$": "api_link", "index$": 2 }, "api_model": { "a": true, "h": "Api Model", "n": "api_model", "r": false, "sh": "REST API resource type or endpoint", "t": "`$ANY`", "key$": "api_model", "index$": 3 }, "audience_id": { "a": true, "h": "Audience Id", "n": "audience_id", "r": false, "sh": "Unique identifier indicating the preferred audience for this event", "t": "`$STRING`", "key$": "audience_id", "index$": 4 }, "buy_button_caption": { "a": true, "h": "Buy Button Caption", "n": "buy_button_caption", "r": false, "sh": "Additional text below the ticket/registration button", "t": "`$ANY`", "key$": "buy_button_caption", "index$": 5 }, "buy_button_text": { "a": true, "h": "Buy Button Text", "n": "buy_button_text", "r": false, "sh": "The text used on the ticket/registration button", "t": "`$ANY`", "key$": "buy_button_text", "index$": 6 }, "date_display": { "a": true, "h": "Date Display", "n": "date_display", "r": false, "sh": "A readable display of the event dates", "t": "`$ANY`", "key$": "date_display", "index$": 7 }, "description": { "a": true, "h": "Description", "n": "description", "r": false, "sh": "All copytext of the event", "t": "`$STRING`", "key$": "description", "index$": 8 }, "door_time": { "a": true, "h": "Door Time", "n": "door_time", "r": false, "sh": "The time the doors open for this event", "t": "`$ANY`", "key$": "door_time", "index$": 9 }, "end_date": { "a": true, "h": "End Date", "n": "end_date", "r": false, "sh": "The date the event ends", "t": "`$ANY`", "key$": "end_date", "index$": 10 }, "end_time": { "a": true, "h": "End Time", "n": "end_time", "r": false, "sh": "The time the event ends", "t": "`$ANY`", "key$": "end_time", "index$": 11 }, "entrance": { "a": true, "h": "Entrance", "n": "entrance", "r": false, "sh": "Which entrance to use for this event", "t": "`$ANY`", "key$": "entrance", "index$": 12 }, "event_host_id": { "a": true, "h": "Event Host Id", "n": "event_host_id", "r": false, "sh": "Unique identifier of the host (cf.", "t": "`$STRING`", "key$": "event_host_id", "index$": 13 }, "event_host_title": { "a": true, "h": "Event Host Title", "n": "event_host_title", "r": false, "sh": "Unique identifier of the host (cf.", "t": "`$ANY`", "key$": "event_host_title", "index$": 14 }, "event_type_id": { "a": true, "h": "Event Type Id", "n": "event_type_id", "r": false, "sh": "Unique identifier indicating the preferred type of this event", "t": "`$STRING`", "key$": "event_type_id", "index$": 15 }, "header_description": { "a": true, "h": "Header Description", "n": "header_description", "r": false, "sh": "Brief description of the event displayed below the title", "t": "`$ANY`", "key$": "header_description", "index$": 16 }, "hero_caption": { "a": true, "h": "Hero Caption", "n": "hero_caption", "r": false, "sh": "Text displayed with the hero image on the event", "t": "`$ANY`", "key$": "hero_caption", "index$": 17 }, "id": { "a": true, "h": "Id", "n": "id", "r": false, "sh": "Unique identifier of this resource.", "t": "`$STRING`", "key$": "id", "index$": 18 }, "image_url": { "a": true, "h": "Image Url", "n": "image_url", "r": false, "sh": "The URL of an image representing this page", "t": "`$ANY`", "key$": "image_url", "index$": 19 }, "is_admission_required": { "a": true, "h": "Is Admission Required", "n": "is_admission_required", "r": false, "sh": "Whether admission to the museum is required to attend this event", "t": "`$BOOLEAN`", "key$": "is_admission_required", "index$": 20 }, "is_after_hours": { "a": true, "h": "Is After Hours", "n": "is_after_hours", "r": false, "sh": "Whether the event is to be held after the museum closes", "t": "`$BOOLEAN`", "key$": "is_after_hours", "index$": 21 }, "is_free": { "a": true, "h": "Is Free", "n": "is_free", "r": false, "sh": "Whether the event is free", "t": "`$BOOLEAN`", "key$": "is_free", "index$": 22 }, "is_member_exclusive": { "a": true, "h": "Is Member Exclusive", "n": "is_member_exclusive", "r": false, "sh": "Whether the event is exclusive to members of the museum", "t": "`$BOOLEAN`", "key$": "is_member_exclusive", "index$": 23 }, "is_private": { "a": true, "h": "Is Private", "n": "is_private", "r": false, "sh": "Whether the event is private", "t": "`$BOOLEAN`", "key$": "is_private", "index$": 24 }, "is_registration_required": { "a": true, "h": "Is Registration Required", "n": "is_registration_required", "r": false, "sh": "Whether registration is required to attend the event", "t": "`$BOOLEAN`", "key$": "is_registration_required", "index$": 25 }, "is_sales_button_hidden": { "a": true, "h": "Is Sales Button Hidden", "n": "is_sales_button_hidden", "r": false, "sh": "Whether the buy tickets button is hidden on the website event page", "t": "`$BOOLEAN`", "key$": "is_sales_button_hidden", "index$": 26 }, "is_sold_out": { "a": true, "h": "Is Sold Out", "n": "is_sold_out", "r": false, "sh": "Whether the event is sold out", "t": "`$BOOLEAN`", "key$": "is_sold_out", "index$": 27 }, "is_ticketed": { "a": true, "h": "Is Ticketed", "n": "is_ticketed", "r": false, "sh": "Whether a ticket is required to attend the event", "t": "`$BOOLEAN`", "key$": "is_ticketed", "index$": 28 }, "is_virtual_event": { "a": true, "h": "Is Virtual Event", "n": "is_virtual_event", "r": false, "sh": "Whether the event is being held virtually", "t": "`$BOOLEAN`", "key$": "is_virtual_event", "index$": 29 }, "join_url": { "a": true, "h": "Join Url", "n": "join_url", "r": false, "sh": "URL to the membership signup page via this event", "t": "`$ANY`", "key$": "join_url", "index$": 30 }, "layout_type": { "a": true, "h": "Layout Type", "n": "layout_type", "r": false, "sh": "Number indicating the type of layout this event page uses", "t": "`$ANY`", "key$": "layout_type", "index$": 31 }, "list_description": { "a": true, "h": "List Description", "n": "list_description", "r": false, "sh": "One-sentence description of the event displayed in listings", "t": "`$ANY`", "key$": "list_description", "index$": 32 }, "location": { "a": true, "h": "Location", "n": "location", "r": false, "sh": "Where the event takes place", "t": "`$ANY`", "key$": "location", "index$": 33 }, "program_ids": { "a": true, "h": "Program Ids", "n": "program_ids", "r": false, "sh": "Unique identifiers indicating the programs this event is a part of", "t": "`$ANY`", "key$": "program_ids", "index$": 34 }, "program_titles": { "a": true, "h": "Program Titles", "n": "program_titles", "r": false, "sh": "Titles of the programs this event is a part of", "t": "`$ANY`", "key$": "program_titles", "index$": 35 }, "rsvp_link": { "a": true, "h": "Rsvp Link", "n": "rsvp_link", "r": false, "sh": "The URL to the sales site for this event", "t": "`$ANY`", "key$": "rsvp_link", "index$": 36 }, "search_tags": { "a": true, "h": "Search Tags", "n": "search_tags", "r": false, "sh": "Editor-specified list of tags to aid in internal search", "t": "`$ANY`", "key$": "search_tags", "index$": 37 }, "short_description": { "a": true, "h": "Short Description", "n": "short_description", "r": false, "sh": "Brief description of the event", "t": "`$ANY`", "key$": "short_description", "index$": 38 }, "slug": { "a": true, "h": "Slug", "n": "slug", "r": false, "sh": "A string used in the URL for this event", "t": "`$STRING`", "key$": "slug", "index$": 39 }, "source_updated_at": { "a": true, "h": "Source Updated At", "n": "source_updated_at", "r": false, "sh": "Date and time the resource was updated in the source system", "t": "`$ANY`", "key$": "source_updated_at", "index$": 40 }, "start_date": { "a": true, "h": "Start Date", "n": "start_date", "r": false, "sh": "The date the event begins", "t": "`$ANY`", "key$": "start_date", "index$": 41 }, "start_time": { "a": true, "h": "Start Time", "n": "start_time", "r": false, "sh": "The time the event starts", "t": "`$ANY`", "key$": "start_time", "index$": 42 }, "suggest_autocomplete_all": { "a": true, "h": "Suggest Autocomplete All", "n": "suggest_autocomplete_all", "r": false, "sh": "Internal field to power the `/autosuggest` endpoint.", "t": "`$ANY`", "key$": "suggest_autocomplete_all", "index$": 43 }, "suggest_autocomplete_boosted": { "a": true, "h": "Suggest Autocomplete Boosted", "n": "suggest_autocomplete_boosted", "r": false, "sh": "Internal field to power the `/autocomplete` endpoint.", "t": "`$ANY`", "key$": "suggest_autocomplete_boosted", "index$": 44 }, "survey_url": { "a": true, "h": "Survey Url", "n": "survey_url", "r": false, "sh": "URL to the survey associated with this event", "t": "`$ANY`", "key$": "survey_url", "index$": 45 }, "ticketed_event_id": { "a": true, "h": "Ticketed Event Id", "n": "ticketed_event_id", "r": false, "sh": "Unique identifier of the event in the ticketing system this website event is tied to", "t": "`$STRING`", "key$": "ticketed_event_id", "index$": 46 }, "timestamp": { "a": true, "h": "Timestamp", "n": "timestamp", "r": false, "sh": "Date and time the record was updated in the aggregator search index", "t": "`$ANY`", "key$": "timestamp", "index$": 47 }, "title": { "a": true, "h": "Title", "n": "title", "r": false, "sh": "The name of this resource", "t": "`$STRING`", "key$": "title", "index$": 48 }, "title_display": { "a": true, "h": "Title Display", "n": "title_display", "r": false, "sh": "The name of this event formatted with HTML (optional)", "t": "`$ANY`", "key$": "title_display", "index$": 49 }, "updated_at": { "a": true, "h": "Updated At", "n": "updated_at", "r": false, "sh": "Date and time the record was updated in the aggregator database", "t": "`$ANY`", "key$": "updated_at", "index$": 50 }, "virtual_event_passcode": { "a": true, "h": "Virtual Event Passcode", "n": "virtual_event_passcode", "r": false, "sh": "Passcode to access the virtual event", "t": "`$ANY`", "key$": "virtual_event_passcode", "index$": 51 }, "virtual_event_url": { "a": true, "h": "Virtual Event Url", "n": "virtual_event_url", "r": false, "sh": "URL to the virtual event", "t": "`$ANY`", "key$": "virtual_event_url", "index$": 52 } }, "id": { "field": "id", "name": "id" }, "name": "event", "op": { "list": { "input": "data", "name": "list", "points": [{ "a": true, "co": { "id": "GET /events", "source": "openapi3", "version": 2 }, "g": {}, "k": "http", "m": "GET", "o": "/events", "q": {}, "r": {}, "s": [{ "lit": "events" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "list" }, "load": { "input": "data", "name": "load", "points": [{ "a": true, "co": { "id": "GET /events/{id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "GET", "o": "/events/{id}", "q": { "exist": ["id"] }, "r": {}, "s": [{ "lit": "events" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "load" } }, "relations": { "ancestors": [] }, "key$": "event", "name__orig": "event", "Name": "Event", "name_": "event", "name-": "event", "NAME": "EVENT", "index$": 12 }, { "active": true, "entity": "event", "key$": "BasicEventFlow", "kind": "basic", "name": "BasicEventFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": {}, "m": {}, "o": "list", "s": [], "v": [{ "apply": "ItemExists", "def": { "ref": "event_ref01" } }], "index$": 0 }, { "a": true, "d": {}, "i": { "ref": "event_ref01", "srcdatavar": "event_ref01_data", "suffix": "_dt0" }, "m": { "id": "event01" }, "o": "load", "s": [], "v": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-event_ref01" } }], "index$": 1 }] }, 'Event', { "GET /events": { "protocol": "http", "parameters": [] }, "GET /events/{id}": { "protocol": "http", "parameters": [{ "name": "id", "in": "path", "type": "string", "required": true, "x-ref": "#/parameters/id", "index$": 0 }] } });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let event_ref01_data = Object.values(setup.data.existing.event)[0];
        // LIST
        const event_ref01_ent = client.Event();
        const event_ref01_match = {};
        const event_ref01_list = (await event_ref01_ent.list(event_ref01_match)).map((e) => e.data());
        // LOAD
        const event_ref01_match_dt0 = {};
        event_ref01_match_dt0.id = event_ref01_data.id;
        const event_ref01_data_dt0 = (await event_ref01_ent.load(event_ref01_match_dt0)).data();
        (0, node_assert_1.default)(event_ref01_data_dt0.id === event_ref01_data.id);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/event/EventTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.ArtInstituteOfChicagoSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['event01', 'event02', 'event03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_ENTID': idmap,
        'ART_INSTITUTE_OF_CHICAGO_TEST_LIVE': 'FALSE',
        'ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN': 'FALSE',
    });
    idmap = env['ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_ENTID'];
    const live = 'TRUE' === env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_ENTID'];
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
//# sourceMappingURL=EventEntity.test.js.map