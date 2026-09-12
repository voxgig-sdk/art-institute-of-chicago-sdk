import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { AgentRole, AgentRoleLoadMatch, AgentRoleListMatch } from '../ArtInstituteOfChicagoTypes';
declare class AgentRoleEntity extends ArtInstituteOfChicagoEntityBase<AgentRole> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: AgentRoleEntity): AgentRoleEntity;
    load(this: any, reqmatch?: AgentRoleLoadMatch, ctrl?: Control): Promise<AgentRoleEntity>;
    list(this: any, reqmatch?: AgentRoleListMatch, ctrl?: Control): Promise<AgentRoleEntity[]>;
}
export { AgentRoleEntity };
