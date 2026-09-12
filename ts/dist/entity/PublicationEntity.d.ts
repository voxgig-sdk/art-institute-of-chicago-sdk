import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { Publication, PublicationLoadMatch, PublicationListMatch } from '../ArtInstituteOfChicagoTypes';
declare class PublicationEntity extends ArtInstituteOfChicagoEntityBase<Publication> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: PublicationEntity): PublicationEntity;
    load(this: any, reqmatch?: PublicationLoadMatch, ctrl?: Control): Promise<PublicationEntity>;
    list(this: any, reqmatch?: PublicationListMatch, ctrl?: Control): Promise<PublicationEntity[]>;
}
export { PublicationEntity };
