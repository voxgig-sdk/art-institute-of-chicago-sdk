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
    const items = struct.items

    if (true === getpath(this._options.feature, 'test.active')) {
      this._mode = 'test'
    }

    this._rootctx.options = this._options

    this._features = []

    const featureAdd = this._utility.featureAdd
    const featureInit = this._utility.featureInit

    items(this._options.feature, (fitem: [string, any]) => {
      const fname = fitem[0]
      const fopts = fitem[1]
      if (fopts.active) {
        featureAdd(this._rootctx, this._rootctx.config.makeFeature(fname))
      }
    })

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


  async direct(fetchargs?: any) {
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



  // Entity access: `client.Agent().list()` / `client.Agent().load({ id })`.
  Agent(data?: any) {
    const self = this
    return new AgentEntity(self,data)
  }


  // Entity access: `client.AgentRole().list()` / `client.AgentRole().load({ id })`.
  AgentRole(data?: any) {
    const self = this
    return new AgentRoleEntity(self,data)
  }


  // Entity access: `client.AgentType().list()` / `client.AgentType().load({ id })`.
  AgentType(data?: any) {
    const self = this
    return new AgentTypeEntity(self,data)
  }


  // Entity access: `client.Article().list()` / `client.Article().load({ id })`.
  Article(data?: any) {
    const self = this
    return new ArticleEntity(self,data)
  }


  // Entity access: `client.Artwork().list()` / `client.Artwork().load({ id })`.
  Artwork(data?: any) {
    const self = this
    return new ArtworkEntity(self,data)
  }


  // Entity access: `client.ArtworkDateQualifier().list()` / `client.ArtworkDateQualifier().load({ id })`.
  ArtworkDateQualifier(data?: any) {
    const self = this
    return new ArtworkDateQualifierEntity(self,data)
  }


  // Entity access: `client.ArtworkPlaceQualifier().list()` / `client.ArtworkPlaceQualifier().load({ id })`.
  ArtworkPlaceQualifier(data?: any) {
    const self = this
    return new ArtworkPlaceQualifierEntity(self,data)
  }


  // Entity access: `client.ArtworkType().list()` / `client.ArtworkType().load({ id })`.
  ArtworkType(data?: any) {
    const self = this
    return new ArtworkTypeEntity(self,data)
  }


  // Entity access: `client.CategoryTerm().list()` / `client.CategoryTerm().load({ id })`.
  CategoryTerm(data?: any) {
    const self = this
    return new CategoryTermEntity(self,data)
  }


  // Entity access: `client.DigitalPublication().list()` / `client.DigitalPublication().load({ id })`.
  DigitalPublication(data?: any) {
    const self = this
    return new DigitalPublicationEntity(self,data)
  }


  // Entity access: `client.DigitalPublicationArticle().list()` / `client.DigitalPublicationArticle().load({ id })`.
  DigitalPublicationArticle(data?: any) {
    const self = this
    return new DigitalPublicationArticleEntity(self,data)
  }


  // Entity access: `client.EducatorResource().list()` / `client.EducatorResource().load({ id })`.
  EducatorResource(data?: any) {
    const self = this
    return new EducatorResourceEntity(self,data)
  }


  // Entity access: `client.Event().list()` / `client.Event().load({ id })`.
  Event(data?: any) {
    const self = this
    return new EventEntity(self,data)
  }


  // Entity access: `client.EventOccurrence().list()` / `client.EventOccurrence().load({ id })`.
  EventOccurrence(data?: any) {
    const self = this
    return new EventOccurrenceEntity(self,data)
  }


  // Entity access: `client.EventProgram().list()` / `client.EventProgram().load({ id })`.
  EventProgram(data?: any) {
    const self = this
    return new EventProgramEntity(self,data)
  }


  // Entity access: `client.Exhibition().list()` / `client.Exhibition().load({ id })`.
  Exhibition(data?: any) {
    const self = this
    return new ExhibitionEntity(self,data)
  }


  // Entity access: `client.Gallery().list()` / `client.Gallery().load({ id })`.
  Gallery(data?: any) {
    const self = this
    return new GalleryEntity(self,data)
  }


  // Entity access: `client.GenericPage().list()` / `client.GenericPage().load({ id })`.
  GenericPage(data?: any) {
    const self = this
    return new GenericPageEntity(self,data)
  }


  // Entity access: `client.Highlight().list()` / `client.Highlight().load({ id })`.
  Highlight(data?: any) {
    const self = this
    return new HighlightEntity(self,data)
  }


  // Entity access: `client.Hour().list()` / `client.Hour().load({ id })`.
  Hour(data?: any) {
    const self = this
    return new HourEntity(self,data)
  }


  // Entity access: `client.Image().list()` / `client.Image().load({ id })`.
  Image(data?: any) {
    const self = this
    return new ImageEntity(self,data)
  }


  // Entity access: `client.LandingPage().list()` / `client.LandingPage().load({ id })`.
  LandingPage(data?: any) {
    const self = this
    return new LandingPageEntity(self,data)
  }


  // Entity access: `client.Place().list()` / `client.Place().load({ id })`.
  Place(data?: any) {
    const self = this
    return new PlaceEntity(self,data)
  }


  // Entity access: `client.PressRelease().list()` / `client.PressRelease().load({ id })`.
  PressRelease(data?: any) {
    const self = this
    return new PressReleaseEntity(self,data)
  }


  // Entity access: `client.PrintedPublication().list()` / `client.PrintedPublication().load({ id })`.
  PrintedPublication(data?: any) {
    const self = this
    return new PrintedPublicationEntity(self,data)
  }


  // Entity access: `client.Product().list()` / `client.Product().load({ id })`.
  Product(data?: any) {
    const self = this
    return new ProductEntity(self,data)
  }


  // Entity access: `client.Publication().list()` / `client.Publication().load({ id })`.
  Publication(data?: any) {
    const self = this
    return new PublicationEntity(self,data)
  }


  // Entity access: `client.Search().list()` / `client.Search().load({ id })`.
  Search(data?: any) {
    const self = this
    return new SearchEntity(self,data)
  }


  // Entity access: `client.Section().list()` / `client.Section().load({ id })`.
  Section(data?: any) {
    const self = this
    return new SectionEntity(self,data)
  }


  // Entity access: `client.Site().list()` / `client.Site().load({ id })`.
  Site(data?: any) {
    const self = this
    return new SiteEntity(self,data)
  }


  // Entity access: `client.Sound().list()` / `client.Sound().load({ id })`.
  Sound(data?: any) {
    const self = this
    return new SoundEntity(self,data)
  }


  // Entity access: `client.StaticPage().list()` / `client.StaticPage().load({ id })`.
  StaticPage(data?: any) {
    const self = this
    return new StaticPageEntity(self,data)
  }


  // Entity access: `client.Text().list()` / `client.Text().load({ id })`.
  Text(data?: any) {
    const self = this
    return new TextEntity(self,data)
  }


  // Entity access: `client.Tour().list()` / `client.Tour().load({ id })`.
  Tour(data?: any) {
    const self = this
    return new TourEntity(self,data)
  }


  // Entity access: `client.Video().list()` / `client.Video().load({ id })`.
  Video(data?: any) {
    const self = this
    return new VideoEntity(self,data)
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

  BaseFeature,
  ArtInstituteOfChicagoEntityBase,

  ArtInstituteOfChicagoSDK,
  SDK,
}


