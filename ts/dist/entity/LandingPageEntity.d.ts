import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { LandingPage, LandingPageLoadMatch, LandingPageListMatch } from '../ArtInstituteOfChicagoTypes';
declare class LandingPageEntity extends ArtInstituteOfChicagoEntityBase<LandingPage> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: LandingPageEntity): LandingPageEntity;
    load(this: any, reqmatch?: LandingPageLoadMatch, ctrl?: Control): Promise<LandingPageEntity>;
    list(this: any, reqmatch?: LandingPageListMatch, ctrl?: Control): Promise<LandingPageEntity[]>;
}
export { LandingPageEntity };
