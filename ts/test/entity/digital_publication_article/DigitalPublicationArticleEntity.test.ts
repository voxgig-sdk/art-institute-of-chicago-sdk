

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { ArtInstituteOfChicagoSDK, BaseFeature, stdutil } from '../../..'

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


describe('DigitalPublicationArticleEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE.
  afterEach(liveDelay('ART_INSTITUTE_OF_CHICAGO_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = ArtInstituteOfChicagoSDK.test()
    const ent = testsdk.DigitalPublicationArticle()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE
    for (const op of ['list', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'digital_publication_article.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"api_link":{"a":true,"h":"Api Link","n":"api_link","r":false,"sh":"REST API link for this resource","t":"`$ANY`","key$":"api_link","index$":0},"api_model":{"a":true,"h":"Api Model","n":"api_model","r":false,"sh":"REST API resource type or endpoint","t":"`$ANY`","key$":"api_model","index$":1},"author_display":{"a":true,"h":"Author Display","n":"author_display","r":false,"sh":"A display-friendly text of the authors of this article","t":"`$ANY`","key$":"author_display","index$":2},"copy":{"a":true,"h":"Copy","n":"copy","r":false,"sh":"The text of the article","t":"`$ANY`","key$":"copy","index$":3},"digital_publication_id":{"a":true,"h":"Digital Publication Id","n":"digital_publication_id","r":false,"sh":"Unique identifier of the digital publication this article belongs to","t":"`$STRING`","key$":"digital_publication_id","index$":4},"id":{"a":true,"h":"Id","n":"id","r":false,"sh":"Unique identifier of this resource.","t":"`$STRING`","key$":"id","index$":5},"source_updated_at":{"a":true,"h":"Source Updated At","n":"source_updated_at","r":false,"sh":"Date and time the resource was updated in the source system","t":"`$ANY`","key$":"source_updated_at","index$":6},"suggest_autocomplete_all":{"a":true,"h":"Suggest Autocomplete All","n":"suggest_autocomplete_all","r":false,"sh":"Internal field to power the `/autosuggest` endpoint.","t":"`$ANY`","key$":"suggest_autocomplete_all","index$":7},"suggest_autocomplete_boosted":{"a":true,"h":"Suggest Autocomplete Boosted","n":"suggest_autocomplete_boosted","r":false,"sh":"Internal field to power the `/autocomplete` endpoint.","t":"`$ANY`","key$":"suggest_autocomplete_boosted","index$":8},"timestamp":{"a":true,"h":"Timestamp","n":"timestamp","r":false,"sh":"Date and time the record was updated in the aggregator search index","t":"`$ANY`","key$":"timestamp","index$":9},"title":{"a":true,"h":"Title","n":"title","r":false,"sh":"The name of this resource","t":"`$STRING`","key$":"title","index$":10},"updated_at":{"a":true,"h":"Updated At","n":"updated_at","r":false,"sh":"Date and time the record was updated in the aggregator database","t":"`$ANY`","key$":"updated_at","index$":11},"web_url":{"a":true,"h":"Web Url","n":"web_url","r":false,"sh":"The URL to this article on our website","t":"`$ANY`","key$":"web_url","index$":12}},"id":{"field":"id","name":"id"},"name":"digital_publication_article","op":{"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /digital-publication-articles","source":"openapi3","version":2},"g":{},"k":"http","m":"GET","o":"/digital-publication-articles","q":{},"r":{},"s":[{"lit":"digital-publication-articles"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /digital-publication-articles/{id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"GET","o":"/digital-publication-articles/{id}","q":{"exist":["id"]},"r":{},"s":[{"lit":"digital-publication-articles"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"digital_publication_article","name__orig":"digital_publication_article","Name":"DigitalPublicationArticle","name_":"digital_publication_article","name-":"digital-publication-article","NAME":"DIGITAL_PUBLICATION_ARTICLE","index$":10}, {"active":true,"entity":"digital_publication_article","key$":"BasicDigitalPublicationArticleFlow","kind":"basic","name":"BasicDigitalPublicationArticleFlow","param":{},"step":[{"a":true,"d":{},"i":{},"m":{},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"digital_publication_article_ref01"}}],"index$":0},{"a":true,"d":{},"i":{"ref":"digital_publication_article_ref01","srcdatavar":"digital_publication_article_ref01_data","suffix":"_dt0"},"m":{"id":"digital_publication_article01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-digital_publication_article_ref01"}}],"index$":1}]}, 'DigitalPublicationArticle', {"GET /digital-publication-articles":{"protocol":"http","responses":{"200":{"description":"Successful operation","schema":{"type":"array","items":{"properties":{"id":{"description":"Unique identifier of this resource. Taken from the source system.","key$":"id"},"api_model":{"description":"REST API resource type or endpoint","key$":"api_model"},"api_link":{"description":"REST API link for this resource","key$":"api_link"},"title":{"description":"The name of this resource","key$":"title"},"web_url":{"description":"The URL to this article on our website","key$":"web_url"},"copy":{"description":"The text of the article","key$":"copy"},"author_display":{"description":"A display-friendly text of the authors of this article","key$":"author_display"},"digital_publication_id":{"description":"Unique identifier of the digital publication this article belongs to","key$":"digital_publication_id"},"suggest_autocomplete_boosted":{"description":"Internal field to power the `/autocomplete` endpoint. Do not use directly.","key$":"suggest_autocomplete_boosted"},"suggest_autocomplete_all":{"description":"Internal field to power the `/autosuggest` endpoint. Do not use directly.","key$":"suggest_autocomplete_all"},"source_updated_at":{"description":"Date and time the resource was updated in the source system","key$":"source_updated_at"},"updated_at":{"description":"Date and time the record was updated in the aggregator database","key$":"updated_at"},"timestamp":{"description":"Date and time the record was updated in the aggregator search index","key$":"timestamp"}},"type":"object","x-ref":"#/definitions/DigitalPublicationArticle","index$":0}}},"default":{"description":"error","schema":{"required":["status","error","detail"],"properties":{"status":{"type":"integer"},"error":{"type":"string"},"detail":{"type":"string"}},"x-ref":"#/definitions/Error"}}},"produces":["application/json"],"parameters":[],"securitySource":"unspecified"},"GET /digital-publication-articles/{id}":{"protocol":"http","responses":{"200":{"description":"Successful operation","schema":{"type":"array","items":{"properties":{"id":{"description":"Unique identifier of this resource. Taken from the source system.","key$":"id"},"api_model":{"description":"REST API resource type or endpoint","key$":"api_model"},"api_link":{"description":"REST API link for this resource","key$":"api_link"},"title":{"description":"The name of this resource","key$":"title"},"web_url":{"description":"The URL to this article on our website","key$":"web_url"},"copy":{"description":"The text of the article","key$":"copy"},"author_display":{"description":"A display-friendly text of the authors of this article","key$":"author_display"},"digital_publication_id":{"description":"Unique identifier of the digital publication this article belongs to","key$":"digital_publication_id"},"suggest_autocomplete_boosted":{"description":"Internal field to power the `/autocomplete` endpoint. Do not use directly.","key$":"suggest_autocomplete_boosted"},"suggest_autocomplete_all":{"description":"Internal field to power the `/autosuggest` endpoint. Do not use directly.","key$":"suggest_autocomplete_all"},"source_updated_at":{"description":"Date and time the resource was updated in the source system","key$":"source_updated_at"},"updated_at":{"description":"Date and time the record was updated in the aggregator database","key$":"updated_at"},"timestamp":{"description":"Date and time the record was updated in the aggregator search index","key$":"timestamp"}},"type":"object","x-ref":"#/definitions/DigitalPublicationArticle","key$":"items"}}},"default":{"description":"error","schema":{"required":["status","error","detail"],"properties":{"status":{"type":"integer"},"error":{"type":"string"},"detail":{"type":"string"}},"x-ref":"#/definitions/Error"}}},"produces":["application/json"],"parameters":[{"name":"id","in":"path","type":"string","required":true,"x-ref":"#/parameters/id","index$":0}],"securitySource":"unspecified"}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let digital_publication_article_ref01_data = Object.values(setup.data.existing.digital_publication_article)[0] as any

    // LIST
    const digital_publication_article_ref01_ent = client.DigitalPublicationArticle()
    const digital_publication_article_ref01_match: any = {}

    const digital_publication_article_ref01_list = (await digital_publication_article_ref01_ent.list(digital_publication_article_ref01_match)).map((e: any) => e.data())


    // LOAD
    const digital_publication_article_ref01_match_dt0: any = {}
    digital_publication_article_ref01_match_dt0.id = digital_publication_article_ref01_data.id
    const digital_publication_article_ref01_data_dt0 = (await digital_publication_article_ref01_ent.load(digital_publication_article_ref01_match_dt0)).data()
    assert(digital_publication_article_ref01_data_dt0.id === digital_publication_article_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/digital_publication_article/DigitalPublicationArticleTestData.json')

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
    ['digital_publication_article01','digital_publication_article02','digital_publication_article03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'ART_INSTITUTE_OF_CHICAGO_TEST_DIGITAL_PUBLICATION_ARTICLE_ENTID': idmap,
    'ART_INSTITUTE_OF_CHICAGO_TEST_LIVE': 'FALSE',
    'ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['ART_INSTITUTE_OF_CHICAGO_TEST_DIGITAL_PUBLICATION_ARTICLE_ENTID']

  const live = 'TRUE' === env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['ART_INSTITUTE_OF_CHICAGO_TEST_DIGITAL_PUBLICATION_ARTICLE_ENTID']
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
  
