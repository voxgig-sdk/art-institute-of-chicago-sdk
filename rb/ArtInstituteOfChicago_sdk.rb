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
    return nil, err if err

    utility.make_fetch_def.call(ctx)
  end

  def direct(fetchargs = {})
    utility = @_utility

    fetchdef, err = prepare(fetchargs)
    return { "ok" => false, "err" => err }, nil if err

    fetchargs ||= {}
    ctrl = ArtInstituteOfChicagoHelpers.to_map(VoxgigStruct.getprop(fetchargs, "ctrl")) || {}

    ctx = utility.make_context.call({
      "opname" => "direct",
      "ctrl" => ctrl,
    }, @_rootctx)

    url = fetchdef["url"] || ""
    fetched, fetch_err = utility.fetcher.call(ctx, url, fetchdef)

    return { "ok" => false, "err" => fetch_err }, nil if fetch_err

    if fetched.nil?
      return {
        "ok" => false,
        "err" => ctx.make_error("direct_no_response", "response: undefined"),
      }, nil
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
      }, nil
    end

    return {
      "ok" => false,
      "err" => ctx.make_error("direct_invalid", "invalid response type"),
    }, nil
  end


  def Agent(data = nil)
    require_relative 'entity/agent_entity'
    AgentEntity.new(self, data)
  end


  def AgentRole(data = nil)
    require_relative 'entity/agent_role_entity'
    AgentRoleEntity.new(self, data)
  end


  def AgentType(data = nil)
    require_relative 'entity/agent_type_entity'
    AgentTypeEntity.new(self, data)
  end


  def Article(data = nil)
    require_relative 'entity/article_entity'
    ArticleEntity.new(self, data)
  end


  def Artwork(data = nil)
    require_relative 'entity/artwork_entity'
    ArtworkEntity.new(self, data)
  end


  def ArtworkDateQualifier(data = nil)
    require_relative 'entity/artwork_date_qualifier_entity'
    ArtworkDateQualifierEntity.new(self, data)
  end


  def ArtworkPlaceQualifier(data = nil)
    require_relative 'entity/artwork_place_qualifier_entity'
    ArtworkPlaceQualifierEntity.new(self, data)
  end


  def ArtworkType(data = nil)
    require_relative 'entity/artwork_type_entity'
    ArtworkTypeEntity.new(self, data)
  end


  def CategoryTerm(data = nil)
    require_relative 'entity/category_term_entity'
    CategoryTermEntity.new(self, data)
  end


  def DigitalPublication(data = nil)
    require_relative 'entity/digital_publication_entity'
    DigitalPublicationEntity.new(self, data)
  end


  def DigitalPublicationArticle(data = nil)
    require_relative 'entity/digital_publication_article_entity'
    DigitalPublicationArticleEntity.new(self, data)
  end


  def EducatorResource(data = nil)
    require_relative 'entity/educator_resource_entity'
    EducatorResourceEntity.new(self, data)
  end


  def Event(data = nil)
    require_relative 'entity/event_entity'
    EventEntity.new(self, data)
  end


  def EventOccurrence(data = nil)
    require_relative 'entity/event_occurrence_entity'
    EventOccurrenceEntity.new(self, data)
  end


  def EventProgram(data = nil)
    require_relative 'entity/event_program_entity'
    EventProgramEntity.new(self, data)
  end


  def Exhibition(data = nil)
    require_relative 'entity/exhibition_entity'
    ExhibitionEntity.new(self, data)
  end


  def Gallery(data = nil)
    require_relative 'entity/gallery_entity'
    GalleryEntity.new(self, data)
  end


  def GenericPage(data = nil)
    require_relative 'entity/generic_page_entity'
    GenericPageEntity.new(self, data)
  end


  def Highlight(data = nil)
    require_relative 'entity/highlight_entity'
    HighlightEntity.new(self, data)
  end


  def Hour(data = nil)
    require_relative 'entity/hour_entity'
    HourEntity.new(self, data)
  end


  def Image(data = nil)
    require_relative 'entity/image_entity'
    ImageEntity.new(self, data)
  end


  def LandingPage(data = nil)
    require_relative 'entity/landing_page_entity'
    LandingPageEntity.new(self, data)
  end


  def Place(data = nil)
    require_relative 'entity/place_entity'
    PlaceEntity.new(self, data)
  end


  def PressRelease(data = nil)
    require_relative 'entity/press_release_entity'
    PressReleaseEntity.new(self, data)
  end


  def PrintedPublication(data = nil)
    require_relative 'entity/printed_publication_entity'
    PrintedPublicationEntity.new(self, data)
  end


  def Product(data = nil)
    require_relative 'entity/product_entity'
    ProductEntity.new(self, data)
  end


  def Publication(data = nil)
    require_relative 'entity/publication_entity'
    PublicationEntity.new(self, data)
  end


  def Search(data = nil)
    require_relative 'entity/search_entity'
    SearchEntity.new(self, data)
  end


  def Section(data = nil)
    require_relative 'entity/section_entity'
    SectionEntity.new(self, data)
  end


  def Site(data = nil)
    require_relative 'entity/site_entity'
    SiteEntity.new(self, data)
  end


  def Sound(data = nil)
    require_relative 'entity/sound_entity'
    SoundEntity.new(self, data)
  end


  def StaticPage(data = nil)
    require_relative 'entity/static_page_entity'
    StaticPageEntity.new(self, data)
  end


  def Text(data = nil)
    require_relative 'entity/text_entity'
    TextEntity.new(self, data)
  end


  def Tour(data = nil)
    require_relative 'entity/tour_entity'
    TourEntity.new(self, data)
  end


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
