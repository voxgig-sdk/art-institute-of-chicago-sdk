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
| `alt_titles` | `Object` | No | Alternate names for this agent |
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `birth_date` | `Object` | No | The year this agent was born |
| `death_date` | `Object` | No | The year this agent died |
| `description` | `String` | No | A biographical description of the agent |
| `id` | `String` | No | Unique identifier of this resource. |
| `is_artist` | `Boolean` | No | Whether the agent is an artist. |
| `sort_title` | `Object` | No | Sortable name for this agent, typically with last name first. |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `ulan_id` | `String` | No | Unique identifier of this agent in Getty's ULAN |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `id` | `String` | No | Unique identifier of this resource. |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `id` | `String` | No | Unique identifier of this resource. |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `copy` | `Object` | No | The text of the article |
| `id` | `String` | No | Unique identifier of this resource. |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |

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
| `alt_artist_ids` | `Object` | No | Unique identifiers of the non-preferred artists/cultures associated with this work |
| `alt_classification_ids` | `Object` | No | Unique identifiers of all other non-preferred classification terms for this work |
| `alt_image_ids` | `Object` | No | Unique identifiers of all non-preferred images of this work. |
| `alt_material_ids` | `Object` | No | Unique identifiers of all other non-preferred material terms for this work |
| `alt_style_ids` | `Object` | No | Unique identifiers of all other non-preferred style terms for this work |
| `alt_subject_ids` | `Object` | No | Unique identifiers of all other non-preferred subject terms for this work |
| `alt_technique_ids` | `Object` | No | Unique identifiers of all other non-preferred technique terms for this work |
| `alt_titles` | `Object` | No | Alternate names for this work |
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `artist_display` | `Object` | No | Readable description of the creator of this work. |
| `artist_id` | `String` | No | Unique identifier of the preferred artist/culture associated with this work |
| `artist_ids` | `Object` | No | Unique identifier of all artist/cultures associated with this work |
| `artist_title` | `Object` | No | Name of the preferred artist/culture associated with this work |
| `artist_titles` | `Object` | No | Names of all artist/cultures associated with this work |
| `artwork_type_id` | `String` | No | Unique identifier of the kind of object or work |
| `artwork_type_title` | `Object` | No | The kind of object or work (e.g. |
| `boost_rank` | `Object` | No | Manual indication of what rank this artwork should take in search results. |
| `catalog_based_search_keyword_titles` | `Object` | No | The keyword search values that would be catalog-based searches on this record |
| `catalogue_display` | `Object` | No | Brief text listing all the catalogues raisonnés which include this work. |
| `category_ids` | `Object` | No | Unique identifiers of the categories this work is a part of |
| `category_titles` | `Object` | No | Names of the categories this artwork is a part of |
| `classification_id` | `String` | No | Unique identifier of the preferred classification term for this work |
| `classification_ids` | `Object` | No | Unique identifiers of all classification terms for this work |
| `classification_title` | `Object` | No | The name of the preferred classification term for this work |
| `classification_titles` | `Object` | No | The names of all classification terms related to this artwork |
| `color` | `Object` | No | Dominant color of this artwork in HSL |
| `colorfulness` | `Object` | No | Unbounded positive float representing an abstract measure of colorfulness. |
| `copyright_notice` | `Object` | No | Statement notifying how the work is protected by copyright. |
| `credit_line` | `Object` | No | Brief statement indicating how the work came into the collection |
| `date_display` | `Object` | No | Readable, free-text description of the period of time associated with the creation of this work. |
| `date_end` | `Object` | No | The year of the period of time associated with the creation of this work |
| `date_qualifier_id` | `String` | No | Unique identifier of the qualifer to the dates provided for this record. |
| `date_qualifier_title` | `Object` | No | Readable, text qualifer to the dates provided for this record. |
| `date_start` | `Object` | No | The year of the period of time associated with the creation of this work |
| `department_id` | `String` | No | Unique identifier of the curatorial department that this work belongs to |
| `department_title` | `Object` | No | Name of the curatorial department that this work belongs to |
| `description` | `String` | No | Longer explanation describing the work |
| `dimensions` | `Object` | No | The size, shape, scale, and dimensions of the work. |
| `dimensions_detail` | `Object` | No | The height, width, depth, and/or diameter of each section of the work in centimeters |
| `document_ids` | `Object` | No | Unique identifiers of assets that serve as documentation for this artwork |
| `edition` | `Object` | No | Edition number if the work is one of many |
| `exhibition_history` | `Object` | No | List of all the places this work has been exhibited |
| `fiscal_year` | `Object` | No | The fiscal year in which the work was acquired. |
| `fiscal_year_deaccession` | `Object` | No | The fiscal year in which the work was deaccessioned. |
| `gallery_id` | `String` | No | Unique identifier of the location of this work in our museum |
| `gallery_title` | `Object` | No | The location of this work in our museum |
| `has_advanced_imaging` | `Boolean` | No | Whether this artwork is enhanced with 3D models, 360 image sequences, Mirador views, etc. |
| `has_educational_resources` | `Boolean` | No | Whether this artwork has any documents tagged as educational |
| `has_multimedia_resources` | `Boolean` | No | Whether this artwork has any associated microsites, digital publications, or documents tagged as multimedia |
| `has_not_been_viewed_much` | `Boolean` | No | Whether the artwork hasn't been visited on our website very much |
| `id` | `String` | No | Unique identifier of this resource. |
| `image_embedding` | `Object` | No | The generated embeddings describing the artwork image |
| `image_id` | `String` | No | Unique identifier of the preferred image to use to represent this work |
| `inscriptions` | `Object` | No | A description of distinguishing or identifying physical markings that are on the work |
| `internal_department_id` | `String` | No | An internal department id we use for analytics. |
| `is_boosted` | `Boolean` | No | Whether this document should be boosted in search |
| `is_on_view` | `Boolean` | No | Whether the work is on display |
| `is_public_domain` | `Boolean` | No | Whether the work is in the public domain, meaning it was created before copyrights existed or has left the copyright term |
| `is_zoomable` | `Boolean` | No | Whether images of the work are allowed to be displayed in a zoomable interface. |
| `latitude` | `Float` | No | Latitude coordinate of the location of this work in our galleries |
| `latlon` | `Object` | No | Latitude and longitude coordinates of the location of this work in our galleries |
| `longitude` | `Float` | No | Longitude coordinate of the location of this work in our galleries |
| `main_reference_number` | `Integer` | No | Unique identifier assigned to the artwork upon acquisition |
| `material_id` | `String` | No | Unique identifier of the preferred material term for this work |
| `material_ids` | `Object` | No | Unique identifiers of all material terms for this work |
| `material_titles` | `Object` | No | The names of all material terms related to this artwork |
| `max_zoom_window_size` | `Object` | No | The maximum size of the window the image is allowed to be viewed in, in pixels. |
| `medium_display` | `Object` | No | The substances or materials used in the creation of a work |
| `nomisma_id` | `String` | No | Unique identifier of this work in the nomisma coin database |
| `on_loan_display` | `Object` | No | If an artwork is on loan, this contains details about the loan |
| `pageviews` | `Object` | No | Approx. |
| `pageviews_recent` | `Object` | No | Approx. |
| `place_of_origin` | `Object` | No | The location where the creation, design, or production of the work took place, or the original location of the work |
| `provenance_text` | `Object` | No | Ownership/collecting history of the work. |
| `publication_history` | `Object` | No | Bibliographic list of all the places this work has been published |
| `publishing_verification_level` | `Object` | No | Indicator of how much metadata on the work in published. |
| `section_ids` | `Object` | No | Unique identifiers of the digital publication chapters this work in included in |
| `section_titles` | `Object` | No | Names of the digital publication chapters this work is included in |
| `short_description` | `Object` | No | Short explanation describing the work |
| `site_ids` | `Object` | No | Unique identifiers of the microsites this work is a part of |
| `sound_ids` | `Object` | No | Unique identifiers of the audio about this work |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `style_id` | `String` | No | Unique identifier of the preferred style term for this work |
| `style_ids` | `Object` | No | Unique identifiers of all style terms for this work |
| `style_title` | `Object` | No | The name of the preferred style term for this work |
| `style_titles` | `Object` | No | The names of all style terms related to this artwork |
| `subject_id` | `String` | No | Unique identifier of the preferred subject term for this work |
| `subject_ids` | `Object` | No | Unique identifiers of all subject terms for this work |
| `subject_titles` | `Object` | No | The names of all subject terms related to this artwork |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `technique_id` | `String` | No | Unique identifier of the preferred technique term for this work |
| `technique_ids` | `Object` | No | Unique identifiers of all technique terms for this work |
| `technique_titles` | `Object` | No | The names of all technique terms related to this artwork |
| `term_titles` | `Object` | No | The names of the taxonomy tags for this work |
| `text_embedding` | `Object` | No | The generated embeddings of artwork text |
| `text_ids` | `Object` | No | Unique identifiers of the texts about this work |
| `theme_titles` | `Object` | No | The names of all thematic publish categories related to this artwork |
| `thumbnail` | `Object` | No | Metadata about the image referenced by `image_id`. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |
| `video_ids` | `Object` | No | Unique identifiers of the videos about this work |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `id` | `String` | No | Unique identifier of this resource. |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `id` | `String` | No | Unique identifier of this resource. |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |

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
| `aat_id` | `String` | No | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `id` | `String` | No | Unique identifier of this resource. |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |

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
| `aat_id` | `String` | No | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `id` | `String` | No | Unique identifier of this resource. |
| `parent_id` | `String` | No | Unique identifier of this category's parent |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `subtype` | `Object` | No | Takes one of the following values: classification, material, technique, style, subject, department, theme |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `copy` | `Object` | No | The text of the page |
| `id` | `String` | No | Unique identifier of this resource. |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | No | The URL to this page on our website |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `author_display` | `Object` | No | A display-friendly text of the authors of this article |
| `copy` | `Object` | No | The text of the article |
| `digital_publication_id` | `String` | No | Unique identifier of the digital publication this article belongs to |
| `id` | `String` | No | Unique identifier of this resource. |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | No | The URL to this article on our website |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `copy` | `Object` | No | The text of the page |
| `id` | `String` | No | Unique identifier of this resource. |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | No | The URL to this page on our website |

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
| `alt_audience_ids` | `Object` | No | Unique identifiers indicating the alternate audiences for this event |
| `alt_event_type_ids` | `Object` | No | Unique identifiers indicating the alternate types of this event |
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `audience_id` | `String` | No | Unique identifier indicating the preferred audience for this event |
| `buy_button_caption` | `Object` | No | Additional text below the ticket/registration button |
| `buy_button_text` | `Object` | No | The text used on the ticket/registration button |
| `date_display` | `Object` | No | A readable display of the event dates |
| `description` | `String` | No | All copytext of the event |
| `door_time` | `Object` | No | The time the doors open for this event |
| `end_date` | `Object` | No | The date the event ends |
| `end_time` | `Object` | No | The time the event ends |
| `entrance` | `Object` | No | Which entrance to use for this event |
| `event_host_id` | `String` | No | Unique identifier of the host (cf. |
| `event_host_title` | `Object` | No | Unique identifier of the host (cf. |
| `event_type_id` | `String` | No | Unique identifier indicating the preferred type of this event |
| `header_description` | `Object` | No | Brief description of the event displayed below the title |
| `hero_caption` | `Object` | No | Text displayed with the hero image on the event |
| `id` | `String` | No | Unique identifier of this resource. |
| `image_url` | `Object` | No | The URL of an image representing this page |
| `is_admission_required` | `Boolean` | No | Whether admission to the museum is required to attend this event |
| `is_after_hours` | `Boolean` | No | Whether the event is to be held after the museum closes |
| `is_free` | `Boolean` | No | Whether the event is free |
| `is_member_exclusive` | `Boolean` | No | Whether the event is exclusive to members of the museum |
| `is_private` | `Boolean` | No | Whether the event is private |
| `is_registration_required` | `Boolean` | No | Whether registration is required to attend the event |
| `is_sales_button_hidden` | `Boolean` | No | Whether the buy tickets button is hidden on the website event page |
| `is_sold_out` | `Boolean` | No | Whether the event is sold out |
| `is_ticketed` | `Boolean` | No | Whether a ticket is required to attend the event |
| `is_virtual_event` | `Boolean` | No | Whether the event is being held virtually |
| `join_url` | `Object` | No | URL to the membership signup page via this event |
| `layout_type` | `Object` | No | Number indicating the type of layout this event page uses |
| `list_description` | `Object` | No | One-sentence description of the event displayed in listings |
| `location` | `Object` | No | Where the event takes place |
| `program_ids` | `Object` | No | Unique identifiers indicating the programs this event is a part of |
| `program_titles` | `Object` | No | Titles of the programs this event is a part of |
| `rsvp_link` | `Object` | No | The URL to the sales site for this event |
| `search_tags` | `Object` | No | Editor-specified list of tags to aid in internal search |
| `short_description` | `Object` | No | Brief description of the event |
| `slug` | `String` | No | A string used in the URL for this event |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `start_date` | `Object` | No | The date the event begins |
| `start_time` | `Object` | No | The time the event starts |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `survey_url` | `Object` | No | URL to the survey associated with this event |
| `ticketed_event_id` | `String` | No | Unique identifier of the event in the ticketing system this website event is tied to |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `title_display` | `Object` | No | The name of this event formatted with HTML (optional) |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |
| `virtual_event_passcode` | `Object` | No | Passcode to access the virtual event |
| `virtual_event_url` | `Object` | No | URL to the virtual event |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `button_caption` | `Object` | No | Additional text below the ticket/registration button |
| `button_text` | `Object` | No | The text used on the ticket/registration button |
| `button_url` | `Object` | No | The URL to the sales site or an RSVP link for this event |
| `description` | `String` | No | Description of the event |
| `end_at` | `Object` | No | The date the event occurrence ends |
| `event_id` | `String` | No | Identifier of the master event of which this is an occurrence |
| `id` | `String` | No | Unique identifier of this resource. |
| `image_url` | `Object` | No | The URL of an image representing this page |
| `is_private` | `Boolean` | No | Whether the event is private. |
| `is_sales_button_hidden` | `Boolean` | No | Whether the buy tickets button is hidden on the website event page |
| `is_ticketed` | `Boolean` | No | Whether a ticket is required to attend the event |
| `location` | `Object` | No | Where the event takes place |
| `off_sale_at` | `Object` | No | Date and time the event goes off sale |
| `on_sale_at` | `Object` | No | Date and time the event goes on sale |
| `short_description` | `Object` | No | Brief description of the event |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `start_at` | `Object` | No | The date the event occurrence begins |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `title_display` | `Object` | No | The name of this event formatted with HTML (optional) |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `id` | `String` | No | Unique identifier of this resource. |
| `is_affiliate_group` | `Boolean` | No | Whether this program represents an affiliate group |
| `is_event_host` | `Boolean` | No | Whether this program represents an event host |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |

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
| `aic_end_at` | `Object` | No | Date the exhibition closed at the Art Institute of Chicago |
| `aic_start_at` | `Object` | No | Date the exhibition opened at the Art Institute of Chicago |
| `alt_image_ids` | `Object` | No | Unique identifiers of all non-preferred images of this exhibition. |
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `artist_ids` | `Object` | No | Unique identifiers of the artist agent records representing who was shown in the exhibition |
| `artwork_ids` | `Object` | No | Unique identifiers of the artworks that were part of the exhibition |
| `artwork_titles` | `Object` | No | Names of the artworks that were part of the exhibition |
| `document_ids` | `Object` | No | Unique identifiers of assets that serve as documentation for this exhibition |
| `gallery_id` | `String` | No | Unique identifier of the gallery that mainly housed the exhibition |
| `gallery_title` | `Object` | No | The name of the gallery that mainly housed the exhibition |
| `id` | `String` | No | Unique identifier of this resource. |
| `image_id` | `String` | No | Unique identifier of the preferred image to use to represent this exhibition |
| `image_url` | `Object` | No | URL to the hero image from the website |
| `is_featured` | `Boolean` | No | Is this exhibition currently featured on our website? |
| `is_published` | `Boolean` | No | Is this exhibition currently published on our website? |
| `position` | `Object` | No | Numering position represnting the order in which this exhibition is featured on the website |
| `short_description` | `Object` | No | Brief explanation of what this exhibition is |
| `site_ids` | `Object` | No | Unique identifiers of the microsites this exhibition is a part of |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `status` | `Object` | No | Whether the exhibition is open or closed |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | No | URL to this exhibition on our website |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `floor` | `Object` | No | The level the gallery is on, e.g., 1, 2, 3, or LL |
| `id` | `String` | No | Unique identifier of this resource. |
| `is_closed` | `Boolean` | No | Whether the gallery is currently closed |
| `latitude` | `Float` | No | Latitude coordinate of the center of the room |
| `latlon` | `Object` | No | Latitude and longitude coordinates of the center of the room |
| `longitude` | `Float` | No | Longitude coordinate of the center of the room |
| `number` | `Object` | No | The gallery's room number. |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | `String` | No | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `copy` | `Object` | No | The text of the page |
| `id` | `String` | No | Unique identifier of this resource. |
| `search_tags` | `Object` | No | Editor-specified list of tags to aid in internal search |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | No | The URL to this page on our website |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `copy` | `Object` | No | The text of the highlight description |
| `id` | `String` | No | Unique identifier of this resource. |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |

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
| `additional_text` | `Object` | No | Additional information about the hours |
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `friday_is_closed` | `Object` | No | Whether the museum is closed on Fridays |
| `friday_member_close` | `Object` | No | The time member hours ends on Fridays |
| `friday_member_open` | `Object` | No | The time member hours starts on Fridays |
| `friday_public_close` | `Object` | No | The time public hours ends on Fridays |
| `friday_public_open` | `Object` | No | The time public hours starts on Fridays |
| `id` | `String` | No | Unique identifier of this resource. |
| `monday_is_closed` | `Object` | No | Whether the museum is closed on Mondays |
| `monday_member_close` | `Object` | No | The time member hours ends on Mondays |
| `monday_member_open` | `Object` | No | The time member hours starts on Mondays |
| `monday_public_close` | `Object` | No | The time public hours ends on Mondays |
| `monday_public_open` | `Object` | No | The time public hours starts on Mondays |
| `saturday_is_closed` | `Object` | No | Whether the museum is closed on Saturdays |
| `saturday_member_close` | `Object` | No | The time member hours ends on Saturdays |
| `saturday_member_open` | `Object` | No | The time member hours starts on Saturdays |
| `saturday_public_close` | `Object` | No | The time public hours ends on Saturdays |
| `saturday_public_open` | `Object` | No | The time public hours starts on Saturdays |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `summary` | `Object` | No | Readable summary of the hours |
| `sunday_is_closed` | `Object` | No | Whether the museum is closed on Sundays |
| `sunday_member_close` | `Object` | No | The time member hours ends on Sundays |
| `sunday_member_open` | `Object` | No | The time member hours starts on Sundays |
| `sunday_public_close` | `Object` | No | The time public hours ends on Sundays |
| `sunday_public_open` | `Object` | No | The time public hours starts on Sundays |
| `thursday_is_closed` | `Object` | No | Whether the museum is closed on Thursdays |
| `thursday_member_close` | `Object` | No | The time member hours ends on Thursdays |
| `thursday_member_open` | `Object` | No | The time member hours starts on Thursdays |
| `thursday_public_close` | `Object` | No | The time public hours ends on Thursdays |
| `thursday_public_open` | `Object` | No | The time public hours starts on Thursdays |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `tuesday_is_closed` | `Object` | No | Whether the museum is closed on Tuesdays |
| `tuesday_member_close` | `Object` | No | The time member hours ends on Tuesdays |
| `tuesday_member_open` | `Object` | No | The time member hours starts on Tuesdays |
| `tuesday_public_close` | `Object` | No | The time public hours ends on Tuesdays |
| `tuesday_public_open` | `Object` | No | The time public hours starts on Tuesdays |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |
| `wednesday_is_closed` | `Object` | No | Whether the museum is closed on Wednesdays |
| `wednesday_member_close` | `Object` | No | The time member hours ends on Wednesdays |
| `wednesday_member_open` | `Object` | No | The time member hours starts on Wednesdays |
| `wednesday_public_close` | `Object` | No | The time public hours ends on Wednesdays |
| `wednesday_public_open` | `Object` | No | The time public hours starts on Wednesdays |

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
| `ahash` | `Object` | No | Image hash generated using ahash algorithm with 64 boolean subfields |
| `alt_text` | `Object` | No | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `artwork_ids` | `Object` | No | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `Object` | No | Names of the artworks associated with this asset |
| `color` | `Object` | No | Dominant color of this image in HSL |
| `colorfulness` | `Object` | No | Unbounded positive float representing an abstract measure of colorfulness. |
| `content` | `Object` | No | Text of or URL to the contents of this asset |
| `content_e_tag` | `Object` | No | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `Object` | No | Asset-specific copyright information |
| `fingerprint` | `Object` | No | Image hashes: aHash, dHash, pHash, wHash |
| `height` | `Float` | No | Native height of the image |
| `id` | `String` | No | Unique identifier of this resource. |
| `iiif_url` | `Object` | No | IIIF URL of this image |
| `is_educational_resource` | `Boolean` | No | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `Boolean` | No | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `Boolean` | No | Whether this resource is considered to be educational |
| `lake_guid` | `Object` | No | Unique UUID of this resource in LAKE, our DAMS. |
| `lqip` | `Object` | No | Low-quality image placeholder (LQIP). |
| `phash` | `Object` | No | Image hash generated using phash algorithm with 64 boolean subfields |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `type` | `Object` | No | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |
| `width` | `Float` | No | Native width of the image |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `copy` | `Object` | No | The text of the page |
| `id` | `String` | No | Unique identifier of this resource. |
| `search_tags` | `Object` | No | Editor-specified list of tags to aid in internal search |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | No | The URL to this page on our website |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `id` | `String` | No | Unique identifier of this resource. |
| `latitude` | `Float` | No | Latitude coordinate of the center of the room |
| `longitude` | `Float` | No | Longitude coordinate of the center of the room |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | `String` | No | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `copy` | `Object` | No | The text of the page |
| `id` | `String` | No | Unique identifier of this resource. |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | No | The URL to this page on our website |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `copy` | `Object` | No | The text of the page |
| `id` | `String` | No | Unique identifier of this resource. |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | No | The URL to this page on our website |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `artist_ids` | `Object` | No | Unique identifiers of the artists associated with this product |
| `artwork_ids` | `Object` | No | Unique identifiers of the artworks associated with this product |
| `description` | `String` | No | Explanation of what this product is |
| `exhibition_ids` | `Object` | No | Unique identifiers of the exhibitions associated with this product |
| `external_sku` | `Object` | No | Numeric product identification code of a machine-readable barcode, when the customer sku differs from our internal one |
| `id` | `String` | No | Unique identifier of this resource. |
| `image_url` | `Object` | No | URL of an image for this product |
| `max_compare_at_price` | `Object` | No | Number indicating how much the most expensive variant of a product cost before a sale |
| `max_current_price` | `Object` | No | Number indicating how much the most expensive variant of a product costs right now |
| `min_compare_at_price` | `Object` | No | Number indicating how much the least expensive variant of a product cost before a sale |
| `min_current_price` | `Object` | No | Number indicating how much the least expensive variant of a product costs right now |
| `price_display` | `Object` | No | Explanation of what this product is |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | No | URL of this product in the shop |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `id` | `String` | No | Unique identifier of this resource. |
| `section_ids` | `Object` | No | Unique identifiers of the sections of this publication |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | No | URL to the publication |

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
| `api_id` | `String` | No | API unique identifier |
| `api_link` | `Object` | No | URL to this recource in the API |
| `api_model` | `Object` | No | Name of the model the resource represents |
| `id` | `String` | No | Unique identifier within the search index |
| `is_boosted` | `Boolean` | No | Whether this record has been flagged to be boosted |
| `score` | `Float` | No | Search index ranking of the result |
| `thumbnail` | `Object` | No | Metadata on the image representing this record |
| `timestamp` | `Object` | No | Date this record was last updated in the API |
| `title` | `String` | No | The name of this resource |

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
| `accession` | `Object` | No | An accession number parsed from the title or tombstone |
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `artwork_id` | `String` | No | Unique identifier of the artwork with which this section is associated |
| `content` | `Object` | No | Content of this section in plaintext |
| `generic_page_id` | `String` | No | Unique identifier of the page on the website that represents the publication this section belongs to |
| `id` | `String` | No | Unique identifier of this resource. |
| `publication_id` | `String` | No | Unique identifier of the publication this section belongs to |
| `publication_title` | `Object` | No | Name of the publication this section belongs to |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | No | URL to the section |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `artwork_ids` | `Object` | No | Unique identifiers of the artworks this site is associated with |
| `artwork_titles` | `Object` | No | Names of the artworks this site is associated with |
| `description` | `String` | No | Explanation of what this site is |
| `exhibition_ids` | `Object` | No | Unique identifier of the exhibitions this site is associated with |
| `exhibition_titles` | `Object` | No | Names of the exhibitions this site is associated with |
| `id` | `String` | No | Unique identifier of this resource. |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | No | URL to this site |

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
| `alt_text` | `Object` | No | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `artwork_ids` | `Object` | No | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `Object` | No | Names of the artworks associated with this asset |
| `content` | `Object` | No | Text of or URL to the contents of this asset |
| `content_e_tag` | `Object` | No | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `Object` | No | Asset-specific copyright information |
| `id` | `String` | No | Unique identifier of this resource. |
| `is_educational_resource` | `Boolean` | No | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `Boolean` | No | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `Boolean` | No | Whether this resource is considered to be educational |
| `lake_guid` | `Object` | No | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | Name of this mobile audio file – derived from the artwork and tour titles |
| `transcript` | `Object` | No | Text transcription of the audio file |
| `type` | `Object` | No | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | No | URL to the audio file |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `id` | `String` | No | Unique identifier of this resource. |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | No | The URL to this page on our website |

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
| `alt_text` | `Object` | No | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `artwork_ids` | `Object` | No | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `Object` | No | Names of the artworks associated with this asset |
| `content` | `Object` | No | Text of or URL to the contents of this asset |
| `content_e_tag` | `Object` | No | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `Object` | No | Asset-specific copyright information |
| `id` | `String` | No | Unique identifier of this resource. |
| `is_educational_resource` | `Boolean` | No | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `Boolean` | No | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `Boolean` | No | Whether this resource is considered to be educational |
| `lake_guid` | `Object` | No | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `type` | `Object` | No | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `artist_titles` | `Object` | No | Names of the artists of the artworks featured in this tour's tour stops |
| `artwork_titles` | `Object` | No | Names of the artworks featured in this tour's tour stops |
| `description` | `String` | No | Explanation of what the tour is |
| `id` | `String` | No | Unique identifier of this resource. |
| `image` | `Object` | No | The main image for the tour |
| `intro` | `Object` | No | Text introducing the tour |
| `intro_link` | `Object` | No | Link to the audio file of the introduction |
| `intro_transcript` | `Object` | No | Transcript of the introduction audio to the tour |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |
| `weight` | `Float` | No | Number representing this tour's sort order |

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
| `alt_text` | `Object` | No | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `Object` | No | REST API link for this resource |
| `api_model` | `Object` | No | REST API resource type or endpoint |
| `artwork_ids` | `Object` | No | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `Object` | No | Names of the artworks associated with this asset |
| `content` | `Object` | No | Text of or URL to the contents of this asset |
| `content_e_tag` | `Object` | No | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `Object` | No | Asset-specific copyright information |
| `id` | `String` | No | Unique identifier of this resource. |
| `is_educational_resource` | `Boolean` | No | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `Boolean` | No | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `Boolean` | No | Whether this resource is considered to be educational |
| `lake_guid` | `Object` | No | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `Object` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | No | Date and time the record was updated in the aggregator search index |
| `title` | `String` | No | The name of this resource |
| `type` | `Object` | No | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `Object` | No | Date and time the record was updated in the aggregator database |

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

