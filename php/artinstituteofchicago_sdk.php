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

        // Add features from config.
        $feature_opts = ArtInstituteOfChicagoHelpers::to_map(Struct::getprop($this->options, "feature"));
        if ($feature_opts) {
            $items = Struct::items($feature_opts);
            if ($items) {
                foreach ($items as $item) {
                    $fname = $item[0];
                    $fopts = ArtInstituteOfChicagoHelpers::to_map($item[1]);
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

    public function direct(array $fetchargs = []): mixed
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


    private $_agent = null;

    // Idiomatic facade: $client->agent()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias Agent() (PHP method
    // names are case-insensitive).
    public function agent($data = null)
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

    // Idiomatic facade: $client->agent_role()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias AgentRole() (PHP method
    // names are case-insensitive).
    public function agent_role($data = null)
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

    // Idiomatic facade: $client->agent_type()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias AgentType() (PHP method
    // names are case-insensitive).
    public function agent_type($data = null)
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

    // Idiomatic facade: $client->article()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias Article() (PHP method
    // names are case-insensitive).
    public function article($data = null)
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

    // Idiomatic facade: $client->artwork()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias Artwork() (PHP method
    // names are case-insensitive).
    public function artwork($data = null)
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

    // Idiomatic facade: $client->artwork_date_qualifier()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias ArtworkDateQualifier() (PHP method
    // names are case-insensitive).
    public function artwork_date_qualifier($data = null)
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

    // Idiomatic facade: $client->artwork_place_qualifier()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias ArtworkPlaceQualifier() (PHP method
    // names are case-insensitive).
    public function artwork_place_qualifier($data = null)
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

    // Idiomatic facade: $client->artwork_type()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias ArtworkType() (PHP method
    // names are case-insensitive).
    public function artwork_type($data = null)
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

    // Idiomatic facade: $client->category_term()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias CategoryTerm() (PHP method
    // names are case-insensitive).
    public function category_term($data = null)
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

    // Idiomatic facade: $client->digital_publication()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias DigitalPublication() (PHP method
    // names are case-insensitive).
    public function digital_publication($data = null)
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

    // Idiomatic facade: $client->digital_publication_article()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias DigitalPublicationArticle() (PHP method
    // names are case-insensitive).
    public function digital_publication_article($data = null)
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

    // Idiomatic facade: $client->educator_resource()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias EducatorResource() (PHP method
    // names are case-insensitive).
    public function educator_resource($data = null)
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

    // Idiomatic facade: $client->event()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias Event() (PHP method
    // names are case-insensitive).
    public function event($data = null)
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

    // Idiomatic facade: $client->event_occurrence()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias EventOccurrence() (PHP method
    // names are case-insensitive).
    public function event_occurrence($data = null)
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

    // Idiomatic facade: $client->event_program()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias EventProgram() (PHP method
    // names are case-insensitive).
    public function event_program($data = null)
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

    // Idiomatic facade: $client->exhibition()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias Exhibition() (PHP method
    // names are case-insensitive).
    public function exhibition($data = null)
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

    // Idiomatic facade: $client->gallery()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias Gallery() (PHP method
    // names are case-insensitive).
    public function gallery($data = null)
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

    // Idiomatic facade: $client->generic_page()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias GenericPage() (PHP method
    // names are case-insensitive).
    public function generic_page($data = null)
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

    // Idiomatic facade: $client->highlight()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias Highlight() (PHP method
    // names are case-insensitive).
    public function highlight($data = null)
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

    // Idiomatic facade: $client->hour()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias Hour() (PHP method
    // names are case-insensitive).
    public function hour($data = null)
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

    // Idiomatic facade: $client->image()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias Image() (PHP method
    // names are case-insensitive).
    public function image($data = null)
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

    // Idiomatic facade: $client->landing_page()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias LandingPage() (PHP method
    // names are case-insensitive).
    public function landing_page($data = null)
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

    // Idiomatic facade: $client->place()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias Place() (PHP method
    // names are case-insensitive).
    public function place($data = null)
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

    // Idiomatic facade: $client->press_release()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias PressRelease() (PHP method
    // names are case-insensitive).
    public function press_release($data = null)
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

    // Idiomatic facade: $client->printed_publication()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias PrintedPublication() (PHP method
    // names are case-insensitive).
    public function printed_publication($data = null)
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

    // Idiomatic facade: $client->product()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias Product() (PHP method
    // names are case-insensitive).
    public function product($data = null)
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

    // Idiomatic facade: $client->publication()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias Publication() (PHP method
    // names are case-insensitive).
    public function publication($data = null)
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

    // Idiomatic facade: $client->search()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias Search() (PHP method
    // names are case-insensitive).
    public function search($data = null)
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

    // Idiomatic facade: $client->section()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias Section() (PHP method
    // names are case-insensitive).
    public function section($data = null)
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

    // Idiomatic facade: $client->site()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias Site() (PHP method
    // names are case-insensitive).
    public function site($data = null)
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

    // Idiomatic facade: $client->sound()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias Sound() (PHP method
    // names are case-insensitive).
    public function sound($data = null)
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

    // Idiomatic facade: $client->static_page()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias StaticPage() (PHP method
    // names are case-insensitive).
    public function static_page($data = null)
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

    // Idiomatic facade: $client->text()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias Text() (PHP method
    // names are case-insensitive).
    public function text($data = null)
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

    // Idiomatic facade: $client->tour()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias Tour() (PHP method
    // names are case-insensitive).
    public function tour($data = null)
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

    // Idiomatic facade: $client->video()->list() / ->load(["id" => ...]).
    // Also serves the deprecated PascalCase alias Video() (PHP method
    // names are case-insensitive).
    public function video($data = null)
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
