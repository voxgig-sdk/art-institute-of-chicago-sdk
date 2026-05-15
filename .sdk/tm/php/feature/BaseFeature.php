<?php
declare(strict_types=1);

// ArtInstituteOfChicago SDK base feature

class ArtInstituteOfChicagoBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(ArtInstituteOfChicagoContext $ctx, array $options): void {}
    public function PostConstruct(ArtInstituteOfChicagoContext $ctx): void {}
    public function PostConstructEntity(ArtInstituteOfChicagoContext $ctx): void {}
    public function SetData(ArtInstituteOfChicagoContext $ctx): void {}
    public function GetData(ArtInstituteOfChicagoContext $ctx): void {}
    public function GetMatch(ArtInstituteOfChicagoContext $ctx): void {}
    public function SetMatch(ArtInstituteOfChicagoContext $ctx): void {}
    public function PrePoint(ArtInstituteOfChicagoContext $ctx): void {}
    public function PreSpec(ArtInstituteOfChicagoContext $ctx): void {}
    public function PreRequest(ArtInstituteOfChicagoContext $ctx): void {}
    public function PreResponse(ArtInstituteOfChicagoContext $ctx): void {}
    public function PreResult(ArtInstituteOfChicagoContext $ctx): void {}
    public function PreDone(ArtInstituteOfChicagoContext $ctx): void {}
    public function PreUnexpected(ArtInstituteOfChicagoContext $ctx): void {}
}
