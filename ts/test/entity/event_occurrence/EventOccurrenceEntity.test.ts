

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


describe('EventOccurrenceEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE.
  afterEach(liveDelay('ART_INSTITUTE_OF_CHICAGO_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = ArtInstituteOfChicagoSDK.test()
    const ent = testsdk.EventOccurrence()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE
    for (const op of ['list', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'event_occurrence.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"api_link":{"a":true,"h":"Api Link","n":"api_link","r":false,"sh":"REST API link for this resource","t":"`$ANY`","key$":"api_link","index$":0},"api_model":{"a":true,"h":"Api Model","n":"api_model","r":false,"sh":"REST API resource type or endpoint","t":"`$ANY`","key$":"api_model","index$":1},"button_caption":{"a":true,"h":"Button Caption","n":"button_caption","r":false,"sh":"Additional text below the ticket/registration button","t":"`$ANY`","key$":"button_caption","index$":2},"button_text":{"a":true,"h":"Button Text","n":"button_text","r":false,"sh":"The text used on the ticket/registration button","t":"`$ANY`","key$":"button_text","index$":3},"button_url":{"a":true,"h":"Button Url","n":"button_url","r":false,"sh":"The URL to the sales site or an RSVP link for this event","t":"`$ANY`","key$":"button_url","index$":4},"description":{"a":true,"h":"Description","n":"description","r":false,"sh":"Description of the event","t":"`$STRING`","key$":"description","index$":5},"end_at":{"a":true,"h":"End At","n":"end_at","r":false,"sh":"The date the event occurrence ends","t":"`$ANY`","key$":"end_at","index$":6},"event_id":{"a":true,"h":"Event Id","n":"event_id","r":false,"sh":"Identifier of the master event of which this is an occurrence","t":"`$STRING`","key$":"event_id","index$":7},"id":{"a":true,"h":"Id","n":"id","r":false,"sh":"Unique identifier of this resource.","t":"`$STRING`","key$":"id","index$":8},"image_url":{"a":true,"h":"Image Url","n":"image_url","r":false,"sh":"The URL of an image representing this page","t":"`$ANY`","key$":"image_url","index$":9},"is_private":{"a":true,"h":"Is Private","n":"is_private","r":false,"sh":"Whether the event is private.","t":"`$BOOLEAN`","key$":"is_private","index$":10},"is_sales_button_hidden":{"a":true,"h":"Is Sales Button Hidden","n":"is_sales_button_hidden","r":false,"sh":"Whether the buy tickets button is hidden on the website event page","t":"`$BOOLEAN`","key$":"is_sales_button_hidden","index$":11},"is_ticketed":{"a":true,"h":"Is Ticketed","n":"is_ticketed","r":false,"sh":"Whether a ticket is required to attend the event","t":"`$BOOLEAN`","key$":"is_ticketed","index$":12},"location":{"a":true,"h":"Location","n":"location","r":false,"sh":"Where the event takes place","t":"`$ANY`","key$":"location","index$":13},"off_sale_at":{"a":true,"h":"Off Sale At","n":"off_sale_at","r":false,"sh":"Date and time the event goes off sale","t":"`$ANY`","key$":"off_sale_at","index$":14},"on_sale_at":{"a":true,"h":"On Sale At","n":"on_sale_at","r":false,"sh":"Date and time the event goes on sale","t":"`$ANY`","key$":"on_sale_at","index$":15},"short_description":{"a":true,"h":"Short Description","n":"short_description","r":false,"sh":"Brief description of the event","t":"`$ANY`","key$":"short_description","index$":16},"source_updated_at":{"a":true,"h":"Source Updated At","n":"source_updated_at","r":false,"sh":"Date and time the resource was updated in the source system","t":"`$ANY`","key$":"source_updated_at","index$":17},"start_at":{"a":true,"h":"Start At","n":"start_at","r":false,"sh":"The date the event occurrence begins","t":"`$ANY`","key$":"start_at","index$":18},"suggest_autocomplete_all":{"a":true,"h":"Suggest Autocomplete All","n":"suggest_autocomplete_all","r":false,"sh":"Internal field to power the `/autosuggest` endpoint.","t":"`$ANY`","key$":"suggest_autocomplete_all","index$":19},"suggest_autocomplete_boosted":{"a":true,"h":"Suggest Autocomplete Boosted","n":"suggest_autocomplete_boosted","r":false,"sh":"Internal field to power the `/autocomplete` endpoint.","t":"`$ANY`","key$":"suggest_autocomplete_boosted","index$":20},"timestamp":{"a":true,"h":"Timestamp","n":"timestamp","r":false,"sh":"Date and time the record was updated in the aggregator search index","t":"`$ANY`","key$":"timestamp","index$":21},"title":{"a":true,"h":"Title","n":"title","r":false,"sh":"The name of this resource","t":"`$STRING`","key$":"title","index$":22},"title_display":{"a":true,"h":"Title Display","n":"title_display","r":false,"sh":"The name of this event formatted with HTML (optional)","t":"`$ANY`","key$":"title_display","index$":23},"updated_at":{"a":true,"h":"Updated At","n":"updated_at","r":false,"sh":"Date and time the record was updated in the aggregator database","t":"`$ANY`","key$":"updated_at","index$":24}},"id":{"field":"id","name":"id"},"name":"event_occurrence","op":{"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /event-occurrences","source":"openapi3","version":2},"g":{},"k":"http","m":"GET","o":"/event-occurrences","q":{},"r":{},"s":[{"lit":"event-occurrences"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /event-occurrences/{id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"GET","o":"/event-occurrences/{id}","q":{"exist":["id"]},"r":{},"s":[{"lit":"event-occurrences"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"event_occurrence","name__orig":"event_occurrence","Name":"EventOccurrence","name_":"event_occurrence","name-":"event-occurrence","NAME":"EVENT_OCCURRENCE","index$":13}, {"active":true,"entity":"event_occurrence","key$":"BasicEventOccurrenceFlow","kind":"basic","name":"BasicEventOccurrenceFlow","param":{},"step":[{"a":true,"d":{},"i":{},"m":{},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"event_occurrence_ref01"}}],"index$":0},{"a":true,"d":{},"i":{"ref":"event_occurrence_ref01","srcdatavar":"event_occurrence_ref01_data","suffix":"_dt0"},"m":{"id":"event_occurrence01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-event_occurrence_ref01"}}],"index$":1}]}, 'EventOccurrence', {"GET /event-occurrences":{"protocol":"http","responses":{"200":{"description":"Successful operation","schema":{"type":"array","items":{"properties":{"id":{"description":"Unique identifier of this resource. Taken from the source system.","key$":"id"},"api_model":{"description":"REST API resource type or endpoint","key$":"api_model"},"api_link":{"description":"REST API link for this resource","key$":"api_link"},"title":{"description":"The name of this resource","key$":"title"},"title_display":{"description":"The name of this event formatted with HTML (optional)","key$":"title_display"},"event_id":{"description":"Identifier of the master event of which this is an occurrence","key$":"event_id"},"short_description":{"description":"Brief description of the event","key$":"short_description"},"description":{"description":"Description of the event","key$":"description"},"image_url":{"description":"The URL of an image representing this page","key$":"image_url"},"is_ticketed":{"description":"Whether a ticket is required to attend the event","key$":"is_ticketed"},"is_private":{"description":"Whether the event is private. Private events should be omitted from listings.","key$":"is_private"},"is_sales_button_hidden":{"description":"Whether the buy tickets button is hidden on the website event page","key$":"is_sales_button_hidden"},"start_at":{"description":"The date the event occurrence begins","key$":"start_at"},"end_at":{"description":"The date the event occurrence ends","key$":"end_at"},"on_sale_at":{"description":"Date and time the event goes on sale","key$":"on_sale_at"},"off_sale_at":{"description":"Date and time the event goes off sale","key$":"off_sale_at"},"location":{"description":"Where the event takes place","key$":"location"},"button_url":{"description":"The URL to the sales site or an RSVP link for this event","key$":"button_url"},"button_text":{"description":"The text used on the ticket/registration button","key$":"button_text"},"button_caption":{"description":"Additional text below the ticket/registration button","key$":"button_caption"},"suggest_autocomplete_boosted":{"description":"Internal field to power the `/autocomplete` endpoint. Do not use directly.","key$":"suggest_autocomplete_boosted"},"suggest_autocomplete_all":{"description":"Internal field to power the `/autosuggest` endpoint. Do not use directly.","key$":"suggest_autocomplete_all"},"source_updated_at":{"description":"Date and time the resource was updated in the source system","key$":"source_updated_at"},"updated_at":{"description":"Date and time the record was updated in the aggregator database","key$":"updated_at"},"timestamp":{"description":"Date and time the record was updated in the aggregator search index","key$":"timestamp"}},"type":"object","x-ref":"#/definitions/EventOccurrence","index$":0}}},"default":{"description":"error","schema":{"required":["status","error","detail"],"properties":{"status":{"type":"integer"},"error":{"type":"string"},"detail":{"type":"string"}},"x-ref":"#/definitions/Error"}}},"produces":["application/json"],"parameters":[],"securitySource":"unspecified"},"GET /event-occurrences/{id}":{"protocol":"http","responses":{"200":{"description":"Successful operation","schema":{"type":"array","items":{"properties":{"id":{"description":"Unique identifier of this resource. Taken from the source system.","key$":"id"},"api_model":{"description":"REST API resource type or endpoint","key$":"api_model"},"api_link":{"description":"REST API link for this resource","key$":"api_link"},"title":{"description":"The name of this resource","key$":"title"},"title_display":{"description":"The name of this event formatted with HTML (optional)","key$":"title_display"},"event_id":{"description":"Identifier of the master event of which this is an occurrence","key$":"event_id"},"short_description":{"description":"Brief description of the event","key$":"short_description"},"description":{"description":"Description of the event","key$":"description"},"image_url":{"description":"The URL of an image representing this page","key$":"image_url"},"is_ticketed":{"description":"Whether a ticket is required to attend the event","key$":"is_ticketed"},"is_private":{"description":"Whether the event is private. Private events should be omitted from listings.","key$":"is_private"},"is_sales_button_hidden":{"description":"Whether the buy tickets button is hidden on the website event page","key$":"is_sales_button_hidden"},"start_at":{"description":"The date the event occurrence begins","key$":"start_at"},"end_at":{"description":"The date the event occurrence ends","key$":"end_at"},"on_sale_at":{"description":"Date and time the event goes on sale","key$":"on_sale_at"},"off_sale_at":{"description":"Date and time the event goes off sale","key$":"off_sale_at"},"location":{"description":"Where the event takes place","key$":"location"},"button_url":{"description":"The URL to the sales site or an RSVP link for this event","key$":"button_url"},"button_text":{"description":"The text used on the ticket/registration button","key$":"button_text"},"button_caption":{"description":"Additional text below the ticket/registration button","key$":"button_caption"},"suggest_autocomplete_boosted":{"description":"Internal field to power the `/autocomplete` endpoint. Do not use directly.","key$":"suggest_autocomplete_boosted"},"suggest_autocomplete_all":{"description":"Internal field to power the `/autosuggest` endpoint. Do not use directly.","key$":"suggest_autocomplete_all"},"source_updated_at":{"description":"Date and time the resource was updated in the source system","key$":"source_updated_at"},"updated_at":{"description":"Date and time the record was updated in the aggregator database","key$":"updated_at"},"timestamp":{"description":"Date and time the record was updated in the aggregator search index","key$":"timestamp"}},"type":"object","x-ref":"#/definitions/EventOccurrence","key$":"items"}}},"default":{"description":"error","schema":{"required":["status","error","detail"],"properties":{"status":{"type":"integer"},"error":{"type":"string"},"detail":{"type":"string"}},"x-ref":"#/definitions/Error"}}},"produces":["application/json"],"parameters":[{"name":"id","in":"path","type":"string","required":true,"x-ref":"#/parameters/id","index$":0}],"securitySource":"unspecified"}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let event_occurrence_ref01_data = Object.values(setup.data.existing.event_occurrence)[0] as any

    // LIST
    const event_occurrence_ref01_ent = client.EventOccurrence()
    const event_occurrence_ref01_match: any = {}

    const event_occurrence_ref01_list = (await event_occurrence_ref01_ent.list(event_occurrence_ref01_match)).map((e: any) => e.data())


    // LOAD
    const event_occurrence_ref01_match_dt0: any = {}
    event_occurrence_ref01_match_dt0.id = event_occurrence_ref01_data.id
    const event_occurrence_ref01_data_dt0 = (await event_occurrence_ref01_ent.load(event_occurrence_ref01_match_dt0)).data()
    assert(event_occurrence_ref01_data_dt0.id === event_occurrence_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/event_occurrence/EventOccurrenceTestData.json')

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
    ['event_occurrence01','event_occurrence02','event_occurrence03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_OCCURRENCE_ENTID': idmap,
    'ART_INSTITUTE_OF_CHICAGO_TEST_LIVE': 'FALSE',
    'ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_OCCURRENCE_ENTID']

  const live = 'TRUE' === env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_OCCURRENCE_ENTID']
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
  
