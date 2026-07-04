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
const agent = client.agent
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.agent.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.agent.load({ id: 'agent_id' })
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
const agent_role = client.agent_role
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.agent_role.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.agent_role.load({ id: 'agent_role_id' })
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
const agent_type = client.agent_type
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.agent_type.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.agent_type.load({ id: 'agent_type_id' })
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
const article = client.article
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.article.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.article.load({ id: 'article_id' })
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
const artwork = client.artwork
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
| `copyright_notice` | ``$ANY`` | No |  |
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.artwork.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.artwork.load({ id: 'artwork_id' })
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
const artwork_date_qualifier = client.artwork_date_qualifier
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.artwork_date_qualifier.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.artwork_date_qualifier.load({ id: 'artwork_date_qualifier_id' })
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
const artwork_place_qualifier = client.artwork_place_qualifier
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.artwork_place_qualifier.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.artwork_place_qualifier.load({ id: 'artwork_place_qualifier_id' })
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
const artwork_type = client.artwork_type
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.artwork_type.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.artwork_type.load({ id: 'artwork_type_id' })
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
const category_term = client.category_term
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.category_term.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.category_term.load({ id: 'category_term_id' })
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
const digital_publication = client.digital_publication
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.digital_publication.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.digital_publication.load({ id: 'digital_publication_id' })
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
const digital_publication_article = client.digital_publication_article
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.digital_publication_article.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.digital_publication_article.load({ id: 'digital_publication_article_id' })
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
const educator_resource = client.educator_resource
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.educator_resource.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.educator_resource.load({ id: 'educator_resource_id' })
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
const event = client.event
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.event.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.event.load({ id: 'event_id' })
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
const event_occurrence = client.event_occurrence
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.event_occurrence.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.event_occurrence.load({ id: 'event_occurrence_id' })
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
const event_program = client.event_program
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.event_program.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.event_program.load({ id: 'event_program_id' })
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
const exhibition = client.exhibition
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.exhibition.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.exhibition.load({ id: 'exhibition_id' })
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
const gallery = client.gallery
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.gallery.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.gallery.load({ id: 'gallery_id' })
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
const generic_page = client.generic_page
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.generic_page.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.generic_page.load({ id: 'generic_page_id' })
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
const highlight = client.highlight
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.highlight.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.highlight.load({ id: 'highlight_id' })
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
const hour = client.hour
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.hour.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.hour.load({ id: 'hour_id' })
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
const image = client.image
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.image.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.image.load({ id: 'image_id' })
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
const landing_page = client.landing_page
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.landing_page.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.landing_page.load({ id: 'landing_page_id' })
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
const place = client.place
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.place.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.place.load({ id: 'place_id' })
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
const press_release = client.press_release
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.press_release.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.press_release.load({ id: 'press_release_id' })
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
const printed_publication = client.printed_publication
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.printed_publication.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.printed_publication.load({ id: 'printed_publication_id' })
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
const product = client.product
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.product.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.product.load({ id: 'product_id' })
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
const publication = client.publication
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.publication.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.publication.load({ id: 'publication_id' })
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
const search = client.search
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.search.list()
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
const section = client.section
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.section.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.section.load({ id: 'section_id' })
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
const site = client.site
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.site.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.site.load({ id: 'site_id' })
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
const sound = client.sound
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.sound.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.sound.load({ id: 'sound_id' })
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
const static_page = client.static_page
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.static_page.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.static_page.load({ id: 'static_page_id' })
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
const text = client.text
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.text.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.text.load({ id: 'text_id' })
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
const tour = client.tour
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.tour.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.tour.load({ id: 'tour_id' })
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
const video = client.video
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.video.list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.video.load({ id: 'video_id' })
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

