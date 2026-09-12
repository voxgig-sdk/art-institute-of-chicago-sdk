"use strict";
// ArtInstituteOfChicago Ts SDK
Object.defineProperty(exports, "__esModule", { value: true });
exports.SDK = exports.ArtInstituteOfChicagoSDK = exports.ArtInstituteOfChicagoEntityBase = exports.BaseFeature = exports.config = exports.stdutil = void 0;
const AgentEntity_1 = require("./entity/AgentEntity");
const AgentRoleEntity_1 = require("./entity/AgentRoleEntity");
const AgentTypeEntity_1 = require("./entity/AgentTypeEntity");
const ArticleEntity_1 = require("./entity/ArticleEntity");
const ArtworkEntity_1 = require("./entity/ArtworkEntity");
const ArtworkDateQualifierEntity_1 = require("./entity/ArtworkDateQualifierEntity");
const ArtworkPlaceQualifierEntity_1 = require("./entity/ArtworkPlaceQualifierEntity");
const ArtworkTypeEntity_1 = require("./entity/ArtworkTypeEntity");
const CategoryTermEntity_1 = require("./entity/CategoryTermEntity");
const DigitalPublicationEntity_1 = require("./entity/DigitalPublicationEntity");
const DigitalPublicationArticleEntity_1 = require("./entity/DigitalPublicationArticleEntity");
const EducatorResourceEntity_1 = require("./entity/EducatorResourceEntity");
const EventEntity_1 = require("./entity/EventEntity");
const EventOccurrenceEntity_1 = require("./entity/EventOccurrenceEntity");
const EventProgramEntity_1 = require("./entity/EventProgramEntity");
const ExhibitionEntity_1 = require("./entity/ExhibitionEntity");
const GalleryEntity_1 = require("./entity/GalleryEntity");
const GenericPageEntity_1 = require("./entity/GenericPageEntity");
const HighlightEntity_1 = require("./entity/HighlightEntity");
const HourEntity_1 = require("./entity/HourEntity");
const ImageEntity_1 = require("./entity/ImageEntity");
const LandingPageEntity_1 = require("./entity/LandingPageEntity");
const PlaceEntity_1 = require("./entity/PlaceEntity");
const PressReleaseEntity_1 = require("./entity/PressReleaseEntity");
const PrintedPublicationEntity_1 = require("./entity/PrintedPublicationEntity");
const ProductEntity_1 = require("./entity/ProductEntity");
const PublicationEntity_1 = require("./entity/PublicationEntity");
const SearchEntity_1 = require("./entity/SearchEntity");
const SectionEntity_1 = require("./entity/SectionEntity");
const SiteEntity_1 = require("./entity/SiteEntity");
const SoundEntity_1 = require("./entity/SoundEntity");
const StaticPageEntity_1 = require("./entity/StaticPageEntity");
const TextEntity_1 = require("./entity/TextEntity");
const TourEntity_1 = require("./entity/TourEntity");
const VideoEntity_1 = require("./entity/VideoEntity");
const node_util_1 = require("node:util");
const Config_1 = require("./Config");
Object.defineProperty(exports, "config", { enumerable: true, get: function () { return Config_1.config; } });
const ArtInstituteOfChicagoEntityBase_1 = require("./ArtInstituteOfChicagoEntityBase");
Object.defineProperty(exports, "ArtInstituteOfChicagoEntityBase", { enumerable: true, get: function () { return ArtInstituteOfChicagoEntityBase_1.ArtInstituteOfChicagoEntityBase; } });
const Utility_1 = require("./utility/Utility");
const BaseFeature_1 = require("./feature/base/BaseFeature");
Object.defineProperty(exports, "BaseFeature", { enumerable: true, get: function () { return BaseFeature_1.BaseFeature; } });
const stdutil = new Utility_1.Utility();
exports.stdutil = stdutil;
class ArtInstituteOfChicagoSDK {
    _mode = 'live';
    _options;
    _utility = new Utility_1.Utility();
    _features;
    _rootctx;
    constructor(options) {
        this._rootctx = this._utility.makeContext({
            client: this,
            utility: this._utility,
            config: Config_1.config,
            options,
            shared: new WeakMap()
        });
        this._options = this._utility.makeOptions(this._rootctx);
        const struct = this._utility.struct;
        const getpath = struct.getpath;
        if (true === getpath(this._options.feature, 'test.active')) {
            this._mode = 'test';
        }
        this._rootctx.options = this._options;
        this._features = [];
        const featureAdd = this._utility.featureAdd;
        const featureInit = this._utility.featureInit;
        // Add features in the resolved order (makeOptions puts an explicit
        // array order first, else defaults to test-first). Ordering matters:
        // the `test` feature installs the base mock transport and the transport
        // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is current,
        // so `test` must be added before them to sit at the base of the chain.
        const extend = this._options.extend || [];
        const featureorder = getpath(this._options, '__derived__.featureorder') || [];
        for (const fname of featureorder) {
            const fopts = this._options.feature[fname] || {};
            if (fopts.active) {
                // An active name with no generated class is legal when an
                // extend-supplied instance carries that name (station's adopt
                // path): the instance is added below, positioned by its own
                // __after__ entry, so skip it here rather than fail construction.
                if (!this._rootctx.config.hasFeature(fname) &&
                    extend.some((f) => fname === f.name)) {
                    continue;
                }
                featureAdd(this._rootctx, this._rootctx.config.makeFeature(fname));
            }
        }
        for (let f of extend) {
            featureAdd(this._rootctx, f);
        }
        for (let f of this._features) {
            featureInit(this._rootctx, f);
        }
        const featureHook = this._utility.featureHook;
        featureHook(this._rootctx, 'PostConstruct');
    }
    options() {
        return this._utility.struct.clone(this._options);
    }
    utility() {
        return this._utility.struct.clone(this._utility);
    }
    async prepare(fetchargs) {
        const utility = this._utility;
        const struct = utility.struct;
        const clone = struct.clone;
        const { makeContext, makeFetchDef, prepareHeaders, prepareAuth, } = utility;
        fetchargs = fetchargs || {};
        let ctx = makeContext({
            opname: 'prepare',
            ctrl: fetchargs.ctrl || {},
        }, this._rootctx);
        const options = this._options;
        // Build spec directly from SDK options + user-provided fetch args.
        const spec = {
            base: options.base,
            prefix: options.prefix,
            suffix: options.suffix,
            path: fetchargs.path || '',
            method: fetchargs.method || 'GET',
            params: fetchargs.params || {},
            query: fetchargs.query || {},
            headers: prepareHeaders(ctx),
            body: fetchargs.body,
            step: 'start',
        };
        ctx.spec = spec;
        // Merge user-provided headers over SDK defaults.
        if (fetchargs.headers) {
            const uheaders = fetchargs.headers;
            for (let key in uheaders) {
                spec.headers[key] = uheaders[key];
            }
        }
        // Apply SDK auth (apikey, auth prefix, etc.)
        const authResult = prepareAuth(ctx);
        if (authResult instanceof Error) {
            return authResult;
        }
        return makeFetchDef(ctx);
    }
    // Raw endpoint access is operator-controllable, like every entity op.
    // Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
    // either one reaches the same endpoint.
    async direct(fetchargs) {
        if (!this._options.allow.op.includes('direct')) {
            return {
                ok: false,
                err: new Error('ArtInstituteOfChicagoSDK: direct: operation not allowed by' +
                    ' SDK option allow.op value: "' + this._options.allow.op + '"'),
            };
        }
        return this._rawRequest(fetchargs);
    }
    // Ungated request path shared by direct() and graphql(), each of which
    // checks its own allow.op token first. Private, rather than a flag on
    // fetchargs: a caller-supplied marker would let anyone opt straight back
    // out of the gate by passing it.
    async _rawRequest(fetchargs) {
        const utility = this._utility;
        const fetcher = utility.fetcher;
        const makeContext = utility.makeContext;
        const fetchdef = await this.prepare(fetchargs);
        if (fetchdef instanceof Error) {
            return fetchdef;
        }
        let ctx = makeContext({
            opname: 'direct',
            ctrl: (fetchargs || {}).ctrl || {},
        }, this._rootctx);
        try {
            const fetched = await fetcher(ctx, fetchdef.url, fetchdef);
            if (null == fetched) {
                return { ok: false, err: ctx.error('direct_no_response', 'response: undefined') };
            }
            else if (fetched instanceof Error) {
                return { ok: false, err: fetched };
            }
            const status = fetched.status;
            // No body responses (204 No Content, 304 Not Modified) and explicit
            // zero content-length must skip JSON parsing — fetched.json() would
            // throw `Unexpected end of JSON input` on an empty body.
            const headers = fetched.headers;
            const contentLength = headers && 'function' === typeof headers.get
                ? headers.get('content-length')
                : (headers || {})['content-length'];
            const noBody = 204 === status || 304 === status || '0' === String(contentLength);
            let json = undefined;
            if (!noBody) {
                try {
                    json = 'function' === typeof fetched.json ? await fetched.json() : fetched.json;
                }
                catch (parseErr) {
                    // Body wasn't valid JSON — surface the raw response rather than
                    // throwing. data stays undefined; callers can inspect status/headers.
                    json = undefined;
                }
            }
            return {
                ok: status >= 200 && status < 300,
                status,
                headers: fetched.headers,
                data: json,
            };
        }
        catch (err) {
            return { ok: false, err };
        }
    }
    // Raw GraphQL access: the pressure valve that makes the generated
    // surface's deliberate omissions (per-call selection sets, typed filter
    // builders, batching, subscriptions) livable — the whole schema stays
    // reachable.
    //
    // Thin wrapper over the same prepare/fetch path `direct` uses, with the
    // one thing raw `direct` cannot do for GraphQL: a GraphQL failure rides
    // HTTP 200 as a top-level `errors` array, so status alone would report a
    // failed query as ok.
    //
    // NOTE: like `direct`, this bypasses the feature pipeline — no retry,
    // ratelimit or paging features apply.
    async graphql(query, variables, ctrl) {
        const options = this._options;
        if (!options.allow.op.includes('graphql')) {
            return {
                ok: false,
                err: new Error('ArtInstituteOfChicagoSDK: graphql: operation not allowed by' +
                    ' SDK option allow.op value: "' + options.allow.op + '"'),
            };
        }
        const res = await this._rawRequest({
            method: 'POST',
            headers: { 'content-type': 'application/json' },
            body: { query, variables: variables || {} },
            ctrl,
        });
        if (res instanceof Error) {
            return res;
        }
        // Errors are read BEFORE any status check: a GraphQL parse or validation
        // failure comes back as HTTP 400 carrying the standard { errors: [...] }
        // body, and the raw path represents a non-2xx as { ok: false } with no
        // err — so returning early on status would discard the server's own
        // diagnostics, which are the only useful part of that response.
        const errors = null == res.data ? undefined : res.data.errors;
        if (null != errors && Array.isArray(errors) && 0 < errors.length) {
            const first = errors[0] || {};
            const err = new Error('ArtInstituteOfChicagoSDK: graphql: ' +
                (first.message || 'graphql error'));
            err.graphql = errors;
            return { ok: false, status: res.status, headers: res.headers, err, data: res.data };
        }
        return res;
    }
    // Entity access: `client.Agent().list()` / `client.Agent().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Agent(entopts) {
        const self = this;
        return new AgentEntity_1.AgentEntity(self, entopts);
    }
    // Entity access: `client.AgentRole().list()` / `client.AgentRole().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    AgentRole(entopts) {
        const self = this;
        return new AgentRoleEntity_1.AgentRoleEntity(self, entopts);
    }
    // Entity access: `client.AgentType().list()` / `client.AgentType().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    AgentType(entopts) {
        const self = this;
        return new AgentTypeEntity_1.AgentTypeEntity(self, entopts);
    }
    // Entity access: `client.Article().list()` / `client.Article().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Article(entopts) {
        const self = this;
        return new ArticleEntity_1.ArticleEntity(self, entopts);
    }
    // Entity access: `client.Artwork().list()` / `client.Artwork().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Artwork(entopts) {
        const self = this;
        return new ArtworkEntity_1.ArtworkEntity(self, entopts);
    }
    // Entity access: `client.ArtworkDateQualifier().list()` / `client.ArtworkDateQualifier().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ArtworkDateQualifier(entopts) {
        const self = this;
        return new ArtworkDateQualifierEntity_1.ArtworkDateQualifierEntity(self, entopts);
    }
    // Entity access: `client.ArtworkPlaceQualifier().list()` / `client.ArtworkPlaceQualifier().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ArtworkPlaceQualifier(entopts) {
        const self = this;
        return new ArtworkPlaceQualifierEntity_1.ArtworkPlaceQualifierEntity(self, entopts);
    }
    // Entity access: `client.ArtworkType().list()` / `client.ArtworkType().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ArtworkType(entopts) {
        const self = this;
        return new ArtworkTypeEntity_1.ArtworkTypeEntity(self, entopts);
    }
    // Entity access: `client.CategoryTerm().list()` / `client.CategoryTerm().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    CategoryTerm(entopts) {
        const self = this;
        return new CategoryTermEntity_1.CategoryTermEntity(self, entopts);
    }
    // Entity access: `client.DigitalPublication().list()` / `client.DigitalPublication().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    DigitalPublication(entopts) {
        const self = this;
        return new DigitalPublicationEntity_1.DigitalPublicationEntity(self, entopts);
    }
    // Entity access: `client.DigitalPublicationArticle().list()` / `client.DigitalPublicationArticle().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    DigitalPublicationArticle(entopts) {
        const self = this;
        return new DigitalPublicationArticleEntity_1.DigitalPublicationArticleEntity(self, entopts);
    }
    // Entity access: `client.EducatorResource().list()` / `client.EducatorResource().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EducatorResource(entopts) {
        const self = this;
        return new EducatorResourceEntity_1.EducatorResourceEntity(self, entopts);
    }
    // Entity access: `client.Event().list()` / `client.Event().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Event(entopts) {
        const self = this;
        return new EventEntity_1.EventEntity(self, entopts);
    }
    // Entity access: `client.EventOccurrence().list()` / `client.EventOccurrence().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EventOccurrence(entopts) {
        const self = this;
        return new EventOccurrenceEntity_1.EventOccurrenceEntity(self, entopts);
    }
    // Entity access: `client.EventProgram().list()` / `client.EventProgram().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EventProgram(entopts) {
        const self = this;
        return new EventProgramEntity_1.EventProgramEntity(self, entopts);
    }
    // Entity access: `client.Exhibition().list()` / `client.Exhibition().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Exhibition(entopts) {
        const self = this;
        return new ExhibitionEntity_1.ExhibitionEntity(self, entopts);
    }
    // Entity access: `client.Gallery().list()` / `client.Gallery().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Gallery(entopts) {
        const self = this;
        return new GalleryEntity_1.GalleryEntity(self, entopts);
    }
    // Entity access: `client.GenericPage().list()` / `client.GenericPage().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    GenericPage(entopts) {
        const self = this;
        return new GenericPageEntity_1.GenericPageEntity(self, entopts);
    }
    // Entity access: `client.Highlight().list()` / `client.Highlight().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Highlight(entopts) {
        const self = this;
        return new HighlightEntity_1.HighlightEntity(self, entopts);
    }
    // Entity access: `client.Hour().list()` / `client.Hour().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Hour(entopts) {
        const self = this;
        return new HourEntity_1.HourEntity(self, entopts);
    }
    // Entity access: `client.Image().list()` / `client.Image().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Image(entopts) {
        const self = this;
        return new ImageEntity_1.ImageEntity(self, entopts);
    }
    // Entity access: `client.LandingPage().list()` / `client.LandingPage().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    LandingPage(entopts) {
        const self = this;
        return new LandingPageEntity_1.LandingPageEntity(self, entopts);
    }
    // Entity access: `client.Place().list()` / `client.Place().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Place(entopts) {
        const self = this;
        return new PlaceEntity_1.PlaceEntity(self, entopts);
    }
    // Entity access: `client.PressRelease().list()` / `client.PressRelease().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PressRelease(entopts) {
        const self = this;
        return new PressReleaseEntity_1.PressReleaseEntity(self, entopts);
    }
    // Entity access: `client.PrintedPublication().list()` / `client.PrintedPublication().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PrintedPublication(entopts) {
        const self = this;
        return new PrintedPublicationEntity_1.PrintedPublicationEntity(self, entopts);
    }
    // Entity access: `client.Product().list()` / `client.Product().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Product(entopts) {
        const self = this;
        return new ProductEntity_1.ProductEntity(self, entopts);
    }
    // Entity access: `client.Publication().list()` / `client.Publication().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Publication(entopts) {
        const self = this;
        return new PublicationEntity_1.PublicationEntity(self, entopts);
    }
    // Entity access: `client.Search().list()` / `client.Search().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Search(entopts) {
        const self = this;
        return new SearchEntity_1.SearchEntity(self, entopts);
    }
    // Entity access: `client.Section().list()` / `client.Section().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Section(entopts) {
        const self = this;
        return new SectionEntity_1.SectionEntity(self, entopts);
    }
    // Entity access: `client.Site().list()` / `client.Site().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Site(entopts) {
        const self = this;
        return new SiteEntity_1.SiteEntity(self, entopts);
    }
    // Entity access: `client.Sound().list()` / `client.Sound().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Sound(entopts) {
        const self = this;
        return new SoundEntity_1.SoundEntity(self, entopts);
    }
    // Entity access: `client.StaticPage().list()` / `client.StaticPage().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    StaticPage(entopts) {
        const self = this;
        return new StaticPageEntity_1.StaticPageEntity(self, entopts);
    }
    // Entity access: `client.Text().list()` / `client.Text().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Text(entopts) {
        const self = this;
        return new TextEntity_1.TextEntity(self, entopts);
    }
    // Entity access: `client.Tour().list()` / `client.Tour().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Tour(entopts) {
        const self = this;
        return new TourEntity_1.TourEntity(self, entopts);
    }
    // Entity access: `client.Video().list()` / `client.Video().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Video(entopts) {
        const self = this;
        return new VideoEntity_1.VideoEntity(self, entopts);
    }
    static test(testoptsarg, sdkoptsarg) {
        const struct = stdutil.struct;
        const setpath = struct.setpath;
        const getdef = struct.getdef;
        const clone = struct.clone;
        const setprop = struct.setprop;
        const sdkopts = getdef(clone(sdkoptsarg), {});
        const testopts = getdef(clone(testoptsarg), {});
        setprop(testopts, 'active', true);
        setpath(sdkopts, 'feature.test', testopts);
        const testsdk = new ArtInstituteOfChicagoSDK(sdkopts);
        testsdk._mode = 'test';
        return testsdk;
    }
    tester(testopts, sdkopts) {
        return ArtInstituteOfChicagoSDK.test(testopts, sdkopts);
    }
    toJSON() {
        return { name: 'ArtInstituteOfChicago' };
    }
    toString() {
        return 'ArtInstituteOfChicago ' + this._utility.struct.jsonify(this.toJSON());
    }
    [node_util_1.inspect.custom]() {
        return this.toString();
    }
}
exports.ArtInstituteOfChicagoSDK = ArtInstituteOfChicagoSDK;
const SDK = ArtInstituteOfChicagoSDK;
exports.SDK = SDK;
//# sourceMappingURL=ArtInstituteOfChicagoSDK.js.map