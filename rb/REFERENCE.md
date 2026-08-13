# ArtInstituteOfChicago Ruby SDK Reference

Complete API reference for the ArtInstituteOfChicago Ruby SDK.


## ArtInstituteOfChicagoSDK

### Constructor

```ruby
require_relative 'ArtInstituteOfChicago_sdk'

client = ArtInstituteOfChicagoSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `ArtInstituteOfChicagoSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = ArtInstituteOfChicagoSDK.test
```


### Instance Methods

#### `Agent(data = nil)`

Create a new `Agent` entity instance. Pass `nil` for no initial data.

#### `AgentRole(data = nil)`

Create a new `AgentRole` entity instance. Pass `nil` for no initial data.

#### `AgentType(data = nil)`

Create a new `AgentType` entity instance. Pass `nil` for no initial data.

#### `Article(data = nil)`

Create a new `Article` entity instance. Pass `nil` for no initial data.

#### `Artwork(data = nil)`

Create a new `Artwork` entity instance. Pass `nil` for no initial data.

#### `ArtworkDateQualifier(data = nil)`

Create a new `ArtworkDateQualifier` entity instance. Pass `nil` for no initial data.

#### `ArtworkPlaceQualifier(data = nil)`

Create a new `ArtworkPlaceQualifier` entity instance. Pass `nil` for no initial data.

#### `ArtworkType(data = nil)`

Create a new `ArtworkType` entity instance. Pass `nil` for no initial data.

#### `CategoryTerm(data = nil)`

Create a new `CategoryTerm` entity instance. Pass `nil` for no initial data.

#### `DigitalPublication(data = nil)`

Create a new `DigitalPublication` entity instance. Pass `nil` for no initial data.

#### `DigitalPublicationArticle(data = nil)`

Create a new `DigitalPublicationArticle` entity instance. Pass `nil` for no initial data.

#### `EducatorResource(data = nil)`

Create a new `EducatorResource` entity instance. Pass `nil` for no initial data.

#### `Event(data = nil)`

Create a new `Event` entity instance. Pass `nil` for no initial data.

#### `EventOccurrence(data = nil)`

Create a new `EventOccurrence` entity instance. Pass `nil` for no initial data.

#### `EventProgram(data = nil)`

Create a new `EventProgram` entity instance. Pass `nil` for no initial data.

#### `Exhibition(data = nil)`

Create a new `Exhibition` entity instance. Pass `nil` for no initial data.

#### `Gallery(data = nil)`

Create a new `Gallery` entity instance. Pass `nil` for no initial data.

#### `GenericPage(data = nil)`

Create a new `GenericPage` entity instance. Pass `nil` for no initial data.

#### `Highlight(data = nil)`

Create a new `Highlight` entity instance. Pass `nil` for no initial data.

#### `Hour(data = nil)`

Create a new `Hour` entity instance. Pass `nil` for no initial data.

#### `Image(data = nil)`

Create a new `Image` entity instance. Pass `nil` for no initial data.

#### `LandingPage(data = nil)`

Create a new `LandingPage` entity instance. Pass `nil` for no initial data.

#### `Place(data = nil)`

Create a new `Place` entity instance. Pass `nil` for no initial data.

#### `PressRelease(data = nil)`

Create a new `PressRelease` entity instance. Pass `nil` for no initial data.

#### `PrintedPublication(data = nil)`

Create a new `PrintedPublication` entity instance. Pass `nil` for no initial data.

#### `Product(data = nil)`

Create a new `Product` entity instance. Pass `nil` for no initial data.

#### `Publication(data = nil)`

Create a new `Publication` entity instance. Pass `nil` for no initial data.

#### `Search(data = nil)`

Create a new `Search` entity instance. Pass `nil` for no initial data.

#### `Section(data = nil)`

Create a new `Section` entity instance. Pass `nil` for no initial data.

#### `Site(data = nil)`

Create a new `Site` entity instance. Pass `nil` for no initial data.

#### `Sound(data = nil)`

Create a new `Sound` entity instance. Pass `nil` for no initial data.

#### `StaticPage(data = nil)`

Create a new `StaticPage` entity instance. Pass `nil` for no initial data.

#### `Text(data = nil)`

Create a new `Text` entity instance. Pass `nil` for no initial data.

#### `Tour(data = nil)`

Create a new `Tour` entity instance. Pass `nil` for no initial data.

#### `Video(data = nil)`

Create a new `Video` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## AgentEntity

```ruby
agent = client.Agent
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_titles` | `Object` | No |  |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `birth_date` | `Object` | No |  |
| `death_date` | `Object` | No |  |
| `description` | `String` | No |  |
| `id` | `String` | No |  |
| `is_artist` | `Boolean` | No |  |
| `sort_title` | `Object` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `ulan_id` | `String` | No |  |
| `updated_at` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Agent.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Agent.load({ "id" => "agent_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `AgentEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## AgentRoleEntity

```ruby
agent_role = client.AgentRole
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `id` | `String` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.AgentRole.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.AgentRole.load({ "id" => "agent_role_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `AgentRoleEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## AgentTypeEntity

```ruby
agent_type = client.AgentType
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `id` | `String` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.AgentType.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.AgentType.load({ "id" => "agent_type_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `AgentTypeEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ArticleEntity

```ruby
article = client.Article
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `copy` | `Object` | No |  |
| `id` | `String` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Article.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Article.load({ "id" => "article_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ArticleEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ArtworkEntity

```ruby
artwork = client.Artwork
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_artist_ids` | `Object` | No |  |
| `alt_classification_ids` | `Object` | No |  |
| `alt_image_ids` | `Object` | No |  |
| `alt_material_ids` | `Object` | No |  |
| `alt_style_ids` | `Object` | No |  |
| `alt_subject_ids` | `Object` | No |  |
| `alt_technique_ids` | `Object` | No |  |
| `alt_titles` | `Object` | No |  |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `artist_display` | `Object` | No |  |
| `artist_id` | `String` | No |  |
| `artist_ids` | `Object` | No |  |
| `artist_title` | `Object` | No |  |
| `artist_titles` | `Object` | No |  |
| `artwork_type_id` | `String` | No |  |
| `artwork_type_title` | `Object` | No |  |
| `boost_rank` | `Object` | No |  |
| `catalog_based_search_keyword_titles` | `Object` | No |  |
| `catalogue_display` | `Object` | No |  |
| `category_ids` | `Object` | No |  |
| `category_titles` | `Object` | No |  |
| `classification_id` | `String` | No |  |
| `classification_ids` | `Object` | No |  |
| `classification_title` | `Object` | No |  |
| `classification_titles` | `Object` | No |  |
| `color` | `Object` | No |  |
| `colorfulness` | `Object` | No |  |
| `copyright_notice` | `Object` | No |  |
| `credit_line` | `Object` | No |  |
| `date_display` | `Object` | No |  |
| `date_end` | `Object` | No |  |
| `date_qualifier_id` | `String` | No |  |
| `date_qualifier_title` | `Object` | No |  |
| `date_start` | `Object` | No |  |
| `department_id` | `String` | No |  |
| `department_title` | `Object` | No |  |
| `description` | `String` | No |  |
| `dimensions` | `Object` | No |  |
| `dimensions_detail` | `Object` | No |  |
| `document_ids` | `Object` | No |  |
| `edition` | `Object` | No |  |
| `exhibition_history` | `Object` | No |  |
| `fiscal_year` | `Object` | No |  |
| `fiscal_year_deaccession` | `Object` | No |  |
| `gallery_id` | `String` | No |  |
| `gallery_title` | `Object` | No |  |
| `has_advanced_imaging` | `Boolean` | No |  |
| `has_educational_resources` | `Boolean` | No |  |
| `has_multimedia_resources` | `Boolean` | No |  |
| `has_not_been_viewed_much` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `image_embedding` | `Object` | No |  |
| `image_id` | `String` | No |  |
| `inscriptions` | `Object` | No |  |
| `internal_department_id` | `String` | No |  |
| `is_boosted` | `Boolean` | No |  |
| `is_on_view` | `Boolean` | No |  |
| `is_public_domain` | `Boolean` | No |  |
| `is_zoomable` | `Boolean` | No |  |
| `latitude` | `Float` | No |  |
| `latlon` | `Object` | No |  |
| `longitude` | `Float` | No |  |
| `main_reference_number` | `Integer` | No |  |
| `material_id` | `String` | No |  |
| `material_ids` | `Object` | No |  |
| `material_titles` | `Object` | No |  |
| `max_zoom_window_size` | `Object` | No |  |
| `medium_display` | `Object` | No |  |
| `nomisma_id` | `String` | No |  |
| `on_loan_display` | `Object` | No |  |
| `pageviews` | `Object` | No |  |
| `pageviews_recent` | `Object` | No |  |
| `place_of_origin` | `Object` | No |  |
| `provenance_text` | `Object` | No |  |
| `publication_history` | `Object` | No |  |
| `publishing_verification_level` | `Object` | No |  |
| `section_ids` | `Object` | No |  |
| `section_titles` | `Object` | No |  |
| `short_description` | `Object` | No |  |
| `site_ids` | `Object` | No |  |
| `sound_ids` | `Object` | No |  |
| `source_updated_at` | `Object` | No |  |
| `style_id` | `String` | No |  |
| `style_ids` | `Object` | No |  |
| `style_title` | `Object` | No |  |
| `style_titles` | `Object` | No |  |
| `subject_id` | `String` | No |  |
| `subject_ids` | `Object` | No |  |
| `subject_titles` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `technique_id` | `String` | No |  |
| `technique_ids` | `Object` | No |  |
| `technique_titles` | `Object` | No |  |
| `term_titles` | `Object` | No |  |
| `text_embedding` | `Object` | No |  |
| `text_ids` | `Object` | No |  |
| `theme_titles` | `Object` | No |  |
| `thumbnail` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |
| `video_ids` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Artwork.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Artwork.load({ "id" => "artwork_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ArtworkEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ArtworkDateQualifierEntity

```ruby
artwork_date_qualifier = client.ArtworkDateQualifier
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `id` | `String` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.ArtworkDateQualifier.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.ArtworkDateQualifier.load({ "id" => "artwork_date_qualifier_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ArtworkDateQualifierEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ArtworkPlaceQualifierEntity

```ruby
artwork_place_qualifier = client.ArtworkPlaceQualifier
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `id` | `String` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.ArtworkPlaceQualifier.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.ArtworkPlaceQualifier.load({ "id" => "artwork_place_qualifier_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ArtworkPlaceQualifierEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ArtworkTypeEntity

```ruby
artwork_type = client.ArtworkType
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `aat_id` | `String` | No |  |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `id` | `String` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.ArtworkType.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.ArtworkType.load({ "id" => "artwork_type_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ArtworkTypeEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CategoryTermEntity

```ruby
category_term = client.CategoryTerm
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `aat_id` | `String` | No |  |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `id` | `String` | No |  |
| `parent_id` | `String` | No |  |
| `source_updated_at` | `Object` | No |  |
| `subtype` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.CategoryTerm.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.CategoryTerm.load({ "id" => "category_term_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `CategoryTermEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## DigitalPublicationEntity

```ruby
digital_publication = client.DigitalPublication
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `copy` | `Object` | No |  |
| `id` | `String` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |
| `web_url` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.DigitalPublication.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.DigitalPublication.load({ "id" => "digital_publication_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `DigitalPublicationEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## DigitalPublicationArticleEntity

```ruby
digital_publication_article = client.DigitalPublicationArticle
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `author_display` | `Object` | No |  |
| `copy` | `Object` | No |  |
| `digital_publication_id` | `String` | No |  |
| `id` | `String` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |
| `web_url` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.DigitalPublicationArticle.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.DigitalPublicationArticle.load({ "id" => "digital_publication_article_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `DigitalPublicationArticleEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EducatorResourceEntity

```ruby
educator_resource = client.EducatorResource
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `copy` | `Object` | No |  |
| `id` | `String` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |
| `web_url` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.EducatorResource.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.EducatorResource.load({ "id" => "educator_resource_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EducatorResourceEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EventEntity

```ruby
event = client.Event
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_audience_ids` | `Object` | No |  |
| `alt_event_type_ids` | `Object` | No |  |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `audience_id` | `String` | No |  |
| `buy_button_caption` | `Object` | No |  |
| `buy_button_text` | `Object` | No |  |
| `date_display` | `Object` | No |  |
| `description` | `String` | No |  |
| `door_time` | `Object` | No |  |
| `end_date` | `Object` | No |  |
| `end_time` | `Object` | No |  |
| `entrance` | `Object` | No |  |
| `event_host_id` | `String` | No |  |
| `event_host_title` | `Object` | No |  |
| `event_type_id` | `String` | No |  |
| `header_description` | `Object` | No |  |
| `hero_caption` | `Object` | No |  |
| `id` | `String` | No |  |
| `image_url` | `Object` | No |  |
| `is_admission_required` | `Boolean` | No |  |
| `is_after_hours` | `Boolean` | No |  |
| `is_free` | `Boolean` | No |  |
| `is_member_exclusive` | `Boolean` | No |  |
| `is_private` | `Boolean` | No |  |
| `is_registration_required` | `Boolean` | No |  |
| `is_sales_button_hidden` | `Boolean` | No |  |
| `is_sold_out` | `Boolean` | No |  |
| `is_ticketed` | `Boolean` | No |  |
| `is_virtual_event` | `Boolean` | No |  |
| `join_url` | `Object` | No |  |
| `layout_type` | `Object` | No |  |
| `list_description` | `Object` | No |  |
| `location` | `Object` | No |  |
| `program_ids` | `Object` | No |  |
| `program_titles` | `Object` | No |  |
| `rsvp_link` | `Object` | No |  |
| `search_tags` | `Object` | No |  |
| `short_description` | `Object` | No |  |
| `slug` | `String` | No |  |
| `source_updated_at` | `Object` | No |  |
| `start_date` | `Object` | No |  |
| `start_time` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `survey_url` | `Object` | No |  |
| `ticketed_event_id` | `String` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `title_display` | `Object` | No |  |
| `updated_at` | `Object` | No |  |
| `virtual_event_passcode` | `Object` | No |  |
| `virtual_event_url` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Event.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Event.load({ "id" => "event_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EventEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EventOccurrenceEntity

```ruby
event_occurrence = client.EventOccurrence
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `button_caption` | `Object` | No |  |
| `button_text` | `Object` | No |  |
| `button_url` | `Object` | No |  |
| `description` | `String` | No |  |
| `end_at` | `Object` | No |  |
| `event_id` | `String` | No |  |
| `id` | `String` | No |  |
| `image_url` | `Object` | No |  |
| `is_private` | `Boolean` | No |  |
| `is_sales_button_hidden` | `Boolean` | No |  |
| `is_ticketed` | `Boolean` | No |  |
| `location` | `Object` | No |  |
| `off_sale_at` | `Object` | No |  |
| `on_sale_at` | `Object` | No |  |
| `short_description` | `Object` | No |  |
| `source_updated_at` | `Object` | No |  |
| `start_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `title_display` | `Object` | No |  |
| `updated_at` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.EventOccurrence.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.EventOccurrence.load({ "id" => "event_occurrence_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EventOccurrenceEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EventProgramEntity

```ruby
event_program = client.EventProgram
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `id` | `String` | No |  |
| `is_affiliate_group` | `Boolean` | No |  |
| `is_event_host` | `Boolean` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.EventProgram.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.EventProgram.load({ "id" => "event_program_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EventProgramEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ExhibitionEntity

```ruby
exhibition = client.Exhibition
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `aic_end_at` | `Object` | No |  |
| `aic_start_at` | `Object` | No |  |
| `alt_image_ids` | `Object` | No |  |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `artist_ids` | `Object` | No |  |
| `artwork_ids` | `Object` | No |  |
| `artwork_titles` | `Object` | No |  |
| `document_ids` | `Object` | No |  |
| `gallery_id` | `String` | No |  |
| `gallery_title` | `Object` | No |  |
| `id` | `String` | No |  |
| `image_id` | `String` | No |  |
| `image_url` | `Object` | No |  |
| `is_featured` | `Boolean` | No |  |
| `is_published` | `Boolean` | No |  |
| `position` | `Object` | No |  |
| `short_description` | `Object` | No |  |
| `site_ids` | `Object` | No |  |
| `source_updated_at` | `Object` | No |  |
| `status` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |
| `web_url` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Exhibition.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Exhibition.load({ "id" => "exhibition_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ExhibitionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## GalleryEntity

```ruby
gallery = client.Gallery
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `floor` | `Object` | No |  |
| `id` | `String` | No |  |
| `is_closed` | `Boolean` | No |  |
| `latitude` | `Float` | No |  |
| `latlon` | `Object` | No |  |
| `longitude` | `Float` | No |  |
| `number` | `Object` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `tgn_id` | `String` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Gallery.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Gallery.load({ "id" => "gallery_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `GalleryEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## GenericPageEntity

```ruby
generic_page = client.GenericPage
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `copy` | `Object` | No |  |
| `id` | `String` | No |  |
| `search_tags` | `Object` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |
| `web_url` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.GenericPage.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.GenericPage.load({ "id" => "generic_page_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `GenericPageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## HighlightEntity

```ruby
highlight = client.Highlight
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `copy` | `Object` | No |  |
| `id` | `String` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Highlight.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Highlight.load({ "id" => "highlight_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `HighlightEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## HourEntity

```ruby
hour = client.Hour
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_text` | `Object` | No |  |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `friday_is_closed` | `Object` | No |  |
| `friday_member_close` | `Object` | No |  |
| `friday_member_open` | `Object` | No |  |
| `friday_public_close` | `Object` | No |  |
| `friday_public_open` | `Object` | No |  |
| `id` | `String` | No |  |
| `monday_is_closed` | `Object` | No |  |
| `monday_member_close` | `Object` | No |  |
| `monday_member_open` | `Object` | No |  |
| `monday_public_close` | `Object` | No |  |
| `monday_public_open` | `Object` | No |  |
| `saturday_is_closed` | `Object` | No |  |
| `saturday_member_close` | `Object` | No |  |
| `saturday_member_open` | `Object` | No |  |
| `saturday_public_close` | `Object` | No |  |
| `saturday_public_open` | `Object` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `summary` | `Object` | No |  |
| `sunday_is_closed` | `Object` | No |  |
| `sunday_member_close` | `Object` | No |  |
| `sunday_member_open` | `Object` | No |  |
| `sunday_public_close` | `Object` | No |  |
| `sunday_public_open` | `Object` | No |  |
| `thursday_is_closed` | `Object` | No |  |
| `thursday_member_close` | `Object` | No |  |
| `thursday_member_open` | `Object` | No |  |
| `thursday_public_close` | `Object` | No |  |
| `thursday_public_open` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `tuesday_is_closed` | `Object` | No |  |
| `tuesday_member_close` | `Object` | No |  |
| `tuesday_member_open` | `Object` | No |  |
| `tuesday_public_close` | `Object` | No |  |
| `tuesday_public_open` | `Object` | No |  |
| `updated_at` | `Object` | No |  |
| `wednesday_is_closed` | `Object` | No |  |
| `wednesday_member_close` | `Object` | No |  |
| `wednesday_member_open` | `Object` | No |  |
| `wednesday_public_close` | `Object` | No |  |
| `wednesday_public_open` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Hour.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Hour.load({ "id" => "hour_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `HourEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ImageEntity

```ruby
image = client.Image
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ahash` | `Object` | No |  |
| `alt_text` | `Object` | No |  |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `artwork_ids` | `Object` | No |  |
| `artwork_titles` | `Object` | No |  |
| `color` | `Object` | No |  |
| `colorfulness` | `Object` | No |  |
| `content` | `Object` | No |  |
| `content_e_tag` | `Object` | No |  |
| `credit_line` | `Object` | No |  |
| `fingerprint` | `Object` | No |  |
| `height` | `Float` | No |  |
| `id` | `String` | No |  |
| `iiif_url` | `Object` | No |  |
| `is_educational_resource` | `Boolean` | No |  |
| `is_multimedia_resource` | `Boolean` | No |  |
| `is_teacher_resource` | `Boolean` | No |  |
| `lake_guid` | `Object` | No |  |
| `lqip` | `Object` | No |  |
| `phash` | `Object` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `type` | `Object` | No |  |
| `updated_at` | `Object` | No |  |
| `width` | `Float` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Image.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Image.load({ "id" => "image_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ImageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## LandingPageEntity

```ruby
landing_page = client.LandingPage
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `copy` | `Object` | No |  |
| `id` | `String` | No |  |
| `search_tags` | `Object` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |
| `web_url` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.LandingPage.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.LandingPage.load({ "id" => "landing_page_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `LandingPageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PlaceEntity

```ruby
place = client.Place
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `id` | `String` | No |  |
| `latitude` | `Float` | No |  |
| `longitude` | `Float` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `tgn_id` | `String` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Place.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Place.load({ "id" => "place_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PlaceEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PressReleaseEntity

```ruby
press_release = client.PressRelease
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `copy` | `Object` | No |  |
| `id` | `String` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |
| `web_url` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.PressRelease.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.PressRelease.load({ "id" => "press_release_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PressReleaseEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PrintedPublicationEntity

```ruby
printed_publication = client.PrintedPublication
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `copy` | `Object` | No |  |
| `id` | `String` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |
| `web_url` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.PrintedPublication.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.PrintedPublication.load({ "id" => "printed_publication_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PrintedPublicationEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ProductEntity

```ruby
product = client.Product
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `artist_ids` | `Object` | No |  |
| `artwork_ids` | `Object` | No |  |
| `description` | `String` | No |  |
| `exhibition_ids` | `Object` | No |  |
| `external_sku` | `Object` | No |  |
| `id` | `String` | No |  |
| `image_url` | `Object` | No |  |
| `max_compare_at_price` | `Object` | No |  |
| `max_current_price` | `Object` | No |  |
| `min_compare_at_price` | `Object` | No |  |
| `min_current_price` | `Object` | No |  |
| `price_display` | `Object` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |
| `web_url` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Product.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Product.load({ "id" => "product_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ProductEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PublicationEntity

```ruby
publication = client.Publication
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `id` | `String` | No |  |
| `section_ids` | `Object` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |
| `web_url` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Publication.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Publication.load({ "id" => "publication_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PublicationEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SearchEntity

```ruby
search = client.Search
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_id` | `String` | No |  |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `id` | `String` | No |  |
| `is_boosted` | `Boolean` | No |  |
| `score` | `Float` | No |  |
| `thumbnail` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Search.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SearchEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SectionEntity

```ruby
section = client.Section
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accession` | `Object` | No |  |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `artwork_id` | `String` | No |  |
| `content` | `Object` | No |  |
| `generic_page_id` | `String` | No |  |
| `id` | `String` | No |  |
| `publication_id` | `String` | No |  |
| `publication_title` | `Object` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |
| `web_url` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Section.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Section.load({ "id" => "section_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SectionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SiteEntity

```ruby
site = client.Site
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `artwork_ids` | `Object` | No |  |
| `artwork_titles` | `Object` | No |  |
| `description` | `String` | No |  |
| `exhibition_ids` | `Object` | No |  |
| `exhibition_titles` | `Object` | No |  |
| `id` | `String` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |
| `web_url` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Site.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Site.load({ "id" => "site_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SiteEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SoundEntity

```ruby
sound = client.Sound
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_text` | `Object` | No |  |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `artwork_ids` | `Object` | No |  |
| `artwork_titles` | `Object` | No |  |
| `content` | `Object` | No |  |
| `content_e_tag` | `Object` | No |  |
| `credit_line` | `Object` | No |  |
| `id` | `String` | No |  |
| `is_educational_resource` | `Boolean` | No |  |
| `is_multimedia_resource` | `Boolean` | No |  |
| `is_teacher_resource` | `Boolean` | No |  |
| `lake_guid` | `Object` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `transcript` | `Object` | No |  |
| `type` | `Object` | No |  |
| `updated_at` | `Object` | No |  |
| `web_url` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Sound.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Sound.load({ "id" => "sound_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SoundEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## StaticPageEntity

```ruby
static_page = client.StaticPage
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `id` | `String` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |
| `web_url` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.StaticPage.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.StaticPage.load({ "id" => "static_page_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `StaticPageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## TextEntity

```ruby
text = client.Text
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_text` | `Object` | No |  |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `artwork_ids` | `Object` | No |  |
| `artwork_titles` | `Object` | No |  |
| `content` | `Object` | No |  |
| `content_e_tag` | `Object` | No |  |
| `credit_line` | `Object` | No |  |
| `id` | `String` | No |  |
| `is_educational_resource` | `Boolean` | No |  |
| `is_multimedia_resource` | `Boolean` | No |  |
| `is_teacher_resource` | `Boolean` | No |  |
| `lake_guid` | `Object` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `type` | `Object` | No |  |
| `updated_at` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Text.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Text.load({ "id" => "text_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `TextEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## TourEntity

```ruby
tour = client.Tour
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `artist_titles` | `Object` | No |  |
| `artwork_titles` | `Object` | No |  |
| `description` | `String` | No |  |
| `id` | `String` | No |  |
| `image` | `Object` | No |  |
| `intro` | `Object` | No |  |
| `intro_link` | `Object` | No |  |
| `intro_transcript` | `Object` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `updated_at` | `Object` | No |  |
| `weight` | `Float` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Tour.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Tour.load({ "id" => "tour_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `TourEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## VideoEntity

```ruby
video = client.Video
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_text` | `Object` | No |  |
| `api_link` | `Object` | No |  |
| `api_model` | `Object` | No |  |
| `artwork_ids` | `Object` | No |  |
| `artwork_titles` | `Object` | No |  |
| `content` | `Object` | No |  |
| `content_e_tag` | `Object` | No |  |
| `credit_line` | `Object` | No |  |
| `id` | `String` | No |  |
| `is_educational_resource` | `Boolean` | No |  |
| `is_multimedia_resource` | `Boolean` | No |  |
| `is_teacher_resource` | `Boolean` | No |  |
| `lake_guid` | `Object` | No |  |
| `source_updated_at` | `Object` | No |  |
| `suggest_autocomplete_all` | `Object` | No |  |
| `suggest_autocomplete_boosted` | `Object` | No |  |
| `timestamp` | `Object` | No |  |
| `title` | `String` | No |  |
| `type` | `Object` | No |  |
| `updated_at` | `Object` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Video.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Video.load({ "id" => "video_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `VideoEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ruby
client = ArtInstituteOfChicagoSDK.new({
  "feature" => {
    "test" => { "active" => true },
  },
})
```

