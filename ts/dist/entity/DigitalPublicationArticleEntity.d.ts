import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { DigitalPublicationArticle, DigitalPublicationArticleLoadMatch, DigitalPublicationArticleListMatch } from '../ArtInstituteOfChicagoTypes';
declare class DigitalPublicationArticleEntity extends ArtInstituteOfChicagoEntityBase<DigitalPublicationArticle> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: DigitalPublicationArticleEntity): DigitalPublicationArticleEntity;
    load(this: any, reqmatch?: DigitalPublicationArticleLoadMatch, ctrl?: Control): Promise<DigitalPublicationArticleEntity>;
    list(this: any, reqmatch?: DigitalPublicationArticleListMatch, ctrl?: Control): Promise<DigitalPublicationArticleEntity[]>;
}
export { DigitalPublicationArticleEntity };
