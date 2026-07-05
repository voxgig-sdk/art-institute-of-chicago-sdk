# ArtInstituteOfChicago Python SDK Reference

Complete API reference for the ArtInstituteOfChicago Python SDK.


## ArtInstituteOfChicagoSDK

### Constructor

```python
from artinstituteofchicago_sdk import ArtInstituteOfChicagoSDK

client = ArtInstituteOfChicagoSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `ArtInstituteOfChicagoSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = ArtInstituteOfChicagoSDK.test()
```


### Instance Methods

#### `Agent(data=None)`

Create a new `AgentEntity` instance. Pass `None` for no initial data.

#### `AgentRole(data=None)`

Create a new `AgentRoleEntity` instance. Pass `None` for no initial data.

#### `AgentType(data=None)`

Create a new `AgentTypeEntity` instance. Pass `None` for no initial data.

#### `Article(data=None)`

Create a new `ArticleEntity` instance. Pass `None` for no initial data.

#### `Artwork(data=None)`

Create a new `ArtworkEntity` instance. Pass `None` for no initial data.

#### `ArtworkDateQualifier(data=None)`

Create a new `ArtworkDateQualifierEntity` instance. Pass `None` for no initial data.

#### `ArtworkPlaceQualifier(data=None)`

Create a new `ArtworkPlaceQualifierEntity` instance. Pass `None` for no initial data.

#### `ArtworkType(data=None)`

Create a new `ArtworkTypeEntity` instance. Pass `None` for no initial data.

#### `CategoryTerm(data=None)`

Create a new `CategoryTermEntity` instance. Pass `None` for no initial data.

#### `DigitalPublication(data=None)`

Create a new `DigitalPublicationEntity` instance. Pass `None` for no initial data.

#### `DigitalPublicationArticle(data=None)`

Create a new `DigitalPublicationArticleEntity` instance. Pass `None` for no initial data.

#### `EducatorResource(data=None)`

Create a new `EducatorResourceEntity` instance. Pass `None` for no initial data.

#### `Event(data=None)`

Create a new `EventEntity` instance. Pass `None` for no initial data.

#### `EventOccurrence(data=None)`

Create a new `EventOccurrenceEntity` instance. Pass `None` for no initial data.

#### `EventProgram(data=None)`

Create a new `EventProgramEntity` instance. Pass `None` for no initial data.

#### `Exhibition(data=None)`

Create a new `ExhibitionEntity` instance. Pass `None` for no initial data.

#### `Gallery(data=None)`

Create a new `GalleryEntity` instance. Pass `None` for no initial data.

#### `GenericPage(data=None)`

Create a new `GenericPageEntity` instance. Pass `None` for no initial data.

#### `Highlight(data=None)`

Create a new `HighlightEntity` instance. Pass `None` for no initial data.

#### `Hour(data=None)`

Create a new `HourEntity` instance. Pass `None` for no initial data.

#### `Image(data=None)`

Create a new `ImageEntity` instance. Pass `None` for no initial data.

#### `LandingPage(data=None)`

Create a new `LandingPageEntity` instance. Pass `None` for no initial data.

#### `Place(data=None)`

Create a new `PlaceEntity` instance. Pass `None` for no initial data.

#### `PressRelease(data=None)`

Create a new `PressReleaseEntity` instance. Pass `None` for no initial data.

#### `PrintedPublication(data=None)`

Create a new `PrintedPublicationEntity` instance. Pass `None` for no initial data.

#### `Product(data=None)`

Create a new `ProductEntity` instance. Pass `None` for no initial data.

#### `Publication(data=None)`

Create a new `PublicationEntity` instance. Pass `None` for no initial data.

#### `Search(data=None)`

Create a new `SearchEntity` instance. Pass `None` for no initial data.

#### `Section(data=None)`

Create a new `SectionEntity` instance. Pass `None` for no initial data.

#### `Site(data=None)`

Create a new `SiteEntity` instance. Pass `None` for no initial data.

#### `Sound(data=None)`

Create a new `SoundEntity` instance. Pass `None` for no initial data.

#### `StaticPage(data=None)`

Create a new `StaticPageEntity` instance. Pass `None` for no initial data.

#### `Text(data=None)`

Create a new `TextEntity` instance. Pass `None` for no initial data.

#### `Tour(data=None)`

Create a new `TourEntity` instance. Pass `None` for no initial data.

#### `Video(data=None)`

Create a new `VideoEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## AgentEntity

```python
agent = client.Agent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_title` | `Any` | No |  |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `birth_date` | `Any` | No |  |
| `death_date` | `Any` | No |  |
| `description` | `str` | No |  |
| `id` | `str` | No |  |
| `is_artist` | `bool` | No |  |
| `sort_title` | `Any` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `ulan_id` | `str` | No |  |
| `updated_at` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Agent().list()
for agent in results:
    print(agent)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Agent().load({"id": "agent_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AgentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## AgentRoleEntity

```python
agent_role = client.AgentRole()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `id` | `str` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.AgentRole().list()
for agent_role in results:
    print(agent_role)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.AgentRole().load({"id": "agent_role_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AgentRoleEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## AgentTypeEntity

```python
agent_type = client.AgentType()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `id` | `str` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.AgentType().list()
for agent_type in results:
    print(agent_type)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.AgentType().load({"id": "agent_type_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AgentTypeEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ArticleEntity

```python
article = client.Article()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `copy` | `Any` | No |  |
| `id` | `str` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Article().list()
for article in results:
    print(article)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Article().load({"id": "article_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ArticleEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ArtworkEntity

```python
artwork = client.Artwork()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_artist_id` | `str` | No |  |
| `alt_classification_id` | `str` | No |  |
| `alt_image_id` | `str` | No |  |
| `alt_material_id` | `str` | No |  |
| `alt_style_id` | `str` | No |  |
| `alt_subject_id` | `str` | No |  |
| `alt_technique_id` | `str` | No |  |
| `alt_title` | `Any` | No |  |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `artist_display` | `Any` | No |  |
| `artist_id` | `str` | No |  |
| `artist_title` | `Any` | No |  |
| `artwork_type_id` | `str` | No |  |
| `artwork_type_title` | `Any` | No |  |
| `boost_rank` | `Any` | No |  |
| `catalog_based_search_keyword_title` | `Any` | No |  |
| `catalogue_display` | `Any` | No |  |
| `category_id` | `str` | No |  |
| `category_title` | `Any` | No |  |
| `classification_id` | `str` | No |  |
| `classification_title` | `Any` | No |  |
| `color` | `Any` | No |  |
| `colorfulness` | `Any` | No |  |
| `copyright_notice` | `Any` | No |  |
| `credit_line` | `Any` | No |  |
| `date_display` | `Any` | No |  |
| `date_end` | `Any` | No |  |
| `date_qualifier_id` | `str` | No |  |
| `date_qualifier_title` | `Any` | No |  |
| `date_start` | `Any` | No |  |
| `department_id` | `str` | No |  |
| `department_title` | `Any` | No |  |
| `description` | `str` | No |  |
| `dimension` | `Any` | No |  |
| `dimensions_detail` | `Any` | No |  |
| `document_id` | `str` | No |  |
| `edition` | `Any` | No |  |
| `exhibition_history` | `Any` | No |  |
| `fiscal_year` | `Any` | No |  |
| `fiscal_year_deaccession` | `Any` | No |  |
| `gallery_id` | `str` | No |  |
| `gallery_title` | `Any` | No |  |
| `has_advanced_imaging` | `bool` | No |  |
| `has_educational_resource` | `bool` | No |  |
| `has_multimedia_resource` | `bool` | No |  |
| `has_not_been_viewed_much` | `bool` | No |  |
| `id` | `str` | No |  |
| `image_embedding` | `Any` | No |  |
| `image_id` | `str` | No |  |
| `inscription` | `Any` | No |  |
| `internal_department_id` | `str` | No |  |
| `is_boosted` | `bool` | No |  |
| `is_on_view` | `bool` | No |  |
| `is_public_domain` | `bool` | No |  |
| `is_zoomable` | `bool` | No |  |
| `latitude` | `float` | No |  |
| `latlon` | `Any` | No |  |
| `longitude` | `float` | No |  |
| `main_reference_number` | `int` | No |  |
| `material_id` | `str` | No |  |
| `material_title` | `Any` | No |  |
| `max_zoom_window_size` | `Any` | No |  |
| `medium_display` | `Any` | No |  |
| `nomisma_id` | `str` | No |  |
| `on_loan_display` | `Any` | No |  |
| `pageview` | `Any` | No |  |
| `pageviews_recent` | `Any` | No |  |
| `place_of_origin` | `Any` | No |  |
| `provenance_text` | `Any` | No |  |
| `publication_history` | `Any` | No |  |
| `publishing_verification_level` | `Any` | No |  |
| `section_id` | `str` | No |  |
| `section_title` | `Any` | No |  |
| `short_description` | `Any` | No |  |
| `site_id` | `str` | No |  |
| `sound_id` | `str` | No |  |
| `source_updated_at` | `Any` | No |  |
| `style_id` | `str` | No |  |
| `style_title` | `Any` | No |  |
| `subject_id` | `str` | No |  |
| `subject_title` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `technique_id` | `str` | No |  |
| `technique_title` | `Any` | No |  |
| `term_title` | `Any` | No |  |
| `text_embedding` | `Any` | No |  |
| `text_id` | `str` | No |  |
| `theme_title` | `Any` | No |  |
| `thumbnail` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |
| `video_id` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Artwork().list()
for artwork in results:
    print(artwork)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Artwork().load({"id": "artwork_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ArtworkEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ArtworkDateQualifierEntity

```python
artwork_date_qualifier = client.ArtworkDateQualifier()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `id` | `str` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.ArtworkDateQualifier().list()
for artwork_date_qualifier in results:
    print(artwork_date_qualifier)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.ArtworkDateQualifier().load({"id": "artwork_date_qualifier_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ArtworkDateQualifierEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ArtworkPlaceQualifierEntity

```python
artwork_place_qualifier = client.ArtworkPlaceQualifier()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `id` | `str` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.ArtworkPlaceQualifier().list()
for artwork_place_qualifier in results:
    print(artwork_place_qualifier)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.ArtworkPlaceQualifier().load({"id": "artwork_place_qualifier_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ArtworkPlaceQualifierEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ArtworkTypeEntity

```python
artwork_type = client.ArtworkType()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `aat_id` | `str` | No |  |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `id` | `str` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.ArtworkType().list()
for artwork_type in results:
    print(artwork_type)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.ArtworkType().load({"id": "artwork_type_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ArtworkTypeEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CategoryTermEntity

```python
category_term = client.CategoryTerm()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `aat_id` | `str` | No |  |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `id` | `str` | No |  |
| `parent_id` | `str` | No |  |
| `source_updated_at` | `Any` | No |  |
| `subtype` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.CategoryTerm().list()
for category_term in results:
    print(category_term)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.CategoryTerm().load({"id": "category_term_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CategoryTermEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## DigitalPublicationEntity

```python
digital_publication = client.DigitalPublication()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `copy` | `Any` | No |  |
| `id` | `str` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |
| `web_url` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.DigitalPublication().list()
for digital_publication in results:
    print(digital_publication)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.DigitalPublication().load({"id": "digital_publication_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DigitalPublicationEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## DigitalPublicationArticleEntity

```python
digital_publication_article = client.DigitalPublicationArticle()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `author_display` | `Any` | No |  |
| `copy` | `Any` | No |  |
| `digital_publication_id` | `str` | No |  |
| `id` | `str` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |
| `web_url` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.DigitalPublicationArticle().list()
for digital_publication_article in results:
    print(digital_publication_article)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.DigitalPublicationArticle().load({"id": "digital_publication_article_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DigitalPublicationArticleEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EducatorResourceEntity

```python
educator_resource = client.EducatorResource()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `copy` | `Any` | No |  |
| `id` | `str` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |
| `web_url` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.EducatorResource().list()
for educator_resource in results:
    print(educator_resource)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.EducatorResource().load({"id": "educator_resource_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EducatorResourceEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EventEntity

```python
event = client.Event()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_audience_id` | `str` | No |  |
| `alt_event_type_id` | `str` | No |  |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `audience_id` | `str` | No |  |
| `buy_button_caption` | `Any` | No |  |
| `buy_button_text` | `Any` | No |  |
| `date_display` | `Any` | No |  |
| `description` | `str` | No |  |
| `door_time` | `Any` | No |  |
| `end_date` | `Any` | No |  |
| `end_time` | `Any` | No |  |
| `entrance` | `Any` | No |  |
| `event_host_id` | `str` | No |  |
| `event_host_title` | `Any` | No |  |
| `event_type_id` | `str` | No |  |
| `header_description` | `Any` | No |  |
| `hero_caption` | `Any` | No |  |
| `id` | `str` | No |  |
| `image_url` | `Any` | No |  |
| `is_admission_required` | `bool` | No |  |
| `is_after_hour` | `bool` | No |  |
| `is_free` | `bool` | No |  |
| `is_member_exclusive` | `bool` | No |  |
| `is_private` | `bool` | No |  |
| `is_registration_required` | `bool` | No |  |
| `is_sales_button_hidden` | `bool` | No |  |
| `is_sold_out` | `bool` | No |  |
| `is_ticketed` | `bool` | No |  |
| `is_virtual_event` | `bool` | No |  |
| `join_url` | `Any` | No |  |
| `layout_type` | `Any` | No |  |
| `list_description` | `Any` | No |  |
| `location` | `Any` | No |  |
| `program_id` | `str` | No |  |
| `program_title` | `Any` | No |  |
| `rsvp_link` | `Any` | No |  |
| `search_tag` | `Any` | No |  |
| `short_description` | `Any` | No |  |
| `slug` | `str` | No |  |
| `source_updated_at` | `Any` | No |  |
| `start_date` | `Any` | No |  |
| `start_time` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `survey_url` | `Any` | No |  |
| `ticketed_event_id` | `str` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `title_display` | `Any` | No |  |
| `updated_at` | `Any` | No |  |
| `virtual_event_passcode` | `Any` | No |  |
| `virtual_event_url` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Event().list()
for event in results:
    print(event)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Event().load({"id": "event_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EventEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EventOccurrenceEntity

```python
event_occurrence = client.EventOccurrence()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `button_caption` | `Any` | No |  |
| `button_text` | `Any` | No |  |
| `button_url` | `Any` | No |  |
| `description` | `str` | No |  |
| `end_at` | `Any` | No |  |
| `event_id` | `str` | No |  |
| `id` | `str` | No |  |
| `image_url` | `Any` | No |  |
| `is_private` | `bool` | No |  |
| `is_sales_button_hidden` | `bool` | No |  |
| `is_ticketed` | `bool` | No |  |
| `location` | `Any` | No |  |
| `off_sale_at` | `Any` | No |  |
| `on_sale_at` | `Any` | No |  |
| `short_description` | `Any` | No |  |
| `source_updated_at` | `Any` | No |  |
| `start_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `title_display` | `Any` | No |  |
| `updated_at` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.EventOccurrence().list()
for event_occurrence in results:
    print(event_occurrence)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.EventOccurrence().load({"id": "event_occurrence_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EventOccurrenceEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EventProgramEntity

```python
event_program = client.EventProgram()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `id` | `str` | No |  |
| `is_affiliate_group` | `bool` | No |  |
| `is_event_host` | `bool` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.EventProgram().list()
for event_program in results:
    print(event_program)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.EventProgram().load({"id": "event_program_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EventProgramEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ExhibitionEntity

```python
exhibition = client.Exhibition()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `aic_end_at` | `Any` | No |  |
| `aic_start_at` | `Any` | No |  |
| `alt_image_id` | `str` | No |  |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `artist_id` | `str` | No |  |
| `artwork_id` | `str` | No |  |
| `artwork_title` | `Any` | No |  |
| `document_id` | `str` | No |  |
| `gallery_id` | `str` | No |  |
| `gallery_title` | `Any` | No |  |
| `id` | `str` | No |  |
| `image_id` | `str` | No |  |
| `image_url` | `Any` | No |  |
| `is_featured` | `bool` | No |  |
| `is_published` | `bool` | No |  |
| `position` | `Any` | No |  |
| `short_description` | `Any` | No |  |
| `site_id` | `str` | No |  |
| `source_updated_at` | `Any` | No |  |
| `status` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |
| `web_url` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Exhibition().list()
for exhibition in results:
    print(exhibition)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Exhibition().load({"id": "exhibition_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ExhibitionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## GalleryEntity

```python
gallery = client.Gallery()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `floor` | `Any` | No |  |
| `id` | `str` | No |  |
| `is_closed` | `bool` | No |  |
| `latitude` | `float` | No |  |
| `latlon` | `Any` | No |  |
| `longitude` | `float` | No |  |
| `number` | `Any` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `tgn_id` | `str` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Gallery().list()
for gallery in results:
    print(gallery)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Gallery().load({"id": "gallery_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GalleryEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## GenericPageEntity

```python
generic_page = client.GenericPage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `copy` | `Any` | No |  |
| `id` | `str` | No |  |
| `search_tag` | `Any` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |
| `web_url` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.GenericPage().list()
for generic_page in results:
    print(generic_page)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.GenericPage().load({"id": "generic_page_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GenericPageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## HighlightEntity

```python
highlight = client.Highlight()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `copy` | `Any` | No |  |
| `id` | `str` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Highlight().list()
for highlight in results:
    print(highlight)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Highlight().load({"id": "highlight_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `HighlightEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## HourEntity

```python
hour = client.Hour()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_text` | `Any` | No |  |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `friday_is_closed` | `Any` | No |  |
| `friday_member_close` | `Any` | No |  |
| `friday_member_open` | `Any` | No |  |
| `friday_public_close` | `Any` | No |  |
| `friday_public_open` | `Any` | No |  |
| `id` | `str` | No |  |
| `monday_is_closed` | `Any` | No |  |
| `monday_member_close` | `Any` | No |  |
| `monday_member_open` | `Any` | No |  |
| `monday_public_close` | `Any` | No |  |
| `monday_public_open` | `Any` | No |  |
| `saturday_is_closed` | `Any` | No |  |
| `saturday_member_close` | `Any` | No |  |
| `saturday_member_open` | `Any` | No |  |
| `saturday_public_close` | `Any` | No |  |
| `saturday_public_open` | `Any` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `summary` | `Any` | No |  |
| `sunday_is_closed` | `Any` | No |  |
| `sunday_member_close` | `Any` | No |  |
| `sunday_member_open` | `Any` | No |  |
| `sunday_public_close` | `Any` | No |  |
| `sunday_public_open` | `Any` | No |  |
| `thursday_is_closed` | `Any` | No |  |
| `thursday_member_close` | `Any` | No |  |
| `thursday_member_open` | `Any` | No |  |
| `thursday_public_close` | `Any` | No |  |
| `thursday_public_open` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `tuesday_is_closed` | `Any` | No |  |
| `tuesday_member_close` | `Any` | No |  |
| `tuesday_member_open` | `Any` | No |  |
| `tuesday_public_close` | `Any` | No |  |
| `tuesday_public_open` | `Any` | No |  |
| `updated_at` | `Any` | No |  |
| `wednesday_is_closed` | `Any` | No |  |
| `wednesday_member_close` | `Any` | No |  |
| `wednesday_member_open` | `Any` | No |  |
| `wednesday_public_close` | `Any` | No |  |
| `wednesday_public_open` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Hour().list()
for hour in results:
    print(hour)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Hour().load({"id": "hour_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `HourEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ImageEntity

```python
image = client.Image()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ahash` | `Any` | No |  |
| `alt_text` | `Any` | No |  |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `artwork_id` | `str` | No |  |
| `artwork_title` | `Any` | No |  |
| `color` | `Any` | No |  |
| `colorfulness` | `Any` | No |  |
| `content` | `Any` | No |  |
| `content_e_tag` | `Any` | No |  |
| `credit_line` | `Any` | No |  |
| `fingerprint` | `Any` | No |  |
| `height` | `float` | No |  |
| `id` | `str` | No |  |
| `iiif_url` | `Any` | No |  |
| `is_educational_resource` | `bool` | No |  |
| `is_multimedia_resource` | `bool` | No |  |
| `is_teacher_resource` | `bool` | No |  |
| `lake_guid` | `Any` | No |  |
| `lqip` | `Any` | No |  |
| `phash` | `Any` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `type` | `Any` | No |  |
| `updated_at` | `Any` | No |  |
| `width` | `float` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Image().list()
for image in results:
    print(image)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Image().load({"id": "image_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ImageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## LandingPageEntity

```python
landing_page = client.LandingPage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `copy` | `Any` | No |  |
| `id` | `str` | No |  |
| `search_tag` | `Any` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |
| `web_url` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.LandingPage().list()
for landing_page in results:
    print(landing_page)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.LandingPage().load({"id": "landing_page_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `LandingPageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PlaceEntity

```python
place = client.Place()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `id` | `str` | No |  |
| `latitude` | `float` | No |  |
| `longitude` | `float` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `tgn_id` | `str` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Place().list()
for place in results:
    print(place)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Place().load({"id": "place_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PlaceEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PressReleaseEntity

```python
press_release = client.PressRelease()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `copy` | `Any` | No |  |
| `id` | `str` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |
| `web_url` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PressRelease().list()
for press_release in results:
    print(press_release)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.PressRelease().load({"id": "press_release_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PressReleaseEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PrintedPublicationEntity

```python
printed_publication = client.PrintedPublication()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `copy` | `Any` | No |  |
| `id` | `str` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |
| `web_url` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PrintedPublication().list()
for printed_publication in results:
    print(printed_publication)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.PrintedPublication().load({"id": "printed_publication_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PrintedPublicationEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ProductEntity

```python
product = client.Product()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `artist_id` | `str` | No |  |
| `artwork_id` | `str` | No |  |
| `description` | `str` | No |  |
| `exhibition_id` | `str` | No |  |
| `external_sku` | `Any` | No |  |
| `id` | `str` | No |  |
| `image_url` | `Any` | No |  |
| `max_compare_at_price` | `Any` | No |  |
| `max_current_price` | `Any` | No |  |
| `min_compare_at_price` | `Any` | No |  |
| `min_current_price` | `Any` | No |  |
| `price_display` | `Any` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |
| `web_url` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Product().list()
for product in results:
    print(product)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Product().load({"id": "product_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ProductEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PublicationEntity

```python
publication = client.Publication()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `id` | `str` | No |  |
| `section_id` | `str` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |
| `web_url` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Publication().list()
for publication in results:
    print(publication)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Publication().load({"id": "publication_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PublicationEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SearchEntity

```python
search = client.Search()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_id` | `str` | No |  |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `id` | `str` | No |  |
| `is_boosted` | `bool` | No |  |
| `score` | `float` | No |  |
| `thumbnail` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Search().list()
for search in results:
    print(search)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SearchEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SectionEntity

```python
section = client.Section()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accession` | `Any` | No |  |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `artwork_id` | `str` | No |  |
| `content` | `Any` | No |  |
| `generic_page_id` | `str` | No |  |
| `id` | `str` | No |  |
| `publication_id` | `str` | No |  |
| `publication_title` | `Any` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |
| `web_url` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Section().list()
for section in results:
    print(section)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Section().load({"id": "section_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SectionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SiteEntity

```python
site = client.Site()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `artwork_id` | `str` | No |  |
| `artwork_title` | `Any` | No |  |
| `description` | `str` | No |  |
| `exhibition_id` | `str` | No |  |
| `exhibition_title` | `Any` | No |  |
| `id` | `str` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |
| `web_url` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Site().list()
for site in results:
    print(site)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Site().load({"id": "site_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SiteEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SoundEntity

```python
sound = client.Sound()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_text` | `Any` | No |  |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `artwork_id` | `str` | No |  |
| `artwork_title` | `Any` | No |  |
| `content` | `Any` | No |  |
| `content_e_tag` | `Any` | No |  |
| `credit_line` | `Any` | No |  |
| `id` | `str` | No |  |
| `is_educational_resource` | `bool` | No |  |
| `is_multimedia_resource` | `bool` | No |  |
| `is_teacher_resource` | `bool` | No |  |
| `lake_guid` | `Any` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `transcript` | `Any` | No |  |
| `type` | `Any` | No |  |
| `updated_at` | `Any` | No |  |
| `web_url` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Sound().list()
for sound in results:
    print(sound)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Sound().load({"id": "sound_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SoundEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## StaticPageEntity

```python
static_page = client.StaticPage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `id` | `str` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |
| `web_url` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.StaticPage().list()
for static_page in results:
    print(static_page)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.StaticPage().load({"id": "static_page_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `StaticPageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## TextEntity

```python
text = client.Text()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_text` | `Any` | No |  |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `artwork_id` | `str` | No |  |
| `artwork_title` | `Any` | No |  |
| `content` | `Any` | No |  |
| `content_e_tag` | `Any` | No |  |
| `credit_line` | `Any` | No |  |
| `id` | `str` | No |  |
| `is_educational_resource` | `bool` | No |  |
| `is_multimedia_resource` | `bool` | No |  |
| `is_teacher_resource` | `bool` | No |  |
| `lake_guid` | `Any` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `type` | `Any` | No |  |
| `updated_at` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Text().list()
for text in results:
    print(text)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Text().load({"id": "text_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TextEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## TourEntity

```python
tour = client.Tour()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `artist_title` | `Any` | No |  |
| `artwork_title` | `Any` | No |  |
| `description` | `str` | No |  |
| `id` | `str` | No |  |
| `image` | `Any` | No |  |
| `intro` | `Any` | No |  |
| `intro_link` | `Any` | No |  |
| `intro_transcript` | `Any` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `updated_at` | `Any` | No |  |
| `weight` | `float` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Tour().list()
for tour in results:
    print(tour)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Tour().load({"id": "tour_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TourEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## VideoEntity

```python
video = client.Video()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alt_text` | `Any` | No |  |
| `api_link` | `Any` | No |  |
| `api_model` | `Any` | No |  |
| `artwork_id` | `str` | No |  |
| `artwork_title` | `Any` | No |  |
| `content` | `Any` | No |  |
| `content_e_tag` | `Any` | No |  |
| `credit_line` | `Any` | No |  |
| `id` | `str` | No |  |
| `is_educational_resource` | `bool` | No |  |
| `is_multimedia_resource` | `bool` | No |  |
| `is_teacher_resource` | `bool` | No |  |
| `lake_guid` | `Any` | No |  |
| `source_updated_at` | `Any` | No |  |
| `suggest_autocomplete_all` | `Any` | No |  |
| `suggest_autocomplete_boosted` | `Any` | No |  |
| `timestamp` | `Any` | No |  |
| `title` | `str` | No |  |
| `type` | `Any` | No |  |
| `updated_at` | `Any` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Video().list()
for video in results:
    print(video)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Video().load({"id": "video_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `VideoEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```python
client = ArtInstituteOfChicagoSDK({
    "feature": {
        "test": {"active": True},
    },
})
```

