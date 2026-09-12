import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { Hour, HourLoadMatch, HourListMatch } from '../ArtInstituteOfChicagoTypes';
declare class HourEntity extends ArtInstituteOfChicagoEntityBase<Hour> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: HourEntity): HourEntity;
    load(this: any, reqmatch?: HourLoadMatch, ctrl?: Control): Promise<HourEntity>;
    list(this: any, reqmatch?: HourListMatch, ctrl?: Control): Promise<HourEntity[]>;
}
export { HourEntity };
