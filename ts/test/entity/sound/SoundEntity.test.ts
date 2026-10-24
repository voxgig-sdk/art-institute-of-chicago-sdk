

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


describe('SoundEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE.
  afterEach(liveDelay('ART_INSTITUTE_OF_CHICAGO_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = ArtInstituteOfChicagoSDK.test()
    const ent = testsdk.Sound()
    assert(null != ent)
  })


  class FailHook extends BaseFeature {
    name = 'failhook'
    version = '0.0.1'
    active = true
    unexpected = 0
    init() { }
    PreSpec() { throw new Error('sound hook failed') }
    PreUnexpected() { this.unexpected++ }
  }

  test('stream-error', async () => {
    const offline = { net: { offline: true } }
    await assert.rejects(async () => {
      for await (const _item of ArtInstituteOfChicagoSDK.test(offline).Sound().stream('list')) { }
    }, /offline/)

    for await (const _item of ArtInstituteOfChicagoSDK.test(offline).Sound()
      .stream('list', undefined, { ctrl: { throw: false } })) { }

    if (null != (config as any).feature?.rbac) {
      const denied = ArtInstituteOfChicagoSDK.test(undefined, { feature: { rbac: { active: true, deny: true } } })
      await assert.rejects(async () => {
        for await (const _item of denied.Sound().stream('list')) { }
      }, (err: any) => 'rbac_denied' === err.code)
    }
  })

  test('stream-ctrl', async () => {
    const explain: any = {}
    const ctrl: any = { explain }
    for await (const _item of ArtInstituteOfChicagoSDK.test().Sound().stream('list', undefined, { ctrl })) { }
    assert.deepStrictEqual(Object.keys(ctrl), ['explain'])
    assert(explain === ctrl.explain && 0 < Object.keys(explain).length)
  })

  test('unexpected', async () => {
    const hook = new FailHook()
    const client = new ArtInstituteOfChicagoSDK({ feature: { test: { active: true } }, extend: [hook] })
    await assert.rejects(client.Sound().list(), /hook failed/)
    assert(0 < hook.unexpected)

    const fired = hook.unexpected
    assert.strictEqual(await client.Sound().list(undefined, { throw: false }), undefined)
    assert(fired < hook.unexpected)
  })

  test('validate', async (t) => {
    if (null == (config as any).feature?.validate) {
      t.skip('feature not present in this SDK: validate')
      return
    }
    const client = ArtInstituteOfChicagoSDK.test(undefined, { feature: { validate: { active: true } } })
    await assert.rejects(client.Sound().list({"id":1} as any),
      (err: any) => 'validate_failed' === err.code)
  })



  test('basic', async (t) => {

    const live = 'TRUE' === process.env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE
    for (const op of ['list', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'sound.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"alt_text":{"a":true,"h":"Alt Text","n":"alt_text","r":false,"sh":"Alternative text for the asset to describe it to people with low or no vision","t":"`$ANY`","key$":"alt_text","index$":0},"api_link":{"a":true,"h":"Api Link","n":"api_link","r":false,"sh":"REST API link for this resource","t":"`$ANY`","key$":"api_link","index$":1},"api_model":{"a":true,"h":"Api Model","n":"api_model","r":false,"sh":"REST API resource type or endpoint","t":"`$ANY`","key$":"api_model","index$":2},"artwork_ids":{"a":true,"h":"Artwork Ids","n":"artwork_ids","r":false,"sh":"Unique identifiers of the artworks associated with this asset","t":"`$ANY`","key$":"artwork_ids","index$":3},"artwork_titles":{"a":true,"h":"Artwork Titles","n":"artwork_titles","r":false,"sh":"Names of the artworks associated with this asset","t":"`$ANY`","key$":"artwork_titles","index$":4},"content":{"a":true,"h":"Content","n":"content","r":false,"sh":"Text of or URL to the contents of this asset","t":"`$ANY`","key$":"content","index$":5},"content_e_tag":{"a":true,"h":"Content E Tag","n":"content_e_tag","r":false,"sh":"Arbitrary unique identifier that changes when the binary file gets updated","t":"`$ANY`","key$":"content_e_tag","index$":6},"credit_line":{"a":true,"h":"Credit Line","n":"credit_line","r":false,"sh":"Asset-specific copyright information","t":"`$ANY`","key$":"credit_line","index$":7},"id":{"a":true,"h":"Id","n":"id","r":false,"sh":"Unique identifier of this resource.","t":"`$STRING`","key$":"id","index$":8},"is_educational_resource":{"a":true,"h":"Is Educational Resource","n":"is_educational_resource","r":false,"sh":"Whether this resource is considered to be educational","t":"`$BOOLEAN`","key$":"is_educational_resource","index$":9},"is_multimedia_resource":{"a":true,"h":"Is Multimedia Resource","n":"is_multimedia_resource","r":false,"sh":"Whether this resource is considered to be multimedia","t":"`$BOOLEAN`","key$":"is_multimedia_resource","index$":10},"is_teacher_resource":{"a":true,"h":"Is Teacher Resource","n":"is_teacher_resource","r":false,"sh":"Whether this resource is considered to be educational","t":"`$BOOLEAN`","key$":"is_teacher_resource","index$":11},"lake_guid":{"a":true,"h":"Lake Guid","n":"lake_guid","r":false,"sh":"Unique UUID of this resource in LAKE, our DAMS.","t":"`$ANY`","key$":"lake_guid","index$":12},"source_updated_at":{"a":true,"h":"Source Updated At","n":"source_updated_at","r":false,"sh":"Date and time the resource was updated in the source system","t":"`$ANY`","key$":"source_updated_at","index$":13},"suggest_autocomplete_all":{"a":true,"h":"Suggest Autocomplete All","n":"suggest_autocomplete_all","r":false,"sh":"Internal field to power the `/autosuggest` endpoint.","t":"`$ANY`","key$":"suggest_autocomplete_all","index$":14},"suggest_autocomplete_boosted":{"a":true,"h":"Suggest Autocomplete Boosted","n":"suggest_autocomplete_boosted","r":false,"sh":"Internal field to power the `/autocomplete` endpoint.","t":"`$ANY`","key$":"suggest_autocomplete_boosted","index$":15},"timestamp":{"a":true,"h":"Timestamp","n":"timestamp","r":false,"sh":"Date and time the record was updated in the aggregator search index","t":"`$ANY`","key$":"timestamp","index$":16},"title":{"a":true,"h":"Title","n":"title","r":false,"sh":"Name of this mobile audio file – derived from the artwork and tour titles","t":"`$STRING`","key$":"title","index$":17},"transcript":{"a":true,"h":"Transcript","n":"transcript","r":false,"sh":"Text transcription of the audio file","t":"`$ANY`","key$":"transcript","index$":18},"type":{"a":true,"h":"Type","n":"type","r":false,"sh":"Type always takes one of the following values: image, sound, text, video","t":"`$ANY`","key$":"type","index$":19},"updated_at":{"a":true,"h":"Updated At","n":"updated_at","r":false,"sh":"Date and time the record was updated in the aggregator database","t":"`$ANY`","key$":"updated_at","index$":20},"web_url":{"a":true,"h":"Web Url","n":"web_url","r":false,"sh":"URL to the audio file","t":"`$ANY`","key$":"web_url","index$":21}},"id":{"field":"id","name":"id"},"name":"sound","op":{"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /mobile-sounds","source":"openapi3","version":2},"g":{},"k":"http","m":"GET","o":"/mobile-sounds","q":{},"r":{},"s":[{"lit":"mobile-sounds"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0},{"a":true,"co":{"id":"GET /sounds","source":"openapi3","version":2},"g":{},"k":"http","m":"GET","o":"/sounds","q":{},"r":{},"s":[{"lit":"sounds"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":1}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /mobile-sounds/{id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"GET","o":"/mobile-sounds/{id}","q":{"exist":["id"]},"r":{},"s":[{"lit":"mobile-sounds"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0},{"a":true,"co":{"id":"GET /sounds/{id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"GET","o":"/sounds/{id}","q":{"exist":["id"]},"r":{},"s":[{"lit":"sounds"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":1}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"sound","name__orig":"sound","Name":"Sound","name_":"sound","name-":"sound","NAME":"SOUND","index$":30}, {"active":true,"entity":"sound","key$":"BasicSoundFlow","kind":"basic","name":"BasicSoundFlow","param":{},"step":[{"a":true,"d":{},"i":{},"m":{},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"sound_ref01"}}],"index$":0},{"a":true,"d":{},"i":{"ref":"sound_ref01","srcdatavar":"sound_ref01_data","suffix":"_dt0"},"m":{"id":"sound01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-sound_ref01"}}],"index$":1}]}, 'Sound', {"GET /mobile-sounds":{"protocol":"http","parameters":[]},"GET /sounds":{"protocol":"http","parameters":[]},"GET /mobile-sounds/{id}":{"protocol":"http","parameters":[{"name":"id","in":"path","type":"string","required":true,"x-ref":"#/parameters/id","index$":0}]},"GET /sounds/{id}":{"protocol":"http","parameters":[{"name":"id","in":"path","type":"string","required":true,"x-ref":"#/parameters/id","index$":0}]}}, { strict: LIVE_STRICT, t })
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
  
