# ArtInstituteOfChicago Golang SDK



The Golang SDK for the ArtInstituteOfChicago API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

> Other languages, the CLI, and MCP server live alongside this one — see
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

agent, err := client.Agent(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(agent) // the loaded mock data
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
| `Create` | `(reqdata, ctrl map[string]any) (any, error)` | Create a new entity. |
| `Update` | `(reqdata, ctrl map[string]any) (any, error)` | Update an existing entity. |
| `Remove` | `(reqmatch, ctrl map[string]any) (any, error)` | Remove an entity. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(value, error)`. The `value` is the
operation's data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `Load` / `Create` / `Update` / `Remove` | the entity record (`map[string]any`) |
| `List` | a `[]any` of entity records |

Check `err` first, then use the value directly (or the typed
`...Typed` variants, which return the entity's model struct and a typed
slice):

    agent, err := client.Agent(nil).Load(map[string]any{"id": "example_id"}, nil)
    if err != nil { /* handle */ }
    // agent is the loaded record

Only `Direct()` returns a response envelope — a `map[string]any` with
`"ok"`, `"status"`, `"headers"`, and `"data"` keys.

### Entities

#### Agent

| Field | Description |
| --- | --- |
| `"alt_title"` |  |
| `"api_link"` |  |
| `"api_model"` |  |
| `"birth_date"` |  |
| `"death_date"` |  |
| `"description"` |  |
| `"id"` |  |
| `"is_artist"` |  |
| `"sort_title"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"ulan_id"` |  |
| `"updated_at"` |  |

Operations: List, Load.

API path: `/agents`

#### AgentRole

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"id"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |

Operations: List, Load.

API path: `/agent-roles`

#### AgentType

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"id"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |

Operations: List, Load.

API path: `/agent-types`

#### Article

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"copy"` |  |
| `"id"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |

Operations: List, Load.

API path: `/articles`

#### Artwork

| Field | Description |
| --- | --- |
| `"alt_artist_id"` |  |
| `"alt_classification_id"` |  |
| `"alt_image_id"` |  |
| `"alt_material_id"` |  |
| `"alt_style_id"` |  |
| `"alt_subject_id"` |  |
| `"alt_technique_id"` |  |
| `"alt_title"` |  |
| `"api_link"` |  |
| `"api_model"` |  |
| `"artist_display"` |  |
| `"artist_id"` |  |
| `"artist_title"` |  |
| `"artwork_type_id"` |  |
| `"artwork_type_title"` |  |
| `"boost_rank"` |  |
| `"catalog_based_search_keyword_title"` |  |
| `"catalogue_display"` |  |
| `"category_id"` |  |
| `"category_title"` |  |
| `"classification_id"` |  |
| `"classification_title"` |  |
| `"color"` |  |
| `"colorfulness"` |  |
| `"copyright_notice"` |  |
| `"credit_line"` |  |
| `"date_display"` |  |
| `"date_end"` |  |
| `"date_qualifier_id"` |  |
| `"date_qualifier_title"` |  |
| `"date_start"` |  |
| `"department_id"` |  |
| `"department_title"` |  |
| `"description"` |  |
| `"dimension"` |  |
| `"dimensions_detail"` |  |
| `"document_id"` |  |
| `"edition"` |  |
| `"exhibition_history"` |  |
| `"fiscal_year"` |  |
| `"fiscal_year_deaccession"` |  |
| `"gallery_id"` |  |
| `"gallery_title"` |  |
| `"has_advanced_imaging"` |  |
| `"has_educational_resource"` |  |
| `"has_multimedia_resource"` |  |
| `"has_not_been_viewed_much"` |  |
| `"id"` |  |
| `"image_embedding"` |  |
| `"image_id"` |  |
| `"inscription"` |  |
| `"internal_department_id"` |  |
| `"is_boosted"` |  |
| `"is_on_view"` |  |
| `"is_public_domain"` |  |
| `"is_zoomable"` |  |
| `"latitude"` |  |
| `"latlon"` |  |
| `"longitude"` |  |
| `"main_reference_number"` |  |
| `"material_id"` |  |
| `"material_title"` |  |
| `"max_zoom_window_size"` |  |
| `"medium_display"` |  |
| `"nomisma_id"` |  |
| `"on_loan_display"` |  |
| `"pageview"` |  |
| `"pageviews_recent"` |  |
| `"place_of_origin"` |  |
| `"provenance_text"` |  |
| `"publication_history"` |  |
| `"publishing_verification_level"` |  |
| `"section_id"` |  |
| `"section_title"` |  |
| `"short_description"` |  |
| `"site_id"` |  |
| `"sound_id"` |  |
| `"source_updated_at"` |  |
| `"style_id"` |  |
| `"style_title"` |  |
| `"subject_id"` |  |
| `"subject_title"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"technique_id"` |  |
| `"technique_title"` |  |
| `"term_title"` |  |
| `"text_embedding"` |  |
| `"text_id"` |  |
| `"theme_title"` |  |
| `"thumbnail"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |
| `"video_id"` |  |

Operations: List, Load.

API path: `/artworks`

#### ArtworkDateQualifier

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"id"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |

Operations: List, Load.

API path: `/artwork-date-qualifiers`

#### ArtworkPlaceQualifier

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"id"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |

Operations: List, Load.

API path: `/artwork-place-qualifiers`

#### ArtworkType

| Field | Description |
| --- | --- |
| `"aat_id"` |  |
| `"api_link"` |  |
| `"api_model"` |  |
| `"id"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |

Operations: List, Load.

API path: `/artwork-types`

#### CategoryTerm

| Field | Description |
| --- | --- |
| `"aat_id"` |  |
| `"api_link"` |  |
| `"api_model"` |  |
| `"id"` |  |
| `"parent_id"` |  |
| `"source_updated_at"` |  |
| `"subtype"` |  |
| `"suggest_autocomplete_all"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |

Operations: List, Load.

API path: `/category-terms`

#### DigitalPublication

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"copy"` |  |
| `"id"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |
| `"web_url"` |  |

Operations: List, Load.

API path: `/digital-publications`

#### DigitalPublicationArticle

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"author_display"` |  |
| `"copy"` |  |
| `"digital_publication_id"` |  |
| `"id"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |
| `"web_url"` |  |

Operations: List, Load.

API path: `/digital-publication-articles`

#### EducatorResource

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"copy"` |  |
| `"id"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |
| `"web_url"` |  |

Operations: List, Load.

API path: `/educator-resources`

#### Event

| Field | Description |
| --- | --- |
| `"alt_audience_id"` |  |
| `"alt_event_type_id"` |  |
| `"api_link"` |  |
| `"api_model"` |  |
| `"audience_id"` |  |
| `"buy_button_caption"` |  |
| `"buy_button_text"` |  |
| `"date_display"` |  |
| `"description"` |  |
| `"door_time"` |  |
| `"end_date"` |  |
| `"end_time"` |  |
| `"entrance"` |  |
| `"event_host_id"` |  |
| `"event_host_title"` |  |
| `"event_type_id"` |  |
| `"header_description"` |  |
| `"hero_caption"` |  |
| `"id"` |  |
| `"image_url"` |  |
| `"is_admission_required"` |  |
| `"is_after_hour"` |  |
| `"is_free"` |  |
| `"is_member_exclusive"` |  |
| `"is_private"` |  |
| `"is_registration_required"` |  |
| `"is_sales_button_hidden"` |  |
| `"is_sold_out"` |  |
| `"is_ticketed"` |  |
| `"is_virtual_event"` |  |
| `"join_url"` |  |
| `"layout_type"` |  |
| `"list_description"` |  |
| `"location"` |  |
| `"program_id"` |  |
| `"program_title"` |  |
| `"rsvp_link"` |  |
| `"search_tag"` |  |
| `"short_description"` |  |
| `"slug"` |  |
| `"source_updated_at"` |  |
| `"start_date"` |  |
| `"start_time"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"survey_url"` |  |
| `"ticketed_event_id"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"title_display"` |  |
| `"updated_at"` |  |
| `"virtual_event_passcode"` |  |
| `"virtual_event_url"` |  |

Operations: List, Load.

API path: `/events`

#### EventOccurrence

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"button_caption"` |  |
| `"button_text"` |  |
| `"button_url"` |  |
| `"description"` |  |
| `"end_at"` |  |
| `"event_id"` |  |
| `"id"` |  |
| `"image_url"` |  |
| `"is_private"` |  |
| `"is_sales_button_hidden"` |  |
| `"is_ticketed"` |  |
| `"location"` |  |
| `"off_sale_at"` |  |
| `"on_sale_at"` |  |
| `"short_description"` |  |
| `"source_updated_at"` |  |
| `"start_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"title_display"` |  |
| `"updated_at"` |  |

Operations: List, Load.

API path: `/event-occurrences`

#### EventProgram

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"id"` |  |
| `"is_affiliate_group"` |  |
| `"is_event_host"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |

Operations: List, Load.

API path: `/event-programs`

#### Exhibition

| Field | Description |
| --- | --- |
| `"aic_end_at"` |  |
| `"aic_start_at"` |  |
| `"alt_image_id"` |  |
| `"api_link"` |  |
| `"api_model"` |  |
| `"artist_id"` |  |
| `"artwork_id"` |  |
| `"artwork_title"` |  |
| `"document_id"` |  |
| `"gallery_id"` |  |
| `"gallery_title"` |  |
| `"id"` |  |
| `"image_id"` |  |
| `"image_url"` |  |
| `"is_featured"` |  |
| `"is_published"` |  |
| `"position"` |  |
| `"short_description"` |  |
| `"site_id"` |  |
| `"source_updated_at"` |  |
| `"status"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |
| `"web_url"` |  |

Operations: List, Load.

API path: `/exhibitions`

#### Gallery

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"floor"` |  |
| `"id"` |  |
| `"is_closed"` |  |
| `"latitude"` |  |
| `"latlon"` |  |
| `"longitude"` |  |
| `"number"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"tgn_id"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |

Operations: List, Load.

API path: `/galleries`

#### GenericPage

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"copy"` |  |
| `"id"` |  |
| `"search_tag"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |
| `"web_url"` |  |

Operations: List, Load.

API path: `/generic-pages`

#### Highlight

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"copy"` |  |
| `"id"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |

Operations: List, Load.

API path: `/highlights`

#### Hour

| Field | Description |
| --- | --- |
| `"additional_text"` |  |
| `"api_link"` |  |
| `"api_model"` |  |
| `"friday_is_closed"` |  |
| `"friday_member_close"` |  |
| `"friday_member_open"` |  |
| `"friday_public_close"` |  |
| `"friday_public_open"` |  |
| `"id"` |  |
| `"monday_is_closed"` |  |
| `"monday_member_close"` |  |
| `"monday_member_open"` |  |
| `"monday_public_close"` |  |
| `"monday_public_open"` |  |
| `"saturday_is_closed"` |  |
| `"saturday_member_close"` |  |
| `"saturday_member_open"` |  |
| `"saturday_public_close"` |  |
| `"saturday_public_open"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"summary"` |  |
| `"sunday_is_closed"` |  |
| `"sunday_member_close"` |  |
| `"sunday_member_open"` |  |
| `"sunday_public_close"` |  |
| `"sunday_public_open"` |  |
| `"thursday_is_closed"` |  |
| `"thursday_member_close"` |  |
| `"thursday_member_open"` |  |
| `"thursday_public_close"` |  |
| `"thursday_public_open"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"tuesday_is_closed"` |  |
| `"tuesday_member_close"` |  |
| `"tuesday_member_open"` |  |
| `"tuesday_public_close"` |  |
| `"tuesday_public_open"` |  |
| `"updated_at"` |  |
| `"wednesday_is_closed"` |  |
| `"wednesday_member_close"` |  |
| `"wednesday_member_open"` |  |
| `"wednesday_public_close"` |  |
| `"wednesday_public_open"` |  |

Operations: List, Load.

API path: `/hours`

#### Image

| Field | Description |
| --- | --- |
| `"ahash"` |  |
| `"alt_text"` |  |
| `"api_link"` |  |
| `"api_model"` |  |
| `"artwork_id"` |  |
| `"artwork_title"` |  |
| `"color"` |  |
| `"colorfulness"` |  |
| `"content"` |  |
| `"content_e_tag"` |  |
| `"credit_line"` |  |
| `"fingerprint"` |  |
| `"height"` |  |
| `"id"` |  |
| `"iiif_url"` |  |
| `"is_educational_resource"` |  |
| `"is_multimedia_resource"` |  |
| `"is_teacher_resource"` |  |
| `"lake_guid"` |  |
| `"lqip"` |  |
| `"phash"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"type"` |  |
| `"updated_at"` |  |
| `"width"` |  |

Operations: List, Load.

API path: `/images`

#### LandingPage

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"copy"` |  |
| `"id"` |  |
| `"search_tag"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |
| `"web_url"` |  |

Operations: List, Load.

API path: `/landing-pages`

#### Place

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"id"` |  |
| `"latitude"` |  |
| `"longitude"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"tgn_id"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |

Operations: List, Load.

API path: `/places`

#### PressRelease

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"copy"` |  |
| `"id"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |
| `"web_url"` |  |

Operations: List, Load.

API path: `/press-releases`

#### PrintedPublication

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"copy"` |  |
| `"id"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |
| `"web_url"` |  |

Operations: List, Load.

API path: `/printed-publications`

#### Product

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"artist_id"` |  |
| `"artwork_id"` |  |
| `"description"` |  |
| `"exhibition_id"` |  |
| `"external_sku"` |  |
| `"id"` |  |
| `"image_url"` |  |
| `"max_compare_at_price"` |  |
| `"max_current_price"` |  |
| `"min_compare_at_price"` |  |
| `"min_current_price"` |  |
| `"price_display"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |
| `"web_url"` |  |

Operations: List, Load.

API path: `/products`

#### Publication

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"id"` |  |
| `"section_id"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |
| `"web_url"` |  |

Operations: List, Load.

API path: `/publications`

#### Search

| Field | Description |
| --- | --- |
| `"api_id"` |  |
| `"api_link"` |  |
| `"api_model"` |  |
| `"id"` |  |
| `"is_boosted"` |  |
| `"score"` |  |
| `"thumbnail"` |  |
| `"timestamp"` |  |
| `"title"` |  |

Operations: List.

API path: `/agents/search`

#### Section

| Field | Description |
| --- | --- |
| `"accession"` |  |
| `"api_link"` |  |
| `"api_model"` |  |
| `"artwork_id"` |  |
| `"content"` |  |
| `"generic_page_id"` |  |
| `"id"` |  |
| `"publication_id"` |  |
| `"publication_title"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |
| `"web_url"` |  |

Operations: List, Load.

API path: `/sections`

#### Site

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"artwork_id"` |  |
| `"artwork_title"` |  |
| `"description"` |  |
| `"exhibition_id"` |  |
| `"exhibition_title"` |  |
| `"id"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |
| `"web_url"` |  |

Operations: List, Load.

API path: `/sites`

#### Sound

| Field | Description |
| --- | --- |
| `"alt_text"` |  |
| `"api_link"` |  |
| `"api_model"` |  |
| `"artwork_id"` |  |
| `"artwork_title"` |  |
| `"content"` |  |
| `"content_e_tag"` |  |
| `"credit_line"` |  |
| `"id"` |  |
| `"is_educational_resource"` |  |
| `"is_multimedia_resource"` |  |
| `"is_teacher_resource"` |  |
| `"lake_guid"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"transcript"` |  |
| `"type"` |  |
| `"updated_at"` |  |
| `"web_url"` |  |

Operations: List, Load.

API path: `/mobile-sounds`

#### StaticPage

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"id"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |
| `"web_url"` |  |

Operations: List, Load.

API path: `/static-pages`

#### Text

| Field | Description |
| --- | --- |
| `"alt_text"` |  |
| `"api_link"` |  |
| `"api_model"` |  |
| `"artwork_id"` |  |
| `"artwork_title"` |  |
| `"content"` |  |
| `"content_e_tag"` |  |
| `"credit_line"` |  |
| `"id"` |  |
| `"is_educational_resource"` |  |
| `"is_multimedia_resource"` |  |
| `"is_teacher_resource"` |  |
| `"lake_guid"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"type"` |  |
| `"updated_at"` |  |

Operations: List, Load.

API path: `/texts`

#### Tour

| Field | Description |
| --- | --- |
| `"api_link"` |  |
| `"api_model"` |  |
| `"artist_title"` |  |
| `"artwork_title"` |  |
| `"description"` |  |
| `"id"` |  |
| `"image"` |  |
| `"intro"` |  |
| `"intro_link"` |  |
| `"intro_transcript"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"updated_at"` |  |
| `"weight"` |  |

Operations: List, Load.

API path: `/tours`

#### Video

| Field | Description |
| --- | --- |
| `"alt_text"` |  |
| `"api_link"` |  |
| `"api_model"` |  |
| `"artwork_id"` |  |
| `"artwork_title"` |  |
| `"content"` |  |
| `"content_e_tag"` |  |
| `"credit_line"` |  |
| `"id"` |  |
| `"is_educational_resource"` |  |
| `"is_multimedia_resource"` |  |
| `"is_teacher_resource"` |  |
| `"lake_guid"` |  |
| `"source_updated_at"` |  |
| `"suggest_autocomplete_all"` |  |
| `"suggest_autocomplete_boosted"` |  |
| `"timestamp"` |  |
| `"title"` |  |
| `"type"` |  |
| `"updated_at"` |  |

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

Create an instance: `agent_role := client.AgentRole(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
agent_role, err := client.AgentRole(nil).Load(map[string]any{"id": "agent_role_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(agent_role) // the loaded record
```

#### Example: List

```go
agent_roles, err := client.AgentRole(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(agent_roles) // the array of records
```


### AgentType

Create an instance: `agent_type := client.AgentType(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
agent_type, err := client.AgentType(nil).Load(map[string]any{"id": "agent_type_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(agent_type) // the loaded record
```

#### Example: List

```go
agent_types, err := client.AgentType(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(agent_types) // the array of records
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

Create an instance: `artwork_date_qualifier := client.ArtworkDateQualifier(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
artwork_date_qualifier, err := client.ArtworkDateQualifier(nil).Load(map[string]any{"id": "artwork_date_qualifier_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(artwork_date_qualifier) // the loaded record
```

#### Example: List

```go
artwork_date_qualifiers, err := client.ArtworkDateQualifier(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(artwork_date_qualifiers) // the array of records
```


### ArtworkPlaceQualifier

Create an instance: `artwork_place_qualifier := client.ArtworkPlaceQualifier(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
artwork_place_qualifier, err := client.ArtworkPlaceQualifier(nil).Load(map[string]any{"id": "artwork_place_qualifier_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(artwork_place_qualifier) // the loaded record
```

#### Example: List

```go
artwork_place_qualifiers, err := client.ArtworkPlaceQualifier(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(artwork_place_qualifiers) // the array of records
```


### ArtworkType

Create an instance: `artwork_type := client.ArtworkType(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
artwork_type, err := client.ArtworkType(nil).Load(map[string]any{"id": "artwork_type_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(artwork_type) // the loaded record
```

#### Example: List

```go
artwork_types, err := client.ArtworkType(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(artwork_types) // the array of records
```


### CategoryTerm

Create an instance: `category_term := client.CategoryTerm(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
category_term, err := client.CategoryTerm(nil).Load(map[string]any{"id": "category_term_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(category_term) // the loaded record
```

#### Example: List

```go
category_terms, err := client.CategoryTerm(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(category_terms) // the array of records
```


### DigitalPublication

Create an instance: `digital_publication := client.DigitalPublication(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
digital_publication, err := client.DigitalPublication(nil).Load(map[string]any{"id": "digital_publication_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(digital_publication) // the loaded record
```

#### Example: List

```go
digital_publications, err := client.DigitalPublication(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(digital_publications) // the array of records
```


### DigitalPublicationArticle

Create an instance: `digital_publication_article := client.DigitalPublicationArticle(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
digital_publication_article, err := client.DigitalPublicationArticle(nil).Load(map[string]any{"id": "digital_publication_article_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(digital_publication_article) // the loaded record
```

#### Example: List

```go
digital_publication_articles, err := client.DigitalPublicationArticle(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(digital_publication_articles) // the array of records
```


### EducatorResource

Create an instance: `educator_resource := client.EducatorResource(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
educator_resource, err := client.EducatorResource(nil).Load(map[string]any{"id": "educator_resource_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(educator_resource) // the loaded record
```

#### Example: List

```go
educator_resources, err := client.EducatorResource(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(educator_resources) // the array of records
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

Create an instance: `event_occurrence := client.EventOccurrence(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
event_occurrence, err := client.EventOccurrence(nil).Load(map[string]any{"id": "event_occurrence_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(event_occurrence) // the loaded record
```

#### Example: List

```go
event_occurrences, err := client.EventOccurrence(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(event_occurrences) // the array of records
```


### EventProgram

Create an instance: `event_program := client.EventProgram(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
event_program, err := client.EventProgram(nil).Load(map[string]any{"id": "event_program_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(event_program) // the loaded record
```

#### Example: List

```go
event_programs, err := client.EventProgram(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(event_programs) // the array of records
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

Create an instance: `generic_page := client.GenericPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
generic_page, err := client.GenericPage(nil).Load(map[string]any{"id": "generic_page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(generic_page) // the loaded record
```

#### Example: List

```go
generic_pages, err := client.GenericPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(generic_pages) // the array of records
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

Create an instance: `landing_page := client.LandingPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
landing_page, err := client.LandingPage(nil).Load(map[string]any{"id": "landing_page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(landing_page) // the loaded record
```

#### Example: List

```go
landing_pages, err := client.LandingPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(landing_pages) // the array of records
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

Create an instance: `press_release := client.PressRelease(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
press_release, err := client.PressRelease(nil).Load(map[string]any{"id": "press_release_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(press_release) // the loaded record
```

#### Example: List

```go
press_releases, err := client.PressRelease(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(press_releases) // the array of records
```


### PrintedPublication

Create an instance: `printed_publication := client.PrintedPublication(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
printed_publication, err := client.PrintedPublication(nil).Load(map[string]any{"id": "printed_publication_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(printed_publication) // the loaded record
```

#### Example: List

```go
printed_publications, err := client.PrintedPublication(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(printed_publications) // the array of records
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

Create an instance: `static_page := client.StaticPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
static_page, err := client.StaticPage(nil).Load(map[string]any{"id": "static_page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(static_page) // the loaded record
```

#### Example: List

```go
static_pages, err := client.StaticPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(static_pages) // the array of records
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
error is returned to the caller. An unexpected panic triggers the
`PreUnexpected` hook.

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

Entity instances are stateful. After a successful `Load`, the entity
stores the returned data and match criteria internally.

```go
agent := client.Agent(nil)
agent.Load(map[string]any{"id": "example_id"}, nil)

// agent.Data() now returns the loaded agent data
// agent.Match() returns the last match criteria
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
