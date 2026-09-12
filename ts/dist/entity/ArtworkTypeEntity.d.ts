import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { ArtworkType, ArtworkTypeLoadMatch, ArtworkTypeListMatch } from '../ArtInstituteOfChicagoTypes';
declare class ArtworkTypeEntity extends ArtInstituteOfChicagoEntityBase<ArtworkType> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: ArtworkTypeEntity): ArtworkTypeEntity;
    load(this: any, reqmatch?: ArtworkTypeLoadMatch, ctrl?: Control): Promise<ArtworkTypeEntity>;
    list(this: any, reqmatch?: ArtworkTypeListMatch, ctrl?: Control): Promise<ArtworkTypeEntity[]>;
}
export { ArtworkTypeEntity };
