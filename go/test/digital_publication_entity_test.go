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

func TestDigitalPublicationEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.DigitalPublication(nil)
		if ent == nil {
			t.Fatal("expected non-nil DigitalPublicationEntity")
		}
	})

	// Feature #4: the entity Stream(action, ...) method runs the op pipeline and
	// returns a channel over result items. With the streaming feature active it
	// yields the feature's incremental output; otherwise it falls back to the
	// materialised list so Stream always yields.
	t.Run("stream", func(t *testing.T) {
		seed := map[string]any{
			"entity": map[string]any{
				"digital_publication": map[string]any{
					"s1": map[string]any{"id": "s1"},
					"s2": map[string]any{"id": "s2"},
					"s3": map[string]any{"id": "s3"},
				},
			},
		}

		// Fallback: streaming inactive -> yields the materialised list items.
		base := sdk.TestSDK(seed, nil)
		var seen []any
		for item := range base.DigitalPublication(nil).Stream("list", nil, nil) {
			seen = append(seen, item)
		}
		if len(seen) != 3 {
			t.Fatalf("expected 3 streamed items, got %d", len(seen))
		}

		// Inbound: streaming active -> yields each item from the feature iterator.
		hasStreaming := false
		if fm, ok := core.SharedConfig()["feature"].(map[string]any); ok {
			_, hasStreaming = fm["streaming"]
		}
		if hasStreaming {
			streamSdk := sdk.TestSDK(seed, map[string]any{
				"feature": map[string]any{"streaming": map[string]any{"active": true}},
			})
			var got []any
			for item := range streamSdk.DigitalPublication(nil).Stream("list", nil, nil) {
				if sub, ok := item.([]any); ok {
					got = append(got, sub...)
				} else {
					got = append(got, item)
				}
			}
			if len(got) != 3 {
				t.Fatalf("expected 3 items via streaming feature, got %d", len(got))
			}
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := digital_publicationBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "digital_publication." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set ART_INSTITUTE_OF_CHICAGO_TEST_DIGITAL_PUBLICATION_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		digitalPublicationRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.digital_publication", setup.data)))
		var digitalPublicationRef01Data map[string]any
		if len(digitalPublicationRef01DataRaw) > 0 {
			digitalPublicationRef01Data = core.ToMapAny(digitalPublicationRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = digitalPublicationRef01Data

		// LIST
		digitalPublicationRef01Ent := client.DigitalPublication(nil)
		digitalPublicationRef01Match := map[string]any{}

		digitalPublicationRef01ListResult, err := digitalPublicationRef01Ent.List(digitalPublicationRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, digitalPublicationRef01ListOk := digitalPublicationRef01ListResult.([]any)
		if !digitalPublicationRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", digitalPublicationRef01ListResult)
		}

		// LOAD
		digitalPublicationRef01MatchDt0 := map[string]any{
			"id": digitalPublicationRef01Data["id"],
		}
		digitalPublicationRef01DataDt0Loaded, err := digitalPublicationRef01Ent.Load(digitalPublicationRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		digitalPublicationRef01DataDt0LoadResult := core.ToMapAny(entityData(digitalPublicationRef01DataDt0Loaded))
		if digitalPublicationRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if digitalPublicationRef01DataDt0LoadResult["id"] != digitalPublicationRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

	})
}

func digital_publicationBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "digital_publication", "DigitalPublicationTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read digital_publication test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse digital_publication test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"digital_publication01", "digital_publication02", "digital_publication03"},
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
	entidEnvRaw := os.Getenv("ART_INSTITUTE_OF_CHICAGO_TEST_DIGITAL_PUBLICATION_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"ART_INSTITUTE_OF_CHICAGO_TEST_DIGITAL_PUBLICATION_ENTID": idmap,
		"ART_INSTITUTE_OF_CHICAGO_TEST_LIVE":      "FALSE",
		"ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN":   "FALSE",
	})

	idmapResolved := core.ToMapAny(env["ART_INSTITUTE_OF_CHICAGO_TEST_DIGITAL_PUBLICATION_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["ART_INSTITUTE_OF_CHICAGO_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
			},
			extra,
		})
		client = sdk.NewArtInstituteOfChicagoSDK(core.ToMapAny(mergedOpts))
	}

	live := env["ART_INSTITUTE_OF_CHICAGO_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
