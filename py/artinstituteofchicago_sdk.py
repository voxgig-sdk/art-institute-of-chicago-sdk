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


    @property
    def agent(self):
        """Idiomatic facade: client.agent.list() / client.agent.load({"id": ...})."""
        from entity.agent_entity import AgentEntity
        cached = getattr(self, "_agent", None)
        if cached is None:
            cached = AgentEntity(self, None)
            self._agent = cached
        return cached

    def Agent(self, data=None):
        # Deprecated: use client.agent instead.
        from entity.agent_entity import AgentEntity
        return AgentEntity(self, data)


    @property
    def agent_role(self):
        """Idiomatic facade: client.agent_role.list() / client.agent_role.load({"id": ...})."""
        from entity.agent_role_entity import AgentRoleEntity
        cached = getattr(self, "_agent_role", None)
        if cached is None:
            cached = AgentRoleEntity(self, None)
            self._agent_role = cached
        return cached

    def AgentRole(self, data=None):
        # Deprecated: use client.agent_role instead.
        from entity.agent_role_entity import AgentRoleEntity
        return AgentRoleEntity(self, data)


    @property
    def agent_type(self):
        """Idiomatic facade: client.agent_type.list() / client.agent_type.load({"id": ...})."""
        from entity.agent_type_entity import AgentTypeEntity
        cached = getattr(self, "_agent_type", None)
        if cached is None:
            cached = AgentTypeEntity(self, None)
            self._agent_type = cached
        return cached

    def AgentType(self, data=None):
        # Deprecated: use client.agent_type instead.
        from entity.agent_type_entity import AgentTypeEntity
        return AgentTypeEntity(self, data)


    @property
    def article(self):
        """Idiomatic facade: client.article.list() / client.article.load({"id": ...})."""
        from entity.article_entity import ArticleEntity
        cached = getattr(self, "_article", None)
        if cached is None:
            cached = ArticleEntity(self, None)
            self._article = cached
        return cached

    def Article(self, data=None):
        # Deprecated: use client.article instead.
        from entity.article_entity import ArticleEntity
        return ArticleEntity(self, data)


    @property
    def artwork(self):
        """Idiomatic facade: client.artwork.list() / client.artwork.load({"id": ...})."""
        from entity.artwork_entity import ArtworkEntity
        cached = getattr(self, "_artwork", None)
        if cached is None:
            cached = ArtworkEntity(self, None)
            self._artwork = cached
        return cached

    def Artwork(self, data=None):
        # Deprecated: use client.artwork instead.
        from entity.artwork_entity import ArtworkEntity
        return ArtworkEntity(self, data)


    @property
    def artwork_date_qualifier(self):
        """Idiomatic facade: client.artwork_date_qualifier.list() / client.artwork_date_qualifier.load({"id": ...})."""
        from entity.artwork_date_qualifier_entity import ArtworkDateQualifierEntity
        cached = getattr(self, "_artwork_date_qualifier", None)
        if cached is None:
            cached = ArtworkDateQualifierEntity(self, None)
            self._artwork_date_qualifier = cached
        return cached

    def ArtworkDateQualifier(self, data=None):
        # Deprecated: use client.artwork_date_qualifier instead.
        from entity.artwork_date_qualifier_entity import ArtworkDateQualifierEntity
        return ArtworkDateQualifierEntity(self, data)


    @property
    def artwork_place_qualifier(self):
        """Idiomatic facade: client.artwork_place_qualifier.list() / client.artwork_place_qualifier.load({"id": ...})."""
        from entity.artwork_place_qualifier_entity import ArtworkPlaceQualifierEntity
        cached = getattr(self, "_artwork_place_qualifier", None)
        if cached is None:
            cached = ArtworkPlaceQualifierEntity(self, None)
            self._artwork_place_qualifier = cached
        return cached

    def ArtworkPlaceQualifier(self, data=None):
        # Deprecated: use client.artwork_place_qualifier instead.
        from entity.artwork_place_qualifier_entity import ArtworkPlaceQualifierEntity
        return ArtworkPlaceQualifierEntity(self, data)


    @property
    def artwork_type(self):
        """Idiomatic facade: client.artwork_type.list() / client.artwork_type.load({"id": ...})."""
        from entity.artwork_type_entity import ArtworkTypeEntity
        cached = getattr(self, "_artwork_type", None)
        if cached is None:
            cached = ArtworkTypeEntity(self, None)
            self._artwork_type = cached
        return cached

    def ArtworkType(self, data=None):
        # Deprecated: use client.artwork_type instead.
        from entity.artwork_type_entity import ArtworkTypeEntity
        return ArtworkTypeEntity(self, data)


    @property
    def category_term(self):
        """Idiomatic facade: client.category_term.list() / client.category_term.load({"id": ...})."""
        from entity.category_term_entity import CategoryTermEntity
        cached = getattr(self, "_category_term", None)
        if cached is None:
            cached = CategoryTermEntity(self, None)
            self._category_term = cached
        return cached

    def CategoryTerm(self, data=None):
        # Deprecated: use client.category_term instead.
        from entity.category_term_entity import CategoryTermEntity
        return CategoryTermEntity(self, data)


    @property
    def digital_publication(self):
        """Idiomatic facade: client.digital_publication.list() / client.digital_publication.load({"id": ...})."""
        from entity.digital_publication_entity import DigitalPublicationEntity
        cached = getattr(self, "_digital_publication", None)
        if cached is None:
            cached = DigitalPublicationEntity(self, None)
            self._digital_publication = cached
        return cached

    def DigitalPublication(self, data=None):
        # Deprecated: use client.digital_publication instead.
        from entity.digital_publication_entity import DigitalPublicationEntity
        return DigitalPublicationEntity(self, data)


    @property
    def digital_publication_article(self):
        """Idiomatic facade: client.digital_publication_article.list() / client.digital_publication_article.load({"id": ...})."""
        from entity.digital_publication_article_entity import DigitalPublicationArticleEntity
        cached = getattr(self, "_digital_publication_article", None)
        if cached is None:
            cached = DigitalPublicationArticleEntity(self, None)
            self._digital_publication_article = cached
        return cached

    def DigitalPublicationArticle(self, data=None):
        # Deprecated: use client.digital_publication_article instead.
        from entity.digital_publication_article_entity import DigitalPublicationArticleEntity
        return DigitalPublicationArticleEntity(self, data)


    @property
    def educator_resource(self):
        """Idiomatic facade: client.educator_resource.list() / client.educator_resource.load({"id": ...})."""
        from entity.educator_resource_entity import EducatorResourceEntity
        cached = getattr(self, "_educator_resource", None)
        if cached is None:
            cached = EducatorResourceEntity(self, None)
            self._educator_resource = cached
        return cached

    def EducatorResource(self, data=None):
        # Deprecated: use client.educator_resource instead.
        from entity.educator_resource_entity import EducatorResourceEntity
        return EducatorResourceEntity(self, data)


    @property
    def event(self):
        """Idiomatic facade: client.event.list() / client.event.load({"id": ...})."""
        from entity.event_entity import EventEntity
        cached = getattr(self, "_event", None)
        if cached is None:
            cached = EventEntity(self, None)
            self._event = cached
        return cached

    def Event(self, data=None):
        # Deprecated: use client.event instead.
        from entity.event_entity import EventEntity
        return EventEntity(self, data)


    @property
    def event_occurrence(self):
        """Idiomatic facade: client.event_occurrence.list() / client.event_occurrence.load({"id": ...})."""
        from entity.event_occurrence_entity import EventOccurrenceEntity
        cached = getattr(self, "_event_occurrence", None)
        if cached is None:
            cached = EventOccurrenceEntity(self, None)
            self._event_occurrence = cached
        return cached

    def EventOccurrence(self, data=None):
        # Deprecated: use client.event_occurrence instead.
        from entity.event_occurrence_entity import EventOccurrenceEntity
        return EventOccurrenceEntity(self, data)


    @property
    def event_program(self):
        """Idiomatic facade: client.event_program.list() / client.event_program.load({"id": ...})."""
        from entity.event_program_entity import EventProgramEntity
        cached = getattr(self, "_event_program", None)
        if cached is None:
            cached = EventProgramEntity(self, None)
            self._event_program = cached
        return cached

    def EventProgram(self, data=None):
        # Deprecated: use client.event_program instead.
        from entity.event_program_entity import EventProgramEntity
        return EventProgramEntity(self, data)


    @property
    def exhibition(self):
        """Idiomatic facade: client.exhibition.list() / client.exhibition.load({"id": ...})."""
        from entity.exhibition_entity import ExhibitionEntity
        cached = getattr(self, "_exhibition", None)
        if cached is None:
            cached = ExhibitionEntity(self, None)
            self._exhibition = cached
        return cached

    def Exhibition(self, data=None):
        # Deprecated: use client.exhibition instead.
        from entity.exhibition_entity import ExhibitionEntity
        return ExhibitionEntity(self, data)


    @property
    def gallery(self):
        """Idiomatic facade: client.gallery.list() / client.gallery.load({"id": ...})."""
        from entity.gallery_entity import GalleryEntity
        cached = getattr(self, "_gallery", None)
        if cached is None:
            cached = GalleryEntity(self, None)
            self._gallery = cached
        return cached

    def Gallery(self, data=None):
        # Deprecated: use client.gallery instead.
        from entity.gallery_entity import GalleryEntity
        return GalleryEntity(self, data)


    @property
    def generic_page(self):
        """Idiomatic facade: client.generic_page.list() / client.generic_page.load({"id": ...})."""
        from entity.generic_page_entity import GenericPageEntity
        cached = getattr(self, "_generic_page", None)
        if cached is None:
            cached = GenericPageEntity(self, None)
            self._generic_page = cached
        return cached

    def GenericPage(self, data=None):
        # Deprecated: use client.generic_page instead.
        from entity.generic_page_entity import GenericPageEntity
        return GenericPageEntity(self, data)


    @property
    def highlight(self):
        """Idiomatic facade: client.highlight.list() / client.highlight.load({"id": ...})."""
        from entity.highlight_entity import HighlightEntity
        cached = getattr(self, "_highlight", None)
        if cached is None:
            cached = HighlightEntity(self, None)
            self._highlight = cached
        return cached

    def Highlight(self, data=None):
        # Deprecated: use client.highlight instead.
        from entity.highlight_entity import HighlightEntity
        return HighlightEntity(self, data)


    @property
    def hour(self):
        """Idiomatic facade: client.hour.list() / client.hour.load({"id": ...})."""
        from entity.hour_entity import HourEntity
        cached = getattr(self, "_hour", None)
        if cached is None:
            cached = HourEntity(self, None)
            self._hour = cached
        return cached

    def Hour(self, data=None):
        # Deprecated: use client.hour instead.
        from entity.hour_entity import HourEntity
        return HourEntity(self, data)


    @property
    def image(self):
        """Idiomatic facade: client.image.list() / client.image.load({"id": ...})."""
        from entity.image_entity import ImageEntity
        cached = getattr(self, "_image", None)
        if cached is None:
            cached = ImageEntity(self, None)
            self._image = cached
        return cached

    def Image(self, data=None):
        # Deprecated: use client.image instead.
        from entity.image_entity import ImageEntity
        return ImageEntity(self, data)


    @property
    def landing_page(self):
        """Idiomatic facade: client.landing_page.list() / client.landing_page.load({"id": ...})."""
        from entity.landing_page_entity import LandingPageEntity
        cached = getattr(self, "_landing_page", None)
        if cached is None:
            cached = LandingPageEntity(self, None)
            self._landing_page = cached
        return cached

    def LandingPage(self, data=None):
        # Deprecated: use client.landing_page instead.
        from entity.landing_page_entity import LandingPageEntity
        return LandingPageEntity(self, data)


    @property
    def place(self):
        """Idiomatic facade: client.place.list() / client.place.load({"id": ...})."""
        from entity.place_entity import PlaceEntity
        cached = getattr(self, "_place", None)
        if cached is None:
            cached = PlaceEntity(self, None)
            self._place = cached
        return cached

    def Place(self, data=None):
        # Deprecated: use client.place instead.
        from entity.place_entity import PlaceEntity
        return PlaceEntity(self, data)


    @property
    def press_release(self):
        """Idiomatic facade: client.press_release.list() / client.press_release.load({"id": ...})."""
        from entity.press_release_entity import PressReleaseEntity
        cached = getattr(self, "_press_release", None)
        if cached is None:
            cached = PressReleaseEntity(self, None)
            self._press_release = cached
        return cached

    def PressRelease(self, data=None):
        # Deprecated: use client.press_release instead.
        from entity.press_release_entity import PressReleaseEntity
        return PressReleaseEntity(self, data)


    @property
    def printed_publication(self):
        """Idiomatic facade: client.printed_publication.list() / client.printed_publication.load({"id": ...})."""
        from entity.printed_publication_entity import PrintedPublicationEntity
        cached = getattr(self, "_printed_publication", None)
        if cached is None:
            cached = PrintedPublicationEntity(self, None)
            self._printed_publication = cached
        return cached

    def PrintedPublication(self, data=None):
        # Deprecated: use client.printed_publication instead.
        from entity.printed_publication_entity import PrintedPublicationEntity
        return PrintedPublicationEntity(self, data)


    @property
    def product(self):
        """Idiomatic facade: client.product.list() / client.product.load({"id": ...})."""
        from entity.product_entity import ProductEntity
        cached = getattr(self, "_product", None)
        if cached is None:
            cached = ProductEntity(self, None)
            self._product = cached
        return cached

    def Product(self, data=None):
        # Deprecated: use client.product instead.
        from entity.product_entity import ProductEntity
        return ProductEntity(self, data)


    @property
    def publication(self):
        """Idiomatic facade: client.publication.list() / client.publication.load({"id": ...})."""
        from entity.publication_entity import PublicationEntity
        cached = getattr(self, "_publication", None)
        if cached is None:
            cached = PublicationEntity(self, None)
            self._publication = cached
        return cached

    def Publication(self, data=None):
        # Deprecated: use client.publication instead.
        from entity.publication_entity import PublicationEntity
        return PublicationEntity(self, data)


    @property
    def search(self):
        """Idiomatic facade: client.search.list() / client.search.load({"id": ...})."""
        from entity.search_entity import SearchEntity
        cached = getattr(self, "_search", None)
        if cached is None:
            cached = SearchEntity(self, None)
            self._search = cached
        return cached

    def Search(self, data=None):
        # Deprecated: use client.search instead.
        from entity.search_entity import SearchEntity
        return SearchEntity(self, data)


    @property
    def section(self):
        """Idiomatic facade: client.section.list() / client.section.load({"id": ...})."""
        from entity.section_entity import SectionEntity
        cached = getattr(self, "_section", None)
        if cached is None:
            cached = SectionEntity(self, None)
            self._section = cached
        return cached

    def Section(self, data=None):
        # Deprecated: use client.section instead.
        from entity.section_entity import SectionEntity
        return SectionEntity(self, data)


    @property
    def site(self):
        """Idiomatic facade: client.site.list() / client.site.load({"id": ...})."""
        from entity.site_entity import SiteEntity
        cached = getattr(self, "_site", None)
        if cached is None:
            cached = SiteEntity(self, None)
            self._site = cached
        return cached

    def Site(self, data=None):
        # Deprecated: use client.site instead.
        from entity.site_entity import SiteEntity
        return SiteEntity(self, data)


    @property
    def sound(self):
        """Idiomatic facade: client.sound.list() / client.sound.load({"id": ...})."""
        from entity.sound_entity import SoundEntity
        cached = getattr(self, "_sound", None)
        if cached is None:
            cached = SoundEntity(self, None)
            self._sound = cached
        return cached

    def Sound(self, data=None):
        # Deprecated: use client.sound instead.
        from entity.sound_entity import SoundEntity
        return SoundEntity(self, data)


    @property
    def static_page(self):
        """Idiomatic facade: client.static_page.list() / client.static_page.load({"id": ...})."""
        from entity.static_page_entity import StaticPageEntity
        cached = getattr(self, "_static_page", None)
        if cached is None:
            cached = StaticPageEntity(self, None)
            self._static_page = cached
        return cached

    def StaticPage(self, data=None):
        # Deprecated: use client.static_page instead.
        from entity.static_page_entity import StaticPageEntity
        return StaticPageEntity(self, data)


    @property
    def text(self):
        """Idiomatic facade: client.text.list() / client.text.load({"id": ...})."""
        from entity.text_entity import TextEntity
        cached = getattr(self, "_text", None)
        if cached is None:
            cached = TextEntity(self, None)
            self._text = cached
        return cached

    def Text(self, data=None):
        # Deprecated: use client.text instead.
        from entity.text_entity import TextEntity
        return TextEntity(self, data)


    @property
    def tour(self):
        """Idiomatic facade: client.tour.list() / client.tour.load({"id": ...})."""
        from entity.tour_entity import TourEntity
        cached = getattr(self, "_tour", None)
        if cached is None:
            cached = TourEntity(self, None)
            self._tour = cached
        return cached

    def Tour(self, data=None):
        # Deprecated: use client.tour instead.
        from entity.tour_entity import TourEntity
        return TourEntity(self, data)


    @property
    def video(self):
        """Idiomatic facade: client.video.list() / client.video.load({"id": ...})."""
        from entity.video_entity import VideoEntity
        cached = getattr(self, "_video", None)
        if cached is None:
            cached = VideoEntity(self, None)
            self._video = cached
        return cached

    def Video(self, data=None):
        # Deprecated: use client.video instead.
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
