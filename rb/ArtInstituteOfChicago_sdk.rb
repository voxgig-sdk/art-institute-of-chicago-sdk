# ArtInstituteOfChicago SDK

require_relative 'utility/struct/voxgig_struct'
require_relative 'core/utility_type'
require_relative 'core/spec'
require_relative 'core/helpers'

# Load utility registration
require_relative 'utility/register'

# Load config and features
require_relative 'config'
require_relative 'feature/base_feature'
require_relative 'features'

# Load typed models (Struct value objects).
require_relative 'ArtInstituteOfChicago_types'


class ArtInstituteOfChicagoSDK
  attr_accessor :mode, :features, :options

  def initialize(options = {})
    @mode = "live"
    @features = []
    @options = nil

    utility = ArtInstituteOfChicagoUtility.new
    @_utility = utility

    config = ArtInstituteOfChicagoConfig.make_config

    @_rootctx = utility.make_context.call({
      "client" => self,
      "utility" => utility,
      "config" => config,
      "options" => options || {},
      "shared" => {},
    }, nil)

    @options = utility.make_options.call(@_rootctx)

    if VoxgigStruct.getpath(@options, "feature.test.active") == true
      @mode = "test"
    end

    @_rootctx.options = @options

    # Add features from config.
    feature_opts = ArtInstituteOfChicagoHelpers.to_map(VoxgigStruct.getprop(@options, "feature"))
    if feature_opts
      items = VoxgigStruct.items(feature_opts)
      if items
        items.each do |item|
          fname = item[0]
          fopts = ArtInstituteOfChicagoHelpers.to_map(item[1])
          if fopts && fopts["active"] == true
            utility.feature_add.call(@_rootctx, ArtInstituteOfChicagoFeatures.make_feature(fname))
          end
        end
      end
    end

    # Add extension features.
    extend_val = VoxgigStruct.getprop(@options, "extend")
    if extend_val.is_a?(Array)
      extend_val.each do |f|
        if f.respond_to?(:get_name)
          utility.feature_add.call(@_rootctx, f)
        end
      end
    end

    # Initialize features.
    @features.each do |f|
      utility.feature_init.call(@_rootctx, f)
    end

    utility.feature_hook.call(@_rootctx, "PostConstruct")
  end

  def options_map
    out = VoxgigStruct.clone(@options)
    out.is_a?(Hash) ? out : {}
  end

  def get_utility
    ArtInstituteOfChicagoUtility.copy(@_utility)
  end

  def get_root_ctx
    @_rootctx
  end

  def prepare(fetchargs = {})
    utility = @_utility
    fetchargs ||= {}

    ctrl = ArtInstituteOfChicagoHelpers.to_map(VoxgigStruct.getprop(fetchargs, "ctrl")) || {}

    ctx = utility.make_context.call({
      "opname" => "prepare",
      "ctrl" => ctrl,
    }, @_rootctx)

    opts = @options
    path = VoxgigStruct.getprop(fetchargs, "path") || ""
    path = "" unless path.is_a?(String)
    method_val = VoxgigStruct.getprop(fetchargs, "method") || "GET"
    method_val = "GET" unless method_val.is_a?(String)
    params = ArtInstituteOfChicagoHelpers.to_map(VoxgigStruct.getprop(fetchargs, "params")) || {}
    query = ArtInstituteOfChicagoHelpers.to_map(VoxgigStruct.getprop(fetchargs, "query")) || {}
    headers = utility.prepare_headers.call(ctx)

    base = VoxgigStruct.getprop(opts, "base") || ""
    base = "" unless base.is_a?(String)
    prefix = VoxgigStruct.getprop(opts, "prefix") || ""
    prefix = "" unless prefix.is_a?(String)
    suffix = VoxgigStruct.getprop(opts, "suffix") || ""
    suffix = "" unless suffix.is_a?(String)

    ctx.spec = ArtInstituteOfChicagoSpec.new({
      "base" => base, "prefix" => prefix, "suffix" => suffix,
      "path" => path, "method" => method_val,
      "params" => params, "query" => query, "headers" => headers,
      "body" => VoxgigStruct.getprop(fetchargs, "body"),
      "step" => "start",
    })

    # Merge user-provided headers.
    uh = VoxgigStruct.getprop(fetchargs, "headers")
    if uh.is_a?(Hash)
      uh.each { |k, v| ctx.spec.headers[k] = v }
    end

    _, err = utility.prepare_auth.call(ctx)
    raise err if err

    utility.make_fetch_def.call(ctx)
  end

  def direct(fetchargs = {})
    utility = @_utility

    # direct() is the raw-HTTP escape hatch: it always returns a result hash
    # ({ "ok" => ..., ... }) and never raises. prepare() raises on error, so
    # trap that and surface it in the hash.
    begin
      fetchdef = prepare(fetchargs)
    rescue ArtInstituteOfChicagoError => err
      return { "ok" => false, "err" => err }
    end

    fetchargs ||= {}
    ctrl = ArtInstituteOfChicagoHelpers.to_map(VoxgigStruct.getprop(fetchargs, "ctrl")) || {}

    ctx = utility.make_context.call({
      "opname" => "direct",
      "ctrl" => ctrl,
    }, @_rootctx)

    url = fetchdef["url"] || ""
    fetched, fetch_err = utility.fetcher.call(ctx, url, fetchdef)

    return { "ok" => false, "err" => fetch_err } if fetch_err

    if fetched.nil?
      return {
        "ok" => false,
        "err" => ctx.make_error("direct_no_response", "response: undefined"),
      }
    end

    if fetched.is_a?(Hash)
      status = ArtInstituteOfChicagoHelpers.to_int(VoxgigStruct.getprop(fetched, "status"))
      headers = VoxgigStruct.getprop(fetched, "headers") || {}

      # No-body responses (204, 304) and explicit zero content-length must
      # skip JSON parsing — calling json() on an empty body errors.
      content_length = headers.is_a?(Hash) ? headers["content-length"] : nil
      no_body = status == 204 || status == 304 || content_length.to_s == "0"

      json_data = nil
      unless no_body
        jf = VoxgigStruct.getprop(fetched, "json")
        if jf.is_a?(Proc)
          begin
            json_data = jf.call
          rescue StandardError
            # Non-JSON body — leave data nil, keep status/headers.
            json_data = nil
          end
        end
      end

      return {
        "ok" => status >= 200 && status < 300,
        "status" => status,
        "headers" => headers,
        "data" => json_data,
      }
    end

    return {
      "ok" => false,
      "err" => ctx.make_error("direct_invalid", "invalid response type"),
    }
  end


  # Idiomatic facade: client.agent.list / client.agent.load({ "id" => ... })
  def agent
    require_relative 'entity/agent_entity'
    @agent ||= AgentEntity.new(self, nil)
  end

  # Deprecated: use client.agent instead.
  def Agent(data = nil)
    require_relative 'entity/agent_entity'
    AgentEntity.new(self, data)
  end


  # Idiomatic facade: client.agent_role.list / client.agent_role.load({ "id" => ... })
  def agent_role
    require_relative 'entity/agent_role_entity'
    @agent_role ||= AgentRoleEntity.new(self, nil)
  end

  # Deprecated: use client.agent_role instead.
  def AgentRole(data = nil)
    require_relative 'entity/agent_role_entity'
    AgentRoleEntity.new(self, data)
  end


  # Idiomatic facade: client.agent_type.list / client.agent_type.load({ "id" => ... })
  def agent_type
    require_relative 'entity/agent_type_entity'
    @agent_type ||= AgentTypeEntity.new(self, nil)
  end

  # Deprecated: use client.agent_type instead.
  def AgentType(data = nil)
    require_relative 'entity/agent_type_entity'
    AgentTypeEntity.new(self, data)
  end


  # Idiomatic facade: client.article.list / client.article.load({ "id" => ... })
  def article
    require_relative 'entity/article_entity'
    @article ||= ArticleEntity.new(self, nil)
  end

  # Deprecated: use client.article instead.
  def Article(data = nil)
    require_relative 'entity/article_entity'
    ArticleEntity.new(self, data)
  end


  # Idiomatic facade: client.artwork.list / client.artwork.load({ "id" => ... })
  def artwork
    require_relative 'entity/artwork_entity'
    @artwork ||= ArtworkEntity.new(self, nil)
  end

  # Deprecated: use client.artwork instead.
  def Artwork(data = nil)
    require_relative 'entity/artwork_entity'
    ArtworkEntity.new(self, data)
  end


  # Idiomatic facade: client.artwork_date_qualifier.list / client.artwork_date_qualifier.load({ "id" => ... })
  def artwork_date_qualifier
    require_relative 'entity/artwork_date_qualifier_entity'
    @artwork_date_qualifier ||= ArtworkDateQualifierEntity.new(self, nil)
  end

  # Deprecated: use client.artwork_date_qualifier instead.
  def ArtworkDateQualifier(data = nil)
    require_relative 'entity/artwork_date_qualifier_entity'
    ArtworkDateQualifierEntity.new(self, data)
  end


  # Idiomatic facade: client.artwork_place_qualifier.list / client.artwork_place_qualifier.load({ "id" => ... })
  def artwork_place_qualifier
    require_relative 'entity/artwork_place_qualifier_entity'
    @artwork_place_qualifier ||= ArtworkPlaceQualifierEntity.new(self, nil)
  end

  # Deprecated: use client.artwork_place_qualifier instead.
  def ArtworkPlaceQualifier(data = nil)
    require_relative 'entity/artwork_place_qualifier_entity'
    ArtworkPlaceQualifierEntity.new(self, data)
  end


  # Idiomatic facade: client.artwork_type.list / client.artwork_type.load({ "id" => ... })
  def artwork_type
    require_relative 'entity/artwork_type_entity'
    @artwork_type ||= ArtworkTypeEntity.new(self, nil)
  end

  # Deprecated: use client.artwork_type instead.
  def ArtworkType(data = nil)
    require_relative 'entity/artwork_type_entity'
    ArtworkTypeEntity.new(self, data)
  end


  # Idiomatic facade: client.category_term.list / client.category_term.load({ "id" => ... })
  def category_term
    require_relative 'entity/category_term_entity'
    @category_term ||= CategoryTermEntity.new(self, nil)
  end

  # Deprecated: use client.category_term instead.
  def CategoryTerm(data = nil)
    require_relative 'entity/category_term_entity'
    CategoryTermEntity.new(self, data)
  end


  # Idiomatic facade: client.digital_publication.list / client.digital_publication.load({ "id" => ... })
  def digital_publication
    require_relative 'entity/digital_publication_entity'
    @digital_publication ||= DigitalPublicationEntity.new(self, nil)
  end

  # Deprecated: use client.digital_publication instead.
  def DigitalPublication(data = nil)
    require_relative 'entity/digital_publication_entity'
    DigitalPublicationEntity.new(self, data)
  end


  # Idiomatic facade: client.digital_publication_article.list / client.digital_publication_article.load({ "id" => ... })
  def digital_publication_article
    require_relative 'entity/digital_publication_article_entity'
    @digital_publication_article ||= DigitalPublicationArticleEntity.new(self, nil)
  end

  # Deprecated: use client.digital_publication_article instead.
  def DigitalPublicationArticle(data = nil)
    require_relative 'entity/digital_publication_article_entity'
    DigitalPublicationArticleEntity.new(self, data)
  end


  # Idiomatic facade: client.educator_resource.list / client.educator_resource.load({ "id" => ... })
  def educator_resource
    require_relative 'entity/educator_resource_entity'
    @educator_resource ||= EducatorResourceEntity.new(self, nil)
  end

  # Deprecated: use client.educator_resource instead.
  def EducatorResource(data = nil)
    require_relative 'entity/educator_resource_entity'
    EducatorResourceEntity.new(self, data)
  end


  # Idiomatic facade: client.event.list / client.event.load({ "id" => ... })
  def event
    require_relative 'entity/event_entity'
    @event ||= EventEntity.new(self, nil)
  end

  # Deprecated: use client.event instead.
  def Event(data = nil)
    require_relative 'entity/event_entity'
    EventEntity.new(self, data)
  end


  # Idiomatic facade: client.event_occurrence.list / client.event_occurrence.load({ "id" => ... })
  def event_occurrence
    require_relative 'entity/event_occurrence_entity'
    @event_occurrence ||= EventOccurrenceEntity.new(self, nil)
  end

  # Deprecated: use client.event_occurrence instead.
  def EventOccurrence(data = nil)
    require_relative 'entity/event_occurrence_entity'
    EventOccurrenceEntity.new(self, data)
  end


  # Idiomatic facade: client.event_program.list / client.event_program.load({ "id" => ... })
  def event_program
    require_relative 'entity/event_program_entity'
    @event_program ||= EventProgramEntity.new(self, nil)
  end

  # Deprecated: use client.event_program instead.
  def EventProgram(data = nil)
    require_relative 'entity/event_program_entity'
    EventProgramEntity.new(self, data)
  end


  # Idiomatic facade: client.exhibition.list / client.exhibition.load({ "id" => ... })
  def exhibition
    require_relative 'entity/exhibition_entity'
    @exhibition ||= ExhibitionEntity.new(self, nil)
  end

  # Deprecated: use client.exhibition instead.
  def Exhibition(data = nil)
    require_relative 'entity/exhibition_entity'
    ExhibitionEntity.new(self, data)
  end


  # Idiomatic facade: client.gallery.list / client.gallery.load({ "id" => ... })
  def gallery
    require_relative 'entity/gallery_entity'
    @gallery ||= GalleryEntity.new(self, nil)
  end

  # Deprecated: use client.gallery instead.
  def Gallery(data = nil)
    require_relative 'entity/gallery_entity'
    GalleryEntity.new(self, data)
  end


  # Idiomatic facade: client.generic_page.list / client.generic_page.load({ "id" => ... })
  def generic_page
    require_relative 'entity/generic_page_entity'
    @generic_page ||= GenericPageEntity.new(self, nil)
  end

  # Deprecated: use client.generic_page instead.
  def GenericPage(data = nil)
    require_relative 'entity/generic_page_entity'
    GenericPageEntity.new(self, data)
  end


  # Idiomatic facade: client.highlight.list / client.highlight.load({ "id" => ... })
  def highlight
    require_relative 'entity/highlight_entity'
    @highlight ||= HighlightEntity.new(self, nil)
  end

  # Deprecated: use client.highlight instead.
  def Highlight(data = nil)
    require_relative 'entity/highlight_entity'
    HighlightEntity.new(self, data)
  end


  # Idiomatic facade: client.hour.list / client.hour.load({ "id" => ... })
  def hour
    require_relative 'entity/hour_entity'
    @hour ||= HourEntity.new(self, nil)
  end

  # Deprecated: use client.hour instead.
  def Hour(data = nil)
    require_relative 'entity/hour_entity'
    HourEntity.new(self, data)
  end


  # Idiomatic facade: client.image.list / client.image.load({ "id" => ... })
  def image
    require_relative 'entity/image_entity'
    @image ||= ImageEntity.new(self, nil)
  end

  # Deprecated: use client.image instead.
  def Image(data = nil)
    require_relative 'entity/image_entity'
    ImageEntity.new(self, data)
  end


  # Idiomatic facade: client.landing_page.list / client.landing_page.load({ "id" => ... })
  def landing_page
    require_relative 'entity/landing_page_entity'
    @landing_page ||= LandingPageEntity.new(self, nil)
  end

  # Deprecated: use client.landing_page instead.
  def LandingPage(data = nil)
    require_relative 'entity/landing_page_entity'
    LandingPageEntity.new(self, data)
  end


  # Idiomatic facade: client.place.list / client.place.load({ "id" => ... })
  def place
    require_relative 'entity/place_entity'
    @place ||= PlaceEntity.new(self, nil)
  end

  # Deprecated: use client.place instead.
  def Place(data = nil)
    require_relative 'entity/place_entity'
    PlaceEntity.new(self, data)
  end


  # Idiomatic facade: client.press_release.list / client.press_release.load({ "id" => ... })
  def press_release
    require_relative 'entity/press_release_entity'
    @press_release ||= PressReleaseEntity.new(self, nil)
  end

  # Deprecated: use client.press_release instead.
  def PressRelease(data = nil)
    require_relative 'entity/press_release_entity'
    PressReleaseEntity.new(self, data)
  end


  # Idiomatic facade: client.printed_publication.list / client.printed_publication.load({ "id" => ... })
  def printed_publication
    require_relative 'entity/printed_publication_entity'
    @printed_publication ||= PrintedPublicationEntity.new(self, nil)
  end

  # Deprecated: use client.printed_publication instead.
  def PrintedPublication(data = nil)
    require_relative 'entity/printed_publication_entity'
    PrintedPublicationEntity.new(self, data)
  end


  # Idiomatic facade: client.product.list / client.product.load({ "id" => ... })
  def product
    require_relative 'entity/product_entity'
    @product ||= ProductEntity.new(self, nil)
  end

  # Deprecated: use client.product instead.
  def Product(data = nil)
    require_relative 'entity/product_entity'
    ProductEntity.new(self, data)
  end


  # Idiomatic facade: client.publication.list / client.publication.load({ "id" => ... })
  def publication
    require_relative 'entity/publication_entity'
    @publication ||= PublicationEntity.new(self, nil)
  end

  # Deprecated: use client.publication instead.
  def Publication(data = nil)
    require_relative 'entity/publication_entity'
    PublicationEntity.new(self, data)
  end


  # Idiomatic facade: client.search.list / client.search.load({ "id" => ... })
  def search
    require_relative 'entity/search_entity'
    @search ||= SearchEntity.new(self, nil)
  end

  # Deprecated: use client.search instead.
  def Search(data = nil)
    require_relative 'entity/search_entity'
    SearchEntity.new(self, data)
  end


  # Idiomatic facade: client.section.list / client.section.load({ "id" => ... })
  def section
    require_relative 'entity/section_entity'
    @section ||= SectionEntity.new(self, nil)
  end

  # Deprecated: use client.section instead.
  def Section(data = nil)
    require_relative 'entity/section_entity'
    SectionEntity.new(self, data)
  end


  # Idiomatic facade: client.site.list / client.site.load({ "id" => ... })
  def site
    require_relative 'entity/site_entity'
    @site ||= SiteEntity.new(self, nil)
  end

  # Deprecated: use client.site instead.
  def Site(data = nil)
    require_relative 'entity/site_entity'
    SiteEntity.new(self, data)
  end


  # Idiomatic facade: client.sound.list / client.sound.load({ "id" => ... })
  def sound
    require_relative 'entity/sound_entity'
    @sound ||= SoundEntity.new(self, nil)
  end

  # Deprecated: use client.sound instead.
  def Sound(data = nil)
    require_relative 'entity/sound_entity'
    SoundEntity.new(self, data)
  end


  # Idiomatic facade: client.static_page.list / client.static_page.load({ "id" => ... })
  def static_page
    require_relative 'entity/static_page_entity'
    @static_page ||= StaticPageEntity.new(self, nil)
  end

  # Deprecated: use client.static_page instead.
  def StaticPage(data = nil)
    require_relative 'entity/static_page_entity'
    StaticPageEntity.new(self, data)
  end


  # Idiomatic facade: client.text.list / client.text.load({ "id" => ... })
  def text
    require_relative 'entity/text_entity'
    @text ||= TextEntity.new(self, nil)
  end

  # Deprecated: use client.text instead.
  def Text(data = nil)
    require_relative 'entity/text_entity'
    TextEntity.new(self, data)
  end


  # Idiomatic facade: client.tour.list / client.tour.load({ "id" => ... })
  def tour
    require_relative 'entity/tour_entity'
    @tour ||= TourEntity.new(self, nil)
  end

  # Deprecated: use client.tour instead.
  def Tour(data = nil)
    require_relative 'entity/tour_entity'
    TourEntity.new(self, data)
  end


  # Idiomatic facade: client.video.list / client.video.load({ "id" => ... })
  def video
    require_relative 'entity/video_entity'
    @video ||= VideoEntity.new(self, nil)
  end

  # Deprecated: use client.video instead.
  def Video(data = nil)
    require_relative 'entity/video_entity'
    VideoEntity.new(self, data)
  end



  def self.test(testopts = nil, sdkopts = nil)
    sdkopts = sdkopts || {}
    sdkopts = VoxgigStruct.clone(sdkopts)
    sdkopts = {} unless sdkopts.is_a?(Hash)

    testopts = testopts || {}
    testopts = VoxgigStruct.clone(testopts)
    testopts = {} unless testopts.is_a?(Hash)
    testopts["active"] = true

    VoxgigStruct.setpath(sdkopts, "feature.test", testopts)

    sdk = ArtInstituteOfChicagoSDK.new(sdkopts)
    sdk.mode = "test"
    sdk
  end
end
