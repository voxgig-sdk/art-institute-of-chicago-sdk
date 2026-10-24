package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	sdk "github.com/voxgig-sdk/art-institute-of-chicago-sdk/go"
)

// ListArgs is what an agent sends to art-institute-of-chicago_list.
type ListArgs struct {
	Entity string         `json:"entity" jsonschema:"one of: agent | agent_role | agent_type | article | artwork | artwork_date_qualifier | artwork_place_qualifier | artwork_type | category_term | digital_publication | digital_publication_article | educator_resource | event | event_occurrence | event_program | exhibition | gallery | generic_page | highlight | hour | image | landing_page | place | press_release | printed_publication | product | publication | search | section | site | sound | static_page | text | tour | video"`
	Query  map[string]any `json:"query,omitempty" jsonschema:"optional filter map; omit it for the first page"`
}

// LoadArgs is what an agent sends to art-institute-of-chicago_load.
type LoadArgs struct {
	Entity string         `json:"entity" jsonschema:"one of: agent | agent_role | agent_type | article | artwork | artwork_date_qualifier | artwork_place_qualifier | artwork_type | category_term | digital_publication | digital_publication_article | educator_resource | event | event_occurrence | event_program | exhibition | gallery | generic_page | highlight | hour | image | landing_page | place | press_release | printed_publication | product | publication | section | site | sound | static_page | text | tour | video"`
	Query  map[string]any `json:"query" jsonschema:"match map naming the record, such as {\"id\":1}"`
}

func registerTools(server *mcp.Server, client *sdk.ArtInstituteOfChicagoSDK) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "art-institute-of-chicago_list",
		Description: "List records from ArtInstituteOfChicago. Args: entity, query (optional filter map; omit it for the first page). Returns the first page of records as JSON.",
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true},
		InputSchema: entitySchema[ListArgs]("agent", "agent_role", "agent_type", "article", "artwork", "artwork_date_qualifier", "artwork_place_qualifier", "artwork_type", "category_term", "digital_publication", "digital_publication_article", "educator_resource", "event", "event_occurrence", "event_program", "exhibition", "gallery", "generic_page", "highlight", "hour", "image", "landing_page", "place", "press_release", "printed_publication", "product", "publication", "search", "section", "site", "sound", "static_page", "text", "tour", "video"),
	}, func(ctx context.Context, req *mcp.CallToolRequest, args ListArgs) (*mcp.CallToolResult, any, error) {
		return runOp(ctx, client, "list", args.Entity, args.Query)
	})
	mcp.AddTool(server, &mcp.Tool{
		Name:        "art-institute-of-chicago_load",
		Description: "Load one record from ArtInstituteOfChicago. Args: entity, query (match map naming the record, such as {\"id\":1}). Returns the record as JSON.",
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true},
		InputSchema: entitySchema[LoadArgs]("agent", "agent_role", "agent_type", "article", "artwork", "artwork_date_qualifier", "artwork_place_qualifier", "artwork_type", "category_term", "digital_publication", "digital_publication_article", "educator_resource", "event", "event_occurrence", "event_program", "exhibition", "gallery", "generic_page", "highlight", "hour", "image", "landing_page", "place", "press_release", "printed_publication", "product", "publication", "section", "site", "sound", "static_page", "text", "tour", "video"),
	}, func(ctx context.Context, req *mcp.CallToolRequest, args LoadArgs) (*mcp.CallToolResult, any, error) {
		return runOp(ctx, client, "load", args.Entity, args.Query)
	})
}

// entitySchema is the schema inferred from In, its entity limited to the
// entities the tool serves.
func entitySchema[In any](names ...string) *jsonschema.Schema {
	schema, err := jsonschema.For[In](nil)
	if err != nil {
		panic(err)
	}
	enum := make([]any, len(names))
	for i, name := range names {
		enum[i] = name
	}
	schema.Properties["entity"].Enum = enum
	return schema
}

func runOp(_ context.Context, client *sdk.ArtInstituteOfChicagoSDK, op string, entity string, input map[string]any) (*mcp.CallToolResult, any, error) {
	ent, err := entityFor(client, entity)
	if err != nil {
		return toolError(err.Error())
	}

	var result any
	switch op {
	case "list":
		result, err = ent.List(input, nil)
	case "load":
		result, err = ent.Load(input, nil)
	case "create":
		result, err = ent.Create(input, nil)
	case "update":
		result, err = ent.Update(input, nil)
	case "patch":
		result, err = ent.Patch(input, nil)
	case "remove":
		result, err = ent.Remove(input, nil)
	default:
		return toolError(fmt.Sprintf("unknown op %q", op))
	}
	if err != nil {
		return toolError(err.Error())
	}

	// SDK returns *Entity wrappers; unwrap each via .Data() to get a
	// plain map[string]any (or []any of maps for list) suitable for
	// JSON marshalling.
	data := extractData(result)
	body, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return toolError(fmt.Sprintf("marshal: %v", err))
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(body)},
		},
	}, data, nil
}

// entityFor dispatches on the lowercase entity name. The generator
// emits one `case "<name>":` per entity defined in the SDK model.
func entityFor(client *sdk.ArtInstituteOfChicagoSDK, name string) (sdk.ArtInstituteOfChicagoEntity, error) {
	switch strings.ToLower(name) {
	case "agent":
		return client.Agent(nil), nil
	case "agent_role":
		return client.AgentRole(nil), nil
	case "agent_type":
		return client.AgentType(nil), nil
	case "article":
		return client.Article(nil), nil
	case "artwork":
		return client.Artwork(nil), nil
	case "artwork_date_qualifier":
		return client.ArtworkDateQualifier(nil), nil
	case "artwork_place_qualifier":
		return client.ArtworkPlaceQualifier(nil), nil
	case "artwork_type":
		return client.ArtworkType(nil), nil
	case "category_term":
		return client.CategoryTerm(nil), nil
	case "digital_publication":
		return client.DigitalPublication(nil), nil
	case "digital_publication_article":
		return client.DigitalPublicationArticle(nil), nil
	case "educator_resource":
		return client.EducatorResource(nil), nil
	case "event":
		return client.Event(nil), nil
	case "event_occurrence":
		return client.EventOccurrence(nil), nil
	case "event_program":
		return client.EventProgram(nil), nil
	case "exhibition":
		return client.Exhibition(nil), nil
	case "gallery":
		return client.Gallery(nil), nil
	case "generic_page":
		return client.GenericPage(nil), nil
	case "highlight":
		return client.Highlight(nil), nil
	case "hour":
		return client.Hour(nil), nil
	case "image":
		return client.Image(nil), nil
	case "landing_page":
		return client.LandingPage(nil), nil
	case "place":
		return client.Place(nil), nil
	case "press_release":
		return client.PressRelease(nil), nil
	case "printed_publication":
		return client.PrintedPublication(nil), nil
	case "product":
		return client.Product(nil), nil
	case "publication":
		return client.Publication(nil), nil
	case "search":
		return client.Search(nil), nil
	case "section":
		return client.Section(nil), nil
	case "site":
		return client.Site(nil), nil
	case "sound":
		return client.Sound(nil), nil
	case "static_page":
		return client.StaticPage(nil), nil
	case "text":
		return client.Text(nil), nil
	case "tour":
		return client.Tour(nil), nil
	case "video":
		return client.Video(nil), nil
	}
	return nil, fmt.Errorf("unknown entity %q", name)
}

func extractData(x any) any {
	switch v := x.(type) {
	case sdk.Entity:
		return extractData(v.Data())
	case []any:
		out := make([]any, len(v))
		for i, e := range v {
			out[i] = extractData(e)
		}
		return out
	case map[string]any:
		out := make(map[string]any, len(v))
		for k, vv := range v {
			out[k] = extractData(vv)
		}
		return out
	}
	return x
}

func toolError(msg string) (*mcp.CallToolResult, any, error) {
	return &mcp.CallToolResult{
		IsError: true,
		Content: []mcp.Content{
			&mcp.TextContent{Text: msg},
		},
	}, nil, nil
}

// hint is an MCP annotation that defaults to true unless stated.
func hint(b bool) *bool {
	return &b
}
