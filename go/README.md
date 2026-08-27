# ArtInstituteOfChicago Golang SDK



The Golang SDK for the ArtInstituteOfChicago API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

It exposes the API as capitalised, semantic **Entities** — e.g. `client.Agent(nil)` — each with the same small set of operations (`List`, `Load`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Also generated from this model: `go-cli`, `go-mcp`, `lua`, `php`, `py`, `rb`, `ts` — see
> the [top-level README](../README.md).


## Install
```bash
go get github.com/voxgig-sdk/art-institute-of-chicago-sdk/go@latest
```

The Go module proxy resolves the version from the `go/vX.Y.Z` GitHub
release tag — see [Releases](https://github.com/voxgig-sdk/art-institute-of-chicago-sdk/releases) for the available versions.

To vendor from a local checkout instead, clone this repo alongside your
project and add a `replace` directive pointing at the checked-out
`go/` directory:

```bash
go mod edit -replace github.com/voxgig-sdk/art-institute-of-chicago-sdk/go=../art-institute-of-chicago-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### Quickstart

A complete program: create a client, then call the entity operations.
Each operation returns `(value, error)` — the value is the data itself
(there is no `{ok, data}` wrapper), so check `err` and use the value
directly.

```go
package main

import (
    "fmt"
    sdk "github.com/voxgig-sdk/art-institute-of-chicago-sdk/go"
)

func main() {
    client := sdk.New()

    // List agent records — the value is the array of records itself.
    agents, err := client.Agent(nil).List(nil, nil)
    if err != nil {
        panic(err)
    }
    for _, item := range agents.([]any) {
        fmt.Println(item)
    }

    // Load a single agent — the value is the loaded record.
    agent, err := client.Agent(nil).Load(map[string]any{"id": "example_id"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(agent)
}
```


## Error handling

Every entity operation returns `(value, error)`. Check `err` before
using the value — there is no exception to catch:

```go
publications, err := client.Publication(nil).List(nil, nil)
if err != nil {
    // handle err
    return
}
_ = publications
```

`Direct` follows the same `(value, error)` convention:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example_id"},
})
if err != nil {
    // handle err
}
_ = result
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.Test()

publication, err := client.Publication(nil).List(
    nil, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(publication) // the returned mock data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewArtInstituteOfChicagoSDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
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
cd go && go test ./test/...
```


## Reference

### NewArtInstituteOfChicagoSDK

```go
func NewArtInstituteOfChicagoSDK(options map[string]any) *ArtInstituteOfChicagoSDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *ArtInstituteOfChicagoSDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### ArtInstituteOfChicagoSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `Agent` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create an Agent entity instance. |
| `AgentRole` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create an AgentRole entity instance. |
| `AgentType` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create an AgentType entity instance. |
| `Article` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create an Article entity instance. |
| `Artwork` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create an Artwork entity instance. |
| `ArtworkDateQualifier` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create an ArtworkDateQualifier entity instance. |
| `ArtworkPlaceQualifier` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create an ArtworkPlaceQualifier entity instance. |
| `ArtworkType` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create an ArtworkType entity instance. |
| `CategoryTerm` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create a CategoryTerm entity instance. |
| `DigitalPublication` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create a DigitalPublication entity instance. |
| `DigitalPublicationArticle` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create a DigitalPublicationArticle entity instance. |
| `EducatorResource` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create an EducatorResource entity instance. |
| `Event` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create an Event entity instance. |
| `EventOccurrence` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create an EventOccurrence entity instance. |
| `EventProgram` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create an EventProgram entity instance. |
| `Exhibition` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create an Exhibition entity instance. |
| `Gallery` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create a Gallery entity instance. |
| `GenericPage` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create a GenericPage entity instance. |
| `Highlight` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create a Highlight entity instance. |
| `Hour` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create a Hour entity instance. |
| `Image` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create an Image entity instance. |
| `LandingPage` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create a LandingPage entity instance. |
| `Place` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create a Place entity instance. |
| `PressRelease` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create a PressRelease entity instance. |
| `PrintedPublication` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create a PrintedPublication entity instance. |
| `Product` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create a Product entity instance. |
| `Publication` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create a Publication entity instance. |
| `Search` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create a Search entity instance. |
| `Section` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create a Section entity instance. |
| `Site` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create a Site entity instance. |
| `Sound` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create a Sound entity instance. |
| `StaticPage` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create a StaticPage entity instance. |
| `Text` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create a Text entity instance. |
| `Tour` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create a Tour entity instance. |
| `Video` | `(data map[string]any) ArtInstituteOfChicagoEntity` | Create a Video entity instance. |

### Entity interface (ArtInstituteOfChicagoEntity)

All entities implement the `ArtInstituteOfChicagoEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `Load` | `(reqmatch, ctrl map[string]any) (any, error)` | Load a single entity by match criteria. |
| `List` | `(reqmatch, ctrl map[string]any) (any, error)` | List entities matching the criteria. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(value, error)`. The `value` is the
operation's data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `Load` | the entity record (`map[string]any`) |
| `List` | a `[]any` of entity records |

Check `err` first, then use the value directly (or the typed
`...Typed` variants, which return the entity's model struct and a typed
slice):

    agent, err := client.Agent(nil).List(map[string]any{/* fields */}, nil)
    if err != nil { /* handle */ }
    // agent is the returned record

Only `Direct()` returns a response envelope — a `map[string]any` with
`"ok"`, `"status"`, `"headers"`, and `"data"` keys.

### Entities

#### Agent

| Field | Description |
| --- | --- |
| `"alt_titles"` | Alternate names for this agent |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"birth_date"` | The year this agent was born |
| `"death_date"` | The year this agent died |
| `"description"` | A biographical description of the agent |
| `"id"` | Unique identifier of this resource. |
| `"is_artist"` | Whether the agent is an artist. |
| `"sort_title"` | Sortable name for this agent, typically with last name first. |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"ulan_id"` | Unique identifier of this agent in Getty's ULAN |
| `"updated_at"` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/agents`

#### AgentRole

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"id"` | Unique identifier of this resource. |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/agent-roles`

#### AgentType

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"id"` | Unique identifier of this resource. |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/agent-types`

#### Article

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"copy"` | The text of the article |
| `"id"` | Unique identifier of this resource. |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/articles`

#### Artwork

| Field | Description |
| --- | --- |
| `"alt_artist_ids"` | Unique identifiers of the non-preferred artists/cultures associated with this work |
| `"alt_classification_ids"` | Unique identifiers of all other non-preferred classification terms for this work |
| `"alt_image_ids"` | Unique identifiers of all non-preferred images of this work. |
| `"alt_material_ids"` | Unique identifiers of all other non-preferred material terms for this work |
| `"alt_style_ids"` | Unique identifiers of all other non-preferred style terms for this work |
| `"alt_subject_ids"` | Unique identifiers of all other non-preferred subject terms for this work |
| `"alt_technique_ids"` | Unique identifiers of all other non-preferred technique terms for this work |
| `"alt_titles"` | Alternate names for this work |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"artist_display"` | Readable description of the creator of this work. |
| `"artist_id"` | Unique identifier of the preferred artist/culture associated with this work |
| `"artist_ids"` | Unique identifier of all artist/cultures associated with this work |
| `"artist_title"` | Name of the preferred artist/culture associated with this work |
| `"artist_titles"` | Names of all artist/cultures associated with this work |
| `"artwork_type_id"` | Unique identifier of the kind of object or work |
| `"artwork_type_title"` | The kind of object or work (e.g. |
| `"boost_rank"` | Manual indication of what rank this artwork should take in search results. |
| `"catalog_based_search_keyword_titles"` | The keyword search values that would be catalog-based searches on this record |
| `"catalogue_display"` | Brief text listing all the catalogues raisonnés which include this work. |
| `"category_ids"` | Unique identifiers of the categories this work is a part of |
| `"category_titles"` | Names of the categories this artwork is a part of |
| `"classification_id"` | Unique identifier of the preferred classification term for this work |
| `"classification_ids"` | Unique identifiers of all classification terms for this work |
| `"classification_title"` | The name of the preferred classification term for this work |
| `"classification_titles"` | The names of all classification terms related to this artwork |
| `"color"` | Dominant color of this artwork in HSL |
| `"colorfulness"` | Unbounded positive float representing an abstract measure of colorfulness. |
| `"copyright_notice"` | Statement notifying how the work is protected by copyright. |
| `"credit_line"` | Brief statement indicating how the work came into the collection |
| `"date_display"` | Readable, free-text description of the period of time associated with the creation of this work. |
| `"date_end"` | The year of the period of time associated with the creation of this work |
| `"date_qualifier_id"` | Unique identifier of the qualifer to the dates provided for this record. |
| `"date_qualifier_title"` | Readable, text qualifer to the dates provided for this record. |
| `"date_start"` | The year of the period of time associated with the creation of this work |
| `"department_id"` | Unique identifier of the curatorial department that this work belongs to |
| `"department_title"` | Name of the curatorial department that this work belongs to |
| `"description"` | Longer explanation describing the work |
| `"dimensions"` | The size, shape, scale, and dimensions of the work. |
| `"dimensions_detail"` | The height, width, depth, and/or diameter of each section of the work in centimeters |
| `"document_ids"` | Unique identifiers of assets that serve as documentation for this artwork |
| `"edition"` | Edition number if the work is one of many |
| `"exhibition_history"` | List of all the places this work has been exhibited |
| `"fiscal_year"` | The fiscal year in which the work was acquired. |
| `"fiscal_year_deaccession"` | The fiscal year in which the work was deaccessioned. |
| `"gallery_id"` | Unique identifier of the location of this work in our museum |
| `"gallery_title"` | The location of this work in our museum |
| `"has_advanced_imaging"` | Whether this artwork is enhanced with 3D models, 360 image sequences, Mirador views, etc. |
| `"has_educational_resources"` | Whether this artwork has any documents tagged as educational |
| `"has_multimedia_resources"` | Whether this artwork has any associated microsites, digital publications, or documents tagged as multimedia |
| `"has_not_been_viewed_much"` | Whether the artwork hasn't been visited on our website very much |
| `"id"` | Unique identifier of this resource. |
| `"image_embedding"` | The generated embeddings describing the artwork image |
| `"image_id"` | Unique identifier of the preferred image to use to represent this work |
| `"inscriptions"` | A description of distinguishing or identifying physical markings that are on the work |
| `"internal_department_id"` | An internal department id we use for analytics. |
| `"is_boosted"` | Whether this document should be boosted in search |
| `"is_on_view"` | Whether the work is on display |
| `"is_public_domain"` | Whether the work is in the public domain, meaning it was created before copyrights existed or has left the copyright term |
| `"is_zoomable"` | Whether images of the work are allowed to be displayed in a zoomable interface. |
| `"latitude"` | Latitude coordinate of the location of this work in our galleries |
| `"latlon"` | Latitude and longitude coordinates of the location of this work in our galleries |
| `"longitude"` | Longitude coordinate of the location of this work in our galleries |
| `"main_reference_number"` | Unique identifier assigned to the artwork upon acquisition |
| `"material_id"` | Unique identifier of the preferred material term for this work |
| `"material_ids"` | Unique identifiers of all material terms for this work |
| `"material_titles"` | The names of all material terms related to this artwork |
| `"max_zoom_window_size"` | The maximum size of the window the image is allowed to be viewed in, in pixels. |
| `"medium_display"` | The substances or materials used in the creation of a work |
| `"nomisma_id"` | Unique identifier of this work in the nomisma coin database |
| `"on_loan_display"` | If an artwork is on loan, this contains details about the loan |
| `"pageviews"` | Approx. |
| `"pageviews_recent"` | Approx. |
| `"place_of_origin"` | The location where the creation, design, or production of the work took place, or the original location of the work |
| `"provenance_text"` | Ownership/collecting history of the work. |
| `"publication_history"` | Bibliographic list of all the places this work has been published |
| `"publishing_verification_level"` | Indicator of how much metadata on the work in published. |
| `"section_ids"` | Unique identifiers of the digital publication chapters this work in included in |
| `"section_titles"` | Names of the digital publication chapters this work is included in |
| `"short_description"` | Short explanation describing the work |
| `"site_ids"` | Unique identifiers of the microsites this work is a part of |
| `"sound_ids"` | Unique identifiers of the audio about this work |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"style_id"` | Unique identifier of the preferred style term for this work |
| `"style_ids"` | Unique identifiers of all style terms for this work |
| `"style_title"` | The name of the preferred style term for this work |
| `"style_titles"` | The names of all style terms related to this artwork |
| `"subject_id"` | Unique identifier of the preferred subject term for this work |
| `"subject_ids"` | Unique identifiers of all subject terms for this work |
| `"subject_titles"` | The names of all subject terms related to this artwork |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"technique_id"` | Unique identifier of the preferred technique term for this work |
| `"technique_ids"` | Unique identifiers of all technique terms for this work |
| `"technique_titles"` | The names of all technique terms related to this artwork |
| `"term_titles"` | The names of the taxonomy tags for this work |
| `"text_embedding"` | The generated embeddings of artwork text |
| `"text_ids"` | Unique identifiers of the texts about this work |
| `"theme_titles"` | The names of all thematic publish categories related to this artwork |
| `"thumbnail"` | Metadata about the image referenced by `image_id`. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |
| `"video_ids"` | Unique identifiers of the videos about this work |

Operations: List, Load.

API path: `/artworks`

#### ArtworkDateQualifier

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"id"` | Unique identifier of this resource. |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/artwork-date-qualifiers`

#### ArtworkPlaceQualifier

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"id"` | Unique identifier of this resource. |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/artwork-place-qualifiers`

#### ArtworkType

| Field | Description |
| --- | --- |
| `"aat_id"` | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"id"` | Unique identifier of this resource. |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/artwork-types`

#### CategoryTerm

| Field | Description |
| --- | --- |
| `"aat_id"` | Identifier of reconciled (most similar) term in the Getty's Art and Architecture Thesaurus (AAT) |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"id"` | Unique identifier of this resource. |
| `"parent_id"` | Unique identifier of this category's parent |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"subtype"` | Takes one of the following values: classification, material, technique, style, subject, department, theme |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/category-terms`

#### DigitalPublication

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"copy"` | The text of the page |
| `"id"` | Unique identifier of this resource. |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |
| `"web_url"` | The URL to this page on our website |

Operations: List, Load.

API path: `/digital-publications`

#### DigitalPublicationArticle

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"author_display"` | A display-friendly text of the authors of this article |
| `"copy"` | The text of the article |
| `"digital_publication_id"` | Unique identifier of the digital publication this article belongs to |
| `"id"` | Unique identifier of this resource. |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |
| `"web_url"` | The URL to this article on our website |

Operations: List, Load.

API path: `/digital-publication-articles`

#### EducatorResource

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"copy"` | The text of the page |
| `"id"` | Unique identifier of this resource. |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |
| `"web_url"` | The URL to this page on our website |

Operations: List, Load.

API path: `/educator-resources`

#### Event

| Field | Description |
| --- | --- |
| `"alt_audience_ids"` | Unique identifiers indicating the alternate audiences for this event |
| `"alt_event_type_ids"` | Unique identifiers indicating the alternate types of this event |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"audience_id"` | Unique identifier indicating the preferred audience for this event |
| `"buy_button_caption"` | Additional text below the ticket/registration button |
| `"buy_button_text"` | The text used on the ticket/registration button |
| `"date_display"` | A readable display of the event dates |
| `"description"` | All copytext of the event |
| `"door_time"` | The time the doors open for this event |
| `"end_date"` | The date the event ends |
| `"end_time"` | The time the event ends |
| `"entrance"` | Which entrance to use for this event |
| `"event_host_id"` | Unique identifier of the host (cf. |
| `"event_host_title"` | Unique identifier of the host (cf. |
| `"event_type_id"` | Unique identifier indicating the preferred type of this event |
| `"header_description"` | Brief description of the event displayed below the title |
| `"hero_caption"` | Text displayed with the hero image on the event |
| `"id"` | Unique identifier of this resource. |
| `"image_url"` | The URL of an image representing this page |
| `"is_admission_required"` | Whether admission to the museum is required to attend this event |
| `"is_after_hours"` | Whether the event is to be held after the museum closes |
| `"is_free"` | Whether the event is free |
| `"is_member_exclusive"` | Whether the event is exclusive to members of the museum |
| `"is_private"` | Whether the event is private |
| `"is_registration_required"` | Whether registration is required to attend the event |
| `"is_sales_button_hidden"` | Whether the buy tickets button is hidden on the website event page |
| `"is_sold_out"` | Whether the event is sold out |
| `"is_ticketed"` | Whether a ticket is required to attend the event |
| `"is_virtual_event"` | Whether the event is being held virtually |
| `"join_url"` | URL to the membership signup page via this event |
| `"layout_type"` | Number indicating the type of layout this event page uses |
| `"list_description"` | One-sentence description of the event displayed in listings |
| `"location"` | Where the event takes place |
| `"program_ids"` | Unique identifiers indicating the programs this event is a part of |
| `"program_titles"` | Titles of the programs this event is a part of |
| `"rsvp_link"` | The URL to the sales site for this event |
| `"search_tags"` | Editor-specified list of tags to aid in internal search |
| `"short_description"` | Brief description of the event |
| `"slug"` | A string used in the URL for this event |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"start_date"` | The date the event begins |
| `"start_time"` | The time the event starts |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"survey_url"` | URL to the survey associated with this event |
| `"ticketed_event_id"` | Unique identifier of the event in the ticketing system this website event is tied to |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"title_display"` | The name of this event formatted with HTML (optional) |
| `"updated_at"` | Date and time the record was updated in the aggregator database |
| `"virtual_event_passcode"` | Passcode to access the virtual event |
| `"virtual_event_url"` | URL to the virtual event |

Operations: List, Load.

API path: `/events`

#### EventOccurrence

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"button_caption"` | Additional text below the ticket/registration button |
| `"button_text"` | The text used on the ticket/registration button |
| `"button_url"` | The URL to the sales site or an RSVP link for this event |
| `"description"` | Description of the event |
| `"end_at"` | The date the event occurrence ends |
| `"event_id"` | Identifier of the master event of which this is an occurrence |
| `"id"` | Unique identifier of this resource. |
| `"image_url"` | The URL of an image representing this page |
| `"is_private"` | Whether the event is private. |
| `"is_sales_button_hidden"` | Whether the buy tickets button is hidden on the website event page |
| `"is_ticketed"` | Whether a ticket is required to attend the event |
| `"location"` | Where the event takes place |
| `"off_sale_at"` | Date and time the event goes off sale |
| `"on_sale_at"` | Date and time the event goes on sale |
| `"short_description"` | Brief description of the event |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"start_at"` | The date the event occurrence begins |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"title_display"` | The name of this event formatted with HTML (optional) |
| `"updated_at"` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/event-occurrences`

#### EventProgram

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"id"` | Unique identifier of this resource. |
| `"is_affiliate_group"` | Whether this program represents an affiliate group |
| `"is_event_host"` | Whether this program represents an event host |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/event-programs`

#### Exhibition

| Field | Description |
| --- | --- |
| `"aic_end_at"` | Date the exhibition closed at the Art Institute of Chicago |
| `"aic_start_at"` | Date the exhibition opened at the Art Institute of Chicago |
| `"alt_image_ids"` | Unique identifiers of all non-preferred images of this exhibition. |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"artist_ids"` | Unique identifiers of the artist agent records representing who was shown in the exhibition |
| `"artwork_ids"` | Unique identifiers of the artworks that were part of the exhibition |
| `"artwork_titles"` | Names of the artworks that were part of the exhibition |
| `"document_ids"` | Unique identifiers of assets that serve as documentation for this exhibition |
| `"gallery_id"` | Unique identifier of the gallery that mainly housed the exhibition |
| `"gallery_title"` | The name of the gallery that mainly housed the exhibition |
| `"id"` | Unique identifier of this resource. |
| `"image_id"` | Unique identifier of the preferred image to use to represent this exhibition |
| `"image_url"` | URL to the hero image from the website |
| `"is_featured"` | Is this exhibition currently featured on our website? |
| `"is_published"` | Is this exhibition currently published on our website? |
| `"position"` | Numering position represnting the order in which this exhibition is featured on the website |
| `"short_description"` | Brief explanation of what this exhibition is |
| `"site_ids"` | Unique identifiers of the microsites this exhibition is a part of |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"status"` | Whether the exhibition is open or closed |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |
| `"web_url"` | URL to this exhibition on our website |

Operations: List, Load.

API path: `/exhibitions`

#### Gallery

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"floor"` | The level the gallery is on, e.g., 1, 2, 3, or LL |
| `"id"` | Unique identifier of this resource. |
| `"is_closed"` | Whether the gallery is currently closed |
| `"latitude"` | Latitude coordinate of the center of the room |
| `"latlon"` | Latitude and longitude coordinates of the center of the room |
| `"longitude"` | Longitude coordinate of the center of the room |
| `"number"` | The gallery's room number. |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"tgn_id"` | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/galleries`

#### GenericPage

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"copy"` | The text of the page |
| `"id"` | Unique identifier of this resource. |
| `"search_tags"` | Editor-specified list of tags to aid in internal search |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |
| `"web_url"` | The URL to this page on our website |

Operations: List, Load.

API path: `/generic-pages`

#### Highlight

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"copy"` | The text of the highlight description |
| `"id"` | Unique identifier of this resource. |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/highlights`

#### Hour

| Field | Description |
| --- | --- |
| `"additional_text"` | Additional information about the hours |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"friday_is_closed"` | Whether the museum is closed on Fridays |
| `"friday_member_close"` | The time member hours ends on Fridays |
| `"friday_member_open"` | The time member hours starts on Fridays |
| `"friday_public_close"` | The time public hours ends on Fridays |
| `"friday_public_open"` | The time public hours starts on Fridays |
| `"id"` | Unique identifier of this resource. |
| `"monday_is_closed"` | Whether the museum is closed on Mondays |
| `"monday_member_close"` | The time member hours ends on Mondays |
| `"monday_member_open"` | The time member hours starts on Mondays |
| `"monday_public_close"` | The time public hours ends on Mondays |
| `"monday_public_open"` | The time public hours starts on Mondays |
| `"saturday_is_closed"` | Whether the museum is closed on Saturdays |
| `"saturday_member_close"` | The time member hours ends on Saturdays |
| `"saturday_member_open"` | The time member hours starts on Saturdays |
| `"saturday_public_close"` | The time public hours ends on Saturdays |
| `"saturday_public_open"` | The time public hours starts on Saturdays |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"summary"` | Readable summary of the hours |
| `"sunday_is_closed"` | Whether the museum is closed on Sundays |
| `"sunday_member_close"` | The time member hours ends on Sundays |
| `"sunday_member_open"` | The time member hours starts on Sundays |
| `"sunday_public_close"` | The time public hours ends on Sundays |
| `"sunday_public_open"` | The time public hours starts on Sundays |
| `"thursday_is_closed"` | Whether the museum is closed on Thursdays |
| `"thursday_member_close"` | The time member hours ends on Thursdays |
| `"thursday_member_open"` | The time member hours starts on Thursdays |
| `"thursday_public_close"` | The time public hours ends on Thursdays |
| `"thursday_public_open"` | The time public hours starts on Thursdays |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"tuesday_is_closed"` | Whether the museum is closed on Tuesdays |
| `"tuesday_member_close"` | The time member hours ends on Tuesdays |
| `"tuesday_member_open"` | The time member hours starts on Tuesdays |
| `"tuesday_public_close"` | The time public hours ends on Tuesdays |
| `"tuesday_public_open"` | The time public hours starts on Tuesdays |
| `"updated_at"` | Date and time the record was updated in the aggregator database |
| `"wednesday_is_closed"` | Whether the museum is closed on Wednesdays |
| `"wednesday_member_close"` | The time member hours ends on Wednesdays |
| `"wednesday_member_open"` | The time member hours starts on Wednesdays |
| `"wednesday_public_close"` | The time public hours ends on Wednesdays |
| `"wednesday_public_open"` | The time public hours starts on Wednesdays |

Operations: List, Load.

API path: `/hours`

#### Image

| Field | Description |
| --- | --- |
| `"ahash"` | Image hash generated using ahash algorithm with 64 boolean subfields |
| `"alt_text"` | Alternative text for the asset to describe it to people with low or no vision |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"artwork_ids"` | Unique identifiers of the artworks associated with this asset |
| `"artwork_titles"` | Names of the artworks associated with this asset |
| `"color"` | Dominant color of this image in HSL |
| `"colorfulness"` | Unbounded positive float representing an abstract measure of colorfulness. |
| `"content"` | Text of or URL to the contents of this asset |
| `"content_e_tag"` | Arbitrary unique identifier that changes when the binary file gets updated |
| `"credit_line"` | Asset-specific copyright information |
| `"fingerprint"` | Image hashes: aHash, dHash, pHash, wHash |
| `"height"` | Native height of the image |
| `"id"` | Unique identifier of this resource. |
| `"iiif_url"` | IIIF URL of this image |
| `"is_educational_resource"` | Whether this resource is considered to be educational |
| `"is_multimedia_resource"` | Whether this resource is considered to be multimedia |
| `"is_teacher_resource"` | Whether this resource is considered to be educational |
| `"lake_guid"` | Unique UUID of this resource in LAKE, our DAMS. |
| `"lqip"` | Low-quality image placeholder (LQIP). |
| `"phash"` | Image hash generated using phash algorithm with 64 boolean subfields |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"type"` | Type always takes one of the following values: image, sound, text, video |
| `"updated_at"` | Date and time the record was updated in the aggregator database |
| `"width"` | Native width of the image |

Operations: List, Load.

API path: `/images`

#### LandingPage

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"copy"` | The text of the page |
| `"id"` | Unique identifier of this resource. |
| `"search_tags"` | Editor-specified list of tags to aid in internal search |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |
| `"web_url"` | The URL to this page on our website |

Operations: List, Load.

API path: `/landing-pages`

#### Place

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"id"` | Unique identifier of this resource. |
| `"latitude"` | Latitude coordinate of the center of the room |
| `"longitude"` | Longitude coordinate of the center of the room |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"tgn_id"` | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/places`

#### PressRelease

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"copy"` | The text of the page |
| `"id"` | Unique identifier of this resource. |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |
| `"web_url"` | The URL to this page on our website |

Operations: List, Load.

API path: `/press-releases`

#### PrintedPublication

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"copy"` | The text of the page |
| `"id"` | Unique identifier of this resource. |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |
| `"web_url"` | The URL to this page on our website |

Operations: List, Load.

API path: `/printed-publications`

#### Product

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"artist_ids"` | Unique identifiers of the artists associated with this product |
| `"artwork_ids"` | Unique identifiers of the artworks associated with this product |
| `"description"` | Explanation of what this product is |
| `"exhibition_ids"` | Unique identifiers of the exhibitions associated with this product |
| `"external_sku"` | Numeric product identification code of a machine-readable barcode, when the customer sku differs from our internal one |
| `"id"` | Unique identifier of this resource. |
| `"image_url"` | URL of an image for this product |
| `"max_compare_at_price"` | Number indicating how much the most expensive variant of a product cost before a sale |
| `"max_current_price"` | Number indicating how much the most expensive variant of a product costs right now |
| `"min_compare_at_price"` | Number indicating how much the least expensive variant of a product cost before a sale |
| `"min_current_price"` | Number indicating how much the least expensive variant of a product costs right now |
| `"price_display"` | Explanation of what this product is |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |
| `"web_url"` | URL of this product in the shop |

Operations: List, Load.

API path: `/products`

#### Publication

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"id"` | Unique identifier of this resource. |
| `"section_ids"` | Unique identifiers of the sections of this publication |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |
| `"web_url"` | URL to the publication |

Operations: List, Load.

API path: `/publications`

#### Search

| Field | Description |
| --- | --- |
| `"api_id"` | API unique identifier |
| `"api_link"` | URL to this recource in the API |
| `"api_model"` | Name of the model the resource represents |
| `"id"` | Unique identifier within the search index |
| `"is_boosted"` | Whether this record has been flagged to be boosted |
| `"score"` | Search index ranking of the result |
| `"thumbnail"` | Metadata on the image representing this record |
| `"timestamp"` | Date this record was last updated in the API |
| `"title"` | The name of this resource |

Operations: List.

API path: `/agents/search`

#### Section

| Field | Description |
| --- | --- |
| `"accession"` | An accession number parsed from the title or tombstone |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"artwork_id"` | Unique identifier of the artwork with which this section is associated |
| `"content"` | Content of this section in plaintext |
| `"generic_page_id"` | Unique identifier of the page on the website that represents the publication this section belongs to |
| `"id"` | Unique identifier of this resource. |
| `"publication_id"` | Unique identifier of the publication this section belongs to |
| `"publication_title"` | Name of the publication this section belongs to |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |
| `"web_url"` | URL to the section |

Operations: List, Load.

API path: `/sections`

#### Site

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"artwork_ids"` | Unique identifiers of the artworks this site is associated with |
| `"artwork_titles"` | Names of the artworks this site is associated with |
| `"description"` | Explanation of what this site is |
| `"exhibition_ids"` | Unique identifier of the exhibitions this site is associated with |
| `"exhibition_titles"` | Names of the exhibitions this site is associated with |
| `"id"` | Unique identifier of this resource. |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |
| `"web_url"` | URL to this site |

Operations: List, Load.

API path: `/sites`

#### Sound

| Field | Description |
| --- | --- |
| `"alt_text"` | Alternative text for the asset to describe it to people with low or no vision |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"artwork_ids"` | Unique identifiers of the artworks associated with this asset |
| `"artwork_titles"` | Names of the artworks associated with this asset |
| `"content"` | Text of or URL to the contents of this asset |
| `"content_e_tag"` | Arbitrary unique identifier that changes when the binary file gets updated |
| `"credit_line"` | Asset-specific copyright information |
| `"id"` | Unique identifier of this resource. |
| `"is_educational_resource"` | Whether this resource is considered to be educational |
| `"is_multimedia_resource"` | Whether this resource is considered to be multimedia |
| `"is_teacher_resource"` | Whether this resource is considered to be educational |
| `"lake_guid"` | Unique UUID of this resource in LAKE, our DAMS. |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | Name of this mobile audio file – derived from the artwork and tour titles |
| `"transcript"` | Text transcription of the audio file |
| `"type"` | Type always takes one of the following values: image, sound, text, video |
| `"updated_at"` | Date and time the record was updated in the aggregator database |
| `"web_url"` | URL to the audio file |

Operations: List, Load.

API path: `/mobile-sounds`

#### StaticPage

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"id"` | Unique identifier of this resource. |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |
| `"web_url"` | The URL to this page on our website |

Operations: List, Load.

API path: `/static-pages`

#### Text

| Field | Description |
| --- | --- |
| `"alt_text"` | Alternative text for the asset to describe it to people with low or no vision |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"artwork_ids"` | Unique identifiers of the artworks associated with this asset |
| `"artwork_titles"` | Names of the artworks associated with this asset |
| `"content"` | Text of or URL to the contents of this asset |
| `"content_e_tag"` | Arbitrary unique identifier that changes when the binary file gets updated |
| `"credit_line"` | Asset-specific copyright information |
| `"id"` | Unique identifier of this resource. |
| `"is_educational_resource"` | Whether this resource is considered to be educational |
| `"is_multimedia_resource"` | Whether this resource is considered to be multimedia |
| `"is_teacher_resource"` | Whether this resource is considered to be educational |
| `"lake_guid"` | Unique UUID of this resource in LAKE, our DAMS. |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"type"` | Type always takes one of the following values: image, sound, text, video |
| `"updated_at"` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/texts`

#### Tour

| Field | Description |
| --- | --- |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"artist_titles"` | Names of the artists of the artworks featured in this tour's tour stops |
| `"artwork_titles"` | Names of the artworks featured in this tour's tour stops |
| `"description"` | Explanation of what the tour is |
| `"id"` | Unique identifier of this resource. |
| `"image"` | The main image for the tour |
| `"intro"` | Text introducing the tour |
| `"intro_link"` | Link to the audio file of the introduction |
| `"intro_transcript"` | Transcript of the introduction audio to the tour |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"updated_at"` | Date and time the record was updated in the aggregator database |
| `"weight"` | Number representing this tour's sort order |

Operations: List, Load.

API path: `/tours`

#### Video

| Field | Description |
| --- | --- |
| `"alt_text"` | Alternative text for the asset to describe it to people with low or no vision |
| `"api_link"` | REST API link for this resource |
| `"api_model"` | REST API resource type or endpoint |
| `"artwork_ids"` | Unique identifiers of the artworks associated with this asset |
| `"artwork_titles"` | Names of the artworks associated with this asset |
| `"content"` | Text of or URL to the contents of this asset |
| `"content_e_tag"` | Arbitrary unique identifier that changes when the binary file gets updated |
| `"credit_line"` | Asset-specific copyright information |
| `"id"` | Unique identifier of this resource. |
| `"is_educational_resource"` | Whether this resource is considered to be educational |
| `"is_multimedia_resource"` | Whether this resource is considered to be multimedia |
| `"is_teacher_resource"` | Whether this resource is considered to be educational |
| `"lake_guid"` | Unique UUID of this resource in LAKE, our DAMS. |
| `"source_updated_at"` | Date and time the resource was updated in the source system |
| `"suggest_autocomplete_all"` | Internal field to power the `/autosuggest` endpoint. |
| `"suggest_autocomplete_boosted"` | Internal field to power the `/autocomplete` endpoint. |
| `"timestamp"` | Date and time the record was updated in the aggregator search index |
| `"title"` | The name of this resource |
| `"type"` | Type always takes one of the following values: image, sound, text, video |
| `"updated_at"` | Date and time the record was updated in the aggregator database |

Operations: List, Load.

API path: `/videos`



## Entities


### Agent

Create an instance: `agent := client.Agent(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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
| `is_artist` | `bool` | Whether the agent is an artist. |
| `sort_title` | `any` | Sortable name for this agent, typically with last name first. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `ulan_id` | `string` | Unique identifier of this agent in Getty's ULAN |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |

#### Example: Load

```go
agent, err := client.Agent(nil).Load(map[string]any{"id": "agent_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(agent) // the loaded record
```

#### Example: List

```go
agents, err := client.Agent(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(agents) // the array of records
```


### AgentRole

Create an instance: `agentRole := client.AgentRole(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
agentRole, err := client.AgentRole(nil).Load(map[string]any{"id": "agent_role_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(agentRole) // the loaded record
```

#### Example: List

```go
agentRoles, err := client.AgentRole(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(agentRoles) // the array of records
```


### AgentType

Create an instance: `agentType := client.AgentType(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
agentType, err := client.AgentType(nil).Load(map[string]any{"id": "agent_type_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(agentType) // the loaded record
```

#### Example: List

```go
agentTypes, err := client.AgentType(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(agentTypes) // the array of records
```


### Article

Create an instance: `article := client.Article(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
article, err := client.Article(nil).Load(map[string]any{"id": "article_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(article) // the loaded record
```

#### Example: List

```go
articles, err := client.Article(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(articles) // the array of records
```


### Artwork

Create an instance: `artwork := client.Artwork(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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
| `has_advanced_imaging` | `bool` | Whether this artwork is enhanced with 3D models, 360 image sequences, Mirador views, etc. |
| `has_educational_resources` | `bool` | Whether this artwork has any documents tagged as educational |
| `has_multimedia_resources` | `bool` | Whether this artwork has any associated microsites, digital publications, or documents tagged as multimedia |
| `has_not_been_viewed_much` | `bool` | Whether the artwork hasn't been visited on our website very much |
| `id` | `string` | Unique identifier of this resource. |
| `image_embedding` | `any` | The generated embeddings describing the artwork image |
| `image_id` | `string` | Unique identifier of the preferred image to use to represent this work |
| `inscriptions` | `any` | A description of distinguishing or identifying physical markings that are on the work |
| `internal_department_id` | `string` | An internal department id we use for analytics. |
| `is_boosted` | `bool` | Whether this document should be boosted in search |
| `is_on_view` | `bool` | Whether the work is on display |
| `is_public_domain` | `bool` | Whether the work is in the public domain, meaning it was created before copyrights existed or has left the copyright term |
| `is_zoomable` | `bool` | Whether images of the work are allowed to be displayed in a zoomable interface. |
| `latitude` | `float64` | Latitude coordinate of the location of this work in our galleries |
| `latlon` | `any` | Latitude and longitude coordinates of the location of this work in our galleries |
| `longitude` | `float64` | Longitude coordinate of the location of this work in our galleries |
| `main_reference_number` | `int` | Unique identifier assigned to the artwork upon acquisition |
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

```go
artwork, err := client.Artwork(nil).Load(map[string]any{"id": "artwork_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(artwork) // the loaded record
```

#### Example: List

```go
artworks, err := client.Artwork(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(artworks) // the array of records
```


### ArtworkDateQualifier

Create an instance: `artworkDateQualifier := client.ArtworkDateQualifier(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
artworkDateQualifier, err := client.ArtworkDateQualifier(nil).Load(map[string]any{"id": "artwork_date_qualifier_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(artworkDateQualifier) // the loaded record
```

#### Example: List

```go
artworkDateQualifiers, err := client.ArtworkDateQualifier(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(artworkDateQualifiers) // the array of records
```


### ArtworkPlaceQualifier

Create an instance: `artworkPlaceQualifier := client.ArtworkPlaceQualifier(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
artworkPlaceQualifier, err := client.ArtworkPlaceQualifier(nil).Load(map[string]any{"id": "artwork_place_qualifier_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(artworkPlaceQualifier) // the loaded record
```

#### Example: List

```go
artworkPlaceQualifiers, err := client.ArtworkPlaceQualifier(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(artworkPlaceQualifiers) // the array of records
```


### ArtworkType

Create an instance: `artworkType := client.ArtworkType(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
artworkType, err := client.ArtworkType(nil).Load(map[string]any{"id": "artwork_type_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(artworkType) // the loaded record
```

#### Example: List

```go
artworkTypes, err := client.ArtworkType(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(artworkTypes) // the array of records
```


### CategoryTerm

Create an instance: `categoryTerm := client.CategoryTerm(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
categoryTerm, err := client.CategoryTerm(nil).Load(map[string]any{"id": "category_term_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(categoryTerm) // the loaded record
```

#### Example: List

```go
categoryTerms, err := client.CategoryTerm(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(categoryTerms) // the array of records
```


### DigitalPublication

Create an instance: `digitalPublication := client.DigitalPublication(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
digitalPublication, err := client.DigitalPublication(nil).Load(map[string]any{"id": "digital_publication_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(digitalPublication) // the loaded record
```

#### Example: List

```go
digitalPublications, err := client.DigitalPublication(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(digitalPublications) // the array of records
```


### DigitalPublicationArticle

Create an instance: `digitalPublicationArticle := client.DigitalPublicationArticle(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
digitalPublicationArticle, err := client.DigitalPublicationArticle(nil).Load(map[string]any{"id": "digital_publication_article_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(digitalPublicationArticle) // the loaded record
```

#### Example: List

```go
digitalPublicationArticles, err := client.DigitalPublicationArticle(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(digitalPublicationArticles) // the array of records
```


### EducatorResource

Create an instance: `educatorResource := client.EducatorResource(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
educatorResource, err := client.EducatorResource(nil).Load(map[string]any{"id": "educator_resource_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(educatorResource) // the loaded record
```

#### Example: List

```go
educatorResources, err := client.EducatorResource(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(educatorResources) // the array of records
```


### Event

Create an instance: `event := client.Event(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
event, err := client.Event(nil).Load(map[string]any{"id": "event_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(event) // the loaded record
```

#### Example: List

```go
events, err := client.Event(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(events) // the array of records
```


### EventOccurrence

Create an instance: `eventOccurrence := client.EventOccurrence(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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
| `is_private` | `bool` | Whether the event is private. |
| `is_sales_button_hidden` | `bool` | Whether the buy tickets button is hidden on the website event page |
| `is_ticketed` | `bool` | Whether a ticket is required to attend the event |
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

```go
eventOccurrence, err := client.EventOccurrence(nil).Load(map[string]any{"id": "event_occurrence_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(eventOccurrence) // the loaded record
```

#### Example: List

```go
eventOccurrences, err := client.EventOccurrence(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(eventOccurrences) // the array of records
```


### EventProgram

Create an instance: `eventProgram := client.EventProgram(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `is_affiliate_group` | `bool` | Whether this program represents an affiliate group |
| `is_event_host` | `bool` | Whether this program represents an event host |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |

#### Example: Load

```go
eventProgram, err := client.EventProgram(nil).Load(map[string]any{"id": "event_program_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(eventProgram) // the loaded record
```

#### Example: List

```go
eventPrograms, err := client.EventProgram(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(eventPrograms) // the array of records
```


### Exhibition

Create an instance: `exhibition := client.Exhibition(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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
| `is_featured` | `bool` | Is this exhibition currently featured on our website? |
| `is_published` | `bool` | Is this exhibition currently published on our website? |
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

```go
exhibition, err := client.Exhibition(nil).Load(map[string]any{"id": "exhibition_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(exhibition) // the loaded record
```

#### Example: List

```go
exhibitions, err := client.Exhibition(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(exhibitions) // the array of records
```


### Gallery

Create an instance: `gallery := client.Gallery(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `floor` | `any` | The level the gallery is on, e.g., 1, 2, 3, or LL |
| `id` | `string` | Unique identifier of this resource. |
| `is_closed` | `bool` | Whether the gallery is currently closed |
| `latitude` | `float64` | Latitude coordinate of the center of the room |
| `latlon` | `any` | Latitude and longitude coordinates of the center of the room |
| `longitude` | `float64` | Longitude coordinate of the center of the room |
| `number` | `any` | The gallery's room number. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | `string` | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |

#### Example: Load

```go
gallery, err := client.Gallery(nil).Load(map[string]any{"id": "gallery_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(gallery) // the loaded record
```

#### Example: List

```go
gallerys, err := client.Gallery(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(gallerys) // the array of records
```


### GenericPage

Create an instance: `genericPage := client.GenericPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
genericPage, err := client.GenericPage(nil).Load(map[string]any{"id": "generic_page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(genericPage) // the loaded record
```

#### Example: List

```go
genericPages, err := client.GenericPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(genericPages) // the array of records
```


### Highlight

Create an instance: `highlight := client.Highlight(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
highlight, err := client.Highlight(nil).Load(map[string]any{"id": "highlight_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(highlight) // the loaded record
```

#### Example: List

```go
highlights, err := client.Highlight(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(highlights) // the array of records
```


### Hour

Create an instance: `hour := client.Hour(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
hour, err := client.Hour(nil).Load(map[string]any{"id": "hour_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(hour) // the loaded record
```

#### Example: List

```go
hours, err := client.Hour(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(hours) // the array of records
```


### Image

Create an instance: `image := client.Image(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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
| `height` | `float64` | Native height of the image |
| `id` | `string` | Unique identifier of this resource. |
| `iiif_url` | `any` | IIIF URL of this image |
| `is_educational_resource` | `bool` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | Whether this resource is considered to be educational |
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
| `width` | `float64` | Native width of the image |

#### Example: Load

```go
image, err := client.Image(nil).Load(map[string]any{"id": "image_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(image) // the loaded record
```

#### Example: List

```go
images, err := client.Image(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(images) // the array of records
```


### LandingPage

Create an instance: `landingPage := client.LandingPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
landingPage, err := client.LandingPage(nil).Load(map[string]any{"id": "landing_page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(landingPage) // the loaded record
```

#### Example: List

```go
landingPages, err := client.LandingPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(landingPages) // the array of records
```


### Place

Create an instance: `place := client.Place(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_link` | `any` | REST API link for this resource |
| `api_model` | `any` | REST API resource type or endpoint |
| `id` | `string` | Unique identifier of this resource. |
| `latitude` | `float64` | Latitude coordinate of the center of the room |
| `longitude` | `float64` | Longitude coordinate of the center of the room |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `tgn_id` | `string` | Reconciled identifier of this object in the Getty's Thesauraus of Geographic Names (TGN) |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |

#### Example: Load

```go
place, err := client.Place(nil).Load(map[string]any{"id": "place_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(place) // the loaded record
```

#### Example: List

```go
places, err := client.Place(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(places) // the array of records
```


### PressRelease

Create an instance: `pressRelease := client.PressRelease(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
pressRelease, err := client.PressRelease(nil).Load(map[string]any{"id": "press_release_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(pressRelease) // the loaded record
```

#### Example: List

```go
pressReleases, err := client.PressRelease(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(pressReleases) // the array of records
```


### PrintedPublication

Create an instance: `printedPublication := client.PrintedPublication(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
printedPublication, err := client.PrintedPublication(nil).Load(map[string]any{"id": "printed_publication_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(printedPublication) // the loaded record
```

#### Example: List

```go
printedPublications, err := client.PrintedPublication(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(printedPublications) // the array of records
```


### Product

Create an instance: `product := client.Product(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
product, err := client.Product(nil).Load(map[string]any{"id": "product_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(product) // the loaded record
```

#### Example: List

```go
products, err := client.Product(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(products) // the array of records
```


### Publication

Create an instance: `publication := client.Publication(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
publication, err := client.Publication(nil).Load(map[string]any{"id": "publication_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(publication) // the loaded record
```

#### Example: List

```go
publications, err := client.Publication(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(publications) // the array of records
```


### Search

Create an instance: `search := client.Search(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_id` | `string` | API unique identifier |
| `api_link` | `any` | URL to this recource in the API |
| `api_model` | `any` | Name of the model the resource represents |
| `id` | `string` | Unique identifier within the search index |
| `is_boosted` | `bool` | Whether this record has been flagged to be boosted |
| `score` | `float64` | Search index ranking of the result |
| `thumbnail` | `any` | Metadata on the image representing this record |
| `timestamp` | `any` | Date this record was last updated in the API |
| `title` | `string` | The name of this resource |

#### Example: List

```go
searchs, err := client.Search(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(searchs) // the array of records
```


### Section

Create an instance: `section := client.Section(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
section, err := client.Section(nil).Load(map[string]any{"id": "section_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(section) // the loaded record
```

#### Example: List

```go
sections, err := client.Section(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(sections) // the array of records
```


### Site

Create an instance: `site := client.Site(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
site, err := client.Site(nil).Load(map[string]any{"id": "site_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(site) // the loaded record
```

#### Example: List

```go
sites, err := client.Site(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(sites) // the array of records
```


### Sound

Create an instance: `sound := client.Sound(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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
| `is_educational_resource` | `bool` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | Whether this resource is considered to be educational |
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

```go
sound, err := client.Sound(nil).Load(map[string]any{"id": "sound_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(sound) // the loaded record
```

#### Example: List

```go
sounds, err := client.Sound(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(sounds) // the array of records
```


### StaticPage

Create an instance: `staticPage := client.StaticPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
staticPage, err := client.StaticPage(nil).Load(map[string]any{"id": "static_page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(staticPage) // the loaded record
```

#### Example: List

```go
staticPages, err := client.StaticPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(staticPages) // the array of records
```


### Text

Create an instance: `text := client.Text(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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
| `is_educational_resource` | `bool` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | Whether this resource is considered to be educational |
| `lake_guid` | `any` | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `type` | `any` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |

#### Example: Load

```go
text, err := client.Text(nil).Load(map[string]any{"id": "text_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(text) // the loaded record
```

#### Example: List

```go
texts, err := client.Text(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(texts) // the array of records
```


### Tour

Create an instance: `tour := client.Tour(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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
| `weight` | `float64` | Number representing this tour's sort order |

#### Example: Load

```go
tour, err := client.Tour(nil).Load(map[string]any{"id": "tour_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(tour) // the loaded record
```

#### Example: List

```go
tours, err := client.Tour(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(tours) // the array of records
```


### Video

Create an instance: `video := client.Video(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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
| `is_educational_resource` | `bool` | Whether this resource is considered to be educational |
| `is_multimedia_resource` | `bool` | Whether this resource is considered to be multimedia |
| `is_teacher_resource` | `bool` | Whether this resource is considered to be educational |
| `lake_guid` | `any` | Unique UUID of this resource in LAKE, our DAMS. |
| `source_updated_at` | `any` | Date and time the resource was updated in the source system |
| `suggest_autocomplete_all` | `any` | Internal field to power the `/autosuggest` endpoint. |
| `suggest_autocomplete_boosted` | `any` | Internal field to power the `/autocomplete` endpoint. |
| `timestamp` | `any` | Date and time the record was updated in the aggregator search index |
| `title` | `string` | The name of this resource |
| `type` | `any` | Type always takes one of the following values: image, sound, text, video |
| `updated_at` | `any` | Date and time the record was updated in the aggregator database |

#### Example: Load

```go
video, err := client.Video(nil).Load(map[string]any{"id": "video_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(video) // the loaded record
```

#### Example: List

```go
videos, err := client.Video(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(videos) // the array of records
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

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

Use `core.ToMapAny()` to safely cast results and nested data.

### Package structure

```
github.com/voxgig-sdk/art-institute-of-chicago-sdk/go/
├── art-institute-of-chicago.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/voxgig-sdk/art-institute-of-chicago-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `List`, the entity
stores the returned data and match criteria internally.

```go
publication := client.Publication(nil)
publication.List(nil, nil)

// publication.Data() now returns the publication data from the last list
// publication.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
