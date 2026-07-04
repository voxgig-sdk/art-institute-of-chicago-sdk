# ArtInstituteOfChicago Ruby SDK



The Ruby SDK for the ArtInstituteOfChicago API — an entity-oriented client using idiomatic Ruby conventions.

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
    puts "#{item["id"]} #{item["name"]}"
  end
rescue => err
  warn "list failed: #{err}"
end
```

### 3. Load an agent

```ruby
begin
  # load returns the bare Agent record (raises on error).
  agent = client.Agent.load({ "id" => "example_id" })
  puts agent
rescue => err
  warn "load failed: #{err}"
end
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
  warn result["err"]
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
  "entity" => { "agent" => { "test01" => { "id" => "test01" } } },
})

# load returns the bare mock record (raises on error).
agent = client.Agent.load({ "id" => "test01" })
puts agent
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
| `list` | `(reqmatch, ctrl) -> Array` | List entities matching the criteria. Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `update` | `(reqdata, ctrl) -> any` | Update an existing entity. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity. Raises on error. |
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
| `search_tag` |  |
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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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

Operations: List, Load.

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
| `alt_title` | ``$ANY`` |  |
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `birth_date` | ``$ANY`` |  |
| `death_date` | ``$ANY`` |  |
| `description` | ``$STRING`` |  |
| `id` | ``$STRING`` |  |
| `is_artist` | ``$BOOLEAN`` |  |
| `sort_title` | ``$ANY`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `ulan_id` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare Agent record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare AgentRole record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare AgentType record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `copy` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare Article record (raises on error).
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
| `alt_artist_id` | ``$STRING`` |  |
| `alt_classification_id` | ``$STRING`` |  |
| `alt_image_id` | ``$STRING`` |  |
| `alt_material_id` | ``$STRING`` |  |
| `alt_style_id` | ``$STRING`` |  |
| `alt_subject_id` | ``$STRING`` |  |
| `alt_technique_id` | ``$STRING`` |  |
| `alt_title` | ``$ANY`` |  |
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `artist_display` | ``$ANY`` |  |
| `artist_id` | ``$STRING`` |  |
| `artist_title` | ``$ANY`` |  |
| `artwork_type_id` | ``$STRING`` |  |
| `artwork_type_title` | ``$ANY`` |  |
| `boost_rank` | ``$ANY`` |  |
| `catalog_based_search_keyword_title` | ``$ANY`` |  |
| `catalogue_display` | ``$ANY`` |  |
| `category_id` | ``$STRING`` |  |
| `category_title` | ``$ANY`` |  |
| `classification_id` | ``$STRING`` |  |
| `classification_title` | ``$ANY`` |  |
| `color` | ``$ANY`` |  |
| `colorfulness` | ``$ANY`` |  |
| `copyright_notice` | ``$ANY`` |  |
| `credit_line` | ``$ANY`` |  |
| `date_display` | ``$ANY`` |  |
| `date_end` | ``$ANY`` |  |
| `date_qualifier_id` | ``$STRING`` |  |
| `date_qualifier_title` | ``$ANY`` |  |
| `date_start` | ``$ANY`` |  |
| `department_id` | ``$STRING`` |  |
| `department_title` | ``$ANY`` |  |
| `description` | ``$STRING`` |  |
| `dimension` | ``$ANY`` |  |
| `dimensions_detail` | ``$ANY`` |  |
| `document_id` | ``$STRING`` |  |
| `edition` | ``$ANY`` |  |
| `exhibition_history` | ``$ANY`` |  |
| `fiscal_year` | ``$ANY`` |  |
| `fiscal_year_deaccession` | ``$ANY`` |  |
| `gallery_id` | ``$STRING`` |  |
| `gallery_title` | ``$ANY`` |  |
| `has_advanced_imaging` | ``$BOOLEAN`` |  |
| `has_educational_resource` | ``$BOOLEAN`` |  |
| `has_multimedia_resource` | ``$BOOLEAN`` |  |
| `has_not_been_viewed_much` | ``$BOOLEAN`` |  |
| `id` | ``$STRING`` |  |
| `image_embedding` | ``$ANY`` |  |
| `image_id` | ``$STRING`` |  |
| `inscription` | ``$ANY`` |  |
| `internal_department_id` | ``$STRING`` |  |
| `is_boosted` | ``$BOOLEAN`` |  |
| `is_on_view` | ``$BOOLEAN`` |  |
| `is_public_domain` | ``$BOOLEAN`` |  |
| `is_zoomable` | ``$BOOLEAN`` |  |
| `latitude` | ``$NUMBER`` |  |
| `latlon` | ``$ANY`` |  |
| `longitude` | ``$NUMBER`` |  |
| `main_reference_number` | ``$INTEGER`` |  |
| `material_id` | ``$STRING`` |  |
| `material_title` | ``$ANY`` |  |
| `max_zoom_window_size` | ``$ANY`` |  |
| `medium_display` | ``$ANY`` |  |
| `nomisma_id` | ``$STRING`` |  |
| `on_loan_display` | ``$ANY`` |  |
| `pageview` | ``$ANY`` |  |
| `pageviews_recent` | ``$ANY`` |  |
| `place_of_origin` | ``$ANY`` |  |
| `provenance_text` | ``$ANY`` |  |
| `publication_history` | ``$ANY`` |  |
| `publishing_verification_level` | ``$ANY`` |  |
| `section_id` | ``$STRING`` |  |
| `section_title` | ``$ANY`` |  |
| `short_description` | ``$ANY`` |  |
| `site_id` | ``$STRING`` |  |
| `sound_id` | ``$STRING`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `style_id` | ``$STRING`` |  |
| `style_title` | ``$ANY`` |  |
| `subject_id` | ``$STRING`` |  |
| `subject_title` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `technique_id` | ``$STRING`` |  |
| `technique_title` | ``$ANY`` |  |
| `term_title` | ``$ANY`` |  |
| `text_embedding` | ``$ANY`` |  |
| `text_id` | ``$STRING`` |  |
| `theme_title` | ``$ANY`` |  |
| `thumbnail` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |
| `video_id` | ``$STRING`` |  |

#### Example: Load

```ruby
# load returns the bare Artwork record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare ArtworkDateQualifier record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare ArtworkPlaceQualifier record (raises on error).
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
| `aat_id` | ``$STRING`` |  |
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare ArtworkType record (raises on error).
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
| `aat_id` | ``$STRING`` |  |
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `parent_id` | ``$STRING`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `subtype` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare CategoryTerm record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `copy` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |
| `web_url` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare DigitalPublication record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `author_display` | ``$ANY`` |  |
| `copy` | ``$ANY`` |  |
| `digital_publication_id` | ``$STRING`` |  |
| `id` | ``$STRING`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |
| `web_url` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare DigitalPublicationArticle record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `copy` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |
| `web_url` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare EducatorResource record (raises on error).
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
| `alt_audience_id` | ``$STRING`` |  |
| `alt_event_type_id` | ``$STRING`` |  |
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `audience_id` | ``$STRING`` |  |
| `buy_button_caption` | ``$ANY`` |  |
| `buy_button_text` | ``$ANY`` |  |
| `date_display` | ``$ANY`` |  |
| `description` | ``$STRING`` |  |
| `door_time` | ``$ANY`` |  |
| `end_date` | ``$ANY`` |  |
| `end_time` | ``$ANY`` |  |
| `entrance` | ``$ANY`` |  |
| `event_host_id` | ``$STRING`` |  |
| `event_host_title` | ``$ANY`` |  |
| `event_type_id` | ``$STRING`` |  |
| `header_description` | ``$ANY`` |  |
| `hero_caption` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `image_url` | ``$ANY`` |  |
| `is_admission_required` | ``$BOOLEAN`` |  |
| `is_after_hour` | ``$BOOLEAN`` |  |
| `is_free` | ``$BOOLEAN`` |  |
| `is_member_exclusive` | ``$BOOLEAN`` |  |
| `is_private` | ``$BOOLEAN`` |  |
| `is_registration_required` | ``$BOOLEAN`` |  |
| `is_sales_button_hidden` | ``$BOOLEAN`` |  |
| `is_sold_out` | ``$BOOLEAN`` |  |
| `is_ticketed` | ``$BOOLEAN`` |  |
| `is_virtual_event` | ``$BOOLEAN`` |  |
| `join_url` | ``$ANY`` |  |
| `layout_type` | ``$ANY`` |  |
| `list_description` | ``$ANY`` |  |
| `location` | ``$ANY`` |  |
| `program_id` | ``$STRING`` |  |
| `program_title` | ``$ANY`` |  |
| `rsvp_link` | ``$ANY`` |  |
| `search_tag` | ``$ANY`` |  |
| `short_description` | ``$ANY`` |  |
| `slug` | ``$STRING`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `start_date` | ``$ANY`` |  |
| `start_time` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `survey_url` | ``$ANY`` |  |
| `ticketed_event_id` | ``$STRING`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `title_display` | ``$ANY`` |  |
| `updated_at` | ``$ANY`` |  |
| `virtual_event_passcode` | ``$ANY`` |  |
| `virtual_event_url` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare Event record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `button_caption` | ``$ANY`` |  |
| `button_text` | ``$ANY`` |  |
| `button_url` | ``$ANY`` |  |
| `description` | ``$STRING`` |  |
| `end_at` | ``$ANY`` |  |
| `event_id` | ``$STRING`` |  |
| `id` | ``$STRING`` |  |
| `image_url` | ``$ANY`` |  |
| `is_private` | ``$BOOLEAN`` |  |
| `is_sales_button_hidden` | ``$BOOLEAN`` |  |
| `is_ticketed` | ``$BOOLEAN`` |  |
| `location` | ``$ANY`` |  |
| `off_sale_at` | ``$ANY`` |  |
| `on_sale_at` | ``$ANY`` |  |
| `short_description` | ``$ANY`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `start_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `title_display` | ``$ANY`` |  |
| `updated_at` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare EventOccurrence record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `is_affiliate_group` | ``$BOOLEAN`` |  |
| `is_event_host` | ``$BOOLEAN`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare EventProgram record (raises on error).
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
| `aic_end_at` | ``$ANY`` |  |
| `aic_start_at` | ``$ANY`` |  |
| `alt_image_id` | ``$STRING`` |  |
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `artist_id` | ``$STRING`` |  |
| `artwork_id` | ``$STRING`` |  |
| `artwork_title` | ``$ANY`` |  |
| `document_id` | ``$STRING`` |  |
| `gallery_id` | ``$STRING`` |  |
| `gallery_title` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `image_id` | ``$STRING`` |  |
| `image_url` | ``$ANY`` |  |
| `is_featured` | ``$BOOLEAN`` |  |
| `is_published` | ``$BOOLEAN`` |  |
| `position` | ``$ANY`` |  |
| `short_description` | ``$ANY`` |  |
| `site_id` | ``$STRING`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `status` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |
| `web_url` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare Exhibition record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `floor` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `is_closed` | ``$BOOLEAN`` |  |
| `latitude` | ``$NUMBER`` |  |
| `latlon` | ``$ANY`` |  |
| `longitude` | ``$NUMBER`` |  |
| `number` | ``$ANY`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `tgn_id` | ``$STRING`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare Gallery record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `copy` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `search_tag` | ``$ANY`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |
| `web_url` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare GenericPage record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `copy` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare Highlight record (raises on error).
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
| `additional_text` | ``$ANY`` |  |
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `friday_is_closed` | ``$ANY`` |  |
| `friday_member_close` | ``$ANY`` |  |
| `friday_member_open` | ``$ANY`` |  |
| `friday_public_close` | ``$ANY`` |  |
| `friday_public_open` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `monday_is_closed` | ``$ANY`` |  |
| `monday_member_close` | ``$ANY`` |  |
| `monday_member_open` | ``$ANY`` |  |
| `monday_public_close` | ``$ANY`` |  |
| `monday_public_open` | ``$ANY`` |  |
| `saturday_is_closed` | ``$ANY`` |  |
| `saturday_member_close` | ``$ANY`` |  |
| `saturday_member_open` | ``$ANY`` |  |
| `saturday_public_close` | ``$ANY`` |  |
| `saturday_public_open` | ``$ANY`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `summary` | ``$ANY`` |  |
| `sunday_is_closed` | ``$ANY`` |  |
| `sunday_member_close` | ``$ANY`` |  |
| `sunday_member_open` | ``$ANY`` |  |
| `sunday_public_close` | ``$ANY`` |  |
| `sunday_public_open` | ``$ANY`` |  |
| `thursday_is_closed` | ``$ANY`` |  |
| `thursday_member_close` | ``$ANY`` |  |
| `thursday_member_open` | ``$ANY`` |  |
| `thursday_public_close` | ``$ANY`` |  |
| `thursday_public_open` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `tuesday_is_closed` | ``$ANY`` |  |
| `tuesday_member_close` | ``$ANY`` |  |
| `tuesday_member_open` | ``$ANY`` |  |
| `tuesday_public_close` | ``$ANY`` |  |
| `tuesday_public_open` | ``$ANY`` |  |
| `updated_at` | ``$ANY`` |  |
| `wednesday_is_closed` | ``$ANY`` |  |
| `wednesday_member_close` | ``$ANY`` |  |
| `wednesday_member_open` | ``$ANY`` |  |
| `wednesday_public_close` | ``$ANY`` |  |
| `wednesday_public_open` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare Hour record (raises on error).
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
| `ahash` | ``$ANY`` |  |
| `alt_text` | ``$ANY`` |  |
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `artwork_id` | ``$STRING`` |  |
| `artwork_title` | ``$ANY`` |  |
| `color` | ``$ANY`` |  |
| `colorfulness` | ``$ANY`` |  |
| `content` | ``$ANY`` |  |
| `content_e_tag` | ``$ANY`` |  |
| `credit_line` | ``$ANY`` |  |
| `fingerprint` | ``$ANY`` |  |
| `height` | ``$NUMBER`` |  |
| `id` | ``$STRING`` |  |
| `iiif_url` | ``$ANY`` |  |
| `is_educational_resource` | ``$BOOLEAN`` |  |
| `is_multimedia_resource` | ``$BOOLEAN`` |  |
| `is_teacher_resource` | ``$BOOLEAN`` |  |
| `lake_guid` | ``$ANY`` |  |
| `lqip` | ``$ANY`` |  |
| `phash` | ``$ANY`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `type` | ``$ANY`` |  |
| `updated_at` | ``$ANY`` |  |
| `width` | ``$NUMBER`` |  |

#### Example: Load

```ruby
# load returns the bare Image record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `copy` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `search_tag` | ``$ANY`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |
| `web_url` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare LandingPage record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `latitude` | ``$NUMBER`` |  |
| `longitude` | ``$NUMBER`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `tgn_id` | ``$STRING`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare Place record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `copy` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |
| `web_url` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare PressRelease record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `copy` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |
| `web_url` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare PrintedPublication record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `artist_id` | ``$STRING`` |  |
| `artwork_id` | ``$STRING`` |  |
| `description` | ``$STRING`` |  |
| `exhibition_id` | ``$STRING`` |  |
| `external_sku` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `image_url` | ``$ANY`` |  |
| `max_compare_at_price` | ``$ANY`` |  |
| `max_current_price` | ``$ANY`` |  |
| `min_compare_at_price` | ``$ANY`` |  |
| `min_current_price` | ``$ANY`` |  |
| `price_display` | ``$ANY`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |
| `web_url` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare Product record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `section_id` | ``$STRING`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |
| `web_url` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare Publication record (raises on error).
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
| `api_id` | ``$STRING`` |  |
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `is_boosted` | ``$BOOLEAN`` |  |
| `score` | ``$NUMBER`` |  |
| `thumbnail` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |

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
| `accession` | ``$ANY`` |  |
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `artwork_id` | ``$STRING`` |  |
| `content` | ``$ANY`` |  |
| `generic_page_id` | ``$STRING`` |  |
| `id` | ``$STRING`` |  |
| `publication_id` | ``$STRING`` |  |
| `publication_title` | ``$ANY`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |
| `web_url` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare Section record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `artwork_id` | ``$STRING`` |  |
| `artwork_title` | ``$ANY`` |  |
| `description` | ``$STRING`` |  |
| `exhibition_id` | ``$STRING`` |  |
| `exhibition_title` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |
| `web_url` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare Site record (raises on error).
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
| `alt_text` | ``$ANY`` |  |
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `artwork_id` | ``$STRING`` |  |
| `artwork_title` | ``$ANY`` |  |
| `content` | ``$ANY`` |  |
| `content_e_tag` | ``$ANY`` |  |
| `credit_line` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `is_educational_resource` | ``$BOOLEAN`` |  |
| `is_multimedia_resource` | ``$BOOLEAN`` |  |
| `is_teacher_resource` | ``$BOOLEAN`` |  |
| `lake_guid` | ``$ANY`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `transcript` | ``$ANY`` |  |
| `type` | ``$ANY`` |  |
| `updated_at` | ``$ANY`` |  |
| `web_url` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare Sound record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |
| `web_url` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare StaticPage record (raises on error).
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
| `alt_text` | ``$ANY`` |  |
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `artwork_id` | ``$STRING`` |  |
| `artwork_title` | ``$ANY`` |  |
| `content` | ``$ANY`` |  |
| `content_e_tag` | ``$ANY`` |  |
| `credit_line` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `is_educational_resource` | ``$BOOLEAN`` |  |
| `is_multimedia_resource` | ``$BOOLEAN`` |  |
| `is_teacher_resource` | ``$BOOLEAN`` |  |
| `lake_guid` | ``$ANY`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `type` | ``$ANY`` |  |
| `updated_at` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare Text record (raises on error).
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
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `artist_title` | ``$ANY`` |  |
| `artwork_title` | ``$ANY`` |  |
| `description` | ``$STRING`` |  |
| `id` | ``$STRING`` |  |
| `image` | ``$ANY`` |  |
| `intro` | ``$ANY`` |  |
| `intro_link` | ``$ANY`` |  |
| `intro_transcript` | ``$ANY`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `updated_at` | ``$ANY`` |  |
| `weight` | ``$NUMBER`` |  |

#### Example: Load

```ruby
# load returns the bare Tour record (raises on error).
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
| `alt_text` | ``$ANY`` |  |
| `api_link` | ``$ANY`` |  |
| `api_model` | ``$ANY`` |  |
| `artwork_id` | ``$STRING`` |  |
| `artwork_title` | ``$ANY`` |  |
| `content` | ``$ANY`` |  |
| `content_e_tag` | ``$ANY`` |  |
| `credit_line` | ``$ANY`` |  |
| `id` | ``$STRING`` |  |
| `is_educational_resource` | ``$BOOLEAN`` |  |
| `is_multimedia_resource` | ``$BOOLEAN`` |  |
| `is_teacher_resource` | ``$BOOLEAN`` |  |
| `lake_guid` | ``$ANY`` |  |
| `source_updated_at` | ``$ANY`` |  |
| `suggest_autocomplete_all` | ``$ANY`` |  |
| `suggest_autocomplete_boosted` | ``$ANY`` |  |
| `timestamp` | ``$ANY`` |  |
| `title` | ``$STRING`` |  |
| `type` | ``$ANY`` |  |
| `updated_at` | ``$ANY`` |  |

#### Example: Load

```ruby
# load returns the bare Video record (raises on error).
video = client.Video.load({ "id" => "video_id" })
```

#### Example: List

```ruby
# list returns an Array of Video records (raises on error).
videos = client.Video.list
```


## Explanation

### The operation pipeline

Every entity operation (load, list, create, update, remove) follows a
six-stage pipeline. Each stage fires a feature hook before executing:

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

If any stage returns an error, the pipeline short-circuits and the
error is returned to the caller as a second return value.

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

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```ruby
agent = client.Agent
agent.load({ "id" => "example_id" })

# agent.data_get now returns the loaded agent data
# agent.match_get returns the last match criteria
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
