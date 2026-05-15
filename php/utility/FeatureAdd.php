<?php
declare(strict_types=1);

// ArtInstituteOfChicago SDK utility: feature_add

class ArtInstituteOfChicagoFeatureAdd
{
    public static function call(ArtInstituteOfChicagoContext $ctx, mixed $f): void
    {
        $ctx->client->features[] = $f;
    }
}
