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
| `alt_titles` | `Any` | No | Alternate names for this agent |
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `birth_date` | `Any` | No | The year this agent was born |
| `death_date` | `Any` | No | The year this agent died |
| `description` | `str` | No | A biographical description of the agent |
| `id` | `str` | No | Unique identifier of this resource. |
| `is_artist` | `bool` | No | Whether the agent is an artist. |
| `sort_title` | `Any` | No | Sortable name for this agent, typically with last name first. |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `ulan_id` | `str` | No | Unique identifier of this agent in Getty's ULAN |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `id` | `str` | No | Unique identifier of this resource. |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `id` | `str` | No | Unique identifier of this resource. |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `copy` | `Any` | No | The text of the article |
| `id` | `str` | No | Unique identifier of this resource. |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |

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
| `alt_artist_ids` | `Any` | No | Unique identifiers of the non-preferred artists/cultures associated with this work |
| `alt_classification_ids` | `Any` | No | Unique identifiers of all other non-preferred classification terms for this work |
| `alt_image_ids` | `Any` | No | Unique identifiers of all non-preferred images of this work. |
| `alt_material_ids` | `Any` | No | Unique identifiers of all other non-preferred material terms for this work |
| `alt_style_ids` | `Any` | No | Unique identifiers of all other non-preferred style terms for this work |
| `alt_subject_ids` | `Any` | No | Unique identifiers of all other non-preferred subject terms for this work |
| `alt_technique_ids` | `Any` | No | Unique identifiers of all other non-preferred technique terms for this work |
| `alt_titles` | `Any` | No | Alternate names for this work |
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `artist_display` | `Any` | No | Readable description of the creator of this work. |
| `artist_id` | `str` | No | Unique identifier of the preferred artist/culture associated with this work |
| `artist_ids` | `Any` | No | Unique identifier of all artist/cultures associated with this work |
| `artist_title` | `Any` | No | Name of the preferred artist/culture associated with this work |
| `artist_titles` | `Any` | No | Names of all artist/cultures associated with this work |
| `artwork_type_id` | `str` | No | Unique identifier of the kind of object or work |
| `artwork_type_title` | `Any` | No | The kind of object or work (e.g. |
| `boost_rank` | `Any` | No | Manual indication of what rank this artwork should take in search results. |
| `catalog_based_search_keyword_titles` | `Any` | No | The keyword search values that would be catalog-based searches on this record |
| `catalogue_display` | `Any` | No | Brief text listing all the catalogues raisonnés which include this work. |
| `category_ids` | `Any` | No | Unique identifiers of the categories this work is a part of |
| `category_titles` | `Any` | No | Names of the categories this artwork is a part of |
| `classification_id` | `str` | No | Unique identifier of the preferred classification term for this work |
| `classification_ids` | `Any` | No | Unique identifiers of all classification terms for this work |
| `classification_title` | `Any` | No | The name of the preferred classification term for this work |
| `classification_titles` | `Any` | No | The names of all classification terms related to this artwork |
| `color` | `Any` | No | Dominant color of this artwork in HSL |
| `colorfulness` | `Any` | No | Unbounded positive float representing an abstract measure of colorfulness. |
| `copyright_notice` | `Any` | No | Statement notifying how the work is protected by copyright. |
| `credit_line` | `Any` | No | Brief statement indicating how the work came into the collection |
| `date_display` | `Any` | No | Readable, free-text description of the period of time associated with the creation of this work. |
| `date_end` | `Any` | No | The year of the period of time associated with the creation of this work |
| `date_qualifier_id` | `str` | No | Unique identifier of the qualifer to the dates provided for this record. |
| `date_qualifier_title` | `Any` | No | Readable, text qualifer to the dates provided for this record. |
| `date_start` | `Any` | No | The year of the period of time associated with the creation of this work |
| `department_id` | `str` | No | Unique identifier of the curatorial department that this work belongs to |
| `department_title` | `Any` | No | Name of the curatorial department that this work belongs to |
| `description` | `str` | No | Longer explanation describing the work |
| `dimensions` | `Any` | No | The size, shape, scale, and dimensions of the work. |
| `dimensions_detail` | `Any` | No | The height, width, depth, and/or diameter of each section of the work in centimeters |
| `document_ids` | `Any` | No | Unique identifiers of assets that serve as documentation for this artwork |
| `edition` | `Any` | No | Edition number if the work is one of many |
| `exhibition_history` | `Any` | No | List of all the places this work has been exhibited |
| `fiscal_year` | `Any` | No | The fiscal year in which the work was acquired. |
| `fiscal_year_deaccession` | `Any` | No | The fiscal year in which the work was deaccessioned. |
| `gallery_id` | `str` | No | Unique identifier of the location of this work in our museum |
| `gallery_title` | `Any` | No | The location of this work in our museum |
| `has_advanced_imaging` | `bool` | No | Whether this artwork is enhanced with 3D models, 360 image sequences, Mirador views, etc. |
| `has_educational_resources` | `bool` | No | Whether this artwork has any documents tagged as educational |
| `has_multimedia_resources` | `bool` | No | Whether this artwork has any associated microsites, digital publications, or documents tagged as multimedia |
| `has_not_been_viewed_much` | `bool` | No | Whether the artwork hasn't been visited on our website very much |
| `id` | `str` | No | Unique identifier of this resource. |
| `image_embedding` | `Any` | No | The generated embeddings describing the artwork image |
| `image_id` | `str` | No | Unique identifier of the preferred image to use to represent this work |
| `inscriptions` | `Any` | No | A description of distinguishing or identifying physical markings that are on the work |
| `internal_department_id` | `str` | No | An internal department id we use for analytics. |
| `is_boosted` | `bool` | No | Whether this document should be boosted in search |
| `is_on_view` | `bool` | No | Whether the work is on display |
| `is_public_domain` | `bool` | No | Whether the work is in the public domain, meaning it was created before copyrights existed or has left the copyright term |
| `is_zoomable` | `bool` | No | Whether images of the work are allowed to be displayed in a zoomable interface. |
| `latitude` | `float` | No | Latitude coordinate of the location of this work in our galleries |
| `latlon` | `Any` | No | Latitude and longitude coordinates of the location of this work in our galleries |
| `longitude` | `float` | No | Longitude coordinate of the location of this work in our galleries |
| `main_reference_number` | `int` | No | Unique identifier assigned to the artwork upon acquisition |
| `material_id` | `str` | No | Unique identifier of the preferred material term for this work |
| `material_ids` | `Any` | No | Unique identifiers of all material terms for this work |
| `material_titles` | `Any` | No | The names of all material terms related to this artwork |
| `max_zoom_window_size` | `Any` | No | The maximum size of the window the image is allowed to be viewed in, in pixels. |
| `medium_display` | `Any` | No | The substances or materials used in the creation of a work |
| `nomisma_id` | `str` | No | Unique identifier of this work in the nomisma coin database |
| `on_loan_display` | `Any` | No | If an artwork is on loan, this contains details about the loan |
| `pageviews` | `Any` | No | Approx. |
| `pageviews_recent` | `Any` | No | Approx. |
| `place_of_origin` | `Any` | No | The location where the creation, design, or production of the work took place, or the original location of the work |
| `provenance_text` | `Any` | No | Ownership/collecting history of the work. |
| `publication_history` | `Any` | No | Bibliographic list of all the places this work has been published |
| `publishing_verification_level` | `Any` | No | Indicator of how much metadata on the work in published. |
| `section_ids` | `Any` | No | Unique identifiers of the digital publication chapters this work in included in |
| `section_titles` | `Any` | No | Names of the digital publication chapters this work is included in |
| `short_description` | `Any` | No | Short explanation describing the work |
| `site_ids` | `Any` | No | Unique identifiers of the microsites this work is a part of |
| `sound_ids` | `Any` | No | Unique identifiers of the audio about this work |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `style_id` | `str` | No | Unique identifier of the preferred style term for this work |
| `style_ids` | `Any` | No | Unique identifiers of all style terms for this work |
| `style_title` | `Any` | No | The name of the preferred style term for this work |
| `style_titles` | `Any` | No | The names of all style terms related to this artwork |
| `subject_id` | `str` | No | Unique identifier of the preferred subject term for this work |
| `subject_ids` | `Any` | No | Unique identifiers of all subject terms for this work |
| `subject_titles` | `Any` | No | The names of all subject terms related to this artwork |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `technique_id` | `str` | No | Unique identifier of the preferred technique term for this work |
| `technique_ids` | `Any` | No | Unique identifiers of all technique terms for this work |
| `technique_titles` | `Any` | No | The names of all technique terms related to this artwork |
| `term_titles` | `Any` | No | The names of the taxonomy tags for this work |
| `text_embedding` | `Any` | No | The generated embeddings of artwork text |
| `text_ids` | `Any` | No | Unique identifiers of the texts about this work |
| `theme_titles` | `Any` | No | The names of all thematic publish categories related to this artwork |
| `thumbnail` | `Any` | No | Metadata about the image referenced by `image_id`. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |
| `video_ids` | `Any` | No | Unique identifiers of the videos about this work |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `id` | `str` | No | Unique identifier of this resource. |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `id` | `str` | No | Unique identifier of this resource. |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |

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
| `aat_id` | `str` | No | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `id` | `str` | No | Unique identifier of this resource. |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |

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
| `aat_id` | `str` | No | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `id` | `str` | No | Unique identifier of this resource. |
| `parent_id` | `str` | No | Unique identifier of this category's parent |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `subtype` | `Any` | No | Takes one of the following values: classification, material, technique, style, subject, department, theme |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `copy` | `Any` | No | The text of the page |
| `id` | `str` | No | Unique identifier of this resource. |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | No | The URL to this page on our website |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `author_display` | `Any` | No | A display-friendly text of the authors of this article |
| `copy` | `Any` | No | The text of the article |
| `digital_publication_id` | `str` | No | Unique identifier of the digital publication this article belongs to |
| `id` | `str` | No | Unique identifier of this resource. |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | No | The URL to this article on our website |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `copy` | `Any` | No | The text of the page |
| `id` | `str` | No | Unique identifier of this resource. |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | No | The URL to this page on our website |

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
| `alt_audience_ids` | `Any` | No | Unique identifiers indicating the alternate audiences for this event |
| `alt_event_type_ids` | `Any` | No | Unique identifiers indicating the alternate types of this event |
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `audience_id` | `str` | No | Unique identifier indicating the preferred audience for this event |
| `buy_button_caption` | `Any` | No | Additional text below the ticket/registration button |
| `buy_button_text` | `Any` | No | The text used on the ticket/registration button |
| `date_display` | `Any` | No | A readable display of the event dates |
| `description` | `str` | No | All copytext of the event |
| `door_time` | `Any` | No | The time the doors open for this event |
| `end_date` | `Any` | No | The date the event ends |
| `end_time` | `Any` | No | The time the event ends |
| `entrance` | `Any` | No | Which entrance to use for this event |
| `event_host_id` | `str` | No | Unique identifier of the host (cf. |
| `event_host_title` | `Any` | No | Unique identifier of the host (cf. |
| `event_type_id` | `str` | No | Unique identifier indicating the preferred type of this event |
| `header_description` | `Any` | No | Brief description of the event displayed below the title |
| `hero_caption` | `Any` | No | Text displayed with the hero image on the event |
| `id` | `str` | No | Unique identifier of this resource. |
| `image_url` | `Any` | No | The URL of an image representing this page |
| `is_admission_required` | `bool` | No | Whether admission to the museum is required to attend this event |
| `is_after_hours` | `bool` | No | Whether the event is to be held after the museum closes |
| `is_free` | `bool` | No | Whether the event is free |
| `is_member_exclusive` | `bool` | No | Whether the event is exclusive to members of the museum |
| `is_private` | `bool` | No | Whether the event is private |
| `is_registration_required` | `bool` | No | Whether registration is required to attend the event |
| `is_sales_button_hidden` | `bool` | No | Whether the buy tickets button is hidden on the website event page |
| `is_sold_out` | `bool` | No | Whether the event is sold out |
| `is_ticketed` | `bool` | No | Whether a ticket is required to attend the event |
| `is_virtual_event` | `bool` | No | Whether the event is being held virtually |
| `join_url` | `Any` | No | URL to the membership signup page via this event |
| `layout_type` | `Any` | No | Number indicating the type of layout this event page uses |
| `list_description` | `Any` | No | One-sentence description of the event displayed in listings |
| `location` | `Any` | No | Where the event takes place |
| `program_ids` | `Any` | No | Unique identifiers indicating the programs this event is a part of |
| `program_titles` | `Any` | No | Titles of the programs this event is a part of |
| `rsvp_link` | `Any` | No | The URL to the sales site for this event |
| `search_tags` | `Any` | No | Editor-specified list of tags to aid in internal search |
| `short_description` | `Any` | No | Brief description of the event |
| `slug` | `str` | No | A string used in the URL for this event |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `start_date` | `Any` | No | The date the event begins |
| `start_time` | `Any` | No | The time the event starts |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `survey_url` | `Any` | No | URL to the survey associated with this event |
| `ticketed_event_id` | `str` | No | Unique identifier of the event in the ticketing system this website event is tied to |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `title_display` | `Any` | No | The name of this event formatted with HTML (optional) |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |
| `virtual_event_passcode` | `Any` | No | Passcode to access the virtual event |
| `virtual_event_url` | `Any` | No | URL to the virtual event |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `button_caption` | `Any` | No | Additional text below the ticket/registration button |
| `button_text` | `Any` | No | The text used on the ticket/registration button |
| `button_url` | `Any` | No | The URL to the sales site or an RSVP link for this event |
| `description` | `str` | No | Description of the event |
| `end_at` | `Any` | No | The date the event occurrence ends |
| `event_id` | `str` | No | Identifier of the master event of which this is an occurrence |
| `id` | `str` | No | Unique identifier of this resource. |
| `image_url` | `Any` | No | The URL of an image representing this page |
| `is_private` | `bool` | No | Whether the event is private. |
| `is_sales_button_hidden` | `bool` | No | Whether the buy tickets button is hidden on the website event page |
| `is_ticketed` | `bool` | No | Whether a ticket is required to attend the event |
| `location` | `Any` | No | Where the event takes place |
| `off_sale_at` | `Any` | No | Date and time the event goes off sale |
| `on_sale_at` | `Any` | No | Date and time the event goes on sale |
| `short_description` | `Any` | No | Brief description of the event |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `start_at` | `Any` | No | The date the event occurrence begins |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `title_display` | `Any` | No | The name of this event formatted with HTML (optional) |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `id` | `str` | No | Unique identifier of this resource. |
| `is_affiliate_group` | `bool` | No | Whether this program represents an affiliate group |
| `is_event_host` | `bool` | No | Whether this program represents an event host |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |

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
| `aic_end_at` | `Any` | No | Date the exhibition closed at the Art Institute of Chicago |
| `aic_start_at` | `Any` | No | Date the exhibition opened at the Art Institute of Chicago |
| `alt_image_ids` | `Any` | No | Unique identifiers of all non-preferred images of this exhibition. |
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `artist_ids` | `Any` | No | Unique identifiers of the artist agent records representing who was shown in the exhibition |
| `artwork_ids` | `Any` | No | Unique identifiers of the artworks that were part of the exhibition |
| `artwork_titles` | `Any` | No | Names of the artworks that were part of the exhibition |
| `document_ids` | `Any` | No | Unique identifiers of assets that serve as documentation for this exhibition |
| `gallery_id` | `str` | No | Unique identifier of the gallery that mainly housed the exhibition |
| `gallery_title` | `Any` | No | The name of the gallery that mainly housed the exhibition |
| `id` | `str` | No | Unique identifier of this resource. |
| `image_id` | `str` | No | Unique identifier of the preferred image to use to represent this exhibition |
| `image_url` | `Any` | No | URL to the hero image from the website |
| `is_featured` | `bool` | No | Is this exhibition currently featured on our website? |
| `is_published` | `bool` | No | Is this exhibition currently published on our website? |
| `position` | `Any` | No | Numering position represnting the order in which this exhibition is featured on the website |
| `short_description` | `Any` | No | Brief explanation of what this exhibition is |
| `site_ids` | `Any` | No | Unique identifiers of the microsites this exhibition is a part of |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `status` | `Any` | No | Whether the exhibition is open or closed |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | No | URL to this exhibition on our website |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `floor` | `Any` | No | The level the gallery is on, e.g., 1, 2, 3, or LL |
| `id` | `str` | No | Unique identifier of this resource. |
| `is_closed` | `bool` | No | Whether the gallery is currently closed |
| `latitude` | `float` | No | Latitude coordinate of the center of the room |
| `latlon` | `Any` | No | Latitude and longitude coordinates of the center of the room |
| `longitude` | `float` | No | Longitude coordinate of the center of the room |
| `number` | `Any` | No | The gallery's room number. |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | `str` | No | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `copy` | `Any` | No | The text of the page |
| `id` | `str` | No | Unique identifier of this resource. |
| `search_tags` | `Any` | No | Editor-specified list of tags to aid in internal search |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | No | The URL to this page on our website |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `copy` | `Any` | No | The text of the highlight description |
| `id` | `str` | No | Unique identifier of this resource. |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |

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
| `additional_text` | `Any` | No | Additional information about the hours |
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `friday_is_closed` | `Any` | No | Whether the museum is closed on Fridays |
| `friday_member_close` | `Any` | No | The time member hours ends on Fridays |
| `friday_member_open` | `Any` | No | The time member hours starts on Fridays |
| `friday_public_close` | `Any` | No | The time public hours ends on Fridays |
| `friday_public_open` | `Any` | No | The time public hours starts on Fridays |
| `id` | `str` | No | Unique identifier of this resource. |
| `monday_is_closed` | `Any` | No | Whether the museum is closed on Mondays |
| `monday_member_close` | `Any` | No | The time member hours ends on Mondays |
| `monday_member_open` | `Any` | No | The time member hours starts on Mondays |
| `monday_public_close` | `Any` | No | The time public hours ends on Mondays |
| `monday_public_open` | `Any` | No | The time public hours starts on Mondays |
| `saturday_is_closed` | `Any` | No | Whether the museum is closed on Saturdays |
| `saturday_member_close` | `Any` | No | The time member hours ends on Saturdays |
| `saturday_member_open` | `Any` | No | The time member hours starts on Saturdays |
| `saturday_public_close` | `Any` | No | The time public hours ends on Saturdays |
| `saturday_public_open` | `Any` | No | The time public hours starts on Saturdays |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `summary` | `Any` | No | Readable summary of the hours |
| `sunday_is_closed` | `Any` | No | Whether the museum is closed on Sundays |
| `sunday_member_close` | `Any` | No | The time member hours ends on Sundays |
| `sunday_member_open` | `Any` | No | The time member hours starts on Sundays |
| `sunday_public_close` | `Any` | No | The time public hours ends on Sundays |
| `sunday_public_open` | `Any` | No | The time public hours starts on Sundays |
| `thursday_is_closed` | `Any` | No | Whether the museum is closed on Thursdays |
| `thursday_member_close` | `Any` | No | The time member hours ends on Thursdays |
| `thursday_member_open` | `Any` | No | The time member hours starts on Thursdays |
| `thursday_public_close` | `Any` | No | The time public hours ends on Thursdays |
| `thursday_public_open` | `Any` | No | The time public hours starts on Thursdays |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `tuesday_is_closed` | `Any` | No | Whether the museum is closed on Tuesdays |
| `tuesday_member_close` | `Any` | No | The time member hours ends on Tuesdays |
| `tuesday_member_open` | `Any` | No | The time member hours starts on Tuesdays |
| `tuesday_public_close` | `Any` | No | The time public hours ends on Tuesdays |
| `tuesday_public_open` | `Any` | No | The time public hours starts on Tuesdays |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |
| `wednesday_is_closed` | `Any` | No | Whether the museum is closed on Wednesdays |
| `wednesday_member_close` | `Any` | No | The time member hours ends on Wednesdays |
| `wednesday_member_open` | `Any` | No | The time member hours starts on Wednesdays |
| `wednesday_public_close` | `Any` | No | The time public hours ends on Wednesdays |
| `wednesday_public_open` | `Any` | No | The time public hours starts on Wednesdays |

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
| `ahash` | `Any` | No | Image hash generated using ahash algorithm with 64 boolean subfields |
| `alt_text` | `Any` | No | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `artwork_ids` | `Any` | No | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `Any` | No | Names of the artworks associated with this asset |
| `color` | `Any` | No | Dominant color of this image in HSL |
| `colorfulness` | `Any` | No | Unbounded positive float representing an abstract measure of colorfulness. |
| `content` | `Any` | No | Text of or URL to the contents of this asset |
| `content_e_tag` | `Any` | No | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `Any` | No | Asset-specific copyright information |
| `fingerprint` | `Any` | No | Image hashes: aHash, dHash, pHash, wHash |
| `height` | `float` | No | Native height of the image |
| `id` | `str` | No | Unique identifier of this resource. |
| `iiif_url` | `Any` | No | IIIF URL of this image |
| `is_educational_resource` | `bool` | No | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | No | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | No | Whether this resource is considered to be educational |
| `lake_guid` | `Any` | No | Unique UUID of this resource in LAKE, our DAMS. |
| `lqip` | `Any` | No | Low-quality image placeholder (LQIP). |
| `phash` | `Any` | No | Image hash generated using phash algorithm with 64 boolean subfields |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `type` | `Any` | No | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |
| `width` | `float` | No | Native width of the image |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `copy` | `Any` | No | The text of the page |
| `id` | `str` | No | Unique identifier of this resource. |
| `search_tags` | `Any` | No | Editor-specified list of tags to aid in internal search |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | No | The URL to this page on our website |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `id` | `str` | No | Unique identifier of this resource. |
| `latitude` | `float` | No | Latitude coordinate of the center of the room |
| `longitude` | `float` | No | Longitude coordinate of the center of the room |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | `str` | No | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `copy` | `Any` | No | The text of the page |
| `id` | `str` | No | Unique identifier of this resource. |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | No | The URL to this page on our website |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `copy` | `Any` | No | The text of the page |
| `id` | `str` | No | Unique identifier of this resource. |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | No | The URL to this page on our website |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `artist_ids` | `Any` | No | Unique identifiers of the artists associated with this product |
| `artwork_ids` | `Any` | No | Unique identifiers of the artworks associated with this product |
| `description` | `str` | No | Explanation of what this product is |
| `exhibition_ids` | `Any` | No | Unique identifiers of the exhibitions associated with this product |
| `external_sku` | `Any` | No | Numeric product identification code of a machine-readable barcode, when the customer sku differs from our internal one |
| `id` | `str` | No | Unique identifier of this resource. |
| `image_url` | `Any` | No | URL of an image for this product |
| `max_compare_at_price` | `Any` | No | Number indicating how much the most expensive variant of a product cost before a sale |
| `max_current_price` | `Any` | No | Number indicating how much the most expensive variant of a product costs right now |
| `min_compare_at_price` | `Any` | No | Number indicating how much the least expensive variant of a product cost before a sale |
| `min_current_price` | `Any` | No | Number indicating how much the least expensive variant of a product costs right now |
| `price_display` | `Any` | No | Explanation of what this product is |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | No | URL of this product in the shop |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `id` | `str` | No | Unique identifier of this resource. |
| `section_ids` | `Any` | No | Unique identifiers of the sections of this publication |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | No | URL to the publication |

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
| `api_id` | `str` | No | API unique identifier |
| `api_link` | `Any` | No | URL to this recource in the API |
| `api_model` | `Any` | No | Name of the model the resource represents |
| `id` | `str` | No | Unique identifier within the search index |
| `is_boosted` | `bool` | No | Whether this record has been flagged to be boosted |
| `score` | `float` | No | Search index ranking of the result |
| `thumbnail` | `Any` | No | Metadata on the image representing this record |
| `timestamp` | `Any` | No | Date this record was last updated in the API |
| `title` | `str` | No | The name of this resource |

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
| `accession` | `Any` | No | An accession number parsed from the title or tombstone |
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `artwork_id` | `str` | No | Unique identifier of the artwork with which this section is associated |
| `content` | `Any` | No | Content of this section in plaintext |
| `generic_page_id` | `str` | No | Unique identifier of the page on the website that represents the publication this section belongs to |
| `id` | `str` | No | Unique identifier of this resource. |
| `publication_id` | `str` | No | Unique identifier of the publication this section belongs to |
| `publication_title` | `Any` | No | Name of the publication this section belongs to |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | No | URL to the section |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `artwork_ids` | `Any` | No | Unique identifiers of the artworks this site is associated with |
| `artwork_titles` | `Any` | No | Names of the artworks this site is associated with |
| `description` | `str` | No | Explanation of what this site is |
| `exhibition_ids` | `Any` | No | Unique identifier of the exhibitions this site is associated with |
| `exhibition_titles` | `Any` | No | Names of the exhibitions this site is associated with |
| `id` | `str` | No | Unique identifier of this resource. |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | No | URL to this site |

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
| `alt_text` | `Any` | No | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `artwork_ids` | `Any` | No | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `Any` | No | Names of the artworks associated with this asset |
| `content` | `Any` | No | Text of or URL to the contents of this asset |
| `content_e_tag` | `Any` | No | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `Any` | No | Asset-specific copyright information |
| `id` | `str` | No | Unique identifier of this resource. |
| `is_educational_resource` | `bool` | No | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | No | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | No | Whether this resource is considered to be educational |
| `lake_guid` | `Any` | No | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | Name of this mobile audio file – derived from the artwork and tour titles |
| `transcript` | `Any` | No | Text transcription of the audio file |
| `type` | `Any` | No | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | No | URL to the audio file |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `id` | `str` | No | Unique identifier of this resource. |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |
| `web_url` | `Any` | No | The URL to this page on our website |

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
| `alt_text` | `Any` | No | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `artwork_ids` | `Any` | No | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `Any` | No | Names of the artworks associated with this asset |
| `content` | `Any` | No | Text of or URL to the contents of this asset |
| `content_e_tag` | `Any` | No | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `Any` | No | Asset-specific copyright information |
| `id` | `str` | No | Unique identifier of this resource. |
| `is_educational_resource` | `bool` | No | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | No | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | No | Whether this resource is considered to be educational |
| `lake_guid` | `Any` | No | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `type` | `Any` | No | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |

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
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `artist_titles` | `Any` | No | Names of the artists of the artworks featured in this tour's tour stops |
| `artwork_titles` | `Any` | No | Names of the artworks featured in this tour's tour stops |
| `description` | `str` | No | Explanation of what the tour is |
| `id` | `str` | No | Unique identifier of this resource. |
| `image` | `Any` | No | The main image for the tour |
| `intro` | `Any` | No | Text introducing the tour |
| `intro_link` | `Any` | No | Link to the audio file of the introduction |
| `intro_transcript` | `Any` | No | Transcript of the introduction audio to the tour |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |
| `weight` | `float` | No | Number representing this tour's sort order |

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
| `alt_text` | `Any` | No | Alternative text for the asset to describe it to people with low or no vision |
| `api_link` | `Any` | No | REST API link for this resource |
| `api_model` | `Any` | No | REST API resource type or endpoint |
| `artwork_ids` | `Any` | No | Unique identifiers of the artworks associated with this asset |
| `artwork_titles` | `Any` | No | Names of the artworks associated with this asset |
| `content` | `Any` | No | Text of or URL to the contents of this asset |
| `content_e_tag` | `Any` | No | Arbitrary unique identifier that changes when the binary file gets updated |
| `credit_line` | `Any` | No | Asset-specific copyright information |
| `id` | `str` | No | Unique identifier of this resource. |
| `is_educational_resource` | `bool` | No | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | No | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | No | Whether this resource is considered to be educational |
| `lake_guid` | `Any` | No | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `Any` | No | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `Any` | No | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `Any` | No | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `Any` | No | Date and time the record was updated in the aggregator search index |
| `title` | `str` | No | The name of this resource |
| `type` | `Any` | No | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `Any` | No | Date and time the record was updated in the aggregator database |

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


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### `test`

In-memory mock transport for testing without a live server.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

Options above are those the model carries a default for. A feature may
also accept callback options — a `sink` to receive each record, for
instance — which have no default and are covered in the full feature
reference.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

