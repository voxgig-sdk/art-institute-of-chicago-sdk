# ArtInstituteOfChicago SDK utility: make_context

from projectname_sdk.core.context import ArtInstituteOfChicagoContext


def make_context_util(ctxmap, basectx):
    return ArtInstituteOfChicagoContext(ctxmap, basectx)
