<?php
declare(strict_types=1);

// ArtInstituteOfChicago SDK utility: prepare_body

class ArtInstituteOfChicagoPrepareBody
{
    public static function call(ArtInstituteOfChicagoContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
