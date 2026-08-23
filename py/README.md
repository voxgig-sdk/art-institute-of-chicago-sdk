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

Create an instance: `agent = client.Agent()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alt_titles` | `Any` | Alternate names for this agent |
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `birth_date` | `Any` | The year this agent was born |
| `death_date` | `Any` | The year this agent died |
| `description` | `str` | A biographical description of the agent |
| `id` | `str` | Unique identifier of this resource. |
| `is_artist` | `bool` | Whether the agent is an artist. |
| `sort_title` | `Any` | Sortable name for this agent, typically with last name first. |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `ulan_id` | `str` | Unique identifier of this agent in Getty's ULAN |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `id` | `str` | Unique identifier of this resource. |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `id` | `str` | Unique identifier of this resource. |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `copy` | `Any` | The text of the article |
| `id` | `str` | Unique identifier of this resource. |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |

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
| `alt_artist_ids` | `Any` | Unique identifiers of the non-preferred artists/cultures associated with this work |
| `alt_classification_ids` | `Any` | Unique identifiers of all other non-preferred classification terms for this work |
| `alt_image_ids` | `Any` | Unique identifiers of all non-preferred images of this work. |
| `alt_material_ids` | `Any` | Unique identifiers of all other non-preferred material terms for this work |
| `alt_style_ids` | `Any` | Unique identifiers of all other non-preferred style terms for this work |
| `alt_subject_ids` | `Any` | Unique identifiers of all other non-preferred subject terms for this work |
| `alt_technique_ids` | `Any` | Unique identifiers of all other non-preferred technique terms for this work |
| `alt_titles` | `Any` | Alternate names for this work |
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `artist_display` | `Any` | Readable description of the creator of this work. |
| `artist_id` | `str` | Unique identifier of the preferred artist/culture associated with this work |
| `artist_ids` | `Any` | Unique identifier of all artist/cultures associated with this work |
| `artist_title` | `Any` | Name of the preferred artist/culture associated with this work |
| `artist_titles` | `Any` | Names of all artist/cultures associated with this work |
| `artwork_type_id` | `str` | Unique identifier of the kind of object or work |
| `artwork_type_title` | `Any` | The kind of object or work (e.g. |
| `boost_rank` | `Any` | Manual indication of what rank this artwork should take in search results. |
| `catalog_based_search_keyword_titles` | `Any` | The keyword search values that would be catalog-based searches on this record |
| `catalogue_display` | `Any` | Brief text listing all the catalogues raisonnés which include this work. |
| `category_ids` | `Any` | Unique identifiers of the categories this work is a part of |
| `category_titles` | `Any` | Names of the categories this artwork is a part of |
| `classification_id` | `str` | Unique identifier of the preferred classification term for this work |
| `classification_ids` | `Any` | Unique identifiers of all classification terms for this work |
| `classification_title` | `Any` | The name of the preferred classification term for this work |
| `classification_titles` | `Any` | The names of all classification terms related to this artwork |
| `color` | `Any` | Dominant color of this artwork in HSL |
| `colorfulness` | `Any` | Unbounded positive float representing an abstract measure of colorfulness. |
| `copyright_notice` | `Any` | Statement notifying how the work is protected by copyright. |
| `credit_line` | `Any` | Brief statement indicating how the work came into the collection |
| `date_display` | `Any` | Readable, free-text description of the period of time associated with the creation of this work. |
| `date_end` | `Any` | The year of the period of time associated with the creation of this work |
| `date_qualifier_id` | `str` | Unique identifier of the qualifer to the dates provided for this record. |
| `date_qualifier_title` | `Any` | Readable, text qualifer to the dates provided for this record. |
| `date_start` | `Any` | The year of the period of time associated with the creation of this work |
| `department_id` | `str` | Unique identifier of the curatorial department that this work belongs to |
| `department_title` | `Any` | Name of the curatorial department that this work belongs to |
| `description` | `str` | Longer explanation describing the work |
| `dimensions` | `Any` | The size, shape, scale, and dimensions of the work. |
| `dimensions_detail` | `Any` | The height, width, depth, and/or diameter of each section of the work in centimeters |
| `document_ids` | `Any` | Unique identifiers of assets that serve as documentation for this artwork |
| `edition` | `Any` | Edition number if the work is one of many |
| `exhibition_history` | `Any` | List of all the places this work has been exhibited |
| `fiscal_year` | `Any` | The fiscal year in which the work was acquired. |
| `fiscal_year_deaccession` | `Any` | The fiscal year in which the work was deaccessioned. |
| `gallery_id` | `str` | Unique identifier of the location of this work in our museum |
| `gallery_title` | `Any` | The location of this work in our museum |
| `has_advanced_imaging` | `bool` | Whether this artwork is enhanced with 3D models, 360 image sequences, Mirador views, etc. |
| `has_educational_resources` | `bool` | Whether this artwork has any documents tagged as educational |
| `has_multimedia_resources` | `bool` | Whether this artwork has any associated microsites, digital publications, or documents tagged as multimedia |
| `has_not_been_viewed_much` | `bool` | Whether the artwork hasn't been visited on our website very much |
| `id` | `str` | Unique identifier of this resource. |
| `image_embedding` | `Any` | The generated embeddings describing the artwork image |
| `image_id` | `str` | Unique identifier of the preferred image to use to represent this work |
| `inscriptions` | `Any` | A description of distinguishing or identifying physical markings that are on the work |
| `internal_department_id` | `str` | An internal department id we use for analytics. |
| `is_boosted` | `bool` | Whether this document should be boosted in search |
| `is_on_view` | `bool` | Whether the work is on display |
| `is_public_domain` | `bool` | Whether the work is in the public domain, meaning it was created before copyrights existed or has left the copyright term |
| `is_zoomable` | `bool` | Whether images of the work are allowed to be displayed in a zoomable interface. |
| `latitude` | `float` | Latitude coordinate of the location of this work in our galleries |
| `latlon` | `Any` | Latitude and longitude coordinates of the location of this work in our galleries |
| `longitude` | `float` | Longitude coordinate of the location of this work in our galleries |
| `main_reference_number` | `int` | Unique identifier assigned to the artwork upon acquisition |
| `material_id` | `str` | Unique identifier of the preferred material term for this work |
| `material_ids` | `Any` | Unique identifiers of all material terms for this work |
| `material_titles` | `Any` | The names of all material terms related to this artwork |
| `max_zoom_window_size` | `Any` | The maximum size of the window the image is allowed to be viewed in, in pixels. |
| `medium_display` | `Any` | The substances or materials used in the creation of a work |
| `nomisma_id` | `str` | Unique identifier of this work in the nomisma coin database |
| `on_loan_display` | `Any` | If an artwork is on loan, this contains details about the loan |
| `pageviews` | `Any` | Approx. |
| `pageviews_recent` | `Any` | Approx. |
| `place_of_origin` | `Any` | The location where the creation, design, or production of the work took place, or the original location of the work |
| `provenance_text` | `Any` | Ownership/collecting history of the work. |
| `publication_history` | `Any` | Bibliographic list of all the places this work has been published |
| `publishing_verification_level` | `Any` | Indicator of how much metadata on the work in published. |
| `section_ids` | `Any` | Unique identifiers of the digital publication chapters this work in included in |
| `section_titles` | `Any` | Names of the digital publication chapters this work is included in |
| `short_description` | `Any` | Short explanation describing the work |
| `site_ids` | `Any` | Unique identifiers of the microsites this work is a part of |
| `sound_ids` | `Any` | Unique identifiers of the audio about this work |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `style_id` | `str` | Unique identifier of the preferred style term for this work |
| `style_ids` | `Any` | Unique identifiers of all style terms for this work |
| `style_title` | `Any` | The name of the preferred style term for this work |
| `style_titles` | `Any` | The names of all style terms related to this artwork |
| `subject_id` | `str` | Unique identifier of the preferred subject term for this work |
| `subject_ids` | `Any` | Unique identifiers of all subject terms for this work |
| `subject_titles` | `Any` | The names of all subject terms related to this artwork |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `technique_id` | `str` | Unique identifier of the preferred technique term for this work |
| `technique_ids` | `Any` | Unique identifiers of all technique terms for this work |
| `technique_titles` | `Any` | The names of all technique terms related to this artwork |
| `term_titles` | `Any` | The names of the taxonomy tags for this work |
| `text_embedding` | `Any` | The generated embeddings of artwork text |
| `text_ids` | `Any` | Unique identifiers of the texts about this work |
| `theme_titles` | `Any` | The names of all thematic publish categories related to this artwork |
| `thumbnail` | `Any` | Metadata about the image referenced by `image_id`. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |
| `video_ids` | `Any` | Unique identifiers of the videos about this work |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `id` | `str` | Unique identifier of this resource. |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `id` | `str` | Unique identifier of this resource. |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |

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
| `aat_id` | `str` | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `id` | `str` | Unique identifier of this resource. |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |

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
| `aat_id` | `str` | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `id` | `str` | Unique identifier of this resource. |
| `parent_id` | `str` | Unique identifier of this category's parent |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `subtype` | `Any` | Takes one of the following values: classification, material, technique, style, subject, department, theme |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `copy` | `Any` | The text of the page |
| `id` | `str` | Unique identifier of this resource. |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | The URL to this page on our website |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `author_display` | `Any` | A display-friendly text of the authors of this article |
| `copy` | `Any` | The text of the article |
| `digital_publication_id` | `str` | Unique identifier of the digital publication this article belongs to |
| `id` | `str` | Unique identifier of this resource. |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | The URL to this article on our website |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `copy` | `Any` | The text of the page |
| `id` | `str` | Unique identifier of this resource. |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | The URL to this page on our website |

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
| `alt_audience_ids` | `Any` | Unique identifiers indicating the alternate audiences for this event |
| `alt_event_type_ids` | `Any` | Unique identifiers indicating the alternate types of this event |
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `audience_id` | `str` | Unique identifier indicating the preferred audience for this event |
| `buy_button_caption` | `Any` | Additional text below the ticket/registration button |
| `buy_button_text` | `Any` | The text used on the ticket/registration button |
| `date_display` | `Any` | A readable display of the event dates |
| `description` | `str` | All copytext of the event |
| `door_time` | `Any` | The time the doors open for this event |
| `end_date` | `Any` | The date the event ends |
| `end_time` | `Any` | The time the event ends |
| `entrance` | `Any` | Which entrance to use for this event |
| `event_host_id` | `str` | Unique identifier of the host (cf. |
| `event_host_title` | `Any` | Unique identifier of the host (cf. |
| `event_type_id` | `str` | Unique identifier indicating the preferred type of this event |
| `header_description` | `Any` | Brief description of the event displayed below the title |
| `hero_caption` | `Any` | Text displayed with the hero image on the event |
| `id` | `str` | Unique identifier of this resource. |
| `image_url` | `Any` | The URL of an image representing this page |
| `is_admission_required` | `bool` | Whether admission to the museum is required to attend this event |
| `is_after_hours` | `bool` | Whether the event is to be held after the museum closes |
| `is_free` | `bool` | Whether the event is free |
| `is_member_exclusive` | `bool` | Whether the event is exclusive to members of the museum |
| `is_private` | `bool` | Whether the event is private |
| `is_registration_required` | `bool` | Whether registration is required to attend the event |
| `is_sales_button_hidden` | `bool` | Whether the buy tickets button is hidden on the website event page |
| `is_sold_out` | `bool` | Whether the event is sold out |
| `is_ticketed` | `bool` | Whether a ticket is required to attend the event |
| `is_virtual_event` | `bool` | Whether the event is being held virtually |
| `join_url` | `Any` | URL to the membership signup page via this event |
| `layout_type` | `Any` | Number indicating the type of layout this event page uses |
| `list_description` | `Any` | One-sentence description of the event displayed in listings |
| `location` | `Any` | Where the event takes place |
| `program_ids` | `Any` | Unique identifiers indicating the programs this event is a part of |
| `program_titles` | `Any` | Titles of the programs this event is a part of |
| `rsvp_link` | `Any` | The URL to the sales site for this event |
| `search_tags` | `Any` | Editor-specified list of tags to aid in internal search |
| `short_description` | `Any` | Brief description of the event |
| `slug` | `str` | A string used in the URL for this event |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `start_date` | `Any` | The date the event begins |
| `start_time` | `Any` | The time the event starts |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `survey_url` | `Any` | URL to the survey associated with this event |
| `ticketed_event_id` | `str` | Unique identifier of the event in the ticketing system this website event is tied to |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `title_display` | `Any` | The name of this event formatted with HTML (optional) |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |
| `virtual_event_passcode` | `Any` | Passcode to access the virtual event |
| `virtual_event_url` | `Any` | URL to the virtual event |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `button_caption` | `Any` | Additional text below the ticket/registration button |
| `button_text` | `Any` | The text used on the ticket/registration button |
| `button_url` | `Any` | The URL to the sales site or an RSVP link for this event |
| `description` | `str` | Description of the event |
| `end_at` | `Any` | The date the event occurrence ends |
| `event_id` | `str` | Identifier of the master event of which this is an occurrence |
| `id` | `str` | Unique identifier of this resource. |
| `image_url` | `Any` | The URL of an image representing this page |
| `is_private` | `bool` | Whether the event is private. |
| `is_sales_button_hidden` | `bool` | Whether the buy tickets button is hidden on the website event page |
| `is_ticketed` | `bool` | Whether a ticket is required to attend the event |
| `location` | `Any` | Where the event takes place |
| `off_sale_at` | `Any` | Date and time the event goes off sale |
| `on_sale_at` | `Any` | Date and time the event goes on sale |
| `short_description` | `Any` | Brief description of the event |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `start_at` | `Any` | The date the event occurrence begins |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `title_display` | `Any` | The name of this event formatted with HTML (optional) |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `id` | `str` | Unique identifier of this resource. |
| `is_affiliate_group` | `bool` | Whether this program represents an affiliate group |
| `is_event_host` | `bool` | Whether this program represents an event host |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |

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
| `aic_end_at` | `Any` | Date the exhibition closed at the Art Institute of Chicago |
| `aic_start_at` | `Any` | Date the exhibition opened at the Art Institute of Chicago |
| `alt_image_ids` | `Any` | Unique identifiers of all non-preferred images of this exhibition. |
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `artist_ids` | `Any` | Unique identifiers of the artist agent records representing who was shown in the exhibition |
| `artwork_ids` | `Any` | Unique identifiers of the artworks that were part of the exhibition |
| `artwork_titles` | `Any` | Names of the artworks that were part of the exhibition |
| `document_ids` | `Any` | Unique identifiers of assets that serve as documentation for this exhibition |
| `gallery_id` | `str` | Unique identifier of the gallery that mainly housed the exhibition |
| `gallery_title` | `Any` | The name of the gallery that mainly housed the exhibition |
| `id` | `str` | Unique identifier of this resource. |
| `image_id` | `str` | Unique identifier of the preferred image to use to represent this exhibition |
| `image_url` | `Any` | URL to the hero image from the website |
| `is_featured` | `bool` | Is this exhibition currently featured on our website? |
| `is_published` | `bool` | Is this exhibition currently published on our website? |
| `position` | `Any` | Numering position represnting the order in which this exhibition is featured on the website |
| `short_description` | `Any` | Brief explanation of what this exhibition is |
| `site_ids` | `Any` | Unique identifiers of the microsites this exhibition is a part of |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `status` | `Any` | Whether the exhibition is open or closed |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | URL to this exhibition on our website |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `floor` | `Any` | The level the gallery is on, e.g., 1, 2, 3, or LL |
| `id` | `str` | Unique identifier of this resource. |
| `is_closed` | `bool` | Whether the gallery is currently closed |
| `latitude` | `float` | Latitude coordinate of the center of the room |
| `latlon` | `Any` | Latitude and longitude coordinates of the center of the room |
| `longitude` | `float` | Longitude coordinate of the center of the room |
| `number` | `Any` | The gallery's room number. |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | `str` | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `copy` | `Any` | The text of the page |
| `id` | `str` | Unique identifier of this resource. |
| `search_tags` | `Any` | Editor-specified list of tags to aid in internal search |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | The URL to this page on our website |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `copy` | `Any` | The text of the highlight description |
| `id` | `str` | Unique identifier of this resource. |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |

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
| `additional_text` | `Any` | Additional information about the hours |
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `friday_is_closed` | `Any` | Whether the museum is closed on Fridays |
| `friday_member_close` | `Any` | The time member hours ends on Fridays |
| `friday_member_open` | `Any` | The time member hours starts on Fridays |
| `friday_public_close` | `Any` | The time public hours ends on Fridays |
| `friday_public_open` | `Any` | The time public hours starts on Fridays |
| `id` | `str` | Unique identifier of this resource. |
| `monday_is_closed` | `Any` | Whether the museum is closed on Mondays |
| `monday_member_close` | `Any` | The time member hours ends on Mondays |
| `monday_member_open` | `Any` | The time member hours starts on Mondays |
| `monday_public_close` | `Any` | The time public hours ends on Mondays |
| `monday_public_open` | `Any` | The time public hours starts on Mondays |
| `saturday_is_closed` | `Any` | Whether the museum is closed on Saturdays |
| `saturday_member_close` | `Any` | The time member hours ends on Saturdays |
| `saturday_member_open` | `Any` | The time member hours starts on Saturdays |
| `saturday_public_close` | `Any` | The time public hours ends on Saturdays |
| `saturday_public_open` | `Any` | The time public hours starts on Saturdays |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `summary` | `Any` | Readable summary of the hours |
| `sunday_is_closed` | `Any` | Whether the museum is closed on Sundays |
| `sunday_member_close` | `Any` | The time member hours ends on Sundays |
| `sunday_member_open` | `Any` | The time member hours starts on Sundays |
| `sunday_public_close` | `Any` | The time public hours ends on Sundays |
| `sunday_public_open` | `Any` | The time public hours starts on Sundays |
| `thursday_is_closed` | `Any` | Whether the museum is closed on Thursdays |
| `thursday_member_close` | `Any` | The time member hours ends on Thursdays |
| `thursday_member_open` | `Any` | The time member hours starts on Thursdays |
| `thursday_public_close` | `Any` | The time public hours ends on Thursdays |
| `thursday_public_open` | `Any` | The time public hours starts on Thursdays |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `tuesday_is_closed` | `Any` | Whether the museum is closed on Tuesdays |
| `tuesday_member_close` | `Any` | The time member hours ends on Tuesdays |
| `tuesday_member_open` | `Any` | The time member hours starts on Tuesdays |
| `tuesday_public_close` | `Any` | The time public hours ends on Tuesdays |
| `tuesday_public_open` | `Any` | The time public hours starts on Tuesdays |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |
| `wednesday_is_closed` | `Any` | Whether the museum is closed on Wednesdays |
| `wednesday_member_close` | `Any` | The time member hours ends on Wednesdays |
| `wednesday_member_open` | `Any` | The time member hours starts on Wednesdays |
| `wednesday_public_close` | `Any` | The time public hours ends on Wednesdays |
| `wednesday_public_open` | `Any` | The time public hours starts on Wednesdays |

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
| `ahash` | `Any` | Image hash generated using ahash algorithm with 64 boolean subfields |
| `alt_text` | `Any` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `artwork_ids` | `Any` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `Any` | Names of the artworks associated with this asset |
| `color` | `Any` | Dominant color of this image in HSL |
| `colorfulness` | `Any` | Unbounded positive float representing an abstract measure of colorfulness. |
| `content` | `Any` | Text of or URL to the contents of this asset |
| `content_e_tag` | `Any` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `Any` | Asset-specific copyright information |
| `fingerprint` | `Any` | Image hashes: aHash, dHash, pHash, wHash |
| `height` | `float` | Native height of the image |
| `id` | `str` | Unique identifier of this resource. |
| `iiif_url` | `Any` | IIIF URL of this image |
| `is_educational_resource` | `bool` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | Whether this resource is considered to be educational |
| `lake_guid` | `Any` | Unique UUID of this resource in LAKE, our DAMS. |
| `lqip` | `Any` | Low-quality image placeholder (LQIP). |
| `phash` | `Any` | Image hash generated using phash algorithm with 64 boolean subfields |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `type` | `Any` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |
| `width` | `float` | Native width of the image |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `copy` | `Any` | The text of the page |
| `id` | `str` | Unique identifier of this resource. |
| `search_tags` | `Any` | Editor-specified list of tags to aid in internal search |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | The URL to this page on our website |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `id` | `str` | Unique identifier of this resource. |
| `latitude` | `float` | Latitude coordinate of the center of the room |
| `longitude` | `float` | Longitude coordinate of the center of the room |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | `str` | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `copy` | `Any` | The text of the page |
| `id` | `str` | Unique identifier of this resource. |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | The URL to this page on our website |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `copy` | `Any` | The text of the page |
| `id` | `str` | Unique identifier of this resource. |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | The URL to this page on our website |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `artist_ids` | `Any` | Unique identifiers of the artists associated with this product |
| `artwork_ids` | `Any` | Unique identifiers of the artworks associated with this product |
| `description` | `str` | Explanation of what this product is |
| `exhibition_ids` | `Any` | Unique identifiers of the exhibitions associated with this product |
| `external_sku` | `Any` | Numeric product identification code of a machine-readable barcode, when the customer sku differs from our internal one |
| `id` | `str` | Unique identifier of this resource. |
| `image_url` | `Any` | URL of an image for this product |
| `max_compare_at_price` | `Any` | Number indicating how much the most expensive variant of a product cost before a sale |
| `max_current_price` | `Any` | Number indicating how much the most expensive variant of a product costs right now |
| `min_compare_at_price` | `Any` | Number indicating how much the least expensive variant of a product cost before a sale |
| `min_current_price` | `Any` | Number indicating how much the least expensive variant of a product costs right now |
| `price_display` | `Any` | Explanation of what this product is |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | URL of this product in the shop |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `id` | `str` | Unique identifier of this resource. |
| `section_ids` | `Any` | Unique identifiers of the sections of this publication |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | URL to the publication |

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
| `api_id` | `str` | API unique identifier |
| `api_link` | `Any` | URL to this recource in the API |
| `api_model` | `Any` | Name of the model the resource represents |
| `id` | `str` | Unique identifier within the search index |
| `is_boosted` | `bool` | Whether this record has been flagged to be boosted |
| `score` | `float` | Search index ranking of the result |
| `thumbnail` | `Any` | Metadata on the image representing this record |
| `timestamp` | `Any` | Date this record was last updated in the API |
| `title` | `str` | The name of this resource |

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
| `accession` | `Any` | An accession number parsed from the title or tombstone |
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `artwork_id` | `str` | Unique identifier of the artwork with which this section is associated |
| `content` | `Any` | Content of this section in plaintext |
| `generic_page_id` | `str` | Unique identifier of the page on the website that represents the publication this section belongs to |
| `id` | `str` | Unique identifier of this resource. |
| `publication_id` | `str` | Unique identifier of the publication this section belongs to |
| `publication_title` | `Any` | Name of the publication this section belongs to |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | URL to the section |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `artwork_ids` | `Any` | Unique identifiers of the artworks this site is associated with |
| `artwork_titles` | `Any` | Names of the artworks this site is associated with |
| `description` | `str` | Explanation of what this site is |
| `exhibition_ids` | `Any` | Unique identifier of the exhibitions this site is associated with |
| `exhibition_titles` | `Any` | Names of the exhibitions this site is associated with |
| `id` | `str` | Unique identifier of this resource. |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | URL to this site |

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
| `alt_text` | `Any` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `artwork_ids` | `Any` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `Any` | Names of the artworks associated with this asset |
| `content` | `Any` | Text of or URL to the contents of this asset |
| `content_e_tag` | `Any` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `Any` | Asset-specific copyright information |
| `id` | `str` | Unique identifier of this resource. |
| `is_educational_resource` | `bool` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | Whether this resource is considered to be educational |
| `lake_guid` | `Any` | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | Name of this mobile audio file – derived from the artwork and tour titles |
| `transcript` | `Any` | Text transcription of the audio file |
| `type` | `Any` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | URL to the audio file |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `id` | `str` | Unique identifier of this resource. |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | The URL to this page on our website |

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
| `alt_text` | `Any` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `artwork_ids` | `Any` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `Any` | Names of the artworks associated with this asset |
| `content` | `Any` | Text of or URL to the contents of this asset |
| `content_e_tag` | `Any` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `Any` | Asset-specific copyright information |
| `id` | `str` | Unique identifier of this resource. |
| `is_educational_resource` | `bool` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | Whether this resource is considered to be educational |
| `lake_guid` | `Any` | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `type` | `Any` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `artist_titles` | `Any` | Names of the artists of the artworks featured in this tour's tour stops |
| `artwork_titles` | `Any` | Names of the artworks featured in this tour's tour stops |
| `description` | `str` | Explanation of what the tour is |
| `id` | `str` | Unique identifier of this resource. |
| `image` | `Any` | The main image for the tour |
| `intro` | `Any` | Text introducing the tour |
| `intro_link` | `Any` | Link to the audio file of the introduction |
| `intro_transcript` | `Any` | Transcript of the introduction audio to the tour |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |
| `weight` | `float` | Number representing this tour's sort order |

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
| `alt_text` | `Any` | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `Any` | REST API link for this resource |
| `api_model` | `Any` | REST API resource type or endpoint |
| `artwork_ids` | `Any` | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `Any` | Names of the artworks associated with this asset |
| `content` | `Any` | Text of or URL to the contents of this asset |
| `content_e_tag` | `Any` | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `Any` | Asset-specific copyright information |
| `id` | `str` | Unique identifier of this resource. |
| `is_educational_resource` | `bool` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | Whether this resource is considered to be educational |
| `lake_guid` | `Any` | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `Any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | Date and time the record was updated in the aggregator search index |
| `title` | `str` | The name of this resource |
| `type` | `Any` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `Any` | Date and time the record was updated in the aggregator database |

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
