import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { Highlight, HighlightLoadMatch, HighlightListMatch } from '../ArtInstituteOfChicagoTypes';
declare class HighlightEntity extends ArtInstituteOfChicagoEntityBase<Highlight> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: HighlightEntity): HighlightEntity;
    load(this: any, reqmatch?: HighlightLoadMatch, ctrl?: Control): Promise<HighlightEntity>;
    list(this: any, reqmatch?: HighlightListMatch, ctrl?: Control): Promise<HighlightEntity[]>;
}
export { HighlightEntity };
