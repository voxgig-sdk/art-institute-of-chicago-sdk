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
| `alt_titles` | `any` | No |  |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `birth_date` | `any` | No |  |
| `death_date` | `any` | No |  |
| `description` | `string` | No |  |
| `id` | `string` | No |  |
| `is_artist` | `bool` | No |  |
| `sort_title` | `any` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `ulan_id` | `string` | No |  |
| `updated_at` | `any` | No |  |

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
| `has_advanced_imaging` | `bool` | No |  |
| `has_educational_resources` | `bool` | No |  |
| `has_multimedia_resources` | `bool` | No |  |
| `has_not_been_viewed_much` | `bool` | No |  |
| `id` | `string` | No |  |
| `image_embedding` | `any` | No |  |
| `image_id` | `string` | No |  |
| `inscriptions` | `any` | No |  |
| `internal_department_id` | `string` | No |  |
| `is_boosted` | `bool` | No |  |
| `is_on_view` | `bool` | No |  |
| `is_public_domain` | `bool` | No |  |
| `is_zoomable` | `bool` | No |  |
| `latitude` | `float64` | No |  |
| `latlon` | `any` | No |  |
| `longitude` | `float64` | No |  |
| `main_reference_number` | `int` | No |  |
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
| `is_private` | `bool` | No |  |
| `is_sales_button_hidden` | `bool` | No |  |
| `is_ticketed` | `bool` | No |  |
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
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `id` | `string` | No |  |
| `is_affiliate_group` | `bool` | No |  |
| `is_event_host` | `bool` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |

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
| `is_featured` | `bool` | No |  |
| `is_published` | `bool` | No |  |
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
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `floor` | `any` | No |  |
| `id` | `string` | No |  |
| `is_closed` | `bool` | No |  |
| `latitude` | `float64` | No |  |
| `latlon` | `any` | No |  |
| `longitude` | `float64` | No |  |
| `number` | `any` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `tgn_id` | `string` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |

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
| `height` | `float64` | No |  |
| `id` | `string` | No |  |
| `iiif_url` | `any` | No |  |
| `is_educational_resource` | `bool` | No |  |
| `is_multimedia_resource` | `bool` | No |  |
| `is_teacher_resource` | `bool` | No |  |
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
| `width` | `float64` | No |  |

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
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `id` | `string` | No |  |
| `latitude` | `float64` | No |  |
| `longitude` | `float64` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `tgn_id` | `string` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `updated_at` | `any` | No |  |

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
| `api_id` | `string` | No |  |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `id` | `string` | No |  |
| `is_boosted` | `bool` | No |  |
| `score` | `float64` | No |  |
| `thumbnail` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |

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
| `alt_text` | `any` | No |  |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `artwork_ids` | `any` | No |  |
| `artwork_titles` | `any` | No |  |
| `content` | `any` | No |  |
| `content_e_tag` | `any` | No |  |
| `credit_line` | `any` | No |  |
| `id` | `string` | No |  |
| `is_educational_resource` | `bool` | No |  |
| `is_multimedia_resource` | `bool` | No |  |
| `is_teacher_resource` | `bool` | No |  |
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
| `alt_text` | `any` | No |  |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `artwork_ids` | `any` | No |  |
| `artwork_titles` | `any` | No |  |
| `content` | `any` | No |  |
| `content_e_tag` | `any` | No |  |
| `credit_line` | `any` | No |  |
| `id` | `string` | No |  |
| `is_educational_resource` | `bool` | No |  |
| `is_multimedia_resource` | `bool` | No |  |
| `is_teacher_resource` | `bool` | No |  |
| `lake_guid` | `any` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `type` | `any` | No |  |
| `updated_at` | `any` | No |  |

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
| `weight` | `float64` | No |  |

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
| `alt_text` | `any` | No |  |
| `api_link` | `any` | No |  |
| `api_model` | `any` | No |  |
| `artwork_ids` | `any` | No |  |
| `artwork_titles` | `any` | No |  |
| `content` | `any` | No |  |
| `content_e_tag` | `any` | No |  |
| `credit_line` | `any` | No |  |
| `id` | `string` | No |  |
| `is_educational_resource` | `bool` | No |  |
| `is_multimedia_resource` | `bool` | No |  |
| `is_teacher_resource` | `bool` | No |  |
| `lake_guid` | `any` | No |  |
| `source_updated_at` | `any` | No |  |
| `suggest_autocomplete_all` | `any` | No |  |
| `suggest_autocomplete_boosted` | `any` | No |  |
| `timestamp` | `any` | No |  |
| `title` | `string` | No |  |
| `type` | `any` | No |  |
| `updated_at` | `any` | No |  |

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

