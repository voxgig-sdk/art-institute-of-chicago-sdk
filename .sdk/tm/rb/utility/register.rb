# ArtInstituteOfChicago SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

ArtInstituteOfChicagoUtility.registrar = ->(u) {
  u.clean = ArtInstituteOfChicagoUtilities::Clean
  u.done = ArtInstituteOfChicagoUtilities::Done
  u.make_error = ArtInstituteOfChicagoUtilities::MakeError
  u.feature_add = ArtInstituteOfChicagoUtilities::FeatureAdd
  u.feature_hook = ArtInstituteOfChicagoUtilities::FeatureHook
  u.feature_init = ArtInstituteOfChicagoUtilities::FeatureInit
  u.fetcher = ArtInstituteOfChicagoUtilities::Fetcher
  u.make_fetch_def = ArtInstituteOfChicagoUtilities::MakeFetchDef
  u.make_context = ArtInstituteOfChicagoUtilities::MakeContext
  u.make_options = ArtInstituteOfChicagoUtilities::MakeOptions
  u.make_request = ArtInstituteOfChicagoUtilities::MakeRequest
  u.make_response = ArtInstituteOfChicagoUtilities::MakeResponse
  u.make_result = ArtInstituteOfChicagoUtilities::MakeResult
  u.make_point = ArtInstituteOfChicagoUtilities::MakePoint
  u.make_spec = ArtInstituteOfChicagoUtilities::MakeSpec
  u.make_url = ArtInstituteOfChicagoUtilities::MakeUrl
  u.param = ArtInstituteOfChicagoUtilities::Param
  u.prepare_auth = ArtInstituteOfChicagoUtilities::PrepareAuth
  u.prepare_body = ArtInstituteOfChicagoUtilities::PrepareBody
  u.prepare_headers = ArtInstituteOfChicagoUtilities::PrepareHeaders
  u.prepare_method = ArtInstituteOfChicagoUtilities::PrepareMethod
  u.prepare_params = ArtInstituteOfChicagoUtilities::PrepareParams
  u.prepare_path = ArtInstituteOfChicagoUtilities::PreparePath
  u.prepare_query = ArtInstituteOfChicagoUtilities::PrepareQuery
  u.result_basic = ArtInstituteOfChicagoUtilities::ResultBasic
  u.result_body = ArtInstituteOfChicagoUtilities::ResultBody
  u.result_headers = ArtInstituteOfChicagoUtilities::ResultHeaders
  u.transform_request = ArtInstituteOfChicagoUtilities::TransformRequest
  u.transform_response = ArtInstituteOfChicagoUtilities::TransformResponse
}
