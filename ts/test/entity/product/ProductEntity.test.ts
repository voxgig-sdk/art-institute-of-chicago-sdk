

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


describe('ProductEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE.
  afterEach(liveDelay('ART_INSTITUTE_OF_CHICAGO_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = ArtInstituteOfChicagoSDK.test()
    const ent = testsdk.Product()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE
    for (const op of ['list', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'product.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"api_link":{"a":true,"h":"Api Link","n":"api_link","r":false,"sh":"REST API link for this resource","t":"`$ANY`","key$":"api_link","index$":0},"api_model":{"a":true,"h":"Api Model","n":"api_model","r":false,"sh":"REST API resource type or endpoint","t":"`$ANY`","key$":"api_model","index$":1},"artist_ids":{"a":true,"h":"Artist Ids","n":"artist_ids","r":false,"sh":"Unique identifiers of the artists associated with this product","t":"`$ANY`","key$":"artist_ids","index$":2},"artwork_ids":{"a":true,"h":"Artwork Ids","n":"artwork_ids","r":false,"sh":"Unique identifiers of the artworks associated with this product","t":"`$ANY`","key$":"artwork_ids","index$":3},"description":{"a":true,"h":"Description","n":"description","r":false,"sh":"Explanation of what this product is","t":"`$STRING`","key$":"description","index$":4},"exhibition_ids":{"a":true,"h":"Exhibition Ids","n":"exhibition_ids","r":false,"sh":"Unique identifiers of the exhibitions associated with this product","t":"`$ANY`","key$":"exhibition_ids","index$":5},"external_sku":{"a":true,"h":"External Sku","n":"external_sku","r":false,"sh":"Numeric product identification code of a machine-readable barcode, when the customer sku differs from our internal one","t":"`$ANY`","key$":"external_sku","index$":6},"id":{"a":true,"h":"Id","n":"id","r":false,"sh":"Unique identifier of this resource.","t":"`$STRING`","key$":"id","index$":7},"image_url":{"a":true,"h":"Image Url","n":"image_url","r":false,"sh":"URL of an image for this product","t":"`$ANY`","key$":"image_url","index$":8},"max_compare_at_price":{"a":true,"h":"Max Compare At Price","n":"max_compare_at_price","r":false,"sh":"Number indicating how much the most expensive variant of a product cost before a sale","t":"`$ANY`","key$":"max_compare_at_price","index$":9},"max_current_price":{"a":true,"h":"Max Current Price","n":"max_current_price","r":false,"sh":"Number indicating how much the most expensive variant of a product costs right now","t":"`$ANY`","key$":"max_current_price","index$":10},"min_compare_at_price":{"a":true,"h":"Min Compare At Price","n":"min_compare_at_price","r":false,"sh":"Number indicating how much the least expensive variant of a product cost before a sale","t":"`$ANY`","key$":"min_compare_at_price","index$":11},"min_current_price":{"a":true,"h":"Min Current Price","n":"min_current_price","r":false,"sh":"Number indicating how much the least expensive variant of a product costs right now","t":"`$ANY`","key$":"min_current_price","index$":12},"price_display":{"a":true,"h":"Price Display","n":"price_display","r":false,"sh":"Explanation of what this product is","t":"`$ANY`","key$":"price_display","index$":13},"source_updated_at":{"a":true,"h":"Source Updated At","n":"source_updated_at","r":false,"sh":"Date and time the resource was updated in the source system","t":"`$ANY`","key$":"source_updated_at","index$":14},"suggest_autocomplete_all":{"a":true,"h":"Suggest Autocomplete All","n":"suggest_autocomplete_all","r":false,"sh":"Internal field to power the `/autosuggest` endpoint.","t":"`$ANY`","key$":"suggest_autocomplete_all","index$":15},"suggest_autocomplete_boosted":{"a":true,"h":"Suggest Autocomplete Boosted","n":"suggest_autocomplete_boosted","r":false,"sh":"Internal field to power the `/autocomplete` endpoint.","t":"`$ANY`","key$":"suggest_autocomplete_boosted","index$":16},"timestamp":{"a":true,"h":"Timestamp","n":"timestamp","r":false,"sh":"Date and time the record was updated in the aggregator search index","t":"`$ANY`","key$":"timestamp","index$":17},"title":{"a":true,"h":"Title","n":"title","r":false,"sh":"The name of this resource","t":"`$STRING`","key$":"title","index$":18},"updated_at":{"a":true,"h":"Updated At","n":"updated_at","r":false,"sh":"Date and time the record was updated in the aggregator database","t":"`$ANY`","key$":"updated_at","index$":19},"web_url":{"a":true,"h":"Web Url","n":"web_url","r":false,"sh":"URL of this product in the shop","t":"`$ANY`","key$":"web_url","index$":20}},"id":{"field":"id","name":"id"},"name":"product","op":{"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /products","source":"openapi3","version":2},"g":{},"k":"http","m":"GET","o":"/products","q":{},"r":{},"s":[{"lit":"products"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /products/{id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"GET","o":"/products/{id}","q":{"exist":["id"]},"r":{},"s":[{"lit":"products"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"product","name__orig":"product","Name":"Product","name_":"product","name-":"product","NAME":"PRODUCT","index$":25}, {"active":true,"entity":"product","key$":"BasicProductFlow","kind":"basic","name":"BasicProductFlow","param":{},"step":[{"a":true,"d":{},"i":{},"m":{},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"product_ref01"}}],"index$":0},{"a":true,"d":{},"i":{"ref":"product_ref01","srcdatavar":"product_ref01_data","suffix":"_dt0"},"m":{"id":"product01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-product_ref01"}}],"index$":1}]}, 'Product', {"GET /products":{"protocol":"http","responses":{"200":{"description":"Successful operation","schema":{"type":"array","items":{"properties":{"id":{"description":"Unique identifier of this resource. Taken from the source system.","key$":"id"},"api_model":{"description":"REST API resource type or endpoint","key$":"api_model"},"api_link":{"description":"REST API link for this resource","key$":"api_link"},"title":{"description":"The name of this resource","key$":"title"},"external_sku":{"description":"Numeric product identification code of a machine-readable barcode, when the customer sku differs from our internal one","key$":"external_sku"},"image_url":{"description":"URL of an image for this product","key$":"image_url"},"web_url":{"description":"URL of this product in the shop","key$":"web_url"},"description":{"description":"Explanation of what this product is","key$":"description"},"price_display":{"description":"Explanation of what this product is","key$":"price_display"},"min_compare_at_price":{"description":"Number indicating how much the least expensive variant of a product cost before a sale","key$":"min_compare_at_price"},"max_compare_at_price":{"description":"Number indicating how much the most expensive variant of a product cost before a sale","key$":"max_compare_at_price"},"min_current_price":{"description":"Number indicating how much the least expensive variant of a product costs right now","key$":"min_current_price"},"max_current_price":{"description":"Number indicating how much the most expensive variant of a product costs right now","key$":"max_current_price"},"artist_ids":{"description":"Unique identifiers of the artists associated with this product","key$":"artist_ids"},"artwork_ids":{"description":"Unique identifiers of the artworks associated with this product","key$":"artwork_ids"},"exhibition_ids":{"description":"Unique identifiers of the exhibitions associated with this product","key$":"exhibition_ids"},"suggest_autocomplete_boosted":{"description":"Internal field to power the `/autocomplete` endpoint. Do not use directly.","key$":"suggest_autocomplete_boosted"},"suggest_autocomplete_all":{"description":"Internal field to power the `/autosuggest` endpoint. Do not use directly.","key$":"suggest_autocomplete_all"},"source_updated_at":{"description":"Date and time the resource was updated in the source system","key$":"source_updated_at"},"updated_at":{"description":"Date and time the record was updated in the aggregator database","key$":"updated_at"},"timestamp":{"description":"Date and time the record was updated in the aggregator search index","key$":"timestamp"}},"type":"object","x-ref":"#/definitions/Product","index$":0}}},"default":{"description":"error","schema":{"required":["status","error","detail"],"properties":{"status":{"type":"integer"},"error":{"type":"string"},"detail":{"type":"string"}},"x-ref":"#/definitions/Error"}}},"produces":["application/json"],"parameters":[],"securitySource":"unspecified"},"GET /products/{id}":{"protocol":"http","responses":{"200":{"description":"Successful operation","schema":{"type":"array","items":{"properties":{"id":{"description":"Unique identifier of this resource. Taken from the source system.","key$":"id"},"api_model":{"description":"REST API resource type or endpoint","key$":"api_model"},"api_link":{"description":"REST API link for this resource","key$":"api_link"},"title":{"description":"The name of this resource","key$":"title"},"external_sku":{"description":"Numeric product identification code of a machine-readable barcode, when the customer sku differs from our internal one","key$":"external_sku"},"image_url":{"description":"URL of an image for this product","key$":"image_url"},"web_url":{"description":"URL of this product in the shop","key$":"web_url"},"description":{"description":"Explanation of what this product is","key$":"description"},"price_display":{"description":"Explanation of what this product is","key$":"price_display"},"min_compare_at_price":{"description":"Number indicating how much the least expensive variant of a product cost before a sale","key$":"min_compare_at_price"},"max_compare_at_price":{"description":"Number indicating how much the most expensive variant of a product cost before a sale","key$":"max_compare_at_price"},"min_current_price":{"description":"Number indicating how much the least expensive variant of a product costs right now","key$":"min_current_price"},"max_current_price":{"description":"Number indicating how much the most expensive variant of a product costs right now","key$":"max_current_price"},"artist_ids":{"description":"Unique identifiers of the artists associated with this product","key$":"artist_ids"},"artwork_ids":{"description":"Unique identifiers of the artworks associated with this product","key$":"artwork_ids"},"exhibition_ids":{"description":"Unique identifiers of the exhibitions associated with this product","key$":"exhibition_ids"},"suggest_autocomplete_boosted":{"description":"Internal field to power the `/autocomplete` endpoint. Do not use directly.","key$":"suggest_autocomplete_boosted"},"suggest_autocomplete_all":{"description":"Internal field to power the `/autosuggest` endpoint. Do not use directly.","key$":"suggest_autocomplete_all"},"source_updated_at":{"description":"Date and time the resource was updated in the source system","key$":"source_updated_at"},"updated_at":{"description":"Date and time the record was updated in the aggregator database","key$":"updated_at"},"timestamp":{"description":"Date and time the record was updated in the aggregator search index","key$":"timestamp"}},"type":"object","x-ref":"#/definitions/Product","key$":"items"}}},"default":{"description":"error","schema":{"required":["status","error","detail"],"properties":{"status":{"type":"integer"},"error":{"type":"string"},"detail":{"type":"string"}},"x-ref":"#/definitions/Error"}}},"produces":["application/json"],"parameters":[{"name":"id","in":"path","type":"string","required":true,"x-ref":"#/parameters/id","index$":0}],"securitySource":"unspecified"}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let product_ref01_data = Object.values(setup.data.existing.product)[0] as any

    // LIST
    const product_ref01_ent = client.Product()
    const product_ref01_match: any = {}

    const product_ref01_list = (await product_ref01_ent.list(product_ref01_match)).map((e: any) => e.data())


    // LOAD
    const product_ref01_match_dt0: any = {}
    product_ref01_match_dt0.id = product_ref01_data.id
    const product_ref01_data_dt0 = (await product_ref01_ent.load(product_ref01_match_dt0)).data()
    assert(product_ref01_data_dt0.id === product_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/product/ProductTestData.json')

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
    ['product01','product02','product03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'ART_INSTITUTE_OF_CHICAGO_TEST_PRODUCT_ENTID': idmap,
    'ART_INSTITUTE_OF_CHICAGO_TEST_LIVE': 'FALSE',
    'ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['ART_INSTITUTE_OF_CHICAGO_TEST_PRODUCT_ENTID']

  const live = 'TRUE' === env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['ART_INSTITUTE_OF_CHICAGO_TEST_PRODUCT_ENTID']
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
  
