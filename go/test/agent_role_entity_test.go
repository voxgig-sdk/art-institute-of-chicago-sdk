package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/art-institute-of-chicago-sdk/go"
	"github.com/voxgig-sdk/art-institute-of-chicago-sdk/go/core"

	vs "github.com/voxgig-sdk/art-institute-of-chicago-sdk/go/utility/struct"
)

func TestAgentRoleEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.AgentRole(nil)
		if ent == nil {
			t.Fatal("expected non-nil AgentRoleEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := agent_roleBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "agent_role." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set ARTINSTITUTEOFCHICAGO_TEST_AGENT_ROLE_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		agentRoleRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.agent_role", setup.data)))
		var agentRoleRef01Data map[string]any
		if len(agentRoleRef01DataRaw) > 0 {
			agentRoleRef01Data = core.ToMapAny(agentRoleRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = agentRoleRef01Data

		// LIST
		agentRoleRef01Ent := client.AgentRole(nil)
		agentRoleRef01Match := map[string]any{}

		agentRoleRef01ListResult, err := agentRoleRef01Ent.List(agentRoleRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, agentRoleRef01ListOk := agentRoleRef01ListResult.([]any)
		if !agentRoleRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", agentRoleRef01ListResult)
		}

		// LOAD
		agentRoleRef01MatchDt0 := map[string]any{
			"id": agentRoleRef01Data["id"],
		}
		agentRoleRef01DataDt0Loaded, err := agentRoleRef01Ent.Load(agentRoleRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		agentRoleRef01DataDt0LoadResult := core.ToMapAny(agentRoleRef01DataDt0Loaded)
		if agentRoleRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if agentRoleRef01DataDt0LoadResult["id"] != agentRoleRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

	})
}

func agent_roleBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "agent_role", "AgentRoleTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read agent_role test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse agent_role test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"agent_role01", "agent_role02", "agent_role03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("ARTINSTITUTEOFCHICAGO_TEST_AGENT_ROLE_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"ARTINSTITUTEOFCHICAGO_TEST_AGENT_ROLE_ENTID": idmap,
		"ARTINSTITUTEOFCHICAGO_TEST_LIVE":      "FALSE",
		"ARTINSTITUTEOFCHICAGO_TEST_EXPLAIN":   "FALSE",
	})

	idmapResolved := core.ToMapAny(env["ARTINSTITUTEOFCHICAGO_TEST_AGENT_ROLE_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["ARTINSTITUTEOFCHICAGO_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
			},
			extra,
		})
		client = sdk.NewArtInstituteOfChicagoSDK(core.ToMapAny(mergedOpts))
	}

	live := env["ARTINSTITUTEOFCHICAGO_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["ARTINSTITUTEOFCHICAGO_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
