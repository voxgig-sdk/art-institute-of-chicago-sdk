import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { PrintedPublication, PrintedPublicationLoadMatch, PrintedPublicationListMatch } from '../ArtInstituteOfChicagoTypes';
declare class PrintedPublicationEntity extends ArtInstituteOfChicagoEntityBase<PrintedPublication> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: PrintedPublicationEntity): PrintedPublicationEntity;
    load(this: any, reqmatch?: PrintedPublicationLoadMatch, ctrl?: Control): Promise<PrintedPublicationEntity>;
    list(this: any, reqmatch?: PrintedPublicationListMatch, ctrl?: Control): Promise<PrintedPublicationEntity[]>;
}
export { PrintedPublicationEntity };
