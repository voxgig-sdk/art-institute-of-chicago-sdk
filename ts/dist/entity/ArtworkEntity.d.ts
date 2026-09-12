import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { Artwork, ArtworkLoadMatch, ArtworkListMatch } from '../ArtInstituteOfChicagoTypes';
declare class ArtworkEntity extends ArtInstituteOfChicagoEntityBase<Artwork> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: ArtworkEntity): ArtworkEntity;
    load(this: any, reqmatch?: ArtworkLoadMatch, ctrl?: Control): Promise<ArtworkEntity>;
    list(this: any, reqmatch?: ArtworkListMatch, ctrl?: Control): Promise<ArtworkEntity[]>;
}
export { ArtworkEntity };
