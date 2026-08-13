# ArtInstituteOfChicago Lua SDK Reference

Complete API reference for the ArtInstituteOfChicago Lua SDK.


## ArtInstituteOfChicagoSDK

### Constructor

```lua
local sdk = require("art-institute-of-chicago_sdk")
local client = sdk.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `table` | SDK configuration options. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts?, sdkopts?)`

Create a test client with mock features active. Both arguments are optional.

```lua
local client = sdk.test()
```


### Instance Methods

#### `Agent(data)`

Create a new `Agent` entity instance. Pass `nil` for no initial data.

#### `AgentRole(data)`

Create a new `AgentRole` entity instance. Pass `nil` for no initial data.

#### `AgentType(data)`

Create a new `AgentType` entity instance. Pass `nil` for no initial data.

#### `Article(data)`

Create a new `Article` entity instance. Pass `nil` for no initial data.

#### `Artwork(data)`

Create a new `Artwork` entity instance. Pass `nil` for no initial data.

#### `ArtworkDateQualifier(data)`

Create a new `ArtworkDateQualifier` entity instance. Pass `nil` for no initial data.

#### `ArtworkPlaceQualifier(data)`

Create a new `ArtworkPlaceQualifier` entity instance. Pass `nil` for no initial data.

#### `ArtworkType(data)`

Create a new `ArtworkType` entity instance. Pass `nil` for no initial data.

#### `CategoryTerm(data)`

Create a new `CategoryTerm` entity instance. Pass `nil` for no initial data.

#### `DigitalPublication(data)`

Create a new `DigitalPublication` entity instance. Pass `nil` for no initial data.

#### `DigitalPublicationArticle(data)`

Create a new `DigitalPublicationArticle` entity instance. Pass `nil` for no initial data.

#### `EducatorResource(data)`

Create a new `EducatorResource` entity instance. Pass `nil` for no initial data.

#### `Event(data)`

Create a new `Event` entity instance. Pass `nil` for no initial data.

#### `EventOccurrence(data)`

Create a new `EventOccurrence` entity instance. Pass `nil` for no initial data.

#### `EventProgram(data)`

Create a new `EventProgram` entity instance. Pass `nil` for no initial data.

#### `Exhibition(data)`

Create a new `Exhibition` entity instance. Pass `nil` for no initial data.

#### `Gallery(data)`

Create a new `Gallery` entity instance. Pass `nil` for no initial data.

#### `GenericPage(data)`

Create a new `GenericPage` entity instance. Pass `nil` for no initial data.

#### `Highlight(data)`

Create a new `Highlight` entity instance. Pass `nil` for no initial data.

#### `Hour(data)`

Create a new `Hour` entity instance. Pass `nil` for no initial data.

#### `Image(data)`

Create a new `Image` entity instance. Pass `nil` for no initial data.

#### `LandingPage(data)`

Create a new `LandingPage` entity instance. Pass `nil` for no initial data.

#### `Place(data)`

Create a new `Place` entity instance. Pass `nil` for no initial data.

#### `PressRelease(data)`

Create a new `PressRelease` entity instance. Pass `nil` for no initial data.

#### `PrintedPublication(data)`

Create a new `PrintedPublication` entity instance. Pass `nil` for no initial data.

#### `Product(data)`

Create a new `Product` entity instance. Pass `nil` for no initial data.

#### `Publication(data)`

Create a new `Publication` entity instance. Pass `nil` for no initial data.

#### `Search(data)`

Create a new `Search` entity instance. Pass `nil` for no initial data.

#### `Section(data)`

Create a new `Section` entity instance. Pass `nil` for no initial data.

#### `Site(data)`

Create a new `Site` entity instance. Pass `nil` for no initial data.

#### `Sound(data)`

Create a new `Sound` entity instance. Pass `nil` for no initial data.

#### `StaticPage(data)`

Create a new `StaticPage` entity instance. Pass `nil` for no initial data.

#### `Text(data)`

Create a new `Text` entity instance. Pass `nil` for no initial data.

#### `Tour(data)`

Create a new `Tour` entity instance. Pass `nil` for no initial data.

#### `Video(data)`

Create a new `Video` entity instance. Pass `nil` for no initial data.

#### `options_map() -> table`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs) -> table, err`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs.params` | `table` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `table` | Query string parameters. |
| `fetchargs.headers` | `table` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (tables are JSON-serialized). |
| `fetchargs.ctrl` | `table` | Control options (e.g. `{ explain = true }`). |

**Returns:** `table, err`

#### `prepare(fetchargs) -> table, err`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `table, err`


---

## AgentEntity

```lua
local agent = client:Agent(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_titles` | `any` | No |  |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `birth_date` | `any` | No |  |
| `death_date` | `any` | No |  |
| `description` | `string` | No |  |
| `id` | `string` | No |  |
| `is_artist` | `boolean` | No |  |
| `sort_title` | `any` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `ulan_id` | `string` | No |  |
| `updated_at` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Agent():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Agent():load({ id = "agent_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AgentEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## AgentRoleEntity

```lua
local agent_role = client:AgentRole(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:AgentRole():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:AgentRole():load({ id = "agent_role_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AgentRoleEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## AgentTypeEntity

```lua
local agent_type = client:AgentType(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:AgentType():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:AgentType():load({ id = "agent_type_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AgentTypeEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ArticleEntity

```lua
local article = client:Article(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `copy` | `any` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Article():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Article():load({ id = "article_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ArticleEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ArtworkEntity

```lua
local artwork = client:Artwork(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_artist_ids` | `any` | No |  |
| `alt_classification_ids` | `any` | No |  |
| `alt_image_ids` | `any` | No |  |
| `alt_material_ids` | `any` | No |  |
| `alt_style_ids` | `any` | No |  |
| `alt_subject_ids` | `any` | No |  |
| `alt_technique_ids` | `any` | No |  |
| `alt_titles` | `any` | No |  |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `artist_display` | `any` | No |  |
| `artist_id` | `string` | No |  |
| `artist_ids` | `any` | No |  |
| `artist_title` | `any` | No |  |
| `artist_titles` | `any` | No |  |
| `artwork_type_id` | `string` | No |  |
| `artwork_type_title` | `any` | No |  |
| `boost_rank` | `any` | No |  |
| `catalog_based_search_keyword_titles` | `any` | No |  |
| `catalogue_display` | `any` | No |  |
| `category_ids` | `any` | No |  |
| `category_titles` | `any` | No |  |
| `classification_id` | `string` | No |  |
| `classification_ids` | `any` | No |  |
| `classification_title` | `any` | No |  |
| `classification_titles` | `any` | No |  |
| `color` | `any` | No |  |
| `colorfulness` | `any` | No |  |
| `copyright_notice` | `any` | No |  |
| `credit_line` | `any` | No |  |
| `date_display` | `any` | No |  |
| `date_end` | `any` | No |  |
| `date_qualifier_id` | `string` | No |  |
| `date_qualifier_title` | `any` | No |  |
| `date_start` | `any` | No |  |
| `department_id` | `string` | No |  |
| `department_title` | `any` | No |  |
| `description` | `string` | No |  |
| `dimensions` | `any` | No |  |
| `dimensions_detail` | `any` | No |  |
| `document_ids` | `any` | No |  |
| `edition` | `any` | No |  |
| `exhibition_history` | `any` | No |  |
| `fiscal_year` | `any` | No |  |
| `fiscal_year_deaccession` | `any` | No |  |
| `gallery_id` | `string` | No |  |
| `gallery_title` | `any` | No |  |
| `has_advanced_imaging` | `boolean` | No |  |
| `has_educational_resources` | `boolean` | No |  |
| `has_multimedia_resources` | `boolean` | No |  |
| `has_not_been_viewed_much` | `boolean` | No |  |
| `id` | `string` | No |  |
| `image_embedding` | `any` | No |  |
| `image_id` | `string` | No |  |
| `inscriptions` | `any` | No |  |
| `internal_department_id` | `string` | No |  |
| `is_boosted` | `boolean` | No |  |
| `is_on_view` | `boolean` | No |  |
| `is_public_domain` | `boolean` | No |  |
| `is_zoomable` | `boolean` | No |  |
| `latitude` | `number` | No |  |
| `latlon` | `any` | No |  |
| `longitude` | `number` | No |  |
| `main_reference_number` | `number` | No |  |
| `material_id` | `string` | No |  |
| `material_ids` | `any` | No |  |
| `material_titles` | `any` | No |  |
| `max_zoom_window_size` | `any` | No |  |
| `medium_display` | `any` | No |  |
| `nomisma_id` | `string` | No |  |
| `on_loan_display` | `any` | No |  |
| `pageviews` | `any` | No |  |
| `pageviews_recent` | `any` | No |  |
| `place_of_origin` | `any` | No |  |
| `provenance_text` | `any` | No |  |
| `publication_history` | `any` | No |  |
| `publishing_verification_level` | `any` | No |  |
| `section_ids` | `any` | No |  |
| `section_titles` | `any` | No |  |
| `short_description` | `any` | No |  |
| `site_ids` | `any` | No |  |
| `sound_ids` | `any` | No |  |
| `source_updated_at` | `any` | No |  |
| `style_id` | `string` | No |  |
| `style_ids` | `any` | No |  |
| `style_title` | `any` | No |  |
| `style_titles` | `any` | No |  |
| `subject_id` | `string` | No |  |
| `subject_ids` | `any` | No |  |
| `subject_titles` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `technique_id` | `string` | No |  |
| `technique_ids` | `any` | No |  |
| `technique_titles` | `any` | No |  |
| `term_titles` | `any` | No |  |
| `text_embedding` | `any` | No |  |
| `text_ids` | `any` | No |  |
| `theme_titles` | `any` | No |  |
| `thumbnail` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |
| `video_ids` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Artwork():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Artwork():load({ id = "artwork_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ArtworkEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ArtworkDateQualifierEntity

```lua
local artwork_date_qualifier = client:ArtworkDateQualifier(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:ArtworkDateQualifier():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:ArtworkDateQualifier():load({ id = "artwork_date_qualifier_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ArtworkDateQualifierEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ArtworkPlaceQualifierEntity

```lua
local artwork_place_qualifier = client:ArtworkPlaceQualifier(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:ArtworkPlaceQualifier():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:ArtworkPlaceQualifier():load({ id = "artwork_place_qualifier_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ArtworkPlaceQualifierEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ArtworkTypeEntity

```lua
local artwork_type = client:ArtworkType(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `aat_id` | `string` | No |  |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:ArtworkType():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:ArtworkType():load({ id = "artwork_type_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ArtworkTypeEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CategoryTermEntity

```lua
local category_term = client:CategoryTerm(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `aat_id` | `string` | No |  |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `id` | `string` | No |  |
| `parent_id` | `string` | No |  |
| `source_updated_at` | `any` | No |  |
| `subtype` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:CategoryTerm():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:CategoryTerm():load({ id = "category_term_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CategoryTermEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## DigitalPublicationEntity

```lua
local digital_publication = client:DigitalPublication(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `copy` | `any` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |
| `web_url` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:DigitalPublication():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:DigitalPublication():load({ id = "digital_publication_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DigitalPublicationEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## DigitalPublicationArticleEntity

```lua
local digital_publication_article = client:DigitalPublicationArticle(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `author_display` | `any` | No |  |
| `copy` | `any` | No |  |
| `digital_publication_id` | `string` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |
| `web_url` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:DigitalPublicationArticle():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:DigitalPublicationArticle():load({ id = "digital_publication_article_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DigitalPublicationArticleEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EducatorResourceEntity

```lua
local educator_resource = client:EducatorResource(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `copy` | `any` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |
| `web_url` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:EducatorResource():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EducatorResource():load({ id = "educator_resource_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EducatorResourceEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EventEntity

```lua
local event = client:Event(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_audience_ids` | `any` | No |  |
| `alt_event_type_ids` | `any` | No |  |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `audience_id` | `string` | No |  |
| `buy_button_caption` | `any` | No |  |
| `buy_button_text` | `any` | No |  |
| `date_display` | `any` | No |  |
| `description` | `string` | No |  |
| `door_time` | `any` | No |  |
| `end_date` | `any` | No |  |
| `end_time` | `any` | No |  |
| `entrance` | `any` | No |  |
| `event_host_id` | `string` | No |  |
| `event_host_title` | `any` | No |  |
| `event_type_id` | `string` | No |  |
| `header_description` | `any` | No |  |
| `hero_caption` | `any` | No |  |
| `id` | `string` | No |  |
| `image_url` | `any` | No |  |
| `is_admission_required` | `boolean` | No |  |
| `is_after_hours` | `boolean` | No |  |
| `is_free` | `boolean` | No |  |
| `is_member_exclusive` | `boolean` | No |  |
| `is_private` | `boolean` | No |  |
| `is_registration_required` | `boolean` | No |  |
| `is_sales_button_hidden` | `boolean` | No |  |
| `is_sold_out` | `boolean` | No |  |
| `is_ticketed` | `boolean` | No |  |
| `is_virtual_event` | `boolean` | No |  |
| `join_url` | `any` | No |  |
| `layout_type` | `any` | No |  |
| `list_description` | `any` | No |  |
| `location` | `any` | No |  |
| `program_ids` | `any` | No |  |
| `program_titles` | `any` | No |  |
| `rsvp_link` | `any` | No |  |
| `search_tags` | `any` | No |  |
| `short_description` | `any` | No |  |
| `slug` | `string` | No |  |
| `source_updated_at` | `any` | No |  |
| `start_date` | `any` | No |  |
| `start_time` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `survey_url` | `any` | No |  |
| `ticketed_event_id` | `string` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `title_display` | `any` | No |  |
| `updated_at` | `any` | No |  |
| `virtual_event_passcode` | `any` | No |  |
| `virtual_event_url` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Event():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Event():load({ id = "event_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EventEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EventOccurrenceEntity

```lua
local event_occurrence = client:EventOccurrence(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `button_caption` | `any` | No |  |
| `button_text` | `any` | No |  |
| `button_url` | `any` | No |  |
| `description` | `string` | No |  |
| `end_at` | `any` | No |  |
| `event_id` | `string` | No |  |
| `id` | `string` | No |  |
| `image_url` | `any` | No |  |
| `is_private` | `boolean` | No |  |
| `is_sales_button_hidden` | `boolean` | No |  |
| `is_ticketed` | `boolean` | No |  |
| `location` | `any` | No |  |
| `off_sale_at` | `any` | No |  |
| `on_sale_at` | `any` | No |  |
| `short_description` | `any` | No |  |
| `source_updated_at` | `any` | No |  |
| `start_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `title_display` | `any` | No |  |
| `updated_at` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:EventOccurrence():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EventOccurrence():load({ id = "event_occurrence_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EventOccurrenceEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EventProgramEntity

```lua
local event_program = client:EventProgram(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `id` | `string` | No |  |
| `is_affiliate_group` | `boolean` | No |  |
| `is_event_host` | `boolean` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:EventProgram():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EventProgram():load({ id = "event_program_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EventProgramEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ExhibitionEntity

```lua
local exhibition = client:Exhibition(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `aic_end_at` | `any` | No |  |
| `aic_start_at` | `any` | No |  |
| `alt_image_ids` | `any` | No |  |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `artist_ids` | `any` | No |  |
| `artwork_ids` | `any` | No |  |
| `artwork_titles` | `any` | No |  |
| `document_ids` | `any` | No |  |
| `gallery_id` | `string` | No |  |
| `gallery_title` | `any` | No |  |
| `id` | `string` | No |  |
| `image_id` | `string` | No |  |
| `image_url` | `any` | No |  |
| `is_featured` | `boolean` | No |  |
| `is_published` | `boolean` | No |  |
| `position` | `any` | No |  |
| `short_description` | `any` | No |  |
| `site_ids` | `any` | No |  |
| `source_updated_at` | `any` | No |  |
| `status` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |
| `web_url` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Exhibition():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Exhibition():load({ id = "exhibition_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ExhibitionEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## GalleryEntity

```lua
local gallery = client:Gallery(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `floor` | `any` | No |  |
| `id` | `string` | No |  |
| `is_closed` | `boolean` | No |  |
| `latitude` | `number` | No |  |
| `latlon` | `any` | No |  |
| `longitude` | `number` | No |  |
| `number` | `any` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `tgn_id` | `string` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Gallery():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Gallery():load({ id = "gallery_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GalleryEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## GenericPageEntity

```lua
local generic_page = client:GenericPage(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `copy` | `any` | No |  |
| `id` | `string` | No |  |
| `search_tags` | `any` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |
| `web_url` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:GenericPage():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:GenericPage():load({ id = "generic_page_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GenericPageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## HighlightEntity

```lua
local highlight = client:Highlight(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `copy` | `any` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Highlight():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Highlight():load({ id = "highlight_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `HighlightEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## HourEntity

```lua
local hour = client:Hour(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_text` | `any` | No |  |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `friday_is_closed` | `any` | No |  |
| `friday_member_close` | `any` | No |  |
| `friday_member_open` | `any` | No |  |
| `friday_public_close` | `any` | No |  |
| `friday_public_open` | `any` | No |  |
| `id` | `string` | No |  |
| `monday_is_closed` | `any` | No |  |
| `monday_member_close` | `any` | No |  |
| `monday_member_open` | `any` | No |  |
| `monday_public_close` | `any` | No |  |
| `monday_public_open` | `any` | No |  |
| `saturday_is_closed` | `any` | No |  |
| `saturday_member_close` | `any` | No |  |
| `saturday_member_open` | `any` | No |  |
| `saturday_public_close` | `any` | No |  |
| `saturday_public_open` | `any` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `summary` | `any` | No |  |
| `sunday_is_closed` | `any` | No |  |
| `sunday_member_close` | `any` | No |  |
| `sunday_member_open` | `any` | No |  |
| `sunday_public_close` | `any` | No |  |
| `sunday_public_open` | `any` | No |  |
| `thursday_is_closed` | `any` | No |  |
| `thursday_member_close` | `any` | No |  |
| `thursday_member_open` | `any` | No |  |
| `thursday_public_close` | `any` | No |  |
| `thursday_public_open` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `tuesday_is_closed` | `any` | No |  |
| `tuesday_member_close` | `any` | No |  |
| `tuesday_member_open` | `any` | No |  |
| `tuesday_public_close` | `any` | No |  |
| `tuesday_public_open` | `any` | No |  |
| `updated_at` | `any` | No |  |
| `wednesday_is_closed` | `any` | No |  |
| `wednesday_member_close` | `any` | No |  |
| `wednesday_member_open` | `any` | No |  |
| `wednesday_public_close` | `any` | No |  |
| `wednesday_public_open` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Hour():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Hour():load({ id = "hour_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `HourEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ImageEntity

```lua
local image = client:Image(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ahash` | `any` | No |  |
| `alt_text` | `any` | No |  |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `artwork_ids` | `any` | No |  |
| `artwork_titles` | `any` | No |  |
| `color` | `any` | No |  |
| `colorfulness` | `any` | No |  |
| `content` | `any` | No |  |
| `content_e_tag` | `any` | No |  |
| `credit_line` | `any` | No |  |
| `fingerprint` | `any` | No |  |
| `height` | `number` | No |  |
| `id` | `string` | No |  |
| `iiif_url` | `any` | No |  |
| `is_educational_resource` | `boolean` | No |  |
| `is_multimedia_resource` | `boolean` | No |  |
| `is_teacher_resource` | `boolean` | No |  |
| `lake_guid` | `any` | No |  |
| `lqip` | `any` | No |  |
| `phash` | `any` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `type` | `any` | No |  |
| `updated_at` | `any` | No |  |
| `width` | `number` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Image():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Image():load({ id = "image_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ImageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## LandingPageEntity

```lua
local landing_page = client:LandingPage(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `copy` | `any` | No |  |
| `id` | `string` | No |  |
| `search_tags` | `any` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |
| `web_url` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:LandingPage():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:LandingPage():load({ id = "landing_page_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `LandingPageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PlaceEntity

```lua
local place = client:Place(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `id` | `string` | No |  |
| `latitude` | `number` | No |  |
| `longitude` | `number` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `tgn_id` | `string` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Place():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Place():load({ id = "place_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PlaceEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PressReleaseEntity

```lua
local press_release = client:PressRelease(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `copy` | `any` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |
| `web_url` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:PressRelease():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:PressRelease():load({ id = "press_release_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PressReleaseEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PrintedPublicationEntity

```lua
local printed_publication = client:PrintedPublication(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `copy` | `any` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |
| `web_url` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:PrintedPublication():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:PrintedPublication():load({ id = "printed_publication_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PrintedPublicationEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ProductEntity

```lua
local product = client:Product(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `artist_ids` | `any` | No |  |
| `artwork_ids` | `any` | No |  |
| `description` | `string` | No |  |
| `exhibition_ids` | `any` | No |  |
| `external_sku` | `any` | No |  |
| `id` | `string` | No |  |
| `image_url` | `any` | No |  |
| `max_compare_at_price` | `any` | No |  |
| `max_current_price` | `any` | No |  |
| `min_compare_at_price` | `any` | No |  |
| `min_current_price` | `any` | No |  |
| `price_display` | `any` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |
| `web_url` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Product():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Product():load({ id = "product_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ProductEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PublicationEntity

```lua
local publication = client:Publication(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `id` | `string` | No |  |
| `section_ids` | `any` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |
| `web_url` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Publication():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Publication():load({ id = "publication_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PublicationEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SearchEntity

```lua
local search = client:Search(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_id` | `string` | No |  |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `id` | `string` | No |  |
| `is_boosted` | `boolean` | No |  |
| `score` | `number` | No |  |
| `thumbnail` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Search():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SearchEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SectionEntity

```lua
local section = client:Section(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accession` | `any` | No |  |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `artwork_id` | `string` | No |  |
| `content` | `any` | No |  |
| `generic_page_id` | `string` | No |  |
| `id` | `string` | No |  |
| `publication_id` | `string` | No |  |
| `publication_title` | `any` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |
| `web_url` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Section():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Section():load({ id = "section_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SectionEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SiteEntity

```lua
local site = client:Site(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `artwork_ids` | `any` | No |  |
| `artwork_titles` | `any` | No |  |
| `description` | `string` | No |  |
| `exhibition_ids` | `any` | No |  |
| `exhibition_titles` | `any` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |
| `web_url` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Site():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Site():load({ id = "site_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SiteEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SoundEntity

```lua
local sound = client:Sound(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_text` | `any` | No |  |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `artwork_ids` | `any` | No |  |
| `artwork_titles` | `any` | No |  |
| `content` | `any` | No |  |
| `content_e_tag` | `any` | No |  |
| `credit_line` | `any` | No |  |
| `id` | `string` | No |  |
| `is_educational_resource` | `boolean` | No |  |
| `is_multimedia_resource` | `boolean` | No |  |
| `is_teacher_resource` | `boolean` | No |  |
| `lake_guid` | `any` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `transcript` | `any` | No |  |
| `type` | `any` | No |  |
| `updated_at` | `any` | No |  |
| `web_url` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Sound():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Sound():load({ id = "sound_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SoundEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## StaticPageEntity

```lua
local static_page = client:StaticPage(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `id` | `string` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |
| `web_url` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:StaticPage():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:StaticPage():load({ id = "static_page_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `StaticPageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## TextEntity

```lua
local text = client:Text(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_text` | `any` | No |  |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `artwork_ids` | `any` | No |  |
| `artwork_titles` | `any` | No |  |
| `content` | `any` | No |  |
| `content_e_tag` | `any` | No |  |
| `credit_line` | `any` | No |  |
| `id` | `string` | No |  |
| `is_educational_resource` | `boolean` | No |  |
| `is_multimedia_resource` | `boolean` | No |  |
| `is_teacher_resource` | `boolean` | No |  |
| `lake_guid` | `any` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `type` | `any` | No |  |
| `updated_at` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Text():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Text():load({ id = "text_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TextEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## TourEntity

```lua
local tour = client:Tour(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `artist_titles` | `any` | No |  |
| `artwork_titles` | `any` | No |  |
| `description` | `string` | No |  |
| `id` | `string` | No |  |
| `image` | `any` | No |  |
| `intro` | `any` | No |  |
| `intro_link` | `any` | No |  |
| `intro_transcript` | `any` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |
| `weight` | `number` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Tour():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Tour():load({ id = "tour_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TourEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## VideoEntity

```lua
local video = client:Video(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_text` | `any` | No |  |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `artwork_ids` | `any` | No |  |
| `artwork_titles` | `any` | No |  |
| `content` | `any` | No |  |
| `content_e_tag` | `any` | No |  |
| `credit_line` | `any` | No |  |
| `id` | `string` | No |  |
| `is_educational_resource` | `boolean` | No |  |
| `is_multimedia_resource` | `boolean` | No |  |
| `is_teacher_resource` | `boolean` | No |  |
| `lake_guid` | `any` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `type` | `any` | No |  |
| `updated_at` | `any` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Video():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Video():load({ id = "video_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `VideoEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```lua
local client = sdk.new({
  feature = {
    test = { active = true },
  },
})
```

