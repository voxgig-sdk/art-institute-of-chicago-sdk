<?php
declare(strict_types=1);

// ArtInstituteOfChicago SDK utility: result_headers

class ArtInstituteOfChicagoResultHeaders
{
    public static function call(ArtInstituteOfChicagoContext $ctx): ?ArtInstituteOfChicagoResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
