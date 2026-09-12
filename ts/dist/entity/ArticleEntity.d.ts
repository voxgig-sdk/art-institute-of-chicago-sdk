import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { Article, ArticleLoadMatch, ArticleListMatch } from '../ArtInstituteOfChicagoTypes';
declare class ArticleEntity extends ArtInstituteOfChicagoEntityBase<Article> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: ArticleEntity): ArticleEntity;
    load(this: any, reqmatch?: ArticleLoadMatch, ctrl?: Control): Promise<ArticleEntity>;
    list(this: any, reqmatch?: ArticleListMatch, ctrl?: Control): Promise<ArticleEntity[]>;
}
export { ArticleEntity };
