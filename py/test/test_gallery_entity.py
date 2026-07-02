# Gallery entity test

import json
import os
import time

import pytest

from utility.voxgig_struct import voxgig_struct as vs
from artinstituteofchicago_sdk import ArtInstituteOfChicagoSDK
from core import helpers

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class TestGalleryEntity:

    def test_should_create_instance(self):
        testsdk = ArtInstituteOfChicagoSDK.test(None, None)
        ent = testsdk.Gallery(None)
        assert ent is not None

    def test_should_run_basic_flow(self):
        setup = _gallery_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["list", "load"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "gallery." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        # The basic flow consumes synthetic IDs from the fixture. In live mode
        # without an *_ENTID env override, those IDs hit the live API and 4xx.
        if setup.get("synthetic_only"):
            pytest.skip("live entity test uses synthetic IDs from fixture — "
                        "set ARTINSTITUTEOFCHICAGO_TEST_GALLERY_ENTID JSON to run live")
        client = setup["client"]

        # Bootstrap entity data from existing test data.
        gallery_ref01_data_raw = vs.items(helpers.to_map(
            vs.getpath(setup["data"], "existing.gallery")))
        gallery_ref01_data = None
        if len(gallery_ref01_data_raw) > 0:
            gallery_ref01_data = helpers.to_map(gallery_ref01_data_raw[0][1])

        # LIST
        gallery_ref01_ent = client.Gallery(None)
        gallery_ref01_match = {}

        gallery_ref01_list_result, err = gallery_ref01_ent.list(gallery_ref01_match, None)
        assert err is None
        assert isinstance(gallery_ref01_list_result, list)

        # LOAD
        gallery_ref01_match_dt0 = {
            "id": gallery_ref01_data["id"],
        }
        gallery_ref01_data_dt0_loaded, err = gallery_ref01_ent.load(gallery_ref01_match_dt0, None)
        assert err is None
        gallery_ref01_data_dt0_load_result = helpers.to_map(gallery_ref01_data_dt0_loaded)
        assert gallery_ref01_data_dt0_load_result is not None
        assert gallery_ref01_data_dt0_load_result["id"] == gallery_ref01_data["id"]



def _gallery_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/gallery/GalleryTestData.json")
    with open(entity_data_file, "r") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = ArtInstituteOfChicagoSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["gallery01", "gallery02", "gallery03"],
        {
            "`$PACK`": ["", {
                "`$KEY`": "`$COPY`",
                "`$VAL`": ["`$FORMAT`", "upper", "`$COPY`"],
            }],
        }
    )

    # Detect ENTID env override before envOverride consumes it. When live
    # mode is on without a real override, the basic test runs against synthetic
    # IDs from the fixture and 4xx's. We surface this so the test can skip.
    _entid_env_raw = os.environ.get(
        "ARTINSTITUTEOFCHICAGO_TEST_GALLERY_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "ARTINSTITUTEOFCHICAGO_TEST_GALLERY_ENTID": idmap,
        "ARTINSTITUTEOFCHICAGO_TEST_LIVE": "FALSE",
        "ARTINSTITUTEOFCHICAGO_TEST_EXPLAIN": "FALSE",
        "ARTINSTITUTEOFCHICAGO_APIKEY": "NONE",
    })

    idmap_resolved = helpers.to_map(
        env.get("ARTINSTITUTEOFCHICAGO_TEST_GALLERY_ENTID"))
    if idmap_resolved is None:
        idmap_resolved = helpers.to_map(idmap)

    if env.get("ARTINSTITUTEOFCHICAGO_TEST_LIVE") == "TRUE":
        merged_opts = vs.merge([
            {
                "apikey": env.get("ARTINSTITUTEOFCHICAGO_APIKEY"),
            },
            extra or {},
        ])
        client = ArtInstituteOfChicagoSDK(helpers.to_map(merged_opts))

    _live = env.get("ARTINSTITUTEOFCHICAGO_TEST_LIVE") == "TRUE"
    return {
        "client": client,
        "data": entity_data,
        "idmap": idmap_resolved,
        "env": env,
        "explain": env.get("ARTINSTITUTEOFCHICAGO_TEST_EXPLAIN") == "TRUE",
        "live": _live,
        "synthetic_only": _live and not _idmap_overridden,
        "now": int(time.time() * 1000),
    }
