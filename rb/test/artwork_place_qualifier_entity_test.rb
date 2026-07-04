# ArtworkPlaceQualifier entity test

require "minitest/autorun"
require "json"
require_relative "../ArtInstituteOfChicago_sdk"
require_relative "runner"

class ArtworkPlaceQualifierEntityTest < Minitest::Test
  def test_create_instance
    testsdk = ArtInstituteOfChicagoSDK.test(nil, nil)
    ent = testsdk.ArtworkPlaceQualifier(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = artwork_place_qualifier_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["list", "load"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "artwork_place_qualifier." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set ARTINSTITUTEOFCHICAGO_TEST_ARTWORK_PLACE_QUALIFIER_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # Bootstrap entity data from existing test data.
    artwork_place_qualifier_ref01_data_raw = Vs.items(Helpers.to_map(
      Vs.getpath(setup[:data], "existing.artwork_place_qualifier")))
    artwork_place_qualifier_ref01_data = nil
    if artwork_place_qualifier_ref01_data_raw.length > 0
      artwork_place_qualifier_ref01_data = Helpers.to_map(artwork_place_qualifier_ref01_data_raw[0][1])
    end

    # LIST
    artwork_place_qualifier_ref01_ent = client.ArtworkPlaceQualifier(nil)
    artwork_place_qualifier_ref01_match = {}

    artwork_place_qualifier_ref01_list_result = artwork_place_qualifier_ref01_ent.list(artwork_place_qualifier_ref01_match, nil)
    assert artwork_place_qualifier_ref01_list_result.is_a?(Array)

    # LOAD
    artwork_place_qualifier_ref01_match_dt0 = {
      "id" => artwork_place_qualifier_ref01_data["id"],
    }
    artwork_place_qualifier_ref01_data_dt0_loaded = artwork_place_qualifier_ref01_ent.load(artwork_place_qualifier_ref01_match_dt0, nil)
    artwork_place_qualifier_ref01_data_dt0_load_result = Helpers.to_map(artwork_place_qualifier_ref01_data_dt0_loaded)
    assert !artwork_place_qualifier_ref01_data_dt0_load_result.nil?
    assert_equal artwork_place_qualifier_ref01_data_dt0_load_result["id"], artwork_place_qualifier_ref01_data["id"]

  end
end

def artwork_place_qualifier_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "artwork_place_qualifier", "ArtworkPlaceQualifierTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = ArtInstituteOfChicagoSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["artwork_place_qualifier01", "artwork_place_qualifier02", "artwork_place_qualifier03"],
    {
      "`$PACK`" => ["", {
        "`$KEY`" => "`$COPY`",
        "`$VAL`" => ["`$FORMAT`", "upper", "`$COPY`"],
      }],
    }
  )

  # Detect ENTID env override before envOverride consumes it. When live
  # mode is on without a real override, the basic test runs against synthetic
  # IDs from the fixture and 4xx's. Surface this so the test can skip.
  entid_env_raw = ENV["ARTINSTITUTEOFCHICAGO_TEST_ARTWORK_PLACE_QUALIFIER_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "ARTINSTITUTEOFCHICAGO_TEST_ARTWORK_PLACE_QUALIFIER_ENTID" => idmap,
    "ARTINSTITUTEOFCHICAGO_TEST_LIVE" => "FALSE",
    "ARTINSTITUTEOFCHICAGO_TEST_EXPLAIN" => "FALSE",
  })

  idmap_resolved = Helpers.to_map(
    env["ARTINSTITUTEOFCHICAGO_TEST_ARTWORK_PLACE_QUALIFIER_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end

  if env["ARTINSTITUTEOFCHICAGO_TEST_LIVE"] == "TRUE"
    merged_opts = Vs.merge([
      {
      },
      extra || {},
    ])
    client = ArtInstituteOfChicagoSDK.new(Helpers.to_map(merged_opts))
  end

  live = env["ARTINSTITUTEOFCHICAGO_TEST_LIVE"] == "TRUE"
  {
    client: client,
    data: entity_data,
    idmap: idmap_resolved,
    env: env,
    explain: env["ARTINSTITUTEOFCHICAGO_TEST_EXPLAIN"] == "TRUE",
    live: live,
    synthetic_only: live && !idmap_overridden,
    now: (Time.now.to_f * 1000).to_i,
  }
end
