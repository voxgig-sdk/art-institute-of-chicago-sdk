

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


// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
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
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"api_link","req":false,"short":"REST API link for this resource","type":"`$ANY`","index$":0},{"active":true,"name":"api_model","req":false,"short":"REST API resource type or endpoint","type":"`$ANY`","index$":1},{"active":true,"name":"author_display","req":false,"short":"A display-friendly text of the authors of this article","type":"`$ANY`","index$":2},{"active":true,"name":"copy","req":false,"short":"The text of the article","type":"`$ANY`","index$":3},{"active":true,"name":"digital_publication_id","req":false,"short":"Unique identifier of the digital publication this article belongs to","type":"`$STRING`","index$":4},{"active":true,"name":"id","req":false,"short":"Unique identifier of this resource.","type":"`$STRING`","index$":5},{"active":true,"name":"source_updated_at","req":false,"short":"Date and time the resource was updated in the source system","type":"`$ANY`","index$":6},{"active":true,"name":"suggest_autocomplete_all","req":false,"short":"Internal field to power the `/autosuggest` endpoint.","type":"`$ANY`","index$":7},{"active":true,"name":"suggest_autocomplete_boosted","req":false,"short":"Internal field to power the `/autocomplete` endpoint.","type":"`$ANY`","index$":8},{"active":true,"name":"timestamp","req":false,"short":"Date and time the record was updated in the aggregator search index","type":"`$ANY`","index$":9},{"active":true,"name":"title","req":false,"short":"The name of this resource","type":"`$STRING`","index$":10},{"active":true,"name":"updated_at","req":false,"short":"Date and time the record was updated in the aggregator database","type":"`$ANY`","index$":11},{"active":true,"name":"web_url","req":false,"short":"The URL to this article on our website","type":"`$ANY`","index$":12}],"id":{"field":"id","name":"id"},"name":"digital_publication_article","op":{"list":{"input":"data","name":"list","points":[{"active":true,"args":{},"contract":{"id":"GET /digital-publication-articles","json":"{\"parameters\":[],\"produces\":[\"application/json\"],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"Successful operation\",\"schema\":{\"items\":{\"properties\":{\"api_link\":{\"description\":\"REST API link for this resource\"},\"api_model\":{\"description\":\"REST API resource type or endpoint\"},\"author_display\":{\"description\":\"A display-friendly text of the authors of this article\"},\"copy\":{\"description\":\"The text of the article\"},\"digital_publication_id\":{\"description\":\"Unique identifier of the digital publication this article belongs to\"},\"id\":{\"description\":\"Unique identifier of this resource. Taken from the source system.\"},\"source_updated_at\":{\"description\":\"Date and time the resource was updated in the source system\"},\"suggest_autocomplete_all\":{\"description\":\"Internal field to power the `/autosuggest` endpoint. Do not use directly.\"},\"suggest_autocomplete_boosted\":{\"description\":\"Internal field to power the `/autocomplete` endpoint. Do not use directly.\"},\"timestamp\":{\"description\":\"Date and time the record was updated in the aggregator search index\"},\"title\":{\"description\":\"The name of this resource\"},\"updated_at\":{\"description\":\"Date and time the record was updated in the aggregator database\"},\"web_url\":{\"description\":\"The URL to this article on our website\"}},\"type\":\"object\"},\"type\":\"array\"}},\"default\":{\"description\":\"error\",\"schema\":{\"properties\":{\"detail\":{\"type\":\"string\"},\"error\":{\"type\":\"string\"},\"status\":{\"type\":\"integer\"}},\"required\":[\"status\",\"error\",\"detail\"]}}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/digital-publication-articles","segments":[{"lit":"digital-publication-articles"}],"select":{},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"id","orig":"id","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"GET /digital-publication-articles/{id}","json":"{\"parameters\":[{\"in\":\"path\",\"name\":\"id\",\"required\":true,\"type\":\"string\"}],\"produces\":[\"application/json\"],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"Successful operation\",\"schema\":{\"items\":{\"properties\":{\"api_link\":{\"description\":\"REST API link for this resource\"},\"api_model\":{\"description\":\"REST API resource type or endpoint\"},\"author_display\":{\"description\":\"A display-friendly text of the authors of this article\"},\"copy\":{\"description\":\"The text of the article\"},\"digital_publication_id\":{\"description\":\"Unique identifier of the digital publication this article belongs to\"},\"id\":{\"description\":\"Unique identifier of this resource. Taken from the source system.\"},\"source_updated_at\":{\"description\":\"Date and time the resource was updated in the source system\"},\"suggest_autocomplete_all\":{\"description\":\"Internal field to power the `/autosuggest` endpoint. Do not use directly.\"},\"suggest_autocomplete_boosted\":{\"description\":\"Internal field to power the `/autocomplete` endpoint. Do not use directly.\"},\"timestamp\":{\"description\":\"Date and time the record was updated in the aggregator search index\"},\"title\":{\"description\":\"The name of this resource\"},\"updated_at\":{\"description\":\"Date and time the record was updated in the aggregator database\"},\"web_url\":{\"description\":\"The URL to this article on our website\"}},\"type\":\"object\"},\"type\":\"array\"}},\"default\":{\"description\":\"error\",\"schema\":{\"properties\":{\"detail\":{\"type\":\"string\"},\"error\":{\"type\":\"string\"},\"status\":{\"type\":\"integer\"}},\"required\":[\"status\",\"error\",\"detail\"]}}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/digital-publication-articles/{id}","segments":[{"lit":"digital-publication-articles"},{"var":"id"}],"select":{"exist":["id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"digital_publication_article","name__orig":"digital_publication_article","Name":"DigitalPublicationArticle","name_":"digital_publication_article","name-":"digital-publication-article","NAME":"DIGITAL_PUBLICATION_ARTICLE","index$":10}, {"active":true,"entity":"digital_publication_article","key$":"BasicDigitalPublicationArticleFlow","kind":"basic","name":"BasicDigitalPublicationArticleFlow","param":{},"step":[{"active":true,"data":{},"input":{},"match":{},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"digital_publication_article_ref01"}}],"index$":0},{"active":true,"data":{},"input":{"ref":"digital_publication_article_ref01","srcdatavar":"digital_publication_article_ref01_data","suffix":"_dt0"},"match":{"id":"digital_publication_article01"},"op":"load","spec":[],"valid":[{"apply":"TextFieldMark","def":{"mark":"Mark01-digital_publication_article_ref01"}}],"index$":1}]}, 'DigitalPublicationArticle')
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
  
