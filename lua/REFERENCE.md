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
| `alt_titles` | `any` | No | Alternate names for this agent |
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `birth_date` | `any` | No | The year this agent was born |
| `death_date` | `any` | No | The year this agent died |
| `description` | `string` | No | A biographical description of the agent |
| `id` | `string` | No | Unique identifier of this resource. |
| `is_artist` | `boolean` | No | Whether the agent is an artist. |
| `sort_title` | `any` | No | Sortable name for this agent, typically with last name first. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `ulan_id` | `string` | No | Unique identifier of this agent in Getty's ULAN |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `copy` | `any` | No | The text of the article |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |

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
| `alt_artist_ids` | `any` | No | Unique identifiers of the non-preferred artists/cultures associated with this work |
| `alt_classification_ids` | `any` | No | Unique identifiers of all other non-preferred classification terms for this work |
| `alt_image_ids` | `any` | No | Unique identifiers of all non-preferred images of this work. |
| `alt_material_ids` | `any` | No | Unique identifiers of all other non-preferred material terms for this work |
| `alt_style_ids` | `any` | No | Unique identifiers of all other non-preferred style terms for this work |
| `alt_subject_ids` | `any` | No | Unique identifiers of all other non-preferred subject terms for this work |
| `alt_technique_ids` | `any` | No | Unique identifiers of all other non-preferred technique terms for this work |
| `alt_titles` | `any` | No | Alternate names for this work |
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `artist_display` | `any` | No | Readable description of the creator of this work. |
| `artist_id` | `string` | No | Unique identifier of the preferred artist/culture associated with this work |
| `artist_ids` | `any` | No | Unique identifier of all artist/cultures associated with this work |
| `artist_title` | `any` | No | Name of the preferred artist/culture associated with this work |
| `artist_titles` | `any` | No | Names of all artist/cultures associated with this work |
| `artwork_type_id` | `string` | No | Unique identifier of the kind of object or work |
| `artwork_type_title` | `any` | No | The kind of object or work (e.g. |
| `boost_rank` | `any` | No | Manual indication of what rank this artwork should take in search results. |
| `catalog_based_search_keyword_titles` | `any` | No | The keyword search values that would be catalog-based searches on this record |
| `catalogue_display` | `any` | No | Brief text listing all the catalogues raisonnés which include this work. |
| `category_ids` | `any` | No | Unique identifiers of the categories this work is a part of |
| `category_titles` | `any` | No | Names of the categories this artwork is a part of |
| `classification_id` | `string` | No | Unique identifier of the preferred classification term for this work |
| `classification_ids` | `any` | No | Unique identifiers of all classification terms for this work |
| `classification_title` | `any` | No | The name of the preferred classification term for this work |
| `classification_titles` | `any` | No | The names of all classification terms related to this artwork |
| `color` | `any` | No | Dominant color of this artwork in HSL |
| `colorfulness` | `any` | No | Unbounded positive float representing an abstract measure of colorfulness. |
| `copyright_notice` | `any` | No | Statement notifying how the work is protected by copyright. |
| `credit_line` | `any` | No | Brief statement indicating how the work came into the collection |
| `date_display` | `any` | No | Readable, free-text description of the period of time associated with the creation of this work. |
| `date_end` | `any` | No | The year of the period of time associated with the creation of this work |
| `date_qualifier_id` | `string` | No | Unique identifier of the qualifer to the dates provided for this record. |
| `date_qualifier_title` | `any` | No | Readable, text qualifer to the dates provided for this record. |
| `date_start` | `any` | No | The year of the period of time associated with the creation of this work |
| `department_id` | `string` | No | Unique identifier of the curatorial department that this work belongs to |
| `department_title` | `any` | No | Name of the curatorial department that this work belongs to |
| `description` | `string` | No | Longer explanation describing the work |
| `dimensions` | `any` | No | The size, shape, scale, and dimensions of the work. |
| `dimensions_detail` | `any` | No | The height, width, depth, and/or diameter of each section of the work in centimeters |
| `document_ids` | `any` | No | Unique identifiers of assets that serve as documentation for this artwork |
| `edition` | `any` | No | Edition number if the work is one of many |
| `exhibition_history` | `any` | No | List of all the places this work has been exhibited |
| `fiscal_year` | `any` | No | The fiscal year in which the work was acquired. |
| `fiscal_year_deaccession` | `any` | No | The fiscal year in which the work was deaccessioned. |
| `gallery_id` | `string` | No | Unique identifier of the location of this work in our museum |
| `gallery_title` | `any` | No | The location of this work in our museum |
| `has_advanced_imaging` | `boolean` | No | Whether this artwork is enhanced with 3D models, 360 image sequences, Mirador views, etc. |
| `has_educational_resources` | `boolean` | No | Whether this artwork has any documents tagged as educational |
| `has_multimedia_resources` | `boolean` | No | Whether this artwork has any associated microsites, digital publications, or documents tagged as multimedia |
| `has_not_been_viewed_much` | `boolean` | No | Whether the artwork hasn't been visited on our website very much |
| `id` | `string` | No | Unique identifier of this resource. |
| `image_embedding` | `any` | No | The generated embeddings describing the artwork image |
| `image_id` | `string` | No | Unique identifier of the preferred image to use to represent this work |
| `inscriptions` | `any` | No | A description of distinguishing or identifying physical markings that are on the work |
| `internal_department_id` | `string` | No | An internal department id we use for analytics. |
| `is_boosted` | `boolean` | No | Whether this document should be boosted in search |
| `is_on_view` | `boolean` | No | Whether the work is on display |
| `is_public_domain` | `boolean` | No | Whether the work is in the public domain, meaning it was created before copyrights existed or has left the copyright term |
| `is_zoomable` | `boolean` | No | Whether images of the work are allowed to be displayed in a zoomable interface. |
| `latitude` | `number` | No | Latitude coordinate of the location of this work in our galleries |
| `latlon` | `any` | No | Latitude and longitude coordinates of the location of this work in our galleries |
| `longitude` | `number` | No | Longitude coordinate of the location of this work in our galleries |
| `main_reference_number` | `number` | No | Unique identifier assigned to the artwork upon acquisition |
| `material_id` | `string` | No | Unique identifier of the preferred material term for this work |
| `material_ids` | `any` | No | Unique identifiers of all material terms for this work |
| `material_titles` | `any` | No | The names of all material terms related to this artwork |
| `max_zoom_window_size` | `any` | No | The maximum size of the window the image is allowed to be viewed in, in pixels. |
| `medium_display` | `any` | No | The substances or materials used in the creation of a work |
| `nomisma_id` | `string` | No | Unique identifier of this work in the nomisma coin database |
| `on_loan_display` | `any` | No | If an artwork is on loan, this contains details about the loan |
| `pageviews` | `any` | No | Approx. |
| `pageviews_recent` | `any` | No | Approx. |
| `place_of_origin` | `any` | No | The location where the creation, design, or production of the work took place, or the original location of the work |
| `provenance_text` | `any` | No | Ownership/collecting history of the work. |
| `publication_history` | `any` | No | Bibliographic list of all the places this work has been published |
| `publishing_verification_level` | `any` | No | Indicator of how much metadata on the work in published. |
| `section_ids` | `any` | No | Unique identifiers of the digital publication chapters this work in included in |
| `section_titles` | `any` | No | Names of the digital publication chapters this work is included in |
| `short_description` | `any` | No | Short explanation describing the work |
| `site_ids` | `any` | No | Unique identifiers of the microsites this work is a part of |
| `sound_ids` | `any` | No | Unique identifiers of the audio about this work |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `style_id` | `string` | No | Unique identifier of the preferred style term for this work |
| `style_ids` | `any` | No | Unique identifiers of all style terms for this work |
| `style_title` | `any` | No | The name of the preferred style term for this work |
| `style_titles` | `any` | No | The names of all style terms related to this artwork |
| `subject_id` | `string` | No | Unique identifier of the preferred subject term for this work |
| `subject_ids` | `any` | No | Unique identifiers of all subject terms for this work |
| `subject_titles` | `any` | No | The names of all subject terms related to this artwork |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `technique_id` | `string` | No | Unique identifier of the preferred technique term for this work |
| `technique_ids` | `any` | No | Unique identifiers of all technique terms for this work |
| `technique_titles` | `any` | No | The names of all technique terms related to this artwork |
| `term_titles` | `any` | No | The names of the taxonomy tags for this work |
| `text_embedding` | `any` | No | The generated embeddings of artwork text |
| `text_ids` | `any` | No | Unique identifiers of the texts about this work |
| `theme_titles` | `any` | No | The names of all thematic publish categories related to this artwork |
| `thumbnail` | `any` | No | Metadata about the image referenced by `image_id`. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |
| `video_ids` | `any` | No | Unique identifiers of the videos about this work |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |

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
| `aat_id` | `string` | No | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |

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
| `aat_id` | `string` | No | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `parent_id` | `string` | No | Unique identifier of this category's parent |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `subtype` | `any` | No | Takes one of the following values: classification, material, technique, style, subject, department, theme |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `copy` | `any` | No | The text of the page |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | No | The URL to this page on our website |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `author_display` | `any` | No | A display-friendly text of the authors of this article |
| `copy` | `any` | No | The text of the article |
| `digital_publication_id` | `string` | No | Unique identifier of the digital publication this article belongs to |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | No | The URL to this article on our website |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `copy` | `any` | No | The text of the page |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | No | The URL to this page on our website |

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
| `alt_audience_ids` | `any` | No | Unique identifiers indicating the alternate audiences for this event |
| `alt_event_type_ids` | `any` | No | Unique identifiers indicating the alternate types of this event |
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `audience_id` | `string` | No | Unique identifier indicating the preferred audience for this event |
| `buy_button_caption` | `any` | No | Additional text below the ticket/registration button |
| `buy_button_text` | `any` | No | The text used on the ticket/registration button |
| `date_display` | `any` | No | A readable display of the event dates |
| `description` | `string` | No | All copytext of the event |
| `door_time` | `any` | No | The time the doors open for this event |
| `end_date` | `any` | No | The date the event ends |
| `end_time` | `any` | No | The time the event ends |
| `entrance` | `any` | No | Which entrance to use for this event |
| `event_host_id` | `string` | No | Unique identifier of the host (cf. |
| `event_host_title` | `any` | No | Unique identifier of the host (cf. |
| `event_type_id` | `string` | No | Unique identifier indicating the preferred type of this event |
| `header_description` | `any` | No | Brief description of the event displayed below the title |
| `hero_caption` | `any` | No | Text displayed with the hero image on the event |
| `id` | `string` | No | Unique identifier of this resource. |
| `image_url` | `any` | No | The URL of an image representing this page |
| `is_admission_required` | `boolean` | No | Whether admission to the museum is required to attend this event |
| `is_after_hours` | `boolean` | No | Whether the event is to be held after the museum closes |
| `is_free` | `boolean` | No | Whether the event is free |
| `is_member_exclusive` | `boolean` | No | Whether the event is exclusive to members of the museum |
| `is_private` | `boolean` | No | Whether the event is private |
| `is_registration_required` | `boolean` | No | Whether registration is required to attend the event |
| `is_sales_button_hidden` | `boolean` | No | Whether the buy tickets button is hidden on the website event page |
| `is_sold_out` | `boolean` | No | Whether the event is sold out |
| `is_ticketed` | `boolean` | No | Whether a ticket is required to attend the event |
| `is_virtual_event` | `boolean` | No | Whether the event is being held virtually |
| `join_url` | `any` | No | URL to the membership signup page via this event |
| `layout_type` | `any` | No | Number indicating the type of layout this event page uses |
| `list_description` | `any` | No | One-sentence description of the event displayed in listings |
| `location` | `any` | No | Where the event takes place |
| `program_ids` | `any` | No | Unique identifiers indicating the programs this event is a part of |
| `program_titles` | `any` | No | Titles of the programs this event is a part of |
| `rsvp_link` | `any` | No | The URL to the sales site for this event |
| `search_tags` | `any` | No | Editor-specified list of tags to aid in internal search |
| `short_description` | `any` | No | Brief description of the event |
| `slug` | `string` | No | A string used in the URL for this event |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `start_date` | `any` | No | The date the event begins |
| `start_time` | `any` | No | The time the event starts |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `survey_url` | `any` | No | URL to the survey associated with this event |
| `ticketed_event_id` | `string` | No | Unique identifier of the event in the ticketing system this website event is tied to |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `title_display` | `any` | No | The name of this event formatted with HTML (optional) |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |
| `virtual_event_passcode` | `any` | No | Passcode to access the virtual event |
| `virtual_event_url` | `any` | No | URL to the virtual event |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `button_caption` | `any` | No | Additional text below the ticket/registration button |
| `button_text` | `any` | No | The text used on the ticket/registration button |
| `button_url` | `any` | No | The URL to the sales site or an RSVP link for this event |
| `description` | `string` | No | Description of the event |
| `end_at` | `any` | No | The date the event occurrence ends |
| `event_id` | `string` | No | Identifier of the master event of which this is an occurrence |
| `id` | `string` | No | Unique identifier of this resource. |
| `image_url` | `any` | No | The URL of an image representing this page |
| `is_private` | `boolean` | No | Whether the event is private. |
| `is_sales_button_hidden` | `boolean` | No | Whether the buy tickets button is hidden on the website event page |
| `is_ticketed` | `boolean` | No | Whether a ticket is required to attend the event |
| `location` | `any` | No | Where the event takes place |
| `off_sale_at` | `any` | No | Date and time the event goes off sale |
| `on_sale_at` | `any` | No | Date and time the event goes on sale |
| `short_description` | `any` | No | Brief description of the event |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `start_at` | `any` | No | The date the event occurrence begins |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `title_display` | `any` | No | The name of this event formatted with HTML (optional) |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `is_affiliate_group` | `boolean` | No | Whether this program represents an affiliate group |
| `is_event_host` | `boolean` | No | Whether this program represents an event host |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |

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
| `aic_end_at` | `any` | No | Date the exhibition closed at the Art Institute of Chicago |
| `aic_start_at` | `any` | No | Date the exhibition opened at the Art Institute of Chicago |
| `alt_image_ids` | `any` | No | Unique identifiers of all non-preferred images of this exhibition. |
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `artist_ids` | `any` | No | Unique identifiers of the artist agent records representing who was shown in the exhibition |
| `artwork_ids` | `any` | No | Unique identifiers of the artworks that were part of the exhibition |
| `artwork_titles` | `any` | No | Names of the artworks that were part of the exhibition |
| `document_ids` | `any` | No | Unique identifiers of assets that serve as documentation for this exhibition |
| `gallery_id` | `string` | No | Unique identifier of the gallery that mainly housed the exhibition |
| `gallery_title` | `any` | No | The name of the gallery that mainly housed the exhibition |
| `id` | `string` | No | Unique identifier of this resource. |
| `image_id` | `string` | No | Unique identifier of the preferred image to use to represent this exhibition |
| `image_url` | `any` | No | URL to the hero image from the website |
| `is_featured` | `boolean` | No | Is this exhibition currently featured on our website? |
| `is_published` | `boolean` | No | Is this exhibition currently published on our website? |
| `position` | `any` | No | Numering position represnting the order in which this exhibition is featured on the website |
| `short_description` | `any` | No | Brief explanation of what this exhibition is |
| `site_ids` | `any` | No | Unique identifiers of the microsites this exhibition is a part of |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `status` | `any` | No | Whether the exhibition is open or closed |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | No | URL to this exhibition on our website |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `floor` | `any` | No | The level the gallery is on, e.g., 1, 2, 3, or LL |
| `id` | `string` | No | Unique identifier of this resource. |
| `is_closed` | `boolean` | No | Whether the gallery is currently closed |
| `latitude` | `number` | No | Latitude coordinate of the center of the room |
| `latlon` | `any` | No | Latitude and longitude coordinates of the center of the room |
| `longitude` | `number` | No | Longitude coordinate of the center of the room |
| `number` | `any` | No | The gallery's room number. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | `string` | No | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `copy` | `any` | No | The text of the page |
| `id` | `string` | No | Unique identifier of this resource. |
| `search_tags` | `any` | No | Editor-specified list of tags to aid in internal search |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | No | The URL to this page on our website |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `copy` | `any` | No | The text of the highlight description |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |

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
| `additional_text` | `any` | No | Additional information about the hours |
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `friday_is_closed` | `any` | No | Whether the museum is closed on Fridays |
| `friday_member_close` | `any` | No | The time member hours ends on Fridays |
| `friday_member_open` | `any` | No | The time member hours starts on Fridays |
| `friday_public_close` | `any` | No | The time public hours ends on Fridays |
| `friday_public_open` | `any` | No | The time public hours starts on Fridays |
| `id` | `string` | No | Unique identifier of this resource. |
| `monday_is_closed` | `any` | No | Whether the museum is closed on Mondays |
| `monday_member_close` | `any` | No | The time member hours ends on Mondays |
| `monday_member_open` | `any` | No | The time member hours starts on Mondays |
| `monday_public_close` | `any` | No | The time public hours ends on Mondays |
| `monday_public_open` | `any` | No | The time public hours starts on Mondays |
| `saturday_is_closed` | `any` | No | Whether the museum is closed on Saturdays |
| `saturday_member_close` | `any` | No | The time member hours ends on Saturdays |
| `saturday_member_open` | `any` | No | The time member hours starts on Saturdays |
| `saturday_public_close` | `any` | No | The time public hours ends on Saturdays |
| `saturday_public_open` | `any` | No | The time public hours starts on Saturdays |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `summary` | `any` | No | Readable summary of the hours |
| `sunday_is_closed` | `any` | No | Whether the museum is closed on Sundays |
| `sunday_member_close` | `any` | No | The time member hours ends on Sundays |
| `sunday_member_open` | `any` | No | The time member hours starts on Sundays |
| `sunday_public_close` | `any` | No | The time public hours ends on Sundays |
| `sunday_public_open` | `any` | No | The time public hours starts on Sundays |
| `thursday_is_closed` | `any` | No | Whether the museum is closed on Thursdays |
| `thursday_member_close` | `any` | No | The time member hours ends on Thursdays |
| `thursday_member_open` | `any` | No | The time member hours starts on Thursdays |
| `thursday_public_close` | `any` | No | The time public hours ends on Thursdays |
| `thursday_public_open` | `any` | No | The time public hours starts on Thursdays |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `tuesday_is_closed` | `any` | No | Whether the museum is closed on Tuesdays |
| `tuesday_member_close` | `any` | No | The time member hours ends on Tuesdays |
| `tuesday_member_open` | `any` | No | The time member hours starts on Tuesdays |
| `tuesday_public_close` | `any` | No | The time public hours ends on Tuesdays |
| `tuesday_public_open` | `any` | No | The time public hours starts on Tuesdays |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |
| `wednesday_is_closed` | `any` | No | Whether the museum is closed on Wednesdays |
| `wednesday_member_close` | `any` | No | The time member hours ends on Wednesdays |
| `wednesday_member_open` | `any` | No | The time member hours starts on Wednesdays |
| `wednesday_public_close` | `any` | No | The time public hours ends on Wednesdays |
| `wednesday_public_open` | `any` | No | The time public hours starts on Wednesdays |

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
| `ahash` | `any` | No | Image hash generated using ahash algorithm with 64 boolean subfields |
| `alt_text` | `any` | No | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `artwork_ids` | `any` | No | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `any` | No | Names of the artworks associated with this asset |
| `color` | `any` | No | Dominant color of this image in HSL |
| `colorfulness` | `any` | No | Unbounded positive float representing an abstract measure of colorfulness. |
| `content` | `any` | No | Text of or URL to the contents of this asset |
| `content_e_tag` | `any` | No | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `any` | No | Asset-specific copyright information |
| `fingerprint` | `any` | No | Image hashes: aHash, dHash, pHash, wHash |
| `height` | `number` | No | Native height of the image |
| `id` | `string` | No | Unique identifier of this resource. |
| `iiif_url` | `any` | No | IIIF URL of this image |
| `is_educational_resource` | `boolean` | No | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `boolean` | No | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `boolean` | No | Whether this resource is considered to be educational |
| `lake_guid` | `any` | No | Unique UUID of this resource in LAKE, our DAMS. |
| `lqip` | `any` | No | Low-quality image placeholder (LQIP). |
| `phash` | `any` | No | Image hash generated using phash algorithm with 64 boolean subfields |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `type` | `any` | No | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |
| `width` | `number` | No | Native width of the image |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `copy` | `any` | No | The text of the page |
| `id` | `string` | No | Unique identifier of this resource. |
| `search_tags` | `any` | No | Editor-specified list of tags to aid in internal search |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | No | The URL to this page on our website |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `latitude` | `number` | No | Latitude coordinate of the center of the room |
| `longitude` | `number` | No | Longitude coordinate of the center of the room |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | `string` | No | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `copy` | `any` | No | The text of the page |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | No | The URL to this page on our website |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `copy` | `any` | No | The text of the page |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | No | The URL to this page on our website |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `artist_ids` | `any` | No | Unique identifiers of the artists associated with this product |
| `artwork_ids` | `any` | No | Unique identifiers of the artworks associated with this product |
| `description` | `string` | No | Explanation of what this product is |
| `exhibition_ids` | `any` | No | Unique identifiers of the exhibitions associated with this product |
| `external_sku` | `any` | No | Numeric product identification code of a machine-readable barcode, when the customer sku differs from our internal one |
| `id` | `string` | No | Unique identifier of this resource. |
| `image_url` | `any` | No | URL of an image for this product |
| `max_compare_at_price` | `any` | No | Number indicating how much the most expensive variant of a product cost before a sale |
| `max_current_price` | `any` | No | Number indicating how much the most expensive variant of a product costs right now |
| `min_compare_at_price` | `any` | No | Number indicating how much the least expensive variant of a product cost before a sale |
| `min_current_price` | `any` | No | Number indicating how much the least expensive variant of a product costs right now |
| `price_display` | `any` | No | Explanation of what this product is |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | No | URL of this product in the shop |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `section_ids` | `any` | No | Unique identifiers of the sections of this publication |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | No | URL to the publication |

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
| `api_id` | `string` | No | API unique identifier |
| `api_link` | `any` | No | URL to this recource in the API |
| `api_model` | `any` | No | Name of the model the resource represents |
| `id` | `string` | No | Unique identifier within the search index |
| `is_boosted` | `boolean` | No | Whether this record has been flagged to be boosted |
| `score` | `number` | No | Search index ranking of the result |
| `thumbnail` | `any` | No | Metadata on the image representing this record |
| `timestamp` | `any` | No | Date this record was last updated in the API |
| `title` | `string` | No | The name of this resource |

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
| `accession` | `any` | No | An accession number parsed from the title or tombstone |
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `artwork_id` | `string` | No | Unique identifier of the artwork with which this section is associated |
| `content` | `any` | No | Content of this section in plaintext |
| `generic_page_id` | `string` | No | Unique identifier of the page on the website that represents the publication this section belongs to |
| `id` | `string` | No | Unique identifier of this resource. |
| `publication_id` | `string` | No | Unique identifier of the publication this section belongs to |
| `publication_title` | `any` | No | Name of the publication this section belongs to |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | No | URL to the section |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `artwork_ids` | `any` | No | Unique identifiers of the artworks this site is associated with |
| `artwork_titles` | `any` | No | Names of the artworks this site is associated with |
| `description` | `string` | No | Explanation of what this site is |
| `exhibition_ids` | `any` | No | Unique identifier of the exhibitions this site is associated with |
| `exhibition_titles` | `any` | No | Names of the exhibitions this site is associated with |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | No | URL to this site |

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
| `alt_text` | `any` | No | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `artwork_ids` | `any` | No | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `any` | No | Names of the artworks associated with this asset |
| `content` | `any` | No | Text of or URL to the contents of this asset |
| `content_e_tag` | `any` | No | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `any` | No | Asset-specific copyright information |
| `id` | `string` | No | Unique identifier of this resource. |
| `is_educational_resource` | `boolean` | No | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `boolean` | No | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `boolean` | No | Whether this resource is considered to be educational |
| `lake_guid` | `any` | No | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | Name of this mobile audio file – derived from the artwork and tour titles |
| `transcript` | `any` | No | Text transcription of the audio file |
| `type` | `any` | No | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | No | URL to the audio file |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | No | The URL to this page on our website |

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
| `alt_text` | `any` | No | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `artwork_ids` | `any` | No | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `any` | No | Names of the artworks associated with this asset |
| `content` | `any` | No | Text of or URL to the contents of this asset |
| `content_e_tag` | `any` | No | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `any` | No | Asset-specific copyright information |
| `id` | `string` | No | Unique identifier of this resource. |
| `is_educational_resource` | `boolean` | No | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `boolean` | No | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `boolean` | No | Whether this resource is considered to be educational |
| `lake_guid` | `any` | No | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `type` | `any` | No | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `artist_titles` | `any` | No | Names of the artists of the artworks featured in this tour's tour stops |
| `artwork_titles` | `any` | No | Names of the artworks featured in this tour's tour stops |
| `description` | `string` | No | Explanation of what the tour is |
| `id` | `string` | No | Unique identifier of this resource. |
| `image` | `any` | No | The main image for the tour |
| `intro` | `any` | No | Text introducing the tour |
| `intro_link` | `any` | No | Link to the audio file of the introduction |
| `intro_transcript` | `any` | No | Transcript of the introduction audio to the tour |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |
| `weight` | `number` | No | Number representing this tour's sort order |

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
| `alt_text` | `any` | No | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `artwork_ids` | `any` | No | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `any` | No | Names of the artworks associated with this asset |
| `content` | `any` | No | Text of or URL to the contents of this asset |
| `content_e_tag` | `any` | No | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `any` | No | Asset-specific copyright information |
| `id` | `string` | No | Unique identifier of this resource. |
| `is_educational_resource` | `boolean` | No | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `boolean` | No | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `boolean` | No | Whether this resource is considered to be educational |
| `lake_guid` | `any` | No | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `type` | `any` | No | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |

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


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### `test`

In-memory mock transport for testing without a live server.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

Options above are those the model carries a default for. A feature may
also accept callback options — a `sink` to receive each record, for
instance — which have no default and are covered in the full feature
reference.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

