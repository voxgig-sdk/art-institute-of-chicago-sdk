# EventOccurrence entity test

import json
import os
import time

import pytest

from artinstituteofchicago_sdk.utility.voxgig_struct import voxgig_struct as vs
from artinstituteofchicago_sdk import ArtInstituteOfChicagoSDK
from artinstituteofchicago_sdk.core import helpers
from artinstituteofchicago_sdk.config import shared_config
from artinstituteofchicago_sdk.feature.base_feature import ArtInstituteOfChicagoBaseFeature

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class _FailHook(ArtInstituteOfChicagoBaseFeature):
    def __init__(self):
        super().__init__()
        self.name = "failhook"
        self.unexpected = 0

    def init(self, ctx, options):
        pass

    def PreSpec(self, ctx):
        raise RuntimeError("event_occurrence hook failed")

    def PreUnexpected(self, ctx):
        self.unexpected += 1



# main.kit.test.live.strict is true (the default is true): a live
# request that fails, or a live test missing an input it needs,
# fails the test.
# An account with no record for a test to read skips it either way.
LIVE_STRICT = True


class TestEventOccurrenceEntity:

    def test_should_create_instance(self):
        testsdk = ArtInstituteOfChicagoSDK.test(None, None)
        ent = testsdk.EventOccurrence(None)
        assert ent is not None

    def test_should_stream(self):
        # Feature #4: the entity stream(action, ...) method runs the op
        # pipeline and yields result items. With the streaming feature active
        # it yields the feature's incremental output; otherwise it falls back
        # to the materialised list so stream always yields.
        seed = {
            "entity": {
                "event_occurrence": {
                    "s1": {"id": "s1"},
                    "s2": {"id": "s2"},
                    "s3": {"id": "s3"},
                }
            }
        }

        # Fallback: streaming inactive -> yields the materialised list items.
        base = ArtInstituteOfChicagoSDK.test(seed, None)
        seen = list(base.EventOccurrence(None).stream("list", None, None))
        assert len(seen) == 3

        # Inbound: streaming active -> yields each item from the feature.
        from artinstituteofchicago_sdk.config import shared_config
        cfg = shared_config()
        if isinstance(cfg.get("feature"), dict) and "streaming" in cfg["feature"]:
            sdk = ArtInstituteOfChicagoSDK.test(
                seed, {"feature": {"streaming": {"active": True}}})
            got = []
            for item in sdk.EventOccurrence(None).stream("list", None, None):
                if isinstance(item, list):
                    got.extend(item)
                else:
                    got.append(item)
            assert len(got) == 3

    def test_should_report_a_failed_stream(self):
        offline = {"net": {"offline": True}}
        with pytest.raises(Exception, match="offline"):
            list(ArtInstituteOfChicagoSDK.test(offline, None).EventOccurrence(None).stream("list", None, None))

        quiet = {"ctrl": {"throw": False}}
        list(ArtInstituteOfChicagoSDK.test(offline, None).EventOccurrence(None).stream("list", None, quiet))

        if "rbac" in (shared_config().get("feature") or {}):
            denied = ArtInstituteOfChicagoSDK.test(
                None, {"feature": {"rbac": {"active": True, "deny": True}}})
            with pytest.raises(Exception) as err:
                list(denied.EventOccurrence(None).stream("list", None, None))
            assert "rbac_denied" == getattr(err.value, "code", None)

    def test_should_leave_the_callers_ctrl(self):
        explain = {}
        ctrl = {"explain": explain}
        list(ArtInstituteOfChicagoSDK.test(None, None).EventOccurrence(None).stream("list", None, {"ctrl": ctrl}))
        assert ["explain"] == list(ctrl.keys())
        assert explain is ctrl["explain"] and 0 < len(explain)

    def test_should_fire_pre_unexpected(self):
        hook = _FailHook()
        client = ArtInstituteOfChicagoSDK({"feature": {"test": {"active": True}}, "extend": [hook]})
        with pytest.raises(Exception, match="hook failed"):
            client.EventOccurrence(None).list(None, None)
        assert 0 < hook.unexpected

        fired = hook.unexpected
        assert client.EventOccurrence(None).list(None, {"throw": False}) is None
        assert fired < hook.unexpected

    def test_should_validate_each_listed_record(self):
        if "validate" not in (shared_config().get("feature") or {}):
            pytest.skip("feature not present in this SDK: validate")
        opts = {"feature": {"validate": {"active": True, "request": False, "response": True}}}
        listed = ArtInstituteOfChicagoSDK.test({"entity": {"event_occurrence": {"v1": {"api_link": "x", "api_model": "x", "button_caption": "x", "button_text": "x", "button_url": "x", "end_at": "x", "image_url": "x", "location": "x", "off_sale_at": "x", "on_sale_at": "x", "short_description": "x", "source_updated_at": "x", "start_at": "x", "suggest_autocomplete_all": "x", "suggest_autocomplete_boosted": "x", "timestamp": "x", "title_display": "x", "updated_at": "x"}}}}, opts).EventOccurrence(None).list(None, None)
        assert ["v1"] == [ent.data_get().get("id") for ent in listed]
        with pytest.raises(Exception) as err:
            ArtInstituteOfChicagoSDK.test({"entity": {"event_occurrence": {"v1": {"api_link": "x", "api_model": "x", "button_caption": "x", "button_text": "x", "button_url": "x", "end_at": "x", "image_url": "x", "location": "x", "off_sale_at": "x", "on_sale_at": "x", "short_description": "x", "source_updated_at": "x", "start_at": "x", "suggest_autocomplete_all": "x", "suggest_autocomplete_boosted": "x", "timestamp": "x", "title_display": "x", "updated_at": "x", "description": 1}}}}, opts).EventOccurrence(None).list(None, None)
        assert "validate_failed" == getattr(err.value, "code", None)

    def test_should_refuse_an_invalid_request(self):
        if "validate" not in (shared_config().get("feature") or {}):
            pytest.skip("feature not present in this SDK: validate")
        client = ArtInstituteOfChicagoSDK.test(
            None, {"feature": {"validate": {"active": True}}})
        with pytest.raises(Exception) as err:
            client.EventOccurrence(None).list({"description": 1}, None)
        assert "validate_failed" == getattr(err.value, "code", None)

    def test_should_run_basic_flow(self):
        setup = _event_occurrence_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["list", "load"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "event_occurrence." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        client = setup["client"]
        if setup["live"]:
            runner.live_existing(setup, LIVE_STRICT, "event_occurrence",
                                 lambda: client.EventOccurrence(None).list({}, None))

        # Bootstrap entity data from existing test data.
        event_occurrence_ref01_data_raw = vs.items(helpers.to_map(
            vs.getpath(setup["data"], "existing.event_occurrence")))
        event_occurrence_ref01_data = None
        if len(event_occurrence_ref01_data_raw) > 0:
            event_occurrence_ref01_data = helpers.to_map(event_occurrence_ref01_data_raw[0][1])

        # LIST
        event_occurrence_ref01_ent = client.EventOccurrence(None)
        event_occurrence_ref01_match = {}

        event_occurrence_ref01_list_result = event_occurrence_ref01_ent.list(event_occurrence_ref01_match, None)
        assert isinstance(event_occurrence_ref01_list_result, list)

        # LOAD
        event_occurrence_ref01_match_dt0 = {
            "id": event_occurrence_ref01_data["id"],
        }
        event_occurrence_ref01_data_dt0_loaded = event_occurrence_ref01_ent.load(event_occurrence_ref01_match_dt0, None)
        event_occurrence_ref01_data_dt0_load_result = helpers.to_map(runner.entity_data(event_occurrence_ref01_data_dt0_loaded))
        assert event_occurrence_ref01_data_dt0_load_result is not None
        assert event_occurrence_ref01_data_dt0_load_result["id"] == event_occurrence_ref01_data["id"]



def _event_occurrence_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/event_occurrence/EventOccurrenceTestData.json")
    with open(entity_data_file, "r", encoding="utf-8") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = ArtInstituteOfChicagoSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["event_occurrence01", "event_occurrence02", "event_occurrence03"],
        {
            "`$PACK`": ["", {
                "`$KEY`": "`$COPY`",
                "`$VAL`": ["`$FORMAT`", "upper", "`$COPY`"],
            }],
        }
    )

    # Whether *_ENTID supplied the idmap, read before env_override consumes
    # it: without it, the ids a live flow binds are the fixture's synthetic ones.
    _entid_env_raw = os.environ.get(
        "ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_OCCURRENCE_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_OCCURRENCE_ENTID": idmap,
        "ART_INSTITUTE_OF_CHICAGO_TEST_LIVE": "FALSE",
        "ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN": "FALSE",
    })

    idmap_resolved = helpers.to_map(
        env.get("ART_INSTITUTE_OF_CHICAGO_TEST_EVENT_OCCURRENCE_ENTID"))
    if idmap_resolved is None:
        idmap_resolved = helpers.to_map(idmap)

    if env.get("ART_INSTITUTE_OF_CHICAGO_TEST_LIVE") == "TRUE":
        merged_opts = vs.merge([
            # FIRST, so the generated fields below win: sdk-test-control.json's
            # test.client.options adds to the live client, it does not
            # redirect it.
            runner.live_client_options(),
            {
            },
            extra or {},
        ])
        client = ArtInstituteOfChicagoSDK(helpers.to_map(merged_opts))

    _live = env.get("ART_INSTITUTE_OF_CHICAGO_TEST_LIVE") == "TRUE"
    return {
        "client": client,
        "data": entity_data,
        "idmap": idmap_resolved,
        "env": env,
        "explain": env.get("ART_INSTITUTE_OF_CHICAGO_TEST_EXPLAIN") == "TRUE",
        "live": _live,
        "synthetic_only": _live and not _idmap_overridden,
        "now": int(time.time() * 1000),
    }
