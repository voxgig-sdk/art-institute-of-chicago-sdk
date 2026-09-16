

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


describe('TourEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE.
  afterEach(liveDelay('ART_INSTITUTE_OF_CHICAGO_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = ArtInstituteOfChicagoSDK.test()
    const ent = testsdk.Tour()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE
    for (const op of ['list', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'tour.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"api_link","req":false,"short":"REST API link for this resource","type":"`$ANY`","index$":0},{"active":true,"name":"api_model","req":false,"short":"REST API resource type or endpoint","type":"`$ANY`","index$":1},{"active":true,"name":"artist_titles","req":false,"short":"Names of the artists of the artworks featured in this tour's tour stops","type":"`$ANY`","index$":2},{"active":true,"name":"artwork_titles","req":false,"short":"Names of the artworks featured in this tour's tour stops","type":"`$ANY`","index$":3},{"active":true,"name":"description","req":false,"short":"Explanation of what the tour is","type":"`$STRING`","index$":4},{"active":true,"name":"id","req":false,"short":"Unique identifier of this resource.","type":"`$STRING`","index$":5},{"active":true,"name":"image","req":false,"short":"The main image for the tour","type":"`$ANY`","index$":6},{"active":true,"name":"intro","req":false,"short":"Text introducing the tour","type":"`$ANY`","index$":7},{"active":true,"name":"intro_link","req":false,"short":"Link to the audio file of the introduction","type":"`$ANY`","index$":8},{"active":true,"name":"intro_transcript","req":false,"short":"Transcript of the introduction audio to the tour","type":"`$ANY`","index$":9},{"active":true,"name":"source_updated_at","req":false,"short":"Date and time the resource was updated in the source system","type":"`$ANY`","index$":10},{"active":true,"name":"suggest_autocomplete_all","req":false,"short":"Internal field to power the `/autosuggest` endpoint.","type":"`$ANY`","index$":11},{"active":true,"name":"suggest_autocomplete_boosted","req":false,"short":"Internal field to power the `/autocomplete` endpoint.","type":"`$ANY`","index$":12},{"active":true,"name":"timestamp","req":false,"short":"Date and time the record was updated in the aggregator search index","type":"`$ANY`","index$":13},{"active":true,"name":"title","req":false,"short":"The name of this resource","type":"`$STRING`","index$":14},{"active":true,"name":"updated_at","req":false,"short":"Date and time the record was updated in the aggregator database","type":"`$ANY`","index$":15},{"active":true,"name":"weight","req":false,"short":"Number representing this tour's sort order","type":"`$NUMBER`","index$":16}],"id":{"field":"id","name":"id"},"name":"tour","op":{"list":{"input":"data","name":"list","points":[{"active":true,"args":{},"contract":{"id":"GET /tours","json":"{\"parameters\":[],\"produces\":[\"application/json\"],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"Successful operation\",\"schema\":{\"items\":{\"properties\":{\"api_link\":{\"description\":\"REST API link for this resource\"},\"api_model\":{\"description\":\"REST API resource type or endpoint\"},\"artist_titles\":{\"description\":\"Names of the artists of the artworks featured in this tour's tour stops\"},\"artwork_titles\":{\"description\":\"Names of the artworks featured in this tour's tour stops\"},\"description\":{\"description\":\"Explanation of what the tour is\"},\"id\":{\"description\":\"Unique identifier of this resource. Taken from the source system.\"},\"image\":{\"description\":\"The main image for the tour\"},\"intro\":{\"description\":\"Text introducing the tour\"},\"intro_link\":{\"description\":\"Link to the audio file of the introduction\"},\"intro_transcript\":{\"description\":\"Transcript of the introduction audio to the tour\"},\"source_updated_at\":{\"description\":\"Date and time the resource was updated in the source system\"},\"suggest_autocomplete_all\":{\"description\":\"Internal field to power the `/autosuggest` endpoint. Do not use directly.\"},\"suggest_autocomplete_boosted\":{\"description\":\"Internal field to power the `/autocomplete` endpoint. Do not use directly.\"},\"timestamp\":{\"description\":\"Date and time the record was updated in the aggregator search index\"},\"title\":{\"description\":\"The name of this resource\"},\"updated_at\":{\"description\":\"Date and time the record was updated in the aggregator database\"},\"weight\":{\"description\":\"Number representing this tour's sort order\"}},\"type\":\"object\"},\"type\":\"array\"}},\"default\":{\"description\":\"error\",\"schema\":{\"properties\":{\"detail\":{\"type\":\"string\"},\"error\":{\"type\":\"string\"},\"status\":{\"type\":\"integer\"}},\"required\":[\"status\",\"error\",\"detail\"]}}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/tours","segments":[{"lit":"tours"}],"select":{},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"id","orig":"id","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"GET /tours/{id}","json":"{\"parameters\":[{\"in\":\"path\",\"name\":\"id\",\"required\":true,\"type\":\"string\"}],\"produces\":[\"application/json\"],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"Successful operation\",\"schema\":{\"items\":{\"properties\":{\"api_link\":{\"description\":\"REST API link for this resource\"},\"api_model\":{\"description\":\"REST API resource type or endpoint\"},\"artist_titles\":{\"description\":\"Names of the artists of the artworks featured in this tour's tour stops\"},\"artwork_titles\":{\"description\":\"Names of the artworks featured in this tour's tour stops\"},\"description\":{\"description\":\"Explanation of what the tour is\"},\"id\":{\"description\":\"Unique identifier of this resource. Taken from the source system.\"},\"image\":{\"description\":\"The main image for the tour\"},\"intro\":{\"description\":\"Text introducing the tour\"},\"intro_link\":{\"description\":\"Link to the audio file of the introduction\"},\"intro_transcript\":{\"description\":\"Transcript of the introduction audio to the tour\"},\"source_updated_at\":{\"description\":\"Date and time the resource was updated in the source system\"},\"suggest_autocomplete_all\":{\"description\":\"Internal field to power the `/autosuggest` endpoint. Do not use directly.\"},\"suggest_autocomplete_boosted\":{\"description\":\"Internal field to power the `/autocomplete` endpoint. Do not use directly.\"},\"timestamp\":{\"description\":\"Date and time the record was updated in the aggregator search index\"},\"title\":{\"description\":\"The name of this resource\"},\"updated_at\":{\"description\":\"Date and time the record was updated in the aggregator database\"},\"weight\":{\"description\":\"Number representing this tour's sort order\"}},\"type\":\"object\"},\"type\":\"array\"}},\"default\":{\"description\":\"error\",\"schema\":{\"properties\":{\"detail\":{\"type\":\"string\"},\"error\":{\"type\":\"string\"},\"status\":{\"type\":\"integer\"}},\"required\":[\"status\",\"error\",\"detail\"]}}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/tours/{id}","segments":[{"lit":"tours"},{"var":"id"}],"select":{"exist":["id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"tour","name__orig":"tour","Name":"Tour","name_":"tour","name-":"tour","NAME":"TOUR","index$":33}, {"active":true,"entity":"tour","key$":"BasicTourFlow","kind":"basic","name":"BasicTourFlow","param":{},"step":[{"active":true,"data":{},"input":{},"match":{},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"tour_ref01"}}],"index$":0},{"active":true,"data":{},"input":{"ref":"tour_ref01","srcdatavar":"tour_ref01_data","suffix":"_dt0"},"match":{"id":"tour01"},"op":"load","spec":[],"valid":[{"apply":"TextFieldMark","def":{"mark":"Mark01-tour_ref01"}}],"index$":1}]}, 'Tour')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let tour_ref01_data = Object.values(setup.data.existing.tour)[0] as any

    // LIST
    const tour_ref01_ent = client.Tour()
    const tour_ref01_match: any = {}

    const tour_ref01_list = (await tour_ref01_ent.list(tour_ref01_match)).map((e: any) => e.data())


    // LOAD
    const tour_ref01_match_dt0: any = {}
    tour_ref01_match_dt0.id = tour_ref01_data.id
    const tour_ref01_data_dt0 = (await tour_ref01_ent.load(tour_ref01_match_dt0)).data()
    assert(tour_ref01_data_dt0.id === tour_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/tour/TourTestData.json')

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
    ['tour01','tour02','tour03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'ART_INSTITUTE_OF_CHICAGO_TEST_TOUR_ENTID': idmap,
    'ART_INSTITUTE_OF_CHICAGO_TEST_LIVE': 'FALSE',
    'ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['ART_INSTITUTE_OF_CHICAGO_TEST_TOUR_ENTID']

  const live = 'TRUE' === env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['ART_INSTITUTE_OF_CHICAGO_TEST_TOUR_ENTID']
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
  
