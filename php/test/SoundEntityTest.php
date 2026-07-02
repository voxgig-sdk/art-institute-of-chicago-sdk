<?php
declare(strict_types=1);

// Sound entity test

require_once __DIR__ . '/../artinstituteofchicago_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class SoundEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = ArtInstituteOfChicagoSDK::test(null, null);
        $ent = $testsdk->Sound(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = sound_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["list", "load"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "sound." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set ARTINSTITUTEOFCHICAGO_TEST_SOUND_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // Bootstrap entity data from existing test data.
        $sound_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.sound")));
        $sound_ref01_data = null;
        if (count($sound_ref01_data_raw) > 0) {
            $sound_ref01_data = Helpers::to_map($sound_ref01_data_raw[0][1]);
        }

        // LIST
        $sound_ref01_ent = $client->Sound(null);
        $sound_ref01_match = [];

        [$sound_ref01_list_result, $err] = $sound_ref01_ent->list($sound_ref01_match, null);
        $this->assertNull($err);
        $this->assertIsArray($sound_ref01_list_result);

        // LOAD
        $sound_ref01_match_dt0 = [
            "id" => $sound_ref01_data["id"],
        ];
        [$sound_ref01_data_dt0_loaded, $err] = $sound_ref01_ent->load($sound_ref01_match_dt0, null);
        $this->assertNull($err);
        $sound_ref01_data_dt0_load_result = Helpers::to_map($sound_ref01_data_dt0_loaded);
        $this->assertNotNull($sound_ref01_data_dt0_load_result);
        $this->assertEquals($sound_ref01_data_dt0_load_result["id"], $sound_ref01_data["id"]);

    }
}

function sound_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/sound/SoundTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = ArtInstituteOfChicagoSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["sound01", "sound02", "sound03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("ARTINSTITUTEOFCHICAGO_TEST_SOUND_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "ARTINSTITUTEOFCHICAGO_TEST_SOUND_ENTID" => $idmap,
        "ARTINSTITUTEOFCHICAGO_TEST_LIVE" => "FALSE",
        "ARTINSTITUTEOFCHICAGO_TEST_EXPLAIN" => "FALSE",
        "ARTINSTITUTEOFCHICAGO_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["ARTINSTITUTEOFCHICAGO_TEST_SOUND_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["ARTINSTITUTEOFCHICAGO_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            [
                "apikey" => $env["ARTINSTITUTEOFCHICAGO_APIKEY"],
            ],
            $extra ?? [],
        ]);
        $client = new ArtInstituteOfChicagoSDK(Helpers::to_map($merged_opts));
    }

    $live = $env["ARTINSTITUTEOFCHICAGO_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["ARTINSTITUTEOFCHICAGO_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}
