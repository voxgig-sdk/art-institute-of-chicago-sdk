import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { CategoryTerm, CategoryTermLoadMatch, CategoryTermListMatch } from '../ArtInstituteOfChicagoTypes';
declare class CategoryTermEntity extends ArtInstituteOfChicagoEntityBase<CategoryTerm> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: CategoryTermEntity): CategoryTermEntity;
    load(this: any, reqmatch?: CategoryTermLoadMatch, ctrl?: Control): Promise<CategoryTermEntity>;
    list(this: any, reqmatch?: CategoryTermListMatch, ctrl?: Control): Promise<CategoryTermEntity[]>;
}
export { CategoryTermEntity };
