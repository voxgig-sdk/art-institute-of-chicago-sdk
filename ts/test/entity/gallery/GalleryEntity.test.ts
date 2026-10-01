

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


describe('GalleryEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE.
  afterEach(liveDelay('ART_INSTITUTE_OF_CHICAGO_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = ArtInstituteOfChicagoSDK.test()
    const ent = testsdk.Gallery()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE
    for (const op of ['list', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'gallery.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"api_link":{"a":true,"h":"Api Link","n":"api_link","r":false,"sh":"REST API link for this resource","t":"`$ANY`","key$":"api_link","index$":0},"api_model":{"a":true,"h":"Api Model","n":"api_model","r":false,"sh":"REST API resource type or endpoint","t":"`$ANY`","key$":"api_model","index$":1},"floor":{"a":true,"h":"Floor","n":"floor","r":false,"sh":"The level the gallery is on, e.g., 1, 2, 3, or LL","t":"`$ANY`","key$":"floor","index$":2},"id":{"a":true,"h":"Id","n":"id","r":false,"sh":"Unique identifier of this resource.","t":"`$STRING`","key$":"id","index$":3},"is_closed":{"a":true,"h":"Is Closed","n":"is_closed","r":false,"sh":"Whether the gallery is currently closed","t":"`$BOOLEAN`","key$":"is_closed","index$":4},"latitude":{"a":true,"h":"Latitude","n":"latitude","r":false,"sh":"Latitude coordinate of the center of the room","t":"`$NUMBER`","key$":"latitude","index$":5},"latlon":{"a":true,"h":"Latlon","n":"latlon","r":false,"sh":"Latitude and longitude coordinates of the center of the room","t":"`$ANY`","key$":"latlon","index$":6},"longitude":{"a":true,"h":"Longitude","n":"longitude","r":false,"sh":"Longitude coordinate of the center of the room","t":"`$NUMBER`","key$":"longitude","index$":7},"number":{"a":true,"h":"Number","n":"number","r":false,"sh":"The gallery's room number.","t":"`$ANY`","key$":"number","index$":8},"source_updated_at":{"a":true,"h":"Source Updated At","n":"source_updated_at","r":false,"sh":"Date and time the resource was updated in the source system","t":"`$ANY`","key$":"source_updated_at","index$":9},"suggest_autocomplete_all":{"a":true,"h":"Suggest Autocomplete All","n":"suggest_autocomplete_all","r":false,"sh":"Internal field to power the `/autosuggest` endpoint.","t":"`$ANY`","key$":"suggest_autocomplete_all","index$":10},"suggest_autocomplete_boosted":{"a":true,"h":"Suggest Autocomplete Boosted","n":"suggest_autocomplete_boosted","r":false,"sh":"Internal field to power the `/autocomplete` endpoint.","t":"`$ANY`","key$":"suggest_autocomplete_boosted","index$":11},"tgn_id":{"a":true,"h":"Tgn Id","n":"tgn_id","r":false,"sh":"Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN)","t":"`$STRING`","key$":"tgn_id","index$":12},"timestamp":{"a":true,"h":"Timestamp","n":"timestamp","r":false,"sh":"Date and time the record was updated in the aggregator search index","t":"`$ANY`","key$":"timestamp","index$":13},"title":{"a":true,"h":"Title","n":"title","r":false,"sh":"The name of this resource","t":"`$STRING`","key$":"title","index$":14},"updated_at":{"a":true,"h":"Updated At","n":"updated_at","r":false,"sh":"Date and time the record was updated in the aggregator database","t":"`$ANY`","key$":"updated_at","index$":15}},"id":{"field":"id","name":"id"},"name":"gallery","op":{"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /galleries","source":"openapi3","version":2},"g":{},"k":"http","m":"GET","o":"/galleries","q":{},"r":{},"s":[{"lit":"galleries"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /galleries/{id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"GET","o":"/galleries/{id}","q":{"exist":["id"]},"r":{},"s":[{"lit":"galleries"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"gallery","name__orig":"gallery","Name":"Gallery","name_":"gallery","name-":"gallery","NAME":"GALLERY","index$":16}, {"active":true,"entity":"gallery","key$":"BasicGalleryFlow","kind":"basic","name":"BasicGalleryFlow","param":{},"step":[{"a":true,"d":{},"i":{},"m":{},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"gallery_ref01"}}],"index$":0},{"a":true,"d":{},"i":{"ref":"gallery_ref01","srcdatavar":"gallery_ref01_data","suffix":"_dt0"},"m":{"id":"gallery01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-gallery_ref01"}}],"index$":1}]}, 'Gallery', {"GET /galleries":{"protocol":"http","parameters":[]},"GET /galleries/{id}":{"protocol":"http","parameters":[{"name":"id","in":"path","type":"string","required":true,"x-ref":"#/parameters/id","index$":0}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let gallery_ref01_data = Object.values(setup.data.existing.gallery)[0] as any

    // LIST
    const gallery_ref01_ent = client.Gallery()
    const gallery_ref01_match: any = {}

    const gallery_ref01_list = (await gallery_ref01_ent.list(gallery_ref01_match)).map((e: any) => e.data())


    // LOAD
    const gallery_ref01_match_dt0: any = {}
    gallery_ref01_match_dt0.id = gallery_ref01_data.id
    const gallery_ref01_data_dt0 = (await gallery_ref01_ent.load(gallery_ref01_match_dt0)).data()
    assert(gallery_ref01_data_dt0.id === gallery_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/gallery/GalleryTestData.json')

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
    ['gallery01','gallery02','gallery03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'ART_INSTITUTE_OF_CHICAGO_TEST_GALLERY_ENTID': idmap,
    'ART_INSTITUTE_OF_CHICAGO_TEST_LIVE': 'FALSE',
    'ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['ART_INSTITUTE_OF_CHICAGO_TEST_GALLERY_ENTID']

  const live = 'TRUE' === env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['ART_INSTITUTE_OF_CHICAGO_TEST_GALLERY_ENTID']
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
  
