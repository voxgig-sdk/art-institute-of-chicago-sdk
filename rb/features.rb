# ArtInstituteOfChicago SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/ratelimit_feature'
require_relative 'feature/retry_feature'
require_relative 'feature/test_feature'
require_relative 'feature/timeout_feature'


module ArtInstituteOfChicagoFeatures
  def self.make_feature(name)
    case name
    when "base"
      ArtInstituteOfChicagoBaseFeature.new
    when "ratelimit"
      ArtInstituteOfChicagoRatelimitFeature.new
    when "retry"
      ArtInstituteOfChicagoRetryFeature.new
    when "test"
      ArtInstituteOfChicagoTestFeature.new
    when "timeout"
      ArtInstituteOfChicagoTimeoutFeature.new
    else
      ArtInstituteOfChicagoBaseFeature.new
    end
  end
end
