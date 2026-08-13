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
| `alt_titles` |  |
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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

API path: `/articles`

#### Artwork

| Field | Description |
| --- | --- |
| `alt_artist_ids` |  |
| `alt_classification_ids` |  |
| `alt_image_ids` |  |
| `alt_material_ids` |  |
| `alt_style_ids` |  |
| `alt_subject_ids` |  |
| `alt_technique_ids` |  |
| `alt_titles` |  |
| `api_link` |  |
| `api_model` |  |
| `artist_display` |  |
| `artist_id` |  |
| `artist_ids` |  |
| `artist_title` |  |
| `artist_titles` |  |
| `artwork_type_id` |  |
| `artwork_type_title` |  |
| `boost_rank` |  |
| `catalog_based_search_keyword_titles` |  |
| `catalogue_display` |  |
| `category_ids` |  |
| `category_titles` |  |
| `classification_id` |  |
| `classification_ids` |  |
| `classification_title` |  |
| `classification_titles` |  |
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
| `dimensions` |  |
| `dimensions_detail` |  |
| `document_ids` |  |
| `edition` |  |
| `exhibition_history` |  |
| `fiscal_year` |  |
| `fiscal_year_deaccession` |  |
| `gallery_id` |  |
| `gallery_title` |  |
| `has_advanced_imaging` |  |
| `has_educational_resources` |  |
| `has_multimedia_resources` |  |
| `has_not_been_viewed_much` |  |
| `id` |  |
| `image_embedding` |  |
| `image_id` |  |
| `inscriptions` |  |
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
| `material_ids` |  |
| `material_titles` |  |
| `max_zoom_window_size` |  |
| `medium_display` |  |
| `nomisma_id` |  |
| `on_loan_display` |  |
| `pageviews` |  |
| `pageviews_recent` |  |
| `place_of_origin` |  |
| `provenance_text` |  |
| `publication_history` |  |
| `publishing_verification_level` |  |
| `section_ids` |  |
| `section_titles` |  |
| `short_description` |  |
| `site_ids` |  |
| `sound_ids` |  |
| `source_updated_at` |  |
| `style_id` |  |
| `style_ids` |  |
| `style_title` |  |
| `style_titles` |  |
| `subject_id` |  |
| `subject_ids` |  |
| `subject_titles` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `technique_id` |  |
| `technique_ids` |  |
| `technique_titles` |  |
| `term_titles` |  |
| `text_embedding` |  |
| `text_ids` |  |
| `theme_titles` |  |
| `thumbnail` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `video_ids` |  |

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

API path: `/educator-resources`

#### Event

| Field | Description |
| --- | --- |
| `alt_audience_ids` |  |
| `alt_event_type_ids` |  |
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
| `is_after_hours` |  |
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
| `program_ids` |  |
| `program_titles` |  |
| `rsvp_link` |  |
| `search_tags` |  |
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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

API path: `/event-programs`

#### Exhibition

| Field | Description |
| --- | --- |
| `aic_end_at` |  |
| `aic_start_at` |  |
| `alt_image_ids` |  |
| `api_link` |  |
| `api_model` |  |
| `artist_ids` |  |
| `artwork_ids` |  |
| `artwork_titles` |  |
| `document_ids` |  |
| `gallery_id` |  |
| `gallery_title` |  |
| `id` |  |
| `image_id` |  |
| `image_url` |  |
| `is_featured` |  |
| `is_published` |  |
| `position` |  |
| `short_description` |  |
| `site_ids` |  |
| `source_updated_at` |  |
| `status` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: List, Load.

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

Operations: List, Load.

API path: `/galleries`

#### GenericPage

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `copy` |  |
| `id` |  |
| `search_tags` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

API path: `/hours`

#### Image

| Field | Description |
| --- | --- |
| `ahash` |  |
| `alt_text` |  |
| `api_link` |  |
| `api_model` |  |
| `artwork_ids` |  |
| `artwork_titles` |  |
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

Operations: List, Load.

API path: `/images`

#### LandingPage

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `copy` |  |
| `id` |  |
| `search_tags` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

API path: `/printed-publications`

#### Product

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `artist_ids` |  |
| `artwork_ids` |  |
| `description` |  |
| `exhibition_ids` |  |
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

Operations: List, Load.

API path: `/products`

#### Publication

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `id` |  |
| `section_ids` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: List, Load.

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

Operations: List.

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

Operations: List, Load.

API path: `/sections`

#### Site

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `artwork_ids` |  |
| `artwork_titles` |  |
| `description` |  |
| `exhibition_ids` |  |
| `exhibition_titles` |  |
| `id` |  |
| `source_updated_at` |  |
| `suggest_autocomplete_all` |  |
| `suggest_autocomplete_boosted` |  |
| `timestamp` |  |
| `title` |  |
| `updated_at` |  |
| `web_url` |  |

Operations: List, Load.

API path: `/sites`

#### Sound

| Field | Description |
| --- | --- |
| `alt_text` |  |
| `api_link` |  |
| `api_model` |  |
| `artwork_ids` |  |
| `artwork_titles` |  |
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

Operations: List, Load.

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

Operations: List, Load.

API path: `/static-pages`

#### Text

| Field | Description |
| --- | --- |
| `alt_text` |  |
| `api_link` |  |
| `api_model` |  |
| `artwork_ids` |  |
| `artwork_titles` |  |
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

Operations: List, Load.

API path: `/texts`

#### Tour

| Field | Description |
| --- | --- |
| `api_link` |  |
| `api_model` |  |
| `artist_titles` |  |
| `artwork_titles` |  |
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

Operations: List, Load.

API path: `/tours`

#### Video

| Field | Description |
| --- | --- |
| `alt_text` |  |
| `api_link` |  |
| `api_model` |  |
| `artwork_ids` |  |
| `artwork_titles` |  |
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
| `alt_titles` | `Object` |  |
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `birth_date` | `Object` |  |
| `death_date` | `Object` |  |
| `description` | `String` |  |
| `id` | `String` |  |
| `is_artist` | `Boolean` |  |
| `sort_title` | `Object` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `ulan_id` | `String` |  |
| `updated_at` | `Object` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `id` | `String` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `id` | `String` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `copy` | `Object` |  |
| `id` | `String` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |

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
| `alt_artist_ids` | `Object` |  |
| `alt_classification_ids` | `Object` |  |
| `alt_image_ids` | `Object` |  |
| `alt_material_ids` | `Object` |  |
| `alt_style_ids` | `Object` |  |
| `alt_subject_ids` | `Object` |  |
| `alt_technique_ids` | `Object` |  |
| `alt_titles` | `Object` |  |
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `artist_display` | `Object` |  |
| `artist_id` | `String` |  |
| `artist_ids` | `Object` |  |
| `artist_title` | `Object` |  |
| `artist_titles` | `Object` |  |
| `artwork_type_id` | `String` |  |
| `artwork_type_title` | `Object` |  |
| `boost_rank` | `Object` |  |
| `catalog_based_search_keyword_titles` | `Object` |  |
| `catalogue_display` | `Object` |  |
| `category_ids` | `Object` |  |
| `category_titles` | `Object` |  |
| `classification_id` | `String` |  |
| `classification_ids` | `Object` |  |
| `classification_title` | `Object` |  |
| `classification_titles` | `Object` |  |
| `color` | `Object` |  |
| `colorfulness` | `Object` |  |
| `copyright_notice` | `Object` |  |
| `credit_line` | `Object` |  |
| `date_display` | `Object` |  |
| `date_end` | `Object` |  |
| `date_qualifier_id` | `String` |  |
| `date_qualifier_title` | `Object` |  |
| `date_start` | `Object` |  |
| `department_id` | `String` |  |
| `department_title` | `Object` |  |
| `description` | `String` |  |
| `dimensions` | `Object` |  |
| `dimensions_detail` | `Object` |  |
| `document_ids` | `Object` |  |
| `edition` | `Object` |  |
| `exhibition_history` | `Object` |  |
| `fiscal_year` | `Object` |  |
| `fiscal_year_deaccession` | `Object` |  |
| `gallery_id` | `String` |  |
| `gallery_title` | `Object` |  |
| `has_advanced_imaging` | `Boolean` |  |
| `has_educational_resources` | `Boolean` |  |
| `has_multimedia_resources` | `Boolean` |  |
| `has_not_been_viewed_much` | `Boolean` |  |
| `id` | `String` |  |
| `image_embedding` | `Object` |  |
| `image_id` | `String` |  |
| `inscriptions` | `Object` |  |
| `internal_department_id` | `String` |  |
| `is_boosted` | `Boolean` |  |
| `is_on_view` | `Boolean` |  |
| `is_public_domain` | `Boolean` |  |
| `is_zoomable` | `Boolean` |  |
| `latitude` | `Float` |  |
| `latlon` | `Object` |  |
| `longitude` | `Float` |  |
| `main_reference_number` | `Integer` |  |
| `material_id` | `String` |  |
| `material_ids` | `Object` |  |
| `material_titles` | `Object` |  |
| `max_zoom_window_size` | `Object` |  |
| `medium_display` | `Object` |  |
| `nomisma_id` | `String` |  |
| `on_loan_display` | `Object` |  |
| `pageviews` | `Object` |  |
| `pageviews_recent` | `Object` |  |
| `place_of_origin` | `Object` |  |
| `provenance_text` | `Object` |  |
| `publication_history` | `Object` |  |
| `publishing_verification_level` | `Object` |  |
| `section_ids` | `Object` |  |
| `section_titles` | `Object` |  |
| `short_description` | `Object` |  |
| `site_ids` | `Object` |  |
| `sound_ids` | `Object` |  |
| `source_updated_at` | `Object` |  |
| `style_id` | `String` |  |
| `style_ids` | `Object` |  |
| `style_title` | `Object` |  |
| `style_titles` | `Object` |  |
| `subject_id` | `String` |  |
| `subject_ids` | `Object` |  |
| `subject_titles` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `technique_id` | `String` |  |
| `technique_ids` | `Object` |  |
| `technique_titles` | `Object` |  |
| `term_titles` | `Object` |  |
| `text_embedding` | `Object` |  |
| `text_ids` | `Object` |  |
| `theme_titles` | `Object` |  |
| `thumbnail` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |
| `video_ids` | `Object` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `id` | `String` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `id` | `String` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |

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
| `aat_id` | `String` |  |
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `id` | `String` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |

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
| `aat_id` | `String` |  |
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `id` | `String` |  |
| `parent_id` | `String` |  |
| `source_updated_at` | `Object` |  |
| `subtype` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `copy` | `Object` |  |
| `id` | `String` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |
| `web_url` | `Object` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `author_display` | `Object` |  |
| `copy` | `Object` |  |
| `digital_publication_id` | `String` |  |
| `id` | `String` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |
| `web_url` | `Object` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `copy` | `Object` |  |
| `id` | `String` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |
| `web_url` | `Object` |  |

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
| `alt_audience_ids` | `Object` |  |
| `alt_event_type_ids` | `Object` |  |
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `audience_id` | `String` |  |
| `buy_button_caption` | `Object` |  |
| `buy_button_text` | `Object` |  |
| `date_display` | `Object` |  |
| `description` | `String` |  |
| `door_time` | `Object` |  |
| `end_date` | `Object` |  |
| `end_time` | `Object` |  |
| `entrance` | `Object` |  |
| `event_host_id` | `String` |  |
| `event_host_title` | `Object` |  |
| `event_type_id` | `String` |  |
| `header_description` | `Object` |  |
| `hero_caption` | `Object` |  |
| `id` | `String` |  |
| `image_url` | `Object` |  |
| `is_admission_required` | `Boolean` |  |
| `is_after_hours` | `Boolean` |  |
| `is_free` | `Boolean` |  |
| `is_member_exclusive` | `Boolean` |  |
| `is_private` | `Boolean` |  |
| `is_registration_required` | `Boolean` |  |
| `is_sales_button_hidden` | `Boolean` |  |
| `is_sold_out` | `Boolean` |  |
| `is_ticketed` | `Boolean` |  |
| `is_virtual_event` | `Boolean` |  |
| `join_url` | `Object` |  |
| `layout_type` | `Object` |  |
| `list_description` | `Object` |  |
| `location` | `Object` |  |
| `program_ids` | `Object` |  |
| `program_titles` | `Object` |  |
| `rsvp_link` | `Object` |  |
| `search_tags` | `Object` |  |
| `short_description` | `Object` |  |
| `slug` | `String` |  |
| `source_updated_at` | `Object` |  |
| `start_date` | `Object` |  |
| `start_time` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `survey_url` | `Object` |  |
| `ticketed_event_id` | `String` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `title_display` | `Object` |  |
| `updated_at` | `Object` |  |
| `virtual_event_passcode` | `Object` |  |
| `virtual_event_url` | `Object` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `button_caption` | `Object` |  |
| `button_text` | `Object` |  |
| `button_url` | `Object` |  |
| `description` | `String` |  |
| `end_at` | `Object` |  |
| `event_id` | `String` |  |
| `id` | `String` |  |
| `image_url` | `Object` |  |
| `is_private` | `Boolean` |  |
| `is_sales_button_hidden` | `Boolean` |  |
| `is_ticketed` | `Boolean` |  |
| `location` | `Object` |  |
| `off_sale_at` | `Object` |  |
| `on_sale_at` | `Object` |  |
| `short_description` | `Object` |  |
| `source_updated_at` | `Object` |  |
| `start_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `title_display` | `Object` |  |
| `updated_at` | `Object` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `id` | `String` |  |
| `is_affiliate_group` | `Boolean` |  |
| `is_event_host` | `Boolean` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |

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
| `aic_end_at` | `Object` |  |
| `aic_start_at` | `Object` |  |
| `alt_image_ids` | `Object` |  |
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `artist_ids` | `Object` |  |
| `artwork_ids` | `Object` |  |
| `artwork_titles` | `Object` |  |
| `document_ids` | `Object` |  |
| `gallery_id` | `String` |  |
| `gallery_title` | `Object` |  |
| `id` | `String` |  |
| `image_id` | `String` |  |
| `image_url` | `Object` |  |
| `is_featured` | `Boolean` |  |
| `is_published` | `Boolean` |  |
| `position` | `Object` |  |
| `short_description` | `Object` |  |
| `site_ids` | `Object` |  |
| `source_updated_at` | `Object` |  |
| `status` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |
| `web_url` | `Object` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `floor` | `Object` |  |
| `id` | `String` |  |
| `is_closed` | `Boolean` |  |
| `latitude` | `Float` |  |
| `latlon` | `Object` |  |
| `longitude` | `Float` |  |
| `number` | `Object` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `tgn_id` | `String` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `copy` | `Object` |  |
| `id` | `String` |  |
| `search_tags` | `Object` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |
| `web_url` | `Object` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `copy` | `Object` |  |
| `id` | `String` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |

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
| `additional_text` | `Object` |  |
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `friday_is_closed` | `Object` |  |
| `friday_member_close` | `Object` |  |
| `friday_member_open` | `Object` |  |
| `friday_public_close` | `Object` |  |
| `friday_public_open` | `Object` |  |
| `id` | `String` |  |
| `monday_is_closed` | `Object` |  |
| `monday_member_close` | `Object` |  |
| `monday_member_open` | `Object` |  |
| `monday_public_close` | `Object` |  |
| `monday_public_open` | `Object` |  |
| `saturday_is_closed` | `Object` |  |
| `saturday_member_close` | `Object` |  |
| `saturday_member_open` | `Object` |  |
| `saturday_public_close` | `Object` |  |
| `saturday_public_open` | `Object` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `summary` | `Object` |  |
| `sunday_is_closed` | `Object` |  |
| `sunday_member_close` | `Object` |  |
| `sunday_member_open` | `Object` |  |
| `sunday_public_close` | `Object` |  |
| `sunday_public_open` | `Object` |  |
| `thursday_is_closed` | `Object` |  |
| `thursday_member_close` | `Object` |  |
| `thursday_member_open` | `Object` |  |
| `thursday_public_close` | `Object` |  |
| `thursday_public_open` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `tuesday_is_closed` | `Object` |  |
| `tuesday_member_close` | `Object` |  |
| `tuesday_member_open` | `Object` |  |
| `tuesday_public_close` | `Object` |  |
| `tuesday_public_open` | `Object` |  |
| `updated_at` | `Object` |  |
| `wednesday_is_closed` | `Object` |  |
| `wednesday_member_close` | `Object` |  |
| `wednesday_member_open` | `Object` |  |
| `wednesday_public_close` | `Object` |  |
| `wednesday_public_open` | `Object` |  |

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
| `ahash` | `Object` |  |
| `alt_text` | `Object` |  |
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `artwork_ids` | `Object` |  |
| `artwork_titles` | `Object` |  |
| `color` | `Object` |  |
| `colorfulness` | `Object` |  |
| `content` | `Object` |  |
| `content_e_tag` | `Object` |  |
| `credit_line` | `Object` |  |
| `fingerprint` | `Object` |  |
| `height` | `Float` |  |
| `id` | `String` |  |
| `iiif_url` | `Object` |  |
| `is_educational_resource` | `Boolean` |  |
| `is_multimedia_resource` | `Boolean` |  |
| `is_teacher_resource` | `Boolean` |  |
| `lake_guid` | `Object` |  |
| `lqip` | `Object` |  |
| `phash` | `Object` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `type` | `Object` |  |
| `updated_at` | `Object` |  |
| `width` | `Float` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `copy` | `Object` |  |
| `id` | `String` |  |
| `search_tags` | `Object` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |
| `web_url` | `Object` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `id` | `String` |  |
| `latitude` | `Float` |  |
| `longitude` | `Float` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `tgn_id` | `String` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `copy` | `Object` |  |
| `id` | `String` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |
| `web_url` | `Object` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `copy` | `Object` |  |
| `id` | `String` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |
| `web_url` | `Object` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `artist_ids` | `Object` |  |
| `artwork_ids` | `Object` |  |
| `description` | `String` |  |
| `exhibition_ids` | `Object` |  |
| `external_sku` | `Object` |  |
| `id` | `String` |  |
| `image_url` | `Object` |  |
| `max_compare_at_price` | `Object` |  |
| `max_current_price` | `Object` |  |
| `min_compare_at_price` | `Object` |  |
| `min_current_price` | `Object` |  |
| `price_display` | `Object` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |
| `web_url` | `Object` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `id` | `String` |  |
| `section_ids` | `Object` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |
| `web_url` | `Object` |  |

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
| `api_id` | `String` |  |
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `id` | `String` |  |
| `is_boosted` | `Boolean` |  |
| `score` | `Float` |  |
| `thumbnail` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |

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
| `accession` | `Object` |  |
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `artwork_id` | `String` |  |
| `content` | `Object` |  |
| `generic_page_id` | `String` |  |
| `id` | `String` |  |
| `publication_id` | `String` |  |
| `publication_title` | `Object` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |
| `web_url` | `Object` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `artwork_ids` | `Object` |  |
| `artwork_titles` | `Object` |  |
| `description` | `String` |  |
| `exhibition_ids` | `Object` |  |
| `exhibition_titles` | `Object` |  |
| `id` | `String` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |
| `web_url` | `Object` |  |

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
| `alt_text` | `Object` |  |
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `artwork_ids` | `Object` |  |
| `artwork_titles` | `Object` |  |
| `content` | `Object` |  |
| `content_e_tag` | `Object` |  |
| `credit_line` | `Object` |  |
| `id` | `String` |  |
| `is_educational_resource` | `Boolean` |  |
| `is_multimedia_resource` | `Boolean` |  |
| `is_teacher_resource` | `Boolean` |  |
| `lake_guid` | `Object` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `transcript` | `Object` |  |
| `type` | `Object` |  |
| `updated_at` | `Object` |  |
| `web_url` | `Object` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `id` | `String` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |
| `web_url` | `Object` |  |

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
| `alt_text` | `Object` |  |
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `artwork_ids` | `Object` |  |
| `artwork_titles` | `Object` |  |
| `content` | `Object` |  |
| `content_e_tag` | `Object` |  |
| `credit_line` | `Object` |  |
| `id` | `String` |  |
| `is_educational_resource` | `Boolean` |  |
| `is_multimedia_resource` | `Boolean` |  |
| `is_teacher_resource` | `Boolean` |  |
| `lake_guid` | `Object` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `type` | `Object` |  |
| `updated_at` | `Object` |  |

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
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `artist_titles` | `Object` |  |
| `artwork_titles` | `Object` |  |
| `description` | `String` |  |
| `id` | `String` |  |
| `image` | `Object` |  |
| `intro` | `Object` |  |
| `intro_link` | `Object` |  |
| `intro_transcript` | `Object` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `updated_at` | `Object` |  |
| `weight` | `Float` |  |

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
| `alt_text` | `Object` |  |
| `api_link` | `Object` |  |
| `api_model` | `Object` |  |
| `artwork_ids` | `Object` |  |
| `artwork_titles` | `Object` |  |
| `content` | `Object` |  |
| `content_e_tag` | `Object` |  |
| `credit_line` | `Object` |  |
| `id` | `String` |  |
| `is_educational_resource` | `Boolean` |  |
| `is_multimedia_resource` | `Boolean` |  |
| `is_teacher_resource` | `Boolean` |  |
| `lake_guid` | `Object` |  |
| `source_updated_at` | `Object` |  |
| `suggest_autocomplete_all` | `Object` |  |
| `suggest_autocomplete_boosted` | `Object` |  |
| `timestamp` | `Object` |  |
| `title` | `String` |  |
| `type` | `Object` |  |
| `updated_at` | `Object` |  |

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
