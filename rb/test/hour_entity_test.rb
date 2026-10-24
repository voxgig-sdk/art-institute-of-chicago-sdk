# Hour entity test

require "minitest/autorun"
require "json"
require_relative "../ArtInstituteOfChicago_sdk"
require_relative "runner"

class HourEntityTest < Minitest::Test
  # main.kit.test.live.strict is true (the default is true): a live
  # request that fails, or a live test missing an input it needs,
  # fails the test.
  # An account with no record for a test to read skips it either way.
  LIVE_STRICT = true

  def test_create_instance
    testsdk = ArtInstituteOfChicagoSDK.test(nil, nil)
    ent = testsdk.Hour(nil)
    assert !ent.nil?
  end

  def test_list_entities
    seed = {
      "entity" => {
        "hour" => {
          "l1" => { "id" => "l1" },
          "l2" => { "id" => "l2" },
        },
      },
    }
    items = ArtInstituteOfChicagoSDK.test(seed, nil).Hour(nil).list(nil, nil)
    # list resolves to one entity per record; data_get reads the record.
    assert_equal 2, items.length
    items.each do |item|
      assert item.respond_to?(:data_get)
      assert item.data_get.is_a?(Hash)
    end
  end

  # Feature #4: the entity stream(action, ...) method runs the op pipeline and
  # returns an Enumerator over result items. With the streaming feature active
  # it yields the feature's incremental output; otherwise it falls back to the
  # materialised list so stream always yields.
  def test_stream
    seed = {
      "entity" => {
        "hour" => {
          "s1" => { "id" => "s1" },
          "s2" => { "id" => "s2" },
          "s3" => { "id" => "s3" },
        },
      },
    }

    # Fallback: streaming inactive -> yields the materialised list items.
    base = ArtInstituteOfChicagoSDK.test(seed, nil)
    seen = base.Hour(nil).stream("list", nil, nil).to_a
    assert_equal 3, seen.length

    # Inbound: streaming active -> yields each item from the feature.
    cfg = ArtInstituteOfChicagoConfig.shared_config
    if cfg["feature"].is_a?(Hash) && cfg["feature"].key?("streaming")
      sdk = ArtInstituteOfChicagoSDK.test(seed, { "feature" => { "streaming" => { "active" => true } } })
      got = []
      sdk.Hour(nil).stream("list", nil, nil).each do |item|
        if item.is_a?(Array)
          got.concat(item)
        else
          got << item
        end
      end
      assert_equal 3, got.length
    end
  end

  class FailHook < ArtInstituteOfChicagoBaseFeature
    attr_reader :unexpected

    def initialize
      super()
      @name = "failhook"
      @unexpected = 0
    end

    def PreSpec(ctx)
      raise "hour hook failed"
    end

    def PreUnexpected(ctx)
      @unexpected += 1
    end
  end

  def test_stream_error
    offline = { "net" => { "offline" => true } }
    err = assert_raises(StandardError) do
      ArtInstituteOfChicagoSDK.test(offline, nil).Hour(nil).stream("list", nil, nil).to_a
    end
    assert_match(/offline/, err.message)

    ArtInstituteOfChicagoSDK.test(offline, nil).Hour(nil)
      .stream("list", nil, { "ctrl" => { "throw" => false } }).to_a

    cfg = ArtInstituteOfChicagoConfig.shared_config
    if cfg["feature"].is_a?(Hash) && cfg["feature"].key?("rbac")
      denied = ArtInstituteOfChicagoSDK.test(nil, { "feature" => { "rbac" => { "active" => true, "deny" => true } } })
      err = assert_raises(StandardError) do
        denied.Hour(nil).stream("list", nil, nil).to_a
      end
      assert_equal "rbac_denied", err.code
    end
  end

  def test_stream_ctrl
    explain = {}
    ctrl = { "explain" => explain }
    ArtInstituteOfChicagoSDK.test(nil, nil).Hour(nil).stream("list", nil, { "ctrl" => ctrl }).to_a
    assert_equal ["explain"], ctrl.keys
    assert_same explain, ctrl["explain"]
    refute_empty explain
  end

  def test_unexpected
    hook = FailHook.new
    client = ArtInstituteOfChicagoSDK.new({ "feature" => { "test" => { "active" => true } }, "extend" => [hook] })

    err = assert_raises(StandardError) do
      client.Hour(nil).list(nil, nil)
    end
    assert_match(/hook failed/, err.message)
    assert_operator hook.unexpected, :>, 0

    fired = hook.unexpected
    assert_nil client.Hour(nil).list(nil, { "throw" => false })
    assert_operator hook.unexpected, :>, fired
  end

  def test_validate_list_response
    cfg = ArtInstituteOfChicagoConfig.shared_config
    unless cfg["feature"].is_a?(Hash) && cfg["feature"].key?("validate")
      skip("feature not present in this SDK: validate")
    end
    opts = { "feature" => { "validate" => { "active" => true, "request" => false, "response" => true } } }
    listed = ArtInstituteOfChicagoSDK.test({ "entity" => { "hour" => { "v1" => { "additional_text" => "x", "api_link" => "x", "api_model" => "x", "friday_is_closed" => "x", "friday_member_close" => "x", "friday_member_open" => "x", "friday_public_close" => "x", "friday_public_open" => "x", "monday_is_closed" => "x", "monday_member_close" => "x", "monday_member_open" => "x", "monday_public_close" => "x", "monday_public_open" => "x", "saturday_is_closed" => "x", "saturday_member_close" => "x", "saturday_member_open" => "x", "saturday_public_close" => "x", "saturday_public_open" => "x", "source_updated_at" => "x", "suggest_autocomplete_all" => "x", "suggest_autocomplete_boosted" => "x", "summary" => "x", "sunday_is_closed" => "x", "sunday_member_close" => "x", "sunday_member_open" => "x", "sunday_public_close" => "x", "sunday_public_open" => "x", "thursday_is_closed" => "x", "thursday_member_close" => "x", "thursday_member_open" => "x", "thursday_public_close" => "x", "thursday_public_open" => "x", "timestamp" => "x", "tuesday_is_closed" => "x", "tuesday_member_close" => "x", "tuesday_member_open" => "x", "tuesday_public_close" => "x", "tuesday_public_open" => "x", "updated_at" => "x", "wednesday_is_closed" => "x", "wednesday_member_close" => "x", "wednesday_member_open" => "x", "wednesday_public_close" => "x", "wednesday_public_open" => "x" } } } }, opts).Hour(nil).list(nil, nil)
    assert_equal ["v1"], listed.map { |ent| ent.data_get["id"] }
    err = assert_raises(StandardError) do
      ArtInstituteOfChicagoSDK.test({ "entity" => { "hour" => { "v1" => { "additional_text" => "x", "api_link" => "x", "api_model" => "x", "friday_is_closed" => "x", "friday_member_close" => "x", "friday_member_open" => "x", "friday_public_close" => "x", "friday_public_open" => "x", "monday_is_closed" => "x", "monday_member_close" => "x", "monday_member_open" => "x", "monday_public_close" => "x", "monday_public_open" => "x", "saturday_is_closed" => "x", "saturday_member_close" => "x", "saturday_member_open" => "x", "saturday_public_close" => "x", "saturday_public_open" => "x", "source_updated_at" => "x", "suggest_autocomplete_all" => "x", "suggest_autocomplete_boosted" => "x", "summary" => "x", "sunday_is_closed" => "x", "sunday_member_close" => "x", "sunday_member_open" => "x", "sunday_public_close" => "x", "sunday_public_open" => "x", "thursday_is_closed" => "x", "thursday_member_close" => "x", "thursday_member_open" => "x", "thursday_public_close" => "x", "thursday_public_open" => "x", "timestamp" => "x", "tuesday_is_closed" => "x", "tuesday_member_close" => "x", "tuesday_member_open" => "x", "tuesday_public_close" => "x", "tuesday_public_open" => "x", "updated_at" => "x", "wednesday_is_closed" => "x", "wednesday_member_close" => "x", "wednesday_member_open" => "x", "wednesday_public_close" => "x", "wednesday_public_open" => "x", "title" => 1 } } } }, opts).Hour(nil).list(nil, nil)
    end
    assert_equal "validate_failed", err.code
  end

  def test_validate
    cfg = ArtInstituteOfChicagoConfig.shared_config
    unless cfg["feature"].is_a?(Hash) && cfg["feature"].key?("validate")
      skip("feature not present in this SDK: validate")
    end
    client = ArtInstituteOfChicagoSDK.test(nil, { "feature" => { "validate" => { "active" => true } } })
    err = assert_raises(StandardError) do
      client.Hour(nil).list({ "id" => 1 }, nil)
    end
    assert_equal "validate_failed", err.code
  end

  def test_basic_flow
    setup = hour_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["list", "load"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "hour." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    client = setup[:client]
    if setup[:live]
      Runner.live_existing(setup, LIVE_STRICT, "hour") do
        client.Hour(nil).list({}, nil)
      end
    end

    # Bootstrap entity data from existing test data.
    hour_ref01_data_raw = Vs.items(Helpers.to_map(
      Vs.getpath(setup[:data], "existing.hour")))
    hour_ref01_data = nil
    if hour_ref01_data_raw.length > 0
      hour_ref01_data = Helpers.to_map(hour_ref01_data_raw[0][1])
    end

    # LIST
    hour_ref01_ent = client.Hour(nil)
    hour_ref01_match = {}

    hour_ref01_list_result = hour_ref01_ent.list(hour_ref01_match, nil)
    assert hour_ref01_list_result.is_a?(Array)

    # LOAD
    hour_ref01_match_dt0 = {
      "id" => hour_ref01_data["id"],
    }
    hour_ref01_data_dt0_loaded = hour_ref01_ent.load(hour_ref01_match_dt0, nil)
    hour_ref01_data_dt0_load_result = Helpers.to_map(hour_ref01_data_dt0_loaded.respond_to?(:data_get) ? hour_ref01_data_dt0_loaded.data_get : hour_ref01_data_dt0_loaded)
    assert !hour_ref01_data_dt0_load_result.nil?
    assert_equal hour_ref01_data_dt0_load_result["id"], hour_ref01_data["id"]

  end
end

def hour_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "hour", "HourTestData.json")
  entity_data_source = File.read(entity_data_file, encoding: "UTF-8")
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = ArtInstituteOfChicagoSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["hour01", "hour02", "hour03"],
    {
      "`$PACK`" => ["", {
        "`$KEY`" => "`$COPY`",
        "`$VAL`" => ["`$FORMAT`", "upper", "`$COPY`"],
      }],
    }
  )

  # Whether *_ENTID supplied the idmap, read before env_override consumes
  # it: without it, the ids a live flow binds are the fixture's synthetic ones.
  entid_env_raw = ENV["ART_INSTITUTE_OF_CHICAGO_TEST_HOUR_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "ART_INSTITUTE_OF_CHICAGO_TEST_HOUR_ENTID" => idmap,
    "ART_INSTITUTE_OF_CHICAGO_TEST_LIVE" => "FALSE",
    "ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN" => "FALSE",
  })

  idmap_resolved = Helpers.to_map(
    env["ART_INSTITUTE_OF_CHICAGO_TEST_HOUR_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end

  if env["ART_INSTITUTE_OF_CHICAGO_TEST_LIVE"] == "TRUE"
    merged_opts = Vs.merge([
      # FIRST, so the generated fields below win: sdk-test-control.json's
      # test.client.options adds to the live client, it does not redirect it.
      Runner.live_client_options,
      {
      },
      extra || {},
    ])
    client = ArtInstituteOfChicagoSDK.new(Helpers.to_map(merged_opts))
  end

  live = env["ART_INSTITUTE_OF_CHICAGO_TEST_LIVE"] == "TRUE"
  {
    client: client,
    data: entity_data,
    idmap: idmap_resolved,
    env: env,
    explain: env["ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN"] == "TRUE",
    live: live,
    synthetic_only: live && !idmap_overridden,
    now: (Time.now.to_f * 1000).to_i,
  }
end
