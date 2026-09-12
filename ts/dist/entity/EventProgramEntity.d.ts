import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { EventProgram, EventProgramLoadMatch, EventProgramListMatch } from '../ArtInstituteOfChicagoTypes';
declare class EventProgramEntity extends ArtInstituteOfChicagoEntityBase<EventProgram> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: EventProgramEntity): EventProgramEntity;
    load(this: any, reqmatch?: EventProgramLoadMatch, ctrl?: Control): Promise<EventProgramEntity>;
    list(this: any, reqmatch?: EventProgramListMatch, ctrl?: Control): Promise<EventProgramEntity[]>;
}
export { EventProgramEntity };
