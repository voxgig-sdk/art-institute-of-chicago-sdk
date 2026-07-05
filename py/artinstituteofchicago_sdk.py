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
            raise err

        fetchdef, err = utility.make_fetch_def(ctx)
        if err is not None:
            raise err

        return fetchdef

    def direct(self, fetchargs=None):
        utility = self._utility

        try:
            fetchdef = self.prepare(fetchargs)
        except Exception as err:
            # direct() is the raw-HTTP escape hatch: it never raises, it
            # returns a result object callers branch on via result["ok"].
            return {"ok": False, "err": err}

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
            return {"ok": False, "err": fetch_err}

        if fetched is None:
            return {
                "ok": False,
                "err": ctx.make_error("direct_no_response", "response: undefined"),
            }

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
            }

        return {
            "ok": False,
            "err": ctx.make_error("direct_invalid", "invalid response type"),
        }


    def Agent(self, data=None) -> "AgentEntity":
        """Entity factory: client.Agent().list() / client.Agent().load({"id": ...})."""
        from entity.agent_entity import AgentEntity
        return AgentEntity(self, data)


    def AgentRole(self, data=None) -> "AgentRoleEntity":
        """Entity factory: client.AgentRole().list() / client.AgentRole().load({"id": ...})."""
        from entity.agent_role_entity import AgentRoleEntity
        return AgentRoleEntity(self, data)


    def AgentType(self, data=None) -> "AgentTypeEntity":
        """Entity factory: client.AgentType().list() / client.AgentType().load({"id": ...})."""
        from entity.agent_type_entity import AgentTypeEntity
        return AgentTypeEntity(self, data)


    def Article(self, data=None) -> "ArticleEntity":
        """Entity factory: client.Article().list() / client.Article().load({"id": ...})."""
        from entity.article_entity import ArticleEntity
        return ArticleEntity(self, data)


    def Artwork(self, data=None) -> "ArtworkEntity":
        """Entity factory: client.Artwork().list() / client.Artwork().load({"id": ...})."""
        from entity.artwork_entity import ArtworkEntity
        return ArtworkEntity(self, data)


    def ArtworkDateQualifier(self, data=None) -> "ArtworkDateQualifierEntity":
        """Entity factory: client.ArtworkDateQualifier().list() / client.ArtworkDateQualifier().load({"id": ...})."""
        from entity.artwork_date_qualifier_entity import ArtworkDateQualifierEntity
        return ArtworkDateQualifierEntity(self, data)


    def ArtworkPlaceQualifier(self, data=None) -> "ArtworkPlaceQualifierEntity":
        """Entity factory: client.ArtworkPlaceQualifier().list() / client.ArtworkPlaceQualifier().load({"id": ...})."""
        from entity.artwork_place_qualifier_entity import ArtworkPlaceQualifierEntity
        return ArtworkPlaceQualifierEntity(self, data)


    def ArtworkType(self, data=None) -> "ArtworkTypeEntity":
        """Entity factory: client.ArtworkType().list() / client.ArtworkType().load({"id": ...})."""
        from entity.artwork_type_entity import ArtworkTypeEntity
        return ArtworkTypeEntity(self, data)


    def CategoryTerm(self, data=None) -> "CategoryTermEntity":
        """Entity factory: client.CategoryTerm().list() / client.CategoryTerm().load({"id": ...})."""
        from entity.category_term_entity import CategoryTermEntity
        return CategoryTermEntity(self, data)


    def DigitalPublication(self, data=None) -> "DigitalPublicationEntity":
        """Entity factory: client.DigitalPublication().list() / client.DigitalPublication().load({"id": ...})."""
        from entity.digital_publication_entity import DigitalPublicationEntity
        return DigitalPublicationEntity(self, data)


    def DigitalPublicationArticle(self, data=None) -> "DigitalPublicationArticleEntity":
        """Entity factory: client.DigitalPublicationArticle().list() / client.DigitalPublicationArticle().load({"id": ...})."""
        from entity.digital_publication_article_entity import DigitalPublicationArticleEntity
        return DigitalPublicationArticleEntity(self, data)


    def EducatorResource(self, data=None) -> "EducatorResourceEntity":
        """Entity factory: client.EducatorResource().list() / client.EducatorResource().load({"id": ...})."""
        from entity.educator_resource_entity import EducatorResourceEntity
        return EducatorResourceEntity(self, data)


    def Event(self, data=None) -> "EventEntity":
        """Entity factory: client.Event().list() / client.Event().load({"id": ...})."""
        from entity.event_entity import EventEntity
        return EventEntity(self, data)


    def EventOccurrence(self, data=None) -> "EventOccurrenceEntity":
        """Entity factory: client.EventOccurrence().list() / client.EventOccurrence().load({"id": ...})."""
        from entity.event_occurrence_entity import EventOccurrenceEntity
        return EventOccurrenceEntity(self, data)


    def EventProgram(self, data=None) -> "EventProgramEntity":
        """Entity factory: client.EventProgram().list() / client.EventProgram().load({"id": ...})."""
        from entity.event_program_entity import EventProgramEntity
        return EventProgramEntity(self, data)


    def Exhibition(self, data=None) -> "ExhibitionEntity":
        """Entity factory: client.Exhibition().list() / client.Exhibition().load({"id": ...})."""
        from entity.exhibition_entity import ExhibitionEntity
        return ExhibitionEntity(self, data)


    def Gallery(self, data=None) -> "GalleryEntity":
        """Entity factory: client.Gallery().list() / client.Gallery().load({"id": ...})."""
        from entity.gallery_entity import GalleryEntity
        return GalleryEntity(self, data)


    def GenericPage(self, data=None) -> "GenericPageEntity":
        """Entity factory: client.GenericPage().list() / client.GenericPage().load({"id": ...})."""
        from entity.generic_page_entity import GenericPageEntity
        return GenericPageEntity(self, data)


    def Highlight(self, data=None) -> "HighlightEntity":
        """Entity factory: client.Highlight().list() / client.Highlight().load({"id": ...})."""
        from entity.highlight_entity import HighlightEntity
        return HighlightEntity(self, data)


    def Hour(self, data=None) -> "HourEntity":
        """Entity factory: client.Hour().list() / client.Hour().load({"id": ...})."""
        from entity.hour_entity import HourEntity
        return HourEntity(self, data)


    def Image(self, data=None) -> "ImageEntity":
        """Entity factory: client.Image().list() / client.Image().load({"id": ...})."""
        from entity.image_entity import ImageEntity
        return ImageEntity(self, data)


    def LandingPage(self, data=None) -> "LandingPageEntity":
        """Entity factory: client.LandingPage().list() / client.LandingPage().load({"id": ...})."""
        from entity.landing_page_entity import LandingPageEntity
        return LandingPageEntity(self, data)


    def Place(self, data=None) -> "PlaceEntity":
        """Entity factory: client.Place().list() / client.Place().load({"id": ...})."""
        from entity.place_entity import PlaceEntity
        return PlaceEntity(self, data)


    def PressRelease(self, data=None) -> "PressReleaseEntity":
        """Entity factory: client.PressRelease().list() / client.PressRelease().load({"id": ...})."""
        from entity.press_release_entity import PressReleaseEntity
        return PressReleaseEntity(self, data)


    def PrintedPublication(self, data=None) -> "PrintedPublicationEntity":
        """Entity factory: client.PrintedPublication().list() / client.PrintedPublication().load({"id": ...})."""
        from entity.printed_publication_entity import PrintedPublicationEntity
        return PrintedPublicationEntity(self, data)


    def Product(self, data=None) -> "ProductEntity":
        """Entity factory: client.Product().list() / client.Product().load({"id": ...})."""
        from entity.product_entity import ProductEntity
        return ProductEntity(self, data)


    def Publication(self, data=None) -> "PublicationEntity":
        """Entity factory: client.Publication().list() / client.Publication().load({"id": ...})."""
        from entity.publication_entity import PublicationEntity
        return PublicationEntity(self, data)


    def Search(self, data=None) -> "SearchEntity":
        """Entity factory: client.Search().list() / client.Search().load({"id": ...})."""
        from entity.search_entity import SearchEntity
        return SearchEntity(self, data)


    def Section(self, data=None) -> "SectionEntity":
        """Entity factory: client.Section().list() / client.Section().load({"id": ...})."""
        from entity.section_entity import SectionEntity
        return SectionEntity(self, data)


    def Site(self, data=None) -> "SiteEntity":
        """Entity factory: client.Site().list() / client.Site().load({"id": ...})."""
        from entity.site_entity import SiteEntity
        return SiteEntity(self, data)


    def Sound(self, data=None) -> "SoundEntity":
        """Entity factory: client.Sound().list() / client.Sound().load({"id": ...})."""
        from entity.sound_entity import SoundEntity
        return SoundEntity(self, data)


    def StaticPage(self, data=None) -> "StaticPageEntity":
        """Entity factory: client.StaticPage().list() / client.StaticPage().load({"id": ...})."""
        from entity.static_page_entity import StaticPageEntity
        return StaticPageEntity(self, data)


    def Text(self, data=None) -> "TextEntity":
        """Entity factory: client.Text().list() / client.Text().load({"id": ...})."""
        from entity.text_entity import TextEntity
        return TextEntity(self, data)


    def Tour(self, data=None) -> "TourEntity":
        """Entity factory: client.Tour().list() / client.Tour().load({"id": ...})."""
        from entity.tour_entity import TourEntity
        return TourEntity(self, data)


    def Video(self, data=None) -> "VideoEntity":
        """Entity factory: client.Video().list() / client.Video().load({"id": ...})."""
        from entity.video_entity import VideoEntity
        return VideoEntity(self, data)



    @classmethod
    def test(cls, testopts=None, sdkopts=None) -> "ArtInstituteOfChicagoSDK":
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


from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from entity.agent_entity import AgentEntity
    from entity.agent_role_entity import AgentRoleEntity
    from entity.agent_type_entity import AgentTypeEntity
    from entity.article_entity import ArticleEntity
    from entity.artwork_entity import ArtworkEntity
    from entity.artwork_date_qualifier_entity import ArtworkDateQualifierEntity
    from entity.artwork_place_qualifier_entity import ArtworkPlaceQualifierEntity
    from entity.artwork_type_entity import ArtworkTypeEntity
    from entity.category_term_entity import CategoryTermEntity
    from entity.digital_publication_entity import DigitalPublicationEntity
    from entity.digital_publication_article_entity import DigitalPublicationArticleEntity
    from entity.educator_resource_entity import EducatorResourceEntity
    from entity.event_entity import EventEntity
    from entity.event_occurrence_entity import EventOccurrenceEntity
    from entity.event_program_entity import EventProgramEntity
    from entity.exhibition_entity import ExhibitionEntity
    from entity.gallery_entity import GalleryEntity
    from entity.generic_page_entity import GenericPageEntity
    from entity.highlight_entity import HighlightEntity
    from entity.hour_entity import HourEntity
    from entity.image_entity import ImageEntity
    from entity.landing_page_entity import LandingPageEntity
    from entity.place_entity import PlaceEntity
    from entity.press_release_entity import PressReleaseEntity
    from entity.printed_publication_entity import PrintedPublicationEntity
    from entity.product_entity import ProductEntity
    from entity.publication_entity import PublicationEntity
    from entity.search_entity import SearchEntity
    from entity.section_entity import SectionEntity
    from entity.site_entity import SiteEntity
    from entity.sound_entity import SoundEntity
    from entity.static_page_entity import StaticPageEntity
    from entity.text_entity import TextEntity
    from entity.tour_entity import TourEntity
    from entity.video_entity import VideoEntity
