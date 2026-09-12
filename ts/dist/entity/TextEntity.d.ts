import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { Text, TextLoadMatch, TextListMatch } from '../ArtInstituteOfChicagoTypes';
declare class TextEntity extends ArtInstituteOfChicagoEntityBase<Text> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: TextEntity): TextEntity;
    load(this: any, reqmatch?: TextLoadMatch, ctrl?: Control): Promise<TextEntity>;
    list(this: any, reqmatch?: TextListMatch, ctrl?: Control): Promise<TextEntity[]>;
}
export { TextEntity };
