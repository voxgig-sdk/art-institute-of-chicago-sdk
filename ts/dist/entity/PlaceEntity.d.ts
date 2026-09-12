import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { Place, PlaceLoadMatch, PlaceListMatch } from '../ArtInstituteOfChicagoTypes';
declare class PlaceEntity extends ArtInstituteOfChicagoEntityBase<Place> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: PlaceEntity): PlaceEntity;
    load(this: any, reqmatch?: PlaceLoadMatch, ctrl?: Control): Promise<PlaceEntity>;
    list(this: any, reqmatch?: PlaceListMatch, ctrl?: Control): Promise<PlaceEntity[]>;
}
export { PlaceEntity };
