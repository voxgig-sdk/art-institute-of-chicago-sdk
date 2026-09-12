import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { Gallery, GalleryLoadMatch, GalleryListMatch } from '../ArtInstituteOfChicagoTypes';
declare class GalleryEntity extends ArtInstituteOfChicagoEntityBase<Gallery> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: GalleryEntity): GalleryEntity;
    load(this: any, reqmatch?: GalleryLoadMatch, ctrl?: Control): Promise<GalleryEntity>;
    list(this: any, reqmatch?: GalleryListMatch, ctrl?: Control): Promise<GalleryEntity[]>;
}
export { GalleryEntity };
