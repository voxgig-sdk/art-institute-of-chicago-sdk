# ArtInstituteOfChicago TypeScript SDK



The TypeScript SDK for the ArtInstituteOfChicago API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.Agent()` — each with a small set of operations (`list`, `load`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Also generated from this model: `go`, `go-cli`, `go-mcp`, `lua`, `php`, `py`, `rb` — see
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

`list()` resolves to an array of Agent ENTITIES — every operation
resolves to entities, not raw records. Iterate them directly, and call
`.data()` on one for the record it holds:

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
  const publications = await client.Publication().list()
  console.log(publications)
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

const publication = await client.Publication().list()
// publication is the entity, populated with mock response data
// — call publication.data() for the record itself
console.log(publication)
```

You can also use the instance method:

```ts
const client = new ArtInstituteOfChicagoSDK()
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.Publication()

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
| `alt_titles` | Alternate names for this agent |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `birth_date` | The year this agent was born |
| `death_date` | The year this agent died |
| `description` | A biographical description of the agent |
| `id` | Unique identifier of this resource. |
| `is_artist` | Whether the agent is an artist. |
| `sort_title` | Sortable name for this agent, typically with last name first. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `ulan_id` | Unique identifier of this agent in Getty's ULAN |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: list, load.

API path: `/agents`

#### AgentRole

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: list, load.

API path: `/agent-roles`

#### AgentType

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: list, load.

API path: `/agent-types`

#### Article

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `copy` | The text of the article |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: list, load.

API path: `/articles`

#### Artwork

| Field | Description |
| --- | --- |
| `alt_artist_ids` | Unique identifiers of the non-preferred artists/cultures associated with this work |
| `alt_classification_ids` | Unique identifiers of all other non-preferred classification terms for this work |
| `alt_image_ids` | Unique identifiers of all non-preferred images of this work. |
| `alt_material_ids` | Unique identifiers of all other non-preferred material terms for this work |
| `alt_style_ids` | Unique identifiers of all other non-preferred style terms for this work |
| `alt_subject_ids` | Unique identifiers of all other non-preferred subject terms for this work |
| `alt_technique_ids` | Unique identifiers of all other non-preferred technique terms for this work |
| `alt_titles` | Alternate names for this work |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `artist_display` | Readable description of the creator of this work. |
| `artist_id` | Unique identifier of the preferred artist/culture associated with this work |
| `artist_ids` | Unique identifier of all artist/cultures associated with this work |
| `artist_title` | Name of the preferred artist/culture associated with this work |
| `artist_titles` | Names of all artist/cultures associated with this work |
| `artwork_type_id` | Unique identifier of the kind of object or work |
| `artwork_type_title` | The kind of object or work (e.g. |
| `boost_rank` | Manual indication of what rank this artwork should take in search results. |
| `catalog_based_search_keyword_titles` | The keyword search values that would be catalog-based searches on this record |
| `catalogue_display` | Brief text listing all the catalogues raisonnés which include this work. |
| `category_ids` | Unique identifiers of the categories this work is a part of |
| `category_titles` | Names of the categories this artwork is a part of |
| `classification_id` | Unique identifier of the preferred classification term for this work |
| `classification_ids` | Unique identifiers of all classification terms for this work |
| `classification_title` | The name of the preferred classification term for this work |
| `classification_titles` | The names of all classification terms related to this artwork |
| `color` | Dominant color of this artwork in HSL |
| `colorfulness` | Unbounded positive float representing an abstract measure of colorfulness. |
| `copyright_notice` | Statement notifying how the work is protected by copyright. |
| `credit_line` | Brief statement indicating how the work came into the collection |
| `date_display` | Readable, free-text description of the period of time associated with the creation of this work. |
| `date_end` | The year of the period of time associated with the creation of this work |
| `date_qualifier_id` | Unique identifier of the qualifer to the dates provided for this record. |
| `date_qualifier_title` | Readable, text qualifer to the dates provided for this record. |
| `date_start` | The year of the period of time associated with the creation of this work |
| `department_id` | Unique identifier of the curatorial department that this work belongs to |
| `department_title` | Name of the curatorial department that this work belongs to |
| `description` | Longer explanation describing the work |
| `dimensions` | The size, shape, scale, and dimensions of the work. |
| `dimensions_detail` | The height, width, depth, and/or diameter of each section of the work in centimeters |
| `document_ids` | Unique identifiers of assets that serve as documentation for this artwork |
| `edition` | Edition number if the work is one of many |
| `exhibition_history` | List of all the places this work has been exhibited |
| `fiscal_year` | The fiscal year in which the work was acquired. |
| `fiscal_year_deaccession` | The fiscal year in which the work was deaccessioned. |
| `gallery_id` | Unique identifier of the location of this work in our museum |
| `gallery_title` | The location of this work in our museum |
| `has_advanced_imaging` | Whether this artwork is enhanced with 3D models, 360 image sequences, Mirador views, etc. |
| `has_educational_resources` | Whether this artwork has any documents tagged as educational |
| `has_multimedia_resources` | Whether this artwork has any associated microsites, digital publications, or documents tagged as multimedia |
| `has_not_been_viewed_much` | Whether the artwork hasn't been visited on our website very much |
| `id` | Unique identifier of this resource. |
| `image_embedding` | The generated embeddings describing the artwork image |
| `image_id` | Unique identifier of the preferred image to use to represent this work |
| `inscriptions` | A description of distinguishing or identifying physical markings that are on the work |
| `internal_department_id` | An internal department id we use for analytics. |
| `is_boosted` | Whether this document should be boosted in search |
| `is_on_view` | Whether the work is on display |
| `is_public_domain` | Whether the work is in the public domain, meaning it was created before copyrights existed or has left the copyright term |
| `is_zoomable` | Whether images of the work are allowed to be displayed in a zoomable interface. |
| `latitude` | Latitude coordinate of the location of this work in our galleries |
| `latlon` | Latitude and longitude coordinates of the location of this work in our galleries |
| `longitude` | Longitude coordinate of the location of this work in our galleries |
| `main_reference_number` | Unique identifier assigned to the artwork upon acquisition |
| `material_id` | Unique identifier of the preferred material term for this work |
| `material_ids` | Unique identifiers of all material terms for this work |
| `material_titles` | The names of all material terms related to this artwork |
| `max_zoom_window_size` | The maximum size of the window the image is allowed to be viewed in, in pixels. |
| `medium_display` | The substances or materials used in the creation of a work |
| `nomisma_id` | Unique identifier of this work in the nomisma coin database |
| `on_loan_display` | If an artwork is on loan, this contains details about the loan |
| `pageviews` | Approx. |
| `pageviews_recent` | Approx. |
| `place_of_origin` | The location where the creation, design, or production of the work took place, or the original location of the work |
| `provenance_text` | Ownership/collecting history of the work. |
| `publication_history` | Bibliographic list of all the places this work has been published |
| `publishing_verification_level` | Indicator of how much metadata on the work in published. |
| `section_ids` | Unique identifiers of the digital publication chapters this work in included in |
| `section_titles` | Names of the digital publication chapters this work is included in |
| `short_description` | Short explanation describing the work |
| `site_ids` | Unique identifiers of the microsites this work is a part of |
| `sound_ids` | Unique identifiers of the audio about this work |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `style_id` | Unique identifier of the preferred style term for this work |
| `style_ids` | Unique identifiers of all style terms for this work |
| `style_title` | The name of the preferred style term for this work |
| `style_titles` | The names of all style terms related to this artwork |
| `subject_id` | Unique identifier of the preferred subject term for this work |
| `subject_ids` | Unique identifiers of all subject terms for this work |
| `subject_titles` | The names of all subject terms related to this artwork |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `technique_id` | Unique identifier of the preferred technique term for this work |
| `technique_ids` | Unique identifiers of all technique terms for this work |
| `technique_titles` | The names of all technique terms related to this artwork |
| `term_titles` | The names of the taxonomy tags for this work |
| `text_embedding` | The generated embeddings of artwork text |
| `text_ids` | Unique identifiers of the texts about this work |
| `theme_titles` | The names of all thematic publish categories related to this artwork |
| `thumbnail` | Metadata about the image referenced by `image_id`. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `video_ids` | Unique identifiers of the videos about this work |

Operations: list, load.

API path: `/artworks`

#### ArtworkDateQualifier

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: list, load.

API path: `/artwork-date-qualifiers`

#### ArtworkPlaceQualifier

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: list, load.

API path: `/artwork-place-qualifiers`

#### ArtworkType

| Field | Description |
| --- | --- |
| `aat_id` | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: list, load.

API path: `/artwork-types`

#### CategoryTerm

| Field | Description |
| --- | --- |
| `aat_id` | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `id` | Unique identifier of this resource. |
| `parent_id` | Unique identifier of this category's parent |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `subtype` | Takes one of the following values: classification, material, technique, style, subject, department, theme |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: list, load.

API path: `/category-terms`

#### DigitalPublication

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `copy` | The text of the page |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | The URL to this page on our website |

Operations: list, load.

API path: `/digital-publications`

#### DigitalPublicationArticle

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `author_display` | A display-friendly text of the authors of this article |
| `copy` | The text of the article |
| `digital_publication_id` | Unique identifier of the digital publication this article belongs to |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | The URL to this article on our website |

Operations: list, load.

API path: `/digital-publication-articles`

#### EducatorResource

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `copy` | The text of the page |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | The URL to this page on our website |

Operations: list, load.

API path: `/educator-resources`

#### Event

| Field | Description |
| --- | --- |
| `alt_audience_ids` | Unique identifiers indicating the alternate audiences for this event |
| `alt_event_type_ids` | Unique identifiers indicating the alternate types of this event |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `audience_id` | Unique identifier indicating the preferred audience for this event |
| `buy_button_caption` | Additional text below the ticket/registration button |
| `buy_button_text` | The text used on the ticket/registration button |
| `date_display` | A readable display of the event dates |
| `description` | All copytext of the event |
| `door_time` | The time the doors open for this event |
| `end_date` | The date the event ends |
| `end_time` | The time the event ends |
| `entrance` | Which entrance to use for this event |
| `event_host_id` | Unique identifier of the host (cf. |
| `event_host_title` | Unique identifier of the host (cf. |
| `event_type_id` | Unique identifier indicating the preferred type of this event |
| `header_description` | Brief description of the event displayed below the title |
| `hero_caption` | Text displayed with the hero image on the event |
| `id` | Unique identifier of this resource. |
| `image_url` | The URL of an image representing this page |
| `is_admission_required` | Whether admission to the museum is required to attend this event |
| `is_after_hours` | Whether the event is to be held after the museum closes |
| `is_free` | Whether the event is free |
| `is_member_exclusive` | Whether the event is exclusive to members of the museum |
| `is_private` | Whether the event is private |
| `is_registration_required` | Whether registration is required to attend the event |
| `is_sales_button_hidden` | Whether the buy tickets button is hidden on the website event page |
| `is_sold_out` | Whether the event is sold out |
| `is_ticketed` | Whether a ticket is required to attend the event |
| `is_virtual_event` | Whether the event is being held virtually |
| `join_url` | URL to the membership signup page via this event |
| `layout_type` | Number indicating the type of layout this event page uses |
| `list_description` | One-sentence description of the event displayed in listings |
| `location` | Where the event takes place |
| `program_ids` | Unique identifiers indicating the programs this event is a part of |
| `program_titles` | Titles of the programs this event is a part of |
| `rsvp_link` | The URL to the sales site for this event |
| `search_tags` | Editor-specified list of tags to aid in internal search |
| `short_description` | Brief description of the event |
| `slug` | A string used in the URL for this event |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `start_date` | The date the event begins |
| `start_time` | The time the event starts |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `survey_url` | URL to the survey associated with this event |
| `ticketed_event_id` | Unique identifier of the event in the ticketing system this website event is tied to |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `title_display` | The name of this event formatted with HTML (optional) |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `virtual_event_passcode` | Passcode to access the virtual event |
| `virtual_event_url` | URL to the virtual event |

Operations: list, load.

API path: `/events`

#### EventOccurrence

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `button_caption` | Additional text below the ticket/registration button |
| `button_text` | The text used on the ticket/registration button |
| `button_url` | The URL to the sales site or an RSVP link for this event |
| `description` | Description of the event |
| `end_at` | The date the event occurrence ends |
| `event_id` | Identifier of the master event of which this is an occurrence |
| `id` | Unique identifier of this resource. |
| `image_url` | The URL of an image representing this page |
| `is_private` | Whether the event is private. |
| `is_sales_button_hidden` | Whether the buy tickets button is hidden on the website event page |
| `is_ticketed` | Whether a ticket is required to attend the event |
| `location` | Where the event takes place |
| `off_sale_at` | Date and time the event goes off sale |
| `on_sale_at` | Date and time the event goes on sale |
| `short_description` | Brief description of the event |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `start_at` | The date the event occurrence begins |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `title_display` | The name of this event formatted with HTML (optional) |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: list, load.

API path: `/event-occurrences`

#### EventProgram

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `id` | Unique identifier of this resource. |
| `is_affiliate_group` | Whether this program represents an affiliate group |
| `is_event_host` | Whether this program represents an event host |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: list, load.

API path: `/event-programs`

#### Exhibition

| Field | Description |
| --- | --- |
| `aic_end_at` | Date the exhibition closed at the Art Institute of Chicago |
| `aic_start_at` | Date the exhibition opened at the Art Institute of Chicago |
| `alt_image_ids` | Unique identifiers of all non-preferred images of this exhibition. |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `artist_ids` | Unique identifiers of the artist agent records representing who was shown in the exhibition |
| `artwork_ids` | Unique identifiers of the artworks that were part of the exhibition |
| `artwork_titles` | Names of the artworks that were part of the exhibition |
| `document_ids` | Unique identifiers of assets that serve as documentation for this exhibition |
| `gallery_id` | Unique identifier of the gallery that mainly housed the exhibition |
| `gallery_title` | The name of the gallery that mainly housed the exhibition |
| `id` | Unique identifier of this resource. |
| `image_id` | Unique identifier of the preferred image to use to represent this exhibition |
| `image_url` | URL to the hero image from the website |
| `is_featured` | Is this exhibition currently featured on our website? |
| `is_published` | Is this exhibition currently published on our website? |
| `position` | Numering position represnting the order in which this exhibition is featured on the website |
| `short_description` | Brief explanation of what this exhibition is |
| `site_ids` | Unique identifiers of the microsites this exhibition is a part of |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `status` | Whether the exhibition is open or closed |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | URL to this exhibition on our website |

Operations: list, load.

API path: `/exhibitions`

#### Gallery

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `floor` | The level the gallery is on, e.g., 1, 2, 3, or LL |
| `id` | Unique identifier of this resource. |
| `is_closed` | Whether the gallery is currently closed |
| `latitude` | Latitude coordinate of the center of the room |
| `latlon` | Latitude and longitude coordinates of the center of the room |
| `longitude` | Longitude coordinate of the center of the room |
| `number` | The gallery's room number. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: list, load.

API path: `/galleries`

#### GenericPage

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `copy` | The text of the page |
| `id` | Unique identifier of this resource. |
| `search_tags` | Editor-specified list of tags to aid in internal search |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | The URL to this page on our website |

Operations: list, load.

API path: `/generic-pages`

#### Highlight

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `copy` | The text of the highlight description |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: list, load.

API path: `/highlights`

#### Hour

| Field | Description |
| --- | --- |
| `additional_text` | Additional information about the hours |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `friday_is_closed` | Whether the museum is closed on Fridays |
| `friday_member_close` | The time member hours ends on Fridays |
| `friday_member_open` | The time member hours starts on Fridays |
| `friday_public_close` | The time public hours ends on Fridays |
| `friday_public_open` | The time public hours starts on Fridays |
| `id` | Unique identifier of this resource. |
| `monday_is_closed` | Whether the museum is closed on Mondays |
| `monday_member_close` | The time member hours ends on Mondays |
| `monday_member_open` | The time member hours starts on Mondays |
| `monday_public_close` | The time public hours ends on Mondays |
| `monday_public_open` | The time public hours starts on Mondays |
| `saturday_is_closed` | Whether the museum is closed on Saturdays |
| `saturday_member_close` | The time member hours ends on Saturdays |
| `saturday_member_open` | The time member hours starts on Saturdays |
| `saturday_public_close` | The time public hours ends on Saturdays |
| `saturday_public_open` | The time public hours starts on Saturdays |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `summary` | Readable summary of the hours |
| `sunday_is_closed` | Whether the museum is closed on Sundays |
| `sunday_member_close` | The time member hours ends on Sundays |
| `sunday_member_open` | The time member hours starts on Sundays |
| `sunday_public_close` | The time public hours ends on Sundays |
| `sunday_public_open` | The time public hours starts on Sundays |
| `thursday_is_closed` | Whether the museum is closed on Thursdays |
| `thursday_member_close` | The time member hours ends on Thursdays |
| `thursday_member_open` | The time member hours starts on Thursdays |
| `thursday_public_close` | The time public hours ends on Thursdays |
| `thursday_public_open` | The time public hours starts on Thursdays |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `tuesday_is_closed` | Whether the museum is closed on Tuesdays |
| `tuesday_member_close` | The time member hours ends on Tuesdays |
| `tuesday_member_open` | The time member hours starts on Tuesdays |
| `tuesday_public_close` | The time public hours ends on Tuesdays |
| `tuesday_public_open` | The time public hours starts on Tuesdays |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `wednesday_is_closed` | Whether the museum is closed on Wednesdays |
| `wednesday_member_close` | The time member hours ends on Wednesdays |
| `wednesday_member_open` | The time member hours starts on Wednesdays |
| `wednesday_public_close` | The time public hours ends on Wednesdays |
| `wednesday_public_open` | The time public hours starts on Wednesdays |

Operations: list, load.

API path: `/hours`

#### Image

| Field | Description |
| --- | --- |
| `ahash` | Image hash generated using ahash algorithm with 64 boolean subfields |
| `alt_text` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `artwork_ids` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | Names of the artworks associated with this asset |
| `color` | Dominant color of this image in HSL |
| `colorfulness` | Unbounded positive float representing an abstract measure of colorfulness. |
| `content` | Text of or URL to the contents of this asset |
| `content_e_tag` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | Asset-specific copyright information |
| `fingerprint` | Image hashes: aHash, dHash, pHash, wHash |
| `height` | Native height of the image |
| `id` | Unique identifier of this resource. |
| `iiif_url` | IIIF URL of this image |
| `is_educational_resource` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | Whether this resource is considered to be educational |
| `lake_guid` | Unique UUID of this resource in LAKE, our DAMS. |
| `lqip` | Low-quality image placeholder (LQIP). |
| `phash` | Image hash generated using phash algorithm with 64 boolean subfields |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `type` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `width` | Native width of the image |

Operations: list, load.

API path: `/images`

#### LandingPage

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `copy` | The text of the page |
| `id` | Unique identifier of this resource. |
| `search_tags` | Editor-specified list of tags to aid in internal search |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | The URL to this page on our website |

Operations: list, load.

API path: `/landing-pages`

#### Place

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `id` | Unique identifier of this resource. |
| `latitude` | Latitude coordinate of the center of the room |
| `longitude` | Longitude coordinate of the center of the room |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: list, load.

API path: `/places`

#### PressRelease

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `copy` | The text of the page |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | The URL to this page on our website |

Operations: list, load.

API path: `/press-releases`

#### PrintedPublication

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `copy` | The text of the page |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | The URL to this page on our website |

Operations: list, load.

API path: `/printed-publications`

#### Product

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `artist_ids` | Unique identifiers of the artists associated with this product |
| `artwork_ids` | Unique identifiers of the artworks associated with this product |
| `description` | Explanation of what this product is |
| `exhibition_ids` | Unique identifiers of the exhibitions associated with this product |
| `external_sku` | Numeric product identification code of a machine-readable barcode, when the customer sku differs from our internal one |
| `id` | Unique identifier of this resource. |
| `image_url` | URL of an image for this product |
| `max_compare_at_price` | Number indicating how much the most expensive variant of a product cost before a sale |
| `max_current_price` | Number indicating how much the most expensive variant of a product costs right now |
| `min_compare_at_price` | Number indicating how much the least expensive variant of a product cost before a sale |
| `min_current_price` | Number indicating how much the least expensive variant of a product costs right now |
| `price_display` | Explanation of what this product is |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | URL of this product in the shop |

Operations: list, load.

API path: `/products`

#### Publication

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `id` | Unique identifier of this resource. |
| `section_ids` | Unique identifiers of the sections of this publication |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | URL to the publication |

Operations: list, load.

API path: `/publications`

#### Search

| Field | Description |
| --- | --- |
| `api_id` | API unique identifier |
| `api_link` | URL to this recource in the API |
| `api_model` | Name of the model the resource represents |
| `id` | Unique identifier within the search index |
| `is_boosted` | Whether this record has been flagged to be boosted |
| `score` | Search index ranking of the result |
| `thumbnail` | Metadata on the image representing this record |
| `timestamp` | Date this record was last updated in the API |
| `title` | The name of this resource |

Operations: list.

API path: `/agents/search`

#### Section

| Field | Description |
| --- | --- |
| `accession` | An accession number parsed from the title or tombstone |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `artwork_id` | Unique identifier of the artwork with which this section is associated |
| `content` | Content of this section in plaintext |
| `generic_page_id` | Unique identifier of the page on the website that represents the publication this section belongs to |
| `id` | Unique identifier of this resource. |
| `publication_id` | Unique identifier of the publication this section belongs to |
| `publication_title` | Name of the publication this section belongs to |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | URL to the section |

Operations: list, load.

API path: `/sections`

#### Site

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `artwork_ids` | Unique identifiers of the artworks this site is associated with |
| `artwork_titles` | Names of the artworks this site is associated with |
| `description` | Explanation of what this site is |
| `exhibition_ids` | Unique identifier of the exhibitions this site is associated with |
| `exhibition_titles` | Names of the exhibitions this site is associated with |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | URL to this site |

Operations: list, load.

API path: `/sites`

#### Sound

| Field | Description |
| --- | --- |
| `alt_text` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `artwork_ids` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | Names of the artworks associated with this asset |
| `content` | Text of or URL to the contents of this asset |
| `content_e_tag` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | Asset-specific copyright information |
| `id` | Unique identifier of this resource. |
| `is_educational_resource` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | Whether this resource is considered to be educational |
| `lake_guid` | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | Name of this mobile audio file – derived from the artwork and tour titles |
| `transcript` | Text transcription of the audio file |
| `type` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | URL to the audio file |

Operations: list, load.

API path: `/mobile-sounds`

#### StaticPage

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `id` | Unique identifier of this resource. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `web_url` | The URL to this page on our website |

Operations: list, load.

API path: `/static-pages`

#### Text

| Field | Description |
| --- | --- |
| `alt_text` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `artwork_ids` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | Names of the artworks associated with this asset |
| `content` | Text of or URL to the contents of this asset |
| `content_e_tag` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | Asset-specific copyright information |
| `id` | Unique identifier of this resource. |
| `is_educational_resource` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | Whether this resource is considered to be educational |
| `lake_guid` | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `type` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | Date and time the record was updated in the aggregator database |

Operations: list, load.

API path: `/texts`

#### Tour

| Field | Description |
| --- | --- |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `artist_titles` | Names of the artists of the artworks featured in this tour's tour stops |
| `artwork_titles` | Names of the artworks featured in this tour's tour stops |
| `description` | Explanation of what the tour is |
| `id` | Unique identifier of this resource. |
| `image` | The main image for the tour |
| `intro` | Text introducing the tour |
| `intro_link` | Link to the audio file of the introduction |
| `intro_transcript` | Transcript of the introduction audio to the tour |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `updated_at` | Date and time the record was updated in the aggregator database |
| `weight` | Number representing this tour's sort order |

Operations: list, load.

API path: `/tours`

#### Video

| Field | Description |
| --- | --- |
| `alt_text` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | REST API link for this resource |
| `api_model` | REST API resource type or endpoint |
| `artwork_ids` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | Names of the artworks associated with this asset |
| `content` | Text of or URL to the contents of this asset |
| `content_e_tag` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | Asset-specific copyright information |
| `id` | Unique identifier of this resource. |
| `is_educational_resource` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | Whether this resource is considered to be educational |
| `lake_guid` | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | Date and time the record was updated in the aggregator search index |
| `title` | The name of this resource |
| `type` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | Date and time the record was updated in the aggregator database |

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
| `alt_titles` | `any` | Alternate names for this agent |
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `birth_date` | `any` | The year this agent was born |
| `death_date` | `any` | The year this agent died |
| `description` | `string` | A biographical description of the agent |
| `id` | `string` | Unique identifier of this resource. |
| `is_artist` | `boolean` | Whether the agent is an artist. |
| `sort_title` | `any` | Sortable name for this agent, typically with last name first. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `ulan_id` | `string` | Unique identifier of this agent in Getty's ULAN |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `copy` | `any` | The text of the article |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |

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
| `alt_artist_ids` | `any` | Unique identifiers of the non-preferred artists/cultures associated with this work |
| `alt_classification_ids` | `any` | Unique identifiers of all other non-preferred classification terms for this work |
| `alt_image_ids` | `any` | Unique identifiers of all non-preferred images of this work. |
| `alt_material_ids` | `any` | Unique identifiers of all other non-preferred material terms for this work |
| `alt_style_ids` | `any` | Unique identifiers of all other non-preferred style terms for this work |
| `alt_subject_ids` | `any` | Unique identifiers of all other non-preferred subject terms for this work |
| `alt_technique_ids` | `any` | Unique identifiers of all other non-preferred technique terms for this work |
| `alt_titles` | `any` | Alternate names for this work |
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `artist_display` | `any` | Readable description of the creator of this work. |
| `artist_id` | `string` | Unique identifier of the preferred artist/culture associated with this work |
| `artist_ids` | `any` | Unique identifier of all artist/cultures associated with this work |
| `artist_title` | `any` | Name of the preferred artist/culture associated with this work |
| `artist_titles` | `any` | Names of all artist/cultures associated with this work |
| `artwork_type_id` | `string` | Unique identifier of the kind of object or work |
| `artwork_type_title` | `any` | The kind of object or work (e.g. |
| `boost_rank` | `any` | Manual indication of what rank this artwork should take in search results. |
| `catalog_based_search_keyword_titles` | `any` | The keyword search values that would be catalog-based searches on this record |
| `catalogue_display` | `any` | Brief text listing all the catalogues raisonnés which include this work. |
| `category_ids` | `any` | Unique identifiers of the categories this work is a part of |
| `category_titles` | `any` | Names of the categories this artwork is a part of |
| `classification_id` | `string` | Unique identifier of the preferred classification term for this work |
| `classification_ids` | `any` | Unique identifiers of all classification terms for this work |
| `classification_title` | `any` | The name of the preferred classification term for this work |
| `classification_titles` | `any` | The names of all classification terms related to this artwork |
| `color` | `any` | Dominant color of this artwork in HSL |
| `colorfulness` | `any` | Unbounded positive float representing an abstract measure of colorfulness. |
| `copyright_notice` | `any` | Statement notifying how the work is protected by copyright. |
| `credit_line` | `any` | Brief statement indicating how the work came into the collection |
| `date_display` | `any` | Readable, free-text description of the period of time associated with the creation of this work. |
| `date_end` | `any` | The year of the period of time associated with the creation of this work |
| `date_qualifier_id` | `string` | Unique identifier of the qualifer to the dates provided for this record. |
| `date_qualifier_title` | `any` | Readable, text qualifer to the dates provided for this record. |
| `date_start` | `any` | The year of the period of time associated with the creation of this work |
| `department_id` | `string` | Unique identifier of the curatorial department that this work belongs to |
| `department_title` | `any` | Name of the curatorial department that this work belongs to |
| `description` | `string` | Longer explanation describing the work |
| `dimensions` | `any` | The size, shape, scale, and dimensions of the work. |
| `dimensions_detail` | `any` | The height, width, depth, and/or diameter of each section of the work in centimeters |
| `document_ids` | `any` | Unique identifiers of assets that serve as documentation for this artwork |
| `edition` | `any` | Edition number if the work is one of many |
| `exhibition_history` | `any` | List of all the places this work has been exhibited |
| `fiscal_year` | `any` | The fiscal year in which the work was acquired. |
| `fiscal_year_deaccession` | `any` | The fiscal year in which the work was deaccessioned. |
| `gallery_id` | `string` | Unique identifier of the location of this work in our museum |
| `gallery_title` | `any` | The location of this work in our museum |
| `has_advanced_imaging` | `boolean` | Whether this artwork is enhanced with 3D models, 360 image sequences, Mirador views, etc. |
| `has_educational_resources` | `boolean` | Whether this artwork has any documents tagged as educational |
| `has_multimedia_resources` | `boolean` | Whether this artwork has any associated microsites, digital publications, or documents tagged as multimedia |
| `has_not_been_viewed_much` | `boolean` | Whether the artwork hasn't been visited on our website very much |
| `id` | `string` | Unique identifier of this resource. |
| `image_embedding` | `any` | The generated embeddings describing the artwork image |
| `image_id` | `string` | Unique identifier of the preferred image to use to represent this work |
| `inscriptions` | `any` | A description of distinguishing or identifying physical markings that are on the work |
| `internal_department_id` | `string` | An internal department id we use for analytics. |
| `is_boosted` | `boolean` | Whether this document should be boosted in search |
| `is_on_view` | `boolean` | Whether the work is on display |
| `is_public_domain` | `boolean` | Whether the work is in the public domain, meaning it was created before copyrights existed or has left the copyright term |
| `is_zoomable` | `boolean` | Whether images of the work are allowed to be displayed in a zoomable interface. |
| `latitude` | `number` | Latitude coordinate of the location of this work in our galleries |
| `latlon` | `any` | Latitude and longitude coordinates of the location of this work in our galleries |
| `longitude` | `number` | Longitude coordinate of the location of this work in our galleries |
| `main_reference_number` | `number` | Unique identifier assigned to the artwork upon acquisition |
| `material_id` | `string` | Unique identifier of the preferred material term for this work |
| `material_ids` | `any` | Unique identifiers of all material terms for this work |
| `material_titles` | `any` | The names of all material terms related to this artwork |
| `max_zoom_window_size` | `any` | The maximum size of the window the image is allowed to be viewed in, in pixels. |
| `medium_display` | `any` | The substances or materials used in the creation of a work |
| `nomisma_id` | `string` | Unique identifier of this work in the nomisma coin database |
| `on_loan_display` | `any` | If an artwork is on loan, this contains details about the loan |
| `pageviews` | `any` | Approx. |
| `pageviews_recent` | `any` | Approx. |
| `place_of_origin` | `any` | The location where the creation, design, or production of the work took place, or the original location of the work |
| `provenance_text` | `any` | Ownership/collecting history of the work. |
| `publication_history` | `any` | Bibliographic list of all the places this work has been published |
| `publishing_verification_level` | `any` | Indicator of how much metadata on the work in published. |
| `section_ids` | `any` | Unique identifiers of the digital publication chapters this work in included in |
| `section_titles` | `any` | Names of the digital publication chapters this work is included in |
| `short_description` | `any` | Short explanation describing the work |
| `site_ids` | `any` | Unique identifiers of the microsites this work is a part of |
| `sound_ids` | `any` | Unique identifiers of the audio about this work |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `style_id` | `string` | Unique identifier of the preferred style term for this work |
| `style_ids` | `any` | Unique identifiers of all style terms for this work |
| `style_title` | `any` | The name of the preferred style term for this work |
| `style_titles` | `any` | The names of all style terms related to this artwork |
| `subject_id` | `string` | Unique identifier of the preferred subject term for this work |
| `subject_ids` | `any` | Unique identifiers of all subject terms for this work |
| `subject_titles` | `any` | The names of all subject terms related to this artwork |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `technique_id` | `string` | Unique identifier of the preferred technique term for this work |
| `technique_ids` | `any` | Unique identifiers of all technique terms for this work |
| `technique_titles` | `any` | The names of all technique terms related to this artwork |
| `term_titles` | `any` | The names of the taxonomy tags for this work |
| `text_embedding` | `any` | The generated embeddings of artwork text |
| `text_ids` | `any` | Unique identifiers of the texts about this work |
| `theme_titles` | `any` | The names of all thematic publish categories related to this artwork |
| `thumbnail` | `any` | Metadata about the image referenced by `image_id`. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |
| `video_ids` | `any` | Unique identifiers of the videos about this work |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |

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
| `aat_id` | `string` | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |

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
| `aat_id` | `string` | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `parent_id` | `string` | Unique identifier of this category's parent |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `subtype` | `any` | Takes one of the following values: classification, material, technique, style, subject, department, theme |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `copy` | `any` | The text of the page |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | The URL to this page on our website |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `author_display` | `any` | A display-friendly text of the authors of this article |
| `copy` | `any` | The text of the article |
| `digital_publication_id` | `string` | Unique identifier of the digital publication this article belongs to |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | The URL to this article on our website |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `copy` | `any` | The text of the page |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | The URL to this page on our website |

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
| `alt_audience_ids` | `any` | Unique identifiers indicating the alternate audiences for this event |
| `alt_event_type_ids` | `any` | Unique identifiers indicating the alternate types of this event |
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `audience_id` | `string` | Unique identifier indicating the preferred audience for this event |
| `buy_button_caption` | `any` | Additional text below the ticket/registration button |
| `buy_button_text` | `any` | The text used on the ticket/registration button |
| `date_display` | `any` | A readable display of the event dates |
| `description` | `string` | All copytext of the event |
| `door_time` | `any` | The time the doors open for this event |
| `end_date` | `any` | The date the event ends |
| `end_time` | `any` | The time the event ends |
| `entrance` | `any` | Which entrance to use for this event |
| `event_host_id` | `string` | Unique identifier of the host (cf. |
| `event_host_title` | `any` | Unique identifier of the host (cf. |
| `event_type_id` | `string` | Unique identifier indicating the preferred type of this event |
| `header_description` | `any` | Brief description of the event displayed below the title |
| `hero_caption` | `any` | Text displayed with the hero image on the event |
| `id` | `string` | Unique identifier of this resource. |
| `image_url` | `any` | The URL of an image representing this page |
| `is_admission_required` | `boolean` | Whether admission to the museum is required to attend this event |
| `is_after_hours` | `boolean` | Whether the event is to be held after the museum closes |
| `is_free` | `boolean` | Whether the event is free |
| `is_member_exclusive` | `boolean` | Whether the event is exclusive to members of the museum |
| `is_private` | `boolean` | Whether the event is private |
| `is_registration_required` | `boolean` | Whether registration is required to attend the event |
| `is_sales_button_hidden` | `boolean` | Whether the buy tickets button is hidden on the website event page |
| `is_sold_out` | `boolean` | Whether the event is sold out |
| `is_ticketed` | `boolean` | Whether a ticket is required to attend the event |
| `is_virtual_event` | `boolean` | Whether the event is being held virtually |
| `join_url` | `any` | URL to the membership signup page via this event |
| `layout_type` | `any` | Number indicating the type of layout this event page uses |
| `list_description` | `any` | One-sentence description of the event displayed in listings |
| `location` | `any` | Where the event takes place |
| `program_ids` | `any` | Unique identifiers indicating the programs this event is a part of |
| `program_titles` | `any` | Titles of the programs this event is a part of |
| `rsvp_link` | `any` | The URL to the sales site for this event |
| `search_tags` | `any` | Editor-specified list of tags to aid in internal search |
| `short_description` | `any` | Brief description of the event |
| `slug` | `string` | A string used in the URL for this event |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `start_date` | `any` | The date the event begins |
| `start_time` | `any` | The time the event starts |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `survey_url` | `any` | URL to the survey associated with this event |
| `ticketed_event_id` | `string` | Unique identifier of the event in the ticketing system this website event is tied to |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `title_display` | `any` | The name of this event formatted with HTML (optional) |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |
| `virtual_event_passcode` | `any` | Passcode to access the virtual event |
| `virtual_event_url` | `any` | URL to the virtual event |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `button_caption` | `any` | Additional text below the ticket/registration button |
| `button_text` | `any` | The text used on the ticket/registration button |
| `button_url` | `any` | The URL to the sales site or an RSVP link for this event |
| `description` | `string` | Description of the event |
| `end_at` | `any` | The date the event occurrence ends |
| `event_id` | `string` | Identifier of the master event of which this is an occurrence |
| `id` | `string` | Unique identifier of this resource. |
| `image_url` | `any` | The URL of an image representing this page |
| `is_private` | `boolean` | Whether the event is private. |
| `is_sales_button_hidden` | `boolean` | Whether the buy tickets button is hidden on the website event page |
| `is_ticketed` | `boolean` | Whether a ticket is required to attend the event |
| `location` | `any` | Where the event takes place |
| `off_sale_at` | `any` | Date and time the event goes off sale |
| `on_sale_at` | `any` | Date and time the event goes on sale |
| `short_description` | `any` | Brief description of the event |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `start_at` | `any` | The date the event occurrence begins |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `title_display` | `any` | The name of this event formatted with HTML (optional) |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `is_affiliate_group` | `boolean` | Whether this program represents an affiliate group |
| `is_event_host` | `boolean` | Whether this program represents an event host |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |

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
| `aic_end_at` | `any` | Date the exhibition closed at the Art Institute of Chicago |
| `aic_start_at` | `any` | Date the exhibition opened at the Art Institute of Chicago |
| `alt_image_ids` | `any` | Unique identifiers of all non-preferred images of this exhibition. |
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `artist_ids` | `any` | Unique identifiers of the artist agent records representing who was shown in the exhibition |
| `artwork_ids` | `any` | Unique identifiers of the artworks that were part of the exhibition |
| `artwork_titles` | `any` | Names of the artworks that were part of the exhibition |
| `document_ids` | `any` | Unique identifiers of assets that serve as documentation for this exhibition |
| `gallery_id` | `string` | Unique identifier of the gallery that mainly housed the exhibition |
| `gallery_title` | `any` | The name of the gallery that mainly housed the exhibition |
| `id` | `string` | Unique identifier of this resource. |
| `image_id` | `string` | Unique identifier of the preferred image to use to represent this exhibition |
| `image_url` | `any` | URL to the hero image from the website |
| `is_featured` | `boolean` | Is this exhibition currently featured on our website? |
| `is_published` | `boolean` | Is this exhibition currently published on our website? |
| `position` | `any` | Numering position represnting the order in which this exhibition is featured on the website |
| `short_description` | `any` | Brief explanation of what this exhibition is |
| `site_ids` | `any` | Unique identifiers of the microsites this exhibition is a part of |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `status` | `any` | Whether the exhibition is open or closed |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | URL to this exhibition on our website |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `floor` | `any` | The level the gallery is on, e.g., 1, 2, 3, or LL |
| `id` | `string` | Unique identifier of this resource. |
| `is_closed` | `boolean` | Whether the gallery is currently closed |
| `latitude` | `number` | Latitude coordinate of the center of the room |
| `latlon` | `any` | Latitude and longitude coordinates of the center of the room |
| `longitude` | `number` | Longitude coordinate of the center of the room |
| `number` | `any` | The gallery's room number. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | `string` | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `copy` | `any` | The text of the page |
| `id` | `string` | Unique identifier of this resource. |
| `search_tags` | `any` | Editor-specified list of tags to aid in internal search |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | The URL to this page on our website |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `copy` | `any` | The text of the highlight description |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |

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
| `additional_text` | `any` | Additional information about the hours |
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `friday_is_closed` | `any` | Whether the museum is closed on Fridays |
| `friday_member_close` | `any` | The time member hours ends on Fridays |
| `friday_member_open` | `any` | The time member hours starts on Fridays |
| `friday_public_close` | `any` | The time public hours ends on Fridays |
| `friday_public_open` | `any` | The time public hours starts on Fridays |
| `id` | `string` | Unique identifier of this resource. |
| `monday_is_closed` | `any` | Whether the museum is closed on Mondays |
| `monday_member_close` | `any` | The time member hours ends on Mondays |
| `monday_member_open` | `any` | The time member hours starts on Mondays |
| `monday_public_close` | `any` | The time public hours ends on Mondays |
| `monday_public_open` | `any` | The time public hours starts on Mondays |
| `saturday_is_closed` | `any` | Whether the museum is closed on Saturdays |
| `saturday_member_close` | `any` | The time member hours ends on Saturdays |
| `saturday_member_open` | `any` | The time member hours starts on Saturdays |
| `saturday_public_close` | `any` | The time public hours ends on Saturdays |
| `saturday_public_open` | `any` | The time public hours starts on Saturdays |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `summary` | `any` | Readable summary of the hours |
| `sunday_is_closed` | `any` | Whether the museum is closed on Sundays |
| `sunday_member_close` | `any` | The time member hours ends on Sundays |
| `sunday_member_open` | `any` | The time member hours starts on Sundays |
| `sunday_public_close` | `any` | The time public hours ends on Sundays |
| `sunday_public_open` | `any` | The time public hours starts on Sundays |
| `thursday_is_closed` | `any` | Whether the museum is closed on Thursdays |
| `thursday_member_close` | `any` | The time member hours ends on Thursdays |
| `thursday_member_open` | `any` | The time member hours starts on Thursdays |
| `thursday_public_close` | `any` | The time public hours ends on Thursdays |
| `thursday_public_open` | `any` | The time public hours starts on Thursdays |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `tuesday_is_closed` | `any` | Whether the museum is closed on Tuesdays |
| `tuesday_member_close` | `any` | The time member hours ends on Tuesdays |
| `tuesday_member_open` | `any` | The time member hours starts on Tuesdays |
| `tuesday_public_close` | `any` | The time public hours ends on Tuesdays |
| `tuesday_public_open` | `any` | The time public hours starts on Tuesdays |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |
| `wednesday_is_closed` | `any` | Whether the museum is closed on Wednesdays |
| `wednesday_member_close` | `any` | The time member hours ends on Wednesdays |
| `wednesday_member_open` | `any` | The time member hours starts on Wednesdays |
| `wednesday_public_close` | `any` | The time public hours ends on Wednesdays |
| `wednesday_public_open` | `any` | The time public hours starts on Wednesdays |

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
| `ahash` | `any` | Image hash generated using ahash algorithm with 64 boolean subfields |
| `alt_text` | `any` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `artwork_ids` | `any` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `any` | Names of the artworks associated with this asset |
| `color` | `any` | Dominant color of this image in HSL |
| `colorfulness` | `any` | Unbounded positive float representing an abstract measure of colorfulness. |
| `content` | `any` | Text of or URL to the contents of this asset |
| `content_e_tag` | `any` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `any` | Asset-specific copyright information |
| `fingerprint` | `any` | Image hashes: aHash, dHash, pHash, wHash |
| `height` | `number` | Native height of the image |
| `id` | `string` | Unique identifier of this resource. |
| `iiif_url` | `any` | IIIF URL of this image |
| `is_educational_resource` | `boolean` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `boolean` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `boolean` | Whether this resource is considered to be educational |
| `lake_guid` | `any` | Unique UUID of this resource in LAKE, our DAMS. |
| `lqip` | `any` | Low-quality image placeholder (LQIP). |
| `phash` | `any` | Image hash generated using phash algorithm with 64 boolean subfields |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `type` | `any` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |
| `width` | `number` | Native width of the image |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `copy` | `any` | The text of the page |
| `id` | `string` | Unique identifier of this resource. |
| `search_tags` | `any` | Editor-specified list of tags to aid in internal search |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | The URL to this page on our website |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `latitude` | `number` | Latitude coordinate of the center of the room |
| `longitude` | `number` | Longitude coordinate of the center of the room |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | `string` | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `copy` | `any` | The text of the page |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | The URL to this page on our website |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `copy` | `any` | The text of the page |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | The URL to this page on our website |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `artist_ids` | `any` | Unique identifiers of the artists associated with this product |
| `artwork_ids` | `any` | Unique identifiers of the artworks associated with this product |
| `description` | `string` | Explanation of what this product is |
| `exhibition_ids` | `any` | Unique identifiers of the exhibitions associated with this product |
| `external_sku` | `any` | Numeric product identification code of a machine-readable barcode, when the customer sku differs from our internal one |
| `id` | `string` | Unique identifier of this resource. |
| `image_url` | `any` | URL of an image for this product |
| `max_compare_at_price` | `any` | Number indicating how much the most expensive variant of a product cost before a sale |
| `max_current_price` | `any` | Number indicating how much the most expensive variant of a product costs right now |
| `min_compare_at_price` | `any` | Number indicating how much the least expensive variant of a product cost before a sale |
| `min_current_price` | `any` | Number indicating how much the least expensive variant of a product costs right now |
| `price_display` | `any` | Explanation of what this product is |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | URL of this product in the shop |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `section_ids` | `any` | Unique identifiers of the sections of this publication |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | URL to the publication |

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
| `api_id` | `string` | API unique identifier |
| `api_link` | `any` | URL to this recource in the API |
| `api_model` | `any` | Name of the model the resource represents |
| `id` | `string` | Unique identifier within the search index |
| `is_boosted` | `boolean` | Whether this record has been flagged to be boosted |
| `score` | `number` | Search index ranking of the result |
| `thumbnail` | `any` | Metadata on the image representing this record |
| `timestamp` | `any` | Date this record was last updated in the API |
| `title` | `string` | The name of this resource |

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
| `accession` | `any` | An accession number parsed from the title or tombstone |
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `artwork_id` | `string` | Unique identifier of the artwork with which this section is associated |
| `content` | `any` | Content of this section in plaintext |
| `generic_page_id` | `string` | Unique identifier of the page on the website that represents the publication this section belongs to |
| `id` | `string` | Unique identifier of this resource. |
| `publication_id` | `string` | Unique identifier of the publication this section belongs to |
| `publication_title` | `any` | Name of the publication this section belongs to |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | URL to the section |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `artwork_ids` | `any` | Unique identifiers of the artworks this site is associated with |
| `artwork_titles` | `any` | Names of the artworks this site is associated with |
| `description` | `string` | Explanation of what this site is |
| `exhibition_ids` | `any` | Unique identifier of the exhibitions this site is associated with |
| `exhibition_titles` | `any` | Names of the exhibitions this site is associated with |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | URL to this site |

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
| `alt_text` | `any` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `artwork_ids` | `any` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `any` | Names of the artworks associated with this asset |
| `content` | `any` | Text of or URL to the contents of this asset |
| `content_e_tag` | `any` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `any` | Asset-specific copyright information |
| `id` | `string` | Unique identifier of this resource. |
| `is_educational_resource` | `boolean` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `boolean` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `boolean` | Whether this resource is considered to be educational |
| `lake_guid` | `any` | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | Name of this mobile audio file – derived from the artwork and tour titles |
| `transcript` | `any` | Text transcription of the audio file |
| `type` | `any` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | URL to the audio file |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |
| `web_url` | `any` | The URL to this page on our website |

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
| `alt_text` | `any` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `artwork_ids` | `any` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `any` | Names of the artworks associated with this asset |
| `content` | `any` | Text of or URL to the contents of this asset |
| `content_e_tag` | `any` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `any` | Asset-specific copyright information |
| `id` | `string` | Unique identifier of this resource. |
| `is_educational_resource` | `boolean` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `boolean` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `boolean` | Whether this resource is considered to be educational |
| `lake_guid` | `any` | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `type` | `any` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |

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
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `artist_titles` | `any` | Names of the artists of the artworks featured in this tour's tour stops |
| `artwork_titles` | `any` | Names of the artworks featured in this tour's tour stops |
| `description` | `string` | Explanation of what the tour is |
| `id` | `string` | Unique identifier of this resource. |
| `image` | `any` | The main image for the tour |
| `intro` | `any` | Text introducing the tour |
| `intro_link` | `any` | Link to the audio file of the introduction |
| `intro_transcript` | `any` | Transcript of the introduction audio to the tour |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |
| `weight` | `number` | Number representing this tour's sort order |

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
| `alt_text` | `any` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `artwork_ids` | `any` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `any` | Names of the artworks associated with this asset |
| `content` | `any` | Text of or URL to the contents of this asset |
| `content_e_tag` | `any` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `any` | Asset-specific copyright information |
| `id` | `string` | Unique identifier of this resource. |
| `is_educational_resource` | `boolean` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `boolean` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `boolean` | Whether this resource is considered to be educational |
| `lake_guid` | `any` | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `type` | `any` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |

#### Example: Load

```ts
const video = await client.Video().load({ id: 'video_id' })
```

#### Example: List

```ts
const videos = await client.Video().list()
```

## Features

This SDK ships 1 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`test`](#test) | In-memory mock transport for testing without a live server |

### test

In-memory mock transport for testing without a live server.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.


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
const publication = client.Publication()
await publication.list()

// publication.data() now returns the publication data from the last `list`
// publication.match() returns the last match criteria
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
