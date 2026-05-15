# ArtInstituteOfChicago SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module ArtInstituteOfChicagoFeatures
  def self.make_feature(name)
    case name
    when "base"
      ArtInstituteOfChicagoBaseFeature.new
    when "test"
      ArtInstituteOfChicagoTestFeature.new
    else
      ArtInstituteOfChicagoBaseFeature.new
    end
  end
end
