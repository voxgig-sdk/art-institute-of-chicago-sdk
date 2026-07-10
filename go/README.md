# ArtInstituteOfChicago Golang SDK



The Golang SDK for the ArtInstituteOfChicago API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

It exposes the API as capitalised, semantic **Entities** — e.g. `client.Agent(nil)` — each with the same small set of operations (`List`, `Load`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

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


## Error handling

Every entity operation returns `(value, error)`. Check `err` before
using the value — there is no exception to catch:

```go
agents, err := client.Agent(nil).List(nil, nil)
if err != nil {
    // handle err
    return
}
_ = agents
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

agent, err := client.Agent(nil).List(
    nil, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(agent) // the returned mock data
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
| `alt_title` | `any` |  |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `birth_date` | `any` |  |
| `death_date` | `any` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `is_artist` | `bool` |  |
| `sort_title` | `any` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `ulan_id` | `string` |  |
| `updated_at` | `any` |  |

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
| `has_advanced_imaging` | `bool` |  |
| `has_educational_resource` | `bool` |  |
| `has_multimedia_resource` | `bool` |  |
| `has_not_been_viewed_much` | `bool` |  |
| `id` | `string` |  |
| `image_embedding` | `any` |  |
| `image_id` | `string` |  |
| `inscription` | `any` |  |
| `internal_department_id` | `string` |  |
| `is_boosted` | `bool` |  |
| `is_on_view` | `bool` |  |
| `is_public_domain` | `bool` |  |
| `is_zoomable` | `bool` |  |
| `latitude` | `float64` |  |
| `latlon` | `any` |  |
| `longitude` | `float64` |  |
| `main_reference_number` | `int` |  |
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
| `is_admission_required` | `bool` |  |
| `is_after_hour` | `bool` |  |
| `is_free` | `bool` |  |
| `is_member_exclusive` | `bool` |  |
| `is_private` | `bool` |  |
| `is_registration_required` | `bool` |  |
| `is_sales_button_hidden` | `bool` |  |
| `is_sold_out` | `bool` |  |
| `is_ticketed` | `bool` |  |
| `is_virtual_event` | `bool` |  |
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
| `is_private` | `bool` |  |
| `is_sales_button_hidden` | `bool` |  |
| `is_ticketed` | `bool` |  |
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
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `id` | `string` |  |
| `is_affiliate_group` | `bool` |  |
| `is_event_host` | `bool` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |

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
| `is_featured` | `bool` |  |
| `is_published` | `bool` |  |
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
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `floor` | `any` |  |
| `id` | `string` |  |
| `is_closed` | `bool` |  |
| `latitude` | `float64` |  |
| `latlon` | `any` |  |
| `longitude` | `float64` |  |
| `number` | `any` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `tgn_id` | `string` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |

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
| `height` | `float64` |  |
| `id` | `string` |  |
| `iiif_url` | `any` |  |
| `is_educational_resource` | `bool` |  |
| `is_multimedia_resource` | `bool` |  |
| `is_teacher_resource` | `bool` |  |
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
| `width` | `float64` |  |

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
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `id` | `string` |  |
| `latitude` | `float64` |  |
| `longitude` | `float64` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `tgn_id` | `string` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `updated_at` | `any` |  |

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
| `api_id` | `string` |  |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `id` | `string` |  |
| `is_boosted` | `bool` |  |
| `score` | `float64` |  |
| `thumbnail` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |

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
| `alt_text` | `any` |  |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `artwork_id` | `string` |  |
| `artwork_title` | `any` |  |
| `content` | `any` |  |
| `content_e_tag` | `any` |  |
| `credit_line` | `any` |  |
| `id` | `string` |  |
| `is_educational_resource` | `bool` |  |
| `is_multimedia_resource` | `bool` |  |
| `is_teacher_resource` | `bool` |  |
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
| `alt_text` | `any` |  |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `artwork_id` | `string` |  |
| `artwork_title` | `any` |  |
| `content` | `any` |  |
| `content_e_tag` | `any` |  |
| `credit_line` | `any` |  |
| `id` | `string` |  |
| `is_educational_resource` | `bool` |  |
| `is_multimedia_resource` | `bool` |  |
| `is_teacher_resource` | `bool` |  |
| `lake_guid` | `any` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `type` | `any` |  |
| `updated_at` | `any` |  |

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
| `weight` | `float64` |  |

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
| `alt_text` | `any` |  |
| `api_link` | `any` |  |
| `api_model` | `any` |  |
| `artwork_id` | `string` |  |
| `artwork_title` | `any` |  |
| `content` | `any` |  |
| `content_e_tag` | `any` |  |
| `credit_line` | `any` |  |
| `id` | `string` |  |
| `is_educational_resource` | `bool` |  |
| `is_multimedia_resource` | `bool` |  |
| `is_teacher_resource` | `bool` |  |
| `lake_guid` | `any` |  |
| `source_updated_at` | `any` |  |
| `suggest_autocomplete_all` | `any` |  |
| `suggest_autocomplete_boosted` | `any` |  |
| `timestamp` | `any` |  |
| `title` | `string` |  |
| `type` | `any` |  |
| `updated_at` | `any` |  |

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
agent := client.Agent(nil)
agent.List(nil, nil)

// agent.Data() now returns the agent data from the last list
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
