-- ArtInstituteOfChicago SDK

local vs = require("utility.struct.struct")
local Utility = require("core.utility_type")
local Spec = require("core.spec")
local helpers = require("core.helpers")

-- Load utility registration (populates Utility._registrar)
require("utility.register")

-- Load features
local BaseFeature = require("feature.base_feature")
local features_factory = require("features")


local ArtInstituteOfChicagoSDK = {}
ArtInstituteOfChicagoSDK.__index = ArtInstituteOfChicagoSDK


local function _make_feature(name)
  local factory = features_factory[name]
  if factory ~= nil then
    return factory()
  end
  return features_factory.base()
end

ArtInstituteOfChicagoSDK._make_feature = _make_feature


function ArtInstituteOfChicagoSDK.new(options)
  local self = setmetatable({}, ArtInstituteOfChicagoSDK)
  self.mode = "live"
  self.features = {}
  self.options = nil

  local utility = Utility.new()
  self._utility = utility

  local config = require("config")()

  self._rootctx = utility.make_context({
    client = self,
    utility = utility,
    config = config,
    options = options or {},
    shared = {},
  }, nil)

  self.options = utility.make_options(self._rootctx)

  if vs.getpath(self.options, "feature.test.active") == true then
    self.mode = "test"
  end

  self._rootctx.options = self.options

  -- Add features from config.
  local feature_opts = helpers.to_map(vs.getprop(self.options, "feature"))
  if feature_opts ~= nil then
    local feature_items = vs.items(feature_opts)
    if feature_items ~= nil then
      for _, item in ipairs(feature_items) do
        local fname = item[1]
        local fopts = helpers.to_map(item[2])
        if fopts ~= nil and fopts["active"] == true then
          utility.feature_add(self._rootctx, _make_feature(fname))
        end
      end
    end
  end

  -- Add extension features.
  local extend = vs.getprop(self.options, "extend")
  if type(extend) == "table" then
    for _, f in ipairs(extend) do
      if type(f) == "table" and type(f.get_name) == "function" then
        utility.feature_add(self._rootctx, f)
      end
    end
  end

  -- Initialize features.
  for _, f in ipairs(self.features) do
    utility.feature_init(self._rootctx, f)
  end

  utility.feature_hook(self._rootctx, "PostConstruct")

  -- #BuildFeatures

  return self
end


function ArtInstituteOfChicagoSDK:options_map()
  local out = vs.clone(self.options)
  if type(out) == "table" then
    return out
  end
  return {}
end


function ArtInstituteOfChicagoSDK:get_utility()
  return Utility.copy(self._utility)
end


function ArtInstituteOfChicagoSDK:get_root_ctx()
  return self._rootctx
end


function ArtInstituteOfChicagoSDK:prepare(fetchargs)
  local utility = self._utility

  fetchargs = fetchargs or {}

  local ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl")) or {}

  local ctx = utility.make_context({
    opname = "prepare",
    ctrl = ctrl,
  }, self._rootctx)

  local options = self.options

  local path = vs.getprop(fetchargs, "path") or ""
  if type(path) ~= "string" then path = "" end

  local method = vs.getprop(fetchargs, "method") or "GET"
  if type(method) ~= "string" then method = "GET" end

  local params = helpers.to_map(vs.getprop(fetchargs, "params")) or {}
  local query = helpers.to_map(vs.getprop(fetchargs, "query")) or {}

  local headers = utility.prepare_headers(ctx)

  local base = vs.getprop(options, "base") or ""
  if type(base) ~= "string" then base = "" end
  local prefix = vs.getprop(options, "prefix") or ""
  if type(prefix) ~= "string" then prefix = "" end
  local suffix = vs.getprop(options, "suffix") or ""
  if type(suffix) ~= "string" then suffix = "" end

  ctx.spec = Spec.new({
    base = base,
    prefix = prefix,
    suffix = suffix,
    path = path,
    method = method,
    params = params,
    query = query,
    headers = headers,
    body = vs.getprop(fetchargs, "body"),
    step = "start",
  })

  -- Merge user-provided headers.
  local uh = vs.getprop(fetchargs, "headers")
  if type(uh) == "table" then
    for k, v in pairs(uh) do
      ctx.spec.headers[k] = v
    end
  end

  local _, err = utility.prepare_auth(ctx)
  if err ~= nil then
    return nil, err
  end

  return utility.make_fetch_def(ctx)
end


function ArtInstituteOfChicagoSDK:direct(fetchargs)
  local utility = self._utility

  local fetchdef, err = self:prepare(fetchargs)
  if err ~= nil then
    return { ok = false, err = err }, nil
  end

  fetchargs = fetchargs or {}
  local ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl")) or {}

  local ctx = utility.make_context({
    opname = "direct",
    ctrl = ctrl,
  }, self._rootctx)

  local url = fetchdef["url"] or ""
  local fetched, fetch_err = utility.fetcher(ctx, url, fetchdef)

  if fetch_err ~= nil then
    return { ok = false, err = fetch_err }, nil
  end

  if fetched == nil then
    return {
      ok = false,
      err = ctx:make_error("direct_no_response", "response: undefined"),
    }, nil
  end

  if type(fetched) == "table" then
    local status = helpers.to_int(vs.getprop(fetched, "status"))
    local headers = vs.getprop(fetched, "headers") or {}

    -- No-body responses (204, 304) and explicit zero content-length
    -- must skip JSON parsing — calling json() on an empty body errors.
    local content_length = nil
    if type(headers) == "table" then
      content_length = headers["content-length"]
    end
    local no_body = status == 204 or status == 304 or tostring(content_length) == "0"

    local json_data = nil
    if not no_body then
      local jf = vs.getprop(fetched, "json")
      if type(jf) == "function" then
        local ok, result = pcall(jf)
        if ok then
          json_data = result
        end
        -- Non-JSON body: json_data stays nil, status/headers preserved.
      end
    end

    return {
      ok = status >= 200 and status < 300,
      status = status,
      headers = headers,
      data = json_data,
    }, nil
  end

  return {
    ok = false,
    err = ctx:make_error("direct_invalid", "invalid response type"),
  }, nil
end



-- Idiomatic facade: client:Agent():list() / client:Agent():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:Agent(data)
  local EntityMod = require("entity.agent_entity")
  if data == nil then
    if self._agent == nil then
      self._agent = EntityMod.new(self, nil)
    end
    return self._agent
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:AgentRole():list() / client:AgentRole():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:AgentRole(data)
  local EntityMod = require("entity.agent_role_entity")
  if data == nil then
    if self._agent_role == nil then
      self._agent_role = EntityMod.new(self, nil)
    end
    return self._agent_role
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:AgentType():list() / client:AgentType():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:AgentType(data)
  local EntityMod = require("entity.agent_type_entity")
  if data == nil then
    if self._agent_type == nil then
      self._agent_type = EntityMod.new(self, nil)
    end
    return self._agent_type
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Article():list() / client:Article():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:Article(data)
  local EntityMod = require("entity.article_entity")
  if data == nil then
    if self._article == nil then
      self._article = EntityMod.new(self, nil)
    end
    return self._article
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Artwork():list() / client:Artwork():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:Artwork(data)
  local EntityMod = require("entity.artwork_entity")
  if data == nil then
    if self._artwork == nil then
      self._artwork = EntityMod.new(self, nil)
    end
    return self._artwork
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ArtworkDateQualifier():list() / client:ArtworkDateQualifier():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:ArtworkDateQualifier(data)
  local EntityMod = require("entity.artwork_date_qualifier_entity")
  if data == nil then
    if self._artwork_date_qualifier == nil then
      self._artwork_date_qualifier = EntityMod.new(self, nil)
    end
    return self._artwork_date_qualifier
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ArtworkPlaceQualifier():list() / client:ArtworkPlaceQualifier():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:ArtworkPlaceQualifier(data)
  local EntityMod = require("entity.artwork_place_qualifier_entity")
  if data == nil then
    if self._artwork_place_qualifier == nil then
      self._artwork_place_qualifier = EntityMod.new(self, nil)
    end
    return self._artwork_place_qualifier
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ArtworkType():list() / client:ArtworkType():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:ArtworkType(data)
  local EntityMod = require("entity.artwork_type_entity")
  if data == nil then
    if self._artwork_type == nil then
      self._artwork_type = EntityMod.new(self, nil)
    end
    return self._artwork_type
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:CategoryTerm():list() / client:CategoryTerm():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:CategoryTerm(data)
  local EntityMod = require("entity.category_term_entity")
  if data == nil then
    if self._category_term == nil then
      self._category_term = EntityMod.new(self, nil)
    end
    return self._category_term
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:DigitalPublication():list() / client:DigitalPublication():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:DigitalPublication(data)
  local EntityMod = require("entity.digital_publication_entity")
  if data == nil then
    if self._digital_publication == nil then
      self._digital_publication = EntityMod.new(self, nil)
    end
    return self._digital_publication
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:DigitalPublicationArticle():list() / client:DigitalPublicationArticle():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:DigitalPublicationArticle(data)
  local EntityMod = require("entity.digital_publication_article_entity")
  if data == nil then
    if self._digital_publication_article == nil then
      self._digital_publication_article = EntityMod.new(self, nil)
    end
    return self._digital_publication_article
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EducatorResource():list() / client:EducatorResource():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:EducatorResource(data)
  local EntityMod = require("entity.educator_resource_entity")
  if data == nil then
    if self._educator_resource == nil then
      self._educator_resource = EntityMod.new(self, nil)
    end
    return self._educator_resource
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Event():list() / client:Event():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:Event(data)
  local EntityMod = require("entity.event_entity")
  if data == nil then
    if self._event == nil then
      self._event = EntityMod.new(self, nil)
    end
    return self._event
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EventOccurrence():list() / client:EventOccurrence():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:EventOccurrence(data)
  local EntityMod = require("entity.event_occurrence_entity")
  if data == nil then
    if self._event_occurrence == nil then
      self._event_occurrence = EntityMod.new(self, nil)
    end
    return self._event_occurrence
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EventProgram():list() / client:EventProgram():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:EventProgram(data)
  local EntityMod = require("entity.event_program_entity")
  if data == nil then
    if self._event_program == nil then
      self._event_program = EntityMod.new(self, nil)
    end
    return self._event_program
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Exhibition():list() / client:Exhibition():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:Exhibition(data)
  local EntityMod = require("entity.exhibition_entity")
  if data == nil then
    if self._exhibition == nil then
      self._exhibition = EntityMod.new(self, nil)
    end
    return self._exhibition
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Gallery():list() / client:Gallery():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:Gallery(data)
  local EntityMod = require("entity.gallery_entity")
  if data == nil then
    if self._gallery == nil then
      self._gallery = EntityMod.new(self, nil)
    end
    return self._gallery
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:GenericPage():list() / client:GenericPage():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:GenericPage(data)
  local EntityMod = require("entity.generic_page_entity")
  if data == nil then
    if self._generic_page == nil then
      self._generic_page = EntityMod.new(self, nil)
    end
    return self._generic_page
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Highlight():list() / client:Highlight():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:Highlight(data)
  local EntityMod = require("entity.highlight_entity")
  if data == nil then
    if self._highlight == nil then
      self._highlight = EntityMod.new(self, nil)
    end
    return self._highlight
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Hour():list() / client:Hour():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:Hour(data)
  local EntityMod = require("entity.hour_entity")
  if data == nil then
    if self._hour == nil then
      self._hour = EntityMod.new(self, nil)
    end
    return self._hour
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Image():list() / client:Image():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:Image(data)
  local EntityMod = require("entity.image_entity")
  if data == nil then
    if self._image == nil then
      self._image = EntityMod.new(self, nil)
    end
    return self._image
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:LandingPage():list() / client:LandingPage():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:LandingPage(data)
  local EntityMod = require("entity.landing_page_entity")
  if data == nil then
    if self._landing_page == nil then
      self._landing_page = EntityMod.new(self, nil)
    end
    return self._landing_page
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Place():list() / client:Place():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:Place(data)
  local EntityMod = require("entity.place_entity")
  if data == nil then
    if self._place == nil then
      self._place = EntityMod.new(self, nil)
    end
    return self._place
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PressRelease():list() / client:PressRelease():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:PressRelease(data)
  local EntityMod = require("entity.press_release_entity")
  if data == nil then
    if self._press_release == nil then
      self._press_release = EntityMod.new(self, nil)
    end
    return self._press_release
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PrintedPublication():list() / client:PrintedPublication():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:PrintedPublication(data)
  local EntityMod = require("entity.printed_publication_entity")
  if data == nil then
    if self._printed_publication == nil then
      self._printed_publication = EntityMod.new(self, nil)
    end
    return self._printed_publication
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Product():list() / client:Product():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:Product(data)
  local EntityMod = require("entity.product_entity")
  if data == nil then
    if self._product == nil then
      self._product = EntityMod.new(self, nil)
    end
    return self._product
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Publication():list() / client:Publication():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:Publication(data)
  local EntityMod = require("entity.publication_entity")
  if data == nil then
    if self._publication == nil then
      self._publication = EntityMod.new(self, nil)
    end
    return self._publication
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Search():list() / client:Search():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:Search(data)
  local EntityMod = require("entity.search_entity")
  if data == nil then
    if self._search == nil then
      self._search = EntityMod.new(self, nil)
    end
    return self._search
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Section():list() / client:Section():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:Section(data)
  local EntityMod = require("entity.section_entity")
  if data == nil then
    if self._section == nil then
      self._section = EntityMod.new(self, nil)
    end
    return self._section
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Site():list() / client:Site():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:Site(data)
  local EntityMod = require("entity.site_entity")
  if data == nil then
    if self._site == nil then
      self._site = EntityMod.new(self, nil)
    end
    return self._site
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Sound():list() / client:Sound():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:Sound(data)
  local EntityMod = require("entity.sound_entity")
  if data == nil then
    if self._sound == nil then
      self._sound = EntityMod.new(self, nil)
    end
    return self._sound
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:StaticPage():list() / client:StaticPage():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:StaticPage(data)
  local EntityMod = require("entity.static_page_entity")
  if data == nil then
    if self._static_page == nil then
      self._static_page = EntityMod.new(self, nil)
    end
    return self._static_page
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Text():list() / client:Text():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:Text(data)
  local EntityMod = require("entity.text_entity")
  if data == nil then
    if self._text == nil then
      self._text = EntityMod.new(self, nil)
    end
    return self._text
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Tour():list() / client:Tour():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:Tour(data)
  local EntityMod = require("entity.tour_entity")
  if data == nil then
    if self._tour == nil then
      self._tour = EntityMod.new(self, nil)
    end
    return self._tour
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Video():list() / client:Video():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function ArtInstituteOfChicagoSDK:Video(data)
  local EntityMod = require("entity.video_entity")
  if data == nil then
    if self._video == nil then
      self._video = EntityMod.new(self, nil)
    end
    return self._video
  end
  return EntityMod.new(self, data)
end




function ArtInstituteOfChicagoSDK.test(testopts, sdkopts)
  sdkopts = sdkopts or {}
  sdkopts = vs.clone(sdkopts)
  if type(sdkopts) ~= "table" then
    sdkopts = {}
  end

  testopts = testopts or {}
  testopts = vs.clone(testopts)
  if type(testopts) ~= "table" then
    testopts = {}
  end
  testopts["active"] = true

  vs.setpath(sdkopts, "feature.test", testopts)

  local sdk = ArtInstituteOfChicagoSDK.new(sdkopts)
  sdk.mode = "test"

  return sdk
end


return ArtInstituteOfChicagoSDK
