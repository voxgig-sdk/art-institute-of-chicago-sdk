import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { Sound, SoundLoadMatch, SoundListMatch } from '../ArtInstituteOfChicagoTypes';
declare class SoundEntity extends ArtInstituteOfChicagoEntityBase<Sound> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: SoundEntity): SoundEntity;
    load(this: any, reqmatch?: SoundLoadMatch, ctrl?: Control): Promise<SoundEntity>;
    list(this: any, reqmatch?: SoundListMatch, ctrl?: Control): Promise<SoundEntity[]>;
}
export { SoundEntity };
