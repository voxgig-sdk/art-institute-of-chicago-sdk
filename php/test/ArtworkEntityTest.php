<?php
declare(strict_types=1);

// Artwork entity test

require_once __DIR__ . '/../artinstituteofchicago_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class ArtworkEntityTestFailHook extends ArtInstituteOfChicagoBaseFeature
{
    public int $unexpected = 0;

    public function __construct()
    {
        parent::__construct();
        $this->name = 'failhook';
    }

    public function init(ArtInstituteOfChicagoContext $ctx, array $options): void
    {
    }

    public function PreSpec(ArtInstituteOfChicagoContext $ctx): void
    {
        throw new \RuntimeException('artwork hook failed');
    }

    public function PreUnexpected(ArtInstituteOfChicagoContext $ctx): void
    {
        $this->unexpected++;
    }
}

class ArtworkEntityTest extends TestCase
{
    // main.kit.test.live.strict is true (the default is true): a live
    // request that fails, or a live test missing an input it needs,
    // fails the test.
    // An account with no record for a test to read skips it either way.
    private const LIVE_STRICT = true;

    public function test_create_instance(): void
    {
        $testsdk = ArtInstituteOfChicagoSDK::test(null, null);
        $ent = $testsdk->Artwork(null);
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
                "artwork" => [
                    "s1" => ["id" => "s1"],
                    "s2" => ["id" => "s2"],
                    "s3" => ["id" => "s3"],
                ],
            ],
        ];

        // Fallback: streaming inactive -> yields the materialised list items.
        $base = ArtInstituteOfChicagoSDK::test($seed, null);
        $seen = iterator_to_array($base->Artwork(null)->stream("list", null, null), false);
        $this->assertCount(3, $seen);

        // Inbound: streaming active -> yields each item from the feature.
        $cfg = ArtInstituteOfChicagoConfig::shared_config();
        if (isset($cfg["feature"]) && is_array($cfg["feature"]) && isset($cfg["feature"]["streaming"])) {
            $sdk = ArtInstituteOfChicagoSDK::test($seed, ["feature" => ["streaming" => ["active" => true]]]);
            $got = [];
            foreach ($sdk->Artwork(null)->stream("list", null, null) as $item) {
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

    public function test_stream_error(): void
    {
        $offline = ["net" => ["offline" => true]];
        $streamerr = null;
        try {
            iterator_to_array(ArtInstituteOfChicagoSDK::test($offline, null)->Artwork(null)
                ->stream("list", null, null), false);
        } catch (\Throwable $e) {
            $streamerr = $e;
        }
        $this->assertNotNull($streamerr, 'the stream should raise the transport failure');
        $this->assertStringContainsString('offline', $streamerr->getMessage());

        iterator_to_array(ArtInstituteOfChicagoSDK::test($offline, null)->Artwork(null)
            ->stream("list", null, ["ctrl" => ["throw" => false]]), false);

        $cfg = ArtInstituteOfChicagoConfig::shared_config();
        if (isset($cfg["feature"]["rbac"])) {
            $denied = ArtInstituteOfChicagoSDK::test(null, ["feature" => ["rbac" => ["active" => true, "deny" => true]]]);
            $denyerr = null;
            try {
                iterator_to_array($denied->Artwork(null)->stream("list", null, null), false);
            } catch (\Throwable $e) {
                $denyerr = $e;
            }
            $this->assertSame('rbac_denied', $denyerr->sdk_code ?? null);
        }
    }

    public function test_stream_ctrl(): void
    {
        $ctrl = ["explain" => []];
        iterator_to_array(ArtInstituteOfChicagoSDK::test(null, null)->Artwork(null)
            ->stream("list", null, ["ctrl" => $ctrl]), false);
        $this->assertSame(["explain"], array_keys($ctrl));
    }

    public function test_unexpected(): void
    {
        $hook = new ArtworkEntityTestFailHook();
        $client = new ArtInstituteOfChicagoSDK(["feature" => ["test" => ["active" => true]], "extend" => [$hook]]);

        $err = null;
        try {
            $client->Artwork(null)->list(null, null);
        } catch (\Throwable $e) {
            $err = $e;
        }
        $this->assertNotNull($err, 'the throwing hook should fail the operation');
        $this->assertStringContainsString('hook failed', $err->getMessage());
        $this->assertGreaterThan(0, $hook->unexpected, 'PreUnexpected did not fire');

        $fired = $hook->unexpected;
        $this->assertNull($client->Artwork(null)->list(null, ["throw" => false]));
        $this->assertGreaterThan($fired, $hook->unexpected, 'PreUnexpected did not fire');
    }

    public function test_cost_commits_a_throwing_transport(): void
    {
        $cfg = ArtInstituteOfChicagoConfig::shared_config();
        if (!isset($cfg["feature"]["cost"])) {
            $this->markTestSkipped('feature not present in this SDK: cost');
        }
        $client = new ArtInstituteOfChicagoSDK([
            "test" => ["active" => true],
            "feature" => ["cost" => ["active" => true, "unit" => 1]],
            "utility" => ["fetcher" => function ($ctx, $url, $fetchdef) {
                throw new \RuntimeException('artwork transport failed');
            }],
        ]);

        $err = null;
        try {
            $client->Artwork(null)->list(null, null);
        } catch (\Throwable $e) {
            $err = $e;
        }
        $this->assertInstanceOf(ArtInstituteOfChicagoError::class, $err);
        $this->assertStringContainsString('transport failed', $err->getMessage());

        $client->Artwork(null)->list(null, ["throw" => false]);
        $this->assertSame(2, $client->_cost["total"]["calls"]);
        $this->assertSame(2, $client->_cost["total"]["attempts"]);
    }

    public function test_validate_list_response(): void
    {
        $cfg = ArtInstituteOfChicagoConfig::shared_config();
        if (!isset($cfg["feature"]["validate"])) {
            $this->markTestSkipped('feature not present in this SDK: validate');
        }
        $opts = ["feature" => ["validate" => ["active" => true, "request" => false, "response" => true]]];
        $listed = ArtInstituteOfChicagoSDK::test(['entity' => ['artwork' => ['v1' => ['alt_artist_ids' => 'x', 'alt_classification_ids' => 'x', 'alt_image_ids' => 'x', 'alt_material_ids' => 'x', 'alt_style_ids' => 'x', 'alt_subject_ids' => 'x', 'alt_technique_ids' => 'x', 'alt_titles' => 'x', 'api_link' => 'x', 'api_model' => 'x', 'artist_display' => 'x', 'artist_ids' => 'x', 'artist_title' => 'x', 'artist_titles' => 'x', 'artwork_type_title' => 'x', 'boost_rank' => 'x', 'catalog_based_search_keyword_titles' => 'x', 'catalogue_display' => 'x', 'category_ids' => 'x', 'category_titles' => 'x', 'classification_ids' => 'x', 'classification_title' => 'x', 'classification_titles' => 'x', 'color' => 'x', 'colorfulness' => 'x', 'copyright_notice' => 'x', 'credit_line' => 'x', 'date_display' => 'x', 'date_end' => 'x', 'date_qualifier_title' => 'x', 'date_start' => 'x', 'department_title' => 'x', 'dimensions' => 'x', 'dimensions_detail' => 'x', 'document_ids' => 'x', 'edition' => 'x', 'exhibition_history' => 'x', 'fiscal_year' => 'x', 'fiscal_year_deaccession' => 'x', 'gallery_title' => 'x', 'image_embedding' => 'x', 'inscriptions' => 'x', 'latlon' => 'x', 'material_ids' => 'x', 'material_titles' => 'x', 'max_zoom_window_size' => 'x', 'medium_display' => 'x', 'on_loan_display' => 'x', 'pageviews' => 'x', 'pageviews_recent' => 'x', 'place_of_origin' => 'x', 'provenance_text' => 'x', 'publication_history' => 'x', 'publishing_verification_level' => 'x', 'section_ids' => 'x', 'section_titles' => 'x', 'short_description' => 'x', 'site_ids' => 'x', 'sound_ids' => 'x', 'source_updated_at' => 'x', 'style_ids' => 'x', 'style_title' => 'x', 'style_titles' => 'x', 'subject_ids' => 'x', 'subject_titles' => 'x', 'suggest_autocomplete_all' => 'x', 'suggest_autocomplete_boosted' => 'x', 'technique_ids' => 'x', 'technique_titles' => 'x', 'term_titles' => 'x', 'text_embedding' => 'x', 'text_ids' => 'x', 'theme_titles' => 'x', 'thumbnail' => 'x', 'timestamp' => 'x', 'updated_at' => 'x', 'video_ids' => 'x']]]], $opts)->Artwork(null)->list(null, null);
        $this->assertSame(['v1'], array_map(fn($ent) => $ent->data_get()['id'] ?? null, $listed));
        $err = null;
        try {
            ArtInstituteOfChicagoSDK::test(['entity' => ['artwork' => ['v1' => ['alt_artist_ids' => 'x', 'alt_classification_ids' => 'x', 'alt_image_ids' => 'x', 'alt_material_ids' => 'x', 'alt_style_ids' => 'x', 'alt_subject_ids' => 'x', 'alt_technique_ids' => 'x', 'alt_titles' => 'x', 'api_link' => 'x', 'api_model' => 'x', 'artist_display' => 'x', 'artist_ids' => 'x', 'artist_title' => 'x', 'artist_titles' => 'x', 'artwork_type_title' => 'x', 'boost_rank' => 'x', 'catalog_based_search_keyword_titles' => 'x', 'catalogue_display' => 'x', 'category_ids' => 'x', 'category_titles' => 'x', 'classification_ids' => 'x', 'classification_title' => 'x', 'classification_titles' => 'x', 'color' => 'x', 'colorfulness' => 'x', 'copyright_notice' => 'x', 'credit_line' => 'x', 'date_display' => 'x', 'date_end' => 'x', 'date_qualifier_title' => 'x', 'date_start' => 'x', 'department_title' => 'x', 'dimensions' => 'x', 'dimensions_detail' => 'x', 'document_ids' => 'x', 'edition' => 'x', 'exhibition_history' => 'x', 'fiscal_year' => 'x', 'fiscal_year_deaccession' => 'x', 'gallery_title' => 'x', 'image_embedding' => 'x', 'inscriptions' => 'x', 'latlon' => 'x', 'material_ids' => 'x', 'material_titles' => 'x', 'max_zoom_window_size' => 'x', 'medium_display' => 'x', 'on_loan_display' => 'x', 'pageviews' => 'x', 'pageviews_recent' => 'x', 'place_of_origin' => 'x', 'provenance_text' => 'x', 'publication_history' => 'x', 'publishing_verification_level' => 'x', 'section_ids' => 'x', 'section_titles' => 'x', 'short_description' => 'x', 'site_ids' => 'x', 'sound_ids' => 'x', 'source_updated_at' => 'x', 'style_ids' => 'x', 'style_title' => 'x', 'style_titles' => 'x', 'subject_ids' => 'x', 'subject_titles' => 'x', 'suggest_autocomplete_all' => 'x', 'suggest_autocomplete_boosted' => 'x', 'technique_ids' => 'x', 'technique_titles' => 'x', 'term_titles' => 'x', 'text_embedding' => 'x', 'text_ids' => 'x', 'theme_titles' => 'x', 'thumbnail' => 'x', 'timestamp' => 'x', 'updated_at' => 'x', 'video_ids' => 'x', 'artist_id' => 1]]]], $opts)->Artwork(null)->list(null, null);
        } catch (\Throwable $e) {
            $err = $e;
        }
        $this->assertSame('validate_failed', $err->sdk_code ?? null);
    }

    public function test_validate(): void
    {
        $cfg = ArtInstituteOfChicagoConfig::shared_config();
        if (!isset($cfg["feature"]["validate"])) {
            $this->markTestSkipped('feature not present in this SDK: validate');
        }
        $client = ArtInstituteOfChicagoSDK::test(null, ["feature" => ["validate" => ["active" => true]]]);
        $err = null;
        try {
            $client->Artwork(null)->list(["artist_id" => 1], null);
        } catch (\Throwable $e) {
            $err = $e;
        }
        $this->assertSame('validate_failed', $err->sdk_code ?? null);
    }

    public function test_basic_flow(): void
    {
        $setup = artwork_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["list", "load"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "artwork." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        $client = $setup["client"];
        if (!empty($setup["live"])) {
            Runner::live_existing($setup, self::LIVE_STRICT, "artwork",
                fn() => $client->Artwork(null)->list([], null));
        }

        // Bootstrap entity data from existing test data.
        $artwork_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.artwork")));
        $artwork_ref01_data = null;
        if (count($artwork_ref01_data_raw) > 0) {
            $artwork_ref01_data = Helpers::to_map($artwork_ref01_data_raw[0][1]);
        }

        // LIST
        $artwork_ref01_ent = $client->Artwork(null);
        $artwork_ref01_match = [];

        $artwork_ref01_list_result = $artwork_ref01_ent->list($artwork_ref01_match, null);
        $this->assertIsArray($artwork_ref01_list_result);

        // LOAD
        $artwork_ref01_match_dt0 = [
            "id" => $artwork_ref01_data["id"],
        ];
        $artwork_ref01_data_dt0_loaded = $artwork_ref01_ent->load($artwork_ref01_match_dt0, null);
        $artwork_ref01_data_dt0_load_result = Helpers::to_map(is_object($artwork_ref01_data_dt0_loaded) && method_exists($artwork_ref01_data_dt0_loaded, 'data_get') ? $artwork_ref01_data_dt0_loaded->data_get() : $artwork_ref01_data_dt0_loaded);
        $this->assertNotNull($artwork_ref01_data_dt0_load_result);
        $this->assertEquals($artwork_ref01_data_dt0_load_result["id"], $artwork_ref01_data["id"]);

    }
}

function artwork_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/artwork/ArtworkTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = ArtInstituteOfChicagoSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["artwork01", "artwork02", "artwork03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Whether *_ENTID supplied the idmap, read before env_override consumes
    // it: without it, the ids a live flow binds are the fixture's synthetic ones.
    $entid_env_raw = getenv("ART_INSTITUTE_OF_CHICAGO_TEST_ARTWORK_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "ART_INSTITUTE_OF_CHICAGO_TEST_ARTWORK_ENTID" => $idmap,
        "ART_INSTITUTE_OF_CHICAGO_TEST_LIVE" => "FALSE",
        "ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN" => "FALSE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["ART_INSTITUTE_OF_CHICAGO_TEST_ARTWORK_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["ART_INSTITUTE_OF_CHICAGO_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            Runner::live_client_options(),
            [
            ],
            // ismap, not a plain "?? []" default: an empty PHP array is a
            // LIST, and a non-map later entry REPLACES the accumulated map in
            // merge - so the no-extras call discarded live_client_options()
            // and the apikey/server map above it.
            Vs::ismap($extra) ? $extra : new \stdClass(),
        ]);
        // "?? []" because merge legitimately answers with a stdClass when every
        // contributing entry is an EMPTY map - an SDK with no apikey and no
        // server variables generates an empty middle entry, so that is the
        // common case, not the edge one. to_map returns null for a non-array by
        // design, and the constructor takes a non-nullable array, so without the
        // fallback every such SDK died on "must be of type array, null given"
        // the moment live mode was switched on. Offline mode never reaches this
        // branch, which is why the offline suite stayed green.
        $client = new ArtInstituteOfChicagoSDK(Helpers::to_map($merged_opts) ?? []);
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
