<?php
declare(strict_types=1);

// EventOccurrence entity test

require_once __DIR__ . '/../artinstituteofchicago_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class EventOccurrenceEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = ArtInstituteOfChicagoSDK::test(null, null);
        $ent = $testsdk->EventOccurrence(null);
        $this->assertNotNull($ent);
    }

    // Feature #4: the entity stream(action, ...) method runs the op pipeline
    // and yields result items. With the streaming feature active it yields the
    // feature's incremental output; otherwise it falls back to the materialised
    // list so stream always yields.
    public function test_stream(): void
    {
        $seed = [
            "entity" => [
                "event_occurrence" => [
                    "s1" => ["id" => "s1"],
                    "s2" => ["id" => "s2"],
                    "s3" => ["id" => "s3"],
                ],
            ],
        ];

        // Fallback: streaming inactive -> yields the materialised list items.
        $base = ArtInstituteOfChicagoSDK::test($seed, null);
        $seen = iterator_to_array($base->EventOccurrence(null)->stream("list", null, null), false);
        $this->assertCount(3, $seen);

        // Inbound: streaming active -> yields each item from the feature.
        $cfg = ArtInstituteOfChicagoConfig::shared_config();
        if (isset($cfg["feature"]) && is_array($cfg["feature"]) && isset($cfg["feature"]["streaming"])) {
            $sdk = ArtInstituteOfChicagoSDK::test($seed, ["feature" => ["streaming" => ["active" => true]]]);
            $got = [];
            foreach ($sdk->EventOccurrence(null)->stream("list", null, null) as $item) {
                if (is_array($item) && array_is_list($item)) {
                    foreach ($item as $sub) {
                        $got[] = $sub;
                    }
                } else {
                    $got[] = $item;
                }
            }
            $this->assertCount(3, $got);
        }
    }

    public function test_basic_flow(): void
    {
        $setup = event_occurrence_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["list", "load"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "event_occurrence." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_OCCURRENCE_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // Bootstrap entity data from existing test data.
        $event_occurrence_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.event_occurrence")));
        $event_occurrence_ref01_data = null;
        if (count($event_occurrence_ref01_data_raw) > 0) {
            $event_occurrence_ref01_data = Helpers::to_map($event_occurrence_ref01_data_raw[0][1]);
        }

        // LIST
        $event_occurrence_ref01_ent = $client->EventOccurrence(null);
        $event_occurrence_ref01_match = [];

        $event_occurrence_ref01_list_result = $event_occurrence_ref01_ent->list($event_occurrence_ref01_match, null);
        $this->assertIsArray($event_occurrence_ref01_list_result);

        // LOAD
        $event_occurrence_ref01_match_dt0 = [
            "id" => $event_occurrence_ref01_data["id"],
        ];
        $event_occurrence_ref01_data_dt0_loaded = $event_occurrence_ref01_ent->load($event_occurrence_ref01_match_dt0, null);
        $event_occurrence_ref01_data_dt0_load_result = Helpers::to_map(is_object($event_occurrence_ref01_data_dt0_loaded) && method_exists($event_occurrence_ref01_data_dt0_loaded, 'data_get') ? $event_occurrence_ref01_data_dt0_loaded->data_get() : $event_occurrence_ref01_data_dt0_loaded);
        $this->assertNotNull($event_occurrence_ref01_data_dt0_load_result);
        $this->assertEquals($event_occurrence_ref01_data_dt0_load_result["id"], $event_occurrence_ref01_data["id"]);

    }
}

function event_occurrence_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/event_occurrence/EventOccurrenceTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = ArtInstituteOfChicagoSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["event_occurrence01", "event_occurrence02", "event_occurrence03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_OCCURRENCE_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_OCCURRENCE_ENTID" => $idmap,
        "ART_INSTITUTE_OF_CHICAGO_TEST_LIVE" => "FALSE",
        "ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN" => "FALSE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_OCCURRENCE_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["ART_INSTITUTE_OF_CHICAGO_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            [
            ],
            $extra ?? [],
        ]);
        $client = new ArtInstituteOfChicagoSDK(Helpers::to_map($merged_opts));
    }

    $live = $env["ART_INSTITUTE_OF_CHICAGO_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}
