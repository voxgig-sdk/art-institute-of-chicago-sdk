import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { ArtworkPlaceQualifier, ArtworkPlaceQualifierLoadMatch, ArtworkPlaceQualifierListMatch } from '../ArtInstituteOfChicagoTypes';
declare class ArtworkPlaceQualifierEntity extends ArtInstituteOfChicagoEntityBase<ArtworkPlaceQualifier> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: ArtworkPlaceQualifierEntity): ArtworkPlaceQualifierEntity;
    load(this: any, reqmatch?: ArtworkPlaceQualifierLoadMatch, ctrl?: Control): Promise<ArtworkPlaceQualifierEntity>;
    list(this: any, reqmatch?: ArtworkPlaceQualifierListMatch, ctrl?: Control): Promise<ArtworkPlaceQualifierEntity[]>;
}
export { ArtworkPlaceQualifierEntity };
