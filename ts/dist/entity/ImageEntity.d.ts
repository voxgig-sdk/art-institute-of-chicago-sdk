import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { Image, ImageLoadMatch, ImageListMatch } from '../ArtInstituteOfChicagoTypes';
declare class ImageEntity extends ArtInstituteOfChicagoEntityBase<Image> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: ImageEntity): ImageEntity;
    load(this: any, reqmatch?: ImageLoadMatch, ctrl?: Control): Promise<ImageEntity>;
    list(this: any, reqmatch?: ImageListMatch, ctrl?: Control): Promise<ImageEntity[]>;
}
export { ImageEntity };
