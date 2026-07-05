# ArtInstituteOfChicago TypeScript SDK



The TypeScript SDK for the ArtInstituteOfChicago API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.Agent()` — each with a small set of operations (`list`, `load`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to npm. Install it from the GitHub
release tag (`ts/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/art-institute-of-chicago-sdk/releases](https://github.com/voxgig-sdk/art-institute-of-chicago-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ts
import { ArtInstituteOfChicagoSDK } from '@voxgig-sdk/art-institute-of-chicago'

const client = new ArtInstituteOfChicagoSDK()
```

### 2. List agent records

`list()` resolves to an array of Agent objects — iterate it directly:

```ts
const agents = await client.Agent().list()

for (const agent of agents) {
  console.log(agent)
}
```

### 3. Load an agent

`load()` returns the entity directly and throws on failure:

```ts
try {
  const agent = await client.Agent().load({ id: 'example_id' })
  console.log(agent)
} catch (err) {
  console.error('load failed:', err)
}
```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const agents = await client.Agent().list()
  console.log(agents)
} catch (err) {
  console.error('list failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
value or an `Error`, so check the result before using it:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (result instanceof Error) {
  throw result
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result instanceof Error) {
  throw result
}
if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```ts
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```ts
const client = ArtInstituteOfChicagoSDK.test()

const agent = await client.Agent().list()
// agent is a bare entity populated with mock response data
console.log(agent)
```

You can also use the instance method:

```ts
const client = new ArtInstituteOfChicagoSDK()
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.Agent()

// First call runs the operation and stores its result
await entity.list()

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data.id)
```

### Add custom middleware

Pass features via the `extend` option:

```ts
const logger = {
  hooks: {
    PreRequest: (ctx: any) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx: any) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new ArtInstituteOfChicagoSDK({
  extend: [logger],
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE
```

Then run:

```bash
cd ts && npm test
```


## Reference

### ArtInstituteOfChicagoSDK

#### Constructor

```ts
new ArtInstituteOfChicagoSDK(options?: {
  base?: string
  prefix?: string
  suffix?: string
  feature?: Record<string, { active: boolean }>
  extend?: Feature[]
})
```

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `Agent(data?)` | `AgentEntity` | Create an Agent entity instance. |
| `AgentRole(data?)` | `AgentRoleEntity` | Create an AgentRole entity instance. |
| `AgentType(data?)` | `AgentTypeEntity` | Create an AgentType entity instance. |
| `Article(data?)` | `ArticleEntity` | Create an Article entity instance. |
| `Artwork(data?)` | `ArtworkEntity` | Create an Artwork entity instance. |
| `ArtworkDateQualifier(data?)` | `ArtworkDateQualifierEntity` | Create an ArtworkDateQualifier entity instance. |
| `ArtworkPlaceQualifier(data?)` | `ArtworkPlaceQualifierEntity` | Create an ArtworkPlaceQualifier entity instance. |
| `ArtworkType(data?)` | `ArtworkTypeEntity` | Create an ArtworkType entity instance. |
| `CategoryTerm(data?)` | `CategoryTermEntity` | Create a CategoryTerm entity instance. |
| `DigitalPublication(data?)` | `DigitalPublicationEntity` | Create a DigitalPublication entity instance. |
| `DigitalPublicationArticle(data?)` | `DigitalPublicationArticleEntity` | Create a DigitalPublicationArticle entity instance. |
| `EducatorResource(data?)` | `EducatorResourceEntity` | Create an EducatorResource entity instance. |
| `Event(data?)` | `EventEntity` | Create an Event entity instance. |
| `EventOccurrence(data?)` | `EventOccurrenceEntity` | Create an EventOccurrence entity instance. |
| `EventProgram(data?)` | `EventProgramEntity` | Create an EventProgram entity instance. |
| `Exhibition(data?)` | `ExhibitionEntity` | Create an Exhibition entity instance. |
| `Gallery(data?)` | `GalleryEntity` | Create a Gallery entity instance. |
| `GenericPage(data?)` | `GenericPageEntity` | Create a GenericPage entity instance. |
| `Highlight(data?)` | `HighlightEntity` | Create a Highlight entity instance. |
| `Hour(data?)` | `HourEntity` | Create a Hour entity instance. |
| `Image(data?)` | `ImageEntity` | Create an Image entity instance. |
| `LandingPage(data?)` | `LandingPageEntity` | Create a LandingPage entity instance. |
| `Place(data?)` | `PlaceEntity` | Create a Place entity instance. |
| `PressRelease(data?)` | `PressReleaseEntity` | Create a PressRelease entity instance. |
| `PrintedPublication(data?)` | `PrintedPublicationEntity` | Create a PrintedPublication entity instance. |
| `Product(data?)` | `ProductEntity` | Create a Product entity instance. |
| `Publication(data?)` | `PublicationEntity` | Create a Publication entity instance. |
| `Search(data?)` | `SearchEntity` | Create a Search entity instance. |
| `Section(data?)` | `SectionEntity` | Create a Section entity instance. |
| `Site(data?)` | `SiteEntity` | Create a Site entity instance. |
| `Sound(data?)` | `SoundEntity` | Create a Sound entity instance. |
| `StaticPage(data?)` | `StaticPageEntity` | Create a StaticPage entity instance. |
| `Text(data?)` | `TextEntity` | Create a Text entity instance. |
| `Tour(data?)` | `TourEntity` | Create a Tour entity instance. |
| `Video(data?)` | `VideoEntity` | Create a Video entity instance. |
| `tester(testopts?, sdkopts?)` | `ArtInstituteOfChicagoSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `ArtInstituteOfChicagoSDK.test(testopts?, sdkopts?)` | `ArtInstituteOfChicagoSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `load(reqmatch?, ctrl?): Promise<Entity>` | Load a single entity by match criteria. |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): ArtInstituteOfChicagoSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `load` resolves to a single entity object.
- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```ts
{
  ok: boolean
  status: number
  headers: object
  data: any
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```ts
{
  url: string
  method: string
  headers: Record<string, string>
  body?: any
}
```

### Entities

#### Agent

| Field | Description |
| --- | --- |
| `alt_title` |  |
| `api_link` |  |
| `api_model` |  |
| `birth_date` |  |
| `death_date` |  |
| `description` |  |
| `id` |  |
| `is_artist` |  |
| `sort_title` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `ulan_id` |  |
| `updated_at` |  |

Operations: list, load.

API path: `/agents`

#### AgentRole

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: list, load.

API path: `/agent-roles`

#### AgentType

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: list, load.

API path: `/agent-types`

#### Article

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `copy` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: list, load.

API path: `/articles`

#### Artwork

| Field | Description |
| --- | --- |
| `alt_artist_id` |  |
| `alt_classification_id` |  |
| `alt_image_id` |  |
| `alt_material_id` |  |
| `alt_style_id` |  |
| `alt_subject_id` |  |
| `alt_technique_id` |  |
| `alt_title` |  |
| `api_link` |  |
| `api_model` |  |
| `artist_display` |  |
| `artist_id` |  |
| `artist_title` |  |
| `artwork_type_id` |  |
| `artwork_type_title` |  |
| `boost_rank` |  |
| `catalog_based_search_keyword_title` |  |
| `catalogue_display` |  |
| `category_id` |  |
| `category_title` |  |
| `classification_id` |  |
| `classification_title` |  |
| `color` |  |
| `colorfulness` |  |
| `copyright_notice` |  |
| `credit_line` |  |
| `date_display` |  |
| `date_end` |  |
| `date_qualifier_id` |  |
| `date_qualifier_title` |  |
| `date_start` |  |
| `department_id` |  |
| `department_title` |  |
| `description` |  |
| `dimension` |  |
| `dimensions_detail` |  |
| `document_id` |  |
| `edition` |  |
| `exhibition_history` |  |
| `fiscal_year` |  |
| `fiscal_year_deaccession` |  |
| `gallery_id` |  |
| `gallery_title` |  |
| `has_advanced_imaging` |  |
| `has_educational_resource` |  |
| `has_multimedia_resource` |  |
| `has_not_been_viewed_much` |  |
| `id` |  |
| `image_embedding` |  |
| `image_id` |  |
| `inscription` |  |
| `internal_department_id` |  |
| `is_boosted` |  |
| `is_on_view` |  |
| `is_public_domain` |  |
| `is_zoomable` |  |
| `latitude` |  |
| `latlon` |  |
| `longitude` |  |
| `main_reference_number` |  |
| `material_id` |  |
| `material_title` |  |
| `max_zoom_window_size` |  |
| `medium_display` |  |
| `nomisma_id` |  |
| `on_loan_display` |  |
| `pageview` |  |
| `pageviews_recent` |  |
| `place_of_origin` |  |
| `provenance_text` |  |
| `publication_history` |  |
| `publishing_verification_level` |  |
| `section_id` |  |
| `section_title` |  |
| `short_description` |  |
| `site_id` |  |
| `sound_id` |  |
| `source_updated_at` |  |
| `style_id` |  |
| `style_title` |  |
| `subject_id` |  |
| `subject_title` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `technique_id` |  |
| `technique_title` |  |
| `term_title` |  |
| `text_embedding` |  |
| `text_id` |  |
| `theme_title` |  |
| `thumbnail` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `video_id` |  |

Operations: list, load.

API path: `/artworks`

#### ArtworkDateQualifier

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: list, load.

API path: `/artwork-date-qualifiers`

#### ArtworkPlaceQualifier

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: list, load.

API path: `/artwork-place-qualifiers`

#### ArtworkType

| Field | Description |
| --- | --- |
| `aat_id` |  |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: list, load.

API path: `/artwork-types`

#### CategoryTerm

| Field | Description |
| --- | --- |
| `aat_id` |  |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `parent_id` |  |
| `source_updated_at` |  |
| `subtype` |  |
| `suggest_autocomplete_all` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: list, load.

API path: `/category-terms`

#### DigitalPublication

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `copy` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: list, load.

API path: `/digital-publications`

#### DigitalPublicationArticle

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `author_display` |  |
| `copy` |  |
| `digital_publication_id` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: list, load.

API path: `/digital-publication-articles`

#### EducatorResource

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `copy` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: list, load.

API path: `/educator-resources`

#### Event

| Field | Description |
| --- | --- |
| `alt_audience_id` |  |
| `alt_event_type_id` |  |
| `api_link` |  |
| `api_model` |  |
| `audience_id` |  |
| `buy_button_caption` |  |
| `buy_button_text` |  |
| `date_display` |  |
| `description` |  |
| `door_time` |  |
| `end_date` |  |
| `end_time` |  |
| `entrance` |  |
| `event_host_id` |  |
| `event_host_title` |  |
| `event_type_id` |  |
| `header_description` |  |
| `hero_caption` |  |
| `id` |  |
| `image_url` |  |
| `is_admission_required` |  |
| `is_after_hour` |  |
| `is_free` |  |
| `is_member_exclusive` |  |
| `is_private` |  |
| `is_registration_required` |  |
| `is_sales_button_hidden` |  |
| `is_sold_out` |  |
| `is_ticketed` |  |
| `is_virtual_event` |  |
| `join_url` |  |
| `layout_type` |  |
| `list_description` |  |
| `location` |  |
| `program_id` |  |
| `program_title` |  |
| `rsvp_link` |  |
| `search_tag` |  |
| `short_description` |  |
| `slug` |  |
| `source_updated_at` |  |
| `start_date` |  |
| `start_time` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `survey_url` |  |
| `ticketed_event_id` |  |
| `timestamp` |  |
| `title` |  |
| `title_display` |  |
| `updated_at` |  |
| `virtual_event_passcode` |  |
| `virtual_event_url` |  |

Operations: list, load.

API path: `/events`

#### EventOccurrence

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `button_caption` |  |
| `button_text` |  |
| `button_url` |  |
| `description` |  |
| `end_at` |  |
| `event_id` |  |
| `id` |  |
| `image_url` |  |
| `is_private` |  |
| `is_sales_button_hidden` |  |
| `is_ticketed` |  |
| `location` |  |
| `off_sale_at` |  |
| `on_sale_at` |  |
| `short_description` |  |
| `source_updated_at` |  |
| `start_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `title_display` |  |
| `updated_at` |  |

Operations: list, load.

API path: `/event-occurrences`

#### EventProgram

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `is_affiliate_group` |  |
| `is_event_host` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: list, load.

API path: `/event-programs`

#### Exhibition

| Field | Description |
| --- | --- |
| `aic_end_at` |  |
| `aic_start_at` |  |
| `alt_image_id` |  |
| `api_link` |  |
| `api_model` |  |
| `artist_id` |  |
| `artwork_id` |  |
| `artwork_title` |  |
| `document_id` |  |
| `gallery_id` |  |
| `gallery_title` |  |
| `id` |  |
| `image_id` |  |
| `image_url` |  |
| `is_featured` |  |
| `is_published` |  |
| `position` |  |
| `short_description` |  |
| `site_id` |  |
| `source_updated_at` |  |
| `status` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: list, load.

API path: `/exhibitions`

#### Gallery

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `floor` |  |
| `id` |  |
| `is_closed` |  |
| `latitude` |  |
| `latlon` |  |
| `longitude` |  |
| `number` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `tgn_id` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: list, load.

API path: `/galleries`

#### GenericPage

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `copy` |  |
| `id` |  |
| `search_tag` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: list, load.

API path: `/generic-pages`

#### Highlight

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `copy` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: list, load.

API path: `/highlights`

#### Hour

| Field | Description |
| --- | --- |
| `additional_text` |  |
| `api_link` |  |
| `api_model` |  |
| `friday_is_closed` |  |
| `friday_member_close` |  |
| `friday_member_open` |  |
| `friday_public_close` |  |
| `friday_public_open` |  |
| `id` |  |
| `monday_is_closed` |  |
| `monday_member_close` |  |
| `monday_member_open` |  |
| `monday_public_close` |  |
| `monday_public_open` |  |
| `saturday_is_closed` |  |
| `saturday_member_close` |  |
| `saturday_member_open` |  |
| `saturday_public_close` |  |
| `saturday_public_open` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `summary` |  |
| `sunday_is_closed` |  |
| `sunday_member_close` |  |
| `sunday_member_open` |  |
| `sunday_public_close` |  |
| `sunday_public_open` |  |
| `thursday_is_closed` |  |
| `thursday_member_close` |  |
| `thursday_member_open` |  |
| `thursday_public_close` |  |
| `thursday_public_open` |  |
| `timestamp` |  |
| `title` |  |
| `tuesday_is_closed` |  |
| `tuesday_member_close` |  |
| `tuesday_member_open` |  |
| `tuesday_public_close` |  |
| `tuesday_public_open` |  |
| `updated_at` |  |
| `wednesday_is_closed` |  |
| `wednesday_member_close` |  |
| `wednesday_member_open` |  |
| `wednesday_public_close` |  |
| `wednesday_public_open` |  |

Operations: list, load.

API path: `/hours`

#### Image

| Field | Description |
| --- | --- |
| `ahash` |  |
| `alt_text` |  |
| `api_link` |  |
| `api_model` |  |
| `artwork_id` |  |
| `artwork_title` |  |
| `color` |  |
| `colorfulness` |  |
| `content` |  |
| `content_e_tag` |  |
| `credit_line` |  |
| `fingerprint` |  |
| `height` |  |
| `id` |  |
| `iiif_url` |  |
| `is_educational_resource` |  |
| `is_multimedia_resource` |  |
| `is_teacher_resource` |  |
| `lake_guid` |  |
| `lqip` |  |
| `phash` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `type` |  |
| `updated_at` |  |
| `width` |  |

Operations: list, load.

API path: `/images`

#### LandingPage

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `copy` |  |
| `id` |  |
| `search_tag` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: list, load.

API path: `/landing-pages`

#### Place

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `latitude` |  |
| `longitude` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `tgn_id` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |

Operations: list, load.

API path: `/places`

#### PressRelease

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `copy` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: list, load.

API path: `/press-releases`

#### PrintedPublication

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `copy` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: list, load.

API path: `/printed-publications`

#### Product

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `artist_id` |  |
| `artwork_id` |  |
| `description` |  |
| `exhibition_id` |  |
| `external_sku` |  |
| `id` |  |
| `image_url` |  |
| `max_compare_at_price` |  |
| `max_current_price` |  |
| `min_compare_at_price` |  |
| `min_current_price` |  |
| `price_display` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: list, load.

API path: `/products`

#### Publication

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `section_id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: list, load.

API path: `/publications`

#### Search

| Field | Description |
| --- | --- |
| `api_id` |  |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `is_boosted` |  |
| `score` |  |
| `thumbnail` |  |
| `timestamp` |  |
| `title` |  |

Operations: list.

API path: `/agents/search`

#### Section

| Field | Description |
| --- | --- |
| `accession` |  |
| `api_link` |  |
| `api_model` |  |
| `artwork_id` |  |
| `content` |  |
| `generic_page_id` |  |
| `id` |  |
| `publication_id` |  |
| `publication_title` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: list, load.

API path: `/sections`

#### Site

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `artwork_id` |  |
| `artwork_title` |  |
| `description` |  |
| `exhibition_id` |  |
| `exhibition_title` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: list, load.

API path: `/sites`

#### Sound

| Field | Description |
| --- | --- |
| `alt_text` |  |
| `api_link` |  |
| `api_model` |  |
| `artwork_id` |  |
| `artwork_title` |  |
| `content` |  |
| `content_e_tag` |  |
| `credit_line` |  |
| `id` |  |
| `is_educational_resource` |  |
| `is_multimedia_resource` |  |
| `is_teacher_resource` |  |
| `lake_guid` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `transcript` |  |
| `type` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: list, load.

API path: `/mobile-sounds`

#### StaticPage

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: list, load.

API path: `/static-pages`

#### Text

| Field | Description |
| --- | --- |
| `alt_text` |  |
| `api_link` |  |
| `api_model` |  |
| `artwork_id` |  |
| `artwork_title` |  |
| `content` |  |
| `content_e_tag` |  |
| `credit_line` |  |
| `id` |  |
| `is_educational_resource` |  |
| `is_multimedia_resource` |  |
| `is_teacher_resource` |  |
| `lake_guid` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `type` |  |
| `updated_at` |  |

Operations: list, load.

API path: `/texts`

#### Tour

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `artist_title` |  |
| `artwork_title` |  |
| `description` |  |
| `id` |  |
| `image` |  |
| `intro` |  |
| `intro_link` |  |
| `intro_transcript` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `weight` |  |

Operations: list, load.

API path: `/tours`

#### Video

| Field | Description |
| --- | --- |
| `alt_text` |  |
| `api_link` |  |
| `api_model` |  |
| `artwork_id` |  |
| `artwork_title` |  |
| `content` |  |
| `content_e_tag` |  |
| `credit_line` |  |
| `id` |  |
| `is_educational_resource` |  |
| `is_multimedia_resource` |  |
| `is_teacher_resource` |  |
| `lake_guid` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `type` |  |
| `updated_at` |  |

Operations: list, load.

API path: `/videos`



## Entities


### Agent

Create an instance: `const agent = client.Agent()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_title` | `any` |  |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `birth_date` | `any` |  |
| `death_date` | `any` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `is_artist` | `boolean` |  |
| `sort_title` | `any` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `ulan_id` | `string` |  |
| `updated_at` | `any` |  |

#### Example: Load

```ts
const agent = await client.Agent().load({ id: 'agent_id' })
```

#### Example: List

```ts
const agents = await client.Agent().list()
```


### AgentRole

Create an instance: `const agent_role = client.AgentRole()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `id` | `string` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |

#### Example: Load

```ts
const agent_role = await client.AgentRole().load({ id: 'agent_role_id' })
```

#### Example: List

```ts
const agent_roles = await client.AgentRole().list()
```


### AgentType

Create an instance: `const agent_type = client.AgentType()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `id` | `string` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |

#### Example: Load

```ts
const agent_type = await client.AgentType().load({ id: 'agent_type_id' })
```

#### Example: List

```ts
const agent_types = await client.AgentType().list()
```


### Article

Create an instance: `const article = client.Article()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `copy` | `any` |  |
| `id` | `string` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |

#### Example: Load

```ts
const article = await client.Article().load({ id: 'article_id' })
```

#### Example: List

```ts
const articles = await client.Article().list()
```


### Artwork

Create an instance: `const artwork = client.Artwork()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_artist_id` | `string` |  |
| `alt_classification_id` | `string` |  |
| `alt_image_id` | `string` |  |
| `alt_material_id` | `string` |  |
| `alt_style_id` | `string` |  |
| `alt_subject_id` | `string` |  |
| `alt_technique_id` | `string` |  |
| `alt_title` | `any` |  |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `artist_display` | `any` |  |
| `artist_id` | `string` |  |
| `artist_title` | `any` |  |
| `artwork_type_id` | `string` |  |
| `artwork_type_title` | `any` |  |
| `boost_rank` | `any` |  |
| `catalog_based_search_keyword_title` | `any` |  |
| `catalogue_display` | `any` |  |
| `category_id` | `string` |  |
| `category_title` | `any` |  |
| `classification_id` | `string` |  |
| `classification_title` | `any` |  |
| `color` | `any` |  |
| `colorfulness` | `any` |  |
| `copyright_notice` | `any` |  |
| `credit_line` | `any` |  |
| `date_display` | `any` |  |
| `date_end` | `any` |  |
| `date_qualifier_id` | `string` |  |
| `date_qualifier_title` | `any` |  |
| `date_start` | `any` |  |
| `department_id` | `string` |  |
| `department_title` | `any` |  |
| `description` | `string` |  |
| `dimension` | `any` |  |
| `dimensions_detail` | `any` |  |
| `document_id` | `string` |  |
| `edition` | `any` |  |
| `exhibition_history` | `any` |  |
| `fiscal_year` | `any` |  |
| `fiscal_year_deaccession` | `any` |  |
| `gallery_id` | `string` |  |
| `gallery_title` | `any` |  |
| `has_advanced_imaging` | `boolean` |  |
| `has_educational_resource` | `boolean` |  |
| `has_multimedia_resource` | `boolean` |  |
| `has_not_been_viewed_much` | `boolean` |  |
| `id` | `string` |  |
| `image_embedding` | `any` |  |
| `image_id` | `string` |  |
| `inscription` | `any` |  |
| `internal_department_id` | `string` |  |
| `is_boosted` | `boolean` |  |
| `is_on_view` | `boolean` |  |
| `is_public_domain` | `boolean` |  |
| `is_zoomable` | `boolean` |  |
| `latitude` | `number` |  |
| `latlon` | `any` |  |
| `longitude` | `number` |  |
| `main_reference_number` | `number` |  |
| `material_id` | `string` |  |
| `material_title` | `any` |  |
| `max_zoom_window_size` | `any` |  |
| `medium_display` | `any` |  |
| `nomisma_id` | `string` |  |
| `on_loan_display` | `any` |  |
| `pageview` | `any` |  |
| `pageviews_recent` | `any` |  |
| `place_of_origin` | `any` |  |
| `provenance_text` | `any` |  |
| `publication_history` | `any` |  |
| `publishing_verification_level` | `any` |  |
| `section_id` | `string` |  |
| `section_title` | `any` |  |
| `short_description` | `any` |  |
| `site_id` | `string` |  |
| `sound_id` | `string` |  |
| `source_updated_at` | `any` |  |
| `style_id` | `string` |  |
| `style_title` | `any` |  |
| `subject_id` | `string` |  |
| `subject_title` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `technique_id` | `string` |  |
| `technique_title` | `any` |  |
| `term_title` | `any` |  |
| `text_embedding` | `any` |  |
| `text_id` | `string` |  |
| `theme_title` | `any` |  |
| `thumbnail` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |
| `video_id` | `string` |  |

#### Example: Load

```ts
const artwork = await client.Artwork().load({ id: 'artwork_id' })
```

#### Example: List

```ts
const artworks = await client.Artwork().list()
```


### ArtworkDateQualifier

Create an instance: `const artwork_date_qualifier = client.ArtworkDateQualifier()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `id` | `string` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |

#### Example: Load

```ts
const artwork_date_qualifier = await client.ArtworkDateQualifier().load({ id: 'artwork_date_qualifier_id' })
```

#### Example: List

```ts
const artwork_date_qualifiers = await client.ArtworkDateQualifier().list()
```


### ArtworkPlaceQualifier

Create an instance: `const artwork_place_qualifier = client.ArtworkPlaceQualifier()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `id` | `string` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |

#### Example: Load

```ts
const artwork_place_qualifier = await client.ArtworkPlaceQualifier().load({ id: 'artwork_place_qualifier_id' })
```

#### Example: List

```ts
const artwork_place_qualifiers = await client.ArtworkPlaceQualifier().list()
```


### ArtworkType

Create an instance: `const artwork_type = client.ArtworkType()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `aat_id` | `string` |  |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `id` | `string` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |

#### Example: Load

```ts
const artwork_type = await client.ArtworkType().load({ id: 'artwork_type_id' })
```

#### Example: List

```ts
const artwork_types = await client.ArtworkType().list()
```


### CategoryTerm

Create an instance: `const category_term = client.CategoryTerm()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `aat_id` | `string` |  |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `id` | `string` |  |
| `parent_id` | `string` |  |
| `source_updated_at` | `any` |  |
| `subtype` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |

#### Example: Load

```ts
const category_term = await client.CategoryTerm().load({ id: 'category_term_id' })
```

#### Example: List

```ts
const category_terms = await client.CategoryTerm().list()
```


### DigitalPublication

Create an instance: `const digital_publication = client.DigitalPublication()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `copy` | `any` |  |
| `id` | `string` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |
| `web_url` | `any` |  |

#### Example: Load

```ts
const digital_publication = await client.DigitalPublication().load({ id: 'digital_publication_id' })
```

#### Example: List

```ts
const digital_publications = await client.DigitalPublication().list()
```


### DigitalPublicationArticle

Create an instance: `const digital_publication_article = client.DigitalPublicationArticle()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `author_display` | `any` |  |
| `copy` | `any` |  |
| `digital_publication_id` | `string` |  |
| `id` | `string` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |
| `web_url` | `any` |  |

#### Example: Load

```ts
const digital_publication_article = await client.DigitalPublicationArticle().load({ id: 'digital_publication_article_id' })
```

#### Example: List

```ts
const digital_publication_articles = await client.DigitalPublicationArticle().list()
```


### EducatorResource

Create an instance: `const educator_resource = client.EducatorResource()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `copy` | `any` |  |
| `id` | `string` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |
| `web_url` | `any` |  |

#### Example: Load

```ts
const educator_resource = await client.EducatorResource().load({ id: 'educator_resource_id' })
```

#### Example: List

```ts
const educator_resources = await client.EducatorResource().list()
```


### Event

Create an instance: `const event = client.Event()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_audience_id` | `string` |  |
| `alt_event_type_id` | `string` |  |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `audience_id` | `string` |  |
| `buy_button_caption` | `any` |  |
| `buy_button_text` | `any` |  |
| `date_display` | `any` |  |
| `description` | `string` |  |
| `door_time` | `any` |  |
| `end_date` | `any` |  |
| `end_time` | `any` |  |
| `entrance` | `any` |  |
| `event_host_id` | `string` |  |
| `event_host_title` | `any` |  |
| `event_type_id` | `string` |  |
| `header_description` | `any` |  |
| `hero_caption` | `any` |  |
| `id` | `string` |  |
| `image_url` | `any` |  |
| `is_admission_required` | `boolean` |  |
| `is_after_hour` | `boolean` |  |
| `is_free` | `boolean` |  |
| `is_member_exclusive` | `boolean` |  |
| `is_private` | `boolean` |  |
| `is_registration_required` | `boolean` |  |
| `is_sales_button_hidden` | `boolean` |  |
| `is_sold_out` | `boolean` |  |
| `is_ticketed` | `boolean` |  |
| `is_virtual_event` | `boolean` |  |
| `join_url` | `any` |  |
| `layout_type` | `any` |  |
| `list_description` | `any` |  |
| `location` | `any` |  |
| `program_id` | `string` |  |
| `program_title` | `any` |  |
| `rsvp_link` | `any` |  |
| `search_tag` | `any` |  |
| `short_description` | `any` |  |
| `slug` | `string` |  |
| `source_updated_at` | `any` |  |
| `start_date` | `any` |  |
| `start_time` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `survey_url` | `any` |  |
| `ticketed_event_id` | `string` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `title_display` | `any` |  |
| `updated_at` | `any` |  |
| `virtual_event_passcode` | `any` |  |
| `virtual_event_url` | `any` |  |

#### Example: Load

```ts
const event = await client.Event().load({ id: 'event_id' })
```

#### Example: List

```ts
const events = await client.Event().list()
```


### EventOccurrence

Create an instance: `const event_occurrence = client.EventOccurrence()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `button_caption` | `any` |  |
| `button_text` | `any` |  |
| `button_url` | `any` |  |
| `description` | `string` |  |
| `end_at` | `any` |  |
| `event_id` | `string` |  |
| `id` | `string` |  |
| `image_url` | `any` |  |
| `is_private` | `boolean` |  |
| `is_sales_button_hidden` | `boolean` |  |
| `is_ticketed` | `boolean` |  |
| `location` | `any` |  |
| `off_sale_at` | `any` |  |
| `on_sale_at` | `any` |  |
| `short_description` | `any` |  |
| `source_updated_at` | `any` |  |
| `start_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `title_display` | `any` |  |
| `updated_at` | `any` |  |

#### Example: Load

```ts
const event_occurrence = await client.EventOccurrence().load({ id: 'event_occurrence_id' })
```

#### Example: List

```ts
const event_occurrences = await client.EventOccurrence().list()
```


### EventProgram

Create an instance: `const event_program = client.EventProgram()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `id` | `string` |  |
| `is_affiliate_group` | `boolean` |  |
| `is_event_host` | `boolean` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |

#### Example: Load

```ts
const event_program = await client.EventProgram().load({ id: 'event_program_id' })
```

#### Example: List

```ts
const event_programs = await client.EventProgram().list()
```


### Exhibition

Create an instance: `const exhibition = client.Exhibition()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `aic_end_at` | `any` |  |
| `aic_start_at` | `any` |  |
| `alt_image_id` | `string` |  |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `artist_id` | `string` |  |
| `artwork_id` | `string` |  |
| `artwork_title` | `any` |  |
| `document_id` | `string` |  |
| `gallery_id` | `string` |  |
| `gallery_title` | `any` |  |
| `id` | `string` |  |
| `image_id` | `string` |  |
| `image_url` | `any` |  |
| `is_featured` | `boolean` |  |
| `is_published` | `boolean` |  |
| `position` | `any` |  |
| `short_description` | `any` |  |
| `site_id` | `string` |  |
| `source_updated_at` | `any` |  |
| `status` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |
| `web_url` | `any` |  |

#### Example: Load

```ts
const exhibition = await client.Exhibition().load({ id: 'exhibition_id' })
```

#### Example: List

```ts
const exhibitions = await client.Exhibition().list()
```


### Gallery

Create an instance: `const gallery = client.Gallery()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `floor` | `any` |  |
| `id` | `string` |  |
| `is_closed` | `boolean` |  |
| `latitude` | `number` |  |
| `latlon` | `any` |  |
| `longitude` | `number` |  |
| `number` | `any` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `tgn_id` | `string` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |

#### Example: Load

```ts
const gallery = await client.Gallery().load({ id: 'gallery_id' })
```

#### Example: List

```ts
const gallerys = await client.Gallery().list()
```


### GenericPage

Create an instance: `const generic_page = client.GenericPage()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `copy` | `any` |  |
| `id` | `string` |  |
| `search_tag` | `any` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |
| `web_url` | `any` |  |

#### Example: Load

```ts
const generic_page = await client.GenericPage().load({ id: 'generic_page_id' })
```

#### Example: List

```ts
const generic_pages = await client.GenericPage().list()
```


### Highlight

Create an instance: `const highlight = client.Highlight()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `copy` | `any` |  |
| `id` | `string` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |

#### Example: Load

```ts
const highlight = await client.Highlight().load({ id: 'highlight_id' })
```

#### Example: List

```ts
const highlights = await client.Highlight().list()
```


### Hour

Create an instance: `const hour = client.Hour()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `additional_text` | `any` |  |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `friday_is_closed` | `any` |  |
| `friday_member_close` | `any` |  |
| `friday_member_open` | `any` |  |
| `friday_public_close` | `any` |  |
| `friday_public_open` | `any` |  |
| `id` | `string` |  |
| `monday_is_closed` | `any` |  |
| `monday_member_close` | `any` |  |
| `monday_member_open` | `any` |  |
| `monday_public_close` | `any` |  |
| `monday_public_open` | `any` |  |
| `saturday_is_closed` | `any` |  |
| `saturday_member_close` | `any` |  |
| `saturday_member_open` | `any` |  |
| `saturday_public_close` | `any` |  |
| `saturday_public_open` | `any` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `summary` | `any` |  |
| `sunday_is_closed` | `any` |  |
| `sunday_member_close` | `any` |  |
| `sunday_member_open` | `any` |  |
| `sunday_public_close` | `any` |  |
| `sunday_public_open` | `any` |  |
| `thursday_is_closed` | `any` |  |
| `thursday_member_close` | `any` |  |
| `thursday_member_open` | `any` |  |
| `thursday_public_close` | `any` |  |
| `thursday_public_open` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `tuesday_is_closed` | `any` |  |
| `tuesday_member_close` | `any` |  |
| `tuesday_member_open` | `any` |  |
| `tuesday_public_close` | `any` |  |
| `tuesday_public_open` | `any` |  |
| `updated_at` | `any` |  |
| `wednesday_is_closed` | `any` |  |
| `wednesday_member_close` | `any` |  |
| `wednesday_member_open` | `any` |  |
| `wednesday_public_close` | `any` |  |
| `wednesday_public_open` | `any` |  |

#### Example: Load

```ts
const hour = await client.Hour().load({ id: 'hour_id' })
```

#### Example: List

```ts
const hours = await client.Hour().list()
```


### Image

Create an instance: `const image = client.Image()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ahash` | `any` |  |
| `alt_text` | `any` |  |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `artwork_id` | `string` |  |
| `artwork_title` | `any` |  |
| `color` | `any` |  |
| `colorfulness` | `any` |  |
| `content` | `any` |  |
| `content_e_tag` | `any` |  |
| `credit_line` | `any` |  |
| `fingerprint` | `any` |  |
| `height` | `number` |  |
| `id` | `string` |  |
| `iiif_url` | `any` |  |
| `is_educational_resource` | `boolean` |  |
| `is_multimedia_resource` | `boolean` |  |
| `is_teacher_resource` | `boolean` |  |
| `lake_guid` | `any` |  |
| `lqip` | `any` |  |
| `phash` | `any` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `type` | `any` |  |
| `updated_at` | `any` |  |
| `width` | `number` |  |

#### Example: Load

```ts
const image = await client.Image().load({ id: 'image_id' })
```

#### Example: List

```ts
const images = await client.Image().list()
```


### LandingPage

Create an instance: `const landing_page = client.LandingPage()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `copy` | `any` |  |
| `id` | `string` |  |
| `search_tag` | `any` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |
| `web_url` | `any` |  |

#### Example: Load

```ts
const landing_page = await client.LandingPage().load({ id: 'landing_page_id' })
```

#### Example: List

```ts
const landing_pages = await client.LandingPage().list()
```


### Place

Create an instance: `const place = client.Place()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `id` | `string` |  |
| `latitude` | `number` |  |
| `longitude` | `number` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `tgn_id` | `string` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |

#### Example: Load

```ts
const place = await client.Place().load({ id: 'place_id' })
```

#### Example: List

```ts
const places = await client.Place().list()
```


### PressRelease

Create an instance: `const press_release = client.PressRelease()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `copy` | `any` |  |
| `id` | `string` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |
| `web_url` | `any` |  |

#### Example: Load

```ts
const press_release = await client.PressRelease().load({ id: 'press_release_id' })
```

#### Example: List

```ts
const press_releases = await client.PressRelease().list()
```


### PrintedPublication

Create an instance: `const printed_publication = client.PrintedPublication()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `copy` | `any` |  |
| `id` | `string` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |
| `web_url` | `any` |  |

#### Example: Load

```ts
const printed_publication = await client.PrintedPublication().load({ id: 'printed_publication_id' })
```

#### Example: List

```ts
const printed_publications = await client.PrintedPublication().list()
```


### Product

Create an instance: `const product = client.Product()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `artist_id` | `string` |  |
| `artwork_id` | `string` |  |
| `description` | `string` |  |
| `exhibition_id` | `string` |  |
| `external_sku` | `any` |  |
| `id` | `string` |  |
| `image_url` | `any` |  |
| `max_compare_at_price` | `any` |  |
| `max_current_price` | `any` |  |
| `min_compare_at_price` | `any` |  |
| `min_current_price` | `any` |  |
| `price_display` | `any` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |
| `web_url` | `any` |  |

#### Example: Load

```ts
const product = await client.Product().load({ id: 'product_id' })
```

#### Example: List

```ts
const products = await client.Product().list()
```


### Publication

Create an instance: `const publication = client.Publication()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `id` | `string` |  |
| `section_id` | `string` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |
| `web_url` | `any` |  |

#### Example: Load

```ts
const publication = await client.Publication().load({ id: 'publication_id' })
```

#### Example: List

```ts
const publications = await client.Publication().list()
```


### Search

Create an instance: `const search = client.Search()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_id` | `string` |  |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `id` | `string` |  |
| `is_boosted` | `boolean` |  |
| `score` | `number` |  |
| `thumbnail` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |

#### Example: List

```ts
const searchs = await client.Search().list()
```


### Section

Create an instance: `const section = client.Section()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accession` | `any` |  |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `artwork_id` | `string` |  |
| `content` | `any` |  |
| `generic_page_id` | `string` |  |
| `id` | `string` |  |
| `publication_id` | `string` |  |
| `publication_title` | `any` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |
| `web_url` | `any` |  |

#### Example: Load

```ts
const section = await client.Section().load({ id: 'section_id' })
```

#### Example: List

```ts
const sections = await client.Section().list()
```


### Site

Create an instance: `const site = client.Site()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `artwork_id` | `string` |  |
| `artwork_title` | `any` |  |
| `description` | `string` |  |
| `exhibition_id` | `string` |  |
| `exhibition_title` | `any` |  |
| `id` | `string` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |
| `web_url` | `any` |  |

#### Example: Load

```ts
const site = await client.Site().load({ id: 'site_id' })
```

#### Example: List

```ts
const sites = await client.Site().list()
```


### Sound

Create an instance: `const sound = client.Sound()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_text` | `any` |  |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `artwork_id` | `string` |  |
| `artwork_title` | `any` |  |
| `content` | `any` |  |
| `content_e_tag` | `any` |  |
| `credit_line` | `any` |  |
| `id` | `string` |  |
| `is_educational_resource` | `boolean` |  |
| `is_multimedia_resource` | `boolean` |  |
| `is_teacher_resource` | `boolean` |  |
| `lake_guid` | `any` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `transcript` | `any` |  |
| `type` | `any` |  |
| `updated_at` | `any` |  |
| `web_url` | `any` |  |

#### Example: Load

```ts
const sound = await client.Sound().load({ id: 'sound_id' })
```

#### Example: List

```ts
const sounds = await client.Sound().list()
```


### StaticPage

Create an instance: `const static_page = client.StaticPage()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `id` | `string` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |
| `web_url` | `any` |  |

#### Example: Load

```ts
const static_page = await client.StaticPage().load({ id: 'static_page_id' })
```

#### Example: List

```ts
const static_pages = await client.StaticPage().list()
```


### Text

Create an instance: `const text = client.Text()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_text` | `any` |  |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `artwork_id` | `string` |  |
| `artwork_title` | `any` |  |
| `content` | `any` |  |
| `content_e_tag` | `any` |  |
| `credit_line` | `any` |  |
| `id` | `string` |  |
| `is_educational_resource` | `boolean` |  |
| `is_multimedia_resource` | `boolean` |  |
| `is_teacher_resource` | `boolean` |  |
| `lake_guid` | `any` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `type` | `any` |  |
| `updated_at` | `any` |  |

#### Example: Load

```ts
const text = await client.Text().load({ id: 'text_id' })
```

#### Example: List

```ts
const texts = await client.Text().list()
```


### Tour

Create an instance: `const tour = client.Tour()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `artist_title` | `any` |  |
| `artwork_title` | `any` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `image` | `any` |  |
| `intro` | `any` |  |
| `intro_link` | `any` |  |
| `intro_transcript` | `any` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |
| `weight` | `number` |  |

#### Example: Load

```ts
const tour = await client.Tour().load({ id: 'tour_id' })
```

#### Example: List

```ts
const tours = await client.Tour().list()
```


### Video

Create an instance: `const video = client.Video()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_text` | `any` |  |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `artwork_id` | `string` |  |
| `artwork_title` | `any` |  |
| `content` | `any` |  |
| `content_e_tag` | `any` |  |
| `credit_line` | `any` |  |
| `id` | `string` |  |
| `is_educational_resource` | `boolean` |  |
| `is_multimedia_resource` | `boolean` |  |
| `is_teacher_resource` | `boolean` |  |
| `lake_guid` | `any` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `type` | `any` |  |
| `updated_at` | `any` |  |

#### Example: Load

```ts
const video = await client.Video().load({ id: 'video_id' })
```

#### Example: List

```ts
const videos = await client.Video().list()
```


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Module structure

```
art-institute-of-chicago/
├── src/
│   ├── ArtInstituteOfChicagoSDK.ts        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
├── test/                   # Test suites
└── dist/                   # Compiled output
```

Import the SDK from the package root:

```ts
import { ArtInstituteOfChicagoSDK } from '@voxgig-sdk/art-institute-of-chicago'
```

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const agent = client.Agent()
await agent.list()

// agent.data() now returns the agent data from the last `list`
// agent.match() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
