# ArtInstituteOfChicago Python SDK



The Python SDK for the ArtInstituteOfChicago API — an entity-oriented client following Pythonic conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.Agent()` — each
carrying a small, uniform set of operations (`list`, `load`) instead of raw URL
paths and query strings. You work with named resources and verbs, which
keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/art-institute-of-chicago-sdk/releases)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
from artinstituteofchicago_sdk import ArtInstituteOfChicagoSDK

client = ArtInstituteOfChicagoSDK()
```

### 2. List agent records

`list()` returns a `list` of records (each a `dict`) and raises on
error — iterate it directly.

```python
try:
    agents = client.Agent().list()
    for agent in agents:
        print(agent)
except Exception as err:
    print(f"list failed: {err}")
```

### 3. Load an agent

`load()` returns the ENTITY — call data_get() for the record — and raises on error.

```python
try:
    agent = client.Agent().load({"id": "example_id"})
    print(agent)
except Exception as err:
    print(f"load failed: {err}")
```


## Error handling

Entity operations raise on failure, so wrap them in `try` / `except`:

```python
try:
    publications = client.Publication().list()
    print(publications)
except Exception as err:
    print(f"list failed: {err}")
```

`direct()` does **not** raise — it returns the result envelope. Branch
on `ok`; on failure `status` holds the HTTP status (for error responses)
and `err` holds a transport error, so read both defensively:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example_id"},
})

if not result["ok"]:
    print("request failed:", result.get("status"), result.get("err"))
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    # A non-2xx response carries status + data (the error body); a
    # transport-level failure carries err instead. Only one is present, so
    # read both with .get() rather than indexing a key that may be absent.
    print(result.get("status"), result.get("err"))
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = ArtInstituteOfChicagoSDK.test()

# Entity ops return the ENTITY and raises on error;
# call data_get() for the record.
publication = client.Publication().list()
# publication contains the mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = ArtInstituteOfChicagoSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
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
cd py && pytest test/
```


## Reference

### ArtInstituteOfChicagoSDK

```python
from artinstituteofchicago_sdk import ArtInstituteOfChicagoSDK

client = ArtInstituteOfChicagoSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = ArtInstituteOfChicagoSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### ArtInstituteOfChicagoSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
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
| `list` | `(reqmatch, ctrl) -> list` | List entities matching the criteria. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (a `dict` for single-entity
ops, a `list` for `list`) and raise on error. Wrap calls in
`try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

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

Create an instance: `agent = client.Agent()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_titles` | `Any` |  |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `birth_date` | `Any` |  |
| `death_date` | `Any` |  |
| `description` | `str` |  |
| `id` | `str` |  |
| `is_artist` | `bool` |  |
| `sort_title` | `Any` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `ulan_id` | `str` |  |
| `updated_at` | `Any` |  |

#### Example: Load

```python
agent = client.Agent().load({"id": "agent_id"})
```

#### Example: List

```python
agents = client.Agent().list()
```


### AgentRole

Create an instance: `agent_role = client.AgentRole()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `id` | `str` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |

#### Example: Load

```python
agent_role = client.AgentRole().load({"id": "agent_role_id"})
```

#### Example: List

```python
agent_roles = client.AgentRole().list()
```


### AgentType

Create an instance: `agent_type = client.AgentType()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `id` | `str` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |

#### Example: Load

```python
agent_type = client.AgentType().load({"id": "agent_type_id"})
```

#### Example: List

```python
agent_types = client.AgentType().list()
```


### Article

Create an instance: `article = client.Article()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `copy` | `Any` |  |
| `id` | `str` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |

#### Example: Load

```python
article = client.Article().load({"id": "article_id"})
```

#### Example: List

```python
articles = client.Article().list()
```


### Artwork

Create an instance: `artwork = client.Artwork()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_artist_ids` | `Any` |  |
| `alt_classification_ids` | `Any` |  |
| `alt_image_ids` | `Any` |  |
| `alt_material_ids` | `Any` |  |
| `alt_style_ids` | `Any` |  |
| `alt_subject_ids` | `Any` |  |
| `alt_technique_ids` | `Any` |  |
| `alt_titles` | `Any` |  |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `artist_display` | `Any` |  |
| `artist_id` | `str` |  |
| `artist_ids` | `Any` |  |
| `artist_title` | `Any` |  |
| `artist_titles` | `Any` |  |
| `artwork_type_id` | `str` |  |
| `artwork_type_title` | `Any` |  |
| `boost_rank` | `Any` |  |
| `catalog_based_search_keyword_titles` | `Any` |  |
| `catalogue_display` | `Any` |  |
| `category_ids` | `Any` |  |
| `category_titles` | `Any` |  |
| `classification_id` | `str` |  |
| `classification_ids` | `Any` |  |
| `classification_title` | `Any` |  |
| `classification_titles` | `Any` |  |
| `color` | `Any` |  |
| `colorfulness` | `Any` |  |
| `copyright_notice` | `Any` |  |
| `credit_line` | `Any` |  |
| `date_display` | `Any` |  |
| `date_end` | `Any` |  |
| `date_qualifier_id` | `str` |  |
| `date_qualifier_title` | `Any` |  |
| `date_start` | `Any` |  |
| `department_id` | `str` |  |
| `department_title` | `Any` |  |
| `description` | `str` |  |
| `dimensions` | `Any` |  |
| `dimensions_detail` | `Any` |  |
| `document_ids` | `Any` |  |
| `edition` | `Any` |  |
| `exhibition_history` | `Any` |  |
| `fiscal_year` | `Any` |  |
| `fiscal_year_deaccession` | `Any` |  |
| `gallery_id` | `str` |  |
| `gallery_title` | `Any` |  |
| `has_advanced_imaging` | `bool` |  |
| `has_educational_resources` | `bool` |  |
| `has_multimedia_resources` | `bool` |  |
| `has_not_been_viewed_much` | `bool` |  |
| `id` | `str` |  |
| `image_embedding` | `Any` |  |
| `image_id` | `str` |  |
| `inscriptions` | `Any` |  |
| `internal_department_id` | `str` |  |
| `is_boosted` | `bool` |  |
| `is_on_view` | `bool` |  |
| `is_public_domain` | `bool` |  |
| `is_zoomable` | `bool` |  |
| `latitude` | `float` |  |
| `latlon` | `Any` |  |
| `longitude` | `float` |  |
| `main_reference_number` | `int` |  |
| `material_id` | `str` |  |
| `material_ids` | `Any` |  |
| `material_titles` | `Any` |  |
| `max_zoom_window_size` | `Any` |  |
| `medium_display` | `Any` |  |
| `nomisma_id` | `str` |  |
| `on_loan_display` | `Any` |  |
| `pageviews` | `Any` |  |
| `pageviews_recent` | `Any` |  |
| `place_of_origin` | `Any` |  |
| `provenance_text` | `Any` |  |
| `publication_history` | `Any` |  |
| `publishing_verification_level` | `Any` |  |
| `section_ids` | `Any` |  |
| `section_titles` | `Any` |  |
| `short_description` | `Any` |  |
| `site_ids` | `Any` |  |
| `sound_ids` | `Any` |  |
| `source_updated_at` | `Any` |  |
| `style_id` | `str` |  |
| `style_ids` | `Any` |  |
| `style_title` | `Any` |  |
| `style_titles` | `Any` |  |
| `subject_id` | `str` |  |
| `subject_ids` | `Any` |  |
| `subject_titles` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `technique_id` | `str` |  |
| `technique_ids` | `Any` |  |
| `technique_titles` | `Any` |  |
| `term_titles` | `Any` |  |
| `text_embedding` | `Any` |  |
| `text_ids` | `Any` |  |
| `theme_titles` | `Any` |  |
| `thumbnail` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |
| `video_ids` | `Any` |  |

#### Example: Load

```python
artwork = client.Artwork().load({"id": "artwork_id"})
```

#### Example: List

```python
artworks = client.Artwork().list()
```


### ArtworkDateQualifier

Create an instance: `artwork_date_qualifier = client.ArtworkDateQualifier()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `id` | `str` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |

#### Example: Load

```python
artwork_date_qualifier = client.ArtworkDateQualifier().load({"id": "artwork_date_qualifier_id"})
```

#### Example: List

```python
artwork_date_qualifiers = client.ArtworkDateQualifier().list()
```


### ArtworkPlaceQualifier

Create an instance: `artwork_place_qualifier = client.ArtworkPlaceQualifier()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `id` | `str` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |

#### Example: Load

```python
artwork_place_qualifier = client.ArtworkPlaceQualifier().load({"id": "artwork_place_qualifier_id"})
```

#### Example: List

```python
artwork_place_qualifiers = client.ArtworkPlaceQualifier().list()
```


### ArtworkType

Create an instance: `artwork_type = client.ArtworkType()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `aat_id` | `str` |  |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `id` | `str` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |

#### Example: Load

```python
artwork_type = client.ArtworkType().load({"id": "artwork_type_id"})
```

#### Example: List

```python
artwork_types = client.ArtworkType().list()
```


### CategoryTerm

Create an instance: `category_term = client.CategoryTerm()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `aat_id` | `str` |  |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `id` | `str` |  |
| `parent_id` | `str` |  |
| `source_updated_at` | `Any` |  |
| `subtype` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |

#### Example: Load

```python
category_term = client.CategoryTerm().load({"id": "category_term_id"})
```

#### Example: List

```python
category_terms = client.CategoryTerm().list()
```


### DigitalPublication

Create an instance: `digital_publication = client.DigitalPublication()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `copy` | `Any` |  |
| `id` | `str` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |
| `web_url` | `Any` |  |

#### Example: Load

```python
digital_publication = client.DigitalPublication().load({"id": "digital_publication_id"})
```

#### Example: List

```python
digital_publications = client.DigitalPublication().list()
```


### DigitalPublicationArticle

Create an instance: `digital_publication_article = client.DigitalPublicationArticle()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `author_display` | `Any` |  |
| `copy` | `Any` |  |
| `digital_publication_id` | `str` |  |
| `id` | `str` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |
| `web_url` | `Any` |  |

#### Example: Load

```python
digital_publication_article = client.DigitalPublicationArticle().load({"id": "digital_publication_article_id"})
```

#### Example: List

```python
digital_publication_articles = client.DigitalPublicationArticle().list()
```


### EducatorResource

Create an instance: `educator_resource = client.EducatorResource()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `copy` | `Any` |  |
| `id` | `str` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |
| `web_url` | `Any` |  |

#### Example: Load

```python
educator_resource = client.EducatorResource().load({"id": "educator_resource_id"})
```

#### Example: List

```python
educator_resources = client.EducatorResource().list()
```


### Event

Create an instance: `event = client.Event()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_audience_ids` | `Any` |  |
| `alt_event_type_ids` | `Any` |  |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `audience_id` | `str` |  |
| `buy_button_caption` | `Any` |  |
| `buy_button_text` | `Any` |  |
| `date_display` | `Any` |  |
| `description` | `str` |  |
| `door_time` | `Any` |  |
| `end_date` | `Any` |  |
| `end_time` | `Any` |  |
| `entrance` | `Any` |  |
| `event_host_id` | `str` |  |
| `event_host_title` | `Any` |  |
| `event_type_id` | `str` |  |
| `header_description` | `Any` |  |
| `hero_caption` | `Any` |  |
| `id` | `str` |  |
| `image_url` | `Any` |  |
| `is_admission_required` | `bool` |  |
| `is_after_hours` | `bool` |  |
| `is_free` | `bool` |  |
| `is_member_exclusive` | `bool` |  |
| `is_private` | `bool` |  |
| `is_registration_required` | `bool` |  |
| `is_sales_button_hidden` | `bool` |  |
| `is_sold_out` | `bool` |  |
| `is_ticketed` | `bool` |  |
| `is_virtual_event` | `bool` |  |
| `join_url` | `Any` |  |
| `layout_type` | `Any` |  |
| `list_description` | `Any` |  |
| `location` | `Any` |  |
| `program_ids` | `Any` |  |
| `program_titles` | `Any` |  |
| `rsvp_link` | `Any` |  |
| `search_tags` | `Any` |  |
| `short_description` | `Any` |  |
| `slug` | `str` |  |
| `source_updated_at` | `Any` |  |
| `start_date` | `Any` |  |
| `start_time` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `survey_url` | `Any` |  |
| `ticketed_event_id` | `str` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `title_display` | `Any` |  |
| `updated_at` | `Any` |  |
| `virtual_event_passcode` | `Any` |  |
| `virtual_event_url` | `Any` |  |

#### Example: Load

```python
event = client.Event().load({"id": "event_id"})
```

#### Example: List

```python
events = client.Event().list()
```


### EventOccurrence

Create an instance: `event_occurrence = client.EventOccurrence()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `button_caption` | `Any` |  |
| `button_text` | `Any` |  |
| `button_url` | `Any` |  |
| `description` | `str` |  |
| `end_at` | `Any` |  |
| `event_id` | `str` |  |
| `id` | `str` |  |
| `image_url` | `Any` |  |
| `is_private` | `bool` |  |
| `is_sales_button_hidden` | `bool` |  |
| `is_ticketed` | `bool` |  |
| `location` | `Any` |  |
| `off_sale_at` | `Any` |  |
| `on_sale_at` | `Any` |  |
| `short_description` | `Any` |  |
| `source_updated_at` | `Any` |  |
| `start_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `title_display` | `Any` |  |
| `updated_at` | `Any` |  |

#### Example: Load

```python
event_occurrence = client.EventOccurrence().load({"id": "event_occurrence_id"})
```

#### Example: List

```python
event_occurrences = client.EventOccurrence().list()
```


### EventProgram

Create an instance: `event_program = client.EventProgram()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `id` | `str` |  |
| `is_affiliate_group` | `bool` |  |
| `is_event_host` | `bool` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |

#### Example: Load

```python
event_program = client.EventProgram().load({"id": "event_program_id"})
```

#### Example: List

```python
event_programs = client.EventProgram().list()
```


### Exhibition

Create an instance: `exhibition = client.Exhibition()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `aic_end_at` | `Any` |  |
| `aic_start_at` | `Any` |  |
| `alt_image_ids` | `Any` |  |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `artist_ids` | `Any` |  |
| `artwork_ids` | `Any` |  |
| `artwork_titles` | `Any` |  |
| `document_ids` | `Any` |  |
| `gallery_id` | `str` |  |
| `gallery_title` | `Any` |  |
| `id` | `str` |  |
| `image_id` | `str` |  |
| `image_url` | `Any` |  |
| `is_featured` | `bool` |  |
| `is_published` | `bool` |  |
| `position` | `Any` |  |
| `short_description` | `Any` |  |
| `site_ids` | `Any` |  |
| `source_updated_at` | `Any` |  |
| `status` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |
| `web_url` | `Any` |  |

#### Example: Load

```python
exhibition = client.Exhibition().load({"id": "exhibition_id"})
```

#### Example: List

```python
exhibitions = client.Exhibition().list()
```


### Gallery

Create an instance: `gallery = client.Gallery()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `floor` | `Any` |  |
| `id` | `str` |  |
| `is_closed` | `bool` |  |
| `latitude` | `float` |  |
| `latlon` | `Any` |  |
| `longitude` | `float` |  |
| `number` | `Any` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `tgn_id` | `str` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |

#### Example: Load

```python
gallery = client.Gallery().load({"id": "gallery_id"})
```

#### Example: List

```python
gallerys = client.Gallery().list()
```


### GenericPage

Create an instance: `generic_page = client.GenericPage()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `copy` | `Any` |  |
| `id` | `str` |  |
| `search_tags` | `Any` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |
| `web_url` | `Any` |  |

#### Example: Load

```python
generic_page = client.GenericPage().load({"id": "generic_page_id"})
```

#### Example: List

```python
generic_pages = client.GenericPage().list()
```


### Highlight

Create an instance: `highlight = client.Highlight()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `copy` | `Any` |  |
| `id` | `str` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |

#### Example: Load

```python
highlight = client.Highlight().load({"id": "highlight_id"})
```

#### Example: List

```python
highlights = client.Highlight().list()
```


### Hour

Create an instance: `hour = client.Hour()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `additional_text` | `Any` |  |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `friday_is_closed` | `Any` |  |
| `friday_member_close` | `Any` |  |
| `friday_member_open` | `Any` |  |
| `friday_public_close` | `Any` |  |
| `friday_public_open` | `Any` |  |
| `id` | `str` |  |
| `monday_is_closed` | `Any` |  |
| `monday_member_close` | `Any` |  |
| `monday_member_open` | `Any` |  |
| `monday_public_close` | `Any` |  |
| `monday_public_open` | `Any` |  |
| `saturday_is_closed` | `Any` |  |
| `saturday_member_close` | `Any` |  |
| `saturday_member_open` | `Any` |  |
| `saturday_public_close` | `Any` |  |
| `saturday_public_open` | `Any` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `summary` | `Any` |  |
| `sunday_is_closed` | `Any` |  |
| `sunday_member_close` | `Any` |  |
| `sunday_member_open` | `Any` |  |
| `sunday_public_close` | `Any` |  |
| `sunday_public_open` | `Any` |  |
| `thursday_is_closed` | `Any` |  |
| `thursday_member_close` | `Any` |  |
| `thursday_member_open` | `Any` |  |
| `thursday_public_close` | `Any` |  |
| `thursday_public_open` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `tuesday_is_closed` | `Any` |  |
| `tuesday_member_close` | `Any` |  |
| `tuesday_member_open` | `Any` |  |
| `tuesday_public_close` | `Any` |  |
| `tuesday_public_open` | `Any` |  |
| `updated_at` | `Any` |  |
| `wednesday_is_closed` | `Any` |  |
| `wednesday_member_close` | `Any` |  |
| `wednesday_member_open` | `Any` |  |
| `wednesday_public_close` | `Any` |  |
| `wednesday_public_open` | `Any` |  |

#### Example: Load

```python
hour = client.Hour().load({"id": "hour_id"})
```

#### Example: List

```python
hours = client.Hour().list()
```


### Image

Create an instance: `image = client.Image()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ahash` | `Any` |  |
| `alt_text` | `Any` |  |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `artwork_ids` | `Any` |  |
| `artwork_titles` | `Any` |  |
| `color` | `Any` |  |
| `colorfulness` | `Any` |  |
| `content` | `Any` |  |
| `content_e_tag` | `Any` |  |
| `credit_line` | `Any` |  |
| `fingerprint` | `Any` |  |
| `height` | `float` |  |
| `id` | `str` |  |
| `iiif_url` | `Any` |  |
| `is_educational_resource` | `bool` |  |
| `is_multimedia_resource` | `bool` |  |
| `is_teacher_resource` | `bool` |  |
| `lake_guid` | `Any` |  |
| `lqip` | `Any` |  |
| `phash` | `Any` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `type` | `Any` |  |
| `updated_at` | `Any` |  |
| `width` | `float` |  |

#### Example: Load

```python
image = client.Image().load({"id": "image_id"})
```

#### Example: List

```python
images = client.Image().list()
```


### LandingPage

Create an instance: `landing_page = client.LandingPage()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `copy` | `Any` |  |
| `id` | `str` |  |
| `search_tags` | `Any` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |
| `web_url` | `Any` |  |

#### Example: Load

```python
landing_page = client.LandingPage().load({"id": "landing_page_id"})
```

#### Example: List

```python
landing_pages = client.LandingPage().list()
```


### Place

Create an instance: `place = client.Place()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `id` | `str` |  |
| `latitude` | `float` |  |
| `longitude` | `float` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `tgn_id` | `str` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |

#### Example: Load

```python
place = client.Place().load({"id": "place_id"})
```

#### Example: List

```python
places = client.Place().list()
```


### PressRelease

Create an instance: `press_release = client.PressRelease()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `copy` | `Any` |  |
| `id` | `str` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |
| `web_url` | `Any` |  |

#### Example: Load

```python
press_release = client.PressRelease().load({"id": "press_release_id"})
```

#### Example: List

```python
press_releases = client.PressRelease().list()
```


### PrintedPublication

Create an instance: `printed_publication = client.PrintedPublication()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `copy` | `Any` |  |
| `id` | `str` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |
| `web_url` | `Any` |  |

#### Example: Load

```python
printed_publication = client.PrintedPublication().load({"id": "printed_publication_id"})
```

#### Example: List

```python
printed_publications = client.PrintedPublication().list()
```


### Product

Create an instance: `product = client.Product()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `artist_ids` | `Any` |  |
| `artwork_ids` | `Any` |  |
| `description` | `str` |  |
| `exhibition_ids` | `Any` |  |
| `external_sku` | `Any` |  |
| `id` | `str` |  |
| `image_url` | `Any` |  |
| `max_compare_at_price` | `Any` |  |
| `max_current_price` | `Any` |  |
| `min_compare_at_price` | `Any` |  |
| `min_current_price` | `Any` |  |
| `price_display` | `Any` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |
| `web_url` | `Any` |  |

#### Example: Load

```python
product = client.Product().load({"id": "product_id"})
```

#### Example: List

```python
products = client.Product().list()
```


### Publication

Create an instance: `publication = client.Publication()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `id` | `str` |  |
| `section_ids` | `Any` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |
| `web_url` | `Any` |  |

#### Example: Load

```python
publication = client.Publication().load({"id": "publication_id"})
```

#### Example: List

```python
publications = client.Publication().list()
```


### Search

Create an instance: `search = client.Search()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_id` | `str` |  |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `id` | `str` |  |
| `is_boosted` | `bool` |  |
| `score` | `float` |  |
| `thumbnail` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |

#### Example: List

```python
searchs = client.Search().list()
```


### Section

Create an instance: `section = client.Section()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accession` | `Any` |  |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `artwork_id` | `str` |  |
| `content` | `Any` |  |
| `generic_page_id` | `str` |  |
| `id` | `str` |  |
| `publication_id` | `str` |  |
| `publication_title` | `Any` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |
| `web_url` | `Any` |  |

#### Example: Load

```python
section = client.Section().load({"id": "section_id"})
```

#### Example: List

```python
sections = client.Section().list()
```


### Site

Create an instance: `site = client.Site()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `artwork_ids` | `Any` |  |
| `artwork_titles` | `Any` |  |
| `description` | `str` |  |
| `exhibition_ids` | `Any` |  |
| `exhibition_titles` | `Any` |  |
| `id` | `str` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |
| `web_url` | `Any` |  |

#### Example: Load

```python
site = client.Site().load({"id": "site_id"})
```

#### Example: List

```python
sites = client.Site().list()
```


### Sound

Create an instance: `sound = client.Sound()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_text` | `Any` |  |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `artwork_ids` | `Any` |  |
| `artwork_titles` | `Any` |  |
| `content` | `Any` |  |
| `content_e_tag` | `Any` |  |
| `credit_line` | `Any` |  |
| `id` | `str` |  |
| `is_educational_resource` | `bool` |  |
| `is_multimedia_resource` | `bool` |  |
| `is_teacher_resource` | `bool` |  |
| `lake_guid` | `Any` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `transcript` | `Any` |  |
| `type` | `Any` |  |
| `updated_at` | `Any` |  |
| `web_url` | `Any` |  |

#### Example: Load

```python
sound = client.Sound().load({"id": "sound_id"})
```

#### Example: List

```python
sounds = client.Sound().list()
```


### StaticPage

Create an instance: `static_page = client.StaticPage()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `id` | `str` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |
| `web_url` | `Any` |  |

#### Example: Load

```python
static_page = client.StaticPage().load({"id": "static_page_id"})
```

#### Example: List

```python
static_pages = client.StaticPage().list()
```


### Text

Create an instance: `text = client.Text()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_text` | `Any` |  |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `artwork_ids` | `Any` |  |
| `artwork_titles` | `Any` |  |
| `content` | `Any` |  |
| `content_e_tag` | `Any` |  |
| `credit_line` | `Any` |  |
| `id` | `str` |  |
| `is_educational_resource` | `bool` |  |
| `is_multimedia_resource` | `bool` |  |
| `is_teacher_resource` | `bool` |  |
| `lake_guid` | `Any` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `type` | `Any` |  |
| `updated_at` | `Any` |  |

#### Example: Load

```python
text = client.Text().load({"id": "text_id"})
```

#### Example: List

```python
texts = client.Text().list()
```


### Tour

Create an instance: `tour = client.Tour()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `artist_titles` | `Any` |  |
| `artwork_titles` | `Any` |  |
| `description` | `str` |  |
| `id` | `str` |  |
| `image` | `Any` |  |
| `intro` | `Any` |  |
| `intro_link` | `Any` |  |
| `intro_transcript` | `Any` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `updated_at` | `Any` |  |
| `weight` | `float` |  |

#### Example: Load

```python
tour = client.Tour().load({"id": "tour_id"})
```

#### Example: List

```python
tours = client.Tour().list()
```


### Video

Create an instance: `video = client.Video()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_text` | `Any` |  |
| `api_link` | `Any` |  |
| `api_model` | `Any` |  |
| `artwork_ids` | `Any` |  |
| `artwork_titles` | `Any` |  |
| `content` | `Any` |  |
| `content_e_tag` | `Any` |  |
| `credit_line` | `Any` |  |
| `id` | `str` |  |
| `is_educational_resource` | `bool` |  |
| `is_multimedia_resource` | `bool` |  |
| `is_teacher_resource` | `bool` |  |
| `lake_guid` | `Any` |  |
| `source_updated_at` | `Any` |  |
| `suggest_autocomplete_all` | `Any` |  |
| `suggest_autocomplete_boosted` | `Any` |  |
| `timestamp` | `Any` |  |
| `title` | `str` |  |
| `type` | `Any` |  |
| `updated_at` | `Any` |  |

#### Example: Load

```python
video = client.Video().load({"id": "video_id"})
```

#### Example: List

```python
videos = client.Video().list()
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

Features are the extension mechanism. A feature is a Python class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── artinstituteofchicago_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`artinstituteofchicago_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```python
publication = client.Publication()
publication.list()

# publication.data_get() now returns the publication data from the last list
# publication.match_get() returns the last match criteria
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
