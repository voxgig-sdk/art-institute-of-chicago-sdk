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

func TestPrintedPublicationEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.PrintedPublication(nil)
		if ent == nil {
			t.Fatal("expected non-nil PrintedPublicationEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := printed_publicationBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "printed_publication." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set ARTINSTITUTEOFCHICAGO_TEST_PRINTED_PUBLICATION_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		printedPublicationRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.printed_publication", setup.data)))
		var printedPublicationRef01Data map[string]any
		if len(printedPublicationRef01DataRaw) > 0 {
			printedPublicationRef01Data = core.ToMapAny(printedPublicationRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = printedPublicationRef01Data

		// LIST
		printedPublicationRef01Ent := client.PrintedPublication(nil)
		printedPublicationRef01Match := map[string]any{}

		printedPublicationRef01ListResult, err := printedPublicationRef01Ent.List(printedPublicationRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, printedPublicationRef01ListOk := printedPublicationRef01ListResult.([]any)
		if !printedPublicationRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", printedPublicationRef01ListResult)
		}

		// LOAD
		printedPublicationRef01MatchDt0 := map[string]any{
			"id": printedPublicationRef01Data["id"],
		}
		printedPublicationRef01DataDt0Loaded, err := printedPublicationRef01Ent.Load(printedPublicationRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		printedPublicationRef01DataDt0LoadResult := core.ToMapAny(printedPublicationRef01DataDt0Loaded)
		if printedPublicationRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if printedPublicationRef01DataDt0LoadResult["id"] != printedPublicationRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

	})
}

func printed_publicationBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "printed_publication", "PrintedPublicationTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read printed_publication test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse printed_publication test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"printed_publication01", "printed_publication02", "printed_publication03"},
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
	entidEnvRaw := os.Getenv("ARTINSTITUTEOFCHICAGO_TEST_PRINTED_PUBLICATION_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"ARTINSTITUTEOFCHICAGO_TEST_PRINTED_PUBLICATION_ENTID": idmap,
		"ARTINSTITUTEOFCHICAGO_TEST_LIVE":      "FALSE",
		"ARTINSTITUTEOFCHICAGO_TEST_EXPLAIN":   "FALSE",
	})

	idmapResolved := core.ToMapAny(env["ARTINSTITUTEOFCHICAGO_TEST_PRINTED_PUBLICATION_ENTID"])
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
