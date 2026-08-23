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
        echo $item["id"] . " " . $item["alt_titles"] . "\n";
    }
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

### 3. Load an agent

```php
try {
    // load() returns the ENTITY — call data_get() for the Agent record (throws on error).
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
    $publications = $client->Publication()->list();
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
    "entity" => ["publication" => ["test01" => ["id" => "test01"]]],
]);

// Entity ops return the ENTITY (throws on error);
// call data_get() for the mock record.
$publication = $client->Publication()->list();
print_r($publication);
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

Entity operations return the ENTITY (call data_get() for the record) (an `array` for single-entity
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
| `alt_titles` | Alternate names for this agent |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `birth_date` | The year this agent was born |
| `death_date` | The year this agent died |
| `description` | A biographical description of the agent |
| `id` | Unique identifier of this resource. |
| `is_artist` | Whether the agent is an artist. |
| `sort_title` | Sortable name for this agent, typically with last name first. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `ulan_id` | Unique identifier of this agent in Getty's ULAN |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/agents`

#### AgentRole

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/agent-roles`

#### AgentType

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/agent-types`

#### Article

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `copy` | The text of the article |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/articles`

#### Artwork

| Field | Description |
| --- | --- |
| `alt_artist_ids` | Unique identifiers of the non-preferred artists/cultures associated with this work |
| `alt_classification_ids` | Unique identifiers of all other non-preferred classification terms for this work |
| `alt_image_ids` | Unique identifiers of all non-preferred images of this work. |
| `alt_material_ids` | Unique identifiers of all other non-preferred material terms for this work |
| `alt_style_ids` | Unique identifiers of all other non-preferred style terms for this work |
| `alt_subject_ids` | Unique identifiers of all other non-preferred subject terms for this work |
| `alt_technique_ids` | Unique identifiers of all other non-preferred technique terms for this work |
| `alt_titles` | Alternate names for this work |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `artist_display` | Readable description of the creator of this work. |
| `artist_id` | Unique identifier of the preferred artist/culture associated with this work |
| `artist_ids` | Unique identifier of all artist/cultures associated with this work |
| `artist_title` | Name of the preferred artist/culture associated with this work |
| `artist_titles` | Names of all artist/cultures associated with this work |
| `artwork_type_id` | Unique identifier of the kind of object or work |
| `artwork_type_title` | The kind of object or work (e.g. |
| `boost_rank` | Manual indication of what rank this artwork should take in search results. |
| `catalog_based_search_keyword_titles` | The keyword search values that would be catalog-based searches on this record |
| `catalogue_display` | Brief text listing all the catalogues raisonnés which include this work. |
| `category_ids` | Unique identifiers of the categories this work is a part of |
| `category_titles` | Names of the categories this artwork is a part of |
| `classification_id` | Unique identifier of the preferred classification term for this work |
| `classification_ids` | Unique identifiers of all classification terms for this work |
| `classification_title` | The name of the preferred classification term for this work |
| `classification_titles` | The names of all classification terms related to this artwork |
| `color` | Dominant color of this artwork in HSL |
| `colorfulness` | Unbounded positive float representing an abstract measure of colorfulness. |
| `copyright_notice` | Statement notifying how the work is protected by copyright. |
| `credit_line` | Brief statement indicating how the work came into the collection |
| `date_display` | Readable, free-text description of the period of time associated with the creation of this work. |
| `date_end` | The year of the period of time associated with the creation of this work |
| `date_qualifier_id` | Unique identifier of the qualifer to the dates provided for this record. |
| `date_qualifier_title` | Readable, text qualifer to the dates provided for this record. |
| `date_start` | The year of the period of time associated with the creation of this work |
| `department_id` | Unique identifier of the curatorial department that this work belongs to |
| `department_title` | Name of the curatorial department that this work belongs to |
| `description` | Longer explanation describing the work |
| `dimensions` | The size, shape, scale, and dimensions of the work. |
| `dimensions_detail` | The height, width, depth, and/or diameter of each section of the work in centimeters |
| `document_ids` | Unique identifiers of assets that serve as documentation for this artwork |
| `edition` | Edition number if the work is one of many |
| `exhibition_history` | List of all the places this work has been exhibited |
| `fiscal_year` | The fiscal year in which the work was acquired. |
| `fiscal_year_deaccession` | The fiscal year in which the work was deaccessioned. |
| `gallery_id` | Unique identifier of the location of this work in our museum |
| `gallery_title` | The location of this work in our museum |
| `has_advanced_imaging` | Whether this artwork is enhanced with 3D models, 360 image sequences, Mirador views, etc. |
| `has_educational_resources` | Whether this artwork has any documents tagged as educational |
| `has_multimedia_resources` | Whether this artwork has any associated microsites, digital publications, or documents tagged as multimedia |
| `has_not_been_viewed_much` | Whether the artwork hasn't been visited on our website very much |
| `id` | Unique identifier of this resource. |
| `image_embedding` | The generated embeddings describing the artwork image |
| `image_id` | Unique identifier of the preferred image to use to represent this work |
| `inscriptions` | A description of distinguishing or identifying physical markings that are on the work |
| `internal_department_id` | An internal department id we use for analytics. |
| `is_boosted` | Whether this document should be boosted in search |
| `is_on_view` | Whether the work is on display |
| `is_public_domain` | Whether the work is in the public domain, meaning it was created before copyrights existed or has left the copyright term |
| `is_zoomable` | Whether images of the work are allowed to be displayed in a zoomable interface. |
| `latitude` | Latitude coordinate of the location of this work in our galleries |
| `latlon` | Latitude and longitude coordinates of the location of this work in our galleries |
| `longitude` | Longitude coordinate of the location of this work in our galleries |
| `main_reference_number` | Unique identifier assigned to the artwork upon acquisition |
| `material_id` | Unique identifier of the preferred material term for this work |
| `material_ids` | Unique identifiers of all material terms for this work |
| `material_titles` | The names of all material terms related to this artwork |
| `max_zoom_window_size` | The maximum size of the window the image is allowed to be viewed in, in pixels. |
| `medium_display` | The substances or materials used in the creation of a work |
| `nomisma_id` | Unique identifier of this work in the nomisma coin database |
| `on_loan_display` | If an artwork is on loan, this contains details about the loan |
| `pageviews` | Approx. |
| `pageviews_recent` | Approx. |
| `place_of_origin` | The location where the creation, design, or production of the work took place, or the original location of the work |
| `provenance_text` | Ownership/collecting history of the work. |
| `publication_history` | Bibliographic list of all the places this work has been published |
| `publishing_verification_level` | Indicator of how much metadata on the work in published. |
| `section_ids` | Unique identifiers of the digital publication chapters this work in included in |
| `section_titles` | Names of the digital publication chapters this work is included in |
| `short_description` | Short explanation describing the work |
| `site_ids` | Unique identifiers of the microsites this work is a part of |
| `sound_ids` | Unique identifiers of the audio about this work |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `style_id` | Unique identifier of the preferred style term for this work |
| `style_ids` | Unique identifiers of all style terms for this work |
| `style_title` | The name of the preferred style term for this work |
| `style_titles` | The names of all style terms related to this artwork |
| `subject_id` | Unique identifier of the preferred subject term for this work |
| `subject_ids` | Unique identifiers of all subject terms for this work |
| `subject_titles` | The names of all subject terms related to this artwork |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `technique_id` | Unique identifier of the preferred technique term for this work |
| `technique_ids` | Unique identifiers of all technique terms for this work |
| `technique_titles` | The names of all technique terms related to this artwork |
| `term_titles` | The names of the taxonomy tags for this work |
| `text_embedding` | The generated embeddings of artwork text |
| `text_ids` | Unique identifiers of the texts about this work |
| `theme_titles` | The names of all thematic publish categories related to this artwork |
| `thumbnail` | Metadata about the image referenced by `image_id`. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `video_ids` | Unique identifiers of the videos about this work |

Operations: List, Load.

API path: `/artworks`

#### ArtworkDateQualifier

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/artwork-date-qualifiers`

#### ArtworkPlaceQualifier

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/artwork-place-qualifiers`

#### ArtworkType

| Field | Description |
| --- | --- |
| `aat_id` | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/artwork-types`

#### CategoryTerm

| Field | Description |
| --- | --- |
| `aat_id` | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `id` | Unique identifier of this resource. |
| `parent_id` | Unique identifier of this category's parent |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `subtype` | Takes one of the following values: classification, material, technique, style, subject, department, theme |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/category-terms`

#### DigitalPublication

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `copy` | The text of the page |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | The URL to this page on our website |

Operations: List, Load.

API path: `/digital-publications`

#### DigitalPublicationArticle

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `author_display` | A display-friendly text of the authors of this article |
| `copy` | The text of the article |
| `digital_publication_id` | Unique identifier of the digital publication this article belongs to |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | The URL to this article on our website |

Operations: List, Load.

API path: `/digital-publication-articles`

#### EducatorResource

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `copy` | The text of the page |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | The URL to this page on our website |

Operations: List, Load.

API path: `/educator-resources`

#### Event

| Field | Description |
| --- | --- |
| `alt_audience_ids` | Unique identifiers indicating the alternate audiences for this event |
| `alt_event_type_ids` | Unique identifiers indicating the alternate types of this event |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `audience_id` | Unique identifier indicating the preferred audience for this event |
| `buy_button_caption` | Additional text below the ticket/registration button |
| `buy_button_text` | The text used on the ticket/registration button |
| `date_display` | A readable display of the event dates |
| `description` | All copytext of the event |
| `door_time` | The time the doors open for this event |
| `end_date` | The date the event ends |
| `end_time` | The time the event ends |
| `entrance` | Which entrance to use for this event |
| `event_host_id` | Unique identifier of the host (cf. |
| `event_host_title` | Unique identifier of the host (cf. |
| `event_type_id` | Unique identifier indicating the preferred type of this event |
| `header_description` | Brief description of the event displayed below the title |
| `hero_caption` | Text displayed with the hero image on the event |
| `id` | Unique identifier of this resource. |
| `image_url` | The URL of an image representing this page |
| `is_admission_required` | Whether admission to the museum is required to attend this event |
| `is_after_hours` | Whether the event is to be held after the museum closes |
| `is_free` | Whether the event is free |
| `is_member_exclusive` | Whether the event is exclusive to members of the museum |
| `is_private` | Whether the event is private |
| `is_registration_required` | Whether registration is required to attend the event |
| `is_sales_button_hidden` | Whether the buy tickets button is hidden on the website event page |
| `is_sold_out` | Whether the event is sold out |
| `is_ticketed` | Whether a ticket is required to attend the event |
| `is_virtual_event` | Whether the event is being held virtually |
| `join_url` | URL to the membership signup page via this event |
| `layout_type` | Number indicating the type of layout this event page uses |
| `list_description` | One-sentence description of the event displayed in listings |
| `location` | Where the event takes place |
| `program_ids` | Unique identifiers indicating the programs this event is a part of |
| `program_titles` | Titles of the programs this event is a part of |
| `rsvp_link` | The URL to the sales site for this event |
| `search_tags` | Editor-specified list of tags to aid in internal search |
| `short_description` | Brief description of the event |
| `slug` | A string used in the URL for this event |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `start_date` | The date the event begins |
| `start_time` | The time the event starts |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `survey_url` | URL to the survey associated with this event |
| `ticketed_event_id` | Unique identifier of the event in the ticketing system this website event is tied to |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `title_display` | The name of this event formatted with HTML (optional) |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `virtual_event_passcode` | Passcode to access the virtual event |
| `virtual_event_url` | URL to the virtual event |

Operations: List, Load.

API path: `/events`

#### EventOccurrence

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `button_caption` | Additional text below the ticket/registration button |
| `button_text` | The text used on the ticket/registration button |
| `button_url` | The URL to the sales site or an RSVP link for this event |
| `description` | Description of the event |
| `end_at` | The date the event occurrence ends |
| `event_id` | Identifier of the master event of which this is an occurrence |
| `id` | Unique identifier of this resource. |
| `image_url` | The URL of an image representing this page |
| `is_private` | Whether the event is private. |
| `is_sales_button_hidden` | Whether the buy tickets button is hidden on the website event page |
| `is_ticketed` | Whether a ticket is required to attend the event |
| `location` | Where the event takes place |
| `off_sale_at` | Date and time the event goes off sale |
| `on_sale_at` | Date and time the event goes on sale |
| `short_description` | Brief description of the event |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `start_at` | The date the event occurrence begins |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `title_display` | The name of this event formatted with HTML (optional) |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/event-occurrences`

#### EventProgram

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `id` | Unique identifier of this resource. |
| `is_affiliate_group` | Whether this program represents an affiliate group |
| `is_event_host` | Whether this program represents an event host |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/event-programs`

#### Exhibition

| Field | Description |
| --- | --- |
| `aic_end_at` | Date the exhibition closed at the Art Institute of Chicago |
| `aic_start_at` | Date the exhibition opened at the Art Institute of Chicago |
| `alt_image_ids` | Unique identifiers of all non-preferred images of this exhibition. |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `artist_ids` | Unique identifiers of the artist agent records representing who was shown in the exhibition |
| `artwork_ids` | Unique identifiers of the artworks that were part of the exhibition |
| `artwork_titles` | Names of the artworks that were part of the exhibition |
| `document_ids` | Unique identifiers of assets that serve as documentation for this exhibition |
| `gallery_id` | Unique identifier of the gallery that mainly housed the exhibition |
| `gallery_title` | The name of the gallery that mainly housed the exhibition |
| `id` | Unique identifier of this resource. |
| `image_id` | Unique identifier of the preferred image to use to represent this exhibition |
| `image_url` | URL to the hero image from the website |
| `is_featured` | Is this exhibition currently featured on our website? |
| `is_published` | Is this exhibition currently published on our website? |
| `position` | Numering position represnting the order in which this exhibition is featured on the website |
| `short_description` | Brief explanation of what this exhibition is |
| `site_ids` | Unique identifiers of the microsites this exhibition is a part of |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `status` | Whether the exhibition is open or closed |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | URL to this exhibition on our website |

Operations: List, Load.

API path: `/exhibitions`

#### Gallery

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `floor` | The level the gallery is on, e.g., 1, 2, 3, or LL |
| `id` | Unique identifier of this resource. |
| `is_closed` | Whether the gallery is currently closed |
| `latitude` | Latitude coordinate of the center of the room |
| `latlon` | Latitude and longitude coordinates of the center of the room |
| `longitude` | Longitude coordinate of the center of the room |
| `number` | The gallery's room number. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/galleries`

#### GenericPage

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `copy` | The text of the page |
| `id` | Unique identifier of this resource. |
| `search_tags` | Editor-specified list of tags to aid in internal search |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | The URL to this page on our website |

Operations: List, Load.

API path: `/generic-pages`

#### Highlight

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `copy` | The text of the highlight description |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/highlights`

#### Hour

| Field | Description |
| --- | --- |
| `additional_text` | Additional information about the hours |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `friday_is_closed` | Whether the museum is closed on Fridays |
| `friday_member_close` | The time member hours ends on Fridays |
| `friday_member_open` | The time member hours starts on Fridays |
| `friday_public_close` | The time public hours ends on Fridays |
| `friday_public_open` | The time public hours starts on Fridays |
| `id` | Unique identifier of this resource. |
| `monday_is_closed` | Whether the museum is closed on Mondays |
| `monday_member_close` | The time member hours ends on Mondays |
| `monday_member_open` | The time member hours starts on Mondays |
| `monday_public_close` | The time public hours ends on Mondays |
| `monday_public_open` | The time public hours starts on Mondays |
| `saturday_is_closed` | Whether the museum is closed on Saturdays |
| `saturday_member_close` | The time member hours ends on Saturdays |
| `saturday_member_open` | The time member hours starts on Saturdays |
| `saturday_public_close` | The time public hours ends on Saturdays |
| `saturday_public_open` | The time public hours starts on Saturdays |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `summary` | Readable summary of the hours |
| `sunday_is_closed` | Whether the museum is closed on Sundays |
| `sunday_member_close` | The time member hours ends on Sundays |
| `sunday_member_open` | The time member hours starts on Sundays |
| `sunday_public_close` | The time public hours ends on Sundays |
| `sunday_public_open` | The time public hours starts on Sundays |
| `thursday_is_closed` | Whether the museum is closed on Thursdays |
| `thursday_member_close` | The time member hours ends on Thursdays |
| `thursday_member_open` | The time member hours starts on Thursdays |
| `thursday_public_close` | The time public hours ends on Thursdays |
| `thursday_public_open` | The time public hours starts on Thursdays |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `tuesday_is_closed` | Whether the museum is closed on Tuesdays |
| `tuesday_member_close` | The time member hours ends on Tuesdays |
| `tuesday_member_open` | The time member hours starts on Tuesdays |
| `tuesday_public_close` | The time public hours ends on Tuesdays |
| `tuesday_public_open` | The time public hours starts on Tuesdays |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `wednesday_is_closed` | Whether the museum is closed on Wednesdays |
| `wednesday_member_close` | The time member hours ends on Wednesdays |
| `wednesday_member_open` | The time member hours starts on Wednesdays |
| `wednesday_public_close` | The time public hours ends on Wednesdays |
| `wednesday_public_open` | The time public hours starts on Wednesdays |

Operations: List, Load.

API path: `/hours`

#### Image

| Field | Description |
| --- | --- |
| `ahash` | Image hash generated using ahash algorithm with 64 boolean subfields |
| `alt_text` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `artwork_ids` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | Names of the artworks associated with this asset |
| `color` | Dominant color of this image in HSL |
| `colorfulness` | Unbounded positive float representing an abstract measure of colorfulness. |
| `content` | Text of or URL to the contents of this asset |
| `content_e_tag` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | Asset-specific copyright information |
| `fingerprint` | Image hashes: aHash, dHash, pHash, wHash |
| `height` | Native height of the image |
| `id` | Unique identifier of this resource. |
| `iiif_url` | IIIF URL of this image |
| `is_educational_resource` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | Whether this resource is considered to be educational |
| `lake_guid` | Unique UUID of this resource in LAKE, our DAMS. |
| `lqip` | Low-quality image placeholder (LQIP). |
| `phash` | Image hash generated using phash algorithm with 64 boolean subfields |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `type` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `width` | Native width of the image |

Operations: List, Load.

API path: `/images`

#### LandingPage

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `copy` | The text of the page |
| `id` | Unique identifier of this resource. |
| `search_tags` | Editor-specified list of tags to aid in internal search |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | The URL to this page on our website |

Operations: List, Load.

API path: `/landing-pages`

#### Place

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `id` | Unique identifier of this resource. |
| `latitude` | Latitude coordinate of the center of the room |
| `longitude` | Longitude coordinate of the center of the room |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/places`

#### PressRelease

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `copy` | The text of the page |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | The URL to this page on our website |

Operations: List, Load.

API path: `/press-releases`

#### PrintedPublication

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `copy` | The text of the page |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | The URL to this page on our website |

Operations: List, Load.

API path: `/printed-publications`

#### Product

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `artist_ids` | Unique identifiers of the artists associated with this product |
| `artwork_ids` | Unique identifiers of the artworks associated with this product |
| `description` | Explanation of what this product is |
| `exhibition_ids` | Unique identifiers of the exhibitions associated with this product |
| `external_sku` | Numeric product identification code of a machine-readable barcode, when the customer sku differs from our internal one |
| `id` | Unique identifier of this resource. |
| `image_url` | URL of an image for this product |
| `max_compare_at_price` | Number indicating how much the most expensive variant of a product cost before a sale |
| `max_current_price` | Number indicating how much the most expensive variant of a product costs right now |
| `min_compare_at_price` | Number indicating how much the least expensive variant of a product cost before a sale |
| `min_current_price` | Number indicating how much the least expensive variant of a product costs right now |
| `price_display` | Explanation of what this product is |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | URL of this product in the shop |

Operations: List, Load.

API path: `/products`

#### Publication

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `id` | Unique identifier of this resource. |
| `section_ids` | Unique identifiers of the sections of this publication |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | URL to the publication |

Operations: List, Load.

API path: `/publications`

#### Search

| Field | Description |
| --- | --- |
| `api_id` | API unique identifier |
| `api_link` | URL to this recource in the API |
| `api_model` | Name of the model the resource represents |
| `id` | Unique identifier within the search index |
| `is_boosted` | Whether this record has been flagged to be boosted |
| `score` | Search index ranking of the result |
| `thumbnail` | Metadata on the image representing this record |
| `timestamp` | Date this record was last updated in the API |
| `title` | The name of this resource |

Operations: List.

API path: `/agents/search`

#### Section

| Field | Description |
| --- | --- |
| `accession` | An accession number parsed from the title or tombstone |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `artwork_id` | Unique identifier of the artwork with which this section is associated |
| `content` | Content of this section in plaintext |
| `generic_page_id` | Unique identifier of the page on the website that represents the publication this section belongs to |
| `id` | Unique identifier of this resource. |
| `publication_id` | Unique identifier of the publication this section belongs to |
| `publication_title` | Name of the publication this section belongs to |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | URL to the section |

Operations: List, Load.

API path: `/sections`

#### Site

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `artwork_ids` | Unique identifiers of the artworks this site is associated with |
| `artwork_titles` | Names of the artworks this site is associated with |
| `description` | Explanation of what this site is |
| `exhibition_ids` | Unique identifier of the exhibitions this site is associated with |
| `exhibition_titles` | Names of the exhibitions this site is associated with |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | URL to this site |

Operations: List, Load.

API path: `/sites`

#### Sound

| Field | Description |
| --- | --- |
| `alt_text` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `artwork_ids` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | Names of the artworks associated with this asset |
| `content` | Text of or URL to the contents of this asset |
| `content_e_tag` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | Asset-specific copyright information |
| `id` | Unique identifier of this resource. |
| `is_educational_resource` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | Whether this resource is considered to be educational |
| `lake_guid` | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | Name of this mobile audio file – derived from the artwork and tour titles |
| `transcript` | Text transcription of the audio file |
| `type` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | URL to the audio file |

Operations: List, Load.

API path: `/mobile-sounds`

#### StaticPage

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | The URL to this page on our website |

Operations: List, Load.

API path: `/static-pages`

#### Text

| Field | Description |
| --- | --- |
| `alt_text` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `artwork_ids` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | Names of the artworks associated with this asset |
| `content` | Text of or URL to the contents of this asset |
| `content_e_tag` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | Asset-specific copyright information |
| `id` | Unique identifier of this resource. |
| `is_educational_resource` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | Whether this resource is considered to be educational |
| `lake_guid` | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `type` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/texts`

#### Tour

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `artist_titles` | Names of the artists of the artworks featured in this tour's tour stops |
| `artwork_titles` | Names of the artworks featured in this tour's tour stops |
| `description` | Explanation of what the tour is |
| `id` | Unique identifier of this resource. |
| `image` | The main image for the tour |
| `intro` | Text introducing the tour |
| `intro_link` | Link to the audio file of the introduction |
| `intro_transcript` | Transcript of the introduction audio to the tour |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `weight` | Number representing this tour's sort order |

Operations: List, Load.

API path: `/tours`

#### Video

| Field | Description |
| --- | --- |
| `alt_text` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `artwork_ids` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | Names of the artworks associated with this asset |
| `content` | Text of or URL to the contents of this asset |
| `content_e_tag` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | Asset-specific copyright information |
| `id` | Unique identifier of this resource. |
| `is_educational_resource` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | Whether this resource is considered to be educational |
| `lake_guid` | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `type` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | Date and time the record was updated in the aggregator database |

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
| `alt_titles` | `mixed` | Alternate names for this agent |
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `birth_date` | `mixed` | The year this agent was born |
| `death_date` | `mixed` | The year this agent died |
| `description` | `string` | A biographical description of the agent |
| `id` | `string` | Unique identifier of this resource. |
| `is_artist` | `bool` | Whether the agent is an artist. |
| `sort_title` | `mixed` | Sortable name for this agent, typically with last name first. |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `ulan_id` | `string` | Unique identifier of this agent in Getty's ULAN |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Agent record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the AgentRole record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the AgentType record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `copy` | `mixed` | The text of the article |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Article record (throws on error).
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
| `alt_artist_ids` | `mixed` | Unique identifiers of the non-preferred artists/cultures associated with this work |
| `alt_classification_ids` | `mixed` | Unique identifiers of all other non-preferred classification terms for this work |
| `alt_image_ids` | `mixed` | Unique identifiers of all non-preferred images of this work. |
| `alt_material_ids` | `mixed` | Unique identifiers of all other non-preferred material terms for this work |
| `alt_style_ids` | `mixed` | Unique identifiers of all other non-preferred style terms for this work |
| `alt_subject_ids` | `mixed` | Unique identifiers of all other non-preferred subject terms for this work |
| `alt_technique_ids` | `mixed` | Unique identifiers of all other non-preferred technique terms for this work |
| `alt_titles` | `mixed` | Alternate names for this work |
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `artist_display` | `mixed` | Readable description of the creator of this work. |
| `artist_id` | `string` | Unique identifier of the preferred artist/culture associated with this work |
| `artist_ids` | `mixed` | Unique identifier of all artist/cultures associated with this work |
| `artist_title` | `mixed` | Name of the preferred artist/culture associated with this work |
| `artist_titles` | `mixed` | Names of all artist/cultures associated with this work |
| `artwork_type_id` | `string` | Unique identifier of the kind of object or work |
| `artwork_type_title` | `mixed` | The kind of object or work (e.g. |
| `boost_rank` | `mixed` | Manual indication of what rank this artwork should take in search results. |
| `catalog_based_search_keyword_titles` | `mixed` | The keyword search values that would be catalog-based searches on this record |
| `catalogue_display` | `mixed` | Brief text listing all the catalogues raisonnés which include this work. |
| `category_ids` | `mixed` | Unique identifiers of the categories this work is a part of |
| `category_titles` | `mixed` | Names of the categories this artwork is a part of |
| `classification_id` | `string` | Unique identifier of the preferred classification term for this work |
| `classification_ids` | `mixed` | Unique identifiers of all classification terms for this work |
| `classification_title` | `mixed` | The name of the preferred classification term for this work |
| `classification_titles` | `mixed` | The names of all classification terms related to this artwork |
| `color` | `mixed` | Dominant color of this artwork in HSL |
| `colorfulness` | `mixed` | Unbounded positive float representing an abstract measure of colorfulness. |
| `copyright_notice` | `mixed` | Statement notifying how the work is protected by copyright. |
| `credit_line` | `mixed` | Brief statement indicating how the work came into the collection |
| `date_display` | `mixed` | Readable, free-text description of the period of time associated with the creation of this work. |
| `date_end` | `mixed` | The year of the period of time associated with the creation of this work |
| `date_qualifier_id` | `string` | Unique identifier of the qualifer to the dates provided for this record. |
| `date_qualifier_title` | `mixed` | Readable, text qualifer to the dates provided for this record. |
| `date_start` | `mixed` | The year of the period of time associated with the creation of this work |
| `department_id` | `string` | Unique identifier of the curatorial department that this work belongs to |
| `department_title` | `mixed` | Name of the curatorial department that this work belongs to |
| `description` | `string` | Longer explanation describing the work |
| `dimensions` | `mixed` | The size, shape, scale, and dimensions of the work. |
| `dimensions_detail` | `mixed` | The height, width, depth, and/or diameter of each section of the work in centimeters |
| `document_ids` | `mixed` | Unique identifiers of assets that serve as documentation for this artwork |
| `edition` | `mixed` | Edition number if the work is one of many |
| `exhibition_history` | `mixed` | List of all the places this work has been exhibited |
| `fiscal_year` | `mixed` | The fiscal year in which the work was acquired. |
| `fiscal_year_deaccession` | `mixed` | The fiscal year in which the work was deaccessioned. |
| `gallery_id` | `string` | Unique identifier of the location of this work in our museum |
| `gallery_title` | `mixed` | The location of this work in our museum |
| `has_advanced_imaging` | `bool` | Whether this artwork is enhanced with 3D models, 360 image sequences, Mirador views, etc. |
| `has_educational_resources` | `bool` | Whether this artwork has any documents tagged as educational |
| `has_multimedia_resources` | `bool` | Whether this artwork has any associated microsites, digital publications, or documents tagged as multimedia |
| `has_not_been_viewed_much` | `bool` | Whether the artwork hasn't been visited on our website very much |
| `id` | `string` | Unique identifier of this resource. |
| `image_embedding` | `mixed` | The generated embeddings describing the artwork image |
| `image_id` | `string` | Unique identifier of the preferred image to use to represent this work |
| `inscriptions` | `mixed` | A description of distinguishing or identifying physical markings that are on the work |
| `internal_department_id` | `string` | An internal department id we use for analytics. |
| `is_boosted` | `bool` | Whether this document should be boosted in search |
| `is_on_view` | `bool` | Whether the work is on display |
| `is_public_domain` | `bool` | Whether the work is in the public domain, meaning it was created before copyrights existed or has left the copyright term |
| `is_zoomable` | `bool` | Whether images of the work are allowed to be displayed in a zoomable interface. |
| `latitude` | `float` | Latitude coordinate of the location of this work in our galleries |
| `latlon` | `mixed` | Latitude and longitude coordinates of the location of this work in our galleries |
| `longitude` | `float` | Longitude coordinate of the location of this work in our galleries |
| `main_reference_number` | `int` | Unique identifier assigned to the artwork upon acquisition |
| `material_id` | `string` | Unique identifier of the preferred material term for this work |
| `material_ids` | `mixed` | Unique identifiers of all material terms for this work |
| `material_titles` | `mixed` | The names of all material terms related to this artwork |
| `max_zoom_window_size` | `mixed` | The maximum size of the window the image is allowed to be viewed in, in pixels. |
| `medium_display` | `mixed` | The substances or materials used in the creation of a work |
| `nomisma_id` | `string` | Unique identifier of this work in the nomisma coin database |
| `on_loan_display` | `mixed` | If an artwork is on loan, this contains details about the loan |
| `pageviews` | `mixed` | Approx. |
| `pageviews_recent` | `mixed` | Approx. |
| `place_of_origin` | `mixed` | The location where the creation, design, or production of the work took place, or the original location of the work |
| `provenance_text` | `mixed` | Ownership/collecting history of the work. |
| `publication_history` | `mixed` | Bibliographic list of all the places this work has been published |
| `publishing_verification_level` | `mixed` | Indicator of how much metadata on the work in published. |
| `section_ids` | `mixed` | Unique identifiers of the digital publication chapters this work in included in |
| `section_titles` | `mixed` | Names of the digital publication chapters this work is included in |
| `short_description` | `mixed` | Short explanation describing the work |
| `site_ids` | `mixed` | Unique identifiers of the microsites this work is a part of |
| `sound_ids` | `mixed` | Unique identifiers of the audio about this work |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `style_id` | `string` | Unique identifier of the preferred style term for this work |
| `style_ids` | `mixed` | Unique identifiers of all style terms for this work |
| `style_title` | `mixed` | The name of the preferred style term for this work |
| `style_titles` | `mixed` | The names of all style terms related to this artwork |
| `subject_id` | `string` | Unique identifier of the preferred subject term for this work |
| `subject_ids` | `mixed` | Unique identifiers of all subject terms for this work |
| `subject_titles` | `mixed` | The names of all subject terms related to this artwork |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `technique_id` | `string` | Unique identifier of the preferred technique term for this work |
| `technique_ids` | `mixed` | Unique identifiers of all technique terms for this work |
| `technique_titles` | `mixed` | The names of all technique terms related to this artwork |
| `term_titles` | `mixed` | The names of the taxonomy tags for this work |
| `text_embedding` | `mixed` | The generated embeddings of artwork text |
| `text_ids` | `mixed` | Unique identifiers of the texts about this work |
| `theme_titles` | `mixed` | The names of all thematic publish categories related to this artwork |
| `thumbnail` | `mixed` | Metadata about the image referenced by `image_id`. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |
| `video_ids` | `mixed` | Unique identifiers of the videos about this work |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Artwork record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the ArtworkDateQualifier record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the ArtworkPlaceQualifier record (throws on error).
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
| `aat_id` | `string` | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the ArtworkType record (throws on error).
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
| `aat_id` | `string` | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `parent_id` | `string` | Unique identifier of this category's parent |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `subtype` | `mixed` | Takes one of the following values: classification, material, technique, style, subject, department, theme |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the CategoryTerm record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `copy` | `mixed` | The text of the page |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | The URL to this page on our website |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the DigitalPublication record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `author_display` | `mixed` | A display-friendly text of the authors of this article |
| `copy` | `mixed` | The text of the article |
| `digital_publication_id` | `string` | Unique identifier of the digital publication this article belongs to |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | The URL to this article on our website |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the DigitalPublicationArticle record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `copy` | `mixed` | The text of the page |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | The URL to this page on our website |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the EducatorResource record (throws on error).
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
| `alt_audience_ids` | `mixed` | Unique identifiers indicating the alternate audiences for this event |
| `alt_event_type_ids` | `mixed` | Unique identifiers indicating the alternate types of this event |
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `audience_id` | `string` | Unique identifier indicating the preferred audience for this event |
| `buy_button_caption` | `mixed` | Additional text below the ticket/registration button |
| `buy_button_text` | `mixed` | The text used on the ticket/registration button |
| `date_display` | `mixed` | A readable display of the event dates |
| `description` | `string` | All copytext of the event |
| `door_time` | `mixed` | The time the doors open for this event |
| `end_date` | `mixed` | The date the event ends |
| `end_time` | `mixed` | The time the event ends |
| `entrance` | `mixed` | Which entrance to use for this event |
| `event_host_id` | `string` | Unique identifier of the host (cf. |
| `event_host_title` | `mixed` | Unique identifier of the host (cf. |
| `event_type_id` | `string` | Unique identifier indicating the preferred type of this event |
| `header_description` | `mixed` | Brief description of the event displayed below the title |
| `hero_caption` | `mixed` | Text displayed with the hero image on the event |
| `id` | `string` | Unique identifier of this resource. |
| `image_url` | `mixed` | The URL of an image representing this page |
| `is_admission_required` | `bool` | Whether admission to the museum is required to attend this event |
| `is_after_hours` | `bool` | Whether the event is to be held after the museum closes |
| `is_free` | `bool` | Whether the event is free |
| `is_member_exclusive` | `bool` | Whether the event is exclusive to members of the museum |
| `is_private` | `bool` | Whether the event is private |
| `is_registration_required` | `bool` | Whether registration is required to attend the event |
| `is_sales_button_hidden` | `bool` | Whether the buy tickets button is hidden on the website event page |
| `is_sold_out` | `bool` | Whether the event is sold out |
| `is_ticketed` | `bool` | Whether a ticket is required to attend the event |
| `is_virtual_event` | `bool` | Whether the event is being held virtually |
| `join_url` | `mixed` | URL to the membership signup page via this event |
| `layout_type` | `mixed` | Number indicating the type of layout this event page uses |
| `list_description` | `mixed` | One-sentence description of the event displayed in listings |
| `location` | `mixed` | Where the event takes place |
| `program_ids` | `mixed` | Unique identifiers indicating the programs this event is a part of |
| `program_titles` | `mixed` | Titles of the programs this event is a part of |
| `rsvp_link` | `mixed` | The URL to the sales site for this event |
| `search_tags` | `mixed` | Editor-specified list of tags to aid in internal search |
| `short_description` | `mixed` | Brief description of the event |
| `slug` | `string` | A string used in the URL for this event |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `start_date` | `mixed` | The date the event begins |
| `start_time` | `mixed` | The time the event starts |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `survey_url` | `mixed` | URL to the survey associated with this event |
| `ticketed_event_id` | `string` | Unique identifier of the event in the ticketing system this website event is tied to |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `title_display` | `mixed` | The name of this event formatted with HTML (optional) |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |
| `virtual_event_passcode` | `mixed` | Passcode to access the virtual event |
| `virtual_event_url` | `mixed` | URL to the virtual event |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Event record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `button_caption` | `mixed` | Additional text below the ticket/registration button |
| `button_text` | `mixed` | The text used on the ticket/registration button |
| `button_url` | `mixed` | The URL to the sales site or an RSVP link for this event |
| `description` | `string` | Description of the event |
| `end_at` | `mixed` | The date the event occurrence ends |
| `event_id` | `string` | Identifier of the master event of which this is an occurrence |
| `id` | `string` | Unique identifier of this resource. |
| `image_url` | `mixed` | The URL of an image representing this page |
| `is_private` | `bool` | Whether the event is private. |
| `is_sales_button_hidden` | `bool` | Whether the buy tickets button is hidden on the website event page |
| `is_ticketed` | `bool` | Whether a ticket is required to attend the event |
| `location` | `mixed` | Where the event takes place |
| `off_sale_at` | `mixed` | Date and time the event goes off sale |
| `on_sale_at` | `mixed` | Date and time the event goes on sale |
| `short_description` | `mixed` | Brief description of the event |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `start_at` | `mixed` | The date the event occurrence begins |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `title_display` | `mixed` | The name of this event formatted with HTML (optional) |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the EventOccurrence record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `is_affiliate_group` | `bool` | Whether this program represents an affiliate group |
| `is_event_host` | `bool` | Whether this program represents an event host |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the EventProgram record (throws on error).
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
| `aic_end_at` | `mixed` | Date the exhibition closed at the Art Institute of Chicago |
| `aic_start_at` | `mixed` | Date the exhibition opened at the Art Institute of Chicago |
| `alt_image_ids` | `mixed` | Unique identifiers of all non-preferred images of this exhibition. |
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `artist_ids` | `mixed` | Unique identifiers of the artist agent records representing who was shown in the exhibition |
| `artwork_ids` | `mixed` | Unique identifiers of the artworks that were part of the exhibition |
| `artwork_titles` | `mixed` | Names of the artworks that were part of the exhibition |
| `document_ids` | `mixed` | Unique identifiers of assets that serve as documentation for this exhibition |
| `gallery_id` | `string` | Unique identifier of the gallery that mainly housed the exhibition |
| `gallery_title` | `mixed` | The name of the gallery that mainly housed the exhibition |
| `id` | `string` | Unique identifier of this resource. |
| `image_id` | `string` | Unique identifier of the preferred image to use to represent this exhibition |
| `image_url` | `mixed` | URL to the hero image from the website |
| `is_featured` | `bool` | Is this exhibition currently featured on our website? |
| `is_published` | `bool` | Is this exhibition currently published on our website? |
| `position` | `mixed` | Numering position represnting the order in which this exhibition is featured on the website |
| `short_description` | `mixed` | Brief explanation of what this exhibition is |
| `site_ids` | `mixed` | Unique identifiers of the microsites this exhibition is a part of |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `status` | `mixed` | Whether the exhibition is open or closed |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | URL to this exhibition on our website |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Exhibition record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `floor` | `mixed` | The level the gallery is on, e.g., 1, 2, 3, or LL |
| `id` | `string` | Unique identifier of this resource. |
| `is_closed` | `bool` | Whether the gallery is currently closed |
| `latitude` | `float` | Latitude coordinate of the center of the room |
| `latlon` | `mixed` | Latitude and longitude coordinates of the center of the room |
| `longitude` | `float` | Longitude coordinate of the center of the room |
| `number` | `mixed` | The gallery's room number. |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | `string` | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Gallery record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `copy` | `mixed` | The text of the page |
| `id` | `string` | Unique identifier of this resource. |
| `search_tags` | `mixed` | Editor-specified list of tags to aid in internal search |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | The URL to this page on our website |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the GenericPage record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `copy` | `mixed` | The text of the highlight description |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Highlight record (throws on error).
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
| `additional_text` | `mixed` | Additional information about the hours |
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `friday_is_closed` | `mixed` | Whether the museum is closed on Fridays |
| `friday_member_close` | `mixed` | The time member hours ends on Fridays |
| `friday_member_open` | `mixed` | The time member hours starts on Fridays |
| `friday_public_close` | `mixed` | The time public hours ends on Fridays |
| `friday_public_open` | `mixed` | The time public hours starts on Fridays |
| `id` | `string` | Unique identifier of this resource. |
| `monday_is_closed` | `mixed` | Whether the museum is closed on Mondays |
| `monday_member_close` | `mixed` | The time member hours ends on Mondays |
| `monday_member_open` | `mixed` | The time member hours starts on Mondays |
| `monday_public_close` | `mixed` | The time public hours ends on Mondays |
| `monday_public_open` | `mixed` | The time public hours starts on Mondays |
| `saturday_is_closed` | `mixed` | Whether the museum is closed on Saturdays |
| `saturday_member_close` | `mixed` | The time member hours ends on Saturdays |
| `saturday_member_open` | `mixed` | The time member hours starts on Saturdays |
| `saturday_public_close` | `mixed` | The time public hours ends on Saturdays |
| `saturday_public_open` | `mixed` | The time public hours starts on Saturdays |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `summary` | `mixed` | Readable summary of the hours |
| `sunday_is_closed` | `mixed` | Whether the museum is closed on Sundays |
| `sunday_member_close` | `mixed` | The time member hours ends on Sundays |
| `sunday_member_open` | `mixed` | The time member hours starts on Sundays |
| `sunday_public_close` | `mixed` | The time public hours ends on Sundays |
| `sunday_public_open` | `mixed` | The time public hours starts on Sundays |
| `thursday_is_closed` | `mixed` | Whether the museum is closed on Thursdays |
| `thursday_member_close` | `mixed` | The time member hours ends on Thursdays |
| `thursday_member_open` | `mixed` | The time member hours starts on Thursdays |
| `thursday_public_close` | `mixed` | The time public hours ends on Thursdays |
| `thursday_public_open` | `mixed` | The time public hours starts on Thursdays |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `tuesday_is_closed` | `mixed` | Whether the museum is closed on Tuesdays |
| `tuesday_member_close` | `mixed` | The time member hours ends on Tuesdays |
| `tuesday_member_open` | `mixed` | The time member hours starts on Tuesdays |
| `tuesday_public_close` | `mixed` | The time public hours ends on Tuesdays |
| `tuesday_public_open` | `mixed` | The time public hours starts on Tuesdays |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |
| `wednesday_is_closed` | `mixed` | Whether the museum is closed on Wednesdays |
| `wednesday_member_close` | `mixed` | The time member hours ends on Wednesdays |
| `wednesday_member_open` | `mixed` | The time member hours starts on Wednesdays |
| `wednesday_public_close` | `mixed` | The time public hours ends on Wednesdays |
| `wednesday_public_open` | `mixed` | The time public hours starts on Wednesdays |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Hour record (throws on error).
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
| `ahash` | `mixed` | Image hash generated using ahash algorithm with 64 boolean subfields |
| `alt_text` | `mixed` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `artwork_ids` | `mixed` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `mixed` | Names of the artworks associated with this asset |
| `color` | `mixed` | Dominant color of this image in HSL |
| `colorfulness` | `mixed` | Unbounded positive float representing an abstract measure of colorfulness. |
| `content` | `mixed` | Text of or URL to the contents of this asset |
| `content_e_tag` | `mixed` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `mixed` | Asset-specific copyright information |
| `fingerprint` | `mixed` | Image hashes: aHash, dHash, pHash, wHash |
| `height` | `float` | Native height of the image |
| `id` | `string` | Unique identifier of this resource. |
| `iiif_url` | `mixed` | IIIF URL of this image |
| `is_educational_resource` | `bool` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | Whether this resource is considered to be educational |
| `lake_guid` | `mixed` | Unique UUID of this resource in LAKE, our DAMS. |
| `lqip` | `mixed` | Low-quality image placeholder (LQIP). |
| `phash` | `mixed` | Image hash generated using phash algorithm with 64 boolean subfields |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `type` | `mixed` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |
| `width` | `float` | Native width of the image |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Image record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `copy` | `mixed` | The text of the page |
| `id` | `string` | Unique identifier of this resource. |
| `search_tags` | `mixed` | Editor-specified list of tags to aid in internal search |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | The URL to this page on our website |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the LandingPage record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `latitude` | `float` | Latitude coordinate of the center of the room |
| `longitude` | `float` | Longitude coordinate of the center of the room |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | `string` | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Place record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `copy` | `mixed` | The text of the page |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | The URL to this page on our website |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the PressRelease record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `copy` | `mixed` | The text of the page |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | The URL to this page on our website |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the PrintedPublication record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `artist_ids` | `mixed` | Unique identifiers of the artists associated with this product |
| `artwork_ids` | `mixed` | Unique identifiers of the artworks associated with this product |
| `description` | `string` | Explanation of what this product is |
| `exhibition_ids` | `mixed` | Unique identifiers of the exhibitions associated with this product |
| `external_sku` | `mixed` | Numeric product identification code of a machine-readable barcode, when the customer sku differs from our internal one |
| `id` | `string` | Unique identifier of this resource. |
| `image_url` | `mixed` | URL of an image for this product |
| `max_compare_at_price` | `mixed` | Number indicating how much the most expensive variant of a product cost before a sale |
| `max_current_price` | `mixed` | Number indicating how much the most expensive variant of a product costs right now |
| `min_compare_at_price` | `mixed` | Number indicating how much the least expensive variant of a product cost before a sale |
| `min_current_price` | `mixed` | Number indicating how much the least expensive variant of a product costs right now |
| `price_display` | `mixed` | Explanation of what this product is |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | URL of this product in the shop |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Product record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `section_ids` | `mixed` | Unique identifiers of the sections of this publication |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | URL to the publication |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Publication record (throws on error).
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
| `api_id` | `string` | API unique identifier |
| `api_link` | `mixed` | URL to this recource in the API |
| `api_model` | `mixed` | Name of the model the resource represents |
| `id` | `string` | Unique identifier within the search index |
| `is_boosted` | `bool` | Whether this record has been flagged to be boosted |
| `score` | `float` | Search index ranking of the result |
| `thumbnail` | `mixed` | Metadata on the image representing this record |
| `timestamp` | `mixed` | Date this record was last updated in the API |
| `title` | `string` | The name of this resource |

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
| `accession` | `mixed` | An accession number parsed from the title or tombstone |
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `artwork_id` | `string` | Unique identifier of the artwork with which this section is associated |
| `content` | `mixed` | Content of this section in plaintext |
| `generic_page_id` | `string` | Unique identifier of the page on the website that represents the publication this section belongs to |
| `id` | `string` | Unique identifier of this resource. |
| `publication_id` | `string` | Unique identifier of the publication this section belongs to |
| `publication_title` | `mixed` | Name of the publication this section belongs to |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | URL to the section |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Section record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `artwork_ids` | `mixed` | Unique identifiers of the artworks this site is associated with |
| `artwork_titles` | `mixed` | Names of the artworks this site is associated with |
| `description` | `string` | Explanation of what this site is |
| `exhibition_ids` | `mixed` | Unique identifier of the exhibitions this site is associated with |
| `exhibition_titles` | `mixed` | Names of the exhibitions this site is associated with |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | URL to this site |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Site record (throws on error).
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
| `alt_text` | `mixed` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `artwork_ids` | `mixed` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `mixed` | Names of the artworks associated with this asset |
| `content` | `mixed` | Text of or URL to the contents of this asset |
| `content_e_tag` | `mixed` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `mixed` | Asset-specific copyright information |
| `id` | `string` | Unique identifier of this resource. |
| `is_educational_resource` | `bool` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | Whether this resource is considered to be educational |
| `lake_guid` | `mixed` | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | Name of this mobile audio file – derived from the artwork and tour titles |
| `transcript` | `mixed` | Text transcription of the audio file |
| `type` | `mixed` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | URL to the audio file |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Sound record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | The URL to this page on our website |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the StaticPage record (throws on error).
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
| `alt_text` | `mixed` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `artwork_ids` | `mixed` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `mixed` | Names of the artworks associated with this asset |
| `content` | `mixed` | Text of or URL to the contents of this asset |
| `content_e_tag` | `mixed` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `mixed` | Asset-specific copyright information |
| `id` | `string` | Unique identifier of this resource. |
| `is_educational_resource` | `bool` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | Whether this resource is considered to be educational |
| `lake_guid` | `mixed` | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `type` | `mixed` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Text record (throws on error).
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
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `artist_titles` | `mixed` | Names of the artists of the artworks featured in this tour's tour stops |
| `artwork_titles` | `mixed` | Names of the artworks featured in this tour's tour stops |
| `description` | `string` | Explanation of what the tour is |
| `id` | `string` | Unique identifier of this resource. |
| `image` | `mixed` | The main image for the tour |
| `intro` | `mixed` | Text introducing the tour |
| `intro_link` | `mixed` | Link to the audio file of the introduction |
| `intro_transcript` | `mixed` | Transcript of the introduction audio to the tour |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |
| `weight` | `float` | Number representing this tour's sort order |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Tour record (throws on error).
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
| `alt_text` | `mixed` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `mixed` | REST API link for this resource |
| `api_model` | `mixed` | REST API resource type or endpoint |
| `artwork_ids` | `mixed` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `mixed` | Names of the artworks associated with this asset |
| `content` | `mixed` | Text of or URL to the contents of this asset |
| `content_e_tag` | `mixed` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `mixed` | Asset-specific copyright information |
| `id` | `string` | Unique identifier of this resource. |
| `is_educational_resource` | `bool` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | Whether this resource is considered to be educational |
| `lake_guid` | `mixed` | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `mixed` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `type` | `mixed` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `mixed` | Date and time the record was updated in the aggregator database |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Video record (throws on error).
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
$publication = $client->Publication();
$publication->list();

// $publication->data_get() now returns the publication data from the last list
// $publication->match_get() returns the last match criteria
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
