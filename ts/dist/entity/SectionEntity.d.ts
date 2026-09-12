import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { Section, SectionLoadMatch, SectionListMatch } from '../ArtInstituteOfChicagoTypes';
declare class SectionEntity extends ArtInstituteOfChicagoEntityBase<Section> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: SectionEntity): SectionEntity;
    load(this: any, reqmatch?: SectionLoadMatch, ctrl?: Control): Promise<SectionEntity>;
    list(this: any, reqmatch?: SectionListMatch, ctrl?: Control): Promise<SectionEntity[]>;
}
export { SectionEntity };
