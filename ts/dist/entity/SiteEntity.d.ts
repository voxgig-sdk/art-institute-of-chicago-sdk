import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { Site, SiteLoadMatch, SiteListMatch } from '../ArtInstituteOfChicagoTypes';
declare class SiteEntity extends ArtInstituteOfChicagoEntityBase<Site> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: SiteEntity): SiteEntity;
    load(this: any, reqmatch?: SiteLoadMatch, ctrl?: Control): Promise<SiteEntity>;
    list(this: any, reqmatch?: SiteListMatch, ctrl?: Control): Promise<SiteEntity[]>;
}
export { SiteEntity };
