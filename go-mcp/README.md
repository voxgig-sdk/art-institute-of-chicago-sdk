# art-institute-of-chicago-mcp

[MCP](https://modelcontextprotocol.io) server exposing the ArtInstituteOfChicago SDK as
2 agent tools, `art-institute-of-chicago_list` and `art-institute-of-chicago_load`, built on the
[official Go MCP SDK](https://github.com/modelcontextprotocol/go-sdk) and the
sibling Go SDK at `../go`. Runs over **stdio** (default, for spawnable installs)
or **streamable HTTP** (one shared server for several agents).

The server only reads. Create, update, patch and remove become tools too when the
SDK's own model sets `main: kit: target: 'go-mcp': tool: write: true`; they are off by default, as an agent calling
them changes the API's data.

## Examples

```sh
# 1. Build a native binary (-> dist/<os>-<arch>/art-institute-of-chicago-mcp)
make build

# 2. Provide credentials via the environment
export ART_INSTITUTE_OF_CHICAGO_APIKEY=sk_live_xxx

# 3a. Install into Claude Code over stdio (most common)
claude mcp add --scope user art-institute-of-chicago \
  -- /absolute/path/to/art-institute-of-chicago-mcp -transport stdio

# 3b. …or run a shared HTTP server instead
./art-institute-of-chicago-mcp -transport http -addr :8080
```

Tool-call arguments (what an agent sends):

```jsonc
// art-institute-of-chicago_list: first page of records
{ "entity": "agent" }

// art-institute-of-chicago_load: one record
{ "entity": "agent", "query": { "id": "example_id" } }
```

> The rest of this guide follows the [Diátaxis](https://diataxis.fr) framework:
> a hands-on **Tutorial**, task-focused **How-to guides**, a factual
> **Reference**, and background **Explanation**.

## Tutorial: install and call a tool

1. **Build** the server from this `go-mcp/` directory:

   ```sh
   make build          # -> dist/<os>-<arch>/art-institute-of-chicago-mcp
   ```

2. **Set your API key:**

   ```sh
   export ART_INSTITUTE_OF_CHICAGO_APIKEY=sk_live_xxx
   ```

3. **Install it into Claude Code** (stdio transport):

   ```sh
   claude mcp add --scope user art-institute-of-chicago \
     -- "$PWD"/dist/*/art-institute-of-chicago-mcp -transport stdio
   ```

4. **Restart Claude Code.** The `art-institute-of-chicago_list` and `art-institute-of-chicago_load` tools now appear in new
   sessions. Ask the agent to *"list agent using art-institute-of-chicago"*
   and it calls `art-institute-of-chicago_list` with `{"entity":"agent"}`.

## How-to guides

### Authenticate and choose an environment

Configuration is read from the environment — nothing is written to disk:

```sh
export ART_INSTITUTE_OF_CHICAGO_APIKEY=sk_live_xxx            # API key
export ART_INSTITUTE_OF_CHICAGO_BASE=https://api.example.com  # optional: override the API base URL
```

Set these in the shell that launches the server (or in the `claude mcp add`
environment) so every tool call is authenticated.

### Run as a shared HTTP server

```sh
./art-institute-of-chicago-mcp -transport http -addr :8080
```

Streamable HTTP lets several agents share one running process; stdio (the
default) spawns a fresh process per client.

### Call the `art-institute-of-chicago_list` tool

Args: `entity` (required), `query` (optional: optional filter map; omit it for the first page).
Returns the first page of records as JSON:

```jsonc
{ "entity": "agent" }
```

### Call the `art-institute-of-chicago_load` tool

Args: `entity` (required), `query` (required: match map naming the record, such as {"id":1}).
Returns the record as JSON:

```jsonc
{ "entity": "agent", "query": { "id": "example_id" } }
```

### Turn on the write tools

In the SDK's own model (`.sdk/model/sdk.aontu`), then regenerate:

```
main: kit: target: 'go-mcp': tool: write: true
```

### Cross-compile release binaries

```sh
make build       # native binary for this machine
make build-all   # linux/darwin/windows x amd64/arm64, under dist/<os>-<arch>/
```

## Reference

### Tools

| Tool | Args | Returns | MCP hints |
|------|------|---------|-----------|
| `art-institute-of-chicago_list` | `entity`, `query` (optional map) | The first page of records as JSON | read-only |
| `art-institute-of-chicago_load` | `entity`, `query` (required map) | The record as JSON | read-only |

On error, a tool returns an MCP error result (`isError: true`) whose text is the
failure message (e.g. unknown entity, or an API error).

### Entities

Each tool takes as its `entity` argument one of the entities that has its
operation, of the 35 the SDK has:

| Tool | Entities |
|------|----------|
| `art-institute-of-chicago_list` | agent, agent_role, agent_type, article, artwork, artwork_date_qualifier, artwork_place_qualifier, artwork_type, category_term, digital_publication, digital_publication_article, educator_resource, event, event_occurrence, event_program, exhibition, gallery, generic_page, highlight, hour, image, landing_page, place, press_release, printed_publication, product, publication, search, section, site, sound, static_page, text, tour, video |
| `art-institute-of-chicago_load` | agent, agent_role, agent_type, article, artwork, artwork_date_qualifier, artwork_place_qualifier, artwork_type, category_term, digital_publication, digital_publication_article, educator_resource, event, event_occurrence, event_program, exhibition, gallery, generic_page, highlight, hour, image, landing_page, place, press_release, printed_publication, product, publication, section, site, sound, static_page, text, tour, video |

JSON schemas are emitted by the SDK from each tool's argument struct's
`json` / `jsonschema` tags — no schema is hand-written. Each tool's
`entity` is an `enum` of the entities in its row, so the server refuses
any other before it runs a call.

### Transports & flags

| Flag | Default | Purpose |
|------|---------|---------|
| `-transport` | `stdio` | `stdio` (spawnable) or `http` (streamable HTTP). |
| `-addr` | `:8080` | Listen address for the `http` transport. |

### Environment variables

| Variable | Purpose |
|----------|---------|
| `ART_INSTITUTE_OF_CHICAGO_APIKEY` | API key sent with every request. |
| `ART_INSTITUTE_OF_CHICAGO_BASE` | Optional override of the API base URL. |

### Smoke test via HTTP (raw JSON-RPC)

```sh
./art-institute-of-chicago-mcp -transport http -addr :18080 &

# initialize, grab the session id
curl -sN -X POST http://localhost:18080 \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json, text/event-stream' \
  -D headers \
  -d '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2025-06-18","capabilities":{},"clientInfo":{"name":"smoke","version":"0"}}}'

SESSION=$(awk '/Mcp-Session-Id/ {print $2}' headers | tr -d '\r')

curl -sN -X POST http://localhost:18080 \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json, text/event-stream' \
  -H "Mcp-Session-Id: $SESSION" \
  -d '{"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"art-institute-of-chicago_list","arguments":{"entity":"agent"}}}'
```

## Explanation

### How tools map to the SDK

`main.go` builds the SDK client (configured from the environment) and registers
one tool per operation the SDK's entities have. Each dispatches on the
`entity` argument to the matching entity in the sibling Go SDK at `../go`,
calls its operation, unwraps the `Entity` wrappers to plain data, and returns
it as pretty-printed JSON.

### Why two transports

**stdio** is the standard for agent hosts that spawn a server per client
(Claude Code's `claude mcp add`). **streamable HTTP** keeps one process running
that many agents can share — handy for a long-lived deployment.

### Schema generation

The input schema is derived from each tool's argument struct's `json` /
`jsonschema` tags at registration time, so the advertised tool schema can
never drift from the code that consumes it. The `entity` enum comes from the
same list the tool is registered for.

## Generated by

sdkgen `go-mcp` target. See the target source under `.sdk/src/cmp/go-mcp/` in
this repo, or upstream at
`github.com/voxgig/sdkgen/project/.sdk/src/cmp/go-mcp/`.
