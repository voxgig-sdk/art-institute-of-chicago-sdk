# ArtInstituteOfChicago Golang SDK Reference

Complete API reference for the ArtInstituteOfChicago Golang SDK.


## ArtInstituteOfChicagoSDK

### Constructor

```go
func NewArtInstituteOfChicagoSDK(options map[string]any) *ArtInstituteOfChicagoSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *ArtInstituteOfChicagoSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *ArtInstituteOfChicagoSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `Agent(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `Agent` entity instance. Pass `nil` for no initial data.

#### `AgentRole(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `AgentRole` entity instance. Pass `nil` for no initial data.

#### `AgentType(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `AgentType` entity instance. Pass `nil` for no initial data.

#### `Article(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `Article` entity instance. Pass `nil` for no initial data.

#### `Artwork(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `Artwork` entity instance. Pass `nil` for no initial data.

#### `ArtworkDateQualifier(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `ArtworkDateQualifier` entity instance. Pass `nil` for no initial data.

#### `ArtworkPlaceQualifier(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `ArtworkPlaceQualifier` entity instance. Pass `nil` for no initial data.

#### `ArtworkType(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `ArtworkType` entity instance. Pass `nil` for no initial data.

#### `CategoryTerm(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `CategoryTerm` entity instance. Pass `nil` for no initial data.

#### `DigitalPublication(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `DigitalPublication` entity instance. Pass `nil` for no initial data.

#### `DigitalPublicationArticle(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `DigitalPublicationArticle` entity instance. Pass `nil` for no initial data.

#### `EducatorResource(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `EducatorResource` entity instance. Pass `nil` for no initial data.

#### `Event(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `Event` entity instance. Pass `nil` for no initial data.

#### `EventOccurrence(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `EventOccurrence` entity instance. Pass `nil` for no initial data.

#### `EventProgram(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `EventProgram` entity instance. Pass `nil` for no initial data.

#### `Exhibition(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `Exhibition` entity instance. Pass `nil` for no initial data.

#### `Gallery(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `Gallery` entity instance. Pass `nil` for no initial data.

#### `GenericPage(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `GenericPage` entity instance. Pass `nil` for no initial data.

#### `Highlight(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `Highlight` entity instance. Pass `nil` for no initial data.

#### `Hour(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `Hour` entity instance. Pass `nil` for no initial data.

#### `Image(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `Image` entity instance. Pass `nil` for no initial data.

#### `LandingPage(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `LandingPage` entity instance. Pass `nil` for no initial data.

#### `Place(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `Place` entity instance. Pass `nil` for no initial data.

#### `PressRelease(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `PressRelease` entity instance. Pass `nil` for no initial data.

#### `PrintedPublication(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `PrintedPublication` entity instance. Pass `nil` for no initial data.

#### `Product(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `Product` entity instance. Pass `nil` for no initial data.

#### `Publication(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `Publication` entity instance. Pass `nil` for no initial data.

#### `Search(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `Search` entity instance. Pass `nil` for no initial data.

#### `Section(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `Section` entity instance. Pass `nil` for no initial data.

#### `Site(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `Site` entity instance. Pass `nil` for no initial data.

#### `Sound(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `Sound` entity instance. Pass `nil` for no initial data.

#### `StaticPage(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `StaticPage` entity instance. Pass `nil` for no initial data.

#### `Text(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `Text` entity instance. Pass `nil` for no initial data.

#### `Tour(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `Tour` entity instance. Pass `nil` for no initial data.

#### `Video(data map[string]any) ArtInstituteOfChicagoEntity`

Create a new `Video` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## AgentEntity

```go
agent := client.Agent(nil)
fmt.Println(agent.GetName()) // "agent"
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
| `is_artist` | `bool` | No | Whether the agent is an artist. |
| `sort_title` | `any` | No | Sortable name for this agent, typically with last name first. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `ulan_id` | `string` | No | Unique identifier of this agent in Getty's ULAN |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Agent(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Agent(nil).Load(map[string]any{"id": "agent_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `AgentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## AgentRoleEntity

```go
agentRole := client.AgentRole(nil)
fmt.Println(agentRole.GetName()) // "agent_role"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.AgentRole(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.AgentRole(nil).Load(map[string]any{"id": "agent_role_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `AgentRoleEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## AgentTypeEntity

```go
agentType := client.AgentType(nil)
fmt.Println(agentType.GetName()) // "agent_type"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.AgentType(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.AgentType(nil).Load(map[string]any{"id": "agent_type_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `AgentTypeEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ArticleEntity

```go
article := client.Article(nil)
fmt.Println(article.GetName()) // "article"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Article(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Article(nil).Load(map[string]any{"id": "article_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ArticleEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ArtworkEntity

```go
artwork := client.Artwork(nil)
fmt.Println(artwork.GetName()) // "artwork"
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
| `has_advanced_imaging` | `bool` | No | Whether this artwork is enhanced with 3D models, 360 image sequences, Mirador views, etc. |
| `has_educational_resources` | `bool` | No | Whether this artwork has any documents tagged as educational |
| `has_multimedia_resources` | `bool` | No | Whether this artwork has any associated microsites, digital publications, or documents tagged as multimedia |
| `has_not_been_viewed_much` | `bool` | No | Whether the artwork hasn't been visited on our website very much |
| `id` | `string` | No | Unique identifier of this resource. |
| `image_embedding` | `any` | No | The generated embeddings describing the artwork image |
| `image_id` | `string` | No | Unique identifier of the preferred image to use to represent this work |
| `inscriptions` | `any` | No | A description of distinguishing or identifying physical markings that are on the work |
| `internal_department_id` | `string` | No | An internal department id we use for analytics. |
| `is_boosted` | `bool` | No | Whether this document should be boosted in search |
| `is_on_view` | `bool` | No | Whether the work is on display |
| `is_public_domain` | `bool` | No | Whether the work is in the public domain, meaning it was created before copyrights existed or has left the copyright term |
| `is_zoomable` | `bool` | No | Whether images of the work are allowed to be displayed in a zoomable interface. |
| `latitude` | `float64` | No | Latitude coordinate of the location of this work in our galleries |
| `latlon` | `any` | No | Latitude and longitude coordinates of the location of this work in our galleries |
| `longitude` | `float64` | No | Longitude coordinate of the location of this work in our galleries |
| `main_reference_number` | `int` | No | Unique identifier assigned to the artwork upon acquisition |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Artwork(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Artwork(nil).Load(map[string]any{"id": "artwork_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ArtworkEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ArtworkDateQualifierEntity

```go
artworkDateQualifier := client.ArtworkDateQualifier(nil)
fmt.Println(artworkDateQualifier.GetName()) // "artwork_date_qualifier"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.ArtworkDateQualifier(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.ArtworkDateQualifier(nil).Load(map[string]any{"id": "artwork_date_qualifier_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ArtworkDateQualifierEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ArtworkPlaceQualifierEntity

```go
artworkPlaceQualifier := client.ArtworkPlaceQualifier(nil)
fmt.Println(artworkPlaceQualifier.GetName()) // "artwork_place_qualifier"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.ArtworkPlaceQualifier(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.ArtworkPlaceQualifier(nil).Load(map[string]any{"id": "artwork_place_qualifier_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ArtworkPlaceQualifierEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ArtworkTypeEntity

```go
artworkType := client.ArtworkType(nil)
fmt.Println(artworkType.GetName()) // "artwork_type"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.ArtworkType(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.ArtworkType(nil).Load(map[string]any{"id": "artwork_type_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ArtworkTypeEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CategoryTermEntity

```go
categoryTerm := client.CategoryTerm(nil)
fmt.Println(categoryTerm.GetName()) // "category_term"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.CategoryTerm(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.CategoryTerm(nil).Load(map[string]any{"id": "category_term_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CategoryTermEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## DigitalPublicationEntity

```go
digitalPublication := client.DigitalPublication(nil)
fmt.Println(digitalPublication.GetName()) // "digital_publication"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.DigitalPublication(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.DigitalPublication(nil).Load(map[string]any{"id": "digital_publication_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `DigitalPublicationEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## DigitalPublicationArticleEntity

```go
digitalPublicationArticle := client.DigitalPublicationArticle(nil)
fmt.Println(digitalPublicationArticle.GetName()) // "digital_publication_article"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.DigitalPublicationArticle(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.DigitalPublicationArticle(nil).Load(map[string]any{"id": "digital_publication_article_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `DigitalPublicationArticleEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EducatorResourceEntity

```go
educatorResource := client.EducatorResource(nil)
fmt.Println(educatorResource.GetName()) // "educator_resource"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.EducatorResource(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.EducatorResource(nil).Load(map[string]any{"id": "educator_resource_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EducatorResourceEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EventEntity

```go
event := client.Event(nil)
fmt.Println(event.GetName()) // "event"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Event(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Event(nil).Load(map[string]any{"id": "event_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EventEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EventOccurrenceEntity

```go
eventOccurrence := client.EventOccurrence(nil)
fmt.Println(eventOccurrence.GetName()) // "event_occurrence"
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
| `is_private` | `bool` | No | Whether the event is private. |
| `is_sales_button_hidden` | `bool` | No | Whether the buy tickets button is hidden on the website event page |
| `is_ticketed` | `bool` | No | Whether a ticket is required to attend the event |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.EventOccurrence(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.EventOccurrence(nil).Load(map[string]any{"id": "event_occurrence_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EventOccurrenceEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EventProgramEntity

```go
eventProgram := client.EventProgram(nil)
fmt.Println(eventProgram.GetName()) // "event_program"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `is_affiliate_group` | `bool` | No | Whether this program represents an affiliate group |
| `is_event_host` | `bool` | No | Whether this program represents an event host |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.EventProgram(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.EventProgram(nil).Load(map[string]any{"id": "event_program_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EventProgramEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ExhibitionEntity

```go
exhibition := client.Exhibition(nil)
fmt.Println(exhibition.GetName()) // "exhibition"
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
| `is_featured` | `bool` | No | Is this exhibition currently featured on our website? |
| `is_published` | `bool` | No | Is this exhibition currently published on our website? |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Exhibition(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Exhibition(nil).Load(map[string]any{"id": "exhibition_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ExhibitionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## GalleryEntity

```go
gallery := client.Gallery(nil)
fmt.Println(gallery.GetName()) // "gallery"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `floor` | `any` | No | The level the gallery is on, e.g., 1, 2, 3, or LL |
| `id` | `string` | No | Unique identifier of this resource. |
| `is_closed` | `bool` | No | Whether the gallery is currently closed |
| `latitude` | `float64` | No | Latitude coordinate of the center of the room |
| `latlon` | `any` | No | Latitude and longitude coordinates of the center of the room |
| `longitude` | `float64` | No | Longitude coordinate of the center of the room |
| `number` | `any` | No | The gallery's room number. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | `string` | No | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Gallery(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Gallery(nil).Load(map[string]any{"id": "gallery_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `GalleryEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## GenericPageEntity

```go
genericPage := client.GenericPage(nil)
fmt.Println(genericPage.GetName()) // "generic_page"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.GenericPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.GenericPage(nil).Load(map[string]any{"id": "generic_page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `GenericPageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## HighlightEntity

```go
highlight := client.Highlight(nil)
fmt.Println(highlight.GetName()) // "highlight"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Highlight(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Highlight(nil).Load(map[string]any{"id": "highlight_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `HighlightEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## HourEntity

```go
hour := client.Hour(nil)
fmt.Println(hour.GetName()) // "hour"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Hour(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Hour(nil).Load(map[string]any{"id": "hour_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `HourEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ImageEntity

```go
image := client.Image(nil)
fmt.Println(image.GetName()) // "image"
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
| `height` | `float64` | No | Native height of the image |
| `id` | `string` | No | Unique identifier of this resource. |
| `iiif_url` | `any` | No | IIIF URL of this image |
| `is_educational_resource` | `bool` | No | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | No | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | No | Whether this resource is considered to be educational |
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
| `width` | `float64` | No | Native width of the image |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Image(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Image(nil).Load(map[string]any{"id": "image_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ImageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## LandingPageEntity

```go
landingPage := client.LandingPage(nil)
fmt.Println(landingPage.GetName()) // "landing_page"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.LandingPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.LandingPage(nil).Load(map[string]any{"id": "landing_page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `LandingPageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PlaceEntity

```go
place := client.Place(nil)
fmt.Println(place.GetName()) // "place"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `any` | No | REST API link for this resource |
| `api_model` | `any` | No | REST API resource type or endpoint |
| `id` | `string` | No | Unique identifier of this resource. |
| `latitude` | `float64` | No | Latitude coordinate of the center of the room |
| `longitude` | `float64` | No | Longitude coordinate of the center of the room |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | `string` | No | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Place(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Place(nil).Load(map[string]any{"id": "place_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PlaceEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PressReleaseEntity

```go
pressRelease := client.PressRelease(nil)
fmt.Println(pressRelease.GetName()) // "press_release"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.PressRelease(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.PressRelease(nil).Load(map[string]any{"id": "press_release_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PressReleaseEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PrintedPublicationEntity

```go
printedPublication := client.PrintedPublication(nil)
fmt.Println(printedPublication.GetName()) // "printed_publication"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.PrintedPublication(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.PrintedPublication(nil).Load(map[string]any{"id": "printed_publication_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PrintedPublicationEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ProductEntity

```go
product := client.Product(nil)
fmt.Println(product.GetName()) // "product"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Product(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Product(nil).Load(map[string]any{"id": "product_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ProductEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PublicationEntity

```go
publication := client.Publication(nil)
fmt.Println(publication.GetName()) // "publication"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Publication(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Publication(nil).Load(map[string]any{"id": "publication_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PublicationEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SearchEntity

```go
search := client.Search(nil)
fmt.Println(search.GetName()) // "search"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_id` | `string` | No | API unique identifier |
| `api_link` | `any` | No | URL to this recource in the API |
| `api_model` | `any` | No | Name of the model the resource represents |
| `id` | `string` | No | Unique identifier within the search index |
| `is_boosted` | `bool` | No | Whether this record has been flagged to be boosted |
| `score` | `float64` | No | Search index ranking of the result |
| `thumbnail` | `any` | No | Metadata on the image representing this record |
| `timestamp` | `any` | No | Date this record was last updated in the API |
| `title` | `string` | No | The name of this resource |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Search(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SearchEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SectionEntity

```go
section := client.Section(nil)
fmt.Println(section.GetName()) // "section"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Section(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Section(nil).Load(map[string]any{"id": "section_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SectionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SiteEntity

```go
site := client.Site(nil)
fmt.Println(site.GetName()) // "site"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Site(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Site(nil).Load(map[string]any{"id": "site_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SiteEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SoundEntity

```go
sound := client.Sound(nil)
fmt.Println(sound.GetName()) // "sound"
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
| `is_educational_resource` | `bool` | No | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | No | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | No | Whether this resource is considered to be educational |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Sound(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Sound(nil).Load(map[string]any{"id": "sound_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SoundEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## StaticPageEntity

```go
staticPage := client.StaticPage(nil)
fmt.Println(staticPage.GetName()) // "static_page"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.StaticPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.StaticPage(nil).Load(map[string]any{"id": "static_page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `StaticPageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## TextEntity

```go
text := client.Text(nil)
fmt.Println(text.GetName()) // "text"
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
| `is_educational_resource` | `bool` | No | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | No | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | No | Whether this resource is considered to be educational |
| `lake_guid` | `any` | No | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `type` | `any` | No | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Text(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Text(nil).Load(map[string]any{"id": "text_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `TextEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## TourEntity

```go
tour := client.Tour(nil)
fmt.Println(tour.GetName()) // "tour"
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
| `weight` | `float64` | No | Number representing this tour's sort order |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Tour(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Tour(nil).Load(map[string]any{"id": "tour_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `TourEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## VideoEntity

```go
video := client.Video(nil)
fmt.Println(video.GetName()) // "video"
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
| `is_educational_resource` | `bool` | No | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | No | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | No | Whether this resource is considered to be educational |
| `lake_guid` | `any` | No | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `string` | No | The name of this resource |
| `type` | `any` | No | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `any` | No | Date and time the record was updated in the aggregator database |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Video(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Video(nil).Load(map[string]any{"id": "video_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `VideoEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```go
client := sdk.NewArtInstituteOfChicagoSDK(map[string]any{
    "feature": map[string]any{
        "test": map[string]any{"active": true},
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

