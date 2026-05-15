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
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts, sdkopts)`

Create a test client with mock features active. Both arguments may be `nil`.

```lua
local client = sdk.test(nil, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Agent(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Agent(nil):load({ id = "agent_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:AgentRole(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:AgentRole(nil):load({ id = "agent_role_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:AgentType(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:AgentType(nil):load({ id = "agent_type_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Article(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Article(nil):load({ id = "article_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Artwork(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Artwork(nil):load({ id = "artwork_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:ArtworkDateQualifier(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:ArtworkDateQualifier(nil):load({ id = "artwork_date_qualifier_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:ArtworkPlaceQualifier(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:ArtworkPlaceQualifier(nil):load({ id = "artwork_place_qualifier_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:ArtworkType(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:ArtworkType(nil):load({ id = "artwork_type_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:CategoryTerm(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:CategoryTerm(nil):load({ id = "category_term_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:DigitalPublication(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:DigitalPublication(nil):load({ id = "digital_publication_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:DigitalPublicationArticle(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:DigitalPublicationArticle(nil):load({ id = "digital_publication_article_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:EducatorResource(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EducatorResource(nil):load({ id = "educator_resource_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Event(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Event(nil):load({ id = "event_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:EventOccurrence(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EventOccurrence(nil):load({ id = "event_occurrence_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:EventProgram(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EventProgram(nil):load({ id = "event_program_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Exhibition(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Exhibition(nil):load({ id = "exhibition_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Gallery(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Gallery(nil):load({ id = "gallery_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:GenericPage(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:GenericPage(nil):load({ id = "generic_page_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Highlight(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Highlight(nil):load({ id = "highlight_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Hour(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Hour(nil):load({ id = "hour_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Image(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Image(nil):load({ id = "image_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:LandingPage(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:LandingPage(nil):load({ id = "landing_page_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Place(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Place(nil):load({ id = "place_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:PressRelease(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:PressRelease(nil):load({ id = "press_release_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:PrintedPublication(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:PrintedPublication(nil):load({ id = "printed_publication_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Product(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Product(nil):load({ id = "product_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Publication(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Publication(nil):load({ id = "publication_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Search(nil):list(nil, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Section(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Section(nil):load({ id = "section_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Site(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Site(nil):load({ id = "site_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Sound(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Sound(nil):load({ id = "sound_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:StaticPage(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:StaticPage(nil):load({ id = "static_page_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Text(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Text(nil):load({ id = "text_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Tour(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Tour(nil):load({ id = "tour_id" }, nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Video(nil):list(nil, nil)
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Video(nil):load({ id = "video_id" }, nil)
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

