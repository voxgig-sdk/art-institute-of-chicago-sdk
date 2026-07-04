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



  _agent?: AgentEntity

  // Idiomatic facade: `client.agent.list()` / `client.agent.load({ id })`.
  get agent(): AgentEntity {
    return (this._agent ??= new AgentEntity(this, undefined))
  }

  /** @deprecated Use `client.agent` instead. */
  Agent(data?: any) {
    const self = this
    return new AgentEntity(self,data)
  }


  _agent_role?: AgentRoleEntity

  // Idiomatic facade: `client.agent_role.list()` / `client.agent_role.load({ id })`.
  get agent_role(): AgentRoleEntity {
    return (this._agent_role ??= new AgentRoleEntity(this, undefined))
  }

  /** @deprecated Use `client.agent_role` instead. */
  AgentRole(data?: any) {
    const self = this
    return new AgentRoleEntity(self,data)
  }


  _agent_type?: AgentTypeEntity

  // Idiomatic facade: `client.agent_type.list()` / `client.agent_type.load({ id })`.
  get agent_type(): AgentTypeEntity {
    return (this._agent_type ??= new AgentTypeEntity(this, undefined))
  }

  /** @deprecated Use `client.agent_type` instead. */
  AgentType(data?: any) {
    const self = this
    return new AgentTypeEntity(self,data)
  }


  _article?: ArticleEntity

  // Idiomatic facade: `client.article.list()` / `client.article.load({ id })`.
  get article(): ArticleEntity {
    return (this._article ??= new ArticleEntity(this, undefined))
  }

  /** @deprecated Use `client.article` instead. */
  Article(data?: any) {
    const self = this
    return new ArticleEntity(self,data)
  }


  _artwork?: ArtworkEntity

  // Idiomatic facade: `client.artwork.list()` / `client.artwork.load({ id })`.
  get artwork(): ArtworkEntity {
    return (this._artwork ??= new ArtworkEntity(this, undefined))
  }

  /** @deprecated Use `client.artwork` instead. */
  Artwork(data?: any) {
    const self = this
    return new ArtworkEntity(self,data)
  }


  _artwork_date_qualifier?: ArtworkDateQualifierEntity

  // Idiomatic facade: `client.artwork_date_qualifier.list()` / `client.artwork_date_qualifier.load({ id })`.
  get artwork_date_qualifier(): ArtworkDateQualifierEntity {
    return (this._artwork_date_qualifier ??= new ArtworkDateQualifierEntity(this, undefined))
  }

  /** @deprecated Use `client.artwork_date_qualifier` instead. */
  ArtworkDateQualifier(data?: any) {
    const self = this
    return new ArtworkDateQualifierEntity(self,data)
  }


  _artwork_place_qualifier?: ArtworkPlaceQualifierEntity

  // Idiomatic facade: `client.artwork_place_qualifier.list()` / `client.artwork_place_qualifier.load({ id })`.
  get artwork_place_qualifier(): ArtworkPlaceQualifierEntity {
    return (this._artwork_place_qualifier ??= new ArtworkPlaceQualifierEntity(this, undefined))
  }

  /** @deprecated Use `client.artwork_place_qualifier` instead. */
  ArtworkPlaceQualifier(data?: any) {
    const self = this
    return new ArtworkPlaceQualifierEntity(self,data)
  }


  _artwork_type?: ArtworkTypeEntity

  // Idiomatic facade: `client.artwork_type.list()` / `client.artwork_type.load({ id })`.
  get artwork_type(): ArtworkTypeEntity {
    return (this._artwork_type ??= new ArtworkTypeEntity(this, undefined))
  }

  /** @deprecated Use `client.artwork_type` instead. */
  ArtworkType(data?: any) {
    const self = this
    return new ArtworkTypeEntity(self,data)
  }


  _category_term?: CategoryTermEntity

  // Idiomatic facade: `client.category_term.list()` / `client.category_term.load({ id })`.
  get category_term(): CategoryTermEntity {
    return (this._category_term ??= new CategoryTermEntity(this, undefined))
  }

  /** @deprecated Use `client.category_term` instead. */
  CategoryTerm(data?: any) {
    const self = this
    return new CategoryTermEntity(self,data)
  }


  _digital_publication?: DigitalPublicationEntity

  // Idiomatic facade: `client.digital_publication.list()` / `client.digital_publication.load({ id })`.
  get digital_publication(): DigitalPublicationEntity {
    return (this._digital_publication ??= new DigitalPublicationEntity(this, undefined))
  }

  /** @deprecated Use `client.digital_publication` instead. */
  DigitalPublication(data?: any) {
    const self = this
    return new DigitalPublicationEntity(self,data)
  }


  _digital_publication_article?: DigitalPublicationArticleEntity

  // Idiomatic facade: `client.digital_publication_article.list()` / `client.digital_publication_article.load({ id })`.
  get digital_publication_article(): DigitalPublicationArticleEntity {
    return (this._digital_publication_article ??= new DigitalPublicationArticleEntity(this, undefined))
  }

  /** @deprecated Use `client.digital_publication_article` instead. */
  DigitalPublicationArticle(data?: any) {
    const self = this
    return new DigitalPublicationArticleEntity(self,data)
  }


  _educator_resource?: EducatorResourceEntity

  // Idiomatic facade: `client.educator_resource.list()` / `client.educator_resource.load({ id })`.
  get educator_resource(): EducatorResourceEntity {
    return (this._educator_resource ??= new EducatorResourceEntity(this, undefined))
  }

  /** @deprecated Use `client.educator_resource` instead. */
  EducatorResource(data?: any) {
    const self = this
    return new EducatorResourceEntity(self,data)
  }


  _event?: EventEntity

  // Idiomatic facade: `client.event.list()` / `client.event.load({ id })`.
  get event(): EventEntity {
    return (this._event ??= new EventEntity(this, undefined))
  }

  /** @deprecated Use `client.event` instead. */
  Event(data?: any) {
    const self = this
    return new EventEntity(self,data)
  }


  _event_occurrence?: EventOccurrenceEntity

  // Idiomatic facade: `client.event_occurrence.list()` / `client.event_occurrence.load({ id })`.
  get event_occurrence(): EventOccurrenceEntity {
    return (this._event_occurrence ??= new EventOccurrenceEntity(this, undefined))
  }

  /** @deprecated Use `client.event_occurrence` instead. */
  EventOccurrence(data?: any) {
    const self = this
    return new EventOccurrenceEntity(self,data)
  }


  _event_program?: EventProgramEntity

  // Idiomatic facade: `client.event_program.list()` / `client.event_program.load({ id })`.
  get event_program(): EventProgramEntity {
    return (this._event_program ??= new EventProgramEntity(this, undefined))
  }

  /** @deprecated Use `client.event_program` instead. */
  EventProgram(data?: any) {
    const self = this
    return new EventProgramEntity(self,data)
  }


  _exhibition?: ExhibitionEntity

  // Idiomatic facade: `client.exhibition.list()` / `client.exhibition.load({ id })`.
  get exhibition(): ExhibitionEntity {
    return (this._exhibition ??= new ExhibitionEntity(this, undefined))
  }

  /** @deprecated Use `client.exhibition` instead. */
  Exhibition(data?: any) {
    const self = this
    return new ExhibitionEntity(self,data)
  }


  _gallery?: GalleryEntity

  // Idiomatic facade: `client.gallery.list()` / `client.gallery.load({ id })`.
  get gallery(): GalleryEntity {
    return (this._gallery ??= new GalleryEntity(this, undefined))
  }

  /** @deprecated Use `client.gallery` instead. */
  Gallery(data?: any) {
    const self = this
    return new GalleryEntity(self,data)
  }


  _generic_page?: GenericPageEntity

  // Idiomatic facade: `client.generic_page.list()` / `client.generic_page.load({ id })`.
  get generic_page(): GenericPageEntity {
    return (this._generic_page ??= new GenericPageEntity(this, undefined))
  }

  /** @deprecated Use `client.generic_page` instead. */
  GenericPage(data?: any) {
    const self = this
    return new GenericPageEntity(self,data)
  }


  _highlight?: HighlightEntity

  // Idiomatic facade: `client.highlight.list()` / `client.highlight.load({ id })`.
  get highlight(): HighlightEntity {
    return (this._highlight ??= new HighlightEntity(this, undefined))
  }

  /** @deprecated Use `client.highlight` instead. */
  Highlight(data?: any) {
    const self = this
    return new HighlightEntity(self,data)
  }


  _hour?: HourEntity

  // Idiomatic facade: `client.hour.list()` / `client.hour.load({ id })`.
  get hour(): HourEntity {
    return (this._hour ??= new HourEntity(this, undefined))
  }

  /** @deprecated Use `client.hour` instead. */
  Hour(data?: any) {
    const self = this
    return new HourEntity(self,data)
  }


  _image?: ImageEntity

  // Idiomatic facade: `client.image.list()` / `client.image.load({ id })`.
  get image(): ImageEntity {
    return (this._image ??= new ImageEntity(this, undefined))
  }

  /** @deprecated Use `client.image` instead. */
  Image(data?: any) {
    const self = this
    return new ImageEntity(self,data)
  }


  _landing_page?: LandingPageEntity

  // Idiomatic facade: `client.landing_page.list()` / `client.landing_page.load({ id })`.
  get landing_page(): LandingPageEntity {
    return (this._landing_page ??= new LandingPageEntity(this, undefined))
  }

  /** @deprecated Use `client.landing_page` instead. */
  LandingPage(data?: any) {
    const self = this
    return new LandingPageEntity(self,data)
  }


  _place?: PlaceEntity

  // Idiomatic facade: `client.place.list()` / `client.place.load({ id })`.
  get place(): PlaceEntity {
    return (this._place ??= new PlaceEntity(this, undefined))
  }

  /** @deprecated Use `client.place` instead. */
  Place(data?: any) {
    const self = this
    return new PlaceEntity(self,data)
  }


  _press_release?: PressReleaseEntity

  // Idiomatic facade: `client.press_release.list()` / `client.press_release.load({ id })`.
  get press_release(): PressReleaseEntity {
    return (this._press_release ??= new PressReleaseEntity(this, undefined))
  }

  /** @deprecated Use `client.press_release` instead. */
  PressRelease(data?: any) {
    const self = this
    return new PressReleaseEntity(self,data)
  }


  _printed_publication?: PrintedPublicationEntity

  // Idiomatic facade: `client.printed_publication.list()` / `client.printed_publication.load({ id })`.
  get printed_publication(): PrintedPublicationEntity {
    return (this._printed_publication ??= new PrintedPublicationEntity(this, undefined))
  }

  /** @deprecated Use `client.printed_publication` instead. */
  PrintedPublication(data?: any) {
    const self = this
    return new PrintedPublicationEntity(self,data)
  }


  _product?: ProductEntity

  // Idiomatic facade: `client.product.list()` / `client.product.load({ id })`.
  get product(): ProductEntity {
    return (this._product ??= new ProductEntity(this, undefined))
  }

  /** @deprecated Use `client.product` instead. */
  Product(data?: any) {
    const self = this
    return new ProductEntity(self,data)
  }


  _publication?: PublicationEntity

  // Idiomatic facade: `client.publication.list()` / `client.publication.load({ id })`.
  get publication(): PublicationEntity {
    return (this._publication ??= new PublicationEntity(this, undefined))
  }

  /** @deprecated Use `client.publication` instead. */
  Publication(data?: any) {
    const self = this
    return new PublicationEntity(self,data)
  }


  _search?: SearchEntity

  // Idiomatic facade: `client.search.list()` / `client.search.load({ id })`.
  get search(): SearchEntity {
    return (this._search ??= new SearchEntity(this, undefined))
  }

  /** @deprecated Use `client.search` instead. */
  Search(data?: any) {
    const self = this
    return new SearchEntity(self,data)
  }


  _section?: SectionEntity

  // Idiomatic facade: `client.section.list()` / `client.section.load({ id })`.
  get section(): SectionEntity {
    return (this._section ??= new SectionEntity(this, undefined))
  }

  /** @deprecated Use `client.section` instead. */
  Section(data?: any) {
    const self = this
    return new SectionEntity(self,data)
  }


  _site?: SiteEntity

  // Idiomatic facade: `client.site.list()` / `client.site.load({ id })`.
  get site(): SiteEntity {
    return (this._site ??= new SiteEntity(this, undefined))
  }

  /** @deprecated Use `client.site` instead. */
  Site(data?: any) {
    const self = this
    return new SiteEntity(self,data)
  }


  _sound?: SoundEntity

  // Idiomatic facade: `client.sound.list()` / `client.sound.load({ id })`.
  get sound(): SoundEntity {
    return (this._sound ??= new SoundEntity(this, undefined))
  }

  /** @deprecated Use `client.sound` instead. */
  Sound(data?: any) {
    const self = this
    return new SoundEntity(self,data)
  }


  _static_page?: StaticPageEntity

  // Idiomatic facade: `client.static_page.list()` / `client.static_page.load({ id })`.
  get static_page(): StaticPageEntity {
    return (this._static_page ??= new StaticPageEntity(this, undefined))
  }

  /** @deprecated Use `client.static_page` instead. */
  StaticPage(data?: any) {
    const self = this
    return new StaticPageEntity(self,data)
  }


  _text?: TextEntity

  // Idiomatic facade: `client.text.list()` / `client.text.load({ id })`.
  get text(): TextEntity {
    return (this._text ??= new TextEntity(this, undefined))
  }

  /** @deprecated Use `client.text` instead. */
  Text(data?: any) {
    const self = this
    return new TextEntity(self,data)
  }


  _tour?: TourEntity

  // Idiomatic facade: `client.tour.list()` / `client.tour.load({ id })`.
  get tour(): TourEntity {
    return (this._tour ??= new TourEntity(this, undefined))
  }

  /** @deprecated Use `client.tour` instead. */
  Tour(data?: any) {
    const self = this
    return new TourEntity(self,data)
  }


  _video?: VideoEntity

  // Idiomatic facade: `client.video.list()` / `client.video.load({ id })`.
  get video(): VideoEntity {
    return (this._video ??= new VideoEntity(this, undefined))
  }

  /** @deprecated Use `client.video` instead. */
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


