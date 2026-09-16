

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


describe('EventEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE.
  afterEach(liveDelay('ART_INSTITUTE_OF_CHICAGO_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = ArtInstituteOfChicagoSDK.test()
    const ent = testsdk.Event()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE
    for (const op of ['list', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'event.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"alt_audience_ids","req":false,"short":"Unique identifiers indicating the alternate audiences for this event","type":"`$ANY`","index$":0},{"active":true,"name":"alt_event_type_ids","req":false,"short":"Unique identifiers indicating the alternate types of this event","type":"`$ANY`","index$":1},{"active":true,"name":"api_link","req":false,"short":"REST API link for this resource","type":"`$ANY`","index$":2},{"active":true,"name":"api_model","req":false,"short":"REST API resource type or endpoint","type":"`$ANY`","index$":3},{"active":true,"name":"audience_id","req":false,"short":"Unique identifier indicating the preferred audience for this event","type":"`$STRING`","index$":4},{"active":true,"name":"buy_button_caption","req":false,"short":"Additional text below the ticket/registration button","type":"`$ANY`","index$":5},{"active":true,"name":"buy_button_text","req":false,"short":"The text used on the ticket/registration button","type":"`$ANY`","index$":6},{"active":true,"name":"date_display","req":false,"short":"A readable display of the event dates","type":"`$ANY`","index$":7},{"active":true,"name":"description","req":false,"short":"All copytext of the event","type":"`$STRING`","index$":8},{"active":true,"name":"door_time","req":false,"short":"The time the doors open for this event","type":"`$ANY`","index$":9},{"active":true,"name":"end_date","req":false,"short":"The date the event ends","type":"`$ANY`","index$":10},{"active":true,"name":"end_time","req":false,"short":"The time the event ends","type":"`$ANY`","index$":11},{"active":true,"name":"entrance","req":false,"short":"Which entrance to use for this event","type":"`$ANY`","index$":12},{"active":true,"name":"event_host_id","req":false,"short":"Unique identifier of the host (cf.","type":"`$STRING`","index$":13},{"active":true,"name":"event_host_title","req":false,"short":"Unique identifier of the host (cf.","type":"`$ANY`","index$":14},{"active":true,"name":"event_type_id","req":false,"short":"Unique identifier indicating the preferred type of this event","type":"`$STRING`","index$":15},{"active":true,"name":"header_description","req":false,"short":"Brief description of the event displayed below the title","type":"`$ANY`","index$":16},{"active":true,"name":"hero_caption","req":false,"short":"Text displayed with the hero image on the event","type":"`$ANY`","index$":17},{"active":true,"name":"id","req":false,"short":"Unique identifier of this resource.","type":"`$STRING`","index$":18},{"active":true,"name":"image_url","req":false,"short":"The URL of an image representing this page","type":"`$ANY`","index$":19},{"active":true,"name":"is_admission_required","req":false,"short":"Whether admission to the museum is required to attend this event","type":"`$BOOLEAN`","index$":20},{"active":true,"name":"is_after_hours","req":false,"short":"Whether the event is to be held after the museum closes","type":"`$BOOLEAN`","index$":21},{"active":true,"name":"is_free","req":false,"short":"Whether the event is free","type":"`$BOOLEAN`","index$":22},{"active":true,"name":"is_member_exclusive","req":false,"short":"Whether the event is exclusive to members of the museum","type":"`$BOOLEAN`","index$":23},{"active":true,"name":"is_private","req":false,"short":"Whether the event is private","type":"`$BOOLEAN`","index$":24},{"active":true,"name":"is_registration_required","req":false,"short":"Whether registration is required to attend the event","type":"`$BOOLEAN`","index$":25},{"active":true,"name":"is_sales_button_hidden","req":false,"short":"Whether the buy tickets button is hidden on the website event page","type":"`$BOOLEAN`","index$":26},{"active":true,"name":"is_sold_out","req":false,"short":"Whether the event is sold out","type":"`$BOOLEAN`","index$":27},{"active":true,"name":"is_ticketed","req":false,"short":"Whether a ticket is required to attend the event","type":"`$BOOLEAN`","index$":28},{"active":true,"name":"is_virtual_event","req":false,"short":"Whether the event is being held virtually","type":"`$BOOLEAN`","index$":29},{"active":true,"name":"join_url","req":false,"short":"URL to the membership signup page via this event","type":"`$ANY`","index$":30},{"active":true,"name":"layout_type","req":false,"short":"Number indicating the type of layout this event page uses","type":"`$ANY`","index$":31},{"active":true,"name":"list_description","req":false,"short":"One-sentence description of the event displayed in listings","type":"`$ANY`","index$":32},{"active":true,"name":"location","req":false,"short":"Where the event takes place","type":"`$ANY`","index$":33},{"active":true,"name":"program_ids","req":false,"short":"Unique identifiers indicating the programs this event is a part of","type":"`$ANY`","index$":34},{"active":true,"name":"program_titles","req":false,"short":"Titles of the programs this event is a part of","type":"`$ANY`","index$":35},{"active":true,"name":"rsvp_link","req":false,"short":"The URL to the sales site for this event","type":"`$ANY`","index$":36},{"active":true,"name":"search_tags","req":false,"short":"Editor-specified list of tags to aid in internal search","type":"`$ANY`","index$":37},{"active":true,"name":"short_description","req":false,"short":"Brief description of the event","type":"`$ANY`","index$":38},{"active":true,"name":"slug","req":false,"short":"A string used in the URL for this event","type":"`$STRING`","index$":39},{"active":true,"name":"source_updated_at","req":false,"short":"Date and time the resource was updated in the source system","type":"`$ANY`","index$":40},{"active":true,"name":"start_date","req":false,"short":"The date the event begins","type":"`$ANY`","index$":41},{"active":true,"name":"start_time","req":false,"short":"The time the event starts","type":"`$ANY`","index$":42},{"active":true,"name":"suggest_autocomplete_all","req":false,"short":"Internal field to power the `/autosuggest` endpoint.","type":"`$ANY`","index$":43},{"active":true,"name":"suggest_autocomplete_boosted","req":false,"short":"Internal field to power the `/autocomplete` endpoint.","type":"`$ANY`","index$":44},{"active":true,"name":"survey_url","req":false,"short":"URL to the survey associated with this event","type":"`$ANY`","index$":45},{"active":true,"name":"ticketed_event_id","req":false,"short":"Unique identifier of the event in the ticketing system this website event is tied to","type":"`$STRING`","index$":46},{"active":true,"name":"timestamp","req":false,"short":"Date and time the record was updated in the aggregator search index","type":"`$ANY`","index$":47},{"active":true,"name":"title","req":false,"short":"The name of this resource","type":"`$STRING`","index$":48},{"active":true,"name":"title_display","req":false,"short":"The name of this event formatted with HTML (optional)","type":"`$ANY`","index$":49},{"active":true,"name":"updated_at","req":false,"short":"Date and time the record was updated in the aggregator database","type":"`$ANY`","index$":50},{"active":true,"name":"virtual_event_passcode","req":false,"short":"Passcode to access the virtual event","type":"`$ANY`","index$":51},{"active":true,"name":"virtual_event_url","req":false,"short":"URL to the virtual event","type":"`$ANY`","index$":52}],"id":{"field":"id","name":"id"},"name":"event","op":{"list":{"input":"data","name":"list","points":[{"active":true,"args":{},"contract":{"id":"GET /events","json":"{\"parameters\":[],\"produces\":[\"application/json\"],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"Successful operation\",\"schema\":{\"items\":{\"properties\":{\"alt_audience_ids\":{\"description\":\"Unique identifiers indicating the alternate audiences for this event\"},\"alt_event_type_ids\":{\"description\":\"Unique identifiers indicating the alternate types of this event\"},\"api_link\":{\"description\":\"REST API link for this resource\"},\"api_model\":{\"description\":\"REST API resource type or endpoint\"},\"audience_id\":{\"description\":\"Unique identifier indicating the preferred audience for this event\"},\"buy_button_caption\":{\"description\":\"Additional text below the ticket/registration button\"},\"buy_button_text\":{\"description\":\"The text used on the ticket/registration button\"},\"date_display\":{\"description\":\"A readable display of the event dates\"},\"description\":{\"description\":\"All copytext of the event\"},\"door_time\":{\"description\":\"The time the doors open for this event\"},\"end_date\":{\"description\":\"The date the event ends\"},\"end_time\":{\"description\":\"The time the event ends\"},\"entrance\":{\"description\":\"Which entrance to use for this event\"},\"event_host_id\":{\"description\":\"Unique identifier of the host (cf. event programs) that is presenting this event\"},\"event_host_title\":{\"description\":\"Unique identifier of the host (cf. event programs) that is presenting this event\"},\"event_type_id\":{\"description\":\"Unique identifier indicating the preferred type of this event\"},\"header_description\":{\"description\":\"Brief description of the event displayed below the title\"},\"hero_caption\":{\"description\":\"Text displayed with the hero image on the event\"},\"id\":{\"description\":\"Unique identifier of this resource. Taken from the source system.\"},\"image_url\":{\"description\":\"The URL of an image representing this page\"},\"is_admission_required\":{\"description\":\"Whether admission to the museum is required to attend this event\"},\"is_after_hours\":{\"description\":\"Whether the event is to be held after the museum closes\"},\"is_free\":{\"description\":\"Whether the event is free\"},\"is_member_exclusive\":{\"description\":\"Whether the event is exclusive to members of the museum\"},\"is_private\":{\"description\":\"Whether the event is private\"},\"is_registration_required\":{\"description\":\"Whether registration is required to attend the event\"},\"is_sales_button_hidden\":{\"description\":\"Whether the buy tickets button is hidden on the website event page\"},\"is_sold_out\":{\"description\":\"Whether the event is sold out\"},\"is_ticketed\":{\"description\":\"Whether a ticket is required to attend the event\"},\"is_virtual_event\":{\"description\":\"Whether the event is being held virtually\"},\"join_url\":{\"description\":\"URL to the membership signup page via this event\"},\"layout_type\":{\"description\":\"Number indicating the type of layout this event page uses\"},\"list_description\":{\"description\":\"One-sentence description of the event displayed in listings\"},\"location\":{\"description\":\"Where the event takes place\"},\"program_ids\":{\"description\":\"Unique identifiers indicating the programs this event is a part of\"},\"program_titles\":{\"description\":\"Titles of the programs this event is a part of\"},\"rsvp_link\":{\"description\":\"The URL to the sales site for this event\"},\"search_tags\":{\"description\":\"Editor-specified list of tags to aid in internal search\"},\"short_description\":{\"description\":\"Brief description of the event\"},\"slug\":{\"description\":\"A string used in the URL for this event\"},\"source_updated_at\":{\"description\":\"Date and time the resource was updated in the source system\"},\"start_date\":{\"description\":\"The date the event begins\"},\"start_time\":{\"description\":\"The time the event starts\"},\"suggest_autocomplete_all\":{\"description\":\"Internal field to power the `/autosuggest` endpoint. Do not use directly.\"},\"suggest_autocomplete_boosted\":{\"description\":\"Internal field to power the `/autocomplete` endpoint. Do not use directly.\"},\"survey_url\":{\"description\":\"URL to the survey associated with this event\"},\"ticketed_event_id\":{\"description\":\"Unique identifier of the event in the ticketing system this website event is tied to\"},\"timestamp\":{\"description\":\"Date and time the record was updated in the aggregator search index\"},\"title\":{\"description\":\"The name of this resource\"},\"title_display\":{\"description\":\"The name of this event formatted with HTML (optional)\"},\"updated_at\":{\"description\":\"Date and time the record was updated in the aggregator database\"},\"virtual_event_passcode\":{\"description\":\"Passcode to access the virtual event\"},\"virtual_event_url\":{\"description\":\"URL to the virtual event\"}},\"type\":\"object\"},\"type\":\"array\"}},\"default\":{\"description\":\"error\",\"schema\":{\"properties\":{\"detail\":{\"type\":\"string\"},\"error\":{\"type\":\"string\"},\"status\":{\"type\":\"integer\"}},\"required\":[\"status\",\"error\",\"detail\"]}}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/events","segments":[{"lit":"events"}],"select":{},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"id","orig":"id","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"GET /events/{id}","json":"{\"parameters\":[{\"in\":\"path\",\"name\":\"id\",\"required\":true,\"type\":\"string\"}],\"produces\":[\"application/json\"],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"Successful operation\",\"schema\":{\"items\":{\"properties\":{\"alt_audience_ids\":{\"description\":\"Unique identifiers indicating the alternate audiences for this event\"},\"alt_event_type_ids\":{\"description\":\"Unique identifiers indicating the alternate types of this event\"},\"api_link\":{\"description\":\"REST API link for this resource\"},\"api_model\":{\"description\":\"REST API resource type or endpoint\"},\"audience_id\":{\"description\":\"Unique identifier indicating the preferred audience for this event\"},\"buy_button_caption\":{\"description\":\"Additional text below the ticket/registration button\"},\"buy_button_text\":{\"description\":\"The text used on the ticket/registration button\"},\"date_display\":{\"description\":\"A readable display of the event dates\"},\"description\":{\"description\":\"All copytext of the event\"},\"door_time\":{\"description\":\"The time the doors open for this event\"},\"end_date\":{\"description\":\"The date the event ends\"},\"end_time\":{\"description\":\"The time the event ends\"},\"entrance\":{\"description\":\"Which entrance to use for this event\"},\"event_host_id\":{\"description\":\"Unique identifier of the host (cf. event programs) that is presenting this event\"},\"event_host_title\":{\"description\":\"Unique identifier of the host (cf. event programs) that is presenting this event\"},\"event_type_id\":{\"description\":\"Unique identifier indicating the preferred type of this event\"},\"header_description\":{\"description\":\"Brief description of the event displayed below the title\"},\"hero_caption\":{\"description\":\"Text displayed with the hero image on the event\"},\"id\":{\"description\":\"Unique identifier of this resource. Taken from the source system.\"},\"image_url\":{\"description\":\"The URL of an image representing this page\"},\"is_admission_required\":{\"description\":\"Whether admission to the museum is required to attend this event\"},\"is_after_hours\":{\"description\":\"Whether the event is to be held after the museum closes\"},\"is_free\":{\"description\":\"Whether the event is free\"},\"is_member_exclusive\":{\"description\":\"Whether the event is exclusive to members of the museum\"},\"is_private\":{\"description\":\"Whether the event is private\"},\"is_registration_required\":{\"description\":\"Whether registration is required to attend the event\"},\"is_sales_button_hidden\":{\"description\":\"Whether the buy tickets button is hidden on the website event page\"},\"is_sold_out\":{\"description\":\"Whether the event is sold out\"},\"is_ticketed\":{\"description\":\"Whether a ticket is required to attend the event\"},\"is_virtual_event\":{\"description\":\"Whether the event is being held virtually\"},\"join_url\":{\"description\":\"URL to the membership signup page via this event\"},\"layout_type\":{\"description\":\"Number indicating the type of layout this event page uses\"},\"list_description\":{\"description\":\"One-sentence description of the event displayed in listings\"},\"location\":{\"description\":\"Where the event takes place\"},\"program_ids\":{\"description\":\"Unique identifiers indicating the programs this event is a part of\"},\"program_titles\":{\"description\":\"Titles of the programs this event is a part of\"},\"rsvp_link\":{\"description\":\"The URL to the sales site for this event\"},\"search_tags\":{\"description\":\"Editor-specified list of tags to aid in internal search\"},\"short_description\":{\"description\":\"Brief description of the event\"},\"slug\":{\"description\":\"A string used in the URL for this event\"},\"source_updated_at\":{\"description\":\"Date and time the resource was updated in the source system\"},\"start_date\":{\"description\":\"The date the event begins\"},\"start_time\":{\"description\":\"The time the event starts\"},\"suggest_autocomplete_all\":{\"description\":\"Internal field to power the `/autosuggest` endpoint. Do not use directly.\"},\"suggest_autocomplete_boosted\":{\"description\":\"Internal field to power the `/autocomplete` endpoint. Do not use directly.\"},\"survey_url\":{\"description\":\"URL to the survey associated with this event\"},\"ticketed_event_id\":{\"description\":\"Unique identifier of the event in the ticketing system this website event is tied to\"},\"timestamp\":{\"description\":\"Date and time the record was updated in the aggregator search index\"},\"title\":{\"description\":\"The name of this resource\"},\"title_display\":{\"description\":\"The name of this event formatted with HTML (optional)\"},\"updated_at\":{\"description\":\"Date and time the record was updated in the aggregator database\"},\"virtual_event_passcode\":{\"description\":\"Passcode to access the virtual event\"},\"virtual_event_url\":{\"description\":\"URL to the virtual event\"}},\"type\":\"object\"},\"type\":\"array\"}},\"default\":{\"description\":\"error\",\"schema\":{\"properties\":{\"detail\":{\"type\":\"string\"},\"error\":{\"type\":\"string\"},\"status\":{\"type\":\"integer\"}},\"required\":[\"status\",\"error\",\"detail\"]}}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/events/{id}","segments":[{"lit":"events"},{"var":"id"}],"select":{"exist":["id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"event","name__orig":"event","Name":"Event","name_":"event","name-":"event","NAME":"EVENT","index$":12}, {"active":true,"entity":"event","key$":"BasicEventFlow","kind":"basic","name":"BasicEventFlow","param":{},"step":[{"active":true,"data":{},"input":{},"match":{},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"event_ref01"}}],"index$":0},{"active":true,"data":{},"input":{"ref":"event_ref01","srcdatavar":"event_ref01_data","suffix":"_dt0"},"match":{"id":"event01"},"op":"load","spec":[],"valid":[{"apply":"TextFieldMark","def":{"mark":"Mark01-event_ref01"}}],"index$":1}]}, 'Event')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let event_ref01_data = Object.values(setup.data.existing.event)[0] as any

    // LIST
    const event_ref01_ent = client.Event()
    const event_ref01_match: any = {}

    const event_ref01_list = (await event_ref01_ent.list(event_ref01_match)).map((e: any) => e.data())


    // LOAD
    const event_ref01_match_dt0: any = {}
    event_ref01_match_dt0.id = event_ref01_data.id
    const event_ref01_data_dt0 = (await event_ref01_ent.load(event_ref01_match_dt0)).data()
    assert(event_ref01_data_dt0.id === event_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/event/EventTestData.json')

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
    ['event01','event02','event03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_ENTID': idmap,
    'ART_INSTITUTE_OF_CHICAGO_TEST_LIVE': 'FALSE',
    'ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_ENTID']

  const live = 'TRUE' === env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_ENTID']
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
  
