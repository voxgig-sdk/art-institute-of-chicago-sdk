# ArtInstituteOfChicago SDK utility: feature_add
module ArtInstituteOfChicagoUtilities
  FeatureAdd = ->(ctx, f) {
    ctx.client.features << f
  }
end
