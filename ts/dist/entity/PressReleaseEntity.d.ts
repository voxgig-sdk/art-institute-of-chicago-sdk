import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { PressRelease, PressReleaseLoadMatch, PressReleaseListMatch } from '../ArtInstituteOfChicagoTypes';
declare class PressReleaseEntity extends ArtInstituteOfChicagoEntityBase<PressRelease> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: PressReleaseEntity): PressReleaseEntity;
    load(this: any, reqmatch?: PressReleaseLoadMatch, ctrl?: Control): Promise<PressReleaseEntity>;
    list(this: any, reqmatch?: PressReleaseListMatch, ctrl?: Control): Promise<PressReleaseEntity[]>;
}
export { PressReleaseEntity };
