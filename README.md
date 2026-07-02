# ArtInstituteOfChicago SDK

Art Institution of Chicago API client, generated from the OpenAPI spec.

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

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

## Quickstart

### TypeScript

```ts
import { ArtInstituteOfChicagoSDK } from 'art-institute-of-chicago'

const client = new ArtInstituteOfChicagoSDK({
  apikey: process.env.ART-INSTITUTE-OF-CHICAGO_APIKEY,
})

// List all agents
const agents = await client.Agent().list()
console.log(agents.data)
```

See the [TypeScript README](ts/README.md) for the full guide.

## Surfaces

| Surface | Path |
| --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | `go-cli/` |
| **MCP server** | `go-mcp/` |

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
| **Agent** |  | `/agents` |
| **AgentRole** |  | `/agent-roles` |
| **AgentType** |  | `/agent-types` |
| **Article** |  | `/articles` |
| **Artwork** |  | `/artworks` |
| **ArtworkDateQualifier** |  | `/artwork-date-qualifiers` |
| **ArtworkPlaceQualifier** |  | `/artwork-place-qualifiers` |
| **ArtworkType** |  | `/artwork-types` |
| **CategoryTerm** |  | `/category-terms` |
| **DigitalPublication** |  | `/digital-publications` |
| **DigitalPublicationArticle** |  | `/digital-publication-articles` |
| **EducatorResource** |  | `/educator-resources` |
| **Event** |  | `/events` |
| **EventOccurrence** |  | `/event-occurrences` |
| **EventProgram** |  | `/event-programs` |
| **Exhibition** |  | `/exhibitions` |
| **Gallery** |  | `/galleries` |
| **GenericPage** |  | `/generic-pages` |
| **Highlight** |  | `/highlights` |
| **Hour** |  | `/hours` |
| **Image** |  | `/images` |
| **LandingPage** |  | `/landing-pages` |
| **Place** |  | `/places` |
| **PressRelease** |  | `/press-releases` |
| **PrintedPublication** |  | `/printed-publications` |
| **Product** |  | `/products` |
| **Publication** |  | `/publications` |
| **Search** |  | `/agents/search` |
| **Section** |  | `/sections` |
| **Site** |  | `/sites` |
| **Sound** |  | `/mobile-sounds` |
| **StaticPage** |  | `/static-pages` |
| **Text** |  | `/texts` |
| **Tour** |  | `/tours` |
| **Video** |  | `/videos` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
import os
from artinstituteofchicago_sdk import ArtInstituteOfChicagoSDK

client = ArtInstituteOfChicagoSDK({
    "apikey": os.environ.get("ART-INSTITUTE-OF-CHICAGO_APIKEY"),
})

# List all agents
agents, err = client.Agent().list()
print(agents)

# Load a specific agent
agent, err = client.Agent().load({"id": "example_id"})
print(agent)
```

### PHP

```php
<?php
require_once 'artinstituteofchicago_sdk.php';

$client = new ArtInstituteOfChicagoSDK([
    "apikey" => getenv("ART-INSTITUTE-OF-CHICAGO_APIKEY"),
]);

// List all agents
[$agents, $err] = $client->Agent()->list();
print_r($agents);

// Load a specific agent
[$agent, $err] = $client->Agent()->load(["id" => "example_id"]);
print_r($agent);
```

### Golang

```go
import sdk "github.com/voxgig-sdk/art-institute-of-chicago-sdk/go"

client := sdk.NewArtInstituteOfChicagoSDK(map[string]any{
    "apikey": os.Getenv("ART-INSTITUTE-OF-CHICAGO_APIKEY"),
})

// List all agents
agents, err := client.Agent(nil).List(nil, nil)
fmt.Println(agents)
```

### Ruby

```ruby
require_relative "ArtInstituteOfChicago_sdk"

client = ArtInstituteOfChicagoSDK.new({
  "apikey" => ENV["ART-INSTITUTE-OF-CHICAGO_APIKEY"],
})

# List all agents
agents, err = client.Agent().list
puts agents

# Load a specific agent
agent, err = client.Agent().load({ "id" => "example_id" })
puts agent
```

### Lua

```lua
local sdk = require("art-institute-of-chicago_sdk")

local client = sdk.new({
  apikey = os.getenv("ART-INSTITUTE-OF-CHICAGO_APIKEY"),
})

-- List all agents
local agents, err = client:Agent():list()
print(agents)

-- Load a specific agent
local agent, err = client:Agent():load({ id = "example_id" })
print(agent)
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
client = ArtInstituteOfChicagoSDK.test()
result, err = client.Agent().load({"id": "test01"})
```

### PHP

```php
$client = ArtInstituteOfChicagoSDK::test();
[$result, $err] = $client->Agent()->load(["id" => "test01"]);
```

### Golang

```go
client := sdk.Test()
result, err := client.Agent(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = ArtInstituteOfChicagoSDK.test
result, err = client.Agent().load({ "id" => "test01" })
```

### Lua

```lua
local client = sdk.test()
local result, err = client:Agent():load({ id = "test01" })
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

---

Generated from the Art Institution of Chicago API OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).
