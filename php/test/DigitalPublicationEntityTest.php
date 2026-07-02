<?php
declare(strict_types=1);

// DigitalPublication entity test

require_once __DIR__ . '/../artinstituteofchicago_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class DigitalPublicationEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = ArtInstituteOfChicagoSDK::test(null, null);
        $ent = $testsdk->DigitalPublication(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = digital_publication_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["list", "load"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "digital_publication." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set ARTINSTITUTEOFCHICAGO_TEST_DIGITAL_PUBLICATION_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // Bootstrap entity data from existing test data.
        $digital_publication_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.digital_publication")));
        $digital_publication_ref01_data = null;
        if (count($digital_publication_ref01_data_raw) > 0) {
            $digital_publication_ref01_data = Helpers::to_map($digital_publication_ref01_data_raw[0][1]);
        }

        // LIST
        $digital_publication_ref01_ent = $client->DigitalPublication(null);
        $digital_publication_ref01_match = [];

        [$digital_publication_ref01_list_result, $err] = $digital_publication_ref01_ent->list($digital_publication_ref01_match, null);
        $this->assertNull($err);
        $this->assertIsArray($digital_publication_ref01_list_result);

        // LOAD
        $digital_publication_ref01_match_dt0 = [
            "id" => $digital_publication_ref01_data["id"],
        ];
        [$digital_publication_ref01_data_dt0_loaded, $err] = $digital_publication_ref01_ent->load($digital_publication_ref01_match_dt0, null);
        $this->assertNull($err);
        $digital_publication_ref01_data_dt0_load_result = Helpers::to_map($digital_publication_ref01_data_dt0_loaded);
        $this->assertNotNull($digital_publication_ref01_data_dt0_load_result);
        $this->assertEquals($digital_publication_ref01_data_dt0_load_result["id"], $digital_publication_ref01_data["id"]);

    }
}

function digital_publication_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/digital_publication/DigitalPublicationTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = ArtInstituteOfChicagoSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["digital_publication01", "digital_publication02", "digital_publication03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("ARTINSTITUTEOFCHICAGO_TEST_DIGITAL_PUBLICATION_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "ARTINSTITUTEOFCHICAGO_TEST_DIGITAL_PUBLICATION_ENTID" => $idmap,
        "ARTINSTITUTEOFCHICAGO_TEST_LIVE" => "FALSE",
        "ARTINSTITUTEOFCHICAGO_TEST_EXPLAIN" => "FALSE",
        "ARTINSTITUTEOFCHICAGO_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["ARTINSTITUTEOFCHICAGO_TEST_DIGITAL_PUBLICATION_ENTID"]);
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
