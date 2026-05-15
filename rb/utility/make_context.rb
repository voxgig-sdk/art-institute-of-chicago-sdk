# ArtInstituteOfChicago SDK utility: make_context
require_relative '../core/context'
module ArtInstituteOfChicagoUtilities
  MakeContext = ->(ctxmap, basectx) {
    ArtInstituteOfChicagoContext.new(ctxmap, basectx)
  }
end
