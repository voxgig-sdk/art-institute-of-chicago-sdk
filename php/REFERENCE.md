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
| `alt_titles` | `mixed` | No |  |
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `birth_date` | `mixed` | No |  |
| `death_date` | `mixed` | No |  |
| `description` | `string` | No |  |
| `id` | `string` | No |  |
| `is_artist` | `bool` | No |  |
| `sort_title` | `mixed` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `ulan_id` | `string` | No |  |
| `updated_at` | `mixed` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `copy` | `mixed` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |

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
| `alt_artist_ids` | `mixed` | No |  |
| `alt_classification_ids` | `mixed` | No |  |
| `alt_image_ids` | `mixed` | No |  |
| `alt_material_ids` | `mixed` | No |  |
| `alt_style_ids` | `mixed` | No |  |
| `alt_subject_ids` | `mixed` | No |  |
| `alt_technique_ids` | `mixed` | No |  |
| `alt_titles` | `mixed` | No |  |
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `artist_display` | `mixed` | No |  |
| `artist_id` | `string` | No |  |
| `artist_ids` | `mixed` | No |  |
| `artist_title` | `mixed` | No |  |
| `artist_titles` | `mixed` | No |  |
| `artwork_type_id` | `string` | No |  |
| `artwork_type_title` | `mixed` | No |  |
| `boost_rank` | `mixed` | No |  |
| `catalog_based_search_keyword_titles` | `mixed` | No |  |
| `catalogue_display` | `mixed` | No |  |
| `category_ids` | `mixed` | No |  |
| `category_titles` | `mixed` | No |  |
| `classification_id` | `string` | No |  |
| `classification_ids` | `mixed` | No |  |
| `classification_title` | `mixed` | No |  |
| `classification_titles` | `mixed` | No |  |
| `color` | `mixed` | No |  |
| `colorfulness` | `mixed` | No |  |
| `copyright_notice` | `mixed` | No |  |
| `credit_line` | `mixed` | No |  |
| `date_display` | `mixed` | No |  |
| `date_end` | `mixed` | No |  |
| `date_qualifier_id` | `string` | No |  |
| `date_qualifier_title` | `mixed` | No |  |
| `date_start` | `mixed` | No |  |
| `department_id` | `string` | No |  |
| `department_title` | `mixed` | No |  |
| `description` | `string` | No |  |
| `dimensions` | `mixed` | No |  |
| `dimensions_detail` | `mixed` | No |  |
| `document_ids` | `mixed` | No |  |
| `edition` | `mixed` | No |  |
| `exhibition_history` | `mixed` | No |  |
| `fiscal_year` | `mixed` | No |  |
| `fiscal_year_deaccession` | `mixed` | No |  |
| `gallery_id` | `string` | No |  |
| `gallery_title` | `mixed` | No |  |
| `has_advanced_imaging` | `bool` | No |  |
| `has_educational_resources` | `bool` | No |  |
| `has_multimedia_resources` | `bool` | No |  |
| `has_not_been_viewed_much` | `bool` | No |  |
| `id` | `string` | No |  |
| `image_embedding` | `mixed` | No |  |
| `image_id` | `string` | No |  |
| `inscriptions` | `mixed` | No |  |
| `internal_department_id` | `string` | No |  |
| `is_boosted` | `bool` | No |  |
| `is_on_view` | `bool` | No |  |
| `is_public_domain` | `bool` | No |  |
| `is_zoomable` | `bool` | No |  |
| `latitude` | `float` | No |  |
| `latlon` | `mixed` | No |  |
| `longitude` | `float` | No |  |
| `main_reference_number` | `int` | No |  |
| `material_id` | `string` | No |  |
| `material_ids` | `mixed` | No |  |
| `material_titles` | `mixed` | No |  |
| `max_zoom_window_size` | `mixed` | No |  |
| `medium_display` | `mixed` | No |  |
| `nomisma_id` | `string` | No |  |
| `on_loan_display` | `mixed` | No |  |
| `pageviews` | `mixed` | No |  |
| `pageviews_recent` | `mixed` | No |  |
| `place_of_origin` | `mixed` | No |  |
| `provenance_text` | `mixed` | No |  |
| `publication_history` | `mixed` | No |  |
| `publishing_verification_level` | `mixed` | No |  |
| `section_ids` | `mixed` | No |  |
| `section_titles` | `mixed` | No |  |
| `short_description` | `mixed` | No |  |
| `site_ids` | `mixed` | No |  |
| `sound_ids` | `mixed` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `style_id` | `string` | No |  |
| `style_ids` | `mixed` | No |  |
| `style_title` | `mixed` | No |  |
| `style_titles` | `mixed` | No |  |
| `subject_id` | `string` | No |  |
| `subject_ids` | `mixed` | No |  |
| `subject_titles` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `technique_id` | `string` | No |  |
| `technique_ids` | `mixed` | No |  |
| `technique_titles` | `mixed` | No |  |
| `term_titles` | `mixed` | No |  |
| `text_embedding` | `mixed` | No |  |
| `text_ids` | `mixed` | No |  |
| `theme_titles` | `mixed` | No |  |
| `thumbnail` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |
| `video_ids` | `mixed` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |

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
| `aat_id` | `string` | No |  |
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |

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
| `aat_id` | `string` | No |  |
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `id` | `string` | No |  |
| `parent_id` | `string` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `subtype` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `copy` | `mixed` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |
| `web_url` | `mixed` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `author_display` | `mixed` | No |  |
| `copy` | `mixed` | No |  |
| `digital_publication_id` | `string` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |
| `web_url` | `mixed` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `copy` | `mixed` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |
| `web_url` | `mixed` | No |  |

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
| `alt_audience_ids` | `mixed` | No |  |
| `alt_event_type_ids` | `mixed` | No |  |
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `audience_id` | `string` | No |  |
| `buy_button_caption` | `mixed` | No |  |
| `buy_button_text` | `mixed` | No |  |
| `date_display` | `mixed` | No |  |
| `description` | `string` | No |  |
| `door_time` | `mixed` | No |  |
| `end_date` | `mixed` | No |  |
| `end_time` | `mixed` | No |  |
| `entrance` | `mixed` | No |  |
| `event_host_id` | `string` | No |  |
| `event_host_title` | `mixed` | No |  |
| `event_type_id` | `string` | No |  |
| `header_description` | `mixed` | No |  |
| `hero_caption` | `mixed` | No |  |
| `id` | `string` | No |  |
| `image_url` | `mixed` | No |  |
| `is_admission_required` | `bool` | No |  |
| `is_after_hours` | `bool` | No |  |
| `is_free` | `bool` | No |  |
| `is_member_exclusive` | `bool` | No |  |
| `is_private` | `bool` | No |  |
| `is_registration_required` | `bool` | No |  |
| `is_sales_button_hidden` | `bool` | No |  |
| `is_sold_out` | `bool` | No |  |
| `is_ticketed` | `bool` | No |  |
| `is_virtual_event` | `bool` | No |  |
| `join_url` | `mixed` | No |  |
| `layout_type` | `mixed` | No |  |
| `list_description` | `mixed` | No |  |
| `location` | `mixed` | No |  |
| `program_ids` | `mixed` | No |  |
| `program_titles` | `mixed` | No |  |
| `rsvp_link` | `mixed` | No |  |
| `search_tags` | `mixed` | No |  |
| `short_description` | `mixed` | No |  |
| `slug` | `string` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `start_date` | `mixed` | No |  |
| `start_time` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `survey_url` | `mixed` | No |  |
| `ticketed_event_id` | `string` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `title_display` | `mixed` | No |  |
| `updated_at` | `mixed` | No |  |
| `virtual_event_passcode` | `mixed` | No |  |
| `virtual_event_url` | `mixed` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `button_caption` | `mixed` | No |  |
| `button_text` | `mixed` | No |  |
| `button_url` | `mixed` | No |  |
| `description` | `string` | No |  |
| `end_at` | `mixed` | No |  |
| `event_id` | `string` | No |  |
| `id` | `string` | No |  |
| `image_url` | `mixed` | No |  |
| `is_private` | `bool` | No |  |
| `is_sales_button_hidden` | `bool` | No |  |
| `is_ticketed` | `bool` | No |  |
| `location` | `mixed` | No |  |
| `off_sale_at` | `mixed` | No |  |
| `on_sale_at` | `mixed` | No |  |
| `short_description` | `mixed` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `start_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `title_display` | `mixed` | No |  |
| `updated_at` | `mixed` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `id` | `string` | No |  |
| `is_affiliate_group` | `bool` | No |  |
| `is_event_host` | `bool` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |

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
| `aic_end_at` | `mixed` | No |  |
| `aic_start_at` | `mixed` | No |  |
| `alt_image_ids` | `mixed` | No |  |
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `artist_ids` | `mixed` | No |  |
| `artwork_ids` | `mixed` | No |  |
| `artwork_titles` | `mixed` | No |  |
| `document_ids` | `mixed` | No |  |
| `gallery_id` | `string` | No |  |
| `gallery_title` | `mixed` | No |  |
| `id` | `string` | No |  |
| `image_id` | `string` | No |  |
| `image_url` | `mixed` | No |  |
| `is_featured` | `bool` | No |  |
| `is_published` | `bool` | No |  |
| `position` | `mixed` | No |  |
| `short_description` | `mixed` | No |  |
| `site_ids` | `mixed` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `status` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |
| `web_url` | `mixed` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `floor` | `mixed` | No |  |
| `id` | `string` | No |  |
| `is_closed` | `bool` | No |  |
| `latitude` | `float` | No |  |
| `latlon` | `mixed` | No |  |
| `longitude` | `float` | No |  |
| `number` | `mixed` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `tgn_id` | `string` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `copy` | `mixed` | No |  |
| `id` | `string` | No |  |
| `search_tags` | `mixed` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |
| `web_url` | `mixed` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `copy` | `mixed` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |

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
| `additional_text` | `mixed` | No |  |
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `friday_is_closed` | `mixed` | No |  |
| `friday_member_close` | `mixed` | No |  |
| `friday_member_open` | `mixed` | No |  |
| `friday_public_close` | `mixed` | No |  |
| `friday_public_open` | `mixed` | No |  |
| `id` | `string` | No |  |
| `monday_is_closed` | `mixed` | No |  |
| `monday_member_close` | `mixed` | No |  |
| `monday_member_open` | `mixed` | No |  |
| `monday_public_close` | `mixed` | No |  |
| `monday_public_open` | `mixed` | No |  |
| `saturday_is_closed` | `mixed` | No |  |
| `saturday_member_close` | `mixed` | No |  |
| `saturday_member_open` | `mixed` | No |  |
| `saturday_public_close` | `mixed` | No |  |
| `saturday_public_open` | `mixed` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `summary` | `mixed` | No |  |
| `sunday_is_closed` | `mixed` | No |  |
| `sunday_member_close` | `mixed` | No |  |
| `sunday_member_open` | `mixed` | No |  |
| `sunday_public_close` | `mixed` | No |  |
| `sunday_public_open` | `mixed` | No |  |
| `thursday_is_closed` | `mixed` | No |  |
| `thursday_member_close` | `mixed` | No |  |
| `thursday_member_open` | `mixed` | No |  |
| `thursday_public_close` | `mixed` | No |  |
| `thursday_public_open` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `tuesday_is_closed` | `mixed` | No |  |
| `tuesday_member_close` | `mixed` | No |  |
| `tuesday_member_open` | `mixed` | No |  |
| `tuesday_public_close` | `mixed` | No |  |
| `tuesday_public_open` | `mixed` | No |  |
| `updated_at` | `mixed` | No |  |
| `wednesday_is_closed` | `mixed` | No |  |
| `wednesday_member_close` | `mixed` | No |  |
| `wednesday_member_open` | `mixed` | No |  |
| `wednesday_public_close` | `mixed` | No |  |
| `wednesday_public_open` | `mixed` | No |  |

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
| `ahash` | `mixed` | No |  |
| `alt_text` | `mixed` | No |  |
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `artwork_ids` | `mixed` | No |  |
| `artwork_titles` | `mixed` | No |  |
| `color` | `mixed` | No |  |
| `colorfulness` | `mixed` | No |  |
| `content` | `mixed` | No |  |
| `content_e_tag` | `mixed` | No |  |
| `credit_line` | `mixed` | No |  |
| `fingerprint` | `mixed` | No |  |
| `height` | `float` | No |  |
| `id` | `string` | No |  |
| `iiif_url` | `mixed` | No |  |
| `is_educational_resource` | `bool` | No |  |
| `is_multimedia_resource` | `bool` | No |  |
| `is_teacher_resource` | `bool` | No |  |
| `lake_guid` | `mixed` | No |  |
| `lqip` | `mixed` | No |  |
| `phash` | `mixed` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `type` | `mixed` | No |  |
| `updated_at` | `mixed` | No |  |
| `width` | `float` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `copy` | `mixed` | No |  |
| `id` | `string` | No |  |
| `search_tags` | `mixed` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |
| `web_url` | `mixed` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `id` | `string` | No |  |
| `latitude` | `float` | No |  |
| `longitude` | `float` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `tgn_id` | `string` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `copy` | `mixed` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |
| `web_url` | `mixed` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `copy` | `mixed` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |
| `web_url` | `mixed` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `artist_ids` | `mixed` | No |  |
| `artwork_ids` | `mixed` | No |  |
| `description` | `string` | No |  |
| `exhibition_ids` | `mixed` | No |  |
| `external_sku` | `mixed` | No |  |
| `id` | `string` | No |  |
| `image_url` | `mixed` | No |  |
| `max_compare_at_price` | `mixed` | No |  |
| `max_current_price` | `mixed` | No |  |
| `min_compare_at_price` | `mixed` | No |  |
| `min_current_price` | `mixed` | No |  |
| `price_display` | `mixed` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |
| `web_url` | `mixed` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `id` | `string` | No |  |
| `section_ids` | `mixed` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |
| `web_url` | `mixed` | No |  |

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
| `api_id` | `string` | No |  |
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `id` | `string` | No |  |
| `is_boosted` | `bool` | No |  |
| `score` | `float` | No |  |
| `thumbnail` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |

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
| `accession` | `mixed` | No |  |
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `artwork_id` | `string` | No |  |
| `content` | `mixed` | No |  |
| `generic_page_id` | `string` | No |  |
| `id` | `string` | No |  |
| `publication_id` | `string` | No |  |
| `publication_title` | `mixed` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |
| `web_url` | `mixed` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `artwork_ids` | `mixed` | No |  |
| `artwork_titles` | `mixed` | No |  |
| `description` | `string` | No |  |
| `exhibition_ids` | `mixed` | No |  |
| `exhibition_titles` | `mixed` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |
| `web_url` | `mixed` | No |  |

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
| `alt_text` | `mixed` | No |  |
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `artwork_ids` | `mixed` | No |  |
| `artwork_titles` | `mixed` | No |  |
| `content` | `mixed` | No |  |
| `content_e_tag` | `mixed` | No |  |
| `credit_line` | `mixed` | No |  |
| `id` | `string` | No |  |
| `is_educational_resource` | `bool` | No |  |
| `is_multimedia_resource` | `bool` | No |  |
| `is_teacher_resource` | `bool` | No |  |
| `lake_guid` | `mixed` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `transcript` | `mixed` | No |  |
| `type` | `mixed` | No |  |
| `updated_at` | `mixed` | No |  |
| `web_url` | `mixed` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |
| `web_url` | `mixed` | No |  |

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
| `alt_text` | `mixed` | No |  |
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `artwork_ids` | `mixed` | No |  |
| `artwork_titles` | `mixed` | No |  |
| `content` | `mixed` | No |  |
| `content_e_tag` | `mixed` | No |  |
| `credit_line` | `mixed` | No |  |
| `id` | `string` | No |  |
| `is_educational_resource` | `bool` | No |  |
| `is_multimedia_resource` | `bool` | No |  |
| `is_teacher_resource` | `bool` | No |  |
| `lake_guid` | `mixed` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `type` | `mixed` | No |  |
| `updated_at` | `mixed` | No |  |

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
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `artist_titles` | `mixed` | No |  |
| `artwork_titles` | `mixed` | No |  |
| `description` | `string` | No |  |
| `id` | `string` | No |  |
| `image` | `mixed` | No |  |
| `intro` | `mixed` | No |  |
| `intro_link` | `mixed` | No |  |
| `intro_transcript` | `mixed` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `mixed` | No |  |
| `weight` | `float` | No |  |

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
| `alt_text` | `mixed` | No |  |
| `api_link` | `mixed` | No |  |
| `api_model` | `mixed` | No |  |
| `artwork_ids` | `mixed` | No |  |
| `artwork_titles` | `mixed` | No |  |
| `content` | `mixed` | No |  |
| `content_e_tag` | `mixed` | No |  |
| `credit_line` | `mixed` | No |  |
| `id` | `string` | No |  |
| `is_educational_resource` | `bool` | No |  |
| `is_multimedia_resource` | `bool` | No |  |
| `is_teacher_resource` | `bool` | No |  |
| `lake_guid` | `mixed` | No |  |
| `source_updated_at` | `mixed` | No |  |
| `suggest_autocomplete_all` | `mixed` | No |  |
| `suggest_autocomplete_boosted` | `mixed` | No |  |
| `timestamp` | `mixed` | No |  |
| `title` | `string` | No |  |
| `type` | `mixed` | No |  |
| `updated_at` | `mixed` | No |  |

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

