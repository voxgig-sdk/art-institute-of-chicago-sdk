import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { AgentType, AgentTypeLoadMatch, AgentTypeListMatch } from '../ArtInstituteOfChicagoTypes';
declare class AgentTypeEntity extends ArtInstituteOfChicagoEntityBase<AgentType> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: AgentTypeEntity): AgentTypeEntity;
    load(this: any, reqmatch?: AgentTypeLoadMatch, ctrl?: Control): Promise<AgentTypeEntity>;
    list(this: any, reqmatch?: AgentTypeListMatch, ctrl?: Control): Promise<AgentTypeEntity[]>;
}
export { AgentTypeEntity };
