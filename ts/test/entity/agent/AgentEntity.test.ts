

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


describe('AgentEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE.
  afterEach(liveDelay('ART_INSTITUTE_OF_CHICAGO_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = ArtInstituteOfChicagoSDK.test()
    const ent = testsdk.Agent()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE
    for (const op of ['list', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'agent.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"alt_titles","req":false,"short":"Alternate names for this agent","type":"`$ANY`","index$":0},{"active":true,"name":"api_link","req":false,"short":"REST API link for this resource","type":"`$ANY`","index$":1},{"active":true,"name":"api_model","req":false,"short":"REST API resource type or endpoint","type":"`$ANY`","index$":2},{"active":true,"name":"birth_date","req":false,"short":"The year this agent was born","type":"`$ANY`","index$":3},{"active":true,"name":"death_date","req":false,"short":"The year this agent died","type":"`$ANY`","index$":4},{"active":true,"name":"description","req":false,"short":"A biographical description of the agent","type":"`$STRING`","index$":5},{"active":true,"name":"id","req":false,"short":"Unique identifier of this resource.","type":"`$STRING`","index$":6},{"active":true,"name":"is_artist","req":false,"short":"Whether the agent is an artist.","type":"`$BOOLEAN`","index$":7},{"active":true,"name":"sort_title","req":false,"short":"Sortable name for this agent, typically with last name first.","type":"`$ANY`","index$":8},{"active":true,"name":"source_updated_at","req":false,"short":"Date and time the resource was updated in the source system","type":"`$ANY`","index$":9},{"active":true,"name":"suggest_autocomplete_all","req":false,"short":"Internal field to power the `/autosuggest` endpoint.","type":"`$ANY`","index$":10},{"active":true,"name":"suggest_autocomplete_boosted","req":false,"short":"Internal field to power the `/autocomplete` endpoint.","type":"`$ANY`","index$":11},{"active":true,"name":"timestamp","req":false,"short":"Date and time the record was updated in the aggregator search index","type":"`$ANY`","index$":12},{"active":true,"name":"title","req":false,"short":"The name of this resource","type":"`$STRING`","index$":13},{"active":true,"name":"ulan_id","req":false,"short":"Unique identifier of this agent in Getty's ULAN","type":"`$STRING`","index$":14},{"active":true,"name":"updated_at","req":false,"short":"Date and time the record was updated in the aggregator database","type":"`$ANY`","index$":15}],"id":{"field":"id","name":"id"},"name":"agent","op":{"list":{"input":"data","name":"list","points":[{"active":true,"args":{},"contract":{"id":"GET /agents","json":"{\"parameters\":[],\"produces\":[\"application/json\"],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"Successful operation\",\"schema\":{\"items\":{\"properties\":{\"alt_titles\":{\"description\":\"Alternate names for this agent\"},\"api_link\":{\"description\":\"REST API link for this resource\"},\"api_model\":{\"description\":\"REST API resource type or endpoint\"},\"birth_date\":{\"description\":\"The year this agent was born\"},\"death_date\":{\"description\":\"The year this agent died\"},\"description\":{\"description\":\"A biographical description of the agent\"},\"id\":{\"description\":\"Unique identifier of this resource. Taken from the source system.\"},\"is_artist\":{\"description\":\"Whether the agent is an artist. Solely based on whether the agent is related to an artwork record.\"},\"sort_title\":{\"description\":\"Sortable name for this agent, typically with last name first.\"},\"source_updated_at\":{\"description\":\"Date and time the resource was updated in the source system\"},\"suggest_autocomplete_all\":{\"description\":\"Internal field to power the `/autosuggest` endpoint. Do not use directly.\"},\"suggest_autocomplete_boosted\":{\"description\":\"Internal field to power the `/autocomplete` endpoint. Do not use directly.\"},\"timestamp\":{\"description\":\"Date and time the record was updated in the aggregator search index\"},\"title\":{\"description\":\"The name of this resource\"},\"ulan_id\":{\"description\":\"Unique identifier of this agent in Getty's ULAN\"},\"updated_at\":{\"description\":\"Date and time the record was updated in the aggregator database\"}},\"type\":\"object\"},\"type\":\"array\"}},\"default\":{\"description\":\"error\",\"schema\":{\"properties\":{\"detail\":{\"type\":\"string\"},\"error\":{\"type\":\"string\"},\"status\":{\"type\":\"integer\"}},\"required\":[\"status\",\"error\",\"detail\"]}}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/agents","segments":[{"lit":"agents"}],"select":{},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0},{"active":true,"args":{},"contract":{"id":"GET /artists","json":"{\"parameters\":[],\"produces\":[\"application/json\"],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"Successful operation\",\"schema\":{\"items\":{\"properties\":{\"alt_titles\":{\"description\":\"Alternate names for this agent\"},\"api_link\":{\"description\":\"REST API link for this resource\"},\"api_model\":{\"description\":\"REST API resource type or endpoint\"},\"birth_date\":{\"description\":\"The year this agent was born\"},\"death_date\":{\"description\":\"The year this agent died\"},\"description\":{\"description\":\"A biographical description of the agent\"},\"id\":{\"description\":\"Unique identifier of this resource. Taken from the source system.\"},\"is_artist\":{\"description\":\"Whether the agent is an artist. Solely based on whether the agent is related to an artwork record.\"},\"sort_title\":{\"description\":\"Sortable name for this agent, typically with last name first.\"},\"source_updated_at\":{\"description\":\"Date and time the resource was updated in the source system\"},\"suggest_autocomplete_all\":{\"description\":\"Internal field to power the `/autosuggest` endpoint. Do not use directly.\"},\"suggest_autocomplete_boosted\":{\"description\":\"Internal field to power the `/autocomplete` endpoint. Do not use directly.\"},\"timestamp\":{\"description\":\"Date and time the record was updated in the aggregator search index\"},\"title\":{\"description\":\"The name of this resource\"},\"ulan_id\":{\"description\":\"Unique identifier of this agent in Getty's ULAN\"},\"updated_at\":{\"description\":\"Date and time the record was updated in the aggregator database\"}},\"type\":\"object\"},\"type\":\"array\"}},\"default\":{\"description\":\"error\",\"schema\":{\"properties\":{\"detail\":{\"type\":\"string\"},\"error\":{\"type\":\"string\"},\"status\":{\"type\":\"integer\"}},\"required\":[\"status\",\"error\",\"detail\"]}}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/artists","segments":[{"lit":"artists"}],"select":{},"transform":{"req":"`reqdata`","res":"`body`"},"index$":1}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"id","orig":"id","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"GET /agents/{id}","json":"{\"parameters\":[{\"in\":\"path\",\"name\":\"id\",\"required\":true,\"type\":\"string\"}],\"produces\":[\"application/json\"],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"Successful operation\",\"schema\":{\"items\":{\"properties\":{\"alt_titles\":{\"description\":\"Alternate names for this agent\"},\"api_link\":{\"description\":\"REST API link for this resource\"},\"api_model\":{\"description\":\"REST API resource type or endpoint\"},\"birth_date\":{\"description\":\"The year this agent was born\"},\"death_date\":{\"description\":\"The year this agent died\"},\"description\":{\"description\":\"A biographical description of the agent\"},\"id\":{\"description\":\"Unique identifier of this resource. Taken from the source system.\"},\"is_artist\":{\"description\":\"Whether the agent is an artist. Solely based on whether the agent is related to an artwork record.\"},\"sort_title\":{\"description\":\"Sortable name for this agent, typically with last name first.\"},\"source_updated_at\":{\"description\":\"Date and time the resource was updated in the source system\"},\"suggest_autocomplete_all\":{\"description\":\"Internal field to power the `/autosuggest` endpoint. Do not use directly.\"},\"suggest_autocomplete_boosted\":{\"description\":\"Internal field to power the `/autocomplete` endpoint. Do not use directly.\"},\"timestamp\":{\"description\":\"Date and time the record was updated in the aggregator search index\"},\"title\":{\"description\":\"The name of this resource\"},\"ulan_id\":{\"description\":\"Unique identifier of this agent in Getty's ULAN\"},\"updated_at\":{\"description\":\"Date and time the record was updated in the aggregator database\"}},\"type\":\"object\"},\"type\":\"array\"}},\"default\":{\"description\":\"error\",\"schema\":{\"properties\":{\"detail\":{\"type\":\"string\"},\"error\":{\"type\":\"string\"},\"status\":{\"type\":\"integer\"}},\"required\":[\"status\",\"error\",\"detail\"]}}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/agents/{id}","segments":[{"lit":"agents"},{"var":"id"}],"select":{"exist":["id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0},{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"id","orig":"id","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"GET /artists/{id}","json":"{\"parameters\":[{\"in\":\"path\",\"name\":\"id\",\"required\":true,\"type\":\"string\"}],\"produces\":[\"application/json\"],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"Successful operation\",\"schema\":{\"items\":{\"properties\":{\"alt_titles\":{\"description\":\"Alternate names for this agent\"},\"api_link\":{\"description\":\"REST API link for this resource\"},\"api_model\":{\"description\":\"REST API resource type or endpoint\"},\"birth_date\":{\"description\":\"The year this agent was born\"},\"death_date\":{\"description\":\"The year this agent died\"},\"description\":{\"description\":\"A biographical description of the agent\"},\"id\":{\"description\":\"Unique identifier of this resource. Taken from the source system.\"},\"is_artist\":{\"description\":\"Whether the agent is an artist. Solely based on whether the agent is related to an artwork record.\"},\"sort_title\":{\"description\":\"Sortable name for this agent, typically with last name first.\"},\"source_updated_at\":{\"description\":\"Date and time the resource was updated in the source system\"},\"suggest_autocomplete_all\":{\"description\":\"Internal field to power the `/autosuggest` endpoint. Do not use directly.\"},\"suggest_autocomplete_boosted\":{\"description\":\"Internal field to power the `/autocomplete` endpoint. Do not use directly.\"},\"timestamp\":{\"description\":\"Date and time the record was updated in the aggregator search index\"},\"title\":{\"description\":\"The name of this resource\"},\"ulan_id\":{\"description\":\"Unique identifier of this agent in Getty's ULAN\"},\"updated_at\":{\"description\":\"Date and time the record was updated in the aggregator database\"}},\"type\":\"object\"},\"type\":\"array\"}},\"default\":{\"description\":\"error\",\"schema\":{\"properties\":{\"detail\":{\"type\":\"string\"},\"error\":{\"type\":\"string\"},\"status\":{\"type\":\"integer\"}},\"required\":[\"status\",\"error\",\"detail\"]}}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/artists/{id}","segments":[{"lit":"artists"},{"var":"id"}],"select":{"exist":["id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":1}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"agent","name__orig":"agent","Name":"Agent","name_":"agent","name-":"agent","NAME":"AGENT","index$":0}, {"active":true,"entity":"agent","key$":"BasicAgentFlow","kind":"basic","name":"BasicAgentFlow","param":{},"step":[{"active":true,"data":{},"input":{},"match":{},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"agent_ref01"}}],"index$":0},{"active":true,"data":{},"input":{"ref":"agent_ref01","srcdatavar":"agent_ref01_data","suffix":"_dt0"},"match":{"id":"agent01"},"op":"load","spec":[],"valid":[{"apply":"TextFieldMark","def":{"mark":"Mark01-agent_ref01"}}],"index$":1}]}, 'Agent')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let agent_ref01_data = Object.values(setup.data.existing.agent)[0] as any

    // LIST
    const agent_ref01_ent = client.Agent()
    const agent_ref01_match: any = {}

    const agent_ref01_list = (await agent_ref01_ent.list(agent_ref01_match)).map((e: any) => e.data())


    // LOAD
    const agent_ref01_match_dt0: any = {}
    agent_ref01_match_dt0.id = agent_ref01_data.id
    const agent_ref01_data_dt0 = (await agent_ref01_ent.load(agent_ref01_match_dt0)).data()
    assert(agent_ref01_data_dt0.id === agent_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/agent/AgentTestData.json')

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
    ['agent01','agent02','agent03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'ART_INSTITUTE_OF_CHICAGO_TEST_AGENT_ENTID': idmap,
    'ART_INSTITUTE_OF_CHICAGO_TEST_LIVE': 'FALSE',
    'ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['ART_INSTITUTE_OF_CHICAGO_TEST_AGENT_ENTID']

  const live = 'TRUE' === env.ART_INSTITUTE_OF_CHICAGO_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['ART_INSTITUTE_OF_CHICAGO_TEST_AGENT_ENTID']
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
  
