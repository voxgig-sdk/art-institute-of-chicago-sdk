import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { EventOccurrence, EventOccurrenceLoadMatch, EventOccurrenceListMatch } from '../ArtInstituteOfChicagoTypes';
declare class EventOccurrenceEntity extends ArtInstituteOfChicagoEntityBase<EventOccurrence> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: EventOccurrenceEntity): EventOccurrenceEntity;
    load(this: any, reqmatch?: EventOccurrenceLoadMatch, ctrl?: Control): Promise<EventOccurrenceEntity>;
    list(this: any, reqmatch?: EventOccurrenceListMatch, ctrl?: Control): Promise<EventOccurrenceEntity[]>;
}
export { EventOccurrenceEntity };
