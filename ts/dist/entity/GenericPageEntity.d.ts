import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { GenericPage, GenericPageLoadMatch, GenericPageListMatch } from '../ArtInstituteOfChicagoTypes';
declare class GenericPageEntity extends ArtInstituteOfChicagoEntityBase<GenericPage> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: GenericPageEntity): GenericPageEntity;
    load(this: any, reqmatch?: GenericPageLoadMatch, ctrl?: Control): Promise<GenericPageEntity>;
    list(this: any, reqmatch?: GenericPageListMatch, ctrl?: Control): Promise<GenericPageEntity[]>;
}
export { GenericPageEntity };
