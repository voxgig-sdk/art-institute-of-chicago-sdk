# ArtInstituteOfChicago Ruby SDK



The Ruby SDK for the ArtInstituteOfChicago API — an entity-oriented client using idiomatic Ruby conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.Agent` — with named operations (`list`/`load`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to RubyGems. Install it from the
GitHub release tag (`rb/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/art-institute-of-chicago-sdk/releases](https://github.com/voxgig-sdk/art-institute-of-chicago-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ruby
require_relative "ArtInstituteOfChicago_sdk"

client = ArtInstituteOfChicagoSDK.new
```

### 2. List agent records

```ruby
begin
  # list returns an Array of Agent records — iterate directly.
  agents = client.Agent.list
  agents.each do |item|
    puts "#{item["id"]} #{item["alt_titles"]}"
  end
rescue => err
  warn "list failed: #{err}"
end
```

### 3. Load an agent

```ruby
begin
  # load returns the ENTITY — call data_get for the Agent record (raises on error).
  agent = client.Agent.load({ "id" => "example_id" })
  puts agent
rescue => err
  warn "load failed: #{err}"
end
```


## Error handling

Entity operations raise on failure, so rescue them:

```ruby
begin
  publications = client.Publication.list()
rescue => err
  warn "list failed: #{err}"
end
```

`direct` does **not** raise — it returns the result hash. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example_id" },
})

warn "request failed: #{result["err"] || "HTTP #{result["status"]}"}" unless result["ok"]
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})

if result["ok"]
  puts result["status"]  # 200
  puts result["data"]    # response body
else
  # On an HTTP error status there is no err (only a transport failure sets
  # it), so fall back to the status code.
  warn(result["err"] || "HTTP #{result["status"]}")
end
```

### Prepare a request without sending it

```ruby
begin
  fetchdef = client.prepare({
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => { "id" => "example" },
  })
  puts fetchdef["url"]
  puts fetchdef["method"]
  puts fetchdef["headers"]
rescue => err
  warn "prepare failed: #{err}"
end
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```ruby
client = ArtInstituteOfChicagoSDK.test({
  "entity" => { "publication" => { "test01" => { "id" => "test01" } } },
})

# Entity ops return the ENTITY (raises on error);
# call data_get for the mock record.
publication = client.Publication.list()
puts publication
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```ruby
mock_fetch = ->(url, init) {
  return {
    "status" => 200,
    "statusText" => "OK",
    "headers" => {},
    "json" => ->() { { "id" => "mock01" } },
  }, nil
}

client = ArtInstituteOfChicagoSDK.new({
  "base" => "http://localhost:8080",
  "system" => {
    "fetch" => mock_fetch,
  },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
ART_INSTITUTE_OF_CHICAGO_TEST_LIVE=TRUE
```

Then run:

```bash
cd rb && ruby -Itest -e "Dir['test/*_test.rb'].each { |f| require_relative f }"
```


## Reference

### ArtInstituteOfChicagoSDK

```ruby
require_relative "ArtInstituteOfChicago_sdk"
client = ArtInstituteOfChicagoSDK.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `String` | Base URL of the API server. |
| `prefix` | `String` | URL path prefix prepended to all requests. |
| `suffix` | `String` | URL path suffix appended to all requests. |
| `feature` | `Hash` | Feature activation flags. |
| `extend` | `Hash` | Additional Feature instances to load. |
| `system` | `Hash` | System overrides (e.g. custom `fetch` lambda). |

### test

```ruby
client = ArtInstituteOfChicagoSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### ArtInstituteOfChicagoSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> Hash` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> Hash` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> Hash` | Build and send an HTTP request. Returns a result hash (`result["ok"]`); does not raise. |
| `Agent` | `(data) -> AgentEntity` | Create an Agent entity instance. |
| `AgentRole` | `(data) -> AgentRoleEntity` | Create an AgentRole entity instance. |
| `AgentType` | `(data) -> AgentTypeEntity` | Create an AgentType entity instance. |
| `Article` | `(data) -> ArticleEntity` | Create an Article entity instance. |
| `Artwork` | `(data) -> ArtworkEntity` | Create an Artwork entity instance. |
| `ArtworkDateQualifier` | `(data) -> ArtworkDateQualifierEntity` | Create an ArtworkDateQualifier entity instance. |
| `ArtworkPlaceQualifier` | `(data) -> ArtworkPlaceQualifierEntity` | Create an ArtworkPlaceQualifier entity instance. |
| `ArtworkType` | `(data) -> ArtworkTypeEntity` | Create an ArtworkType entity instance. |
| `CategoryTerm` | `(data) -> CategoryTermEntity` | Create a CategoryTerm entity instance. |
| `DigitalPublication` | `(data) -> DigitalPublicationEntity` | Create a DigitalPublication entity instance. |
| `DigitalPublicationArticle` | `(data) -> DigitalPublicationArticleEntity` | Create a DigitalPublicationArticle entity instance. |
| `EducatorResource` | `(data) -> EducatorResourceEntity` | Create an EducatorResource entity instance. |
| `Event` | `(data) -> EventEntity` | Create an Event entity instance. |
| `EventOccurrence` | `(data) -> EventOccurrenceEntity` | Create an EventOccurrence entity instance. |
| `EventProgram` | `(data) -> EventProgramEntity` | Create an EventProgram entity instance. |
| `Exhibition` | `(data) -> ExhibitionEntity` | Create an Exhibition entity instance. |
| `Gallery` | `(data) -> GalleryEntity` | Create a Gallery entity instance. |
| `GenericPage` | `(data) -> GenericPageEntity` | Create a GenericPage entity instance. |
| `Highlight` | `(data) -> HighlightEntity` | Create a Highlight entity instance. |
| `Hour` | `(data) -> HourEntity` | Create a Hour entity instance. |
| `Image` | `(data) -> ImageEntity` | Create an Image entity instance. |
| `LandingPage` | `(data) -> LandingPageEntity` | Create a LandingPage entity instance. |
| `Place` | `(data) -> PlaceEntity` | Create a Place entity instance. |
| `PressRelease` | `(data) -> PressReleaseEntity` | Create a PressRelease entity instance. |
| `PrintedPublication` | `(data) -> PrintedPublicationEntity` | Create a PrintedPublication entity instance. |
| `Product` | `(data) -> ProductEntity` | Create a Product entity instance. |
| `Publication` | `(data) -> PublicationEntity` | Create a Publication entity instance. |
| `Search` | `(data) -> SearchEntity` | Create a Search entity instance. |
| `Section` | `(data) -> SectionEntity` | Create a Section entity instance. |
| `Site` | `(data) -> SiteEntity` | Create a Site entity instance. |
| `Sound` | `(data) -> SoundEntity` | Create a Sound entity instance. |
| `StaticPage` | `(data) -> StaticPageEntity` | Create a StaticPage entity instance. |
| `Text` | `(data) -> TextEntity` | Create a Text entity instance. |
| `Tour` | `(data) -> TourEntity` | Create a Tour entity instance. |
| `Video` | `(data) -> VideoEntity` | Create a Video entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `list` | `(reqmatch = nil, ctrl) -> Array` | List entities matching the criteria (call with no argument to list all). Raises on error. |
| `data_get` | `() -> Hash` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> Hash` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> String` | Return the entity name. |

### Result shape

Entity operations return the result data directly. On failure they
raise a `ArtInstituteOfChicagoError` (a `StandardError` subclass), so wrap
calls in `begin`/`rescue` where you need to handle errors.

The `direct` escape hatch is the exception: it never raises and instead
returns a result `Hash` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `Boolean` | `true` if the HTTP status is 2xx. |
| `status` | `Integer` | HTTP status code. |
| `headers` | `Hash` | Response headers. |
| `data` | `any` | Parsed JSON response body. |
| `err` | `Error` | Present when `ok` is `false`. |

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

API path: `/videos`



## Entities


### Agent

Create an instance: `agent = client.Agent`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_titles` | `Object` | Alternate names for this agent |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `birth_date` | `Object` | The year this agent was born |
| `death_date` | `Object` | The year this agent died |
| `description` | `String` | A biographical description of the agent |
| `id` | `String` | Unique identifier of this resource. |
| `is_artist` | `Boolean` | Whether the agent is an artist. |
| `sort_title` | `Object` | Sortable name for this agent, typically with last name first. |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `ulan_id` | `String` | Unique identifier of this agent in Getty's ULAN |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Agent record (raises on error).
agent = client.Agent.load({ "id" => "agent_id" })
```

#### Example: List

```ruby
# list returns an Array of Agent records (raises on error).
agents = client.Agent.list
```


### AgentRole

Create an instance: `agent_role = client.AgentRole`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `id` | `String` | Unique identifier of this resource. |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the AgentRole record (raises on error).
agent_role = client.AgentRole.load({ "id" => "agent_role_id" })
```

#### Example: List

```ruby
# list returns an Array of AgentRole records (raises on error).
agent_roles = client.AgentRole.list
```


### AgentType

Create an instance: `agent_type = client.AgentType`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `id` | `String` | Unique identifier of this resource. |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the AgentType record (raises on error).
agent_type = client.AgentType.load({ "id" => "agent_type_id" })
```

#### Example: List

```ruby
# list returns an Array of AgentType records (raises on error).
agent_types = client.AgentType.list
```


### Article

Create an instance: `article = client.Article`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `copy` | `Object` | The text of the article |
| `id` | `String` | Unique identifier of this resource. |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Article record (raises on error).
article = client.Article.load({ "id" => "article_id" })
```

#### Example: List

```ruby
# list returns an Array of Article records (raises on error).
articles = client.Article.list
```


### Artwork

Create an instance: `artwork = client.Artwork`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_artist_ids` | `Object` | Unique identifiers of the non-preferred artists/cultures associated with this work |
| `alt_classification_ids` | `Object` | Unique identifiers of all other non-preferred classification terms for this work |
| `alt_image_ids` | `Object` | Unique identifiers of all non-preferred images of this work. |
| `alt_material_ids` | `Object` | Unique identifiers of all other non-preferred material terms for this work |
| `alt_style_ids` | `Object` | Unique identifiers of all other non-preferred style terms for this work |
| `alt_subject_ids` | `Object` | Unique identifiers of all other non-preferred subject terms for this work |
| `alt_technique_ids` | `Object` | Unique identifiers of all other non-preferred technique terms for this work |
| `alt_titles` | `Object` | Alternate names for this work |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `artist_display` | `Object` | Readable description of the creator of this work. |
| `artist_id` | `String` | Unique identifier of the preferred artist/culture associated with this work |
| `artist_ids` | `Object` | Unique identifier of all artist/cultures associated with this work |
| `artist_title` | `Object` | Name of the preferred artist/culture associated with this work |
| `artist_titles` | `Object` | Names of all artist/cultures associated with this work |
| `artwork_type_id` | `String` | Unique identifier of the kind of object or work |
| `artwork_type_title` | `Object` | The kind of object or work (e.g. |
| `boost_rank` | `Object` | Manual indication of what rank this artwork should take in search results. |
| `catalog_based_search_keyword_titles` | `Object` | The keyword search values that would be catalog-based searches on this record |
| `catalogue_display` | `Object` | Brief text listing all the catalogues raisonnés which include this work. |
| `category_ids` | `Object` | Unique identifiers of the categories this work is a part of |
| `category_titles` | `Object` | Names of the categories this artwork is a part of |
| `classification_id` | `String` | Unique identifier of the preferred classification term for this work |
| `classification_ids` | `Object` | Unique identifiers of all classification terms for this work |
| `classification_title` | `Object` | The name of the preferred classification term for this work |
| `classification_titles` | `Object` | The names of all classification terms related to this artwork |
| `color` | `Object` | Dominant color of this artwork in HSL |
| `colorfulness` | `Object` | Unbounded positive float representing an abstract measure of colorfulness. |
| `copyright_notice` | `Object` | Statement notifying how the work is protected by copyright. |
| `credit_line` | `Object` | Brief statement indicating how the work came into the collection |
| `date_display` | `Object` | Readable, free-text description of the period of time associated with the creation of this work. |
| `date_end` | `Object` | The year of the period of time associated with the creation of this work |
| `date_qualifier_id` | `String` | Unique identifier of the qualifer to the dates provided for this record. |
| `date_qualifier_title` | `Object` | Readable, text qualifer to the dates provided for this record. |
| `date_start` | `Object` | The year of the period of time associated with the creation of this work |
| `department_id` | `String` | Unique identifier of the curatorial department that this work belongs to |
| `department_title` | `Object` | Name of the curatorial department that this work belongs to |
| `description` | `String` | Longer explanation describing the work |
| `dimensions` | `Object` | The size, shape, scale, and dimensions of the work. |
| `dimensions_detail` | `Object` | The height, width, depth, and/or diameter of each section of the work in centimeters |
| `document_ids` | `Object` | Unique identifiers of assets that serve as documentation for this artwork |
| `edition` | `Object` | Edition number if the work is one of many |
| `exhibition_history` | `Object` | List of all the places this work has been exhibited |
| `fiscal_year` | `Object` | The fiscal year in which the work was acquired. |
| `fiscal_year_deaccession` | `Object` | The fiscal year in which the work was deaccessioned. |
| `gallery_id` | `String` | Unique identifier of the location of this work in our museum |
| `gallery_title` | `Object` | The location of this work in our museum |
| `has_advanced_imaging` | `Boolean` | Whether this artwork is enhanced with 3D models, 360 image sequences, Mirador views, etc. |
| `has_educational_resources` | `Boolean` | Whether this artwork has any documents tagged as educational |
| `has_multimedia_resources` | `Boolean` | Whether this artwork has any associated microsites, digital publications, or documents tagged as multimedia |
| `has_not_been_viewed_much` | `Boolean` | Whether the artwork hasn't been visited on our website very much |
| `id` | `String` | Unique identifier of this resource. |
| `image_embedding` | `Object` | The generated embeddings describing the artwork image |
| `image_id` | `String` | Unique identifier of the preferred image to use to represent this work |
| `inscriptions` | `Object` | A description of distinguishing or identifying physical markings that are on the work |
| `internal_department_id` | `String` | An internal department id we use for analytics. |
| `is_boosted` | `Boolean` | Whether this document should be boosted in search |
| `is_on_view` | `Boolean` | Whether the work is on display |
| `is_public_domain` | `Boolean` | Whether the work is in the public domain, meaning it was created before copyrights existed or has left the copyright term |
| `is_zoomable` | `Boolean` | Whether images of the work are allowed to be displayed in a zoomable interface. |
| `latitude` | `Float` | Latitude coordinate of the location of this work in our galleries |
| `latlon` | `Object` | Latitude and longitude coordinates of the location of this work in our galleries |
| `longitude` | `Float` | Longitude coordinate of the location of this work in our galleries |
| `main_reference_number` | `Integer` | Unique identifier assigned to the artwork upon acquisition |
| `material_id` | `String` | Unique identifier of the preferred material term for this work |
| `material_ids` | `Object` | Unique identifiers of all material terms for this work |
| `material_titles` | `Object` | The names of all material terms related to this artwork |
| `max_zoom_window_size` | `Object` | The maximum size of the window the image is allowed to be viewed in, in pixels. |
| `medium_display` | `Object` | The substances or materials used in the creation of a work |
| `nomisma_id` | `String` | Unique identifier of this work in the nomisma coin database |
| `on_loan_display` | `Object` | If an artwork is on loan, this contains details about the loan |
| `pageviews` | `Object` | Approx. |
| `pageviews_recent` | `Object` | Approx. |
| `place_of_origin` | `Object` | The location where the creation, design, or production of the work took place, or the original location of the work |
| `provenance_text` | `Object` | Ownership/collecting history of the work. |
| `publication_history` | `Object` | Bibliographic list of all the places this work has been published |
| `publishing_verification_level` | `Object` | Indicator of how much metadata on the work in published. |
| `section_ids` | `Object` | Unique identifiers of the digital publication chapters this work in included in |
| `section_titles` | `Object` | Names of the digital publication chapters this work is included in |
| `short_description` | `Object` | Short explanation describing the work |
| `site_ids` | `Object` | Unique identifiers of the microsites this work is a part of |
| `sound_ids` | `Object` | Unique identifiers of the audio about this work |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `style_id` | `String` | Unique identifier of the preferred style term for this work |
| `style_ids` | `Object` | Unique identifiers of all style terms for this work |
| `style_title` | `Object` | The name of the preferred style term for this work |
| `style_titles` | `Object` | The names of all style terms related to this artwork |
| `subject_id` | `String` | Unique identifier of the preferred subject term for this work |
| `subject_ids` | `Object` | Unique identifiers of all subject terms for this work |
| `subject_titles` | `Object` | The names of all subject terms related to this artwork |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `technique_id` | `String` | Unique identifier of the preferred technique term for this work |
| `technique_ids` | `Object` | Unique identifiers of all technique terms for this work |
| `technique_titles` | `Object` | The names of all technique terms related to this artwork |
| `term_titles` | `Object` | The names of the taxonomy tags for this work |
| `text_embedding` | `Object` | The generated embeddings of artwork text |
| `text_ids` | `Object` | Unique identifiers of the texts about this work |
| `theme_titles` | `Object` | The names of all thematic publish categories related to this artwork |
| `thumbnail` | `Object` | Metadata about the image referenced by `image_id`. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |
| `video_ids` | `Object` | Unique identifiers of the videos about this work |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Artwork record (raises on error).
artwork = client.Artwork.load({ "id" => "artwork_id" })
```

#### Example: List

```ruby
# list returns an Array of Artwork records (raises on error).
artworks = client.Artwork.list
```


### ArtworkDateQualifier

Create an instance: `artwork_date_qualifier = client.ArtworkDateQualifier`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `id` | `String` | Unique identifier of this resource. |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the ArtworkDateQualifier record (raises on error).
artwork_date_qualifier = client.ArtworkDateQualifier.load({ "id" => "artwork_date_qualifier_id" })
```

#### Example: List

```ruby
# list returns an Array of ArtworkDateQualifier records (raises on error).
artwork_date_qualifiers = client.ArtworkDateQualifier.list
```


### ArtworkPlaceQualifier

Create an instance: `artwork_place_qualifier = client.ArtworkPlaceQualifier`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `id` | `String` | Unique identifier of this resource. |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the ArtworkPlaceQualifier record (raises on error).
artwork_place_qualifier = client.ArtworkPlaceQualifier.load({ "id" => "artwork_place_qualifier_id" })
```

#### Example: List

```ruby
# list returns an Array of ArtworkPlaceQualifier records (raises on error).
artwork_place_qualifiers = client.ArtworkPlaceQualifier.list
```


### ArtworkType

Create an instance: `artwork_type = client.ArtworkType`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `aat_id` | `String` | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `id` | `String` | Unique identifier of this resource. |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the ArtworkType record (raises on error).
artwork_type = client.ArtworkType.load({ "id" => "artwork_type_id" })
```

#### Example: List

```ruby
# list returns an Array of ArtworkType records (raises on error).
artwork_types = client.ArtworkType.list
```


### CategoryTerm

Create an instance: `category_term = client.CategoryTerm`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `aat_id` | `String` | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `id` | `String` | Unique identifier of this resource. |
| `parent_id` | `String` | Unique identifier of this category's parent |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `subtype` | `Object` | Takes one of the following values: classification, material, technique, style, subject, department, theme |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the CategoryTerm record (raises on error).
category_term = client.CategoryTerm.load({ "id" => "category_term_id" })
```

#### Example: List

```ruby
# list returns an Array of CategoryTerm records (raises on error).
category_terms = client.CategoryTerm.list
```


### DigitalPublication

Create an instance: `digital_publication = client.DigitalPublication`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `copy` | `Object` | The text of the page |
| `id` | `String` | Unique identifier of this resource. |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | The URL to this page on our website |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the DigitalPublication record (raises on error).
digital_publication = client.DigitalPublication.load({ "id" => "digital_publication_id" })
```

#### Example: List

```ruby
# list returns an Array of DigitalPublication records (raises on error).
digital_publications = client.DigitalPublication.list
```


### DigitalPublicationArticle

Create an instance: `digital_publication_article = client.DigitalPublicationArticle`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `author_display` | `Object` | A display-friendly text of the authors of this article |
| `copy` | `Object` | The text of the article |
| `digital_publication_id` | `String` | Unique identifier of the digital publication this article belongs to |
| `id` | `String` | Unique identifier of this resource. |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | The URL to this article on our website |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the DigitalPublicationArticle record (raises on error).
digital_publication_article = client.DigitalPublicationArticle.load({ "id" => "digital_publication_article_id" })
```

#### Example: List

```ruby
# list returns an Array of DigitalPublicationArticle records (raises on error).
digital_publication_articles = client.DigitalPublicationArticle.list
```


### EducatorResource

Create an instance: `educator_resource = client.EducatorResource`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `copy` | `Object` | The text of the page |
| `id` | `String` | Unique identifier of this resource. |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | The URL to this page on our website |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the EducatorResource record (raises on error).
educator_resource = client.EducatorResource.load({ "id" => "educator_resource_id" })
```

#### Example: List

```ruby
# list returns an Array of EducatorResource records (raises on error).
educator_resources = client.EducatorResource.list
```


### Event

Create an instance: `event = client.Event`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_audience_ids` | `Object` | Unique identifiers indicating the alternate audiences for this event |
| `alt_event_type_ids` | `Object` | Unique identifiers indicating the alternate types of this event |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `audience_id` | `String` | Unique identifier indicating the preferred audience for this event |
| `buy_button_caption` | `Object` | Additional text below the ticket/registration button |
| `buy_button_text` | `Object` | The text used on the ticket/registration button |
| `date_display` | `Object` | A readable display of the event dates |
| `description` | `String` | All copytext of the event |
| `door_time` | `Object` | The time the doors open for this event |
| `end_date` | `Object` | The date the event ends |
| `end_time` | `Object` | The time the event ends |
| `entrance` | `Object` | Which entrance to use for this event |
| `event_host_id` | `String` | Unique identifier of the host (cf. |
| `event_host_title` | `Object` | Unique identifier of the host (cf. |
| `event_type_id` | `String` | Unique identifier indicating the preferred type of this event |
| `header_description` | `Object` | Brief description of the event displayed below the title |
| `hero_caption` | `Object` | Text displayed with the hero image on the event |
| `id` | `String` | Unique identifier of this resource. |
| `image_url` | `Object` | The URL of an image representing this page |
| `is_admission_required` | `Boolean` | Whether admission to the museum is required to attend this event |
| `is_after_hours` | `Boolean` | Whether the event is to be held after the museum closes |
| `is_free` | `Boolean` | Whether the event is free |
| `is_member_exclusive` | `Boolean` | Whether the event is exclusive to members of the museum |
| `is_private` | `Boolean` | Whether the event is private |
| `is_registration_required` | `Boolean` | Whether registration is required to attend the event |
| `is_sales_button_hidden` | `Boolean` | Whether the buy tickets button is hidden on the website event page |
| `is_sold_out` | `Boolean` | Whether the event is sold out |
| `is_ticketed` | `Boolean` | Whether a ticket is required to attend the event |
| `is_virtual_event` | `Boolean` | Whether the event is being held virtually |
| `join_url` | `Object` | URL to the membership signup page via this event |
| `layout_type` | `Object` | Number indicating the type of layout this event page uses |
| `list_description` | `Object` | One-sentence description of the event displayed in listings |
| `location` | `Object` | Where the event takes place |
| `program_ids` | `Object` | Unique identifiers indicating the programs this event is a part of |
| `program_titles` | `Object` | Titles of the programs this event is a part of |
| `rsvp_link` | `Object` | The URL to the sales site for this event |
| `search_tags` | `Object` | Editor-specified list of tags to aid in internal search |
| `short_description` | `Object` | Brief description of the event |
| `slug` | `String` | A string used in the URL for this event |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `start_date` | `Object` | The date the event begins |
| `start_time` | `Object` | The time the event starts |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `survey_url` | `Object` | URL to the survey associated with this event |
| `ticketed_event_id` | `String` | Unique identifier of the event in the ticketing system this website event is tied to |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `title_display` | `Object` | The name of this event formatted with HTML (optional) |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |
| `virtual_event_passcode` | `Object` | Passcode to access the virtual event |
| `virtual_event_url` | `Object` | URL to the virtual event |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Event record (raises on error).
event = client.Event.load({ "id" => "event_id" })
```

#### Example: List

```ruby
# list returns an Array of Event records (raises on error).
events = client.Event.list
```


### EventOccurrence

Create an instance: `event_occurrence = client.EventOccurrence`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `button_caption` | `Object` | Additional text below the ticket/registration button |
| `button_text` | `Object` | The text used on the ticket/registration button |
| `button_url` | `Object` | The URL to the sales site or an RSVP link for this event |
| `description` | `String` | Description of the event |
| `end_at` | `Object` | The date the event occurrence ends |
| `event_id` | `String` | Identifier of the master event of which this is an occurrence |
| `id` | `String` | Unique identifier of this resource. |
| `image_url` | `Object` | The URL of an image representing this page |
| `is_private` | `Boolean` | Whether the event is private. |
| `is_sales_button_hidden` | `Boolean` | Whether the buy tickets button is hidden on the website event page |
| `is_ticketed` | `Boolean` | Whether a ticket is required to attend the event |
| `location` | `Object` | Where the event takes place |
| `off_sale_at` | `Object` | Date and time the event goes off sale |
| `on_sale_at` | `Object` | Date and time the event goes on sale |
| `short_description` | `Object` | Brief description of the event |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `start_at` | `Object` | The date the event occurrence begins |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `title_display` | `Object` | The name of this event formatted with HTML (optional) |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the EventOccurrence record (raises on error).
event_occurrence = client.EventOccurrence.load({ "id" => "event_occurrence_id" })
```

#### Example: List

```ruby
# list returns an Array of EventOccurrence records (raises on error).
event_occurrences = client.EventOccurrence.list
```


### EventProgram

Create an instance: `event_program = client.EventProgram`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `id` | `String` | Unique identifier of this resource. |
| `is_affiliate_group` | `Boolean` | Whether this program represents an affiliate group |
| `is_event_host` | `Boolean` | Whether this program represents an event host |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the EventProgram record (raises on error).
event_program = client.EventProgram.load({ "id" => "event_program_id" })
```

#### Example: List

```ruby
# list returns an Array of EventProgram records (raises on error).
event_programs = client.EventProgram.list
```


### Exhibition

Create an instance: `exhibition = client.Exhibition`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `aic_end_at` | `Object` | Date the exhibition closed at the Art Institute of Chicago |
| `aic_start_at` | `Object` | Date the exhibition opened at the Art Institute of Chicago |
| `alt_image_ids` | `Object` | Unique identifiers of all non-preferred images of this exhibition. |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `artist_ids` | `Object` | Unique identifiers of the artist agent records representing who was shown in the exhibition |
| `artwork_ids` | `Object` | Unique identifiers of the artworks that were part of the exhibition |
| `artwork_titles` | `Object` | Names of the artworks that were part of the exhibition |
| `document_ids` | `Object` | Unique identifiers of assets that serve as documentation for this exhibition |
| `gallery_id` | `String` | Unique identifier of the gallery that mainly housed the exhibition |
| `gallery_title` | `Object` | The name of the gallery that mainly housed the exhibition |
| `id` | `String` | Unique identifier of this resource. |
| `image_id` | `String` | Unique identifier of the preferred image to use to represent this exhibition |
| `image_url` | `Object` | URL to the hero image from the website |
| `is_featured` | `Boolean` | Is this exhibition currently featured on our website? |
| `is_published` | `Boolean` | Is this exhibition currently published on our website? |
| `position` | `Object` | Numering position represnting the order in which this exhibition is featured on the website |
| `short_description` | `Object` | Brief explanation of what this exhibition is |
| `site_ids` | `Object` | Unique identifiers of the microsites this exhibition is a part of |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `status` | `Object` | Whether the exhibition is open or closed |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | URL to this exhibition on our website |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Exhibition record (raises on error).
exhibition = client.Exhibition.load({ "id" => "exhibition_id" })
```

#### Example: List

```ruby
# list returns an Array of Exhibition records (raises on error).
exhibitions = client.Exhibition.list
```


### Gallery

Create an instance: `gallery = client.Gallery`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `floor` | `Object` | The level the gallery is on, e.g., 1, 2, 3, or LL |
| `id` | `String` | Unique identifier of this resource. |
| `is_closed` | `Boolean` | Whether the gallery is currently closed |
| `latitude` | `Float` | Latitude coordinate of the center of the room |
| `latlon` | `Object` | Latitude and longitude coordinates of the center of the room |
| `longitude` | `Float` | Longitude coordinate of the center of the room |
| `number` | `Object` | The gallery's room number. |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | `String` | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Gallery record (raises on error).
gallery = client.Gallery.load({ "id" => "gallery_id" })
```

#### Example: List

```ruby
# list returns an Array of Gallery records (raises on error).
gallerys = client.Gallery.list
```


### GenericPage

Create an instance: `generic_page = client.GenericPage`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `copy` | `Object` | The text of the page |
| `id` | `String` | Unique identifier of this resource. |
| `search_tags` | `Object` | Editor-specified list of tags to aid in internal search |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | The URL to this page on our website |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the GenericPage record (raises on error).
generic_page = client.GenericPage.load({ "id" => "generic_page_id" })
```

#### Example: List

```ruby
# list returns an Array of GenericPage records (raises on error).
generic_pages = client.GenericPage.list
```


### Highlight

Create an instance: `highlight = client.Highlight`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `copy` | `Object` | The text of the highlight description |
| `id` | `String` | Unique identifier of this resource. |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Highlight record (raises on error).
highlight = client.Highlight.load({ "id" => "highlight_id" })
```

#### Example: List

```ruby
# list returns an Array of Highlight records (raises on error).
highlights = client.Highlight.list
```


### Hour

Create an instance: `hour = client.Hour`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `additional_text` | `Object` | Additional information about the hours |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `friday_is_closed` | `Object` | Whether the museum is closed on Fridays |
| `friday_member_close` | `Object` | The time member hours ends on Fridays |
| `friday_member_open` | `Object` | The time member hours starts on Fridays |
| `friday_public_close` | `Object` | The time public hours ends on Fridays |
| `friday_public_open` | `Object` | The time public hours starts on Fridays |
| `id` | `String` | Unique identifier of this resource. |
| `monday_is_closed` | `Object` | Whether the museum is closed on Mondays |
| `monday_member_close` | `Object` | The time member hours ends on Mondays |
| `monday_member_open` | `Object` | The time member hours starts on Mondays |
| `monday_public_close` | `Object` | The time public hours ends on Mondays |
| `monday_public_open` | `Object` | The time public hours starts on Mondays |
| `saturday_is_closed` | `Object` | Whether the museum is closed on Saturdays |
| `saturday_member_close` | `Object` | The time member hours ends on Saturdays |
| `saturday_member_open` | `Object` | The time member hours starts on Saturdays |
| `saturday_public_close` | `Object` | The time public hours ends on Saturdays |
| `saturday_public_open` | `Object` | The time public hours starts on Saturdays |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `summary` | `Object` | Readable summary of the hours |
| `sunday_is_closed` | `Object` | Whether the museum is closed on Sundays |
| `sunday_member_close` | `Object` | The time member hours ends on Sundays |
| `sunday_member_open` | `Object` | The time member hours starts on Sundays |
| `sunday_public_close` | `Object` | The time public hours ends on Sundays |
| `sunday_public_open` | `Object` | The time public hours starts on Sundays |
| `thursday_is_closed` | `Object` | Whether the museum is closed on Thursdays |
| `thursday_member_close` | `Object` | The time member hours ends on Thursdays |
| `thursday_member_open` | `Object` | The time member hours starts on Thursdays |
| `thursday_public_close` | `Object` | The time public hours ends on Thursdays |
| `thursday_public_open` | `Object` | The time public hours starts on Thursdays |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `tuesday_is_closed` | `Object` | Whether the museum is closed on Tuesdays |
| `tuesday_member_close` | `Object` | The time member hours ends on Tuesdays |
| `tuesday_member_open` | `Object` | The time member hours starts on Tuesdays |
| `tuesday_public_close` | `Object` | The time public hours ends on Tuesdays |
| `tuesday_public_open` | `Object` | The time public hours starts on Tuesdays |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |
| `wednesday_is_closed` | `Object` | Whether the museum is closed on Wednesdays |
| `wednesday_member_close` | `Object` | The time member hours ends on Wednesdays |
| `wednesday_member_open` | `Object` | The time member hours starts on Wednesdays |
| `wednesday_public_close` | `Object` | The time public hours ends on Wednesdays |
| `wednesday_public_open` | `Object` | The time public hours starts on Wednesdays |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Hour record (raises on error).
hour = client.Hour.load({ "id" => "hour_id" })
```

#### Example: List

```ruby
# list returns an Array of Hour records (raises on error).
hours = client.Hour.list
```


### Image

Create an instance: `image = client.Image`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ahash` | `Object` | Image hash generated using ahash algorithm with 64 boolean subfields |
| `alt_text` | `Object` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `artwork_ids` | `Object` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `Object` | Names of the artworks associated with this asset |
| `color` | `Object` | Dominant color of this image in HSL |
| `colorfulness` | `Object` | Unbounded positive float representing an abstract measure of colorfulness. |
| `content` | `Object` | Text of or URL to the contents of this asset |
| `content_e_tag` | `Object` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `Object` | Asset-specific copyright information |
| `fingerprint` | `Object` | Image hashes: aHash, dHash, pHash, wHash |
| `height` | `Float` | Native height of the image |
| `id` | `String` | Unique identifier of this resource. |
| `iiif_url` | `Object` | IIIF URL of this image |
| `is_educational_resource` | `Boolean` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `Boolean` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `Boolean` | Whether this resource is considered to be educational |
| `lake_guid` | `Object` | Unique UUID of this resource in LAKE, our DAMS. |
| `lqip` | `Object` | Low-quality image placeholder (LQIP). |
| `phash` | `Object` | Image hash generated using phash algorithm with 64 boolean subfields |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `type` | `Object` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |
| `width` | `Float` | Native width of the image |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Image record (raises on error).
image = client.Image.load({ "id" => "image_id" })
```

#### Example: List

```ruby
# list returns an Array of Image records (raises on error).
images = client.Image.list
```


### LandingPage

Create an instance: `landing_page = client.LandingPage`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `copy` | `Object` | The text of the page |
| `id` | `String` | Unique identifier of this resource. |
| `search_tags` | `Object` | Editor-specified list of tags to aid in internal search |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | The URL to this page on our website |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the LandingPage record (raises on error).
landing_page = client.LandingPage.load({ "id" => "landing_page_id" })
```

#### Example: List

```ruby
# list returns an Array of LandingPage records (raises on error).
landing_pages = client.LandingPage.list
```


### Place

Create an instance: `place = client.Place`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `id` | `String` | Unique identifier of this resource. |
| `latitude` | `Float` | Latitude coordinate of the center of the room |
| `longitude` | `Float` | Longitude coordinate of the center of the room |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | `String` | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Place record (raises on error).
place = client.Place.load({ "id" => "place_id" })
```

#### Example: List

```ruby
# list returns an Array of Place records (raises on error).
places = client.Place.list
```


### PressRelease

Create an instance: `press_release = client.PressRelease`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `copy` | `Object` | The text of the page |
| `id` | `String` | Unique identifier of this resource. |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | The URL to this page on our website |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the PressRelease record (raises on error).
press_release = client.PressRelease.load({ "id" => "press_release_id" })
```

#### Example: List

```ruby
# list returns an Array of PressRelease records (raises on error).
press_releases = client.PressRelease.list
```


### PrintedPublication

Create an instance: `printed_publication = client.PrintedPublication`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `copy` | `Object` | The text of the page |
| `id` | `String` | Unique identifier of this resource. |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | The URL to this page on our website |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the PrintedPublication record (raises on error).
printed_publication = client.PrintedPublication.load({ "id" => "printed_publication_id" })
```

#### Example: List

```ruby
# list returns an Array of PrintedPublication records (raises on error).
printed_publications = client.PrintedPublication.list
```


### Product

Create an instance: `product = client.Product`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `artist_ids` | `Object` | Unique identifiers of the artists associated with this product |
| `artwork_ids` | `Object` | Unique identifiers of the artworks associated with this product |
| `description` | `String` | Explanation of what this product is |
| `exhibition_ids` | `Object` | Unique identifiers of the exhibitions associated with this product |
| `external_sku` | `Object` | Numeric product identification code of a machine-readable barcode, when the customer sku differs from our internal one |
| `id` | `String` | Unique identifier of this resource. |
| `image_url` | `Object` | URL of an image for this product |
| `max_compare_at_price` | `Object` | Number indicating how much the most expensive variant of a product cost before a sale |
| `max_current_price` | `Object` | Number indicating how much the most expensive variant of a product costs right now |
| `min_compare_at_price` | `Object` | Number indicating how much the least expensive variant of a product cost before a sale |
| `min_current_price` | `Object` | Number indicating how much the least expensive variant of a product costs right now |
| `price_display` | `Object` | Explanation of what this product is |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | URL of this product in the shop |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Product record (raises on error).
product = client.Product.load({ "id" => "product_id" })
```

#### Example: List

```ruby
# list returns an Array of Product records (raises on error).
products = client.Product.list
```


### Publication

Create an instance: `publication = client.Publication`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `id` | `String` | Unique identifier of this resource. |
| `section_ids` | `Object` | Unique identifiers of the sections of this publication |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | URL to the publication |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Publication record (raises on error).
publication = client.Publication.load({ "id" => "publication_id" })
```

#### Example: List

```ruby
# list returns an Array of Publication records (raises on error).
publications = client.Publication.list
```


### Search

Create an instance: `search = client.Search`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_id` | `String` | API unique identifier |
| `api_link` | `Object` | URL to this recource in the API |
| `api_model` | `Object` | Name of the model the resource represents |
| `id` | `String` | Unique identifier within the search index |
| `is_boosted` | `Boolean` | Whether this record has been flagged to be boosted |
| `score` | `Float` | Search index ranking of the result |
| `thumbnail` | `Object` | Metadata on the image representing this record |
| `timestamp` | `Object` | Date this record was last updated in the API |
| `title` | `String` | The name of this resource |

#### Example: List

```ruby
# list returns an Array of Search records (raises on error).
searchs = client.Search.list
```


### Section

Create an instance: `section = client.Section`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accession` | `Object` | An accession number parsed from the title or tombstone |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `artwork_id` | `String` | Unique identifier of the artwork with which this section is associated |
| `content` | `Object` | Content of this section in plaintext |
| `generic_page_id` | `String` | Unique identifier of the page on the website that represents the publication this section belongs to |
| `id` | `String` | Unique identifier of this resource. |
| `publication_id` | `String` | Unique identifier of the publication this section belongs to |
| `publication_title` | `Object` | Name of the publication this section belongs to |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | URL to the section |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Section record (raises on error).
section = client.Section.load({ "id" => "section_id" })
```

#### Example: List

```ruby
# list returns an Array of Section records (raises on error).
sections = client.Section.list
```


### Site

Create an instance: `site = client.Site`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `artwork_ids` | `Object` | Unique identifiers of the artworks this site is associated with |
| `artwork_titles` | `Object` | Names of the artworks this site is associated with |
| `description` | `String` | Explanation of what this site is |
| `exhibition_ids` | `Object` | Unique identifier of the exhibitions this site is associated with |
| `exhibition_titles` | `Object` | Names of the exhibitions this site is associated with |
| `id` | `String` | Unique identifier of this resource. |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | URL to this site |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Site record (raises on error).
site = client.Site.load({ "id" => "site_id" })
```

#### Example: List

```ruby
# list returns an Array of Site records (raises on error).
sites = client.Site.list
```


### Sound

Create an instance: `sound = client.Sound`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_text` | `Object` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `artwork_ids` | `Object` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `Object` | Names of the artworks associated with this asset |
| `content` | `Object` | Text of or URL to the contents of this asset |
| `content_e_tag` | `Object` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `Object` | Asset-specific copyright information |
| `id` | `String` | Unique identifier of this resource. |
| `is_educational_resource` | `Boolean` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `Boolean` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `Boolean` | Whether this resource is considered to be educational |
| `lake_guid` | `Object` | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | Name of this mobile audio file – derived from the artwork and tour titles |
| `transcript` | `Object` | Text transcription of the audio file |
| `type` | `Object` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | URL to the audio file |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Sound record (raises on error).
sound = client.Sound.load({ "id" => "sound_id" })
```

#### Example: List

```ruby
# list returns an Array of Sound records (raises on error).
sounds = client.Sound.list
```


### StaticPage

Create an instance: `static_page = client.StaticPage`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `id` | `String` | Unique identifier of this resource. |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |
| `web_url` | `Object` | The URL to this page on our website |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the StaticPage record (raises on error).
static_page = client.StaticPage.load({ "id" => "static_page_id" })
```

#### Example: List

```ruby
# list returns an Array of StaticPage records (raises on error).
static_pages = client.StaticPage.list
```


### Text

Create an instance: `text = client.Text`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_text` | `Object` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `artwork_ids` | `Object` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `Object` | Names of the artworks associated with this asset |
| `content` | `Object` | Text of or URL to the contents of this asset |
| `content_e_tag` | `Object` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `Object` | Asset-specific copyright information |
| `id` | `String` | Unique identifier of this resource. |
| `is_educational_resource` | `Boolean` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `Boolean` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `Boolean` | Whether this resource is considered to be educational |
| `lake_guid` | `Object` | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `type` | `Object` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Text record (raises on error).
text = client.Text.load({ "id" => "text_id" })
```

#### Example: List

```ruby
# list returns an Array of Text records (raises on error).
texts = client.Text.list
```


### Tour

Create an instance: `tour = client.Tour`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `artist_titles` | `Object` | Names of the artists of the artworks featured in this tour's tour stops |
| `artwork_titles` | `Object` | Names of the artworks featured in this tour's tour stops |
| `description` | `String` | Explanation of what the tour is |
| `id` | `String` | Unique identifier of this resource. |
| `image` | `Object` | The main image for the tour |
| `intro` | `Object` | Text introducing the tour |
| `intro_link` | `Object` | Link to the audio file of the introduction |
| `intro_transcript` | `Object` | Transcript of the introduction audio to the tour |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |
| `weight` | `Float` | Number representing this tour's sort order |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Tour record (raises on error).
tour = client.Tour.load({ "id" => "tour_id" })
```

#### Example: List

```ruby
# list returns an Array of Tour records (raises on error).
tours = client.Tour.list
```


### Video

Create an instance: `video = client.Video`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_text` | `Object` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `Object` | REST API link for this resource |
| `api_model` | `Object` | REST API resource type or endpoint |
| `artwork_ids` | `Object` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `Object` | Names of the artworks associated with this asset |
| `content` | `Object` | Text of or URL to the contents of this asset |
| `content_e_tag` | `Object` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `Object` | Asset-specific copyright information |
| `id` | `String` | Unique identifier of this resource. |
| `is_educational_resource` | `Boolean` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `Boolean` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `Boolean` | Whether this resource is considered to be educational |
| `lake_guid` | `Object` | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `Object` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Object` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Object` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Object` | Date and time the record was updated in the aggregator search index |
| `title` | `String` | The name of this resource |
| `type` | `Object` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `Object` | Date and time the record was updated in the aggregator database |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Video record (raises on error).
video = client.Video.load({ "id" => "video_id" })
```

#### Example: List

```ruby
# list returns an Array of Video records (raises on error).
videos = client.Video.list
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

Features are the extension mechanism. A feature is a Ruby class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as hashes

The Ruby SDK uses plain Ruby hashes throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers.to_map()` to safely validate that a value is a hash.

### Module structure

```
rb/
├── ArtInstituteOfChicago_sdk.rb       -- Main SDK module
├── config.rb                  -- Configuration
├── features.rb                -- Feature factory
├── core/                      -- Core types and context
├── entity/                    -- Entity implementations
├── feature/                   -- Built-in features (Base, Test, Log)
├── utility/                   -- Utility functions and struct library
└── test/                      -- Test suites
```

The main module (`ArtInstituteOfChicago_sdk`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```ruby
publication = client.Publication
publication.list()

# publication.data_get now returns the publication data from the last list
# publication.match_get returns the last match criteria
```

Call `make` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
