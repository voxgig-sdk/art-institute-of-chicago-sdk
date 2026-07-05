# ArtInstituteOfChicago Lua SDK



The Lua SDK for the ArtInstituteOfChicago API — an entity-oriented client using Lua conventions.

It exposes the API as capitalised, semantic **Entities** — e.g. `client:Agent()` — each with the same small set of operations (`list`, `load`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to LuaRocks. Install it from the
GitHub release tag (`lua/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/art-institute-of-chicago-sdk/releases)),
or add the source directory to your `LUA_PATH`:

```bash
export LUA_PATH="path/to/lua/?.lua;path/to/lua/?/init.lua;;"
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```lua
local sdk = require("art-institute-of-chicago_sdk")

local client = sdk.new()
```

### 2. List agent records

Entity operations return `(value, err)`. For `list`, `value` is the
array of records itself — iterate it directly (there is no wrapper).

```lua
local agents, err = client:Agent():list()
if err then error(err) end

for _, item in ipairs(agents) do
  print(item["id"], item["description"])
end
```

### 3. Load an agent

```lua
local agent, err = client:Agent():load({ id = "example_id" })
if err then error(err) end
print(agent)
```


## Error handling

Entity operations return `(value, err)`. Check `err` before using
the value:

```lua
local agents, err = client:Agent():list()
if err then error(err) end
```

`direct` follows the same `(value, err)` convention:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example_id" },
})
if err then error(err) end
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
if err then error(err) end

if result["ok"] then
  print(result["status"])  -- 200
  print(result["data"])    -- response body
end
```

### Prepare a request without sending it

```lua
local fetchdef, err = client:prepare({
  path = "/api/resource/{id}",
  method = "DELETE",
  params = { id = "example" },
})
if err then error(err) end

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```lua
local client = sdk.test()

local result, err = client:Agent():list()
-- result is the returned data; err is set on failure
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```lua
local function mock_fetch(url, init)
  return {
    status = 200,
    statusText = "OK",
    headers = {},
    json = function()
      return { id = "mock01" }
    end,
  }, nil
end

local client = sdk.new({
  base = "http://localhost:8080",
  system = {
    fetch = mock_fetch,
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
cd lua && busted test/
```


## Reference

### ArtInstituteOfChicagoSDK

```lua
local sdk = require("art-institute-of-chicago_sdk")
local client = sdk.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `table` | Feature activation flags. |
| `extend` | `table` | Additional Feature instances to load. |
| `system` | `table` | System overrides (e.g. custom `fetch` function). |

### test

```lua
local client = sdk.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### ArtInstituteOfChicagoSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> table` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> table, err` | Build an HTTP request definition without sending. |
| `direct` | `(fetchargs) -> table, err` | Build and send an HTTP request. |
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
| `load` | `(reqmatch, ctrl) -> any, err` | Load a single entity by match criteria. |
| `list` | `(reqmatch, ctrl) -> any, err` | List entities matching the criteria. |
| `data_get` | `() -> table` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> table` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> string` | Return the entity name. |

### Result shape

Entity operations return `(value, err)`. The `value` is the operation's
data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `load` | the entity record (a `table`) |
| `list` | an array (`table`) of entity records |

Check `err` first (it is non-`nil` on failure), then use `value`:

    local agent, err = client:Agent():load({ id = "example_id" })
    if err then error(err) end
    -- agent is the loaded record

Only `direct()` returns a response envelope — a `table` with `ok`,
`status`, `headers`, and `data` keys.

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

Create an instance: `local agent = client:Agent(nil)`

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

```lua
local agent, err = client:Agent():load({ id = "agent_id" })
```

#### Example: List

```lua
local agents, err = client:Agent():list()
```


### AgentRole

Create an instance: `local agent_role = client:AgentRole(nil)`

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

```lua
local agent_role, err = client:AgentRole():load({ id = "agent_role_id" })
```

#### Example: List

```lua
local agent_roles, err = client:AgentRole():list()
```


### AgentType

Create an instance: `local agent_type = client:AgentType(nil)`

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

```lua
local agent_type, err = client:AgentType():load({ id = "agent_type_id" })
```

#### Example: List

```lua
local agent_types, err = client:AgentType():list()
```


### Article

Create an instance: `local article = client:Article(nil)`

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

```lua
local article, err = client:Article():load({ id = "article_id" })
```

#### Example: List

```lua
local articles, err = client:Article():list()
```


### Artwork

Create an instance: `local artwork = client:Artwork(nil)`

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

```lua
local artwork, err = client:Artwork():load({ id = "artwork_id" })
```

#### Example: List

```lua
local artworks, err = client:Artwork():list()
```


### ArtworkDateQualifier

Create an instance: `local artwork_date_qualifier = client:ArtworkDateQualifier(nil)`

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

```lua
local artwork_date_qualifier, err = client:ArtworkDateQualifier():load({ id = "artwork_date_qualifier_id" })
```

#### Example: List

```lua
local artwork_date_qualifiers, err = client:ArtworkDateQualifier():list()
```


### ArtworkPlaceQualifier

Create an instance: `local artwork_place_qualifier = client:ArtworkPlaceQualifier(nil)`

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

```lua
local artwork_place_qualifier, err = client:ArtworkPlaceQualifier():load({ id = "artwork_place_qualifier_id" })
```

#### Example: List

```lua
local artwork_place_qualifiers, err = client:ArtworkPlaceQualifier():list()
```


### ArtworkType

Create an instance: `local artwork_type = client:ArtworkType(nil)`

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

```lua
local artwork_type, err = client:ArtworkType():load({ id = "artwork_type_id" })
```

#### Example: List

```lua
local artwork_types, err = client:ArtworkType():list()
```


### CategoryTerm

Create an instance: `local category_term = client:CategoryTerm(nil)`

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

```lua
local category_term, err = client:CategoryTerm():load({ id = "category_term_id" })
```

#### Example: List

```lua
local category_terms, err = client:CategoryTerm():list()
```


### DigitalPublication

Create an instance: `local digital_publication = client:DigitalPublication(nil)`

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

```lua
local digital_publication, err = client:DigitalPublication():load({ id = "digital_publication_id" })
```

#### Example: List

```lua
local digital_publications, err = client:DigitalPublication():list()
```


### DigitalPublicationArticle

Create an instance: `local digital_publication_article = client:DigitalPublicationArticle(nil)`

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

```lua
local digital_publication_article, err = client:DigitalPublicationArticle():load({ id = "digital_publication_article_id" })
```

#### Example: List

```lua
local digital_publication_articles, err = client:DigitalPublicationArticle():list()
```


### EducatorResource

Create an instance: `local educator_resource = client:EducatorResource(nil)`

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

```lua
local educator_resource, err = client:EducatorResource():load({ id = "educator_resource_id" })
```

#### Example: List

```lua
local educator_resources, err = client:EducatorResource():list()
```


### Event

Create an instance: `local event = client:Event(nil)`

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

```lua
local event, err = client:Event():load({ id = "event_id" })
```

#### Example: List

```lua
local events, err = client:Event():list()
```


### EventOccurrence

Create an instance: `local event_occurrence = client:EventOccurrence(nil)`

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

```lua
local event_occurrence, err = client:EventOccurrence():load({ id = "event_occurrence_id" })
```

#### Example: List

```lua
local event_occurrences, err = client:EventOccurrence():list()
```


### EventProgram

Create an instance: `local event_program = client:EventProgram(nil)`

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

```lua
local event_program, err = client:EventProgram():load({ id = "event_program_id" })
```

#### Example: List

```lua
local event_programs, err = client:EventProgram():list()
```


### Exhibition

Create an instance: `local exhibition = client:Exhibition(nil)`

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

```lua
local exhibition, err = client:Exhibition():load({ id = "exhibition_id" })
```

#### Example: List

```lua
local exhibitions, err = client:Exhibition():list()
```


### Gallery

Create an instance: `local gallery = client:Gallery(nil)`

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

```lua
local gallery, err = client:Gallery():load({ id = "gallery_id" })
```

#### Example: List

```lua
local gallerys, err = client:Gallery():list()
```


### GenericPage

Create an instance: `local generic_page = client:GenericPage(nil)`

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

```lua
local generic_page, err = client:GenericPage():load({ id = "generic_page_id" })
```

#### Example: List

```lua
local generic_pages, err = client:GenericPage():list()
```


### Highlight

Create an instance: `local highlight = client:Highlight(nil)`

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

```lua
local highlight, err = client:Highlight():load({ id = "highlight_id" })
```

#### Example: List

```lua
local highlights, err = client:Highlight():list()
```


### Hour

Create an instance: `local hour = client:Hour(nil)`

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

```lua
local hour, err = client:Hour():load({ id = "hour_id" })
```

#### Example: List

```lua
local hours, err = client:Hour():list()
```


### Image

Create an instance: `local image = client:Image(nil)`

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

```lua
local image, err = client:Image():load({ id = "image_id" })
```

#### Example: List

```lua
local images, err = client:Image():list()
```


### LandingPage

Create an instance: `local landing_page = client:LandingPage(nil)`

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

```lua
local landing_page, err = client:LandingPage():load({ id = "landing_page_id" })
```

#### Example: List

```lua
local landing_pages, err = client:LandingPage():list()
```


### Place

Create an instance: `local place = client:Place(nil)`

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

```lua
local place, err = client:Place():load({ id = "place_id" })
```

#### Example: List

```lua
local places, err = client:Place():list()
```


### PressRelease

Create an instance: `local press_release = client:PressRelease(nil)`

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

```lua
local press_release, err = client:PressRelease():load({ id = "press_release_id" })
```

#### Example: List

```lua
local press_releases, err = client:PressRelease():list()
```


### PrintedPublication

Create an instance: `local printed_publication = client:PrintedPublication(nil)`

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

```lua
local printed_publication, err = client:PrintedPublication():load({ id = "printed_publication_id" })
```

#### Example: List

```lua
local printed_publications, err = client:PrintedPublication():list()
```


### Product

Create an instance: `local product = client:Product(nil)`

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

```lua
local product, err = client:Product():load({ id = "product_id" })
```

#### Example: List

```lua
local products, err = client:Product():list()
```


### Publication

Create an instance: `local publication = client:Publication(nil)`

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

```lua
local publication, err = client:Publication():load({ id = "publication_id" })
```

#### Example: List

```lua
local publications, err = client:Publication():list()
```


### Search

Create an instance: `local search = client:Search(nil)`

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

```lua
local searchs, err = client:Search():list()
```


### Section

Create an instance: `local section = client:Section(nil)`

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

```lua
local section, err = client:Section():load({ id = "section_id" })
```

#### Example: List

```lua
local sections, err = client:Section():list()
```


### Site

Create an instance: `local site = client:Site(nil)`

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

```lua
local site, err = client:Site():load({ id = "site_id" })
```

#### Example: List

```lua
local sites, err = client:Site():list()
```


### Sound

Create an instance: `local sound = client:Sound(nil)`

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

```lua
local sound, err = client:Sound():load({ id = "sound_id" })
```

#### Example: List

```lua
local sounds, err = client:Sound():list()
```


### StaticPage

Create an instance: `local static_page = client:StaticPage(nil)`

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

```lua
local static_page, err = client:StaticPage():load({ id = "static_page_id" })
```

#### Example: List

```lua
local static_pages, err = client:StaticPage():list()
```


### Text

Create an instance: `local text = client:Text(nil)`

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

```lua
local text, err = client:Text():load({ id = "text_id" })
```

#### Example: List

```lua
local texts, err = client:Text():list()
```


### Tour

Create an instance: `local tour = client:Tour(nil)`

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

```lua
local tour, err = client:Tour():load({ id = "tour_id" })
```

#### Example: List

```lua
local tours, err = client:Tour():list()
```


### Video

Create an instance: `local video = client:Video(nil)`

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

```lua
local video, err = client:Video():load({ id = "video_id" })
```

#### Example: List

```lua
local videos, err = client:Video():list()
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

Features are the extension mechanism. A feature is a Lua table
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as tables

The Lua SDK uses plain Lua tables throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a table.

### Module structure

```
lua/
├── art-institute-of-chicago_sdk.lua    -- Main SDK module
├── config.lua               -- Configuration
├── features.lua             -- Feature factory
├── core/                    -- Core types and context
├── entity/                  -- Entity implementations
├── feature/                 -- Built-in features (Base, Test, Log)
├── utility/                 -- Utility functions and struct library
└── test/                    -- Test suites
```

The main module (`art-institute-of-chicago_sdk`) exports the SDK constructor
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```lua
local agent = client:Agent()
agent:list()

-- agent:data_get() now returns the agent data from the last list
-- agent:match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
