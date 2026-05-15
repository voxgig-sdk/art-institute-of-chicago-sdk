# ArtInstituteOfChicago SDK feature factory

from feature.base_feature import ArtInstituteOfChicagoBaseFeature
from feature.test_feature import ArtInstituteOfChicagoTestFeature


def _make_feature(name):
    features = {
        "base": lambda: ArtInstituteOfChicagoBaseFeature(),
        "test": lambda: ArtInstituteOfChicagoTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
