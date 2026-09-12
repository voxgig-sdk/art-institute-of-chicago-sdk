import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { Agent, AgentLoadMatch, AgentListMatch } from '../ArtInstituteOfChicagoTypes';
declare class AgentEntity extends ArtInstituteOfChicagoEntityBase<Agent> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: AgentEntity): AgentEntity;
    load(this: any, reqmatch?: AgentLoadMatch, ctrl?: Control): Promise<AgentEntity>;
    list(this: any, reqmatch?: AgentListMatch, ctrl?: Control): Promise<AgentEntity[]>;
}
export { AgentEntity };
