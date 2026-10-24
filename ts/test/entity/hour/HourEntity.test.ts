

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


describe('HourEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE.
  afterEach(liveDelay('ART_INSTITUTE_OF_CHICAGO_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = ArtInstituteOfChicagoSDK.test()
    const ent = testsdk.Hour()
    assert(null != ent)
  })


  class FailHook extends BaseFeature {
    name = 'failhook'
    version = '0.0.1'
    active = true
    unexpected = 0
    init() { }
    PreSpec() { throw new Error('hour hook failed') }
    PreUnexpected() { this.unexpected++ }
  }

  test('stream-error', async () => {
    const offline = { net: { offline: true } }
    await assert.rejects(async () => {
      for await (const _item of ArtInstituteOfChicagoSDK.test(offline).Hour().stream('list')) { }
    }, /offline/)

    for await (const _item of ArtInstituteOfChicagoSDK.test(offline).Hour()
      .stream('list', undefined, { ctrl: { throw: false } })) { }

    if (null != (config as any).feature?.rbac) {
      const denied = ArtInstituteOfChicagoSDK.test(undefined, { feature: { rbac: { active: true, deny: true } } })
      await assert.rejects(async () => {
        for await (const _item of denied.Hour().stream('list')) { }
      }, (err: any) => 'rbac_denied' === err.code)
    }
  })

  test('stream-ctrl', async () => {
    const explain: any = {}
    const ctrl: any = { explain }
    for await (const _item of ArtInstituteOfChicagoSDK.test().Hour().stream('list', undefined, { ctrl })) { }
    assert.deepStrictEqual(Object.keys(ctrl), ['explain'])
    assert(explain === ctrl.explain && 0 < Object.keys(explain).length)
  })

  test('unexpected', async () => {
    const hook = new FailHook()
    const client = new ArtInstituteOfChicagoSDK({ feature: { test: { active: true } }, extend: [hook] })
    await assert.rejects(client.Hour().list(), /hook failed/)
    assert(0 < hook.unexpected)

    const fired = hook.unexpected
    assert.strictEqual(await client.Hour().list(undefined, { throw: false }), undefined)
    assert(fired < hook.unexpected)
  })

  test('validate', async (t) => {
    if (null == (config as any).feature?.validate) {
      t.skip('feature not present in this SDK: validate')
      return
    }
    const client = ArtInstituteOfChicagoSDK.test(undefined, { feature: { validate: { active: true } } })
    await assert.rejects(client.Hour().list({"id":1} as any),
      (err: any) => 'validate_failed' === err.code)
  })



  test('basic', async (t) => {

    const live = 'TRUE' === process.env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE
    for (const op of ['list', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'hour.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"additional_text":{"a":true,"h":"Additional Text","n":"additional_text","r":false,"sh":"Additional information about the hours","t":"`$ANY`","key$":"additional_text","index$":0},"api_link":{"a":true,"h":"Api Link","n":"api_link","r":false,"sh":"REST API link for this resource","t":"`$ANY`","key$":"api_link","index$":1},"api_model":{"a":true,"h":"Api Model","n":"api_model","r":false,"sh":"REST API resource type or endpoint","t":"`$ANY`","key$":"api_model","index$":2},"friday_is_closed":{"a":true,"h":"Friday Is Closed","n":"friday_is_closed","r":false,"sh":"Whether the museum is closed on Fridays","t":"`$ANY`","key$":"friday_is_closed","index$":3},"friday_member_close":{"a":true,"h":"Friday Member Close","n":"friday_member_close","r":false,"sh":"The time member hours ends on Fridays","t":"`$ANY`","key$":"friday_member_close","index$":4},"friday_member_open":{"a":true,"h":"Friday Member Open","n":"friday_member_open","r":false,"sh":"The time member hours starts on Fridays","t":"`$ANY`","key$":"friday_member_open","index$":5},"friday_public_close":{"a":true,"h":"Friday Public Close","n":"friday_public_close","r":false,"sh":"The time public hours ends on Fridays","t":"`$ANY`","key$":"friday_public_close","index$":6},"friday_public_open":{"a":true,"h":"Friday Public Open","n":"friday_public_open","r":false,"sh":"The time public hours starts on Fridays","t":"`$ANY`","key$":"friday_public_open","index$":7},"id":{"a":true,"h":"Id","n":"id","r":false,"sh":"Unique identifier of this resource.","t":"`$STRING`","key$":"id","index$":8},"monday_is_closed":{"a":true,"h":"Monday Is Closed","n":"monday_is_closed","r":false,"sh":"Whether the museum is closed on Mondays","t":"`$ANY`","key$":"monday_is_closed","index$":9},"monday_member_close":{"a":true,"h":"Monday Member Close","n":"monday_member_close","r":false,"sh":"The time member hours ends on Mondays","t":"`$ANY`","key$":"monday_member_close","index$":10},"monday_member_open":{"a":true,"h":"Monday Member Open","n":"monday_member_open","r":false,"sh":"The time member hours starts on Mondays","t":"`$ANY`","key$":"monday_member_open","index$":11},"monday_public_close":{"a":true,"h":"Monday Public Close","n":"monday_public_close","r":false,"sh":"The time public hours ends on Mondays","t":"`$ANY`","key$":"monday_public_close","index$":12},"monday_public_open":{"a":true,"h":"Monday Public Open","n":"monday_public_open","r":false,"sh":"The time public hours starts on Mondays","t":"`$ANY`","key$":"monday_public_open","index$":13},"saturday_is_closed":{"a":true,"h":"Saturday Is Closed","n":"saturday_is_closed","r":false,"sh":"Whether the museum is closed on Saturdays","t":"`$ANY`","key$":"saturday_is_closed","index$":14},"saturday_member_close":{"a":true,"h":"Saturday Member Close","n":"saturday_member_close","r":false,"sh":"The time member hours ends on Saturdays","t":"`$ANY`","key$":"saturday_member_close","index$":15},"saturday_member_open":{"a":true,"h":"Saturday Member Open","n":"saturday_member_open","r":false,"sh":"The time member hours starts on Saturdays","t":"`$ANY`","key$":"saturday_member_open","index$":16},"saturday_public_close":{"a":true,"h":"Saturday Public Close","n":"saturday_public_close","r":false,"sh":"The time public hours ends on Saturdays","t":"`$ANY`","key$":"saturday_public_close","index$":17},"saturday_public_open":{"a":true,"h":"Saturday Public Open","n":"saturday_public_open","r":false,"sh":"The time public hours starts on Saturdays","t":"`$ANY`","key$":"saturday_public_open","index$":18},"source_updated_at":{"a":true,"h":"Source Updated At","n":"source_updated_at","r":false,"sh":"Date and time the resource was updated in the source system","t":"`$ANY`","key$":"source_updated_at","index$":19},"suggest_autocomplete_all":{"a":true,"h":"Suggest Autocomplete All","n":"suggest_autocomplete_all","r":false,"sh":"Internal field to power the `/autosuggest` endpoint.","t":"`$ANY`","key$":"suggest_autocomplete_all","index$":20},"suggest_autocomplete_boosted":{"a":true,"h":"Suggest Autocomplete Boosted","n":"suggest_autocomplete_boosted","r":false,"sh":"Internal field to power the `/autocomplete` endpoint.","t":"`$ANY`","key$":"suggest_autocomplete_boosted","index$":21},"summary":{"a":true,"h":"Summary","n":"summary","r":false,"sh":"Readable summary of the hours","t":"`$ANY`","key$":"summary","index$":22},"sunday_is_closed":{"a":true,"h":"Sunday Is Closed","n":"sunday_is_closed","r":false,"sh":"Whether the museum is closed on Sundays","t":"`$ANY`","key$":"sunday_is_closed","index$":23},"sunday_member_close":{"a":true,"h":"Sunday Member Close","n":"sunday_member_close","r":false,"sh":"The time member hours ends on Sundays","t":"`$ANY`","key$":"sunday_member_close","index$":24},"sunday_member_open":{"a":true,"h":"Sunday Member Open","n":"sunday_member_open","r":false,"sh":"The time member hours starts on Sundays","t":"`$ANY`","key$":"sunday_member_open","index$":25},"sunday_public_close":{"a":true,"h":"Sunday Public Close","n":"sunday_public_close","r":false,"sh":"The time public hours ends on Sundays","t":"`$ANY`","key$":"sunday_public_close","index$":26},"sunday_public_open":{"a":true,"h":"Sunday Public Open","n":"sunday_public_open","r":false,"sh":"The time public hours starts on Sundays","t":"`$ANY`","key$":"sunday_public_open","index$":27},"thursday_is_closed":{"a":true,"h":"Thursday Is Closed","n":"thursday_is_closed","r":false,"sh":"Whether the museum is closed on Thursdays","t":"`$ANY`","key$":"thursday_is_closed","index$":28},"thursday_member_close":{"a":true,"h":"Thursday Member Close","n":"thursday_member_close","r":false,"sh":"The time member hours ends on Thursdays","t":"`$ANY`","key$":"thursday_member_close","index$":29},"thursday_member_open":{"a":true,"h":"Thursday Member Open","n":"thursday_member_open","r":false,"sh":"The time member hours starts on Thursdays","t":"`$ANY`","key$":"thursday_member_open","index$":30},"thursday_public_close":{"a":true,"h":"Thursday Public Close","n":"thursday_public_close","r":false,"sh":"The time public hours ends on Thursdays","t":"`$ANY`","key$":"thursday_public_close","index$":31},"thursday_public_open":{"a":true,"h":"Thursday Public Open","n":"thursday_public_open","r":false,"sh":"The time public hours starts on Thursdays","t":"`$ANY`","key$":"thursday_public_open","index$":32},"timestamp":{"a":true,"h":"Timestamp","n":"timestamp","r":false,"sh":"Date and time the record was updated in the aggregator search index","t":"`$ANY`","key$":"timestamp","index$":33},"title":{"a":true,"h":"Title","n":"title","r":false,"sh":"The name of this resource","t":"`$STRING`","key$":"title","index$":34},"tuesday_is_closed":{"a":true,"h":"Tuesday Is Closed","n":"tuesday_is_closed","r":false,"sh":"Whether the museum is closed on Tuesdays","t":"`$ANY`","key$":"tuesday_is_closed","index$":35},"tuesday_member_close":{"a":true,"h":"Tuesday Member Close","n":"tuesday_member_close","r":false,"sh":"The time member hours ends on Tuesdays","t":"`$ANY`","key$":"tuesday_member_close","index$":36},"tuesday_member_open":{"a":true,"h":"Tuesday Member Open","n":"tuesday_member_open","r":false,"sh":"The time member hours starts on Tuesdays","t":"`$ANY`","key$":"tuesday_member_open","index$":37},"tuesday_public_close":{"a":true,"h":"Tuesday Public Close","n":"tuesday_public_close","r":false,"sh":"The time public hours ends on Tuesdays","t":"`$ANY`","key$":"tuesday_public_close","index$":38},"tuesday_public_open":{"a":true,"h":"Tuesday Public Open","n":"tuesday_public_open","r":false,"sh":"The time public hours starts on Tuesdays","t":"`$ANY`","key$":"tuesday_public_open","index$":39},"updated_at":{"a":true,"h":"Updated At","n":"updated_at","r":false,"sh":"Date and time the record was updated in the aggregator database","t":"`$ANY`","key$":"updated_at","index$":40},"wednesday_is_closed":{"a":true,"h":"Wednesday Is Closed","n":"wednesday_is_closed","r":false,"sh":"Whether the museum is closed on Wednesdays","t":"`$ANY`","key$":"wednesday_is_closed","index$":41},"wednesday_member_close":{"a":true,"h":"Wednesday Member Close","n":"wednesday_member_close","r":false,"sh":"The time member hours ends on Wednesdays","t":"`$ANY`","key$":"wednesday_member_close","index$":42},"wednesday_member_open":{"a":true,"h":"Wednesday Member Open","n":"wednesday_member_open","r":false,"sh":"The time member hours starts on Wednesdays","t":"`$ANY`","key$":"wednesday_member_open","index$":43},"wednesday_public_close":{"a":true,"h":"Wednesday Public Close","n":"wednesday_public_close","r":false,"sh":"The time public hours ends on Wednesdays","t":"`$ANY`","key$":"wednesday_public_close","index$":44},"wednesday_public_open":{"a":true,"h":"Wednesday Public Open","n":"wednesday_public_open","r":false,"sh":"The time public hours starts on Wednesdays","t":"`$ANY`","key$":"wednesday_public_open","index$":45}},"id":{"field":"id","name":"id"},"name":"hour","op":{"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /hours","source":"openapi3","version":2},"g":{},"k":"http","m":"GET","o":"/hours","q":{},"r":{},"s":[{"lit":"hours"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /hours/{id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"GET","o":"/hours/{id}","q":{"exist":["id"]},"r":{},"s":[{"lit":"hours"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"hour","name__orig":"hour","Name":"Hour","name_":"hour","name-":"hour","NAME":"HOUR","index$":19}, {"active":true,"entity":"hour","key$":"BasicHourFlow","kind":"basic","name":"BasicHourFlow","param":{},"step":[{"a":true,"d":{},"i":{},"m":{},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"hour_ref01"}}],"index$":0},{"a":true,"d":{},"i":{"ref":"hour_ref01","srcdatavar":"hour_ref01_data","suffix":"_dt0"},"m":{"id":"hour01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-hour_ref01"}}],"index$":1}]}, 'Hour', {"GET /hours":{"protocol":"http","parameters":[]},"GET /hours/{id}":{"protocol":"http","parameters":[{"name":"id","in":"path","type":"string","required":true,"x-ref":"#/parameters/id","index$":0}]}}, { strict: LIVE_STRICT, t })
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let hour_ref01_data = Object.values(setup.data.existing.hour)[0] as any

    // LIST
    const hour_ref01_ent = client.Hour()
    const hour_ref01_match: any = {}

    const hour_ref01_list = (await hour_ref01_ent.list(hour_ref01_match)).map((e: any) => e.data())


    // LOAD
    const hour_ref01_match_dt0: any = {}
    hour_ref01_match_dt0.id = hour_ref01_data.id
    const hour_ref01_data_dt0 = (await hour_ref01_ent.load(hour_ref01_match_dt0)).data()
    assert(hour_ref01_data_dt0.id === hour_ref01_data.id)


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
      '../../../../.sdk/test/entity/hour/HourTestData.json')

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
    ['hour01','hour02','hour03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'ART_INSTITUTE_OF_CHICAGO_TEST_HOUR_ENTID': idmap,
    'ART_INSTITUTE_OF_CHICAGO_TEST_LIVE': 'FALSE',
    'ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['ART_INSTITUTE_OF_CHICAGO_TEST_HOUR_ENTID']

  const live = 'TRUE' === env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['ART_INSTITUTE_OF_CHICAGO_TEST_HOUR_ENTID']
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
  
