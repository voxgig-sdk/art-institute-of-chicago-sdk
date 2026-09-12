import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { StaticPage, StaticPageLoadMatch, StaticPageListMatch } from '../ArtInstituteOfChicagoTypes';
declare class StaticPageEntity extends ArtInstituteOfChicagoEntityBase<StaticPage> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: StaticPageEntity): StaticPageEntity;
    load(this: any, reqmatch?: StaticPageLoadMatch, ctrl?: Control): Promise<StaticPageEntity>;
    list(this: any, reqmatch?: StaticPageListMatch, ctrl?: Control): Promise<StaticPageEntity[]>;
}
export { StaticPageEntity };
