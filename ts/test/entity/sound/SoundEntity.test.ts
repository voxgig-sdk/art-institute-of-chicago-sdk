

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


describe('SoundEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE.
  afterEach(liveDelay('ART_INSTITUTE_OF_CHICAGO_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = ArtInstituteOfChicagoSDK.test()
    const ent = testsdk.Sound()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE
    for (const op of ['list', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'sound.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"alt_text","req":false,"short":"Alternative text for the asset to describe it to people with low or no vision","type":"`$ANY`","index$":0},{"active":true,"name":"api_link","req":false,"short":"REST API link for this resource","type":"`$ANY`","index$":1},{"active":true,"name":"api_model","req":false,"short":"REST API resource type or endpoint","type":"`$ANY`","index$":2},{"active":true,"name":"artwork_ids","req":false,"short":"Unique identifiers of the artworks associated with this asset","type":"`$ANY`","index$":3},{"active":true,"name":"artwork_titles","req":false,"short":"Names of the artworks associated with this asset","type":"`$ANY`","index$":4},{"active":true,"name":"content","req":false,"short":"Text of or URL to the contents of this asset","type":"`$ANY`","index$":5},{"active":true,"name":"content_e_tag","req":false,"short":"Arbitrary unique identifier that changes when the binary file gets updated","type":"`$ANY`","index$":6},{"active":true,"name":"credit_line","req":false,"short":"Asset-specific copyright information","type":"`$ANY`","index$":7},{"active":true,"name":"id","req":false,"short":"Unique identifier of this resource.","type":"`$STRING`","index$":8},{"active":true,"name":"is_educational_resource","req":false,"short":"Whether this resource is considered to be educational","type":"`$BOOLEAN`","index$":9},{"active":true,"name":"is_multimedia_resource","req":false,"short":"Whether this resource is considered to be multimedia","type":"`$BOOLEAN`","index$":10},{"active":true,"name":"is_teacher_resource","req":false,"short":"Whether this resource is considered to be educational","type":"`$BOOLEAN`","index$":11},{"active":true,"name":"lake_guid","req":false,"short":"Unique UUID of this resource in LAKE, our DAMS.","type":"`$ANY`","index$":12},{"active":true,"name":"source_updated_at","req":false,"short":"Date and time the resource was updated in the source system","type":"`$ANY`","index$":13},{"active":true,"name":"suggest_autocomplete_all","req":false,"short":"Internal field to power the `/autosuggest` endpoint.","type":"`$ANY`","index$":14},{"active":true,"name":"suggest_autocomplete_boosted","req":false,"short":"Internal field to power the `/autocomplete` endpoint.","type":"`$ANY`","index$":15},{"active":true,"name":"timestamp","req":false,"short":"Date and time the record was updated in the aggregator search index","type":"`$ANY`","index$":16},{"active":true,"name":"title","req":false,"short":"Name of this mobile audio file – derived from the artwork and tour titles","type":"`$STRING`","index$":17},{"active":true,"name":"transcript","req":false,"short":"Text transcription of the audio file","type":"`$ANY`","index$":18},{"active":true,"name":"type","req":false,"short":"Type always takes one of the following values: image, sound, text, video","type":"`$ANY`","index$":19},{"active":true,"name":"updated_at","req":false,"short":"Date and time the record was updated in the aggregator database","type":"`$ANY`","index$":20},{"active":true,"name":"web_url","req":false,"short":"URL to the audio file","type":"`$ANY`","index$":21}],"id":{"field":"id","name":"id"},"name":"sound","op":{"list":{"input":"data","name":"list","points":[{"active":true,"args":{},"contract":{"id":"GET /mobile-sounds","json":"{\"parameters\":[],\"produces\":[\"application/json\"],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"Successful operation\",\"schema\":{\"items\":{\"properties\":{\"alt_text\":{\"description\":\"Alternative text for the asset to describe it to people with low or no vision\"},\"api_link\":{\"description\":\"REST API link for this resource\"},\"api_model\":{\"description\":\"REST API resource type or endpoint\"},\"artwork_ids\":{\"description\":\"Unique identifiers of the artworks associated with this asset\"},\"artwork_titles\":{\"description\":\"Names of the artworks associated with this asset\"},\"content\":{\"description\":\"Text of or URL to the contents of this asset\"},\"content_e_tag\":{\"description\":\"Arbitrary unique identifier that changes when the binary file gets updated\"},\"credit_line\":{\"description\":\"Asset-specific copyright information\"},\"id\":{\"description\":\"Unique identifier of this resource. Taken from the source system.\"},\"is_educational_resource\":{\"description\":\"Whether this resource is considered to be educational\"},\"is_multimedia_resource\":{\"description\":\"Whether this resource is considered to be multimedia\"},\"is_teacher_resource\":{\"description\":\"Whether this resource is considered to be educational\"},\"lake_guid\":{\"description\":\"Unique UUID of this resource in LAKE, our DAMS.\"},\"source_updated_at\":{\"description\":\"Date and time the resource was updated in the source system\"},\"suggest_autocomplete_all\":{\"description\":\"Internal field to power the `/autosuggest` endpoint. Do not use directly.\"},\"suggest_autocomplete_boosted\":{\"description\":\"Internal field to power the `/autocomplete` endpoint. Do not use directly.\"},\"timestamp\":{\"description\":\"Date and time the record was updated in the aggregator search index\"},\"title\":{\"description\":\"Name of this mobile audio file – derived from the artwork and tour titles\"},\"transcript\":{\"description\":\"Text transcription of the audio file\"},\"type\":{\"description\":\"Type always takes one of the following values: image, sound, text, video\"},\"updated_at\":{\"description\":\"Date and time the record was updated in the aggregator database\"},\"web_url\":{\"description\":\"URL to the audio file\"}},\"type\":\"object\"},\"type\":\"array\"}},\"default\":{\"description\":\"error\",\"schema\":{\"properties\":{\"detail\":{\"type\":\"string\"},\"error\":{\"type\":\"string\"},\"status\":{\"type\":\"integer\"}},\"required\":[\"status\",\"error\",\"detail\"]}}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/mobile-sounds","segments":[{"lit":"mobile-sounds"}],"select":{},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0},{"active":true,"args":{},"contract":{"id":"GET /sounds","json":"{\"parameters\":[],\"produces\":[\"application/json\"],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"Successful operation\",\"schema\":{\"items\":{\"properties\":{\"alt_text\":{\"description\":\"Alternative text for the asset to describe it to people with low or no vision\"},\"api_link\":{\"description\":\"REST API link for this resource\"},\"api_model\":{\"description\":\"REST API resource type or endpoint\"},\"artwork_ids\":{\"description\":\"Unique identifiers of the artworks associated with this asset\"},\"artwork_titles\":{\"description\":\"Names of the artworks associated with this asset\"},\"content\":{\"description\":\"Text of or URL to the contents of this asset\"},\"content_e_tag\":{\"description\":\"Arbitrary unique identifier that changes when the binary file gets updated\"},\"credit_line\":{\"description\":\"Asset-specific copyright information\"},\"id\":{\"description\":\"Unique identifier of this resource. Taken from the source system.\"},\"is_educational_resource\":{\"description\":\"Whether this resource is considered to be educational\"},\"is_multimedia_resource\":{\"description\":\"Whether this resource is considered to be multimedia\"},\"is_teacher_resource\":{\"description\":\"Whether this resource is considered to be educational\"},\"lake_guid\":{\"description\":\"Unique UUID of this resource in LAKE, our DAMS.\"},\"source_updated_at\":{\"description\":\"Date and time the resource was updated in the source system\"},\"suggest_autocomplete_all\":{\"description\":\"Internal field to power the `/autosuggest` endpoint. Do not use directly.\"},\"suggest_autocomplete_boosted\":{\"description\":\"Internal field to power the `/autocomplete` endpoint. Do not use directly.\"},\"timestamp\":{\"description\":\"Date and time the record was updated in the aggregator search index\"},\"title\":{\"description\":\"Name of this mobile audio file – derived from the artwork and tour titles\"},\"transcript\":{\"description\":\"Text transcription of the audio file\"},\"type\":{\"description\":\"Type always takes one of the following values: image, sound, text, video\"},\"updated_at\":{\"description\":\"Date and time the record was updated in the aggregator database\"},\"web_url\":{\"description\":\"URL to the audio file\"}},\"type\":\"object\"},\"type\":\"array\"}},\"default\":{\"description\":\"error\",\"schema\":{\"properties\":{\"detail\":{\"type\":\"string\"},\"error\":{\"type\":\"string\"},\"status\":{\"type\":\"integer\"}},\"required\":[\"status\",\"error\",\"detail\"]}}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/sounds","segments":[{"lit":"sounds"}],"select":{},"transform":{"req":"`reqdata`","res":"`body`"},"index$":1}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"id","orig":"id","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"GET /mobile-sounds/{id}","json":"{\"parameters\":[{\"in\":\"path\",\"name\":\"id\",\"required\":true,\"type\":\"string\"}],\"produces\":[\"application/json\"],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"Successful operation\",\"schema\":{\"items\":{\"properties\":{\"alt_text\":{\"description\":\"Alternative text for the asset to describe it to people with low or no vision\"},\"api_link\":{\"description\":\"REST API link for this resource\"},\"api_model\":{\"description\":\"REST API resource type or endpoint\"},\"artwork_ids\":{\"description\":\"Unique identifiers of the artworks associated with this asset\"},\"artwork_titles\":{\"description\":\"Names of the artworks associated with this asset\"},\"content\":{\"description\":\"Text of or URL to the contents of this asset\"},\"content_e_tag\":{\"description\":\"Arbitrary unique identifier that changes when the binary file gets updated\"},\"credit_line\":{\"description\":\"Asset-specific copyright information\"},\"id\":{\"description\":\"Unique identifier of this resource. Taken from the source system.\"},\"is_educational_resource\":{\"description\":\"Whether this resource is considered to be educational\"},\"is_multimedia_resource\":{\"description\":\"Whether this resource is considered to be multimedia\"},\"is_teacher_resource\":{\"description\":\"Whether this resource is considered to be educational\"},\"lake_guid\":{\"description\":\"Unique UUID of this resource in LAKE, our DAMS.\"},\"source_updated_at\":{\"description\":\"Date and time the resource was updated in the source system\"},\"suggest_autocomplete_all\":{\"description\":\"Internal field to power the `/autosuggest` endpoint. Do not use directly.\"},\"suggest_autocomplete_boosted\":{\"description\":\"Internal field to power the `/autocomplete` endpoint. Do not use directly.\"},\"timestamp\":{\"description\":\"Date and time the record was updated in the aggregator search index\"},\"title\":{\"description\":\"Name of this mobile audio file – derived from the artwork and tour titles\"},\"transcript\":{\"description\":\"Text transcription of the audio file\"},\"type\":{\"description\":\"Type always takes one of the following values: image, sound, text, video\"},\"updated_at\":{\"description\":\"Date and time the record was updated in the aggregator database\"},\"web_url\":{\"description\":\"URL to the audio file\"}},\"type\":\"object\"},\"type\":\"array\"}},\"default\":{\"description\":\"error\",\"schema\":{\"properties\":{\"detail\":{\"type\":\"string\"},\"error\":{\"type\":\"string\"},\"status\":{\"type\":\"integer\"}},\"required\":[\"status\",\"error\",\"detail\"]}}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/mobile-sounds/{id}","segments":[{"lit":"mobile-sounds"},{"var":"id"}],"select":{"exist":["id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0},{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"id","orig":"id","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"GET /sounds/{id}","json":"{\"parameters\":[{\"in\":\"path\",\"name\":\"id\",\"required\":true,\"type\":\"string\"}],\"produces\":[\"application/json\"],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"Successful operation\",\"schema\":{\"items\":{\"properties\":{\"alt_text\":{\"description\":\"Alternative text for the asset to describe it to people with low or no vision\"},\"api_link\":{\"description\":\"REST API link for this resource\"},\"api_model\":{\"description\":\"REST API resource type or endpoint\"},\"artwork_ids\":{\"description\":\"Unique identifiers of the artworks associated with this asset\"},\"artwork_titles\":{\"description\":\"Names of the artworks associated with this asset\"},\"content\":{\"description\":\"Text of or URL to the contents of this asset\"},\"content_e_tag\":{\"description\":\"Arbitrary unique identifier that changes when the binary file gets updated\"},\"credit_line\":{\"description\":\"Asset-specific copyright information\"},\"id\":{\"description\":\"Unique identifier of this resource. Taken from the source system.\"},\"is_educational_resource\":{\"description\":\"Whether this resource is considered to be educational\"},\"is_multimedia_resource\":{\"description\":\"Whether this resource is considered to be multimedia\"},\"is_teacher_resource\":{\"description\":\"Whether this resource is considered to be educational\"},\"lake_guid\":{\"description\":\"Unique UUID of this resource in LAKE, our DAMS.\"},\"source_updated_at\":{\"description\":\"Date and time the resource was updated in the source system\"},\"suggest_autocomplete_all\":{\"description\":\"Internal field to power the `/autosuggest` endpoint. Do not use directly.\"},\"suggest_autocomplete_boosted\":{\"description\":\"Internal field to power the `/autocomplete` endpoint. Do not use directly.\"},\"timestamp\":{\"description\":\"Date and time the record was updated in the aggregator search index\"},\"title\":{\"description\":\"Name of this mobile audio file – derived from the artwork and tour titles\"},\"transcript\":{\"description\":\"Text transcription of the audio file\"},\"type\":{\"description\":\"Type always takes one of the following values: image, sound, text, video\"},\"updated_at\":{\"description\":\"Date and time the record was updated in the aggregator database\"},\"web_url\":{\"description\":\"URL to the audio file\"}},\"type\":\"object\"},\"type\":\"array\"}},\"default\":{\"description\":\"error\",\"schema\":{\"properties\":{\"detail\":{\"type\":\"string\"},\"error\":{\"type\":\"string\"},\"status\":{\"type\":\"integer\"}},\"required\":[\"status\",\"error\",\"detail\"]}}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/sounds/{id}","segments":[{"lit":"sounds"},{"var":"id"}],"select":{"exist":["id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":1}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"sound","name__orig":"sound","Name":"Sound","name_":"sound","name-":"sound","NAME":"SOUND","index$":30}, {"active":true,"entity":"sound","key$":"BasicSoundFlow","kind":"basic","name":"BasicSoundFlow","param":{},"step":[{"active":true,"data":{},"input":{},"match":{},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"sound_ref01"}}],"index$":0},{"active":true,"data":{},"input":{"ref":"sound_ref01","srcdatavar":"sound_ref01_data","suffix":"_dt0"},"match":{"id":"sound01"},"op":"load","spec":[],"valid":[{"apply":"TextFieldMark","def":{"mark":"Mark01-sound_ref01"}}],"index$":1}]}, 'Sound')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let sound_ref01_data = Object.values(setup.data.existing.sound)[0] as any

    // LIST
    const sound_ref01_ent = client.Sound()
    const sound_ref01_match: any = {}

    const sound_ref01_list = (await sound_ref01_ent.list(sound_ref01_match)).map((e: any) => e.data())


    // LOAD
    const sound_ref01_match_dt0: any = {}
    sound_ref01_match_dt0.id = sound_ref01_data.id
    const sound_ref01_data_dt0 = (await sound_ref01_ent.load(sound_ref01_match_dt0)).data()
    assert(sound_ref01_data_dt0.id === sound_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/sound/SoundTestData.json')

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
    ['sound01','sound02','sound03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'ART_INSTITUTE_OF_CHICAGO_TEST_SOUND_ENTID': idmap,
    'ART_INSTITUTE_OF_CHICAGO_TEST_LIVE': 'FALSE',
    'ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['ART_INSTITUTE_OF_CHICAGO_TEST_SOUND_ENTID']

  const live = 'TRUE' === env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['ART_INSTITUTE_OF_CHICAGO_TEST_SOUND_ENTID']
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
  
