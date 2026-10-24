

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { ArtInstituteOfChicagoSDK, BaseFeature, config, stdutil } from '../../..'

import {
  envOverride,
  liveClientOptions,
  liveDelay,
  loadEnvLocal,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


loadEnvLocal(__dirname + '/../../../.env.local')


describe('ArtworkEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE.
  afterEach(liveDelay('ART_INSTITUTE_OF_CHICAGO_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = ArtInstituteOfChicagoSDK.test()
    const ent = testsdk.Artwork()
    assert(null != ent)
  })


  class FailHook extends BaseFeature {
    name = 'failhook'
    version = '0.0.1'
    active = true
    unexpected = 0
    init() { }
    PreSpec() { throw new Error('artwork hook failed') }
    PreUnexpected() { this.unexpected++ }
  }

  test('stream-error', async () => {
    const offline = { net: { offline: true } }
    await assert.rejects(async () => {
      for await (const _item of ArtInstituteOfChicagoSDK.test(offline).Artwork().stream('list')) { }
    }, /offline/)

    for await (const _item of ArtInstituteOfChicagoSDK.test(offline).Artwork()
      .stream('list', undefined, { ctrl: { throw: false } })) { }

    if (null != (config as any).feature?.rbac) {
      const denied = ArtInstituteOfChicagoSDK.test(undefined, { feature: { rbac: { active: true, deny: true } } })
      await assert.rejects(async () => {
        for await (const _item of denied.Artwork().stream('list')) { }
      }, (err: any) => 'rbac_denied' === err.code)
    }
  })

  test('stream-ctrl', async () => {
    const explain: any = {}
    const ctrl: any = { explain }
    for await (const _item of ArtInstituteOfChicagoSDK.test().Artwork().stream('list', undefined, { ctrl })) { }
    assert.deepStrictEqual(Object.keys(ctrl), ['explain'])
    assert(explain === ctrl.explain && 0 < Object.keys(explain).length)
  })

  test('unexpected', async () => {
    const hook = new FailHook()
    const client = new ArtInstituteOfChicagoSDK({ feature: { test: { active: true } }, extend: [hook] })
    await assert.rejects(client.Artwork().list(), /hook failed/)
    assert(0 < hook.unexpected)

    const fired = hook.unexpected
    assert.strictEqual(await client.Artwork().list(undefined, { throw: false }), undefined)
    assert(fired < hook.unexpected)
  })

  test('validate', async (t) => {
    if (null == (config as any).feature?.validate) {
      t.skip('feature not present in this SDK: validate')
      return
    }
    const client = ArtInstituteOfChicagoSDK.test(undefined, { feature: { validate: { active: true } } })
    await assert.rejects(client.Artwork().list({"artist_id":1} as any),
      (err: any) => 'validate_failed' === err.code)
  })



  test('basic', async (t) => {

    const live = 'TRUE' === process.env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE
    for (const op of ['list', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'artwork.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"alt_artist_ids":{"a":true,"h":"Alt Artist Ids","n":"alt_artist_ids","r":false,"sh":"Unique identifiers of the non-preferred artists/cultures associated with this work","t":"`$ANY`","key$":"alt_artist_ids","index$":0},"alt_classification_ids":{"a":true,"h":"Alt Classification Ids","n":"alt_classification_ids","r":false,"sh":"Unique identifiers of all other non-preferred classification terms for this work","t":"`$ANY`","key$":"alt_classification_ids","index$":1},"alt_image_ids":{"a":true,"h":"Alt Image Ids","n":"alt_image_ids","r":false,"sh":"Unique identifiers of all non-preferred images of this work.","t":"`$ANY`","key$":"alt_image_ids","index$":2},"alt_material_ids":{"a":true,"h":"Alt Material Ids","n":"alt_material_ids","r":false,"sh":"Unique identifiers of all other non-preferred material terms for this work","t":"`$ANY`","key$":"alt_material_ids","index$":3},"alt_style_ids":{"a":true,"h":"Alt Style Ids","n":"alt_style_ids","r":false,"sh":"Unique identifiers of all other non-preferred style terms for this work","t":"`$ANY`","key$":"alt_style_ids","index$":4},"alt_subject_ids":{"a":true,"h":"Alt Subject Ids","n":"alt_subject_ids","r":false,"sh":"Unique identifiers of all other non-preferred subject terms for this work","t":"`$ANY`","key$":"alt_subject_ids","index$":5},"alt_technique_ids":{"a":true,"h":"Alt Technique Ids","n":"alt_technique_ids","r":false,"sh":"Unique identifiers of all other non-preferred technique terms for this work","t":"`$ANY`","key$":"alt_technique_ids","index$":6},"alt_titles":{"a":true,"h":"Alt Titles","n":"alt_titles","r":false,"sh":"Alternate names for this work","t":"`$ANY`","key$":"alt_titles","index$":7},"api_link":{"a":true,"h":"Api Link","n":"api_link","r":false,"sh":"REST API link for this resource","t":"`$ANY`","key$":"api_link","index$":8},"api_model":{"a":true,"h":"Api Model","n":"api_model","r":false,"sh":"REST API resource type or endpoint","t":"`$ANY`","key$":"api_model","index$":9},"artist_display":{"a":true,"h":"Artist Display","n":"artist_display","r":false,"sh":"Readable description of the creator of this work.","t":"`$ANY`","key$":"artist_display","index$":10},"artist_id":{"a":true,"h":"Artist Id","n":"artist_id","r":false,"sh":"Unique identifier of the preferred artist/culture associated with this work","t":"`$STRING`","key$":"artist_id","index$":11},"artist_ids":{"a":true,"h":"Artist Ids","n":"artist_ids","r":false,"sh":"Unique identifier of all artist/cultures associated with this work","t":"`$ANY`","key$":"artist_ids","index$":12},"artist_title":{"a":true,"h":"Artist Title","n":"artist_title","r":false,"sh":"Name of the preferred artist/culture associated with this work","t":"`$ANY`","key$":"artist_title","index$":13},"artist_titles":{"a":true,"h":"Artist Titles","n":"artist_titles","r":false,"sh":"Names of all artist/cultures associated with this work","t":"`$ANY`","key$":"artist_titles","index$":14},"artwork_type_id":{"a":true,"h":"Artwork Type Id","n":"artwork_type_id","r":false,"sh":"Unique identifier of the kind of object or work","t":"`$STRING`","key$":"artwork_type_id","index$":15},"artwork_type_title":{"a":true,"h":"Artwork Type Title","n":"artwork_type_title","r":false,"sh":"The kind of object or work (e.g.","t":"`$ANY`","key$":"artwork_type_title","index$":16},"boost_rank":{"a":true,"h":"Boost Rank","n":"boost_rank","r":false,"sh":"Manual indication of what rank this artwork should take in search results.","t":"`$ANY`","key$":"boost_rank","index$":17},"catalog_based_search_keyword_titles":{"a":true,"h":"Catalog Based Search Keyword Titles","n":"catalog_based_search_keyword_titles","r":false,"sh":"The keyword search values that would be catalog-based searches on this record","t":"`$ANY`","key$":"catalog_based_search_keyword_titles","index$":18},"catalogue_display":{"a":true,"h":"Catalogue Display","n":"catalogue_display","r":false,"sh":"Brief text listing all the catalogues raisonnés which include this work.","t":"`$ANY`","key$":"catalogue_display","index$":19},"category_ids":{"a":true,"h":"Category Ids","n":"category_ids","r":false,"sh":"Unique identifiers of the categories this work is a part of","t":"`$ANY`","key$":"category_ids","index$":20},"category_titles":{"a":true,"h":"Category Titles","n":"category_titles","r":false,"sh":"Names of the categories this artwork is a part of","t":"`$ANY`","key$":"category_titles","index$":21},"classification_id":{"a":true,"h":"Classification Id","n":"classification_id","r":false,"sh":"Unique identifier of the preferred classification term for this work","t":"`$STRING`","key$":"classification_id","index$":22},"classification_ids":{"a":true,"h":"Classification Ids","n":"classification_ids","r":false,"sh":"Unique identifiers of all classification terms for this work","t":"`$ANY`","key$":"classification_ids","index$":23},"classification_title":{"a":true,"h":"Classification Title","n":"classification_title","r":false,"sh":"The name of the preferred classification term for this work","t":"`$ANY`","key$":"classification_title","index$":24},"classification_titles":{"a":true,"h":"Classification Titles","n":"classification_titles","r":false,"sh":"The names of all classification terms related to this artwork","t":"`$ANY`","key$":"classification_titles","index$":25},"color":{"a":true,"h":"Color","n":"color","r":false,"sh":"Dominant color of this artwork in HSL","t":"`$ANY`","key$":"color","index$":26},"colorfulness":{"a":true,"h":"Colorfulness","n":"colorfulness","r":false,"sh":"Unbounded positive float representing an abstract measure of colorfulness.","t":"`$ANY`","key$":"colorfulness","index$":27},"copyright_notice":{"a":true,"h":"Copyright Notice","n":"copyright_notice","r":false,"sh":"Statement notifying how the work is protected by copyright.","t":"`$ANY`","key$":"copyright_notice","index$":28},"credit_line":{"a":true,"h":"Credit Line","n":"credit_line","r":false,"sh":"Brief statement indicating how the work came into the collection","t":"`$ANY`","key$":"credit_line","index$":29},"date_display":{"a":true,"h":"Date Display","n":"date_display","r":false,"sh":"Readable, free-text description of the period of time associated with the creation of this work.","t":"`$ANY`","key$":"date_display","index$":30},"date_end":{"a":true,"h":"Date End","n":"date_end","r":false,"sh":"The year of the period of time associated with the creation of this work","t":"`$ANY`","key$":"date_end","index$":31},"date_qualifier_id":{"a":true,"h":"Date Qualifier Id","n":"date_qualifier_id","r":false,"sh":"Unique identifier of the qualifer to the dates provided for this record.","t":"`$STRING`","key$":"date_qualifier_id","index$":32},"date_qualifier_title":{"a":true,"h":"Date Qualifier Title","n":"date_qualifier_title","r":false,"sh":"Readable, text qualifer to the dates provided for this record.","t":"`$ANY`","key$":"date_qualifier_title","index$":33},"date_start":{"a":true,"h":"Date Start","n":"date_start","r":false,"sh":"The year of the period of time associated with the creation of this work","t":"`$ANY`","key$":"date_start","index$":34},"department_id":{"a":true,"h":"Department Id","n":"department_id","r":false,"sh":"Unique identifier of the curatorial department that this work belongs to","t":"`$STRING`","key$":"department_id","index$":35},"department_title":{"a":true,"h":"Department Title","n":"department_title","r":false,"sh":"Name of the curatorial department that this work belongs to","t":"`$ANY`","key$":"department_title","index$":36},"description":{"a":true,"h":"Description","n":"description","r":false,"sh":"Longer explanation describing the work","t":"`$STRING`","key$":"description","index$":37},"dimensions":{"a":true,"h":"Dimensions","n":"dimensions","r":false,"sh":"The size, shape, scale, and dimensions of the work.","t":"`$ANY`","key$":"dimensions","index$":38},"dimensions_detail":{"a":true,"h":"Dimensions Detail","n":"dimensions_detail","r":false,"sh":"The height, width, depth, and/or diameter of each section of the work in centimeters","t":"`$ANY`","key$":"dimensions_detail","index$":39},"document_ids":{"a":true,"h":"Document Ids","n":"document_ids","r":false,"sh":"Unique identifiers of assets that serve as documentation for this artwork","t":"`$ANY`","key$":"document_ids","index$":40},"edition":{"a":true,"h":"Edition","n":"edition","r":false,"sh":"Edition number if the work is one of many","t":"`$ANY`","key$":"edition","index$":41},"exhibition_history":{"a":true,"h":"Exhibition History","n":"exhibition_history","r":false,"sh":"List of all the places this work has been exhibited","t":"`$ANY`","key$":"exhibition_history","index$":42},"fiscal_year":{"a":true,"h":"Fiscal Year","n":"fiscal_year","r":false,"sh":"The fiscal year in which the work was acquired.","t":"`$ANY`","key$":"fiscal_year","index$":43},"fiscal_year_deaccession":{"a":true,"h":"Fiscal Year Deaccession","n":"fiscal_year_deaccession","r":false,"sh":"The fiscal year in which the work was deaccessioned.","t":"`$ANY`","key$":"fiscal_year_deaccession","index$":44},"gallery_id":{"a":true,"h":"Gallery Id","n":"gallery_id","r":false,"sh":"Unique identifier of the location of this work in our museum","t":"`$STRING`","key$":"gallery_id","index$":45},"gallery_title":{"a":true,"h":"Gallery Title","n":"gallery_title","r":false,"sh":"The location of this work in our museum","t":"`$ANY`","key$":"gallery_title","index$":46},"has_advanced_imaging":{"a":true,"h":"Has Advanced Imaging","n":"has_advanced_imaging","r":false,"sh":"Whether this artwork is enhanced with 3D models, 360 image sequences, Mirador views, etc.","t":"`$BOOLEAN`","key$":"has_advanced_imaging","index$":47},"has_educational_resources":{"a":true,"h":"Has Educational Resources","n":"has_educational_resources","r":false,"sh":"Whether this artwork has any documents tagged as educational","t":"`$BOOLEAN`","key$":"has_educational_resources","index$":48},"has_multimedia_resources":{"a":true,"h":"Has Multimedia Resources","n":"has_multimedia_resources","r":false,"sh":"Whether this artwork has any associated microsites, digital publications, or documents tagged as multimedia","t":"`$BOOLEAN`","key$":"has_multimedia_resources","index$":49},"has_not_been_viewed_much":{"a":true,"h":"Has Not Been Viewed Much","n":"has_not_been_viewed_much","r":false,"sh":"Whether the artwork hasn't been visited on our website very much","t":"`$BOOLEAN`","key$":"has_not_been_viewed_much","index$":50},"id":{"a":true,"h":"Id","n":"id","r":false,"sh":"Unique identifier of this resource.","t":"`$STRING`","key$":"id","index$":51},"image_embedding":{"a":true,"h":"Image Embedding","n":"image_embedding","r":false,"sh":"The generated embeddings describing the artwork image","t":"`$ANY`","key$":"image_embedding","index$":52},"image_id":{"a":true,"h":"Image Id","n":"image_id","r":false,"sh":"Unique identifier of the preferred image to use to represent this work","t":"`$STRING`","key$":"image_id","index$":53},"inscriptions":{"a":true,"h":"Inscriptions","n":"inscriptions","r":false,"sh":"A description of distinguishing or identifying physical markings that are on the work","t":"`$ANY`","key$":"inscriptions","index$":54},"internal_department_id":{"a":true,"h":"Internal Department Id","n":"internal_department_id","r":false,"sh":"An internal department id we use for analytics.","t":"`$STRING`","key$":"internal_department_id","index$":55},"is_boosted":{"a":true,"h":"Is Boosted","n":"is_boosted","r":false,"sh":"Whether this document should be boosted in search","t":"`$BOOLEAN`","key$":"is_boosted","index$":56},"is_on_view":{"a":true,"h":"Is On View","n":"is_on_view","r":false,"sh":"Whether the work is on display","t":"`$BOOLEAN`","key$":"is_on_view","index$":57},"is_public_domain":{"a":true,"h":"Is Public Domain","n":"is_public_domain","r":false,"sh":"Whether the work is in the public domain, meaning it was created before copyrights existed or has left the copyright term","t":"`$BOOLEAN`","key$":"is_public_domain","index$":58},"is_zoomable":{"a":true,"h":"Is Zoomable","n":"is_zoomable","r":false,"sh":"Whether images of the work are allowed to be displayed in a zoomable interface.","t":"`$BOOLEAN`","key$":"is_zoomable","index$":59},"latitude":{"a":true,"h":"Latitude","n":"latitude","r":false,"sh":"Latitude coordinate of the location of this work in our galleries","t":"`$NUMBER`","key$":"latitude","index$":60},"latlon":{"a":true,"h":"Latlon","n":"latlon","r":false,"sh":"Latitude and longitude coordinates of the location of this work in our galleries","t":"`$ANY`","key$":"latlon","index$":61},"longitude":{"a":true,"h":"Longitude","n":"longitude","r":false,"sh":"Longitude coordinate of the location of this work in our galleries","t":"`$NUMBER`","key$":"longitude","index$":62},"main_reference_number":{"a":true,"h":"Main Reference Number","n":"main_reference_number","r":false,"sh":"Unique identifier assigned to the artwork upon acquisition","t":"`$INTEGER`","key$":"main_reference_number","index$":63},"material_id":{"a":true,"h":"Material Id","n":"material_id","r":false,"sh":"Unique identifier of the preferred material term for this work","t":"`$STRING`","key$":"material_id","index$":64},"material_ids":{"a":true,"h":"Material Ids","n":"material_ids","r":false,"sh":"Unique identifiers of all material terms for this work","t":"`$ANY`","key$":"material_ids","index$":65},"material_titles":{"a":true,"h":"Material Titles","n":"material_titles","r":false,"sh":"The names of all material terms related to this artwork","t":"`$ANY`","key$":"material_titles","index$":66},"max_zoom_window_size":{"a":true,"h":"Max Zoom Window Size","n":"max_zoom_window_size","r":false,"sh":"The maximum size of the window the image is allowed to be viewed in, in pixels.","t":"`$ANY`","key$":"max_zoom_window_size","index$":67},"medium_display":{"a":true,"h":"Medium Display","n":"medium_display","r":false,"sh":"The substances or materials used in the creation of a work","t":"`$ANY`","key$":"medium_display","index$":68},"nomisma_id":{"a":true,"h":"Nomisma Id","n":"nomisma_id","r":false,"sh":"Unique identifier of this work in the nomisma coin database","t":"`$STRING`","key$":"nomisma_id","index$":69},"on_loan_display":{"a":true,"h":"On Loan Display","n":"on_loan_display","r":false,"sh":"If an artwork is on loan, this contains details about the loan","t":"`$ANY`","key$":"on_loan_display","index$":70},"pageviews":{"a":true,"h":"Pageviews","n":"pageviews","r":false,"sh":"Approx.","t":"`$ANY`","key$":"pageviews","index$":71},"pageviews_recent":{"a":true,"h":"Pageviews Recent","n":"pageviews_recent","r":false,"sh":"Approx.","t":"`$ANY`","key$":"pageviews_recent","index$":72},"place_of_origin":{"a":true,"h":"Place Of Origin","n":"place_of_origin","r":false,"sh":"The location where the creation, design, or production of the work took place, or the original location of the work","t":"`$ANY`","key$":"place_of_origin","index$":73},"provenance_text":{"a":true,"h":"Provenance Text","n":"provenance_text","r":false,"sh":"Ownership/collecting history of the work.","t":"`$ANY`","key$":"provenance_text","index$":74},"publication_history":{"a":true,"h":"Publication History","n":"publication_history","r":false,"sh":"Bibliographic list of all the places this work has been published","t":"`$ANY`","key$":"publication_history","index$":75},"publishing_verification_level":{"a":true,"h":"Publishing Verification Level","n":"publishing_verification_level","r":false,"sh":"Indicator of how much metadata on the work in published.","t":"`$ANY`","key$":"publishing_verification_level","index$":76},"section_ids":{"a":true,"h":"Section Ids","n":"section_ids","r":false,"sh":"Unique identifiers of the digital publication chapters this work in included in","t":"`$ANY`","key$":"section_ids","index$":77},"section_titles":{"a":true,"h":"Section Titles","n":"section_titles","r":false,"sh":"Names of the digital publication chapters this work is included in","t":"`$ANY`","key$":"section_titles","index$":78},"short_description":{"a":true,"h":"Short Description","n":"short_description","r":false,"sh":"Short explanation describing the work","t":"`$ANY`","key$":"short_description","index$":79},"site_ids":{"a":true,"h":"Site Ids","n":"site_ids","r":false,"sh":"Unique identifiers of the microsites this work is a part of","t":"`$ANY`","key$":"site_ids","index$":80},"sound_ids":{"a":true,"h":"Sound Ids","n":"sound_ids","r":false,"sh":"Unique identifiers of the audio about this work","t":"`$ANY`","key$":"sound_ids","index$":81},"source_updated_at":{"a":true,"h":"Source Updated At","n":"source_updated_at","r":false,"sh":"Date and time the resource was updated in the source system","t":"`$ANY`","key$":"source_updated_at","index$":82},"style_id":{"a":true,"h":"Style Id","n":"style_id","r":false,"sh":"Unique identifier of the preferred style term for this work","t":"`$STRING`","key$":"style_id","index$":83},"style_ids":{"a":true,"h":"Style Ids","n":"style_ids","r":false,"sh":"Unique identifiers of all style terms for this work","t":"`$ANY`","key$":"style_ids","index$":84},"style_title":{"a":true,"h":"Style Title","n":"style_title","r":false,"sh":"The name of the preferred style term for this work","t":"`$ANY`","key$":"style_title","index$":85},"style_titles":{"a":true,"h":"Style Titles","n":"style_titles","r":false,"sh":"The names of all style terms related to this artwork","t":"`$ANY`","key$":"style_titles","index$":86},"subject_id":{"a":true,"h":"Subject Id","n":"subject_id","r":false,"sh":"Unique identifier of the preferred subject term for this work","t":"`$STRING`","key$":"subject_id","index$":87},"subject_ids":{"a":true,"h":"Subject Ids","n":"subject_ids","r":false,"sh":"Unique identifiers of all subject terms for this work","t":"`$ANY`","key$":"subject_ids","index$":88},"subject_titles":{"a":true,"h":"Subject Titles","n":"subject_titles","r":false,"sh":"The names of all subject terms related to this artwork","t":"`$ANY`","key$":"subject_titles","index$":89},"suggest_autocomplete_all":{"a":true,"h":"Suggest Autocomplete All","n":"suggest_autocomplete_all","r":false,"sh":"Internal field to power the `/autosuggest` endpoint.","t":"`$ANY`","key$":"suggest_autocomplete_all","index$":90},"suggest_autocomplete_boosted":{"a":true,"h":"Suggest Autocomplete Boosted","n":"suggest_autocomplete_boosted","r":false,"sh":"Internal field to power the `/autocomplete` endpoint.","t":"`$ANY`","key$":"suggest_autocomplete_boosted","index$":91},"technique_id":{"a":true,"h":"Technique Id","n":"technique_id","r":false,"sh":"Unique identifier of the preferred technique term for this work","t":"`$STRING`","key$":"technique_id","index$":92},"technique_ids":{"a":true,"h":"Technique Ids","n":"technique_ids","r":false,"sh":"Unique identifiers of all technique terms for this work","t":"`$ANY`","key$":"technique_ids","index$":93},"technique_titles":{"a":true,"h":"Technique Titles","n":"technique_titles","r":false,"sh":"The names of all technique terms related to this artwork","t":"`$ANY`","key$":"technique_titles","index$":94},"term_titles":{"a":true,"h":"Term Titles","n":"term_titles","r":false,"sh":"The names of the taxonomy tags for this work","t":"`$ANY`","key$":"term_titles","index$":95},"text_embedding":{"a":true,"h":"Text Embedding","n":"text_embedding","r":false,"sh":"The generated embeddings of artwork text","t":"`$ANY`","key$":"text_embedding","index$":96},"text_ids":{"a":true,"h":"Text Ids","n":"text_ids","r":false,"sh":"Unique identifiers of the texts about this work","t":"`$ANY`","key$":"text_ids","index$":97},"theme_titles":{"a":true,"h":"Theme Titles","n":"theme_titles","r":false,"sh":"The names of all thematic publish categories related to this artwork","t":"`$ANY`","key$":"theme_titles","index$":98},"thumbnail":{"a":true,"h":"Thumbnail","n":"thumbnail","r":false,"sh":"Metadata about the image referenced by `image_id`.","t":"`$ANY`","key$":"thumbnail","index$":99},"timestamp":{"a":true,"h":"Timestamp","n":"timestamp","r":false,"sh":"Date and time the record was updated in the aggregator search index","t":"`$ANY`","key$":"timestamp","index$":100},"title":{"a":true,"h":"Title","n":"title","r":false,"sh":"The name of this resource","t":"`$STRING`","key$":"title","index$":101},"updated_at":{"a":true,"h":"Updated At","n":"updated_at","r":false,"sh":"Date and time the record was updated in the aggregator database","t":"`$ANY`","key$":"updated_at","index$":102},"video_ids":{"a":true,"h":"Video Ids","n":"video_ids","r":false,"sh":"Unique identifiers of the videos about this work","t":"`$ANY`","key$":"video_ids","index$":103}},"id":{"field":"id","name":"id"},"name":"artwork","op":{"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /artworks","source":"openapi3","version":2},"g":{},"k":"http","m":"GET","o":"/artworks","q":{},"r":{},"s":[{"lit":"artworks"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /artworks/{id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"GET","o":"/artworks/{id}","q":{"exist":["id"]},"r":{},"s":[{"lit":"artworks"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"artwork","name__orig":"artwork","Name":"Artwork","name_":"artwork","name-":"artwork","NAME":"ARTWORK","index$":4}, {"active":true,"entity":"artwork","key$":"BasicArtworkFlow","kind":"basic","name":"BasicArtworkFlow","param":{},"step":[{"a":true,"d":{},"i":{},"m":{},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"artwork_ref01"}}],"index$":0},{"a":true,"d":{},"i":{"ref":"artwork_ref01","srcdatavar":"artwork_ref01_data","suffix":"_dt0"},"m":{"id":"artwork01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-artwork_ref01"}}],"index$":1}]}, 'Artwork', {"GET /artworks":{"protocol":"http","parameters":[]},"GET /artworks/{id}":{"protocol":"http","parameters":[{"name":"id","in":"path","type":"string","required":true,"x-ref":"#/parameters/id","index$":0}]}}, { strict: LIVE_STRICT, t })
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let artwork_ref01_data = Object.values(setup.data.existing.artwork)[0] as any

    // LIST
    const artwork_ref01_ent = client.Artwork()
    const artwork_ref01_match: any = {}

    const artwork_ref01_list = (await artwork_ref01_ent.list(artwork_ref01_match)).map((e: any) => e.data())


    // LOAD
    const artwork_ref01_match_dt0: any = {}
    artwork_ref01_match_dt0.id = artwork_ref01_data.id
    const artwork_ref01_data_dt0 = (await artwork_ref01_ent.load(artwork_ref01_match_dt0)).data()
    assert(artwork_ref01_data_dt0.id === artwork_ref01_data.id)


  })
})



// main.kit.test.live.strict is true (the default is true): a live
// request that fails, or a live test missing an input it needs,
// fails the test.
// An account with no record for a test to read skips it either way.
const LIVE_STRICT = true

function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/artwork/ArtworkTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = ArtInstituteOfChicagoSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['artwork01','artwork02','artwork03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'ART_INSTITUTE_OF_CHICAGO_TEST_ARTWORK_ENTID': idmap,
    'ART_INSTITUTE_OF_CHICAGO_TEST_LIVE': 'FALSE',
    'ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['ART_INSTITUTE_OF_CHICAGO_TEST_ARTWORK_ENTID']

  const live = 'TRUE' === env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['ART_INSTITUTE_OF_CHICAGO_TEST_ARTWORK_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new ArtInstituteOfChicagoSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
      {
      },
      // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
      // last entry is undefined, and basicSetup is normally called with no
      // argument at all - so a bare 'extra' silently discarded the apikey
      // and server values above and handed the SDK undefined. Harmless
      // while there was nothing in that object; not harmless now.
      extra || {},
      { system: { fetch: transport.fetch } }
    ]))
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
  }

  return setup
}
  
