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

func TestTextEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Text(nil)
		if ent == nil {
			t.Fatal("expected non-nil TextEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := textBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "text." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set ARTINSTITUTEOFCHICAGO_TEST_TEXT_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		textRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.text", setup.data)))
		var textRef01Data map[string]any
		if len(textRef01DataRaw) > 0 {
			textRef01Data = core.ToMapAny(textRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = textRef01Data

		// LIST
		textRef01Ent := client.Text(nil)
		textRef01Match := map[string]any{}

		textRef01ListResult, err := textRef01Ent.List(textRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, textRef01ListOk := textRef01ListResult.([]any)
		if !textRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", textRef01ListResult)
		}

		// LOAD
		textRef01MatchDt0 := map[string]any{
			"id": textRef01Data["id"],
		}
		textRef01DataDt0Loaded, err := textRef01Ent.Load(textRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		textRef01DataDt0LoadResult := core.ToMapAny(textRef01DataDt0Loaded)
		if textRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if textRef01DataDt0LoadResult["id"] != textRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

	})
}

func textBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "text", "TextTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read text test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse text test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"text01", "text02", "text03"},
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
	entidEnvRaw := os.Getenv("ARTINSTITUTEOFCHICAGO_TEST_TEXT_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"ARTINSTITUTEOFCHICAGO_TEST_TEXT_ENTID": idmap,
		"ARTINSTITUTEOFCHICAGO_TEST_LIVE":      "FALSE",
		"ARTINSTITUTEOFCHICAGO_TEST_EXPLAIN":   "FALSE",
		"ARTINSTITUTEOFCHICAGO_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["ARTINSTITUTEOFCHICAGO_TEST_TEXT_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["ARTINSTITUTEOFCHICAGO_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
				"apikey": env["ARTINSTITUTEOFCHICAGO_APIKEY"],
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
