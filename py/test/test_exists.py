# ProjectName SDK exists test

import pytest
from artinstituteofchicago_sdk import ArtInstituteOfChicagoSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = ArtInstituteOfChicagoSDK.test(None, None)
        assert testsdk is not None
