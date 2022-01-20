import {userConfig} from "@/plugins/config"
import store from '../store/index'
import {IVideoExploded} from "@/types"

export class Helper {
    getWebsocketBaseUrl(): string {
        if (userConfig.baseUrl.startsWith('https://')) {
            return 'wss://' + userConfig.baseUrl.substring(8, userConfig.baseUrl.length)
        } else {
            return 'ws://' + userConfig.baseUrl.substring(7, userConfig.baseUrl.length)
        }
    }

    getAnimeTokenFromUrl(url: string) {
        return /\/anime\/(.*)/.exec(url)![1]
    }

    getVideoExplodedFromUrl(url: string) {
        let arr = /\/anime\/(.*?)\/(\d+?)\/(\d+)/.exec(url)!
        return <IVideoExploded>{
            token: arr[1],
            playlist: parseInt(arr[2]),
            episode: parseInt(arr[3])
        }
    }
}
