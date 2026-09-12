import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { Tour, TourLoadMatch, TourListMatch } from '../ArtInstituteOfChicagoTypes';
declare class TourEntity extends ArtInstituteOfChicagoEntityBase<Tour> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: TourEntity): TourEntity;
    load(this: any, reqmatch?: TourLoadMatch, ctrl?: Control): Promise<TourEntity>;
    list(this: any, reqmatch?: TourListMatch, ctrl?: Control): Promise<TourEntity[]>;
}
export { TourEntity };
