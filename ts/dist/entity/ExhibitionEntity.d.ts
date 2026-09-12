import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { Exhibition, ExhibitionLoadMatch, ExhibitionListMatch } from '../ArtInstituteOfChicagoTypes';
declare class ExhibitionEntity extends ArtInstituteOfChicagoEntityBase<Exhibition> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: ExhibitionEntity): ExhibitionEntity;
    load(this: any, reqmatch?: ExhibitionLoadMatch, ctrl?: Control): Promise<ExhibitionEntity>;
    list(this: any, reqmatch?: ExhibitionListMatch, ctrl?: Control): Promise<ExhibitionEntity[]>;
}
export { ExhibitionEntity };
