# ArtInstituteOfChicago SDK

from utility.voxgig_struct import voxgig_struct as vs
from core.utility_type import ArtInstituteOfChicagoUtility
from core.spec import ArtInstituteOfChicagoSpec
from core import helpers

# Load utility registration (populates Utility._registrar)
from utility import register

# Load features
from feature.base_feature import ArtInstituteOfChicagoBaseFeature
from features import _make_feature


class ArtInstituteOfChicagoSDK:

    def __init__(self, options=None):
        self.mode = "live"
        self.features = []
        self.options = None

        utility = ArtInstituteOfChicagoUtility()
        self._utility = utility

        from config import make_config
        config = make_config()

        self._rootctx = utility.make_context({
            "client": self,
            "utility": utility,
            "config": config,
            "options": options if options is not None else {},
            "shared": {},
        }, None)

        self.options = utility.make_options(self._rootctx)

        if vs.getpath(self.options, "feature.test.active") is True:
            self.mode = "test"

        self._rootctx.options = self.options

        # Add features from config.
        feature_opts = helpers.to_map(vs.getprop(self.options, "feature"))
        if feature_opts is not None:
            feature_items = vs.items(feature_opts)
            if feature_items is not None:
                for item in feature_items:
                    fname = item[0]
                    fopts = helpers.to_map(item[1])
                    if fopts is not None and fopts.get("active") is True:
                        utility.feature_add(self._rootctx, _make_feature(fname))

        # Add extension features.
        extend = vs.getprop(self.options, "extend")
        if isinstance(extend, list):
            for f in extend:
                if isinstance(f, dict) or (hasattr(f, "get_name") and callable(f.get_name)):
                    utility.feature_add(self._rootctx, f)

        # Initialize features.
        for f in self.features:
            utility.feature_init(self._rootctx, f)

        utility.feature_hook(self._rootctx, "PostConstruct")

        # #BuildFeatures

    def options_map(self):
        out = vs.clone(self.options)
        if isinstance(out, dict):
            return out
        return {}

    def get_utility(self):
        return ArtInstituteOfChicagoUtility.copy(self._utility)

    def get_root_ctx(self):
        return self._rootctx

    def prepare(self, fetchargs=None):
        utility = self._utility

        if fetchargs is None:
            fetchargs = {}

        ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl"))
        if ctrl is None:
            ctrl = {}

        ctx = utility.make_context({
            "opname": "prepare",
            "ctrl": ctrl,
        }, self._rootctx)

        options = self.options

        path = vs.getprop(fetchargs, "path") or ""
        if not isinstance(path, str):
            path = ""

        method = vs.getprop(fetchargs, "method") or "GET"
        if not isinstance(method, str):
            method = "GET"

        params = helpers.to_map(vs.getprop(fetchargs, "params"))
        if params is None:
            params = {}
        query = helpers.to_map(vs.getprop(fetchargs, "query"))
        if query is None:
            query = {}

        headers = utility.prepare_headers(ctx)

        base = vs.getprop(options, "base") or ""
        if not isinstance(base, str):
            base = ""
        prefix = vs.getprop(options, "prefix") or ""
        if not isinstance(prefix, str):
            prefix = ""
        suffix = vs.getprop(options, "suffix") or ""
        if not isinstance(suffix, str):
            suffix = ""

        ctx.spec = ArtInstituteOfChicagoSpec({
            "base": base,
            "prefix": prefix,
            "suffix": suffix,
            "path": path,
            "method": method,
            "params": params,
            "query": query,
            "headers": headers,
            "body": vs.getprop(fetchargs, "body"),
            "step": "start",
        })

        # Merge user-provided headers.
        uh = vs.getprop(fetchargs, "headers")
        if isinstance(uh, dict):
            for k, v in uh.items():
                ctx.spec.headers[k] = v

        _, err = utility.prepare_auth(ctx)
        if err is not None:
            return None, err

        return utility.make_fetch_def(ctx)

    def direct(self, fetchargs=None):
        utility = self._utility

        fetchdef, err = self.prepare(fetchargs)
        if err is not None:
            return {"ok": False, "err": err}, None

        if fetchargs is None:
            fetchargs = {}
        ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl"))
        if ctrl is None:
            ctrl = {}

        ctx = utility.make_context({
            "opname": "direct",
            "ctrl": ctrl,
        }, self._rootctx)

        url = fetchdef.get("url", "")
        fetched, fetch_err = utility.fetcher(ctx, url, fetchdef)

        if fetch_err is not None:
            return {"ok": False, "err": fetch_err}, None

        if fetched is None:
            return {
                "ok": False,
                "err": ctx.make_error("direct_no_response", "response: undefined"),
            }, None

        if isinstance(fetched, dict):
            status = helpers.to_int(vs.getprop(fetched, "status"))
            headers = vs.getprop(fetched, "headers") or {}

            # No-body responses (204, 304) and explicit zero content-length
            # must skip JSON parsing — calling json() on an empty body raises.
            content_length = None
            if isinstance(headers, dict):
                content_length = headers.get("content-length")
            no_body = status in (204, 304) or str(content_length) == "0"

            json_data = None
            if not no_body:
                jf = vs.getprop(fetched, "json")
                if callable(jf):
                    try:
                        json_data = jf()
                    except Exception:
                        # Non-JSON body (e.g. text/plain, text/html). Surface
                        # status + headers but leave data as None.
                        json_data = None

            return {
                "ok": status >= 200 and status < 300,
                "status": status,
                "headers": headers,
                "data": json_data,
            }, None

        return {
            "ok": False,
            "err": ctx.make_error("direct_invalid", "invalid response type"),
        }, None


    def Agent(self, data=None):
        from entity.agent_entity import AgentEntity
        return AgentEntity(self, data)


    def AgentRole(self, data=None):
        from entity.agent_role_entity import AgentRoleEntity
        return AgentRoleEntity(self, data)


    def AgentType(self, data=None):
        from entity.agent_type_entity import AgentTypeEntity
        return AgentTypeEntity(self, data)


    def Article(self, data=None):
        from entity.article_entity import ArticleEntity
        return ArticleEntity(self, data)


    def Artwork(self, data=None):
        from entity.artwork_entity import ArtworkEntity
        return ArtworkEntity(self, data)


    def ArtworkDateQualifier(self, data=None):
        from entity.artwork_date_qualifier_entity import ArtworkDateQualifierEntity
        return ArtworkDateQualifierEntity(self, data)


    def ArtworkPlaceQualifier(self, data=None):
        from entity.artwork_place_qualifier_entity import ArtworkPlaceQualifierEntity
        return ArtworkPlaceQualifierEntity(self, data)


    def ArtworkType(self, data=None):
        from entity.artwork_type_entity import ArtworkTypeEntity
        return ArtworkTypeEntity(self, data)


    def CategoryTerm(self, data=None):
        from entity.category_term_entity import CategoryTermEntity
        return CategoryTermEntity(self, data)


    def DigitalPublication(self, data=None):
        from entity.digital_publication_entity import DigitalPublicationEntity
        return DigitalPublicationEntity(self, data)


    def DigitalPublicationArticle(self, data=None):
        from entity.digital_publication_article_entity import DigitalPublicationArticleEntity
        return DigitalPublicationArticleEntity(self, data)


    def EducatorResource(self, data=None):
        from entity.educator_resource_entity import EducatorResourceEntity
        return EducatorResourceEntity(self, data)


    def Event(self, data=None):
        from entity.event_entity import EventEntity
        return EventEntity(self, data)


    def EventOccurrence(self, data=None):
        from entity.event_occurrence_entity import EventOccurrenceEntity
        return EventOccurrenceEntity(self, data)


    def EventProgram(self, data=None):
        from entity.event_program_entity import EventProgramEntity
        return EventProgramEntity(self, data)


    def Exhibition(self, data=None):
        from entity.exhibition_entity import ExhibitionEntity
        return ExhibitionEntity(self, data)


    def Gallery(self, data=None):
        from entity.gallery_entity import GalleryEntity
        return GalleryEntity(self, data)


    def GenericPage(self, data=None):
        from entity.generic_page_entity import GenericPageEntity
        return GenericPageEntity(self, data)


    def Highlight(self, data=None):
        from entity.highlight_entity import HighlightEntity
        return HighlightEntity(self, data)


    def Hour(self, data=None):
        from entity.hour_entity import HourEntity
        return HourEntity(self, data)


    def Image(self, data=None):
        from entity.image_entity import ImageEntity
        return ImageEntity(self, data)


    def LandingPage(self, data=None):
        from entity.landing_page_entity import LandingPageEntity
        return LandingPageEntity(self, data)


    def Place(self, data=None):
        from entity.place_entity import PlaceEntity
        return PlaceEntity(self, data)


    def PressRelease(self, data=None):
        from entity.press_release_entity import PressReleaseEntity
        return PressReleaseEntity(self, data)


    def PrintedPublication(self, data=None):
        from entity.printed_publication_entity import PrintedPublicationEntity
        return PrintedPublicationEntity(self, data)


    def Product(self, data=None):
        from entity.product_entity import ProductEntity
        return ProductEntity(self, data)


    def Publication(self, data=None):
        from entity.publication_entity import PublicationEntity
        return PublicationEntity(self, data)


    def Search(self, data=None):
        from entity.search_entity import SearchEntity
        return SearchEntity(self, data)


    def Section(self, data=None):
        from entity.section_entity import SectionEntity
        return SectionEntity(self, data)


    def Site(self, data=None):
        from entity.site_entity import SiteEntity
        return SiteEntity(self, data)


    def Sound(self, data=None):
        from entity.sound_entity import SoundEntity
        return SoundEntity(self, data)


    def StaticPage(self, data=None):
        from entity.static_page_entity import StaticPageEntity
        return StaticPageEntity(self, data)


    def Text(self, data=None):
        from entity.text_entity import TextEntity
        return TextEntity(self, data)


    def Tour(self, data=None):
        from entity.tour_entity import TourEntity
        return TourEntity(self, data)


    def Video(self, data=None):
        from entity.video_entity import VideoEntity
        return VideoEntity(self, data)



    @classmethod
    def test(cls, testopts=None, sdkopts=None):
        if sdkopts is None:
            sdkopts = {}
        sdkopts = vs.clone(sdkopts)
        if not isinstance(sdkopts, dict):
            sdkopts = {}

        if testopts is None:
            testopts = {}
        testopts = vs.clone(testopts)
        if not isinstance(testopts, dict):
            testopts = {}
        testopts["active"] = True

        vs.setpath(sdkopts, "feature.test", testopts)

        sdk = cls(sdkopts)
        sdk.mode = "test"

        return sdk
