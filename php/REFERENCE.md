# ArtInstituteOfChicago PHP SDK Reference

Complete API reference for the ArtInstituteOfChicago PHP SDK.


## ArtInstituteOfChicagoSDK

### Constructor

```php
require_once __DIR__ . '/artinstituteofchicago_sdk.php';

$client = new ArtInstituteOfChicagoSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `ArtInstituteOfChicagoSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = ArtInstituteOfChicagoSDK::test();
```


### Instance Methods

#### `Agent($data = null)`

Create a new `AgentEntity` instance. Pass `null` for no initial data.

#### `AgentRole($data = null)`

Create a new `AgentRoleEntity` instance. Pass `null` for no initial data.

#### `AgentType($data = null)`

Create a new `AgentTypeEntity` instance. Pass `null` for no initial data.

#### `Article($data = null)`

Create a new `ArticleEntity` instance. Pass `null` for no initial data.

#### `Artwork($data = null)`

Create a new `ArtworkEntity` instance. Pass `null` for no initial data.

#### `ArtworkDateQualifier($data = null)`

Create a new `ArtworkDateQualifierEntity` instance. Pass `null` for no initial data.

#### `ArtworkPlaceQualifier($data = null)`

Create a new `ArtworkPlaceQualifierEntity` instance. Pass `null` for no initial data.

#### `ArtworkType($data = null)`

Create a new `ArtworkTypeEntity` instance. Pass `null` for no initial data.

#### `CategoryTerm($data = null)`

Create a new `CategoryTermEntity` instance. Pass `null` for no initial data.

#### `DigitalPublication($data = null)`

Create a new `DigitalPublicationEntity` instance. Pass `null` for no initial data.

#### `DigitalPublicationArticle($data = null)`

Create a new `DigitalPublicationArticleEntity` instance. Pass `null` for no initial data.

#### `EducatorResource($data = null)`

Create a new `EducatorResourceEntity` instance. Pass `null` for no initial data.

#### `Event($data = null)`

Create a new `EventEntity` instance. Pass `null` for no initial data.

#### `EventOccurrence($data = null)`

Create a new `EventOccurrenceEntity` instance. Pass `null` for no initial data.

#### `EventProgram($data = null)`

Create a new `EventProgramEntity` instance. Pass `null` for no initial data.

#### `Exhibition($data = null)`

Create a new `ExhibitionEntity` instance. Pass `null` for no initial data.

#### `Gallery($data = null)`

Create a new `GalleryEntity` instance. Pass `null` for no initial data.

#### `GenericPage($data = null)`

Create a new `GenericPageEntity` instance. Pass `null` for no initial data.

#### `Highlight($data = null)`

Create a new `HighlightEntity` instance. Pass `null` for no initial data.

#### `Hour($data = null)`

Create a new `HourEntity` instance. Pass `null` for no initial data.

#### `Image($data = null)`

Create a new `ImageEntity` instance. Pass `null` for no initial data.

#### `LandingPage($data = null)`

Create a new `LandingPageEntity` instance. Pass `null` for no initial data.

#### `Place($data = null)`

Create a new `PlaceEntity` instance. Pass `null` for no initial data.

#### `PressRelease($data = null)`

Create a new `PressReleaseEntity` instance. Pass `null` for no initial data.

#### `PrintedPublication($data = null)`

Create a new `PrintedPublicationEntity` instance. Pass `null` for no initial data.

#### `Product($data = null)`

Create a new `ProductEntity` instance. Pass `null` for no initial data.

#### `Publication($data = null)`

Create a new `PublicationEntity` instance. Pass `null` for no initial data.

#### `Search($data = null)`

Create a new `SearchEntity` instance. Pass `null` for no initial data.

#### `Section($data = null)`

Create a new `SectionEntity` instance. Pass `null` for no initial data.

#### `Site($data = null)`

Create a new `SiteEntity` instance. Pass `null` for no initial data.

#### `Sound($data = null)`

Create a new `SoundEntity` instance. Pass `null` for no initial data.

#### `StaticPage($data = null)`

Create a new `StaticPageEntity` instance. Pass `null` for no initial data.

#### `Text($data = null)`

Create a new `TextEntity` instance. Pass `null` for no initial data.

#### `Tour($data = null)`

Create a new `TourEntity` instance. Pass `null` for no initial data.

#### `Video($data = null)`

Create a new `VideoEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): ArtInstituteOfChicagoUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## AgentEntity

```php
$agent = $client->Agent();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_titles` | `mixed` | No | Alternate names for this agent |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `birth_date` | `mixed` | No | The year this agent was born |
| `death_date` | `mixed` | No | The year this agent died |
| `description` | `string` | No | A biographical description of the agent |
| `id` | `string` | No | Unique identifier of this resource. |
| `is_artist` | `bool` | No | Whether the agent is an artist. |
| `sort_title` | `mixed` | No | Sortable name for this agent, typically with last name first. |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `ulan_id` | `string` | No | Unique identifier of this agent in Getty's ULAN |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Agent()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Agent()->load(["id" => "agent_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): AgentEntity`

Create a new `AgentEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## AgentRoleEntity

```php
$agent_role = $client->AgentRole();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->AgentRole()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->AgentRole()->load(["id" => "agent_role_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): AgentRoleEntity`

Create a new `AgentRoleEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## AgentTypeEntity

```php
$agent_type = $client->AgentType();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->AgentType()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->AgentType()->load(["id" => "agent_type_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): AgentTypeEntity`

Create a new `AgentTypeEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ArticleEntity

```php
$article = $client->Article();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `copy` | `mixed` | No | The text of the article |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Article()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Article()->load(["id" => "article_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ArticleEntity`

Create a new `ArticleEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ArtworkEntity

```php
$artwork = $client->Artwork();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_artist_ids` | `mixed` | No | Unique identifiers of the non-preferred artists/cultures associated with this work |
| `alt_classification_ids` | `mixed` | No | Unique identifiers of all other non-preferred classification terms for this work |
| `alt_image_ids` | `mixed` | No | Unique identifiers of all non-preferred images of this work. |
| `alt_material_ids` | `mixed` | No | Unique identifiers of all other non-preferred material terms for this work |
| `alt_style_ids` | `mixed` | No | Unique identifiers of all other non-preferred style terms for this work |
| `alt_subject_ids` | `mixed` | No | Unique identifiers of all other non-preferred subject terms for this work |
| `alt_technique_ids` | `mixed` | No | Unique identifiers of all other non-preferred technique terms for this work |
| `alt_titles` | `mixed` | No | Alternate names for this work |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `artist_display` | `mixed` | No | Readable description of the creator of this work. |
| `artist_id` | `string` | No | Unique identifier of the preferred artist/culture associated with this work |
| `artist_ids` | `mixed` | No | Unique identifier of all artist/cultures associated with this work |
| `artist_title` | `mixed` | No | Name of the preferred artist/culture associated with this work |
| `artist_titles` | `mixed` | No | Names of all artist/cultures associated with this work |
| `artwork_type_id` | `string` | No | Unique identifier of the kind of object or work |
| `artwork_type_title` | `mixed` | No | The kind of object or work (e.g. |
| `boost_rank` | `mixed` | No | Manual indication of what rank this artwork should take in search results. |
| `catalog_based_search_keyword_titles` | `mixed` | No | The keyword search values that would be catalog-based searches on this record |
| `catalogue_display` | `mixed` | No | Brief text listing all the catalogues raisonnés which include this work. |
| `category_ids` | `mixed` | No | Unique identifiers of the categories this work is a part of |
| `category_titles` | `mixed` | No | Names of the categories this artwork is a part of |
| `classification_id` | `string` | No | Unique identifier of the preferred classification term for this work |
| `classification_ids` | `mixed` | No | Unique identifiers of all classification terms for this work |
| `classification_title` | `mixed` | No | The name of the preferred classification term for this work |
| `classification_titles` | `mixed` | No | The names of all classification terms related to this artwork |
| `color` | `mixed` | No | Dominant color of this artwork in HSL |
| `colorfulness` | `mixed` | No | Unbounded positive float representing an abstract measure of colorfulness. |
| `copyright_notice` | `mixed` | No | Statement notifying how the work is protected by copyright. |
| `credit_line` | `mixed` | No | Brief statement indicating how the work came into the collection |
| `date_display` | `mixed` | No | Readable, free-text description of the period of time associated with the creation of this work. |
| `date_end` | `mixed` | No | The year of the period of time associated with the creation of this work |
| `date_qualifier_id` | `string` | No | Unique identifier of the qualifer to the dates provided for this record. |
| `date_qualifier_title` | `mixed` | No | Readable, text qualifer to the dates provided for this record. |
| `date_start` | `mixed` | No | The year of the period of time associated with the creation of this work |
| `department_id` | `string` | No | Unique identifier of the curatorial department that this work belongs to |
| `department_title` | `mixed` | No | Name of the curatorial department that this work belongs to |
| `description` | `string` | No | Longer explanation describing the work |
| `dimensions` | `mixed` | No | The size, shape, scale, and dimensions of the work. |
| `dimensions_detail` | `mixed` | No | The height, width, depth, and/or diameter of each section of the work in centimeters |
| `document_ids` | `mixed` | No | Unique identifiers of assets that serve as documentation for this artwork |
| `edition` | `mixed` | No | Edition number if the work is one of many |
| `exhibition_history` | `mixed` | No | List of all the places this work has been exhibited |
| `fiscal_year` | `mixed` | No | The fiscal year in which the work was acquired. |
| `fiscal_year_deaccession` | `mixed` | No | The fiscal year in which the work was deaccessioned. |
| `gallery_id` | `string` | No | Unique identifier of the location of this work in our museum |
| `gallery_title` | `mixed` | No | The location of this work in our museum |
| `has_advanced_imaging` | `bool` | No | Whether this artwork is enhanced with 3D models, 360 image sequences, Mirador views, etc. |
| `has_educational_resources` | `bool` | No | Whether this artwork has any documents tagged as educational |
| `has_multimedia_resources` | `bool` | No | Whether this artwork has any associated microsites, digital publications, or documents tagged as multimedia |
| `has_not_been_viewed_much` | `bool` | No | Whether the artwork hasn't been visited on our website very much |
| `id` | `string` | No | Unique identifier of this resource. |
| `image_embedding` | `mixed` | No | The generated embeddings describing the artwork image |
| `image_id` | `string` | No | Unique identifier of the preferred image to use to represent this work |
| `inscriptions` | `mixed` | No | A description of distinguishing or identifying physical markings that are on the work |
| `internal_department_id` | `string` | No | An internal department id we use for analytics. |
| `is_boosted` | `bool` | No | Whether this document should be boosted in search |
| `is_on_view` | `bool` | No | Whether the work is on display |
| `is_public_domain` | `bool` | No | Whether the work is in the public domain, meaning it was created before copyrights existed or has left the copyright term |
| `is_zoomable` | `bool` | No | Whether images of the work are allowed to be displayed in a zoomable interface. |
| `latitude` | `float` | No | Latitude coordinate of the location of this work in our galleries |
| `latlon` | `mixed` | No | Latitude and longitude coordinates of the location of this work in our galleries |
| `longitude` | `float` | No | Longitude coordinate of the location of this work in our galleries |
| `main_reference_number` | `int` | No | Unique identifier assigned to the artwork upon acquisition |
| `material_id` | `string` | No | Unique identifier of the preferred material term for this work |
| `material_ids` | `mixed` | No | Unique identifiers of all material terms for this work |
| `material_titles` | `mixed` | No | The names of all material terms related to this artwork |
| `max_zoom_window_size` | `mixed` | No | The maximum size of the window the image is allowed to be viewed in, in pixels. |
| `medium_display` | `mixed` | No | The substances or materials used in the creation of a work |
| `nomisma_id` | `string` | No | Unique identifier of this work in the nomisma coin database |
| `on_loan_display` | `mixed` | No | If an artwork is on loan, this contains details about the loan |
| `pageviews` | `mixed` | No | Approx. |
| `pageviews_recent` | `mixed` | No | Approx. |
| `place_of_origin` | `mixed` | No | The location where the creation, design, or production of the work took place, or the original location of the work |
| `provenance_text` | `mixed` | No | Ownership/collecting history of the work. |
| `publication_history` | `mixed` | No | Bibliographic list of all the places this work has been published |
| `publishing_verification_level` | `mixed` | No | Indicator of how much metadata on the work in published. |
| `section_ids` | `mixed` | No | Unique identifiers of the digital publication chapters this work in included in |
| `section_titles` | `mixed` | No | Names of the digital publication chapters this work is included in |
| `short_description` | `mixed` | No | Short explanation describing the work |
| `site_ids` | `mixed` | No | Unique identifiers of the microsites this work is a part of |
| `sound_ids` | `mixed` | No | Unique identifiers of the audio about this work |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `style_id` | `string` | No | Unique identifier of the preferred style term for this work |
| `style_ids` | `mixed` | No | Unique identifiers of all style terms for this work |
| `style_title` | `mixed` | No | The name of the preferred style term for this work |
| `style_titles` | `mixed` | No | The names of all style terms related to this artwork |
| `subject_id` | `string` | No | Unique identifier of the preferred subject term for this work |
| `subject_ids` | `mixed` | No | Unique identifiers of all subject terms for this work |
| `subject_titles` | `mixed` | No | The names of all subject terms related to this artwork |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `technique_id` | `string` | No | Unique identifier of the preferred technique term for this work |
| `technique_ids` | `mixed` | No | Unique identifiers of all technique terms for this work |
| `technique_titles` | `mixed` | No | The names of all technique terms related to this artwork |
| `term_titles` | `mixed` | No | The names of the taxonomy tags for this work |
| `text_embedding` | `mixed` | No | The generated embeddings of artwork text |
| `text_ids` | `mixed` | No | Unique identifiers of the texts about this work |
| `theme_titles` | `mixed` | No | The names of all thematic publish categories related to this artwork |
| `thumbnail` | `mixed` | No | Metadata about the image referenced by `image_id`. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |
| `video_ids` | `mixed` | No | Unique identifiers of the videos about this work |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Artwork()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Artwork()->load(["id" => "artwork_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ArtworkEntity`

Create a new `ArtworkEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ArtworkDateQualifierEntity

```php
$artwork_date_qualifier = $client->ArtworkDateQualifier();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->ArtworkDateQualifier()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->ArtworkDateQualifier()->load(["id" => "artwork_date_qualifier_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ArtworkDateQualifierEntity`

Create a new `ArtworkDateQualifierEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ArtworkPlaceQualifierEntity

```php
$artwork_place_qualifier = $client->ArtworkPlaceQualifier();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->ArtworkPlaceQualifier()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->ArtworkPlaceQualifier()->load(["id" => "artwork_place_qualifier_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ArtworkPlaceQualifierEntity`

Create a new `ArtworkPlaceQualifierEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ArtworkTypeEntity

```php
$artwork_type = $client->ArtworkType();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `aat_id` | `string` | No | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->ArtworkType()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->ArtworkType()->load(["id" => "artwork_type_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ArtworkTypeEntity`

Create a new `ArtworkTypeEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CategoryTermEntity

```php
$category_term = $client->CategoryTerm();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `aat_id` | `string` | No | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `parent_id` | `string` | No | Unique identifier of this category's parent |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `subtype` | `mixed` | No | Takes one of the following values: classification, material, technique, style, subject, department, theme |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->CategoryTerm()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->CategoryTerm()->load(["id" => "category_term_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CategoryTermEntity`

Create a new `CategoryTermEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## DigitalPublicationEntity

```php
$digital_publication = $client->DigitalPublication();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `copy` | `mixed` | No | The text of the page |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | No | The URL to this page on our website |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->DigitalPublication()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->DigitalPublication()->load(["id" => "digital_publication_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): DigitalPublicationEntity`

Create a new `DigitalPublicationEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## DigitalPublicationArticleEntity

```php
$digital_publication_article = $client->DigitalPublicationArticle();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `author_display` | `mixed` | No | A display-friendly text of the authors of this article |
| `copy` | `mixed` | No | The text of the article |
| `digital_publication_id` | `string` | No | Unique identifier of the digital publication this article belongs to |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | No | The URL to this article on our website |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->DigitalPublicationArticle()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->DigitalPublicationArticle()->load(["id" => "digital_publication_article_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): DigitalPublicationArticleEntity`

Create a new `DigitalPublicationArticleEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EducatorResourceEntity

```php
$educator_resource = $client->EducatorResource();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `copy` | `mixed` | No | The text of the page |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | No | The URL to this page on our website |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->EducatorResource()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->EducatorResource()->load(["id" => "educator_resource_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EducatorResourceEntity`

Create a new `EducatorResourceEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EventEntity

```php
$event = $client->Event();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_audience_ids` | `mixed` | No | Unique identifiers indicating the alternate audiences for this event |
| `alt_event_type_ids` | `mixed` | No | Unique identifiers indicating the alternate types of this event |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `audience_id` | `string` | No | Unique identifier indicating the preferred audience for this event |
| `buy_button_caption` | `mixed` | No | Additional text below the ticket/registration button |
| `buy_button_text` | `mixed` | No | The text used on the ticket/registration button |
| `date_display` | `mixed` | No | A readable display of the event dates |
| `description` | `string` | No | All copytext of the event |
| `door_time` | `mixed` | No | The time the doors open for this event |
| `end_date` | `mixed` | No | The date the event ends |
| `end_time` | `mixed` | No | The time the event ends |
| `entrance` | `mixed` | No | Which entrance to use for this event |
| `event_host_id` | `string` | No | Unique identifier of the host (cf. |
| `event_host_title` | `mixed` | No | Unique identifier of the host (cf. |
| `event_type_id` | `string` | No | Unique identifier indicating the preferred type of this event |
| `header_description` | `mixed` | No | Brief description of the event displayed below the title |
| `hero_caption` | `mixed` | No | Text displayed with the hero image on the event |
| `id` | `string` | No | Unique identifier of this resource. |
| `image_url` | `mixed` | No | The URL of an image representing this page |
| `is_admission_required` | `bool` | No | Whether admission to the museum is required to attend this event |
| `is_after_hours` | `bool` | No | Whether the event is to be held after the museum closes |
| `is_free` | `bool` | No | Whether the event is free |
| `is_member_exclusive` | `bool` | No | Whether the event is exclusive to members of the museum |
| `is_private` | `bool` | No | Whether the event is private |
| `is_registration_required` | `bool` | No | Whether registration is required to attend the event |
| `is_sales_button_hidden` | `bool` | No | Whether the buy tickets button is hidden on the website event page |
| `is_sold_out` | `bool` | No | Whether the event is sold out |
| `is_ticketed` | `bool` | No | Whether a ticket is required to attend the event |
| `is_virtual_event` | `bool` | No | Whether the event is being held virtually |
| `join_url` | `mixed` | No | URL to the membership signup page via this event |
| `layout_type` | `mixed` | No | Number indicating the type of layout this event page uses |
| `list_description` | `mixed` | No | One-sentence description of the event displayed in listings |
| `location` | `mixed` | No | Where the event takes place |
| `program_ids` | `mixed` | No | Unique identifiers indicating the programs this event is a part of |
| `program_titles` | `mixed` | No | Titles of the programs this event is a part of |
| `rsvp_link` | `mixed` | No | The URL to the sales site for this event |
| `search_tags` | `mixed` | No | Editor-specified list of tags to aid in internal search |
| `short_description` | `mixed` | No | Brief description of the event |
| `slug` | `string` | No | A string used in the URL for this event |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `start_date` | `mixed` | No | The date the event begins |
| `start_time` | `mixed` | No | The time the event starts |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `survey_url` | `mixed` | No | URL to the survey associated with this event |
| `ticketed_event_id` | `string` | No | Unique identifier of the event in the ticketing system this website event is tied to |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `title_display` | `mixed` | No | The name of this event formatted with HTML (optional) |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |
| `virtual_event_passcode` | `mixed` | No | Passcode to access the virtual event |
| `virtual_event_url` | `mixed` | No | URL to the virtual event |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Event()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Event()->load(["id" => "event_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EventEntity`

Create a new `EventEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EventOccurrenceEntity

```php
$event_occurrence = $client->EventOccurrence();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `button_caption` | `mixed` | No | Additional text below the ticket/registration button |
| `button_text` | `mixed` | No | The text used on the ticket/registration button |
| `button_url` | `mixed` | No | The URL to the sales site or an RSVP link for this event |
| `description` | `string` | No | Description of the event |
| `end_at` | `mixed` | No | The date the event occurrence ends |
| `event_id` | `string` | No | Identifier of the master event of which this is an occurrence |
| `id` | `string` | No | Unique identifier of this resource. |
| `image_url` | `mixed` | No | The URL of an image representing this page |
| `is_private` | `bool` | No | Whether the event is private. |
| `is_sales_button_hidden` | `bool` | No | Whether the buy tickets button is hidden on the website event page |
| `is_ticketed` | `bool` | No | Whether a ticket is required to attend the event |
| `location` | `mixed` | No | Where the event takes place |
| `off_sale_at` | `mixed` | No | Date and time the event goes off sale |
| `on_sale_at` | `mixed` | No | Date and time the event goes on sale |
| `short_description` | `mixed` | No | Brief description of the event |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `start_at` | `mixed` | No | The date the event occurrence begins |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `title_display` | `mixed` | No | The name of this event formatted with HTML (optional) |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->EventOccurrence()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->EventOccurrence()->load(["id" => "event_occurrence_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EventOccurrenceEntity`

Create a new `EventOccurrenceEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EventProgramEntity

```php
$event_program = $client->EventProgram();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `is_affiliate_group` | `bool` | No | Whether this program represents an affiliate group |
| `is_event_host` | `bool` | No | Whether this program represents an event host |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->EventProgram()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->EventProgram()->load(["id" => "event_program_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EventProgramEntity`

Create a new `EventProgramEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ExhibitionEntity

```php
$exhibition = $client->Exhibition();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `aic_end_at` | `mixed` | No | Date the exhibition closed at the Art Institute of Chicago |
| `aic_start_at` | `mixed` | No | Date the exhibition opened at the Art Institute of Chicago |
| `alt_image_ids` | `mixed` | No | Unique identifiers of all non-preferred images of this exhibition. |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `artist_ids` | `mixed` | No | Unique identifiers of the artist agent records representing who was shown in the exhibition |
| `artwork_ids` | `mixed` | No | Unique identifiers of the artworks that were part of the exhibition |
| `artwork_titles` | `mixed` | No | Names of the artworks that were part of the exhibition |
| `document_ids` | `mixed` | No | Unique identifiers of assets that serve as documentation for this exhibition |
| `gallery_id` | `string` | No | Unique identifier of the gallery that mainly housed the exhibition |
| `gallery_title` | `mixed` | No | The name of the gallery that mainly housed the exhibition |
| `id` | `string` | No | Unique identifier of this resource. |
| `image_id` | `string` | No | Unique identifier of the preferred image to use to represent this exhibition |
| `image_url` | `mixed` | No | URL to the hero image from the website |
| `is_featured` | `bool` | No | Is this exhibition currently featured on our website? |
| `is_published` | `bool` | No | Is this exhibition currently published on our website? |
| `position` | `mixed` | No | Numering position represnting the order in which this exhibition is featured on the website |
| `short_description` | `mixed` | No | Brief explanation of what this exhibition is |
| `site_ids` | `mixed` | No | Unique identifiers of the microsites this exhibition is a part of |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `status` | `mixed` | No | Whether the exhibition is open or closed |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | No | URL to this exhibition on our website |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Exhibition()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Exhibition()->load(["id" => "exhibition_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ExhibitionEntity`

Create a new `ExhibitionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## GalleryEntity

```php
$gallery = $client->Gallery();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `floor` | `mixed` | No | The level the gallery is on, e.g., 1, 2, 3, or LL |
| `id` | `string` | No | Unique identifier of this resource. |
| `is_closed` | `bool` | No | Whether the gallery is currently closed |
| `latitude` | `float` | No | Latitude coordinate of the center of the room |
| `latlon` | `mixed` | No | Latitude and longitude coordinates of the center of the room |
| `longitude` | `float` | No | Longitude coordinate of the center of the room |
| `number` | `mixed` | No | The gallery's room number. |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | `string` | No | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Gallery()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Gallery()->load(["id" => "gallery_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): GalleryEntity`

Create a new `GalleryEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## GenericPageEntity

```php
$generic_page = $client->GenericPage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `copy` | `mixed` | No | The text of the page |
| `id` | `string` | No | Unique identifier of this resource. |
| `search_tags` | `mixed` | No | Editor-specified list of tags to aid in internal search |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | No | The URL to this page on our website |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->GenericPage()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->GenericPage()->load(["id" => "generic_page_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): GenericPageEntity`

Create a new `GenericPageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## HighlightEntity

```php
$highlight = $client->Highlight();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `copy` | `mixed` | No | The text of the highlight description |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Highlight()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Highlight()->load(["id" => "highlight_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): HighlightEntity`

Create a new `HighlightEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## HourEntity

```php
$hour = $client->Hour();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_text` | `mixed` | No | Additional information about the hours |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `friday_is_closed` | `mixed` | No | Whether the museum is closed on Fridays |
| `friday_member_close` | `mixed` | No | The time member hours ends on Fridays |
| `friday_member_open` | `mixed` | No | The time member hours starts on Fridays |
| `friday_public_close` | `mixed` | No | The time public hours ends on Fridays |
| `friday_public_open` | `mixed` | No | The time public hours starts on Fridays |
| `id` | `string` | No | Unique identifier of this resource. |
| `monday_is_closed` | `mixed` | No | Whether the museum is closed on Mondays |
| `monday_member_close` | `mixed` | No | The time member hours ends on Mondays |
| `monday_member_open` | `mixed` | No | The time member hours starts on Mondays |
| `monday_public_close` | `mixed` | No | The time public hours ends on Mondays |
| `monday_public_open` | `mixed` | No | The time public hours starts on Mondays |
| `saturday_is_closed` | `mixed` | No | Whether the museum is closed on Saturdays |
| `saturday_member_close` | `mixed` | No | The time member hours ends on Saturdays |
| `saturday_member_open` | `mixed` | No | The time member hours starts on Saturdays |
| `saturday_public_close` | `mixed` | No | The time public hours ends on Saturdays |
| `saturday_public_open` | `mixed` | No | The time public hours starts on Saturdays |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `summary` | `mixed` | No | Readable summary of the hours |
| `sunday_is_closed` | `mixed` | No | Whether the museum is closed on Sundays |
| `sunday_member_close` | `mixed` | No | The time member hours ends on Sundays |
| `sunday_member_open` | `mixed` | No | The time member hours starts on Sundays |
| `sunday_public_close` | `mixed` | No | The time public hours ends on Sundays |
| `sunday_public_open` | `mixed` | No | The time public hours starts on Sundays |
| `thursday_is_closed` | `mixed` | No | Whether the museum is closed on Thursdays |
| `thursday_member_close` | `mixed` | No | The time member hours ends on Thursdays |
| `thursday_member_open` | `mixed` | No | The time member hours starts on Thursdays |
| `thursday_public_close` | `mixed` | No | The time public hours ends on Thursdays |
| `thursday_public_open` | `mixed` | No | The time public hours starts on Thursdays |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `tuesday_is_closed` | `mixed` | No | Whether the museum is closed on Tuesdays |
| `tuesday_member_close` | `mixed` | No | The time member hours ends on Tuesdays |
| `tuesday_member_open` | `mixed` | No | The time member hours starts on Tuesdays |
| `tuesday_public_close` | `mixed` | No | The time public hours ends on Tuesdays |
| `tuesday_public_open` | `mixed` | No | The time public hours starts on Tuesdays |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |
| `wednesday_is_closed` | `mixed` | No | Whether the museum is closed on Wednesdays |
| `wednesday_member_close` | `mixed` | No | The time member hours ends on Wednesdays |
| `wednesday_member_open` | `mixed` | No | The time member hours starts on Wednesdays |
| `wednesday_public_close` | `mixed` | No | The time public hours ends on Wednesdays |
| `wednesday_public_open` | `mixed` | No | The time public hours starts on Wednesdays |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Hour()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Hour()->load(["id" => "hour_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): HourEntity`

Create a new `HourEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ImageEntity

```php
$image = $client->Image();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ahash` | `mixed` | No | Image hash generated using ahash algorithm with 64 boolean subfields |
| `alt_text` | `mixed` | No | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `artwork_ids` | `mixed` | No | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `mixed` | No | Names of the artworks associated with this asset |
| `color` | `mixed` | No | Dominant color of this image in HSL |
| `colorfulness` | `mixed` | No | Unbounded positive float representing an abstract measure of colorfulness. |
| `content` | `mixed` | No | Text of or URL to the contents of this asset |
| `content_e_tag` | `mixed` | No | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `mixed` | No | Asset-specific copyright information |
| `fingerprint` | `mixed` | No | Image hashes: aHash, dHash, pHash, wHash |
| `height` | `float` | No | Native height of the image |
| `id` | `string` | No | Unique identifier of this resource. |
| `iiif_url` | `mixed` | No | IIIF URL of this image |
| `is_educational_resource` | `bool` | No | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | No | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | No | Whether this resource is considered to be educational |
| `lake_guid` | `mixed` | No | Unique UUID of this resource in LAKE, our DAMS. |
| `lqip` | `mixed` | No | Low-quality image placeholder (LQIP). |
| `phash` | `mixed` | No | Image hash generated using phash algorithm with 64 boolean subfields |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `type` | `mixed` | No | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |
| `width` | `float` | No | Native width of the image |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Image()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Image()->load(["id" => "image_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ImageEntity`

Create a new `ImageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## LandingPageEntity

```php
$landing_page = $client->LandingPage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `copy` | `mixed` | No | The text of the page |
| `id` | `string` | No | Unique identifier of this resource. |
| `search_tags` | `mixed` | No | Editor-specified list of tags to aid in internal search |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | No | The URL to this page on our website |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->LandingPage()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->LandingPage()->load(["id" => "landing_page_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): LandingPageEntity`

Create a new `LandingPageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PlaceEntity

```php
$place = $client->Place();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `latitude` | `float` | No | Latitude coordinate of the center of the room |
| `longitude` | `float` | No | Longitude coordinate of the center of the room |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | `string` | No | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Place()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Place()->load(["id" => "place_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PlaceEntity`

Create a new `PlaceEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PressReleaseEntity

```php
$press_release = $client->PressRelease();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `copy` | `mixed` | No | The text of the page |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | No | The URL to this page on our website |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->PressRelease()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->PressRelease()->load(["id" => "press_release_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PressReleaseEntity`

Create a new `PressReleaseEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PrintedPublicationEntity

```php
$printed_publication = $client->PrintedPublication();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `copy` | `mixed` | No | The text of the page |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | No | The URL to this page on our website |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->PrintedPublication()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->PrintedPublication()->load(["id" => "printed_publication_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PrintedPublicationEntity`

Create a new `PrintedPublicationEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ProductEntity

```php
$product = $client->Product();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `artist_ids` | `mixed` | No | Unique identifiers of the artists associated with this product |
| `artwork_ids` | `mixed` | No | Unique identifiers of the artworks associated with this product |
| `description` | `string` | No | Explanation of what this product is |
| `exhibition_ids` | `mixed` | No | Unique identifiers of the exhibitions associated with this product |
| `external_sku` | `mixed` | No | Numeric product identification code of a machine-readable barcode, when the customer sku differs from our internal one |
| `id` | `string` | No | Unique identifier of this resource. |
| `image_url` | `mixed` | No | URL of an image for this product |
| `max_compare_at_price` | `mixed` | No | Number indicating how much the most expensive variant of a product cost before a sale |
| `max_current_price` | `mixed` | No | Number indicating how much the most expensive variant of a product costs right now |
| `min_compare_at_price` | `mixed` | No | Number indicating how much the least expensive variant of a product cost before a sale |
| `min_current_price` | `mixed` | No | Number indicating how much the least expensive variant of a product costs right now |
| `price_display` | `mixed` | No | Explanation of what this product is |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | No | URL of this product in the shop |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Product()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Product()->load(["id" => "product_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ProductEntity`

Create a new `ProductEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PublicationEntity

```php
$publication = $client->Publication();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `section_ids` | `mixed` | No | Unique identifiers of the sections of this publication |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | No | URL to the publication |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Publication()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Publication()->load(["id" => "publication_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PublicationEntity`

Create a new `PublicationEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SearchEntity

```php
$search = $client->Search();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_id` | `string` | No | API unique identifier |
| `api_link` | `mixed` | No | URL to this recource in the API |
| `api_model` | `mixed` | No | Name of the model the resource represents |
| `id` | `string` | No | Unique identifier within the search index |
| `is_boosted` | `bool` | No | Whether this record has been flagged to be boosted |
| `score` | `float` | No | Search index ranking of the result |
| `thumbnail` | `mixed` | No | Metadata on the image representing this record |
| `timestamp` | `mixed` | No | Date this record was last updated in the API |
| `title` | `string` | No | The name of this resource |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Search()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SearchEntity`

Create a new `SearchEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SectionEntity

```php
$section = $client->Section();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accession` | `mixed` | No | An accession number parsed from the title or tombstone |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `artwork_id` | `string` | No | Unique identifier of the artwork with which this section is associated |
| `content` | `mixed` | No | Content of this section in plaintext |
| `generic_page_id` | `string` | No | Unique identifier of the page on the website that represents the publication this section belongs to |
| `id` | `string` | No | Unique identifier of this resource. |
| `publication_id` | `string` | No | Unique identifier of the publication this section belongs to |
| `publication_title` | `mixed` | No | Name of the publication this section belongs to |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | No | URL to the section |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Section()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Section()->load(["id" => "section_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SectionEntity`

Create a new `SectionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SiteEntity

```php
$site = $client->Site();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `artwork_ids` | `mixed` | No | Unique identifiers of the artworks this site is associated with |
| `artwork_titles` | `mixed` | No | Names of the artworks this site is associated with |
| `description` | `string` | No | Explanation of what this site is |
| `exhibition_ids` | `mixed` | No | Unique identifier of the exhibitions this site is associated with |
| `exhibition_titles` | `mixed` | No | Names of the exhibitions this site is associated with |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | No | URL to this site |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Site()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Site()->load(["id" => "site_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SiteEntity`

Create a new `SiteEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SoundEntity

```php
$sound = $client->Sound();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_text` | `mixed` | No | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `artwork_ids` | `mixed` | No | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `mixed` | No | Names of the artworks associated with this asset |
| `content` | `mixed` | No | Text of or URL to the contents of this asset |
| `content_e_tag` | `mixed` | No | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `mixed` | No | Asset-specific copyright information |
| `id` | `string` | No | Unique identifier of this resource. |
| `is_educational_resource` | `bool` | No | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | No | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | No | Whether this resource is considered to be educational |
| `lake_guid` | `mixed` | No | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | Name of this mobile audio file – derived from the artwork and tour titles |
| `transcript` | `mixed` | No | Text transcription of the audio file |
| `type` | `mixed` | No | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | No | URL to the audio file |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Sound()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Sound()->load(["id" => "sound_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SoundEntity`

Create a new `SoundEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## StaticPageEntity

```php
$static_page = $client->StaticPage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `mixed` | No | The URL to this page on our website |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->StaticPage()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->StaticPage()->load(["id" => "static_page_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): StaticPageEntity`

Create a new `StaticPageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## TextEntity

```php
$text = $client->Text();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_text` | `mixed` | No | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `artwork_ids` | `mixed` | No | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `mixed` | No | Names of the artworks associated with this asset |
| `content` | `mixed` | No | Text of or URL to the contents of this asset |
| `content_e_tag` | `mixed` | No | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `mixed` | No | Asset-specific copyright information |
| `id` | `string` | No | Unique identifier of this resource. |
| `is_educational_resource` | `bool` | No | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | No | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | No | Whether this resource is considered to be educational |
| `lake_guid` | `mixed` | No | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `type` | `mixed` | No | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Text()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Text()->load(["id" => "text_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): TextEntity`

Create a new `TextEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## TourEntity

```php
$tour = $client->Tour();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `artist_titles` | `mixed` | No | Names of the artists of the artworks featured in this tour's tour stops |
| `artwork_titles` | `mixed` | No | Names of the artworks featured in this tour's tour stops |
| `description` | `string` | No | Explanation of what the tour is |
| `id` | `string` | No | Unique identifier of this resource. |
| `image` | `mixed` | No | The main image for the tour |
| `intro` | `mixed` | No | Text introducing the tour |
| `intro_link` | `mixed` | No | Link to the audio file of the introduction |
| `intro_transcript` | `mixed` | No | Transcript of the introduction audio to the tour |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |
| `weight` | `float` | No | Number representing this tour's sort order |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Tour()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Tour()->load(["id" => "tour_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): TourEntity`

Create a new `TourEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## VideoEntity

```php
$video = $client->Video();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_text` | `mixed` | No | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `mixed` | No | REST API link for this resource |
| `api_model` | `mixed` | No | REST API resource type or endpoint |
| `artwork_ids` | `mixed` | No | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `mixed` | No | Names of the artworks associated with this asset |
| `content` | `mixed` | No | Text of or URL to the contents of this asset |
| `content_e_tag` | `mixed` | No | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `mixed` | No | Asset-specific copyright information |
| `id` | `string` | No | Unique identifier of this resource. |
| `is_educational_resource` | `bool` | No | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | No | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | No | Whether this resource is considered to be educational |
| `lake_guid` | `mixed` | No | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `mixed` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `mixed` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `mixed` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `mixed` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `type` | `mixed` | No | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `mixed` | No | Date and time the record was updated in the aggregator database |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Video()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Video()->load(["id" => "video_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): VideoEntity`

Create a new `VideoEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```php
$client = new ArtInstituteOfChicagoSDK([
  "feature" => [
    "test" => ["active" => true],
  ],
]);
```

