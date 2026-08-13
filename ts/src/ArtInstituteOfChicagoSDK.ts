// ArtInstituteOfChicago Ts SDK

import { AgentEntity } from './entity/AgentEntity'
import { AgentRoleEntity } from './entity/AgentRoleEntity'
import { AgentTypeEntity } from './entity/AgentTypeEntity'
import { ArticleEntity } from './entity/ArticleEntity'
import { ArtworkEntity } from './entity/ArtworkEntity'
import { ArtworkDateQualifierEntity } from './entity/ArtworkDateQualifierEntity'
import { ArtworkPlaceQualifierEntity } from './entity/ArtworkPlaceQualifierEntity'
import { ArtworkTypeEntity } from './entity/ArtworkTypeEntity'
import { CategoryTermEntity } from './entity/CategoryTermEntity'
import { DigitalPublicationEntity } from './entity/DigitalPublicationEntity'
import { DigitalPublicationArticleEntity } from './entity/DigitalPublicationArticleEntity'
import { EducatorResourceEntity } from './entity/EducatorResourceEntity'
import { EventEntity } from './entity/EventEntity'
import { EventOccurrenceEntity } from './entity/EventOccurrenceEntity'
import { EventProgramEntity } from './entity/EventProgramEntity'
import { ExhibitionEntity } from './entity/ExhibitionEntity'
import { GalleryEntity } from './entity/GalleryEntity'
import { GenericPageEntity } from './entity/GenericPageEntity'
import { HighlightEntity } from './entity/HighlightEntity'
import { HourEntity } from './entity/HourEntity'
import { ImageEntity } from './entity/ImageEntity'
import { LandingPageEntity } from './entity/LandingPageEntity'
import { PlaceEntity } from './entity/PlaceEntity'
import { PressReleaseEntity } from './entity/PressReleaseEntity'
import { PrintedPublicationEntity } from './entity/PrintedPublicationEntity'
import { ProductEntity } from './entity/ProductEntity'
import { PublicationEntity } from './entity/PublicationEntity'
import { SearchEntity } from './entity/SearchEntity'
import { SectionEntity } from './entity/SectionEntity'
import { SiteEntity } from './entity/SiteEntity'
import { SoundEntity } from './entity/SoundEntity'
import { StaticPageEntity } from './entity/StaticPageEntity'
import { TextEntity } from './entity/TextEntity'
import { TourEntity } from './entity/TourEntity'
import { VideoEntity } from './entity/VideoEntity'

export type * from './ArtInstituteOfChicagoTypes'


import { inspect } from 'node:util'

import type { Context, Feature } from './types'

import { config } from './Config'
import { ArtInstituteOfChicagoEntityBase } from './ArtInstituteOfChicagoEntityBase'
import { Utility } from './utility/Utility'


import { BaseFeature } from './feature/base/BaseFeature'


const stdutil = new Utility()


class ArtInstituteOfChicagoSDK {
  _mode: string = 'live'
  _options: any
  _utility = new Utility()
  _features: Feature[]
  _rootctx: Context

  constructor(options?: any) {

    this._rootctx = this._utility.makeContext({
      client: this,
      utility: this._utility,
      config,
      options,
      shared: new WeakMap()
    })

    this._options = this._utility.makeOptions(this._rootctx)

    const struct = this._utility.struct
    const getpath = struct.getpath

    if (true === getpath(this._options.feature, 'test.active')) {
      this._mode = 'test'
    }

    this._rootctx.options = this._options

    this._features = []

    const featureAdd = this._utility.featureAdd
    const featureInit = this._utility.featureInit

    // Add features in the resolved order (makeOptions puts an explicit
    // array order first, else defaults to test-first). Ordering matters:
    // the `test` feature installs the base mock transport and the transport
    // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is current,
    // so `test` must be added before them to sit at the base of the chain.
    const featureorder = getpath(this._options, '__derived__.featureorder') || []
    for (const fname of featureorder) {
      const fopts = this._options.feature[fname] || {}
      if (fopts.active) {
        featureAdd(this._rootctx, this._rootctx.config.makeFeature(fname))
      }
    }

    if (null != this._options.extend) {
      for (let f of this._options.extend) {
        featureAdd(this._rootctx, f)
      }
    }

    for (let f of this._features) {
      featureInit(this._rootctx, f)
    }

    const featureHook = this._utility.featureHook
    featureHook(this._rootctx, 'PostConstruct')
  }


  options() {
    return this._utility.struct.clone(this._options)
  }


  utility() {
    return this._utility.struct.clone(this._utility)
  }


  async prepare(fetchargs?: any) {
    const utility = this._utility
    const struct = utility.struct
    const clone = struct.clone

    const {
      makeContext,
      makeFetchDef,
      prepareHeaders,
      prepareAuth,
    } = utility

    fetchargs = fetchargs || {}

    let ctx: Context = makeContext({
      opname: 'prepare',
      ctrl: fetchargs.ctrl || {},
    }, this._rootctx)

    const options = this._options

    // Build spec directly from SDK options + user-provided fetch args.
    const spec: any = {
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
    }

    ctx.spec = spec

    // Merge user-provided headers over SDK defaults.
    if (fetchargs.headers) {
      const uheaders = fetchargs.headers
      for (let key in uheaders) {
        spec.headers[key] = uheaders[key]
      }
    }

    // Apply SDK auth (apikey, auth prefix, etc.)
    const authResult = prepareAuth(ctx)
    if (authResult instanceof Error) {
      return authResult
    }

    return makeFetchDef(ctx)
  }


  // Raw endpoint access is operator-controllable, like every entity op.
  // Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
  // either one reaches the same endpoint.
  async direct(fetchargs?: any) {
    if (!this._options.allow.op.includes('direct')) {
      return {
        ok: false,
        err: new Error('ArtInstituteOfChicagoSDK: direct: operation not allowed by' +
          ' SDK option allow.op value: "' + this._options.allow.op + '"'),
      }
    }

    return this._rawRequest(fetchargs)
  }


  // Ungated request path shared by direct() and graphql(), each of which
  // checks its own allow.op token first. Private, rather than a flag on
  // fetchargs: a caller-supplied marker would let anyone opt straight back
  // out of the gate by passing it.
  async _rawRequest(fetchargs?: any) {
    const utility = this._utility

    const fetcher = utility.fetcher
    const makeContext = utility.makeContext

    const fetchdef = await this.prepare(fetchargs)
    if (fetchdef instanceof Error) {
      return fetchdef
    }

    let ctx: Context = makeContext({
      opname: 'direct',
      ctrl: (fetchargs || {}).ctrl || {},
    }, this._rootctx)

    try {
      const fetched = await fetcher(ctx, fetchdef.url, fetchdef)

      if (null == fetched) {
        return { ok: false, err: ctx.error('direct_no_response', 'response: undefined') }
      }
      else if (fetched instanceof Error) {
        return { ok: false, err: fetched }
      }

      const status = fetched.status

      // No body responses (204 No Content, 304 Not Modified) and explicit
      // zero content-length must skip JSON parsing — fetched.json() would
      // throw `Unexpected end of JSON input` on an empty body.
      const headers = fetched.headers
      const contentLength = headers && 'function' === typeof headers.get
        ? headers.get('content-length')
        : (headers || {})['content-length']
      const noBody = 204 === status || 304 === status || '0' === String(contentLength)

      let json: any = undefined
      if (!noBody) {
        try {
          json = 'function' === typeof fetched.json ? await fetched.json() : fetched.json
        }
        catch (parseErr) {
          // Body wasn't valid JSON — surface the raw response rather than
          // throwing. data stays undefined; callers can inspect status/headers.
          json = undefined
        }
      }

      return {
        ok: status >= 200 && status < 300,
        status,
        headers: fetched.headers,
        data: json,
      }
    }
    catch (err: any) {
      return { ok: false, err }
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
  async graphql(query: string, variables?: any, ctrl?: any) {
    const options = this._options

    if (!options.allow.op.includes('graphql')) {
      return {
        ok: false,
        err: new Error('ArtInstituteOfChicagoSDK: graphql: operation not allowed by' +
          ' SDK option allow.op value: "' + options.allow.op + '"'),
      }
    }

    const res: any = await this._rawRequest({
      method: 'POST',
      headers: { 'content-type': 'application/json' },
      body: { query, variables: variables || {} },
      ctrl,
    })

    if (res instanceof Error) {
      return res
    }

    // Errors are read BEFORE any status check: a GraphQL parse or validation
    // failure comes back as HTTP 400 carrying the standard { errors: [...] }
    // body, and the raw path represents a non-2xx as { ok: false } with no
    // err — so returning early on status would discard the server's own
    // diagnostics, which are the only useful part of that response.
    const errors = null == res.data ? undefined : res.data.errors

    if (null != errors && Array.isArray(errors) && 0 < errors.length) {
      const first = errors[0] || {}
      const err: any = new Error('ArtInstituteOfChicagoSDK: graphql: ' +
        (first.message || 'graphql error'))
      err.graphql = errors
      return { ok: false, status: res.status, headers: res.headers, err, data: res.data }
    }

    return res
  }



  // Entity access: `client.Agent().list()` / `client.Agent().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Agent(entopts?: Record<string, any>) {
    const self = this
    return new AgentEntity(self, entopts)
  }


  // Entity access: `client.AgentRole().list()` / `client.AgentRole().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  AgentRole(entopts?: Record<string, any>) {
    const self = this
    return new AgentRoleEntity(self, entopts)
  }


  // Entity access: `client.AgentType().list()` / `client.AgentType().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  AgentType(entopts?: Record<string, any>) {
    const self = this
    return new AgentTypeEntity(self, entopts)
  }


  // Entity access: `client.Article().list()` / `client.Article().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Article(entopts?: Record<string, any>) {
    const self = this
    return new ArticleEntity(self, entopts)
  }


  // Entity access: `client.Artwork().list()` / `client.Artwork().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Artwork(entopts?: Record<string, any>) {
    const self = this
    return new ArtworkEntity(self, entopts)
  }


  // Entity access: `client.ArtworkDateQualifier().list()` / `client.ArtworkDateQualifier().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ArtworkDateQualifier(entopts?: Record<string, any>) {
    const self = this
    return new ArtworkDateQualifierEntity(self, entopts)
  }


  // Entity access: `client.ArtworkPlaceQualifier().list()` / `client.ArtworkPlaceQualifier().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ArtworkPlaceQualifier(entopts?: Record<string, any>) {
    const self = this
    return new ArtworkPlaceQualifierEntity(self, entopts)
  }


  // Entity access: `client.ArtworkType().list()` / `client.ArtworkType().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ArtworkType(entopts?: Record<string, any>) {
    const self = this
    return new ArtworkTypeEntity(self, entopts)
  }


  // Entity access: `client.CategoryTerm().list()` / `client.CategoryTerm().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  CategoryTerm(entopts?: Record<string, any>) {
    const self = this
    return new CategoryTermEntity(self, entopts)
  }


  // Entity access: `client.DigitalPublication().list()` / `client.DigitalPublication().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  DigitalPublication(entopts?: Record<string, any>) {
    const self = this
    return new DigitalPublicationEntity(self, entopts)
  }


  // Entity access: `client.DigitalPublicationArticle().list()` / `client.DigitalPublicationArticle().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  DigitalPublicationArticle(entopts?: Record<string, any>) {
    const self = this
    return new DigitalPublicationArticleEntity(self, entopts)
  }


  // Entity access: `client.EducatorResource().list()` / `client.EducatorResource().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EducatorResource(entopts?: Record<string, any>) {
    const self = this
    return new EducatorResourceEntity(self, entopts)
  }


  // Entity access: `client.Event().list()` / `client.Event().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Event(entopts?: Record<string, any>) {
    const self = this
    return new EventEntity(self, entopts)
  }


  // Entity access: `client.EventOccurrence().list()` / `client.EventOccurrence().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EventOccurrence(entopts?: Record<string, any>) {
    const self = this
    return new EventOccurrenceEntity(self, entopts)
  }


  // Entity access: `client.EventProgram().list()` / `client.EventProgram().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EventProgram(entopts?: Record<string, any>) {
    const self = this
    return new EventProgramEntity(self, entopts)
  }


  // Entity access: `client.Exhibition().list()` / `client.Exhibition().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Exhibition(entopts?: Record<string, any>) {
    const self = this
    return new ExhibitionEntity(self, entopts)
  }


  // Entity access: `client.Gallery().list()` / `client.Gallery().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Gallery(entopts?: Record<string, any>) {
    const self = this
    return new GalleryEntity(self, entopts)
  }


  // Entity access: `client.GenericPage().list()` / `client.GenericPage().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  GenericPage(entopts?: Record<string, any>) {
    const self = this
    return new GenericPageEntity(self, entopts)
  }


  // Entity access: `client.Highlight().list()` / `client.Highlight().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Highlight(entopts?: Record<string, any>) {
    const self = this
    return new HighlightEntity(self, entopts)
  }


  // Entity access: `client.Hour().list()` / `client.Hour().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Hour(entopts?: Record<string, any>) {
    const self = this
    return new HourEntity(self, entopts)
  }


  // Entity access: `client.Image().list()` / `client.Image().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Image(entopts?: Record<string, any>) {
    const self = this
    return new ImageEntity(self, entopts)
  }


  // Entity access: `client.LandingPage().list()` / `client.LandingPage().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  LandingPage(entopts?: Record<string, any>) {
    const self = this
    return new LandingPageEntity(self, entopts)
  }


  // Entity access: `client.Place().list()` / `client.Place().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Place(entopts?: Record<string, any>) {
    const self = this
    return new PlaceEntity(self, entopts)
  }


  // Entity access: `client.PressRelease().list()` / `client.PressRelease().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PressRelease(entopts?: Record<string, any>) {
    const self = this
    return new PressReleaseEntity(self, entopts)
  }


  // Entity access: `client.PrintedPublication().list()` / `client.PrintedPublication().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PrintedPublication(entopts?: Record<string, any>) {
    const self = this
    return new PrintedPublicationEntity(self, entopts)
  }


  // Entity access: `client.Product().list()` / `client.Product().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Product(entopts?: Record<string, any>) {
    const self = this
    return new ProductEntity(self, entopts)
  }


  // Entity access: `client.Publication().list()` / `client.Publication().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Publication(entopts?: Record<string, any>) {
    const self = this
    return new PublicationEntity(self, entopts)
  }


  // Entity access: `client.Search().list()` / `client.Search().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Search(entopts?: Record<string, any>) {
    const self = this
    return new SearchEntity(self, entopts)
  }


  // Entity access: `client.Section().list()` / `client.Section().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Section(entopts?: Record<string, any>) {
    const self = this
    return new SectionEntity(self, entopts)
  }


  // Entity access: `client.Site().list()` / `client.Site().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Site(entopts?: Record<string, any>) {
    const self = this
    return new SiteEntity(self, entopts)
  }


  // Entity access: `client.Sound().list()` / `client.Sound().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Sound(entopts?: Record<string, any>) {
    const self = this
    return new SoundEntity(self, entopts)
  }


  // Entity access: `client.StaticPage().list()` / `client.StaticPage().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  StaticPage(entopts?: Record<string, any>) {
    const self = this
    return new StaticPageEntity(self, entopts)
  }


  // Entity access: `client.Text().list()` / `client.Text().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Text(entopts?: Record<string, any>) {
    const self = this
    return new TextEntity(self, entopts)
  }


  // Entity access: `client.Tour().list()` / `client.Tour().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Tour(entopts?: Record<string, any>) {
    const self = this
    return new TourEntity(self, entopts)
  }


  // Entity access: `client.Video().list()` / `client.Video().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Video(entopts?: Record<string, any>) {
    const self = this
    return new VideoEntity(self, entopts)
  }




  static test(testoptsarg?: any, sdkoptsarg?: any) {
    const struct = stdutil.struct
    const setpath = struct.setpath
    const getdef = struct.getdef
    const clone = struct.clone
    const setprop = struct.setprop

    const sdkopts = getdef(clone(sdkoptsarg), {})
    const testopts = getdef(clone(testoptsarg), {})
    setprop(testopts, 'active', true)
    setpath(sdkopts, 'feature.test', testopts)

    const testsdk = new ArtInstituteOfChicagoSDK(sdkopts)
    testsdk._mode = 'test'

    return testsdk
  }


  tester(testopts?: any, sdkopts?: any) {
    return ArtInstituteOfChicagoSDK.test(testopts, sdkopts)
  }


  toJSON() {
    return { name: 'ArtInstituteOfChicago' }
  }

  toString() {
    return 'ArtInstituteOfChicago ' + this._utility.struct.jsonify(this.toJSON())
  }

  [inspect.custom]() {
    return this.toString()
  }

}




const SDK = ArtInstituteOfChicagoSDK


export {
  stdutil,
  config,

  BaseFeature,
  ArtInstituteOfChicagoEntityBase,

  ArtInstituteOfChicagoSDK,
  SDK,
}


