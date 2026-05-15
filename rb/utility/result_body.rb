# ArtInstituteOfChicago SDK utility: result_body
module ArtInstituteOfChicagoUtilities
  ResultBody = ->(ctx) {
    response = ctx.response
    result = ctx.result
    if result && response && response.json_func && response.body
      result.body = response.json_func.call
    end
    result
  }
end
