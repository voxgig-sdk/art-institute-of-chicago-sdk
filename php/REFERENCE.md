# ArtInstituteOfChicago PHP SDK Reference

Complete API reference for the ArtInstituteOfChicago PHP SDK.


## ArtInstituteOfChicagoSDK

### Constructor

```php
require_once __DIR__ . '/art-institute-of-chicago_sdk.php';

$client = new ArtInstituteOfChicagoSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["apikey"]` | `string` | API key for authentication. |
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

#### `optionsMap(): array`

Return a deep copy of the current SDK options.

#### `getUtility(): ProjectNameUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. Returns `[$result, $err]`.

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

**Returns:** `array [$result, $err]`

#### `prepare(array $fetchargs = []): array`

Prepare a fetch definition without sending the request. Returns `[$fetchdef, $err]`.


---

## AgentEntity

```php
$agent = $client->Agent();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_title` | ``$ANY`` | No |  |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `birth_date` | ``$ANY`` | No |  |
| `death_date` | ``$ANY`` | No |  |
| `description` | ``$STRING`` | No |  |
| `id` | ``$STRING`` | No |  |
| `is_artist` | ``$BOOLEAN`` | No |  |
| `sort_title` | ``$ANY`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `ulan_id` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->Agent()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Agent()->load(["id" => "agent_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): AgentEntity`

Create a new `AgentEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## AgentRoleEntity

```php
$agent_role = $client->AgentRole();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->AgentRole()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->AgentRole()->load(["id" => "agent_role_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): AgentRoleEntity`

Create a new `AgentRoleEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## AgentTypeEntity

```php
$agent_type = $client->AgentType();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->AgentType()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->AgentType()->load(["id" => "agent_type_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): AgentTypeEntity`

Create a new `AgentTypeEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## ArticleEntity

```php
$article = $client->Article();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `copy` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->Article()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Article()->load(["id" => "article_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): ArticleEntity`

Create a new `ArticleEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## ArtworkEntity

```php
$artwork = $client->Artwork();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_artist_id` | ``$STRING`` | No |  |
| `alt_classification_id` | ``$STRING`` | No |  |
| `alt_image_id` | ``$STRING`` | No |  |
| `alt_material_id` | ``$STRING`` | No |  |
| `alt_style_id` | ``$STRING`` | No |  |
| `alt_subject_id` | ``$STRING`` | No |  |
| `alt_technique_id` | ``$STRING`` | No |  |
| `alt_title` | ``$ANY`` | No |  |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `artist_display` | ``$ANY`` | No |  |
| `artist_id` | ``$STRING`` | No |  |
| `artist_title` | ``$ANY`` | No |  |
| `artwork_type_id` | ``$STRING`` | No |  |
| `artwork_type_title` | ``$ANY`` | No |  |
| `boost_rank` | ``$ANY`` | No |  |
| `catalog_based_search_keyword_title` | ``$ANY`` | No |  |
| `catalogue_display` | ``$ANY`` | No |  |
| `category_id` | ``$STRING`` | No |  |
| `category_title` | ``$ANY`` | No |  |
| `classification_id` | ``$STRING`` | No |  |
| `classification_title` | ``$ANY`` | No |  |
| `color` | ``$ANY`` | No |  |
| `colorfulness` | ``$ANY`` | No |  |
| `copyright_notices` | ``$ANY`` | No |  |
| `credit_line` | ``$ANY`` | No |  |
| `date_display` | ``$ANY`` | No |  |
| `date_end` | ``$ANY`` | No |  |
| `date_qualifier_id` | ``$STRING`` | No |  |
| `date_qualifier_title` | ``$ANY`` | No |  |
| `date_start` | ``$ANY`` | No |  |
| `department_id` | ``$STRING`` | No |  |
| `department_title` | ``$ANY`` | No |  |
| `description` | ``$STRING`` | No |  |
| `dimension` | ``$ANY`` | No |  |
| `dimensions_detail` | ``$ANY`` | No |  |
| `document_id` | ``$STRING`` | No |  |
| `edition` | ``$ANY`` | No |  |
| `exhibition_history` | ``$ANY`` | No |  |
| `fiscal_year` | ``$ANY`` | No |  |
| `fiscal_year_deaccession` | ``$ANY`` | No |  |
| `gallery_id` | ``$STRING`` | No |  |
| `gallery_title` | ``$ANY`` | No |  |
| `has_advanced_imaging` | ``$BOOLEAN`` | No |  |
| `has_educational_resource` | ``$BOOLEAN`` | No |  |
| `has_multimedia_resource` | ``$BOOLEAN`` | No |  |
| `has_not_been_viewed_much` | ``$BOOLEAN`` | No |  |
| `id` | ``$STRING`` | No |  |
| `image_embedding` | ``$ANY`` | No |  |
| `image_id` | ``$STRING`` | No |  |
| `inscription` | ``$ANY`` | No |  |
| `internal_department_id` | ``$STRING`` | No |  |
| `is_boosted` | ``$BOOLEAN`` | No |  |
| `is_on_view` | ``$BOOLEAN`` | No |  |
| `is_public_domain` | ``$BOOLEAN`` | No |  |
| `is_zoomable` | ``$BOOLEAN`` | No |  |
| `latitude` | ``$NUMBER`` | No |  |
| `latlon` | ``$ANY`` | No |  |
| `longitude` | ``$NUMBER`` | No |  |
| `main_reference_number` | ``$INTEGER`` | No |  |
| `material_id` | ``$STRING`` | No |  |
| `material_title` | ``$ANY`` | No |  |
| `max_zoom_window_size` | ``$ANY`` | No |  |
| `medium_display` | ``$ANY`` | No |  |
| `nomisma_id` | ``$STRING`` | No |  |
| `on_loan_display` | ``$ANY`` | No |  |
| `pageview` | ``$ANY`` | No |  |
| `pageviews_recent` | ``$ANY`` | No |  |
| `place_of_origin` | ``$ANY`` | No |  |
| `provenance_text` | ``$ANY`` | No |  |
| `publication_history` | ``$ANY`` | No |  |
| `publishing_verification_level` | ``$ANY`` | No |  |
| `section_id` | ``$STRING`` | No |  |
| `section_title` | ``$ANY`` | No |  |
| `short_description` | ``$ANY`` | No |  |
| `site_id` | ``$STRING`` | No |  |
| `sound_id` | ``$STRING`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `style_id` | ``$STRING`` | No |  |
| `style_title` | ``$ANY`` | No |  |
| `subject_id` | ``$STRING`` | No |  |
| `subject_title` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `technique_id` | ``$STRING`` | No |  |
| `technique_title` | ``$ANY`` | No |  |
| `term_title` | ``$ANY`` | No |  |
| `text_embedding` | ``$ANY`` | No |  |
| `text_id` | ``$STRING`` | No |  |
| `theme_title` | ``$ANY`` | No |  |
| `thumbnail` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |
| `video_id` | ``$STRING`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->Artwork()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Artwork()->load(["id" => "artwork_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): ArtworkEntity`

Create a new `ArtworkEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## ArtworkDateQualifierEntity

```php
$artwork_date_qualifier = $client->ArtworkDateQualifier();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->ArtworkDateQualifier()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->ArtworkDateQualifier()->load(["id" => "artwork_date_qualifier_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): ArtworkDateQualifierEntity`

Create a new `ArtworkDateQualifierEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## ArtworkPlaceQualifierEntity

```php
$artwork_place_qualifier = $client->ArtworkPlaceQualifier();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->ArtworkPlaceQualifier()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->ArtworkPlaceQualifier()->load(["id" => "artwork_place_qualifier_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): ArtworkPlaceQualifierEntity`

Create a new `ArtworkPlaceQualifierEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## ArtworkTypeEntity

```php
$artwork_type = $client->ArtworkType();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `aat_id` | ``$STRING`` | No |  |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->ArtworkType()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->ArtworkType()->load(["id" => "artwork_type_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): ArtworkTypeEntity`

Create a new `ArtworkTypeEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## CategoryTermEntity

```php
$category_term = $client->CategoryTerm();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `aat_id` | ``$STRING`` | No |  |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `parent_id` | ``$STRING`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `subtype` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->CategoryTerm()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->CategoryTerm()->load(["id" => "category_term_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): CategoryTermEntity`

Create a new `CategoryTermEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## DigitalPublicationEntity

```php
$digital_publication = $client->DigitalPublication();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `copy` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |
| `web_url` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->DigitalPublication()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->DigitalPublication()->load(["id" => "digital_publication_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): DigitalPublicationEntity`

Create a new `DigitalPublicationEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## DigitalPublicationArticleEntity

```php
$digital_publication_article = $client->DigitalPublicationArticle();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `author_display` | ``$ANY`` | No |  |
| `copy` | ``$ANY`` | No |  |
| `digital_publication_id` | ``$STRING`` | No |  |
| `id` | ``$STRING`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |
| `web_url` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->DigitalPublicationArticle()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->DigitalPublicationArticle()->load(["id" => "digital_publication_article_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): DigitalPublicationArticleEntity`

Create a new `DigitalPublicationArticleEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## EducatorResourceEntity

```php
$educator_resource = $client->EducatorResource();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `copy` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |
| `web_url` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->EducatorResource()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->EducatorResource()->load(["id" => "educator_resource_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): EducatorResourceEntity`

Create a new `EducatorResourceEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## EventEntity

```php
$event = $client->Event();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_audience_id` | ``$STRING`` | No |  |
| `alt_event_type_id` | ``$STRING`` | No |  |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `audience_id` | ``$STRING`` | No |  |
| `buy_button_caption` | ``$ANY`` | No |  |
| `buy_button_text` | ``$ANY`` | No |  |
| `date_display` | ``$ANY`` | No |  |
| `description` | ``$STRING`` | No |  |
| `door_time` | ``$ANY`` | No |  |
| `end_date` | ``$ANY`` | No |  |
| `end_time` | ``$ANY`` | No |  |
| `entrance` | ``$ANY`` | No |  |
| `event_host_id` | ``$STRING`` | No |  |
| `event_host_title` | ``$ANY`` | No |  |
| `event_type_id` | ``$STRING`` | No |  |
| `header_description` | ``$ANY`` | No |  |
| `hero_caption` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `image_url` | ``$ANY`` | No |  |
| `is_admission_required` | ``$BOOLEAN`` | No |  |
| `is_after_hour` | ``$BOOLEAN`` | No |  |
| `is_free` | ``$BOOLEAN`` | No |  |
| `is_member_exclusive` | ``$BOOLEAN`` | No |  |
| `is_private` | ``$BOOLEAN`` | No |  |
| `is_registration_required` | ``$BOOLEAN`` | No |  |
| `is_sales_button_hidden` | ``$BOOLEAN`` | No |  |
| `is_sold_out` | ``$BOOLEAN`` | No |  |
| `is_ticketed` | ``$BOOLEAN`` | No |  |
| `is_virtual_event` | ``$BOOLEAN`` | No |  |
| `join_url` | ``$ANY`` | No |  |
| `layout_type` | ``$ANY`` | No |  |
| `list_description` | ``$ANY`` | No |  |
| `location` | ``$ANY`` | No |  |
| `program_id` | ``$STRING`` | No |  |
| `program_title` | ``$ANY`` | No |  |
| `rsvp_link` | ``$ANY`` | No |  |
| `search_tag` | ``$ANY`` | No |  |
| `short_description` | ``$ANY`` | No |  |
| `slug` | ``$STRING`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `start_date` | ``$ANY`` | No |  |
| `start_time` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `survey_url` | ``$ANY`` | No |  |
| `ticketed_event_id` | ``$STRING`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `title_display` | ``$ANY`` | No |  |
| `updated_at` | ``$ANY`` | No |  |
| `virtual_event_passcode` | ``$ANY`` | No |  |
| `virtual_event_url` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->Event()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Event()->load(["id" => "event_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): EventEntity`

Create a new `EventEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## EventOccurrenceEntity

```php
$event_occurrence = $client->EventOccurrence();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `button_caption` | ``$ANY`` | No |  |
| `button_text` | ``$ANY`` | No |  |
| `button_url` | ``$ANY`` | No |  |
| `description` | ``$STRING`` | No |  |
| `end_at` | ``$ANY`` | No |  |
| `event_id` | ``$STRING`` | No |  |
| `id` | ``$STRING`` | No |  |
| `image_url` | ``$ANY`` | No |  |
| `is_private` | ``$BOOLEAN`` | No |  |
| `is_sales_button_hidden` | ``$BOOLEAN`` | No |  |
| `is_ticketed` | ``$BOOLEAN`` | No |  |
| `location` | ``$ANY`` | No |  |
| `off_sale_at` | ``$ANY`` | No |  |
| `on_sale_at` | ``$ANY`` | No |  |
| `short_description` | ``$ANY`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `start_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `title_display` | ``$ANY`` | No |  |
| `updated_at` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->EventOccurrence()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->EventOccurrence()->load(["id" => "event_occurrence_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): EventOccurrenceEntity`

Create a new `EventOccurrenceEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## EventProgramEntity

```php
$event_program = $client->EventProgram();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `is_affiliate_group` | ``$BOOLEAN`` | No |  |
| `is_event_host` | ``$BOOLEAN`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->EventProgram()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->EventProgram()->load(["id" => "event_program_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): EventProgramEntity`

Create a new `EventProgramEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## ExhibitionEntity

```php
$exhibition = $client->Exhibition();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `aic_end_at` | ``$ANY`` | No |  |
| `aic_start_at` | ``$ANY`` | No |  |
| `alt_image_id` | ``$STRING`` | No |  |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `artist_id` | ``$STRING`` | No |  |
| `artwork_id` | ``$STRING`` | No |  |
| `artwork_title` | ``$ANY`` | No |  |
| `document_id` | ``$STRING`` | No |  |
| `gallery_id` | ``$STRING`` | No |  |
| `gallery_title` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `image_id` | ``$STRING`` | No |  |
| `image_url` | ``$ANY`` | No |  |
| `is_featured` | ``$BOOLEAN`` | No |  |
| `is_published` | ``$BOOLEAN`` | No |  |
| `position` | ``$ANY`` | No |  |
| `short_description` | ``$ANY`` | No |  |
| `site_id` | ``$STRING`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `status` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |
| `web_url` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->Exhibition()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Exhibition()->load(["id" => "exhibition_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): ExhibitionEntity`

Create a new `ExhibitionEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## GalleryEntity

```php
$gallery = $client->Gallery();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `floor` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `is_closed` | ``$BOOLEAN`` | No |  |
| `latitude` | ``$NUMBER`` | No |  |
| `latlon` | ``$ANY`` | No |  |
| `longitude` | ``$NUMBER`` | No |  |
| `number` | ``$ANY`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `tgn_id` | ``$STRING`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->Gallery()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Gallery()->load(["id" => "gallery_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): GalleryEntity`

Create a new `GalleryEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## GenericPageEntity

```php
$generic_page = $client->GenericPage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `copy` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `search_tag` | ``$ANY`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |
| `web_url` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->GenericPage()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->GenericPage()->load(["id" => "generic_page_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): GenericPageEntity`

Create a new `GenericPageEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## HighlightEntity

```php
$highlight = $client->Highlight();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `copy` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->Highlight()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Highlight()->load(["id" => "highlight_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): HighlightEntity`

Create a new `HighlightEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## HourEntity

```php
$hour = $client->Hour();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_text` | ``$ANY`` | No |  |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `friday_is_closed` | ``$ANY`` | No |  |
| `friday_member_close` | ``$ANY`` | No |  |
| `friday_member_open` | ``$ANY`` | No |  |
| `friday_public_close` | ``$ANY`` | No |  |
| `friday_public_open` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `monday_is_closed` | ``$ANY`` | No |  |
| `monday_member_close` | ``$ANY`` | No |  |
| `monday_member_open` | ``$ANY`` | No |  |
| `monday_public_close` | ``$ANY`` | No |  |
| `monday_public_open` | ``$ANY`` | No |  |
| `saturday_is_closed` | ``$ANY`` | No |  |
| `saturday_member_close` | ``$ANY`` | No |  |
| `saturday_member_open` | ``$ANY`` | No |  |
| `saturday_public_close` | ``$ANY`` | No |  |
| `saturday_public_open` | ``$ANY`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `summary` | ``$ANY`` | No |  |
| `sunday_is_closed` | ``$ANY`` | No |  |
| `sunday_member_close` | ``$ANY`` | No |  |
| `sunday_member_open` | ``$ANY`` | No |  |
| `sunday_public_close` | ``$ANY`` | No |  |
| `sunday_public_open` | ``$ANY`` | No |  |
| `thursday_is_closed` | ``$ANY`` | No |  |
| `thursday_member_close` | ``$ANY`` | No |  |
| `thursday_member_open` | ``$ANY`` | No |  |
| `thursday_public_close` | ``$ANY`` | No |  |
| `thursday_public_open` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `tuesday_is_closed` | ``$ANY`` | No |  |
| `tuesday_member_close` | ``$ANY`` | No |  |
| `tuesday_member_open` | ``$ANY`` | No |  |
| `tuesday_public_close` | ``$ANY`` | No |  |
| `tuesday_public_open` | ``$ANY`` | No |  |
| `updated_at` | ``$ANY`` | No |  |
| `wednesday_is_closed` | ``$ANY`` | No |  |
| `wednesday_member_close` | ``$ANY`` | No |  |
| `wednesday_member_open` | ``$ANY`` | No |  |
| `wednesday_public_close` | ``$ANY`` | No |  |
| `wednesday_public_open` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->Hour()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Hour()->load(["id" => "hour_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): HourEntity`

Create a new `HourEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## ImageEntity

```php
$image = $client->Image();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ahash` | ``$ANY`` | No |  |
| `alt_text` | ``$ANY`` | No |  |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `artwork_id` | ``$STRING`` | No |  |
| `artwork_title` | ``$ANY`` | No |  |
| `color` | ``$ANY`` | No |  |
| `colorfulness` | ``$ANY`` | No |  |
| `content` | ``$ANY`` | No |  |
| `content_e_tag` | ``$ANY`` | No |  |
| `credit_line` | ``$ANY`` | No |  |
| `fingerprint` | ``$ANY`` | No |  |
| `height` | ``$NUMBER`` | No |  |
| `id` | ``$STRING`` | No |  |
| `iiif_url` | ``$ANY`` | No |  |
| `is_educational_resource` | ``$BOOLEAN`` | No |  |
| `is_multimedia_resource` | ``$BOOLEAN`` | No |  |
| `is_teacher_resource` | ``$BOOLEAN`` | No |  |
| `lake_guid` | ``$ANY`` | No |  |
| `lqip` | ``$ANY`` | No |  |
| `phash` | ``$ANY`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `type` | ``$ANY`` | No |  |
| `updated_at` | ``$ANY`` | No |  |
| `width` | ``$NUMBER`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->Image()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Image()->load(["id" => "image_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): ImageEntity`

Create a new `ImageEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## LandingPageEntity

```php
$landing_page = $client->LandingPage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `copy` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `search_tag` | ``$ANY`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |
| `web_url` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->LandingPage()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->LandingPage()->load(["id" => "landing_page_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): LandingPageEntity`

Create a new `LandingPageEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## PlaceEntity

```php
$place = $client->Place();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `latitude` | ``$NUMBER`` | No |  |
| `longitude` | ``$NUMBER`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `tgn_id` | ``$STRING`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->Place()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Place()->load(["id" => "place_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): PlaceEntity`

Create a new `PlaceEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## PressReleaseEntity

```php
$press_release = $client->PressRelease();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `copy` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |
| `web_url` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->PressRelease()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->PressRelease()->load(["id" => "press_release_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): PressReleaseEntity`

Create a new `PressReleaseEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## PrintedPublicationEntity

```php
$printed_publication = $client->PrintedPublication();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `copy` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |
| `web_url` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->PrintedPublication()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->PrintedPublication()->load(["id" => "printed_publication_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): PrintedPublicationEntity`

Create a new `PrintedPublicationEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## ProductEntity

```php
$product = $client->Product();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `artist_id` | ``$STRING`` | No |  |
| `artwork_id` | ``$STRING`` | No |  |
| `description` | ``$STRING`` | No |  |
| `exhibition_id` | ``$STRING`` | No |  |
| `external_sku` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `image_url` | ``$ANY`` | No |  |
| `max_compare_at_price` | ``$ANY`` | No |  |
| `max_current_price` | ``$ANY`` | No |  |
| `min_compare_at_price` | ``$ANY`` | No |  |
| `min_current_price` | ``$ANY`` | No |  |
| `price_display` | ``$ANY`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |
| `web_url` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->Product()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Product()->load(["id" => "product_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): ProductEntity`

Create a new `ProductEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## PublicationEntity

```php
$publication = $client->Publication();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `section_id` | ``$STRING`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |
| `web_url` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->Publication()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Publication()->load(["id" => "publication_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): PublicationEntity`

Create a new `PublicationEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## SearchEntity

```php
$search = $client->Search();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_id` | ``$STRING`` | No |  |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `is_boosted` | ``$BOOLEAN`` | No |  |
| `score` | ``$NUMBER`` | No |  |
| `thumbnail` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->Search()->list([]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): SearchEntity`

Create a new `SearchEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## SectionEntity

```php
$section = $client->Section();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accession` | ``$ANY`` | No |  |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `artwork_id` | ``$STRING`` | No |  |
| `content` | ``$ANY`` | No |  |
| `generic_page_id` | ``$STRING`` | No |  |
| `id` | ``$STRING`` | No |  |
| `publication_id` | ``$STRING`` | No |  |
| `publication_title` | ``$ANY`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |
| `web_url` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->Section()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Section()->load(["id" => "section_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): SectionEntity`

Create a new `SectionEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## SiteEntity

```php
$site = $client->Site();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `artwork_id` | ``$STRING`` | No |  |
| `artwork_title` | ``$ANY`` | No |  |
| `description` | ``$STRING`` | No |  |
| `exhibition_id` | ``$STRING`` | No |  |
| `exhibition_title` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |
| `web_url` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->Site()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Site()->load(["id" => "site_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): SiteEntity`

Create a new `SiteEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## SoundEntity

```php
$sound = $client->Sound();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_text` | ``$ANY`` | No |  |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `artwork_id` | ``$STRING`` | No |  |
| `artwork_title` | ``$ANY`` | No |  |
| `content` | ``$ANY`` | No |  |
| `content_e_tag` | ``$ANY`` | No |  |
| `credit_line` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `is_educational_resource` | ``$BOOLEAN`` | No |  |
| `is_multimedia_resource` | ``$BOOLEAN`` | No |  |
| `is_teacher_resource` | ``$BOOLEAN`` | No |  |
| `lake_guid` | ``$ANY`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `transcript` | ``$ANY`` | No |  |
| `type` | ``$ANY`` | No |  |
| `updated_at` | ``$ANY`` | No |  |
| `web_url` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->Sound()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Sound()->load(["id" => "sound_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): SoundEntity`

Create a new `SoundEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## StaticPageEntity

```php
$static_page = $client->StaticPage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |
| `web_url` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->StaticPage()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->StaticPage()->load(["id" => "static_page_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): StaticPageEntity`

Create a new `StaticPageEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## TextEntity

```php
$text = $client->Text();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_text` | ``$ANY`` | No |  |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `artwork_id` | ``$STRING`` | No |  |
| `artwork_title` | ``$ANY`` | No |  |
| `content` | ``$ANY`` | No |  |
| `content_e_tag` | ``$ANY`` | No |  |
| `credit_line` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `is_educational_resource` | ``$BOOLEAN`` | No |  |
| `is_multimedia_resource` | ``$BOOLEAN`` | No |  |
| `is_teacher_resource` | ``$BOOLEAN`` | No |  |
| `lake_guid` | ``$ANY`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `type` | ``$ANY`` | No |  |
| `updated_at` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->Text()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Text()->load(["id" => "text_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): TextEntity`

Create a new `TextEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## TourEntity

```php
$tour = $client->Tour();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `artist_title` | ``$ANY`` | No |  |
| `artwork_title` | ``$ANY`` | No |  |
| `description` | ``$STRING`` | No |  |
| `id` | ``$STRING`` | No |  |
| `image` | ``$ANY`` | No |  |
| `intro` | ``$ANY`` | No |  |
| `intro_link` | ``$ANY`` | No |  |
| `intro_transcript` | ``$ANY`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `updated_at` | ``$ANY`` | No |  |
| `weight` | ``$NUMBER`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->Tour()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Tour()->load(["id" => "tour_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): TourEntity`

Create a new `TourEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## VideoEntity

```php
$video = $client->Video();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_text` | ``$ANY`` | No |  |
| `api_link` | ``$ANY`` | No |  |
| `api_model` | ``$ANY`` | No |  |
| `artwork_id` | ``$STRING`` | No |  |
| `artwork_title` | ``$ANY`` | No |  |
| `content` | ``$ANY`` | No |  |
| `content_e_tag` | ``$ANY`` | No |  |
| `credit_line` | ``$ANY`` | No |  |
| `id` | ``$STRING`` | No |  |
| `is_educational_resource` | ``$BOOLEAN`` | No |  |
| `is_multimedia_resource` | ``$BOOLEAN`` | No |  |
| `is_teacher_resource` | ``$BOOLEAN`` | No |  |
| `lake_guid` | ``$ANY`` | No |  |
| `source_updated_at` | ``$ANY`` | No |  |
| `suggest_autocomplete_all` | ``$ANY`` | No |  |
| `suggest_autocomplete_boosted` | ``$ANY`` | No |  |
| `timestamp` | ``$ANY`` | No |  |
| `title` | ``$STRING`` | No |  |
| `type` | ``$ANY`` | No |  |
| `updated_at` | ``$ANY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->Video()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->Video()->load(["id" => "video_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): VideoEntity`

Create a new `VideoEntity` instance with the same client and
options.

#### `getName(): string`

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

