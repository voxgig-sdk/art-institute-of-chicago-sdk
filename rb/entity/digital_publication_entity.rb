# ArtInstituteOfChicago SDK DigitalPublication entity

require_relative '../utility/struct/voxgig_struct'
require_relative '../core/helpers'

class DigitalPublicationEntity
  def initialize(client, entopts = nil)
    entopts ||= {}
    if entopts["active"].nil?
      entopts["active"] = true
    elsif entopts["active"] == false
      # keep false
    else
      entopts["active"] = true
    end

    @_name = "digital_publication"
    @_client = client
    @_utility = client.get_utility
    @_entopts = entopts
    @_data = {}
    @_match = {}

    @_entctx = @_utility.make_context.call({
      "entity" => self,
      "entopts" => entopts,
    }, client.get_root_ctx)

    @_utility.feature_hook.call(@_entctx, "PostConstructEntity")
  end

  def get_name
    @_name
  end

  def make
    opts = @_entopts.dup
    DigitalPublicationEntity.new(@_client, opts)
  end

  def data_set(args)
    if args
      @_data = ArtInstituteOfChicagoHelpers.to_map(VoxgigStruct.clone(args)) || {}
      @_utility.feature_hook.call(@_entctx, "SetData")
    end
  end

  # @return [DigitalPublication, Hash] the current DigitalPublication data
  def data_get
    @_utility.feature_hook.call(@_entctx, "GetData")
    VoxgigStruct.clone(@_data)
  end

  def match_set(args)
    if args
      @_match = ArtInstituteOfChicagoHelpers.to_map(VoxgigStruct.clone(args)) || {}
      @_utility.feature_hook.call(@_entctx, "SetMatch")
    end
  end

  # @return [Hash] the current match filter (any subset of DigitalPublication fields)
  def match_get
    @_utility.feature_hook.call(@_entctx, "GetMatch")
    VoxgigStruct.clone(@_match)
  end

  
  # Load a single DigitalPublication.
  #
  # @param reqmatch [DigitalPublicationLoadMatch, Hash, nil] match criteria (id/query fields)
  # @param ctrl [Object, nil] optional per-call control
  # @return [DigitalPublication, Hash] the loaded DigitalPublication; raises ArtInstituteOfChicagoError on failure
  def load(reqmatch, ctrl = nil)
    utility = @_utility
    ctx = utility.make_context.call({
      "opname" => "load",
      "ctrl" => ctrl,
      "match" => @_match,
      "data" => @_data,
      "reqmatch" => reqmatch,
    }, @_entctx)

    _run_op(ctx) do
      if ctx.result
        @_match = ctx.result.resmatch if ctx.result.resmatch
        if ctx.result.resdata
          @_data = ArtInstituteOfChicagoHelpers.to_map(VoxgigStruct.clone(ctx.result.resdata)) || {}
        end
      end
    end
  end



  
  # List DigitalPublication items matching the given filter.
  #
  # @param reqmatch [DigitalPublicationListMatch, Hash, nil] match filter (any subset of DigitalPublication fields)
  # @param ctrl [Object, nil] optional per-call control
  # @return [Array<DigitalPublication>, Array] the matching DigitalPublication items; raises ArtInstituteOfChicagoError on failure
  def list(reqmatch, ctrl = nil)
    utility = @_utility
    ctx = utility.make_context.call({
      "opname" => "list",
      "ctrl" => ctrl,
      "match" => @_match,
      "data" => @_data,
      "reqmatch" => reqmatch,
    }, @_entctx)

    _run_op(ctx) do
      if ctx.result
        @_match = ctx.result.resmatch if ctx.result.resmatch
      end
    end
  end



  

  

  

  private

  def _run_op(ctx, &post_done)
    utility = @_utility

    # #PrePoint-Hook

    point, err = utility.make_point.call(ctx)
    ctx.out["point"] = point
    return utility.make_error.call(ctx, err) if err

    # #PreSpec-Hook

    spec, err = utility.make_spec.call(ctx)
    ctx.out["spec"] = spec
    return utility.make_error.call(ctx, err) if err

    # #PreRequest-Hook

    resp, err = utility.make_request.call(ctx)
    ctx.out["request"] = resp
    return utility.make_error.call(ctx, err) if err

    # #PreResponse-Hook

    resp2, err = utility.make_response.call(ctx)
    ctx.out["response"] = resp2
    return utility.make_error.call(ctx, err) if err

    # #PreResult-Hook

    result, err = utility.make_result.call(ctx)
    ctx.out["result"] = result
    return utility.make_error.call(ctx, err) if err

    # #PreDone-Hook

    post_done.call

    utility.done.call(ctx)
  end
end
