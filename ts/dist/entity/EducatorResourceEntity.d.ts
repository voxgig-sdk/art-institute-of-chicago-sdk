import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { EducatorResource, EducatorResourceLoadMatch, EducatorResourceListMatch } from '../ArtInstituteOfChicagoTypes';
declare class EducatorResourceEntity extends ArtInstituteOfChicagoEntityBase<EducatorResource> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: EducatorResourceEntity): EducatorResourceEntity;
    load(this: any, reqmatch?: EducatorResourceLoadMatch, ctrl?: Control): Promise<EducatorResourceEntity>;
    list(this: any, reqmatch?: EducatorResourceListMatch, ctrl?: Control): Promise<EducatorResourceEntity[]>;
}
export { EducatorResourceEntity };
