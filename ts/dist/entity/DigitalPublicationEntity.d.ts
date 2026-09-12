import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { DigitalPublication, DigitalPublicationLoadMatch, DigitalPublicationListMatch } from '../ArtInstituteOfChicagoTypes';
declare class DigitalPublicationEntity extends ArtInstituteOfChicagoEntityBase<DigitalPublication> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: DigitalPublicationEntity): DigitalPublicationEntity;
    load(this: any, reqmatch?: DigitalPublicationLoadMatch, ctrl?: Control): Promise<DigitalPublicationEntity>;
    list(this: any, reqmatch?: DigitalPublicationListMatch, ctrl?: Control): Promise<DigitalPublicationEntity[]>;
}
export { DigitalPublicationEntity };
