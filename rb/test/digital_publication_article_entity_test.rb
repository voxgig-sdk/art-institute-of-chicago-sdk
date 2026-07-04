# DigitalPublicationArticle entity test

require "minitest/autorun"
require "json"
require_relative "../ArtInstituteOfChicago_sdk"
require_relative "runner"

class DigitalPublicationArticleEntityTest < Minitest::Test
  def test_create_instance
    testsdk = ArtInstituteOfChicagoSDK.test(nil, nil)
    ent = testsdk.DigitalPublicationArticle(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = digital_publication_article_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["list", "load"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "digital_publication_article." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set ARTINSTITUTEOFCHICAGO_TEST_DIGITAL_PUBLICATION_ARTICLE_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # Bootstrap entity data from existing test data.
    digital_publication_article_ref01_data_raw = Vs.items(Helpers.to_map(
      Vs.getpath(setup[:data], "existing.digital_publication_article")))
    digital_publication_article_ref01_data = nil
    if digital_publication_article_ref01_data_raw.length > 0
      digital_publication_article_ref01_data = Helpers.to_map(digital_publication_article_ref01_data_raw[0][1])
    end

    # LIST
    digital_publication_article_ref01_ent = client.DigitalPublicationArticle(nil)
    digital_publication_article_ref01_match = {}

    digital_publication_article_ref01_list_result = digital_publication_article_ref01_ent.list(digital_publication_article_ref01_match, nil)
    assert digital_publication_article_ref01_list_result.is_a?(Array)

    # LOAD
    digital_publication_article_ref01_match_dt0 = {
      "id" => digital_publication_article_ref01_data["id"],
    }
    digital_publication_article_ref01_data_dt0_loaded = digital_publication_article_ref01_ent.load(digital_publication_article_ref01_match_dt0, nil)
    digital_publication_article_ref01_data_dt0_load_result = Helpers.to_map(digital_publication_article_ref01_data_dt0_loaded)
    assert !digital_publication_article_ref01_data_dt0_load_result.nil?
    assert_equal digital_publication_article_ref01_data_dt0_load_result["id"], digital_publication_article_ref01_data["id"]

  end
end

def digital_publication_article_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "digital_publication_article", "DigitalPublicationArticleTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = ArtInstituteOfChicagoSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["digital_publication_article01", "digital_publication_article02", "digital_publication_article03"],
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
  entid_env_raw = ENV["ARTINSTITUTEOFCHICAGO_TEST_DIGITAL_PUBLICATION_ARTICLE_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "ARTINSTITUTEOFCHICAGO_TEST_DIGITAL_PUBLICATION_ARTICLE_ENTID" => idmap,
    "ARTINSTITUTEOFCHICAGO_TEST_LIVE" => "FALSE",
    "ARTINSTITUTEOFCHICAGO_TEST_EXPLAIN" => "FALSE",
  })

  idmap_resolved = Helpers.to_map(
    env["ARTINSTITUTEOFCHICAGO_TEST_DIGITAL_PUBLICATION_ARTICLE_ENTID"])
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
