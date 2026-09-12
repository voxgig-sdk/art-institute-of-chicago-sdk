import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { Event, EventLoadMatch, EventListMatch } from '../ArtInstituteOfChicagoTypes';
declare class EventEntity extends ArtInstituteOfChicagoEntityBase<Event> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: EventEntity): EventEntity;
    load(this: any, reqmatch?: EventLoadMatch, ctrl?: Control): Promise<EventEntity>;
    list(this: any, reqmatch?: EventListMatch, ctrl?: Control): Promise<EventEntity[]>;
}
export { EventEntity };
