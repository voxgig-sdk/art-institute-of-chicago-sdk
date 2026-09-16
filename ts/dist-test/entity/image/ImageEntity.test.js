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
(0, node_test_1.describe)('ImageEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('ART_INSTITUTE_OF_CHICAGO_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.ArtInstituteOfChicagoSDK.test();
        const ent = testsdk.Image();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE;
        for (const op of ['list', 'load']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'image.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [{ "active": true, "name": "ahash", "req": false, "short": "Image hash generated using ahash algorithm with 64 boolean subfields", "type": "`$ANY`", "index$": 0 }, { "active": true, "name": "alt_text", "req": false, "short": "Alternative text for the asset to describe it to people with low or no vision", "type": "`$ANY`", "index$": 1 }, { "active": true, "name": "api_link", "req": false, "short": "REST API link for this resource", "type": "`$ANY`", "index$": 2 }, { "active": true, "name": "api_model", "req": false, "short": "REST API resource type or endpoint", "type": "`$ANY`", "index$": 3 }, { "active": true, "name": "artwork_ids", "req": false, "short": "Unique identifiers of the artworks associated with this asset", "type": "`$ANY`", "index$": 4 }, { "active": true, "name": "artwork_titles", "req": false, "short": "Names of the artworks associated with this asset", "type": "`$ANY`", "index$": 5 }, { "active": true, "name": "color", "req": false, "short": "Dominant color of this image in HSL", "type": "`$ANY`", "index$": 6 }, { "active": true, "name": "colorfulness", "req": false, "short": "Unbounded positive float representing an abstract measure of colorfulness.", "type": "`$ANY`", "index$": 7 }, { "active": true, "name": "content", "req": false, "short": "Text of or URL to the contents of this asset", "type": "`$ANY`", "index$": 8 }, { "active": true, "name": "content_e_tag", "req": false, "short": "Arbitrary unique identifier that changes when the binary file gets updated", "type": "`$ANY`", "index$": 9 }, { "active": true, "name": "credit_line", "req": false, "short": "Asset-specific copyright information", "type": "`$ANY`", "index$": 10 }, { "active": true, "name": "fingerprint", "req": false, "short": "Image hashes: aHash, dHash, pHash, wHash", "type": "`$ANY`", "index$": 11 }, { "active": true, "name": "height", "req": false, "short": "Native height of the image", "type": "`$NUMBER`", "index$": 12 }, { "active": true, "name": "id", "req": false, "short": "Unique identifier of this resource.", "type": "`$STRING`", "index$": 13 }, { "active": true, "name": "iiif_url", "req": false, "short": "IIIF URL of this image", "type": "`$ANY`", "index$": 14 }, { "active": true, "name": "is_educational_resource", "req": false, "short": "Whether this resource is considered to be educational", "type": "`$BOOLEAN`", "index$": 15 }, { "active": true, "name": "is_multimedia_resource", "req": false, "short": "Whether this resource is considered to be multimedia", "type": "`$BOOLEAN`", "index$": 16 }, { "active": true, "name": "is_teacher_resource", "req": false, "short": "Whether this resource is considered to be educational", "type": "`$BOOLEAN`", "index$": 17 }, { "active": true, "name": "lake_guid", "req": false, "short": "Unique UUID of this resource in LAKE, our DAMS.", "type": "`$ANY`", "index$": 18 }, { "active": true, "name": "lqip", "req": false, "short": "Low-quality image placeholder (LQIP).", "type": "`$ANY`", "index$": 19 }, { "active": true, "name": "phash", "req": false, "short": "Image hash generated using phash algorithm with 64 boolean subfields", "type": "`$ANY`", "index$": 20 }, { "active": true, "name": "source_updated_at", "req": false, "short": "Date and time the resource was updated in the source system", "type": "`$ANY`", "index$": 21 }, { "active": true, "name": "suggest_autocomplete_all", "req": false, "short": "Internal field to power the `/autosuggest` endpoint.", "type": "`$ANY`", "index$": 22 }, { "active": true, "name": "suggest_autocomplete_boosted", "req": false, "short": "Internal field to power the `/autocomplete` endpoint.", "type": "`$ANY`", "index$": 23 }, { "active": true, "name": "timestamp", "req": false, "short": "Date and time the record was updated in the aggregator search index", "type": "`$ANY`", "index$": 24 }, { "active": true, "name": "title", "req": false, "short": "The name of this resource", "type": "`$STRING`", "index$": 25 }, { "active": true, "name": "type", "req": false, "short": "Type always takes one of the following values: image, sound, text, video", "type": "`$ANY`", "index$": 26 }, { "active": true, "name": "updated_at", "req": false, "short": "Date and time the record was updated in the aggregator database", "type": "`$ANY`", "index$": 27 }, { "active": true, "name": "width", "req": false, "short": "Native width of the image", "type": "`$NUMBER`", "index$": 28 }], "id": { "field": "id", "name": "id" }, "name": "image", "op": { "list": { "input": "data", "name": "list", "points": [{ "active": true, "args": {}, "contract": { "id": "GET /images", "json": "{\"parameters\":[],\"produces\":[\"application/json\"],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"Successful operation\",\"schema\":{\"items\":{\"properties\":{\"ahash\":{\"description\":\"Image hash generated using ahash algorithm with 64 boolean subfields\"},\"alt_text\":{\"description\":\"Alternative text for the asset to describe it to people with low or no vision\"},\"api_link\":{\"description\":\"REST API link for this resource\"},\"api_model\":{\"description\":\"REST API resource type or endpoint\"},\"artwork_ids\":{\"description\":\"Unique identifiers of the artworks associated with this asset\"},\"artwork_titles\":{\"description\":\"Names of the artworks associated with this asset\"},\"color\":{\"description\":\"Dominant color of this image in HSL\"},\"colorfulness\":{\"description\":\"Unbounded positive float representing an abstract measure of colorfulness.\"},\"content\":{\"description\":\"Text of or URL to the contents of this asset\"},\"content_e_tag\":{\"description\":\"Arbitrary unique identifier that changes when the binary file gets updated\"},\"credit_line\":{\"description\":\"Asset-specific copyright information\"},\"fingerprint\":{\"description\":\"Image hashes: aHash, dHash, pHash, wHash\"},\"height\":{\"description\":\"Native height of the image\"},\"id\":{\"description\":\"Unique identifier of this resource. Taken from the source system.\"},\"iiif_url\":{\"description\":\"IIIF URL of this image\"},\"is_educational_resource\":{\"description\":\"Whether this resource is considered to be educational\"},\"is_multimedia_resource\":{\"description\":\"Whether this resource is considered to be multimedia\"},\"is_teacher_resource\":{\"description\":\"Whether this resource is considered to be educational\"},\"lake_guid\":{\"description\":\"Unique UUID of this resource in LAKE, our DAMS.\"},\"lqip\":{\"description\":\"Low-quality image placeholder (LQIP). Currently a 5x5-constrained, base64-encoded GIF.\"},\"phash\":{\"description\":\"Image hash generated using phash algorithm with 64 boolean subfields\"},\"source_updated_at\":{\"description\":\"Date and time the resource was updated in the source system\"},\"suggest_autocomplete_all\":{\"description\":\"Internal field to power the `/autosuggest` endpoint. Do not use directly.\"},\"suggest_autocomplete_boosted\":{\"description\":\"Internal field to power the `/autocomplete` endpoint. Do not use directly.\"},\"timestamp\":{\"description\":\"Date and time the record was updated in the aggregator search index\"},\"title\":{\"description\":\"The name of this resource\"},\"type\":{\"description\":\"Type always takes one of the following values: image, sound, text, video\"},\"updated_at\":{\"description\":\"Date and time the record was updated in the aggregator database\"},\"width\":{\"description\":\"Native width of the image\"}},\"type\":\"object\"},\"type\":\"array\"}},\"default\":{\"description\":\"error\",\"schema\":{\"properties\":{\"detail\":{\"type\":\"string\"},\"error\":{\"type\":\"string\"},\"status\":{\"type\":\"integer\"}},\"required\":[\"status\",\"error\",\"detail\"]}}},\"securitySource\":\"unspecified\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/images", "segments": [{ "lit": "images" }], "select": {}, "transform": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "list" }, "load": { "input": "data", "name": "load", "points": [{ "active": true, "args": { "params": [{ "active": true, "kind": "param", "name": "id", "orig": "id", "reqd": true, "type": "`$STRING`", "index$": 0 }] }, "contract": { "id": "GET /images/{id}", "json": "{\"parameters\":[{\"in\":\"path\",\"name\":\"id\",\"required\":true,\"type\":\"string\"}],\"produces\":[\"application/json\"],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"Successful operation\",\"schema\":{\"items\":{\"properties\":{\"ahash\":{\"description\":\"Image hash generated using ahash algorithm with 64 boolean subfields\"},\"alt_text\":{\"description\":\"Alternative text for the asset to describe it to people with low or no vision\"},\"api_link\":{\"description\":\"REST API link for this resource\"},\"api_model\":{\"description\":\"REST API resource type or endpoint\"},\"artwork_ids\":{\"description\":\"Unique identifiers of the artworks associated with this asset\"},\"artwork_titles\":{\"description\":\"Names of the artworks associated with this asset\"},\"color\":{\"description\":\"Dominant color of this image in HSL\"},\"colorfulness\":{\"description\":\"Unbounded positive float representing an abstract measure of colorfulness.\"},\"content\":{\"description\":\"Text of or URL to the contents of this asset\"},\"content_e_tag\":{\"description\":\"Arbitrary unique identifier that changes when the binary file gets updated\"},\"credit_line\":{\"description\":\"Asset-specific copyright information\"},\"fingerprint\":{\"description\":\"Image hashes: aHash, dHash, pHash, wHash\"},\"height\":{\"description\":\"Native height of the image\"},\"id\":{\"description\":\"Unique identifier of this resource. Taken from the source system.\"},\"iiif_url\":{\"description\":\"IIIF URL of this image\"},\"is_educational_resource\":{\"description\":\"Whether this resource is considered to be educational\"},\"is_multimedia_resource\":{\"description\":\"Whether this resource is considered to be multimedia\"},\"is_teacher_resource\":{\"description\":\"Whether this resource is considered to be educational\"},\"lake_guid\":{\"description\":\"Unique UUID of this resource in LAKE, our DAMS.\"},\"lqip\":{\"description\":\"Low-quality image placeholder (LQIP). Currently a 5x5-constrained, base64-encoded GIF.\"},\"phash\":{\"description\":\"Image hash generated using phash algorithm with 64 boolean subfields\"},\"source_updated_at\":{\"description\":\"Date and time the resource was updated in the source system\"},\"suggest_autocomplete_all\":{\"description\":\"Internal field to power the `/autosuggest` endpoint. Do not use directly.\"},\"suggest_autocomplete_boosted\":{\"description\":\"Internal field to power the `/autocomplete` endpoint. Do not use directly.\"},\"timestamp\":{\"description\":\"Date and time the record was updated in the aggregator search index\"},\"title\":{\"description\":\"The name of this resource\"},\"type\":{\"description\":\"Type always takes one of the following values: image, sound, text, video\"},\"updated_at\":{\"description\":\"Date and time the record was updated in the aggregator database\"},\"width\":{\"description\":\"Native width of the image\"}},\"type\":\"object\"},\"type\":\"array\"}},\"default\":{\"description\":\"error\",\"schema\":{\"properties\":{\"detail\":{\"type\":\"string\"},\"error\":{\"type\":\"string\"},\"status\":{\"type\":\"integer\"}},\"required\":[\"status\",\"error\",\"detail\"]}}},\"securitySource\":\"unspecified\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/images/{id}", "segments": [{ "lit": "images" }, { "var": "id" }], "select": { "exist": ["id"] }, "transform": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "load" } }, "relations": { "ancestors": [] }, "key$": "image", "name__orig": "image", "Name": "Image", "name_": "image", "name-": "image", "NAME": "IMAGE", "index$": 20 }, { "active": true, "entity": "image", "key$": "BasicImageFlow", "kind": "basic", "name": "BasicImageFlow", "param": {}, "step": [{ "active": true, "data": {}, "input": {}, "match": {}, "op": "list", "spec": [], "valid": [{ "apply": "ItemExists", "def": { "ref": "image_ref01" } }], "index$": 0 }, { "active": true, "data": {}, "input": { "ref": "image_ref01", "srcdatavar": "image_ref01_data", "suffix": "_dt0" }, "match": { "id": "image01" }, "op": "load", "spec": [], "valid": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-image_ref01" } }], "index$": 1 }] }, 'Image');
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let image_ref01_data = Object.values(setup.data.existing.image)[0];
        // LIST
        const image_ref01_ent = client.Image();
        const image_ref01_match = {};
        const image_ref01_list = (await image_ref01_ent.list(image_ref01_match)).map((e) => e.data());
        // LOAD
        const image_ref01_match_dt0 = {};
        image_ref01_match_dt0.id = image_ref01_data.id;
        const image_ref01_data_dt0 = (await image_ref01_ent.load(image_ref01_match_dt0)).data();
        (0, node_assert_1.default)(image_ref01_data_dt0.id === image_ref01_data.id);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/image/ImageTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.ArtInstituteOfChicagoSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['image01', 'image02', 'image03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'ART_INSTITUTE_OF_CHICAGO_TEST_IMAGE_ENTID': idmap,
        'ART_INSTITUTE_OF_CHICAGO_TEST_LIVE': 'FALSE',
        'ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN': 'FALSE',
    });
    idmap = env['ART_INSTITUTE_OF_CHICAGO_TEST_IMAGE_ENTID'];
    const live = 'TRUE' === env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['ART_INSTITUTE_OF_CHICAGO_TEST_IMAGE_ENTID'];
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
//# sourceMappingURL=ImageEntity.test.js.map