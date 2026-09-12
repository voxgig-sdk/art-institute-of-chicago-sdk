import { ArtInstituteOfChicagoEntityBase } from '../ArtInstituteOfChicagoEntityBase';
import type { ArtInstituteOfChicagoSDK } from '../ArtInstituteOfChicagoSDK';
import type { Control } from '../types';
import type { Video, VideoLoadMatch, VideoListMatch } from '../ArtInstituteOfChicagoTypes';
declare class VideoEntity extends ArtInstituteOfChicagoEntityBase<Video> {
    constructor(client: ArtInstituteOfChicagoSDK, entopts: any);
    make(this: VideoEntity): VideoEntity;
    load(this: any, reqmatch?: VideoLoadMatch, ctrl?: Control): Promise<VideoEntity>;
    list(this: any, reqmatch?: VideoListMatch, ctrl?: Control): Promise<VideoEntity[]>;
}
export { VideoEntity };
