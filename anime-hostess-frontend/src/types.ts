export interface IAnimeBySearch {
    category: string,
    cover_url: string,
    description: string,
    engine: string,
    score: number
    title: string
    url: string
}

export interface IEngineModule {
    anime: IAnimeItemOfEngineModule[],
    danmaku: IDanmakuItemOfEngineModule[],
    comic: any[],
    music: any[]
}

export interface IAnimeItemOfEngineModule {
    enable: boolean,
    module: string,
    name: string,
    notes: string,
    quality: number,
    type: string[]
}

export interface IDanmakuItemOfEngineModule {
    enable: boolean,
    module: string,
    name: string,
    notes: string,
    quality: number
}

export interface IAnimeDetail {
    title: string;
    category: string;
    cover_url: string;
    description: string;
    module: string;
    play_lists: IPlayListOfAnimeDetail[];
}

export interface IPlayListOfAnimeDetail {
    name: string;
    num: number;
    video_list: IVideoListOfAnimeDetail[];
}

export interface IVideoListOfAnimeDetail {
    info: string;
    name: string;
    player: string;
}

export interface IVideoExploded {
    token: string,
    playlist: number,
    episode: number
}

export interface IWatchInfo {
    format: string;
    lifetime: number;
    proxy_url: string;
    raw_url: string;
    size: number;
}

export interface IDanmakuSource {
    module: string;
    num: number;
    score: number;
    title: string;
    url: string;
}

export interface IDanmakuEpisode {
    data: string;
    name: string;
    url: string;
}

export interface ICustomDanmakuOptions {
    disable: boolean,
    blocked: Array<'scroll' | 'top' | 'bottom' | 'color'>,
    fontsize: number,
    fontsizeScale: number,
    opacity: number,
    speed: number,
    area: 0.25 | 0.5 | 0.75 | 1,
    unlimited: boolean,
    bottomUp: boolean,
}
