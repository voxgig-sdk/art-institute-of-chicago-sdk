<?php
declare(strict_types=1);

// ArtInstituteOfChicago SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class ArtInstituteOfChicagoMakeContext
{
    public static function call(array $ctxmap, ?ArtInstituteOfChicagoContext $basectx): ArtInstituteOfChicagoContext
    {
        return new ArtInstituteOfChicagoContext($ctxmap, $basectx);
    }
}
