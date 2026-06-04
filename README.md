# ArtInstituteOfChicago SDK

Explore the Art Institute of Chicago's public collection — artworks, artists, exhibitions, galleries, and IIIF images

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

## About Art Institution of Chicago API

The Art Institute of Chicago API is a read-only aggregator of public data from the [Art Institute of Chicago](https://www.artic.edu/), exposing the museum's collection, programming, and editorial content over a single JSON HTTP interface rooted at `https://api.artic.edu/api/v1`.

What you get from the API:

- Collection records — artworks, the agents (artists, makers, organisations) connected to them, the places and galleries where they live, and the controlled vocabularies that describe them.
- Programming and editorial — exhibitions, events, tours, articles, press releases, digital and printed publications, educator resources, and static site pages.
- Media — IIIF-backed images plus videos, sounds, and texts associated with the above.
- Search — every resource type has its own `/search` endpoint, and a unified `/search` endpoint queries across the whole dataset with faceting, pagination, and sorting.

No authentication is required for basic use. Anonymous clients are throttled to roughly 60 requests per minute per IP; the docs ask that you send an `AIC-User-Agent` header identifying your project and a contact email, and direct higher-volume users to engineering@artic.edu. Most endpoints have CORS enabled and listings default to sorting by last-updated descending.

## Try it

**TypeScript**
```bash
npm install art-institute-of-chicago
```

**Python**
```bash
pip install art-institute-of-chicago-sdk
```

**PHP**
```bash
composer require voxgig/art-institute-of-chicago-sdk
```

**Golang**
```bash
go get github.com/voxgig-sdk/art-institute-of-chicago-sdk/go
```

**Ruby**
```bash
gem install art-institute-of-chicago-sdk
```

**Lua**
```bash
luarocks install art-institute-of-chicago-sdk
```

## 30-second quickstart

### TypeScript

```ts
import { ArtInstituteOfChicagoSDK } from 'art-institute-of-chicago'

const client = new ArtInstituteOfChicagoSDK({})

// List all agents
const agents = await client.Agent().list()
```

See the [TypeScript README](ts/README.md) for the
full guide, or scroll down for the same example in other languages.

## What's in the box

| Surface | Use it for | Path |
| --- | --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | App integration | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | Scripts, CI, ops, one-off API calls | `go-cli/` |
| **MCP server** | AI agents (Claude, Cursor, Cline) | `go-mcp/` |

## Use it from an AI agent (MCP)

The generated MCP server exposes every operation in this SDK as an
[MCP](https://modelcontextprotocol.io) tool that Claude, Cursor or Cline
can call directly. Build and register it:

```bash
cd go-mcp && go build -o art-institute-of-chicago-mcp .
```

Then add it to your agent's MCP config (Claude Desktop, Cursor, etc.):

```json
{
  "mcpServers": {
    "art-institute-of-chicago": {
      "command": "/abs/path/to/art-institute-of-chicago-mcp"
    }
  }
}
```

## Entities

The API exposes 35 entities:

| Entity | Description | API path |
| --- | --- | --- |
| **Agent** | A person or organisation (artist, maker, donor, etc.) related to collection records. Exposed at `/agents`, `/agents/{id}`, and `/agents/search`. | `/agents` |
| **AgentRole** | A controlled vocabulary describing how an agent relates to an artwork (e.g. artist, donor). | `/agent-roles` |
| **AgentType** | A classification of agents (individual, corporate body, and so on). | `/agent-types` |
| **Article** | Editorial articles published by the museum. | `/articles` |
| **Artwork** | An object in the museum's collection — paintings, sculptures, prints, and so on. Exposed at `/artworks`, `/artworks/{id}`, and `/artworks/search`. | `/artworks` |
| **ArtworkDateQualifier** | Controlled vocabulary qualifying how an artwork's date is interpreted (circa, before, after, etc.). | `/artwork-date-qualifiers` |
| **ArtworkPlaceQualifier** | Controlled vocabulary qualifying how an artwork's place relates to it (made in, found in, etc.). | `/artwork-place-qualifiers` |
| **ArtworkType** | Classification of artworks by medium or form (painting, drawing, sculpture, etc.). | `/artwork-types` |
| **CategoryTerm** | Hierarchical category and classification terms used to tag collection records. | `/category-terms` |
| **DigitalPublication** | A digital publication produced by the museum. | `/digital-publications` |
| **DigitalPublicationArticle** | An individual article within a digital publication. | `/digital-publication-articles` |
| **EducatorResource** | Teaching resources made available for educators. | `/educator-resources` |
| **Event** | A scheduled public event at the museum. | `/events` |
| **EventOccurrence** | A specific dated instance of a recurring event. | `/event-occurrences` |
| **EventProgram** | A program grouping related events. | `/event-programs` |
| **Exhibition** | Past, current, and upcoming exhibitions. Exposed at `/exhibitions`, `/exhibitions/{id}`, and `/exhibitions/search`. | `/exhibitions` |
| **Gallery** | Physical galleries within the museum building. Exposed at `/galleries`, `/galleries/{id}`, and `/galleries/search`. | `/galleries` |
| **GenericPage** | Generic CMS-managed pages on the museum website. | `/generic-pages` |
| **Highlight** | Curated highlight entries promoting selected content. | `/highlights` |
| **Hour** | Opening-hours records for the museum. | `/hours` |
| **Image** | Media records for images, served through the IIIF Image API 2.0. Exposed at `/images`, `/images/{id}`, and `/images/search`. | `/images` |
| **LandingPage** | Landing pages on the museum website. | `/landing-pages` |
| **Place** | Geographic places associated with collection records; uses the Getty Thesaurus of Geographic Names. | `/places` |
| **PressRelease** | Museum press releases. | `/press-releases` |
| **PrintedPublication** | Print publications such as books and catalogues produced by the museum. | `/printed-publications` |
| **Product** | Products listed by the museum (for example shop items referenced in editorial content). | `/products` |
| **Publication** | Publication records that may span print and digital formats. | `/publications` |
| **Search** | Unified cross-resource search endpoint at `/search`, plus per-resource `/{resource}/search` endpoints supporting facets, pagination, and sorting. | `/agents/search` |
| **Section** | Sections within larger publications or site structures. | `/sections` |
| **Site** | Site grouping used to scope content across artic.edu properties. | `/sites` |
| **Sound** | Audio media records associated with collection or editorial content. | `/mobile-sounds` |
| **StaticPage** | Static informational pages on the museum website. | `/static-pages` |
| **Text** | Text media records (essays, labels, and other written content). | `/texts` |
| **Tour** | Self-guided or curated tours through the collection. | `/tours` |
| **Video** | Video media records associated with collection or editorial content. | `/videos` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
from artinstituteofchicago_sdk import ArtInstituteOfChicagoSDK

client = ArtInstituteOfChicagoSDK({})

# List all agents
agents, err = client.Agent(None).list(None, None)

# Load a specific agent
agent, err = client.Agent(None).load(
    {"id": "example_id"}, None
)
```

### PHP

```php
<?php
require_once 'artinstituteofchicago_sdk.php';

$client = new ArtInstituteOfChicagoSDK([]);

// List all agents
[$agents, $err] = $client->Agent(null)->list(null, null);

// Load a specific agent
[$agent, $err] = $client->Agent(null)->load(
    ["id" => "example_id"], null
);
```

### Golang

```go
import sdk "github.com/voxgig-sdk/art-institute-of-chicago-sdk/go"

client := sdk.NewArtInstituteOfChicagoSDK(map[string]any{})

// List all agents
agents, err := client.Agent(nil).List(nil, nil)
```

### Ruby

```ruby
require_relative "ArtInstituteOfChicago_sdk"

client = ArtInstituteOfChicagoSDK.new({})

# List all agents
agents, err = client.Agent(nil).list(nil, nil)

# Load a specific agent
agent, err = client.Agent(nil).load(
  { "id" => "example_id" }, nil
)
```

### Lua

```lua
local sdk = require("art-institute-of-chicago_sdk")

local client = sdk.new({})

-- List all agents
local agents, err = client:Agent(nil):list(nil, nil)

-- Load a specific agent
local agent, err = client:Agent(nil):load(
  { id = "example_id" }, nil
)
```

## Unit testing in offline mode

Every SDK ships a test mode that swaps the HTTP transport for an
in-memory mock, so unit tests run offline.

### TypeScript

```ts
const client = ArtInstituteOfChicagoSDK.test()
const result = await client.Agent().load({ id: 'test01' })
// result.ok === true, result.data contains mock data
```

### Python

```python
client = ArtInstituteOfChicagoSDK.test(None, None)
result, err = client.Agent(None).load(
    {"id": "test01"}, None
)
```

### PHP

```php
$client = ArtInstituteOfChicagoSDK::test(null, null);
[$result, $err] = $client->Agent(null)->load(
    ["id" => "test01"], null
);
```

### Golang

```go
client := sdk.TestSDK(nil, nil)
result, err := client.Agent(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = ArtInstituteOfChicagoSDK.test(nil, nil)
result, err = client.Agent(nil).load(
  { "id" => "test01" }, nil
)
```

### Lua

```lua
local client = sdk.test(nil, nil)
local result, err = client:Agent(nil):load(
  { id = "test01" }, nil
)
```

## How it works

Every SDK call runs the same five-stage pipeline:

1. **Point** — resolve the API endpoint from the operation definition.
2. **Spec** — build the HTTP specification (URL, method, headers, body).
3. **Request** — send the HTTP request.
4. **Response** — receive and parse the response.
5. **Result** — extract the result data for the caller.

A feature hook fires at each stage (e.g. `PrePoint`, `PreSpec`,
`PreRequest`), so features can inspect or modify the pipeline without
forking the SDK.

### Features

| Feature | Purpose |
| --- | --- |
| **TestFeature** | In-memory mock transport for testing without a live server |

Pass custom features via the `extend` option at construction time.

### Direct and Prepare

For endpoints the entity model doesn't cover, use the low-level methods:

- **`direct(fetchargs)`** — build and send an HTTP request in one step.
- **`prepare(fetchargs)`** — build the request without sending it.

Both accept a map with `path`, `method`, `params`, `query`,
`headers`, and `body`. See the [How-to guides](#how-to-guides) below.

## How-to guides

### Make a direct API call

When the entity interface does not cover an endpoint, use `direct`:

**TypeScript:**
```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
console.log(result.data)
```

**Python:**
```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})
```

**PHP:**
```php
[$result, $err] = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);
```

**Go:**
```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
```

**Ruby:**
```ruby
result, err = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})
```

**Lua:**
```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
```

## Per-language documentation

- [TypeScript](ts/README.md)
- [Python](py/README.md)
- [PHP](php/README.md)
- [Golang](go/README.md)
- [Ruby](rb/README.md)
- [Lua](lua/README.md)

## Using the Art Institution of Chicago API

- Upstream: [https://api.artic.edu/docs/](https://api.artic.edu/docs/)

- Most metadata is released under a Creative Commons Zero (CC0) 1.0 designation.
- Artwork descriptions are licensed CC-BY 4.0; places data is CC-BY with attribution to the Getty Thesaurus of Geographic Names.
- Images are served via the IIIF Image API 2.0; the docs recommend only using images of works tagged as public domain, and users remain responsible for individual copyright status.
- Use is subject to the museum's [terms and conditions](https://www.artic.edu/terms/terms-and-conditions).

---

Generated from the Art Institution of Chicago API OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).
