# ArtInstituteOfChicago SDK



Available for [Golang](go/) and [Go CLI](go-cli/) and [Go MCP server](go-mcp/) and [Lua](lua/) and [PHP](php/) and [Python](py/) and [Ruby](rb/) and [TypeScript](ts/).


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

Each entity supports the following operations where available: **load**, **list**, **create**,
**update**, and **remove**.


## Architecture

### Entity-operation model

Every SDK call follows the same pipeline:

1. **Point** — resolve the API endpoint from the operation definition.
2. **Spec** — build the HTTP specification (URL, method, headers, body).
3. **Request** — send the HTTP request.
4. **Response** — receive and parse the response.
5. **Result** — extract the result data for the caller.

At each stage a feature hook fires (e.g. `PrePoint`, `PreSpec`,
`PreRequest`), allowing features to inspect or modify the pipeline.

### Features

Features are hook-based middleware that extend SDK behaviour.

| Feature | Purpose |
| --- | --- |
| **TestFeature** | In-memory mock transport for testing without a live server |

You can add custom features by passing them in the `extend` option at
construction time.

### Direct and Prepare

For endpoints not covered by the entity model, use the low-level methods:

- **`direct(fetchargs)`** — build and send an HTTP request in one step.
- **`prepare(fetchargs)`** — build the request without sending it.

Both accept a map with `path`, `method`, `params`, `query`, `headers`,
and `body`.


## Quick start

### Golang

```go
import sdk "github.com/voxgig-sdk/art-institute-of-chicago-sdk/go"

client := sdk.NewArtInstituteOfChicagoSDK(map[string]any{
    "apikey": os.Getenv("ART-INSTITUTE-OF-CHICAGO_APIKEY"),
})

// List all agents
agents, err := client.Agent(nil).List(nil, nil)
```

### Lua

```lua
local sdk = require("art-institute-of-chicago_sdk")

local client = sdk.new({
  apikey = os.getenv("ART-INSTITUTE-OF-CHICAGO_APIKEY"),
})

-- List all agents
local agents, err = client:Agent(nil):list(nil, nil)

-- Load a specific agent
local agent, err = client:Agent(nil):load(
  { id = "example_id" }, nil
)
```

### PHP

```php
<?php
require_once 'artinstituteofchicago_sdk.php';

$client = new ArtInstituteOfChicagoSDK([
    "apikey" => getenv("ART-INSTITUTE-OF-CHICAGO_APIKEY"),
]);

// List all agents
[$agents, $err] = $client->Agent(null)->list(null, null);

// Load a specific agent
[$agent, $err] = $client->Agent(null)->load(
    ["id" => "example_id"], null
);
```

### Python

```python
import os
from artinstituteofchicago_sdk import ArtInstituteOfChicagoSDK

client = ArtInstituteOfChicagoSDK({
    "apikey": os.environ.get("ART-INSTITUTE-OF-CHICAGO_APIKEY"),
})

# List all agents
agents, err = client.Agent(None).list(None, None)

# Load a specific agent
agent, err = client.Agent(None).load(
    {"id": "example_id"}, None
)
```

### Ruby

```ruby
require_relative "ArtInstituteOfChicago_sdk"

client = ArtInstituteOfChicagoSDK.new({
  "apikey" => ENV["ART-INSTITUTE-OF-CHICAGO_APIKEY"],
})

# List all agents
agents, err = client.Agent(nil).list(nil, nil)

# Load a specific agent
agent, err = client.Agent(nil).load(
  { "id" => "example_id" }, nil
)
```

### TypeScript

```ts
import { ArtInstituteOfChicagoSDK } from 'art-institute-of-chicago'

const client = new ArtInstituteOfChicagoSDK({
  apikey: process.env.ART-INSTITUTE-OF-CHICAGO_APIKEY,
})

// List all agents
const agents = await client.Agent().list()
```


## Testing

Both SDKs provide a test mode that replaces the HTTP transport with an
in-memory mock, so tests run without a network connection.

### Golang

```go
client := sdk.TestSDK(nil, nil)
result, err := client.Agent(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Lua

```lua
local client = sdk.test(nil, nil)
local result, err = client:Agent(nil):load(
  { id = "test01" }, nil
)
```

### PHP

```php
$client = ArtInstituteOfChicagoSDK::test(null, null);
[$result, $err] = $client->Agent(null)->load(
    ["id" => "test01"], null
);
```

### Python

```python
client = ArtInstituteOfChicagoSDK.test(None, None)
result, err = client.Agent(None).load(
    {"id": "test01"}, None
)
```

### Ruby

```ruby
client = ArtInstituteOfChicagoSDK.test(nil, nil)
result, err = client.Agent(nil).load(
  { "id" => "test01" }, nil
)
```

### TypeScript

```ts
const client = ArtInstituteOfChicagoSDK.test()
const result = await client.Agent().load({ id: 'test01' })
// result.ok === true, result.data contains mock data
```


## How-to guides

### Make a direct API call

When the entity interface does not cover an endpoint, use `direct`:

**Go:**
```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
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

**PHP:**
```php
[$result, $err] = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);
```

**Python:**
```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
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

**TypeScript:**
```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
console.log(result.data)
```


## Language-specific documentation

- [Golang SDK](go/README.md)
- [Go CLI SDK](go-cli/README.md)
- [Go MCP server SDK](go-mcp/README.md)
- [Lua SDK](lua/README.md)
- [PHP SDK](php/README.md)
- [Python SDK](py/README.md)
- [Ruby SDK](rb/README.md)
- [TypeScript SDK](ts/README.md)

