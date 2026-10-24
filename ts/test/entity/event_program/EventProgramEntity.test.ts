

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


describe('EventProgramEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE.
  afterEach(liveDelay('ART_INSTITUTE_OF_CHICAGO_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = ArtInstituteOfChicagoSDK.test()
    const ent = testsdk.EventProgram()
    assert(null != ent)
  })


  class FailHook extends BaseFeature {
    name = 'failhook'
    version = '0.0.1'
    active = true
    unexpected = 0
    init() { }
    PreSpec() { throw new Error('event_program hook failed') }
    PreUnexpected() { this.unexpected++ }
  }

  test('stream-error', async () => {
    const offline = { net: { offline: true } }
    await assert.rejects(async () => {
      for await (const _item of ArtInstituteOfChicagoSDK.test(offline).EventProgram().stream('list')) { }
    }, /offline/)

    for await (const _item of ArtInstituteOfChicagoSDK.test(offline).EventProgram()
      .stream('list', undefined, { ctrl: { throw: false } })) { }

    if (null != (config as any).feature?.rbac) {
      const denied = ArtInstituteOfChicagoSDK.test(undefined, { feature: { rbac: { active: true, deny: true } } })
      await assert.rejects(async () => {
        for await (const _item of denied.EventProgram().stream('list')) { }
      }, (err: any) => 'rbac_denied' === err.code)
    }
  })

  test('stream-ctrl', async () => {
    const explain: any = {}
    const ctrl: any = { explain }
    for await (const _item of ArtInstituteOfChicagoSDK.test().EventProgram().stream('list', undefined, { ctrl })) { }
    assert.deepStrictEqual(Object.keys(ctrl), ['explain'])
    assert(explain === ctrl.explain && 0 < Object.keys(explain).length)
  })

  test('unexpected', async () => {
    const hook = new FailHook()
    const client = new ArtInstituteOfChicagoSDK({ feature: { test: { active: true } }, extend: [hook] })
    await assert.rejects(client.EventProgram().list(), /hook failed/)
    assert(0 < hook.unexpected)

    const fired = hook.unexpected
    assert.strictEqual(await client.EventProgram().list(undefined, { throw: false }), undefined)
    assert(fired < hook.unexpected)
  })

  test('validate', async (t) => {
    if (null == (config as any).feature?.validate) {
      t.skip('feature not present in this SDK: validate')
      return
    }
    const client = ArtInstituteOfChicagoSDK.test(undefined, { feature: { validate: { active: true } } })
    await assert.rejects(client.EventProgram().list({"id":1} as any),
      (err: any) => 'validate_failed' === err.code)
  })



  test('basic', async (t) => {

    const live = 'TRUE' === process.env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE
    for (const op of ['list', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'event_program.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"api_link":{"a":true,"h":"Api Link","n":"api_link","r":false,"sh":"REST API link for this resource","t":"`$ANY`","key$":"api_link","index$":0},"api_model":{"a":true,"h":"Api Model","n":"api_model","r":false,"sh":"REST API resource type or endpoint","t":"`$ANY`","key$":"api_model","index$":1},"id":{"a":true,"h":"Id","n":"id","r":false,"sh":"Unique identifier of this resource.","t":"`$STRING`","key$":"id","index$":2},"is_affiliate_group":{"a":true,"h":"Is Affiliate Group","n":"is_affiliate_group","r":false,"sh":"Whether this program represents an affiliate group","t":"`$BOOLEAN`","key$":"is_affiliate_group","index$":3},"is_event_host":{"a":true,"h":"Is Event Host","n":"is_event_host","r":false,"sh":"Whether this program represents an event host","t":"`$BOOLEAN`","key$":"is_event_host","index$":4},"source_updated_at":{"a":true,"h":"Source Updated At","n":"source_updated_at","r":false,"sh":"Date and time the resource was updated in the source system","t":"`$ANY`","key$":"source_updated_at","index$":5},"suggest_autocomplete_all":{"a":true,"h":"Suggest Autocomplete All","n":"suggest_autocomplete_all","r":false,"sh":"Internal field to power the `/autosuggest` endpoint.","t":"`$ANY`","key$":"suggest_autocomplete_all","index$":6},"suggest_autocomplete_boosted":{"a":true,"h":"Suggest Autocomplete Boosted","n":"suggest_autocomplete_boosted","r":false,"sh":"Internal field to power the `/autocomplete` endpoint.","t":"`$ANY`","key$":"suggest_autocomplete_boosted","index$":7},"timestamp":{"a":true,"h":"Timestamp","n":"timestamp","r":false,"sh":"Date and time the record was updated in the aggregator search index","t":"`$ANY`","key$":"timestamp","index$":8},"title":{"a":true,"h":"Title","n":"title","r":false,"sh":"The name of this resource","t":"`$STRING`","key$":"title","index$":9},"updated_at":{"a":true,"h":"Updated At","n":"updated_at","r":false,"sh":"Date and time the record was updated in the aggregator database","t":"`$ANY`","key$":"updated_at","index$":10}},"id":{"field":"id","name":"id"},"name":"event_program","op":{"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /event-programs","source":"openapi3","version":2},"g":{},"k":"http","m":"GET","o":"/event-programs","q":{},"r":{},"s":[{"lit":"event-programs"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /event-programs/{id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"GET","o":"/event-programs/{id}","q":{"exist":["id"]},"r":{},"s":[{"lit":"event-programs"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"event_program","name__orig":"event_program","Name":"EventProgram","name_":"event_program","name-":"event-program","NAME":"EVENT_PROGRAM","index$":14}, {"active":true,"entity":"event_program","key$":"BasicEventProgramFlow","kind":"basic","name":"BasicEventProgramFlow","param":{},"step":[{"a":true,"d":{},"i":{},"m":{},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"event_program_ref01"}}],"index$":0},{"a":true,"d":{},"i":{"ref":"event_program_ref01","srcdatavar":"event_program_ref01_data","suffix":"_dt0"},"m":{"id":"event_program01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-event_program_ref01"}}],"index$":1}]}, 'EventProgram', {"GET /event-programs":{"protocol":"http","parameters":[]},"GET /event-programs/{id}":{"protocol":"http","parameters":[{"name":"id","in":"path","type":"string","required":true,"x-ref":"#/parameters/id","index$":0}]}}, { strict: LIVE_STRICT, t })
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let event_program_ref01_data = Object.values(setup.data.existing.event_program)[0] as any

    // LIST
    const event_program_ref01_ent = client.EventProgram()
    const event_program_ref01_match: any = {}

    const event_program_ref01_list = (await event_program_ref01_ent.list(event_program_ref01_match)).map((e: any) => e.data())


    // LOAD
    const event_program_ref01_match_dt0: any = {}
    event_program_ref01_match_dt0.id = event_program_ref01_data.id
    const event_program_ref01_data_dt0 = (await event_program_ref01_ent.load(event_program_ref01_match_dt0)).data()
    assert(event_program_ref01_data_dt0.id === event_program_ref01_data.id)


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
      '../../../../.sdk/test/entity/event_program/EventProgramTestData.json')

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
    ['event_program01','event_program02','event_program03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_PROGRAM_ENTID': idmap,
    'ART_INSTITUTE_OF_CHICAGO_TEST_LIVE': 'FALSE',
    'ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_PROGRAM_ENTID']

  const live = 'TRUE' === env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_PROGRAM_ENTID']
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
  
