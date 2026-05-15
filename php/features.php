<?php
declare(strict_types=1);

// ArtInstituteOfChicago SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class ArtInstituteOfChicagoFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new ArtInstituteOfChicagoBaseFeature();
            case "test":
                return new ArtInstituteOfChicagoTestFeature();
            default:
                return new ArtInstituteOfChicagoBaseFeature();
        }
    }
}
