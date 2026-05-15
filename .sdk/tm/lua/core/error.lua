-- ArtInstituteOfChicago SDK error

local ArtInstituteOfChicagoError = {}
ArtInstituteOfChicagoError.__index = ArtInstituteOfChicagoError


function ArtInstituteOfChicagoError.new(code, msg, ctx)
  local self = setmetatable({}, ArtInstituteOfChicagoError)
  self.is_sdk_error = true
  self.sdk = "ArtInstituteOfChicago"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function ArtInstituteOfChicagoError:error()
  return self.msg
end


function ArtInstituteOfChicagoError:__tostring()
  return self.msg
end


return ArtInstituteOfChicagoError
