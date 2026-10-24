

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


describe('ImageEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE.
  afterEach(liveDelay('ART_INSTITUTE_OF_CHICAGO_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = ArtInstituteOfChicagoSDK.test()
    const ent = testsdk.Image()
    assert(null != ent)
  })


  class FailHook extends BaseFeature {
    name = 'failhook'
    version = '0.0.1'
    active = true
    unexpected = 0
    init() { }
    PreSpec() { throw new Error('image hook failed') }
    PreUnexpected() { this.unexpected++ }
  }

  test('stream-error', async () => {
    const offline = { net: { offline: true } }
    await assert.rejects(async () => {
      for await (const _item of ArtInstituteOfChicagoSDK.test(offline).Image().stream('list')) { }
    }, /offline/)

    for await (const _item of ArtInstituteOfChicagoSDK.test(offline).Image()
      .stream('list', undefined, { ctrl: { throw: false } })) { }

    if (null != (config as any).feature?.rbac) {
      const denied = ArtInstituteOfChicagoSDK.test(undefined, { feature: { rbac: { active: true, deny: true } } })
      await assert.rejects(async () => {
        for await (const _item of denied.Image().stream('list')) { }
      }, (err: any) => 'rbac_denied' === err.code)
    }
  })

  test('stream-ctrl', async () => {
    const explain: any = {}
    const ctrl: any = { explain }
    for await (const _item of ArtInstituteOfChicagoSDK.test().Image().stream('list', undefined, { ctrl })) { }
    assert.deepStrictEqual(Object.keys(ctrl), ['explain'])
    assert(explain === ctrl.explain && 0 < Object.keys(explain).length)
  })

  test('unexpected', async () => {
    const hook = new FailHook()
    const client = new ArtInstituteOfChicagoSDK({ feature: { test: { active: true } }, extend: [hook] })
    await assert.rejects(client.Image().list(), /hook failed/)
    assert(0 < hook.unexpected)

    const fired = hook.unexpected
    assert.strictEqual(await client.Image().list(undefined, { throw: false }), undefined)
    assert(fired < hook.unexpected)
  })

  test('validate', async (t) => {
    if (null == (config as any).feature?.validate) {
      t.skip('feature not present in this SDK: validate')
      return
    }
    const client = ArtInstituteOfChicagoSDK.test(undefined, { feature: { validate: { active: true } } })
    await assert.rejects(client.Image().list({"height":"x"} as any),
      (err: any) => 'validate_failed' === err.code)
  })



  test('basic', async (t) => {

    const live = 'TRUE' === process.env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE
    for (const op of ['list', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'image.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"ahash":{"a":true,"h":"Ahash","n":"ahash","r":false,"sh":"Image hash generated using ahash algorithm with 64 boolean subfields","t":"`$ANY`","key$":"ahash","index$":0},"alt_text":{"a":true,"h":"Alt Text","n":"alt_text","r":false,"sh":"Alternative text for the asset to describe it to people with low or no vision","t":"`$ANY`","key$":"alt_text","index$":1},"api_link":{"a":true,"h":"Api Link","n":"api_link","r":false,"sh":"REST API link for this resource","t":"`$ANY`","key$":"api_link","index$":2},"api_model":{"a":true,"h":"Api Model","n":"api_model","r":false,"sh":"REST API resource type or endpoint","t":"`$ANY`","key$":"api_model","index$":3},"artwork_ids":{"a":true,"h":"Artwork Ids","n":"artwork_ids","r":false,"sh":"Unique identifiers of the artworks associated with this asset","t":"`$ANY`","key$":"artwork_ids","index$":4},"artwork_titles":{"a":true,"h":"Artwork Titles","n":"artwork_titles","r":false,"sh":"Names of the artworks associated with this asset","t":"`$ANY`","key$":"artwork_titles","index$":5},"color":{"a":true,"h":"Color","n":"color","r":false,"sh":"Dominant color of this image in HSL","t":"`$ANY`","key$":"color","index$":6},"colorfulness":{"a":true,"h":"Colorfulness","n":"colorfulness","r":false,"sh":"Unbounded positive float representing an abstract measure of colorfulness.","t":"`$ANY`","key$":"colorfulness","index$":7},"content":{"a":true,"h":"Content","n":"content","r":false,"sh":"Text of or URL to the contents of this asset","t":"`$ANY`","key$":"content","index$":8},"content_e_tag":{"a":true,"h":"Content E Tag","n":"content_e_tag","r":false,"sh":"Arbitrary unique identifier that changes when the binary file gets updated","t":"`$ANY`","key$":"content_e_tag","index$":9},"credit_line":{"a":true,"h":"Credit Line","n":"credit_line","r":false,"sh":"Asset-specific copyright information","t":"`$ANY`","key$":"credit_line","index$":10},"fingerprint":{"a":true,"h":"Fingerprint","n":"fingerprint","r":false,"sh":"Image hashes: aHash, dHash, pHash, wHash","t":"`$ANY`","key$":"fingerprint","index$":11},"height":{"a":true,"h":"Height","n":"height","r":false,"sh":"Native height of the image","t":"`$NUMBER`","key$":"height","index$":12},"id":{"a":true,"h":"Id","n":"id","r":false,"sh":"Unique identifier of this resource.","t":"`$STRING`","key$":"id","index$":13},"iiif_url":{"a":true,"h":"Iiif Url","n":"iiif_url","r":false,"sh":"IIIF URL of this image","t":"`$ANY`","key$":"iiif_url","index$":14},"is_educational_resource":{"a":true,"h":"Is Educational Resource","n":"is_educational_resource","r":false,"sh":"Whether this resource is considered to be educational","t":"`$BOOLEAN`","key$":"is_educational_resource","index$":15},"is_multimedia_resource":{"a":true,"h":"Is Multimedia Resource","n":"is_multimedia_resource","r":false,"sh":"Whether this resource is considered to be multimedia","t":"`$BOOLEAN`","key$":"is_multimedia_resource","index$":16},"is_teacher_resource":{"a":true,"h":"Is Teacher Resource","n":"is_teacher_resource","r":false,"sh":"Whether this resource is considered to be educational","t":"`$BOOLEAN`","key$":"is_teacher_resource","index$":17},"lake_guid":{"a":true,"h":"Lake Guid","n":"lake_guid","r":false,"sh":"Unique UUID of this resource in LAKE, our DAMS.","t":"`$ANY`","key$":"lake_guid","index$":18},"lqip":{"a":true,"h":"Lqip","n":"lqip","r":false,"sh":"Low-quality image placeholder (LQIP).","t":"`$ANY`","key$":"lqip","index$":19},"phash":{"a":true,"h":"Phash","n":"phash","r":false,"sh":"Image hash generated using phash algorithm with 64 boolean subfields","t":"`$ANY`","key$":"phash","index$":20},"source_updated_at":{"a":true,"h":"Source Updated At","n":"source_updated_at","r":false,"sh":"Date and time the resource was updated in the source system","t":"`$ANY`","key$":"source_updated_at","index$":21},"suggest_autocomplete_all":{"a":true,"h":"Suggest Autocomplete All","n":"suggest_autocomplete_all","r":false,"sh":"Internal field to power the `/autosuggest` endpoint.","t":"`$ANY`","key$":"suggest_autocomplete_all","index$":22},"suggest_autocomplete_boosted":{"a":true,"h":"Suggest Autocomplete Boosted","n":"suggest_autocomplete_boosted","r":false,"sh":"Internal field to power the `/autocomplete` endpoint.","t":"`$ANY`","key$":"suggest_autocomplete_boosted","index$":23},"timestamp":{"a":true,"h":"Timestamp","n":"timestamp","r":false,"sh":"Date and time the record was updated in the aggregator search index","t":"`$ANY`","key$":"timestamp","index$":24},"title":{"a":true,"h":"Title","n":"title","r":false,"sh":"The name of this resource","t":"`$STRING`","key$":"title","index$":25},"type":{"a":true,"h":"Type","n":"type","r":false,"sh":"Type always takes one of the following values: image, sound, text, video","t":"`$ANY`","key$":"type","index$":26},"updated_at":{"a":true,"h":"Updated At","n":"updated_at","r":false,"sh":"Date and time the record was updated in the aggregator database","t":"`$ANY`","key$":"updated_at","index$":27},"width":{"a":true,"h":"Width","n":"width","r":false,"sh":"Native width of the image","t":"`$NUMBER`","key$":"width","index$":28}},"id":{"field":"id","name":"id"},"name":"image","op":{"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /images","source":"openapi3","version":2},"g":{},"k":"http","m":"GET","o":"/images","q":{},"r":{},"s":[{"lit":"images"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /images/{id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"GET","o":"/images/{id}","q":{"exist":["id"]},"r":{},"s":[{"lit":"images"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"image","name__orig":"image","Name":"Image","name_":"image","name-":"image","NAME":"IMAGE","index$":20}, {"active":true,"entity":"image","key$":"BasicImageFlow","kind":"basic","name":"BasicImageFlow","param":{},"step":[{"a":true,"d":{},"i":{},"m":{},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"image_ref01"}}],"index$":0},{"a":true,"d":{},"i":{"ref":"image_ref01","srcdatavar":"image_ref01_data","suffix":"_dt0"},"m":{"id":"image01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-image_ref01"}}],"index$":1}]}, 'Image', {"GET /images":{"protocol":"http","parameters":[]},"GET /images/{id}":{"protocol":"http","parameters":[{"name":"id","in":"path","type":"string","required":true,"x-ref":"#/parameters/id","index$":0}]}}, { strict: LIVE_STRICT, t })
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let image_ref01_data = Object.values(setup.data.existing.image)[0] as any

    // LIST
    const image_ref01_ent = client.Image()
    const image_ref01_match: any = {}

    const image_ref01_list = (await image_ref01_ent.list(image_ref01_match)).map((e: any) => e.data())


    // LOAD
    const image_ref01_match_dt0: any = {}
    image_ref01_match_dt0.id = image_ref01_data.id
    const image_ref01_data_dt0 = (await image_ref01_ent.load(image_ref01_match_dt0)).data()
    assert(image_ref01_data_dt0.id === image_ref01_data.id)


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
      '../../../../.sdk/test/entity/image/ImageTestData.json')

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
    ['image01','image02','image03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'ART_INSTITUTE_OF_CHICAGO_TEST_IMAGE_ENTID': idmap,
    'ART_INSTITUTE_OF_CHICAGO_TEST_LIVE': 'FALSE',
    'ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['ART_INSTITUTE_OF_CHICAGO_TEST_IMAGE_ENTID']

  const live = 'TRUE' === env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['ART_INSTITUTE_OF_CHICAGO_TEST_IMAGE_ENTID']
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
  
