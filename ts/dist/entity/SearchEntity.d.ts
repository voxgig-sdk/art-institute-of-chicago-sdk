import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { Search, SearchListMatch } from '../ArtInstituteOfChicagoTypes';
declare class SearchEntity extends ArtInstituteOfChicagoEntityBase<Search> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: SearchEntity): SearchEntity;
    list(this: any, reqmatch?: SearchListMatch, ctrl?: Control): Promise<SearchEntity[]>;
}
export { SearchEntity };
