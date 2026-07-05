# ArtInstituteOfChicago PHP SDK



The PHP SDK for the ArtInstituteOfChicago API — an entity-oriented client using PHP conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `$client->Agent()` — with named operations (`list`/`load`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to Packagist. Install it from the
GitHub release tag (`php/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/art-institute-of-chicago-sdk/releases](https://github.com/voxgig-sdk/art-institute-of-chicago-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```php
<?php
require_once 'artinstituteofchicago_sdk.php';

$client = new ArtInstituteOfChicagoSDK();
```

### 2. List agent records

```php
try {
    // list() returns an array of Agent records — iterate directly.
    $agents = $client->Agent()->list();
    foreach ($agents as $item) {
        echo $item["id"] . " " . $item["alt_title"] . "\n";
    }
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

### 3. Load an agent

```php
try {
    // load() returns the bare Agent record (throws on error).
    $agent = $client->Agent()->load(["id" => "example_id"]);
    print_r($agent);
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```


## Error handling

Entity operations throw a `\Throwable` on failure, so wrap them in
`try` / `catch`:

```php
try {
    $agents = $client->Agent()->list();
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

`direct()` does **not** throw — it returns the result array. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```php
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example_id"],
]);

if (! $result["ok"]) {
    $err = $result["err"] ?? null;
    echo "request failed: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```php
// direct() is the raw-HTTP escape hatch: it returns a result array
// (it does not throw). Branch on $result["ok"].
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);

if ($result["ok"]) {
    echo $result["status"];  // 200
    print_r($result["data"]);  // response body
} else {
    // On an HTTP error status there is no err (only a transport failure sets
    // it), so fall back to the status code.
    $err = $result["err"] ?? null;
    echo "Error: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```

### Prepare a request without sending it

```php
// prepare() throws on error and returns the fetch definition.
$fetchdef = $client->prepare([
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => ["id" => "example"],
]);

echo $fetchdef["url"];
echo $fetchdef["method"];
print_r($fetchdef["headers"]);
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```php
$client = ArtInstituteOfChicagoSDK::test([
    "entity" => ["agent" => ["test01" => ["id" => "test01"]]],
]);

// Entity ops return the bare mock record (throws on error).
$agent = $client->Agent()->list();
print_r($agent);
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```php
$mock_fetch = function ($url, $init) {
    return [
        [
            "status" => 200,
            "statusText" => "OK",
            "headers" => [],
            "json" => function () { return ["id" => "mock01"]; },
        ],
        null,
    ];
};

$client = new ArtInstituteOfChicagoSDK([
    "base" => "http://localhost:8080",
    "system" => [
        "fetch" => $mock_fetch,
    ],
]);
```

### Run live tests

Create a `.env.local` file at the project root:

```
ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE
```

Then run:

```bash
cd php && ./vendor/bin/phpunit test/
```


## Reference

### ArtInstituteOfChicagoSDK

```php
require_once 'artinstituteofchicago_sdk.php';
$client = new ArtInstituteOfChicagoSDK($options);
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `array` | Feature activation flags. |
| `extend` | `array` | Additional Feature instances to load. |
| `system` | `array` | System overrides (e.g. custom `fetch` callable). |

### test

```php
$client = ArtInstituteOfChicagoSDK::test($testopts, $sdkopts);
```

Creates a test-mode client with mock transport. Both arguments may be `null`.

### ArtInstituteOfChicagoSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `(): array` | Deep copy of current SDK options. |
| `get_utility` | `(): Utility` | Copy of the SDK utility object. |
| `prepare` | `(array $fetchargs): array` | Build an HTTP request definition without sending. |
| `direct` | `(array $fetchargs): array` | Build and send an HTTP request. |
| `Agent` | `($data): AgentEntity` | Create an Agent entity instance. |
| `AgentRole` | `($data): AgentRoleEntity` | Create an AgentRole entity instance. |
| `AgentType` | `($data): AgentTypeEntity` | Create an AgentType entity instance. |
| `Article` | `($data): ArticleEntity` | Create an Article entity instance. |
| `Artwork` | `($data): ArtworkEntity` | Create an Artwork entity instance. |
| `ArtworkDateQualifier` | `($data): ArtworkDateQualifierEntity` | Create an ArtworkDateQualifier entity instance. |
| `ArtworkPlaceQualifier` | `($data): ArtworkPlaceQualifierEntity` | Create an ArtworkPlaceQualifier entity instance. |
| `ArtworkType` | `($data): ArtworkTypeEntity` | Create an ArtworkType entity instance. |
| `CategoryTerm` | `($data): CategoryTermEntity` | Create a CategoryTerm entity instance. |
| `DigitalPublication` | `($data): DigitalPublicationEntity` | Create a DigitalPublication entity instance. |
| `DigitalPublicationArticle` | `($data): DigitalPublicationArticleEntity` | Create a DigitalPublicationArticle entity instance. |
| `EducatorResource` | `($data): EducatorResourceEntity` | Create an EducatorResource entity instance. |
| `Event` | `($data): EventEntity` | Create an Event entity instance. |
| `EventOccurrence` | `($data): EventOccurrenceEntity` | Create an EventOccurrence entity instance. |
| `EventProgram` | `($data): EventProgramEntity` | Create an EventProgram entity instance. |
| `Exhibition` | `($data): ExhibitionEntity` | Create an Exhibition entity instance. |
| `Gallery` | `($data): GalleryEntity` | Create a Gallery entity instance. |
| `GenericPage` | `($data): GenericPageEntity` | Create a GenericPage entity instance. |
| `Highlight` | `($data): HighlightEntity` | Create a Highlight entity instance. |
| `Hour` | `($data): HourEntity` | Create a Hour entity instance. |
| `Image` | `($data): ImageEntity` | Create an Image entity instance. |
| `LandingPage` | `($data): LandingPageEntity` | Create a LandingPage entity instance. |
| `Place` | `($data): PlaceEntity` | Create a Place entity instance. |
| `PressRelease` | `($data): PressReleaseEntity` | Create a PressRelease entity instance. |
| `PrintedPublication` | `($data): PrintedPublicationEntity` | Create a PrintedPublication entity instance. |
| `Product` | `($data): ProductEntity` | Create a Product entity instance. |
| `Publication` | `($data): PublicationEntity` | Create a Publication entity instance. |
| `Search` | `($data): SearchEntity` | Create a Search entity instance. |
| `Section` | `($data): SectionEntity` | Create a Section entity instance. |
| `Site` | `($data): SiteEntity` | Create a Site entity instance. |
| `Sound` | `($data): SoundEntity` | Create a Sound entity instance. |
| `StaticPage` | `($data): StaticPageEntity` | Create a StaticPage entity instance. |
| `Text` | `($data): TextEntity` | Create a Text entity instance. |
| `Tour` | `($data): TourEntity` | Create a Tour entity instance. |
| `Video` | `($data): VideoEntity` | Create a Video entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `($reqmatch, $ctrl): array` | Load a single entity by match criteria. |
| `list` | `(?array $reqmatch = null, $ctrl): array` | List entities matching the criteria (call with no argument to list all). |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the bare result data (an `array` for single-entity
ops, a `list` for `list`) and throw on error. Wrap calls in
`try`/`catch` to handle failures.

The `direct()` escape hatch never throws — it returns a result `array`
you branch on via `$result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `true` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `array` | Response headers. |
| `data` | `mixed` | Parsed JSON response body. |

On error, `ok` is `false` and `$err` contains the error value.

### Entities

#### Agent

| Field | Description |
| --- | --- |
| `alt_title` |  |
| `api_link` |  |
| `api_model` |  |
| `birth_date` |  |
| `death_date` |  |
| `description` |  |
| `id` |  |
| `is_artist` |  |
| `sort_title` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `ulan_id` |  |
| `updated_at` |  |

Operations: List, Load.

API path: `/agents`

#### AgentRole

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: List, Load.

API path: `/agent-roles`

#### AgentType

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: List, Load.

API path: `/agent-types`

#### Article

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `copy` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: List, Load.

API path: `/articles`

#### Artwork

| Field | Description |
| --- | --- |
| `alt_artist_id` |  |
| `alt_classification_id` |  |
| `alt_image_id` |  |
| `alt_material_id` |  |
| `alt_style_id` |  |
| `alt_subject_id` |  |
| `alt_technique_id` |  |
| `alt_title` |  |
| `api_link` |  |
| `api_model` |  |
| `artist_display` |  |
| `artist_id` |  |
| `artist_title` |  |
| `artwork_type_id` |  |
| `artwork_type_title` |  |
| `boost_rank` |  |
| `catalog_based_search_keyword_title` |  |
| `catalogue_display` |  |
| `category_id` |  |
| `category_title` |  |
| `classification_id` |  |
| `classification_title` |  |
| `color` |  |
| `colorfulness` |  |
| `copyright_notice` |  |
| `credit_line` |  |
| `date_display` |  |
| `date_end` |  |
| `date_qualifier_id` |  |
| `date_qualifier_title` |  |
| `date_start` |  |
| `department_id` |  |
| `department_title` |  |
| `description` |  |
| `dimension` |  |
| `dimensions_detail` |  |
| `document_id` |  |
| `edition` |  |
| `exhibition_history` |  |
| `fiscal_year` |  |
| `fiscal_year_deaccession` |  |
| `gallery_id` |  |
| `gallery_title` |  |
| `has_advanced_imaging` |  |
| `has_educational_resource` |  |
| `has_multimedia_resource` |  |
| `has_not_been_viewed_much` |  |
| `id` |  |
| `image_embedding` |  |
| `image_id` |  |
| `inscription` |  |
| `internal_department_id` |  |
| `is_boosted` |  |
| `is_on_view` |  |
| `is_public_domain` |  |
| `is_zoomable` |  |
| `latitude` |  |
| `latlon` |  |
| `longitude` |  |
| `main_reference_number` |  |
| `material_id` |  |
| `material_title` |  |
| `max_zoom_window_size` |  |
| `medium_display` |  |
| `nomisma_id` |  |
| `on_loan_display` |  |
| `pageview` |  |
| `pageviews_recent` |  |
| `place_of_origin` |  |
| `provenance_text` |  |
| `publication_history` |  |
| `publishing_verification_level` |  |
| `section_id` |  |
| `section_title` |  |
| `short_description` |  |
| `site_id` |  |
| `sound_id` |  |
| `source_updated_at` |  |
| `style_id` |  |
| `style_title` |  |
| `subject_id` |  |
| `subject_title` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `technique_id` |  |
| `technique_title` |  |
| `term_title` |  |
| `text_embedding` |  |
| `text_id` |  |
| `theme_title` |  |
| `thumbnail` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `video_id` |  |

Operations: List, Load.

API path: `/artworks`

#### ArtworkDateQualifier

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: List, Load.

API path: `/artwork-date-qualifiers`

#### ArtworkPlaceQualifier

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: List, Load.

API path: `/artwork-place-qualifiers`

#### ArtworkType

| Field | Description |
| --- | --- |
| `aat_id` |  |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: List, Load.

API path: `/artwork-types`

#### CategoryTerm

| Field | Description |
| --- | --- |
| `aat_id` |  |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `parent_id` |  |
| `source_updated_at` |  |
| `subtype` |  |
| `suggest_autocomplete_all` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: List, Load.

API path: `/category-terms`

#### DigitalPublication

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `copy` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: List, Load.

API path: `/digital-publications`

#### DigitalPublicationArticle

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `author_display` |  |
| `copy` |  |
| `digital_publication_id` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: List, Load.

API path: `/digital-publication-articles`

#### EducatorResource

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `copy` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: List, Load.

API path: `/educator-resources`

#### Event

| Field | Description |
| --- | --- |
| `alt_audience_id` |  |
| `alt_event_type_id` |  |
| `api_link` |  |
| `api_model` |  |
| `audience_id` |  |
| `buy_button_caption` |  |
| `buy_button_text` |  |
| `date_display` |  |
| `description` |  |
| `door_time` |  |
| `end_date` |  |
| `end_time` |  |
| `entrance` |  |
| `event_host_id` |  |
| `event_host_title` |  |
| `event_type_id` |  |
| `header_description` |  |
| `hero_caption` |  |
| `id` |  |
| `image_url` |  |
| `is_admission_required` |  |
| `is_after_hour` |  |
| `is_free` |  |
| `is_member_exclusive` |  |
| `is_private` |  |
| `is_registration_required` |  |
| `is_sales_button_hidden` |  |
| `is_sold_out` |  |
| `is_ticketed` |  |
| `is_virtual_event` |  |
| `join_url` |  |
| `layout_type` |  |
| `list_description` |  |
| `location` |  |
| `program_id` |  |
| `program_title` |  |
| `rsvp_link` |  |
| `search_tag` |  |
| `short_description` |  |
| `slug` |  |
| `source_updated_at` |  |
| `start_date` |  |
| `start_time` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `survey_url` |  |
| `ticketed_event_id` |  |
| `timestamp` |  |
| `title` |  |
| `title_display` |  |
| `updated_at` |  |
| `virtual_event_passcode` |  |
| `virtual_event_url` |  |

Operations: List, Load.

API path: `/events`

#### EventOccurrence

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `button_caption` |  |
| `button_text` |  |
| `button_url` |  |
| `description` |  |
| `end_at` |  |
| `event_id` |  |
| `id` |  |
| `image_url` |  |
| `is_private` |  |
| `is_sales_button_hidden` |  |
| `is_ticketed` |  |
| `location` |  |
| `off_sale_at` |  |
| `on_sale_at` |  |
| `short_description` |  |
| `source_updated_at` |  |
| `start_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `title_display` |  |
| `updated_at` |  |

Operations: List, Load.

API path: `/event-occurrences`

#### EventProgram

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `is_affiliate_group` |  |
| `is_event_host` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: List, Load.

API path: `/event-programs`

#### Exhibition

| Field | Description |
| --- | --- |
| `aic_end_at` |  |
| `aic_start_at` |  |
| `alt_image_id` |  |
| `api_link` |  |
| `api_model` |  |
| `artist_id` |  |
| `artwork_id` |  |
| `artwork_title` |  |
| `document_id` |  |
| `gallery_id` |  |
| `gallery_title` |  |
| `id` |  |
| `image_id` |  |
| `image_url` |  |
| `is_featured` |  |
| `is_published` |  |
| `position` |  |
| `short_description` |  |
| `site_id` |  |
| `source_updated_at` |  |
| `status` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: List, Load.

API path: `/exhibitions`

#### Gallery

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `floor` |  |
| `id` |  |
| `is_closed` |  |
| `latitude` |  |
| `latlon` |  |
| `longitude` |  |
| `number` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `tgn_id` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: List, Load.

API path: `/galleries`

#### GenericPage

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `copy` |  |
| `id` |  |
| `search_tag` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: List, Load.

API path: `/generic-pages`

#### Highlight

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `copy` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: List, Load.

API path: `/highlights`

#### Hour

| Field | Description |
| --- | --- |
| `additional_text` |  |
| `api_link` |  |
| `api_model` |  |
| `friday_is_closed` |  |
| `friday_member_close` |  |
| `friday_member_open` |  |
| `friday_public_close` |  |
| `friday_public_open` |  |
| `id` |  |
| `monday_is_closed` |  |
| `monday_member_close` |  |
| `monday_member_open` |  |
| `monday_public_close` |  |
| `monday_public_open` |  |
| `saturday_is_closed` |  |
| `saturday_member_close` |  |
| `saturday_member_open` |  |
| `saturday_public_close` |  |
| `saturday_public_open` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `summary` |  |
| `sunday_is_closed` |  |
| `sunday_member_close` |  |
| `sunday_member_open` |  |
| `sunday_public_close` |  |
| `sunday_public_open` |  |
| `thursday_is_closed` |  |
| `thursday_member_close` |  |
| `thursday_member_open` |  |
| `thursday_public_close` |  |
| `thursday_public_open` |  |
| `timestamp` |  |
| `title` |  |
| `tuesday_is_closed` |  |
| `tuesday_member_close` |  |
| `tuesday_member_open` |  |
| `tuesday_public_close` |  |
| `tuesday_public_open` |  |
| `updated_at` |  |
| `wednesday_is_closed` |  |
| `wednesday_member_close` |  |
| `wednesday_member_open` |  |
| `wednesday_public_close` |  |
| `wednesday_public_open` |  |

Operations: List, Load.

API path: `/hours`

#### Image

| Field | Description |
| --- | --- |
| `ahash` |  |
| `alt_text` |  |
| `api_link` |  |
| `api_model` |  |
| `artwork_id` |  |
| `artwork_title` |  |
| `color` |  |
| `colorfulness` |  |
| `content` |  |
| `content_e_tag` |  |
| `credit_line` |  |
| `fingerprint` |  |
| `height` |  |
| `id` |  |
| `iiif_url` |  |
| `is_educational_resource` |  |
| `is_multimedia_resource` |  |
| `is_teacher_resource` |  |
| `lake_guid` |  |
| `lqip` |  |
| `phash` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `type` |  |
| `updated_at` |  |
| `width` |  |

Operations: List, Load.

API path: `/images`

#### LandingPage

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `copy` |  |
| `id` |  |
| `search_tag` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: List, Load.

API path: `/landing-pages`

#### Place

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `latitude` |  |
| `longitude` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `tgn_id` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: List, Load.

API path: `/places`

#### PressRelease

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `copy` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: List, Load.

API path: `/press-releases`

#### PrintedPublication

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `copy` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: List, Load.

API path: `/printed-publications`

#### Product

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `artist_id` |  |
| `artwork_id` |  |
| `description` |  |
| `exhibition_id` |  |
| `external_sku` |  |
| `id` |  |
| `image_url` |  |
| `max_compare_at_price` |  |
| `max_current_price` |  |
| `min_compare_at_price` |  |
| `min_current_price` |  |
| `price_display` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: List, Load.

API path: `/products`

#### Publication

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `section_id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: List, Load.

API path: `/publications`

#### Search

| Field | Description |
| --- | --- |
| `api_id` |  |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `is_boosted` |  |
| `score` |  |
| `thumbnail` |  |
| `timestamp` |  |
| `title` |  |

Operations: List.

API path: `/agents/search`

#### Section

| Field | Description |
| --- | --- |
| `accession` |  |
| `api_link` |  |
| `api_model` |  |
| `artwork_id` |  |
| `content` |  |
| `generic_page_id` |  |
| `id` |  |
| `publication_id` |  |
| `publication_title` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: List, Load.

API path: `/sections`

#### Site

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `artwork_id` |  |
| `artwork_title` |  |
| `description` |  |
| `exhibition_id` |  |
| `exhibition_title` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: List, Load.

API path: `/sites`

#### Sound

| Field | Description |
| --- | --- |
| `alt_text` |  |
| `api_link` |  |
| `api_model` |  |
| `artwork_id` |  |
| `artwork_title` |  |
| `content` |  |
| `content_e_tag` |  |
| `credit_line` |  |
| `id` |  |
| `is_educational_resource` |  |
| `is_multimedia_resource` |  |
| `is_teacher_resource` |  |
| `lake_guid` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `transcript` |  |
| `type` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: List, Load.

API path: `/mobile-sounds`

#### StaticPage

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: List, Load.

API path: `/static-pages`

#### Text

| Field | Description |
| --- | --- |
| `alt_text` |  |
| `api_link` |  |
| `api_model` |  |
| `artwork_id` |  |
| `artwork_title` |  |
| `content` |  |
| `content_e_tag` |  |
| `credit_line` |  |
| `id` |  |
| `is_educational_resource` |  |
| `is_multimedia_resource` |  |
| `is_teacher_resource` |  |
| `lake_guid` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `type` |  |
| `updated_at` |  |

Operations: List, Load.

API path: `/texts`

#### Tour

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `artist_title` |  |
| `artwork_title` |  |
| `description` |  |
| `id` |  |
| `image` |  |
| `intro` |  |
| `intro_link` |  |
| `intro_transcript` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `weight` |  |

Operations: List, Load.

API path: `/tours`

#### Video

| Field | Description |
| --- | --- |
| `alt_text` |  |
| `api_link` |  |
| `api_model` |  |
| `artwork_id` |  |
| `artwork_title` |  |
| `content` |  |
| `content_e_tag` |  |
| `credit_line` |  |
| `id` |  |
| `is_educational_resource` |  |
| `is_multimedia_resource` |  |
| `is_teacher_resource` |  |
| `lake_guid` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `type` |  |
| `updated_at` |  |

Operations: List, Load.

API path: `/videos`



## Entities


### Agent

Create an instance: `$agent = $client->Agent();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_title` | `mixed` |  |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `birth_date` | `mixed` |  |
| `death_date` | `mixed` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `is_artist` | `bool` |  |
| `sort_title` | `mixed` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `ulan_id` | `string` |  |
| `updated_at` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare Agent record (throws on error).
$agent = $client->Agent()->load(["id" => "agent_id"]);
```

#### Example: List

```php
// list() returns an array of Agent records (throws on error).
$agents = $client->Agent()->list();
```


### AgentRole

Create an instance: `$agent_role = $client->AgentRole();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `id` | `string` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare AgentRole record (throws on error).
$agent_role = $client->AgentRole()->load(["id" => "agent_role_id"]);
```

#### Example: List

```php
// list() returns an array of AgentRole records (throws on error).
$agent_roles = $client->AgentRole()->list();
```


### AgentType

Create an instance: `$agent_type = $client->AgentType();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `id` | `string` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare AgentType record (throws on error).
$agent_type = $client->AgentType()->load(["id" => "agent_type_id"]);
```

#### Example: List

```php
// list() returns an array of AgentType records (throws on error).
$agent_types = $client->AgentType()->list();
```


### Article

Create an instance: `$article = $client->Article();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `copy` | `mixed` |  |
| `id` | `string` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare Article record (throws on error).
$article = $client->Article()->load(["id" => "article_id"]);
```

#### Example: List

```php
// list() returns an array of Article records (throws on error).
$articles = $client->Article()->list();
```


### Artwork

Create an instance: `$artwork = $client->Artwork();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_artist_id` | `string` |  |
| `alt_classification_id` | `string` |  |
| `alt_image_id` | `string` |  |
| `alt_material_id` | `string` |  |
| `alt_style_id` | `string` |  |
| `alt_subject_id` | `string` |  |
| `alt_technique_id` | `string` |  |
| `alt_title` | `mixed` |  |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `artist_display` | `mixed` |  |
| `artist_id` | `string` |  |
| `artist_title` | `mixed` |  |
| `artwork_type_id` | `string` |  |
| `artwork_type_title` | `mixed` |  |
| `boost_rank` | `mixed` |  |
| `catalog_based_search_keyword_title` | `mixed` |  |
| `catalogue_display` | `mixed` |  |
| `category_id` | `string` |  |
| `category_title` | `mixed` |  |
| `classification_id` | `string` |  |
| `classification_title` | `mixed` |  |
| `color` | `mixed` |  |
| `colorfulness` | `mixed` |  |
| `copyright_notice` | `mixed` |  |
| `credit_line` | `mixed` |  |
| `date_display` | `mixed` |  |
| `date_end` | `mixed` |  |
| `date_qualifier_id` | `string` |  |
| `date_qualifier_title` | `mixed` |  |
| `date_start` | `mixed` |  |
| `department_id` | `string` |  |
| `department_title` | `mixed` |  |
| `description` | `string` |  |
| `dimension` | `mixed` |  |
| `dimensions_detail` | `mixed` |  |
| `document_id` | `string` |  |
| `edition` | `mixed` |  |
| `exhibition_history` | `mixed` |  |
| `fiscal_year` | `mixed` |  |
| `fiscal_year_deaccession` | `mixed` |  |
| `gallery_id` | `string` |  |
| `gallery_title` | `mixed` |  |
| `has_advanced_imaging` | `bool` |  |
| `has_educational_resource` | `bool` |  |
| `has_multimedia_resource` | `bool` |  |
| `has_not_been_viewed_much` | `bool` |  |
| `id` | `string` |  |
| `image_embedding` | `mixed` |  |
| `image_id` | `string` |  |
| `inscription` | `mixed` |  |
| `internal_department_id` | `string` |  |
| `is_boosted` | `bool` |  |
| `is_on_view` | `bool` |  |
| `is_public_domain` | `bool` |  |
| `is_zoomable` | `bool` |  |
| `latitude` | `float` |  |
| `latlon` | `mixed` |  |
| `longitude` | `float` |  |
| `main_reference_number` | `int` |  |
| `material_id` | `string` |  |
| `material_title` | `mixed` |  |
| `max_zoom_window_size` | `mixed` |  |
| `medium_display` | `mixed` |  |
| `nomisma_id` | `string` |  |
| `on_loan_display` | `mixed` |  |
| `pageview` | `mixed` |  |
| `pageviews_recent` | `mixed` |  |
| `place_of_origin` | `mixed` |  |
| `provenance_text` | `mixed` |  |
| `publication_history` | `mixed` |  |
| `publishing_verification_level` | `mixed` |  |
| `section_id` | `string` |  |
| `section_title` | `mixed` |  |
| `short_description` | `mixed` |  |
| `site_id` | `string` |  |
| `sound_id` | `string` |  |
| `source_updated_at` | `mixed` |  |
| `style_id` | `string` |  |
| `style_title` | `mixed` |  |
| `subject_id` | `string` |  |
| `subject_title` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `technique_id` | `string` |  |
| `technique_title` | `mixed` |  |
| `term_title` | `mixed` |  |
| `text_embedding` | `mixed` |  |
| `text_id` | `string` |  |
| `theme_title` | `mixed` |  |
| `thumbnail` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |
| `video_id` | `string` |  |

#### Example: Load

```php
// load() returns the bare Artwork record (throws on error).
$artwork = $client->Artwork()->load(["id" => "artwork_id"]);
```

#### Example: List

```php
// list() returns an array of Artwork records (throws on error).
$artworks = $client->Artwork()->list();
```


### ArtworkDateQualifier

Create an instance: `$artwork_date_qualifier = $client->ArtworkDateQualifier();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `id` | `string` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare ArtworkDateQualifier record (throws on error).
$artwork_date_qualifier = $client->ArtworkDateQualifier()->load(["id" => "artwork_date_qualifier_id"]);
```

#### Example: List

```php
// list() returns an array of ArtworkDateQualifier records (throws on error).
$artwork_date_qualifiers = $client->ArtworkDateQualifier()->list();
```


### ArtworkPlaceQualifier

Create an instance: `$artwork_place_qualifier = $client->ArtworkPlaceQualifier();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `id` | `string` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare ArtworkPlaceQualifier record (throws on error).
$artwork_place_qualifier = $client->ArtworkPlaceQualifier()->load(["id" => "artwork_place_qualifier_id"]);
```

#### Example: List

```php
// list() returns an array of ArtworkPlaceQualifier records (throws on error).
$artwork_place_qualifiers = $client->ArtworkPlaceQualifier()->list();
```


### ArtworkType

Create an instance: `$artwork_type = $client->ArtworkType();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `aat_id` | `string` |  |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `id` | `string` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare ArtworkType record (throws on error).
$artwork_type = $client->ArtworkType()->load(["id" => "artwork_type_id"]);
```

#### Example: List

```php
// list() returns an array of ArtworkType records (throws on error).
$artwork_types = $client->ArtworkType()->list();
```


### CategoryTerm

Create an instance: `$category_term = $client->CategoryTerm();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `aat_id` | `string` |  |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `id` | `string` |  |
| `parent_id` | `string` |  |
| `source_updated_at` | `mixed` |  |
| `subtype` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare CategoryTerm record (throws on error).
$category_term = $client->CategoryTerm()->load(["id" => "category_term_id"]);
```

#### Example: List

```php
// list() returns an array of CategoryTerm records (throws on error).
$category_terms = $client->CategoryTerm()->list();
```


### DigitalPublication

Create an instance: `$digital_publication = $client->DigitalPublication();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `copy` | `mixed` |  |
| `id` | `string` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |
| `web_url` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare DigitalPublication record (throws on error).
$digital_publication = $client->DigitalPublication()->load(["id" => "digital_publication_id"]);
```

#### Example: List

```php
// list() returns an array of DigitalPublication records (throws on error).
$digital_publications = $client->DigitalPublication()->list();
```


### DigitalPublicationArticle

Create an instance: `$digital_publication_article = $client->DigitalPublicationArticle();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `author_display` | `mixed` |  |
| `copy` | `mixed` |  |
| `digital_publication_id` | `string` |  |
| `id` | `string` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |
| `web_url` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare DigitalPublicationArticle record (throws on error).
$digital_publication_article = $client->DigitalPublicationArticle()->load(["id" => "digital_publication_article_id"]);
```

#### Example: List

```php
// list() returns an array of DigitalPublicationArticle records (throws on error).
$digital_publication_articles = $client->DigitalPublicationArticle()->list();
```


### EducatorResource

Create an instance: `$educator_resource = $client->EducatorResource();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `copy` | `mixed` |  |
| `id` | `string` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |
| `web_url` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare EducatorResource record (throws on error).
$educator_resource = $client->EducatorResource()->load(["id" => "educator_resource_id"]);
```

#### Example: List

```php
// list() returns an array of EducatorResource records (throws on error).
$educator_resources = $client->EducatorResource()->list();
```


### Event

Create an instance: `$event = $client->Event();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_audience_id` | `string` |  |
| `alt_event_type_id` | `string` |  |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `audience_id` | `string` |  |
| `buy_button_caption` | `mixed` |  |
| `buy_button_text` | `mixed` |  |
| `date_display` | `mixed` |  |
| `description` | `string` |  |
| `door_time` | `mixed` |  |
| `end_date` | `mixed` |  |
| `end_time` | `mixed` |  |
| `entrance` | `mixed` |  |
| `event_host_id` | `string` |  |
| `event_host_title` | `mixed` |  |
| `event_type_id` | `string` |  |
| `header_description` | `mixed` |  |
| `hero_caption` | `mixed` |  |
| `id` | `string` |  |
| `image_url` | `mixed` |  |
| `is_admission_required` | `bool` |  |
| `is_after_hour` | `bool` |  |
| `is_free` | `bool` |  |
| `is_member_exclusive` | `bool` |  |
| `is_private` | `bool` |  |
| `is_registration_required` | `bool` |  |
| `is_sales_button_hidden` | `bool` |  |
| `is_sold_out` | `bool` |  |
| `is_ticketed` | `bool` |  |
| `is_virtual_event` | `bool` |  |
| `join_url` | `mixed` |  |
| `layout_type` | `mixed` |  |
| `list_description` | `mixed` |  |
| `location` | `mixed` |  |
| `program_id` | `string` |  |
| `program_title` | `mixed` |  |
| `rsvp_link` | `mixed` |  |
| `search_tag` | `mixed` |  |
| `short_description` | `mixed` |  |
| `slug` | `string` |  |
| `source_updated_at` | `mixed` |  |
| `start_date` | `mixed` |  |
| `start_time` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `survey_url` | `mixed` |  |
| `ticketed_event_id` | `string` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `title_display` | `mixed` |  |
| `updated_at` | `mixed` |  |
| `virtual_event_passcode` | `mixed` |  |
| `virtual_event_url` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare Event record (throws on error).
$event = $client->Event()->load(["id" => "event_id"]);
```

#### Example: List

```php
// list() returns an array of Event records (throws on error).
$events = $client->Event()->list();
```


### EventOccurrence

Create an instance: `$event_occurrence = $client->EventOccurrence();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `button_caption` | `mixed` |  |
| `button_text` | `mixed` |  |
| `button_url` | `mixed` |  |
| `description` | `string` |  |
| `end_at` | `mixed` |  |
| `event_id` | `string` |  |
| `id` | `string` |  |
| `image_url` | `mixed` |  |
| `is_private` | `bool` |  |
| `is_sales_button_hidden` | `bool` |  |
| `is_ticketed` | `bool` |  |
| `location` | `mixed` |  |
| `off_sale_at` | `mixed` |  |
| `on_sale_at` | `mixed` |  |
| `short_description` | `mixed` |  |
| `source_updated_at` | `mixed` |  |
| `start_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `title_display` | `mixed` |  |
| `updated_at` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare EventOccurrence record (throws on error).
$event_occurrence = $client->EventOccurrence()->load(["id" => "event_occurrence_id"]);
```

#### Example: List

```php
// list() returns an array of EventOccurrence records (throws on error).
$event_occurrences = $client->EventOccurrence()->list();
```


### EventProgram

Create an instance: `$event_program = $client->EventProgram();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `id` | `string` |  |
| `is_affiliate_group` | `bool` |  |
| `is_event_host` | `bool` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare EventProgram record (throws on error).
$event_program = $client->EventProgram()->load(["id" => "event_program_id"]);
```

#### Example: List

```php
// list() returns an array of EventProgram records (throws on error).
$event_programs = $client->EventProgram()->list();
```


### Exhibition

Create an instance: `$exhibition = $client->Exhibition();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `aic_end_at` | `mixed` |  |
| `aic_start_at` | `mixed` |  |
| `alt_image_id` | `string` |  |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `artist_id` | `string` |  |
| `artwork_id` | `string` |  |
| `artwork_title` | `mixed` |  |
| `document_id` | `string` |  |
| `gallery_id` | `string` |  |
| `gallery_title` | `mixed` |  |
| `id` | `string` |  |
| `image_id` | `string` |  |
| `image_url` | `mixed` |  |
| `is_featured` | `bool` |  |
| `is_published` | `bool` |  |
| `position` | `mixed` |  |
| `short_description` | `mixed` |  |
| `site_id` | `string` |  |
| `source_updated_at` | `mixed` |  |
| `status` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |
| `web_url` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare Exhibition record (throws on error).
$exhibition = $client->Exhibition()->load(["id" => "exhibition_id"]);
```

#### Example: List

```php
// list() returns an array of Exhibition records (throws on error).
$exhibitions = $client->Exhibition()->list();
```


### Gallery

Create an instance: `$gallery = $client->Gallery();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `floor` | `mixed` |  |
| `id` | `string` |  |
| `is_closed` | `bool` |  |
| `latitude` | `float` |  |
| `latlon` | `mixed` |  |
| `longitude` | `float` |  |
| `number` | `mixed` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `tgn_id` | `string` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare Gallery record (throws on error).
$gallery = $client->Gallery()->load(["id" => "gallery_id"]);
```

#### Example: List

```php
// list() returns an array of Gallery records (throws on error).
$gallerys = $client->Gallery()->list();
```


### GenericPage

Create an instance: `$generic_page = $client->GenericPage();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `copy` | `mixed` |  |
| `id` | `string` |  |
| `search_tag` | `mixed` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |
| `web_url` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare GenericPage record (throws on error).
$generic_page = $client->GenericPage()->load(["id" => "generic_page_id"]);
```

#### Example: List

```php
// list() returns an array of GenericPage records (throws on error).
$generic_pages = $client->GenericPage()->list();
```


### Highlight

Create an instance: `$highlight = $client->Highlight();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `copy` | `mixed` |  |
| `id` | `string` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare Highlight record (throws on error).
$highlight = $client->Highlight()->load(["id" => "highlight_id"]);
```

#### Example: List

```php
// list() returns an array of Highlight records (throws on error).
$highlights = $client->Highlight()->list();
```


### Hour

Create an instance: `$hour = $client->Hour();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `additional_text` | `mixed` |  |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `friday_is_closed` | `mixed` |  |
| `friday_member_close` | `mixed` |  |
| `friday_member_open` | `mixed` |  |
| `friday_public_close` | `mixed` |  |
| `friday_public_open` | `mixed` |  |
| `id` | `string` |  |
| `monday_is_closed` | `mixed` |  |
| `monday_member_close` | `mixed` |  |
| `monday_member_open` | `mixed` |  |
| `monday_public_close` | `mixed` |  |
| `monday_public_open` | `mixed` |  |
| `saturday_is_closed` | `mixed` |  |
| `saturday_member_close` | `mixed` |  |
| `saturday_member_open` | `mixed` |  |
| `saturday_public_close` | `mixed` |  |
| `saturday_public_open` | `mixed` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `summary` | `mixed` |  |
| `sunday_is_closed` | `mixed` |  |
| `sunday_member_close` | `mixed` |  |
| `sunday_member_open` | `mixed` |  |
| `sunday_public_close` | `mixed` |  |
| `sunday_public_open` | `mixed` |  |
| `thursday_is_closed` | `mixed` |  |
| `thursday_member_close` | `mixed` |  |
| `thursday_member_open` | `mixed` |  |
| `thursday_public_close` | `mixed` |  |
| `thursday_public_open` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `tuesday_is_closed` | `mixed` |  |
| `tuesday_member_close` | `mixed` |  |
| `tuesday_member_open` | `mixed` |  |
| `tuesday_public_close` | `mixed` |  |
| `tuesday_public_open` | `mixed` |  |
| `updated_at` | `mixed` |  |
| `wednesday_is_closed` | `mixed` |  |
| `wednesday_member_close` | `mixed` |  |
| `wednesday_member_open` | `mixed` |  |
| `wednesday_public_close` | `mixed` |  |
| `wednesday_public_open` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare Hour record (throws on error).
$hour = $client->Hour()->load(["id" => "hour_id"]);
```

#### Example: List

```php
// list() returns an array of Hour records (throws on error).
$hours = $client->Hour()->list();
```


### Image

Create an instance: `$image = $client->Image();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ahash` | `mixed` |  |
| `alt_text` | `mixed` |  |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `artwork_id` | `string` |  |
| `artwork_title` | `mixed` |  |
| `color` | `mixed` |  |
| `colorfulness` | `mixed` |  |
| `content` | `mixed` |  |
| `content_e_tag` | `mixed` |  |
| `credit_line` | `mixed` |  |
| `fingerprint` | `mixed` |  |
| `height` | `float` |  |
| `id` | `string` |  |
| `iiif_url` | `mixed` |  |
| `is_educational_resource` | `bool` |  |
| `is_multimedia_resource` | `bool` |  |
| `is_teacher_resource` | `bool` |  |
| `lake_guid` | `mixed` |  |
| `lqip` | `mixed` |  |
| `phash` | `mixed` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `type` | `mixed` |  |
| `updated_at` | `mixed` |  |
| `width` | `float` |  |

#### Example: Load

```php
// load() returns the bare Image record (throws on error).
$image = $client->Image()->load(["id" => "image_id"]);
```

#### Example: List

```php
// list() returns an array of Image records (throws on error).
$images = $client->Image()->list();
```


### LandingPage

Create an instance: `$landing_page = $client->LandingPage();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `copy` | `mixed` |  |
| `id` | `string` |  |
| `search_tag` | `mixed` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |
| `web_url` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare LandingPage record (throws on error).
$landing_page = $client->LandingPage()->load(["id" => "landing_page_id"]);
```

#### Example: List

```php
// list() returns an array of LandingPage records (throws on error).
$landing_pages = $client->LandingPage()->list();
```


### Place

Create an instance: `$place = $client->Place();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `id` | `string` |  |
| `latitude` | `float` |  |
| `longitude` | `float` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `tgn_id` | `string` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare Place record (throws on error).
$place = $client->Place()->load(["id" => "place_id"]);
```

#### Example: List

```php
// list() returns an array of Place records (throws on error).
$places = $client->Place()->list();
```


### PressRelease

Create an instance: `$press_release = $client->PressRelease();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `copy` | `mixed` |  |
| `id` | `string` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |
| `web_url` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare PressRelease record (throws on error).
$press_release = $client->PressRelease()->load(["id" => "press_release_id"]);
```

#### Example: List

```php
// list() returns an array of PressRelease records (throws on error).
$press_releases = $client->PressRelease()->list();
```


### PrintedPublication

Create an instance: `$printed_publication = $client->PrintedPublication();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `copy` | `mixed` |  |
| `id` | `string` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |
| `web_url` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare PrintedPublication record (throws on error).
$printed_publication = $client->PrintedPublication()->load(["id" => "printed_publication_id"]);
```

#### Example: List

```php
// list() returns an array of PrintedPublication records (throws on error).
$printed_publications = $client->PrintedPublication()->list();
```


### Product

Create an instance: `$product = $client->Product();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `artist_id` | `string` |  |
| `artwork_id` | `string` |  |
| `description` | `string` |  |
| `exhibition_id` | `string` |  |
| `external_sku` | `mixed` |  |
| `id` | `string` |  |
| `image_url` | `mixed` |  |
| `max_compare_at_price` | `mixed` |  |
| `max_current_price` | `mixed` |  |
| `min_compare_at_price` | `mixed` |  |
| `min_current_price` | `mixed` |  |
| `price_display` | `mixed` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |
| `web_url` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare Product record (throws on error).
$product = $client->Product()->load(["id" => "product_id"]);
```

#### Example: List

```php
// list() returns an array of Product records (throws on error).
$products = $client->Product()->list();
```


### Publication

Create an instance: `$publication = $client->Publication();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `id` | `string` |  |
| `section_id` | `string` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |
| `web_url` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare Publication record (throws on error).
$publication = $client->Publication()->load(["id" => "publication_id"]);
```

#### Example: List

```php
// list() returns an array of Publication records (throws on error).
$publications = $client->Publication()->list();
```


### Search

Create an instance: `$search = $client->Search();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_id` | `string` |  |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `id` | `string` |  |
| `is_boosted` | `bool` |  |
| `score` | `float` |  |
| `thumbnail` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |

#### Example: List

```php
// list() returns an array of Search records (throws on error).
$searchs = $client->Search()->list();
```


### Section

Create an instance: `$section = $client->Section();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accession` | `mixed` |  |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `artwork_id` | `string` |  |
| `content` | `mixed` |  |
| `generic_page_id` | `string` |  |
| `id` | `string` |  |
| `publication_id` | `string` |  |
| `publication_title` | `mixed` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |
| `web_url` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare Section record (throws on error).
$section = $client->Section()->load(["id" => "section_id"]);
```

#### Example: List

```php
// list() returns an array of Section records (throws on error).
$sections = $client->Section()->list();
```


### Site

Create an instance: `$site = $client->Site();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `artwork_id` | `string` |  |
| `artwork_title` | `mixed` |  |
| `description` | `string` |  |
| `exhibition_id` | `string` |  |
| `exhibition_title` | `mixed` |  |
| `id` | `string` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |
| `web_url` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare Site record (throws on error).
$site = $client->Site()->load(["id" => "site_id"]);
```

#### Example: List

```php
// list() returns an array of Site records (throws on error).
$sites = $client->Site()->list();
```


### Sound

Create an instance: `$sound = $client->Sound();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_text` | `mixed` |  |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `artwork_id` | `string` |  |
| `artwork_title` | `mixed` |  |
| `content` | `mixed` |  |
| `content_e_tag` | `mixed` |  |
| `credit_line` | `mixed` |  |
| `id` | `string` |  |
| `is_educational_resource` | `bool` |  |
| `is_multimedia_resource` | `bool` |  |
| `is_teacher_resource` | `bool` |  |
| `lake_guid` | `mixed` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `transcript` | `mixed` |  |
| `type` | `mixed` |  |
| `updated_at` | `mixed` |  |
| `web_url` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare Sound record (throws on error).
$sound = $client->Sound()->load(["id" => "sound_id"]);
```

#### Example: List

```php
// list() returns an array of Sound records (throws on error).
$sounds = $client->Sound()->list();
```


### StaticPage

Create an instance: `$static_page = $client->StaticPage();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `id` | `string` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |
| `web_url` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare StaticPage record (throws on error).
$static_page = $client->StaticPage()->load(["id" => "static_page_id"]);
```

#### Example: List

```php
// list() returns an array of StaticPage records (throws on error).
$static_pages = $client->StaticPage()->list();
```


### Text

Create an instance: `$text = $client->Text();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_text` | `mixed` |  |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `artwork_id` | `string` |  |
| `artwork_title` | `mixed` |  |
| `content` | `mixed` |  |
| `content_e_tag` | `mixed` |  |
| `credit_line` | `mixed` |  |
| `id` | `string` |  |
| `is_educational_resource` | `bool` |  |
| `is_multimedia_resource` | `bool` |  |
| `is_teacher_resource` | `bool` |  |
| `lake_guid` | `mixed` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `type` | `mixed` |  |
| `updated_at` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare Text record (throws on error).
$text = $client->Text()->load(["id" => "text_id"]);
```

#### Example: List

```php
// list() returns an array of Text records (throws on error).
$texts = $client->Text()->list();
```


### Tour

Create an instance: `$tour = $client->Tour();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `artist_title` | `mixed` |  |
| `artwork_title` | `mixed` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `image` | `mixed` |  |
| `intro` | `mixed` |  |
| `intro_link` | `mixed` |  |
| `intro_transcript` | `mixed` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `updated_at` | `mixed` |  |
| `weight` | `float` |  |

#### Example: Load

```php
// load() returns the bare Tour record (throws on error).
$tour = $client->Tour()->load(["id" => "tour_id"]);
```

#### Example: List

```php
// list() returns an array of Tour records (throws on error).
$tours = $client->Tour()->list();
```


### Video

Create an instance: `$video = $client->Video();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_text` | `mixed` |  |
| `api_link` | `mixed` |  |
| `api_model` | `mixed` |  |
| `artwork_id` | `string` |  |
| `artwork_title` | `mixed` |  |
| `content` | `mixed` |  |
| `content_e_tag` | `mixed` |  |
| `credit_line` | `mixed` |  |
| `id` | `string` |  |
| `is_educational_resource` | `bool` |  |
| `is_multimedia_resource` | `bool` |  |
| `is_teacher_resource` | `bool` |  |
| `lake_guid` | `mixed` |  |
| `source_updated_at` | `mixed` |  |
| `suggest_autocomplete_all` | `mixed` |  |
| `suggest_autocomplete_boosted` | `mixed` |  |
| `timestamp` | `mixed` |  |
| `title` | `string` |  |
| `type` | `mixed` |  |
| `updated_at` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare Video record (throws on error).
$video = $client->Video()->load(["id" => "video_id"]);
```

#### Example: List

```php
// list() returns an array of Video records (throws on error).
$videos = $client->Video()->list();
```


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is a PHP class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as arrays

The PHP SDK uses plain PHP associative arrays throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers::to_map()` to safely validate that a value is an array.

### Directory structure

```
php/
├── artinstituteofchicago_sdk.php          -- Main SDK class
├── config.php                     -- Configuration
├── features.php                   -- Feature factory
├── core/                          -- Core types and context
├── entity/                        -- Entity implementations
├── feature/                       -- Built-in features (Base, Test, Log)
├── utility/                       -- Utility functions and struct library
└── test/                          -- Test suites
```

The main class (`artinstituteofchicago_sdk.php`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```php
$agent = $client->Agent();
$agent->list();

// $agent->data_get() now returns the agent data from the last list
// $agent->match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
