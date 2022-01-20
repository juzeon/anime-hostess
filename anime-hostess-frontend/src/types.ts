export interface IAnimeBySearch {
    category: string,
    cover_url: string,
    description: string,
    engine: string,
    score: number
    title: string
    url: string
}

export interface IAnimeDetail {
    name: string;
    videos: IEpisode[];
    hash: string;
}

export interface IEpisode {
    seriesName: string;
    name: string;
    path: string;
    hash: string;
}

export interface ISeries {
    name: string;
    videos: IVideo[];
    hash: string;
}

export interface IVideo {
    seriesName: string;
    name: string;
    path: string;
    hash: string;
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

export interface IBulletSeries {
    seasonID: number;
    title: string;
}

export interface IDanmakuEpisode {
    cid: number
    title: string
}

export interface IBulletRaw {
    time?: number;
    type?: number;
    size?: number;
    color?: number;
    sentAt?: number;
    sender?: string;
    dmid?: number;
    level?: number;
    text?: string;
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
