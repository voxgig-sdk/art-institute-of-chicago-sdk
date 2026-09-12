import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { ArtworkDateQualifier, ArtworkDateQualifierLoadMatch, ArtworkDateQualifierListMatch } from '../ArtInstituteOfChicagoTypes';
declare class ArtworkDateQualifierEntity extends ArtInstituteOfChicagoEntityBase<ArtworkDateQualifier> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: ArtworkDateQualifierEntity): ArtworkDateQualifierEntity;
    load(this: any, reqmatch?: ArtworkDateQualifierLoadMatch, ctrl?: Control): Promise<ArtworkDateQualifierEntity>;
    list(this: any, reqmatch?: ArtworkDateQualifierListMatch, ctrl?: Control): Promise<ArtworkDateQualifierEntity[]>;
}
export { ArtworkDateQualifierEntity };
