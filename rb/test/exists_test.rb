# ArtInstituteOfChicago SDK exists test

require "minitest/autorun"
require_relative "../ArtInstituteOfChicago_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = ArtInstituteOfChicagoSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
