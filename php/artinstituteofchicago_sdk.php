<?php
declare(strict_types=1);

// ArtInstituteOfChicago SDK

require_once __DIR__ . '/utility/struct/Struct.php';
require_once __DIR__ . '/core/UtilityType.php';
require_once __DIR__ . '/core/Spec.php';
require_once __DIR__ . '/core/Helpers.php';

// Load utility registration
require_once __DIR__ . '/utility/Register.php';

// Load config and features
require_once __DIR__ . '/config.php';
require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/features.php';

use Voxgig\Struct\Struct;

// Features record diagnostic state on the client as dynamic properties
// (_retry, _cache, _metrics, ...); allow them explicitly (PHP 8.2+
// deprecates implicit dynamic properties).
#[\AllowDynamicProperties]
class ArtInstituteOfChicagoSDK
{
    public string $mode;
    public array $features;
    public ?array $options;

    private $_utility;
    private $_rootctx;

    public function __construct(array $options = [])
    {
        $this->mode = "live";
        $this->features = [];
        $this->options = null;

        $utility = new ArtInstituteOfChicagoUtility();
        $this->_utility = $utility;

        $config = ArtInstituteOfChicagoConfig::make_config();

        $this->_rootctx = ($utility->make_context)([
            "client" => $this,
            "utility" => $utility,
            "config" => $config,
            "options" => $options ?? [],
            "shared" => [],
        ], null);

        $this->options = ($utility->make_options)($this->_rootctx);

        if (Struct::getpath($this->options, "feature.test.active") === true) {
            $this->mode = "test";
        }

        $this->_rootctx->options = $this->options;

        // Add features in the resolved order (make_options puts an explicit
        // list order first, else defaults to test-first). Ordering matters: the
        // `test` feature installs the base mock transport and the transport
        // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is
        // current, so `test` must be added before them to sit at the base.
        $feature_opts = ArtInstituteOfChicagoHelpers::to_map(Struct::getprop($this->options, "feature"));
        if ($feature_opts) {
            $featureorder = Struct::getpath($this->options, "__derived__.featureorder");
            if (is_array($featureorder)) {
                foreach ($featureorder as $fname) {
                    $fopts = ArtInstituteOfChicagoHelpers::to_map($feature_opts[$fname] ?? null);
                    if ($fopts && isset($fopts["active"]) && $fopts["active"] === true) {
                        ($utility->feature_add)($this->_rootctx, ArtInstituteOfChicagoFeatures::make_feature($fname));
                    }
                }
            }
        }

        // Add extension features.
        $extend_val = Struct::getprop($this->options, "extend");
        if (is_array($extend_val)) {
            foreach ($extend_val as $f) {
                if (is_object($f) && method_exists($f, 'get_name')) {
                    ($utility->feature_add)($this->_rootctx, $f);
                }
            }
        }

        // Initialize features.
        foreach ($this->features as $f) {
            ($utility->feature_init)($this->_rootctx, $f);
        }

        ($utility->feature_hook)($this->_rootctx, "PostConstruct");
    }

    public function options_map(): array
    {
        $out = Struct::clone($this->options);
        return is_array($out) ? $out : [];
    }

    public function get_utility()
    {
        return ArtInstituteOfChicagoUtility::copy($this->_utility);
    }

    public function get_root_ctx()
    {
        return $this->_rootctx;
    }

    public function prepare(array $fetchargs = []): mixed
    {
        $utility = $this->_utility;
        $fetchargs = $fetchargs ?? [];

        $ctrl = ArtInstituteOfChicagoHelpers::to_map(Struct::getprop($fetchargs, "ctrl")) ?? [];

        $ctx = ($utility->make_context)([
            "opname" => "prepare",
            "ctrl" => $ctrl,
        ], $this->_rootctx);

        $opts = $this->options;
        $path = Struct::getprop($fetchargs, "path") ?? "";
        $path = is_string($path) ? $path : "";
        $method_val = Struct::getprop($fetchargs, "method") ?? "GET";
        $method_val = is_string($method_val) ? $method_val : "GET";
        $params = ArtInstituteOfChicagoHelpers::to_map(Struct::getprop($fetchargs, "params")) ?? [];
        $query = ArtInstituteOfChicagoHelpers::to_map(Struct::getprop($fetchargs, "query")) ?? [];
        $headers = ($utility->prepare_headers)($ctx);

        $base = Struct::getprop($opts, "base") ?? "";
        $base = is_string($base) ? $base : "";
        $prefix = Struct::getprop($opts, "prefix") ?? "";
        $prefix = is_string($prefix) ? $prefix : "";
        $suffix = Struct::getprop($opts, "suffix") ?? "";
        $suffix = is_string($suffix) ? $suffix : "";

        $ctx->spec = new ArtInstituteOfChicagoSpec([
            "base" => $base, "prefix" => $prefix, "suffix" => $suffix,
            "path" => $path, "method" => $method_val,
            "params" => $params, "query" => $query, "headers" => $headers,
            "body" => Struct::getprop($fetchargs, "body"),
            "step" => "start",
        ]);

        // Merge user-provided headers.
        $uh = Struct::getprop($fetchargs, "headers");
        if (is_array($uh)) {
            foreach ($uh as $k => $v) {
                $ctx->spec->headers[$k] = $v;
            }
        }

        [$_, $err] = ($utility->prepare_auth)($ctx);
        if ($err) {
            return ($utility->make_error)($ctx, $err);
        }

        [$fetchdef, $fd_err] = ($utility->make_fetch_def)($ctx);
        if ($fd_err) {
            return ($utility->make_error)($ctx, $fd_err);
        }
        return $fetchdef;
    }

    // Raw endpoint access is operator-controllable, like every entity op.
    // Blocking it means denying BOTH the 'direct' and 'graphql' tokens,
    // since either one reaches the same endpoint.
    public function direct(array $fetchargs = []): mixed
    {
        if (!$this->op_allowed("direct")) {
            return $this->op_denied("direct");
        }

        return $this->raw_request($fetchargs);
    }

    // Is this raw-access op permitted by the SDK's allow.op option?
    private function op_allowed(string $op): bool
    {
        $allow_op = Struct::getpath($this->options, "allow.op");
        return is_string($allow_op) && str_contains($allow_op, $op);
    }

    private function op_denied(string $op): array
    {
        $allow_op = Struct::getpath($this->options, "allow.op");
        return [
            "ok" => false,
            "err" => new ArtInstituteOfChicagoError($op . "_allow",
                "ArtInstituteOfChicagoSDK: " . $op . ": operation not allowed by" .
                " SDK option allow.op value: \"" . (string)$allow_op . "\""),
        ];
    }

    // Ungated request path shared by direct and graphql, each of which
    // checks its own allow.op token first. Private, rather than a flag on
    // fetchargs: a caller-supplied marker would let anyone opt straight back
    // out of the gate by passing it.
    private function raw_request(array $fetchargs = []): mixed
    {
        $utility = $this->_utility;

        // direct() is the raw-HTTP escape hatch: it never throws, it returns
        // an {ok, err, ...} dict. prepare() now raises on error, so catch it
        // and surface the failure through the dict instead.
        try {
            $fetchdef = $this->prepare($fetchargs);
        } catch (\Throwable $err) {
            return ["ok" => false, "err" => $err];
        }

        $fetchargs = $fetchargs ?? [];
        $ctrl = ArtInstituteOfChicagoHelpers::to_map(Struct::getprop($fetchargs, "ctrl")) ?? [];

        $ctx = ($utility->make_context)([
            "opname" => "direct",
            "ctrl" => $ctrl,
        ], $this->_rootctx);

        $url = $fetchdef["url"] ?? "";
        [$fetched, $fetch_err] = ($utility->fetcher)($ctx, $url, $fetchdef);

        if ($fetch_err) {
            return ["ok" => false, "err" => $fetch_err];
        }

        if ($fetched === null) {
            return [
                "ok" => false,
                "err" => $ctx->make_error("direct_no_response", "response: undefined"),
            ];
        }

        if (is_array($fetched)) {
            $status = ArtInstituteOfChicagoHelpers::to_int(Struct::getprop($fetched, "status"));
            $headers = Struct::getprop($fetched, "headers") ?? [];

            // No-body responses (204, 304) and explicit zero content-length
            // must skip JSON parsing — calling json() on an empty body errors.
            $content_length = is_array($headers) ? ($headers["content-length"] ?? null) : null;
            $no_body = $status === 204 || $status === 304 || (string)$content_length === "0";

            $json_data = null;
            if (!$no_body) {
                $jf = Struct::getprop($fetched, "json");
                if (is_callable($jf)) {
                    try {
                        $json_data = $jf();
                    } catch (\Throwable $e) {
                        // Non-JSON body — leave data null but keep status/ok.
                        $json_data = null;
                    }
                }
            }

            return [
                "ok" => $status >= 200 && $status < 300,
                "status" => $status,
                "headers" => Struct::getprop($fetched, "headers"),
                "data" => $json_data,
            ];
        }

        return [
            "ok" => false,
            "err" => $ctx->make_error("direct_invalid", "invalid response type"),
        ];
    }

    // Raw GraphQL access: the pressure valve that makes the generated
    // surface's deliberate omissions (per-call selection sets, typed filter
    // builders, batching, subscriptions) livable — the whole schema stays
    // reachable.
    //
    // Thin wrapper over the same prepare/fetch path direct uses, with the
    // one thing raw direct cannot do for GraphQL: a GraphQL failure rides
    // HTTP 200 as a top-level `errors` array, so status alone would report
    // a failed query as ok.
    //
    // NOTE: like direct, this bypasses the feature pipeline — no retry,
    // ratelimit or paging features apply.
    public function graphql(string $query, ?array $variables = null, ?array $ctrl = null): mixed
    {
        if (!$this->op_allowed("graphql")) {
            return $this->op_denied("graphql");
        }

        $res = $this->raw_request([
            "method" => "POST",
            "headers" => ["content-type" => "application/json"],
            "body" => ["query" => $query, "variables" => $variables ?? []],
            "ctrl" => $ctrl ?? [],
        ]);

        if (!is_array($res)) {
            return $res;
        }

        // Errors are read BEFORE any status check: a GraphQL parse or
        // validation failure comes back as HTTP 400 carrying the standard
        // { errors: [...] } body, and the raw path represents a non-2xx as
        // ok:false with no err — so returning early on status would discard
        // the server's own diagnostics, which are the only useful part of
        // that response.
        $errors = Struct::getpath($res, "data.errors");

        if (is_array($errors) && 0 < count($errors)) {
            $first = is_array($errors[0]) ? $errors[0] : [];
            $msg = $first["message"] ?? "";
            if (!is_string($msg) || "" === $msg) {
                $msg = "graphql error";
            }
            $res["ok"] = false;
            $res["err"] = new ArtInstituteOfChicagoError("graphql_error",
                "ArtInstituteOfChicagoSDK: graphql: " . $msg);
            $res["graphql"] = $errors;
        }

        return $res;
    }


    private $_agent = null;

    // Canonical facade: $client->Agent()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->agent()
    // resolves here too.
    public function Agent($data = null)
    {
        require_once __DIR__ . '/entity/agent_entity.php';
        if ($data === null) {
            if ($this->_agent === null) {
                $this->_agent = new AgentEntity($this, null);
            }
            return $this->_agent;
        }
        return new AgentEntity($this, $data);
    }


    private $_agent_role = null;

    // Canonical facade: $client->AgentRole()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->agent_role()
    // resolves here too.
    public function AgentRole($data = null)
    {
        require_once __DIR__ . '/entity/agent_role_entity.php';
        if ($data === null) {
            if ($this->_agent_role === null) {
                $this->_agent_role = new AgentRoleEntity($this, null);
            }
            return $this->_agent_role;
        }
        return new AgentRoleEntity($this, $data);
    }


    private $_agent_type = null;

    // Canonical facade: $client->AgentType()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->agent_type()
    // resolves here too.
    public function AgentType($data = null)
    {
        require_once __DIR__ . '/entity/agent_type_entity.php';
        if ($data === null) {
            if ($this->_agent_type === null) {
                $this->_agent_type = new AgentTypeEntity($this, null);
            }
            return $this->_agent_type;
        }
        return new AgentTypeEntity($this, $data);
    }


    private $_article = null;

    // Canonical facade: $client->Article()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->article()
    // resolves here too.
    public function Article($data = null)
    {
        require_once __DIR__ . '/entity/article_entity.php';
        if ($data === null) {
            if ($this->_article === null) {
                $this->_article = new ArticleEntity($this, null);
            }
            return $this->_article;
        }
        return new ArticleEntity($this, $data);
    }


    private $_artwork = null;

    // Canonical facade: $client->Artwork()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->artwork()
    // resolves here too.
    public function Artwork($data = null)
    {
        require_once __DIR__ . '/entity/artwork_entity.php';
        if ($data === null) {
            if ($this->_artwork === null) {
                $this->_artwork = new ArtworkEntity($this, null);
            }
            return $this->_artwork;
        }
        return new ArtworkEntity($this, $data);
    }


    private $_artwork_date_qualifier = null;

    // Canonical facade: $client->ArtworkDateQualifier()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->artwork_date_qualifier()
    // resolves here too.
    public function ArtworkDateQualifier($data = null)
    {
        require_once __DIR__ . '/entity/artwork_date_qualifier_entity.php';
        if ($data === null) {
            if ($this->_artwork_date_qualifier === null) {
                $this->_artwork_date_qualifier = new ArtworkDateQualifierEntity($this, null);
            }
            return $this->_artwork_date_qualifier;
        }
        return new ArtworkDateQualifierEntity($this, $data);
    }


    private $_artwork_place_qualifier = null;

    // Canonical facade: $client->ArtworkPlaceQualifier()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->artwork_place_qualifier()
    // resolves here too.
    public function ArtworkPlaceQualifier($data = null)
    {
        require_once __DIR__ . '/entity/artwork_place_qualifier_entity.php';
        if ($data === null) {
            if ($this->_artwork_place_qualifier === null) {
                $this->_artwork_place_qualifier = new ArtworkPlaceQualifierEntity($this, null);
            }
            return $this->_artwork_place_qualifier;
        }
        return new ArtworkPlaceQualifierEntity($this, $data);
    }


    private $_artwork_type = null;

    // Canonical facade: $client->ArtworkType()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->artwork_type()
    // resolves here too.
    public function ArtworkType($data = null)
    {
        require_once __DIR__ . '/entity/artwork_type_entity.php';
        if ($data === null) {
            if ($this->_artwork_type === null) {
                $this->_artwork_type = new ArtworkTypeEntity($this, null);
            }
            return $this->_artwork_type;
        }
        return new ArtworkTypeEntity($this, $data);
    }


    private $_category_term = null;

    // Canonical facade: $client->CategoryTerm()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->category_term()
    // resolves here too.
    public function CategoryTerm($data = null)
    {
        require_once __DIR__ . '/entity/category_term_entity.php';
        if ($data === null) {
            if ($this->_category_term === null) {
                $this->_category_term = new CategoryTermEntity($this, null);
            }
            return $this->_category_term;
        }
        return new CategoryTermEntity($this, $data);
    }


    private $_digital_publication = null;

    // Canonical facade: $client->DigitalPublication()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->digital_publication()
    // resolves here too.
    public function DigitalPublication($data = null)
    {
        require_once __DIR__ . '/entity/digital_publication_entity.php';
        if ($data === null) {
            if ($this->_digital_publication === null) {
                $this->_digital_publication = new DigitalPublicationEntity($this, null);
            }
            return $this->_digital_publication;
        }
        return new DigitalPublicationEntity($this, $data);
    }


    private $_digital_publication_article = null;

    // Canonical facade: $client->DigitalPublicationArticle()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->digital_publication_article()
    // resolves here too.
    public function DigitalPublicationArticle($data = null)
    {
        require_once __DIR__ . '/entity/digital_publication_article_entity.php';
        if ($data === null) {
            if ($this->_digital_publication_article === null) {
                $this->_digital_publication_article = new DigitalPublicationArticleEntity($this, null);
            }
            return $this->_digital_publication_article;
        }
        return new DigitalPublicationArticleEntity($this, $data);
    }


    private $_educator_resource = null;

    // Canonical facade: $client->EducatorResource()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->educator_resource()
    // resolves here too.
    public function EducatorResource($data = null)
    {
        require_once __DIR__ . '/entity/educator_resource_entity.php';
        if ($data === null) {
            if ($this->_educator_resource === null) {
                $this->_educator_resource = new EducatorResourceEntity($this, null);
            }
            return $this->_educator_resource;
        }
        return new EducatorResourceEntity($this, $data);
    }


    private $_event = null;

    // Canonical facade: $client->Event()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->event()
    // resolves here too.
    public function Event($data = null)
    {
        require_once __DIR__ . '/entity/event_entity.php';
        if ($data === null) {
            if ($this->_event === null) {
                $this->_event = new EventEntity($this, null);
            }
            return $this->_event;
        }
        return new EventEntity($this, $data);
    }


    private $_event_occurrence = null;

    // Canonical facade: $client->EventOccurrence()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->event_occurrence()
    // resolves here too.
    public function EventOccurrence($data = null)
    {
        require_once __DIR__ . '/entity/event_occurrence_entity.php';
        if ($data === null) {
            if ($this->_event_occurrence === null) {
                $this->_event_occurrence = new EventOccurrenceEntity($this, null);
            }
            return $this->_event_occurrence;
        }
        return new EventOccurrenceEntity($this, $data);
    }


    private $_event_program = null;

    // Canonical facade: $client->EventProgram()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->event_program()
    // resolves here too.
    public function EventProgram($data = null)
    {
        require_once __DIR__ . '/entity/event_program_entity.php';
        if ($data === null) {
            if ($this->_event_program === null) {
                $this->_event_program = new EventProgramEntity($this, null);
            }
            return $this->_event_program;
        }
        return new EventProgramEntity($this, $data);
    }


    private $_exhibition = null;

    // Canonical facade: $client->Exhibition()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->exhibition()
    // resolves here too.
    public function Exhibition($data = null)
    {
        require_once __DIR__ . '/entity/exhibition_entity.php';
        if ($data === null) {
            if ($this->_exhibition === null) {
                $this->_exhibition = new ExhibitionEntity($this, null);
            }
            return $this->_exhibition;
        }
        return new ExhibitionEntity($this, $data);
    }


    private $_gallery = null;

    // Canonical facade: $client->Gallery()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->gallery()
    // resolves here too.
    public function Gallery($data = null)
    {
        require_once __DIR__ . '/entity/gallery_entity.php';
        if ($data === null) {
            if ($this->_gallery === null) {
                $this->_gallery = new GalleryEntity($this, null);
            }
            return $this->_gallery;
        }
        return new GalleryEntity($this, $data);
    }


    private $_generic_page = null;

    // Canonical facade: $client->GenericPage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->generic_page()
    // resolves here too.
    public function GenericPage($data = null)
    {
        require_once __DIR__ . '/entity/generic_page_entity.php';
        if ($data === null) {
            if ($this->_generic_page === null) {
                $this->_generic_page = new GenericPageEntity($this, null);
            }
            return $this->_generic_page;
        }
        return new GenericPageEntity($this, $data);
    }


    private $_highlight = null;

    // Canonical facade: $client->Highlight()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->highlight()
    // resolves here too.
    public function Highlight($data = null)
    {
        require_once __DIR__ . '/entity/highlight_entity.php';
        if ($data === null) {
            if ($this->_highlight === null) {
                $this->_highlight = new HighlightEntity($this, null);
            }
            return $this->_highlight;
        }
        return new HighlightEntity($this, $data);
    }


    private $_hour = null;

    // Canonical facade: $client->Hour()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->hour()
    // resolves here too.
    public function Hour($data = null)
    {
        require_once __DIR__ . '/entity/hour_entity.php';
        if ($data === null) {
            if ($this->_hour === null) {
                $this->_hour = new HourEntity($this, null);
            }
            return $this->_hour;
        }
        return new HourEntity($this, $data);
    }


    private $_image = null;

    // Canonical facade: $client->Image()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->image()
    // resolves here too.
    public function Image($data = null)
    {
        require_once __DIR__ . '/entity/image_entity.php';
        if ($data === null) {
            if ($this->_image === null) {
                $this->_image = new ImageEntity($this, null);
            }
            return $this->_image;
        }
        return new ImageEntity($this, $data);
    }


    private $_landing_page = null;

    // Canonical facade: $client->LandingPage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->landing_page()
    // resolves here too.
    public function LandingPage($data = null)
    {
        require_once __DIR__ . '/entity/landing_page_entity.php';
        if ($data === null) {
            if ($this->_landing_page === null) {
                $this->_landing_page = new LandingPageEntity($this, null);
            }
            return $this->_landing_page;
        }
        return new LandingPageEntity($this, $data);
    }


    private $_place = null;

    // Canonical facade: $client->Place()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->place()
    // resolves here too.
    public function Place($data = null)
    {
        require_once __DIR__ . '/entity/place_entity.php';
        if ($data === null) {
            if ($this->_place === null) {
                $this->_place = new PlaceEntity($this, null);
            }
            return $this->_place;
        }
        return new PlaceEntity($this, $data);
    }


    private $_press_release = null;

    // Canonical facade: $client->PressRelease()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->press_release()
    // resolves here too.
    public function PressRelease($data = null)
    {
        require_once __DIR__ . '/entity/press_release_entity.php';
        if ($data === null) {
            if ($this->_press_release === null) {
                $this->_press_release = new PressReleaseEntity($this, null);
            }
            return $this->_press_release;
        }
        return new PressReleaseEntity($this, $data);
    }


    private $_printed_publication = null;

    // Canonical facade: $client->PrintedPublication()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->printed_publication()
    // resolves here too.
    public function PrintedPublication($data = null)
    {
        require_once __DIR__ . '/entity/printed_publication_entity.php';
        if ($data === null) {
            if ($this->_printed_publication === null) {
                $this->_printed_publication = new PrintedPublicationEntity($this, null);
            }
            return $this->_printed_publication;
        }
        return new PrintedPublicationEntity($this, $data);
    }


    private $_product = null;

    // Canonical facade: $client->Product()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->product()
    // resolves here too.
    public function Product($data = null)
    {
        require_once __DIR__ . '/entity/product_entity.php';
        if ($data === null) {
            if ($this->_product === null) {
                $this->_product = new ProductEntity($this, null);
            }
            return $this->_product;
        }
        return new ProductEntity($this, $data);
    }


    private $_publication = null;

    // Canonical facade: $client->Publication()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->publication()
    // resolves here too.
    public function Publication($data = null)
    {
        require_once __DIR__ . '/entity/publication_entity.php';
        if ($data === null) {
            if ($this->_publication === null) {
                $this->_publication = new PublicationEntity($this, null);
            }
            return $this->_publication;
        }
        return new PublicationEntity($this, $data);
    }


    private $_search = null;

    // Canonical facade: $client->Search()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->search()
    // resolves here too.
    public function Search($data = null)
    {
        require_once __DIR__ . '/entity/search_entity.php';
        if ($data === null) {
            if ($this->_search === null) {
                $this->_search = new SearchEntity($this, null);
            }
            return $this->_search;
        }
        return new SearchEntity($this, $data);
    }


    private $_section = null;

    // Canonical facade: $client->Section()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->section()
    // resolves here too.
    public function Section($data = null)
    {
        require_once __DIR__ . '/entity/section_entity.php';
        if ($data === null) {
            if ($this->_section === null) {
                $this->_section = new SectionEntity($this, null);
            }
            return $this->_section;
        }
        return new SectionEntity($this, $data);
    }


    private $_site = null;

    // Canonical facade: $client->Site()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->site()
    // resolves here too.
    public function Site($data = null)
    {
        require_once __DIR__ . '/entity/site_entity.php';
        if ($data === null) {
            if ($this->_site === null) {
                $this->_site = new SiteEntity($this, null);
            }
            return $this->_site;
        }
        return new SiteEntity($this, $data);
    }


    private $_sound = null;

    // Canonical facade: $client->Sound()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->sound()
    // resolves here too.
    public function Sound($data = null)
    {
        require_once __DIR__ . '/entity/sound_entity.php';
        if ($data === null) {
            if ($this->_sound === null) {
                $this->_sound = new SoundEntity($this, null);
            }
            return $this->_sound;
        }
        return new SoundEntity($this, $data);
    }


    private $_static_page = null;

    // Canonical facade: $client->StaticPage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->static_page()
    // resolves here too.
    public function StaticPage($data = null)
    {
        require_once __DIR__ . '/entity/static_page_entity.php';
        if ($data === null) {
            if ($this->_static_page === null) {
                $this->_static_page = new StaticPageEntity($this, null);
            }
            return $this->_static_page;
        }
        return new StaticPageEntity($this, $data);
    }


    private $_text = null;

    // Canonical facade: $client->Text()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->text()
    // resolves here too.
    public function Text($data = null)
    {
        require_once __DIR__ . '/entity/text_entity.php';
        if ($data === null) {
            if ($this->_text === null) {
                $this->_text = new TextEntity($this, null);
            }
            return $this->_text;
        }
        return new TextEntity($this, $data);
    }


    private $_tour = null;

    // Canonical facade: $client->Tour()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->tour()
    // resolves here too.
    public function Tour($data = null)
    {
        require_once __DIR__ . '/entity/tour_entity.php';
        if ($data === null) {
            if ($this->_tour === null) {
                $this->_tour = new TourEntity($this, null);
            }
            return $this->_tour;
        }
        return new TourEntity($this, $data);
    }


    private $_video = null;

    // Canonical facade: $client->Video()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->video()
    // resolves here too.
    public function Video($data = null)
    {
        require_once __DIR__ . '/entity/video_entity.php';
        if ($data === null) {
            if ($this->_video === null) {
                $this->_video = new VideoEntity($this, null);
            }
            return $this->_video;
        }
        return new VideoEntity($this, $data);
    }



    public static function test(?array $testopts = null, ?array $sdkopts = null): self
    {
        $sdkopts = $sdkopts ?? [];
        $sdkopts = Struct::clone($sdkopts);
        $sdkopts = is_array($sdkopts) ? $sdkopts : [];

        $testopts = $testopts ?? [];
        $testopts = Struct::clone($testopts);
        $testopts = is_array($testopts) ? $testopts : [];
        $testopts["active"] = true;

        if (!isset($sdkopts["feature"])) {
            $sdkopts["feature"] = [];
        }
        $sdkopts["feature"]["test"] = $testopts;

        $sdk = new ArtInstituteOfChicagoSDK($sdkopts);
        $sdk->mode = "test";
        return $sdk;
    }
}
