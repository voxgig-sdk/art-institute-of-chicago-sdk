# ArtInstituteOfChicago TypeScript SDK Reference

Complete API reference for the ArtInstituteOfChicago TypeScript SDK.


## ArtInstituteOfChicagoSDK

### Constructor

```ts
new ArtInstituteOfChicagoSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `ArtInstituteOfChicagoSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = ArtInstituteOfChicagoSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `ArtInstituteOfChicagoSDK` instance in test mode.


### Instance Methods

#### `Agent(data?: object)`

Create a new `Agent` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AgentEntity` instance.

#### `AgentRole(data?: object)`

Create a new `AgentRole` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AgentRoleEntity` instance.

#### `AgentType(data?: object)`

Create a new `AgentType` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AgentTypeEntity` instance.

#### `Article(data?: object)`

Create a new `Article` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ArticleEntity` instance.

#### `Artwork(data?: object)`

Create a new `Artwork` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ArtworkEntity` instance.

#### `ArtworkDateQualifier(data?: object)`

Create a new `ArtworkDateQualifier` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ArtworkDateQualifierEntity` instance.

#### `ArtworkPlaceQualifier(data?: object)`

Create a new `ArtworkPlaceQualifier` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ArtworkPlaceQualifierEntity` instance.

#### `ArtworkType(data?: object)`

Create a new `ArtworkType` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ArtworkTypeEntity` instance.

#### `CategoryTerm(data?: object)`

Create a new `CategoryTerm` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CategoryTermEntity` instance.

#### `DigitalPublication(data?: object)`

Create a new `DigitalPublication` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `DigitalPublicationEntity` instance.

#### `DigitalPublicationArticle(data?: object)`

Create a new `DigitalPublicationArticle` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `DigitalPublicationArticleEntity` instance.

#### `EducatorResource(data?: object)`

Create a new `EducatorResource` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EducatorResourceEntity` instance.

#### `Event(data?: object)`

Create a new `Event` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EventEntity` instance.

#### `EventOccurrence(data?: object)`

Create a new `EventOccurrence` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EventOccurrenceEntity` instance.

#### `EventProgram(data?: object)`

Create a new `EventProgram` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EventProgramEntity` instance.

#### `Exhibition(data?: object)`

Create a new `Exhibition` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ExhibitionEntity` instance.

#### `Gallery(data?: object)`

Create a new `Gallery` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `GalleryEntity` instance.

#### `GenericPage(data?: object)`

Create a new `GenericPage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `GenericPageEntity` instance.

#### `Highlight(data?: object)`

Create a new `Highlight` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `HighlightEntity` instance.

#### `Hour(data?: object)`

Create a new `Hour` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `HourEntity` instance.

#### `Image(data?: object)`

Create a new `Image` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ImageEntity` instance.

#### `LandingPage(data?: object)`

Create a new `LandingPage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `LandingPageEntity` instance.

#### `Place(data?: object)`

Create a new `Place` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PlaceEntity` instance.

#### `PressRelease(data?: object)`

Create a new `PressRelease` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PressReleaseEntity` instance.

#### `PrintedPublication(data?: object)`

Create a new `PrintedPublication` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PrintedPublicationEntity` instance.

#### `Product(data?: object)`

Create a new `Product` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ProductEntity` instance.

#### `Publication(data?: object)`

Create a new `Publication` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PublicationEntity` instance.

#### `Search(data?: object)`

Create a new `Search` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SearchEntity` instance.

#### `Section(data?: object)`

Create a new `Section` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SectionEntity` instance.

#### `Site(data?: object)`

Create a new `Site` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SiteEntity` instance.

#### `Sound(data?: object)`

Create a new `Sound` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SoundEntity` instance.

#### `StaticPage(data?: object)`

Create a new `StaticPage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `StaticPageEntity` instance.

#### `Text(data?: object)`

Create a new `Text` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `TextEntity` instance.

#### `Tour(data?: object)`

Create a new `Tour` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `TourEntity` instance.

#### `Video(data?: object)`

Create a new `Video` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `VideoEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `ArtInstituteOfChicagoSDK.test()`.

**Returns:** `ArtInstituteOfChicagoSDK` instance in test mode.


---

## AgentEntity

```ts
const agent = client.Agent()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Agent().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Agent().load({ id: 'agent_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AgentEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## AgentRoleEntity

```ts
const agent_role = client.AgentRole()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.AgentRole().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.AgentRole().load({ id: 'agent_role_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AgentRoleEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## AgentTypeEntity

```ts
const agent_type = client.AgentType()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.AgentType().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.AgentType().load({ id: 'agent_type_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AgentTypeEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ArticleEntity

```ts
const article = client.Article()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Article().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Article().load({ id: 'article_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ArticleEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ArtworkEntity

```ts
const artwork = client.Artwork()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Artwork().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Artwork().load({ id: 'artwork_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ArtworkEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ArtworkDateQualifierEntity

```ts
const artwork_date_qualifier = client.ArtworkDateQualifier()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.ArtworkDateQualifier().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.ArtworkDateQualifier().load({ id: 'artwork_date_qualifier_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ArtworkDateQualifierEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ArtworkPlaceQualifierEntity

```ts
const artwork_place_qualifier = client.ArtworkPlaceQualifier()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.ArtworkPlaceQualifier().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.ArtworkPlaceQualifier().load({ id: 'artwork_place_qualifier_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ArtworkPlaceQualifierEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ArtworkTypeEntity

```ts
const artwork_type = client.ArtworkType()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.ArtworkType().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.ArtworkType().load({ id: 'artwork_type_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ArtworkTypeEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CategoryTermEntity

```ts
const category_term = client.CategoryTerm()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.CategoryTerm().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.CategoryTerm().load({ id: 'category_term_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CategoryTermEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## DigitalPublicationEntity

```ts
const digital_publication = client.DigitalPublication()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.DigitalPublication().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.DigitalPublication().load({ id: 'digital_publication_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `DigitalPublicationEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## DigitalPublicationArticleEntity

```ts
const digital_publication_article = client.DigitalPublicationArticle()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.DigitalPublicationArticle().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.DigitalPublicationArticle().load({ id: 'digital_publication_article_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `DigitalPublicationArticleEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EducatorResourceEntity

```ts
const educator_resource = client.EducatorResource()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.EducatorResource().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.EducatorResource().load({ id: 'educator_resource_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EducatorResourceEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EventEntity

```ts
const event = client.Event()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Event().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Event().load({ id: 'event_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EventEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EventOccurrenceEntity

```ts
const event_occurrence = client.EventOccurrence()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.EventOccurrence().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.EventOccurrence().load({ id: 'event_occurrence_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EventOccurrenceEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EventProgramEntity

```ts
const event_program = client.EventProgram()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.EventProgram().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.EventProgram().load({ id: 'event_program_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EventProgramEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ExhibitionEntity

```ts
const exhibition = client.Exhibition()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Exhibition().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Exhibition().load({ id: 'exhibition_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ExhibitionEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## GalleryEntity

```ts
const gallery = client.Gallery()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Gallery().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Gallery().load({ id: 'gallery_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `GalleryEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## GenericPageEntity

```ts
const generic_page = client.GenericPage()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.GenericPage().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.GenericPage().load({ id: 'generic_page_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `GenericPageEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## HighlightEntity

```ts
const highlight = client.Highlight()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Highlight().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Highlight().load({ id: 'highlight_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `HighlightEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## HourEntity

```ts
const hour = client.Hour()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Hour().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Hour().load({ id: 'hour_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `HourEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ImageEntity

```ts
const image = client.Image()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Image().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Image().load({ id: 'image_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ImageEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## LandingPageEntity

```ts
const landing_page = client.LandingPage()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.LandingPage().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.LandingPage().load({ id: 'landing_page_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `LandingPageEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PlaceEntity

```ts
const place = client.Place()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Place().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Place().load({ id: 'place_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PlaceEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PressReleaseEntity

```ts
const press_release = client.PressRelease()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.PressRelease().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.PressRelease().load({ id: 'press_release_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PressReleaseEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PrintedPublicationEntity

```ts
const printed_publication = client.PrintedPublication()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.PrintedPublication().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.PrintedPublication().load({ id: 'printed_publication_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PrintedPublicationEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ProductEntity

```ts
const product = client.Product()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Product().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Product().load({ id: 'product_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ProductEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PublicationEntity

```ts
const publication = client.Publication()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Publication().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Publication().load({ id: 'publication_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PublicationEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SearchEntity

```ts
const search = client.Search()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Search().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SearchEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SectionEntity

```ts
const section = client.Section()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Section().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Section().load({ id: 'section_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SectionEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SiteEntity

```ts
const site = client.Site()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Site().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Site().load({ id: 'site_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SiteEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SoundEntity

```ts
const sound = client.Sound()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Sound().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Sound().load({ id: 'sound_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SoundEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## StaticPageEntity

```ts
const static_page = client.StaticPage()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.StaticPage().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.StaticPage().load({ id: 'static_page_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `StaticPageEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## TextEntity

```ts
const text = client.Text()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Text().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Text().load({ id: 'text_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `TextEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## TourEntity

```ts
const tour = client.Tour()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Tour().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Tour().load({ id: 'tour_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `TourEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## VideoEntity

```ts
const video = client.Video()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Video().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Video().load({ id: 'video_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `VideoEntity` instance with the same client and
options.

#### `client()`

Return the parent `ArtInstituteOfChicagoSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ts
const client = new ArtInstituteOfChicagoSDK({
  feature: {
    test: { active: true },
  }
})
```

