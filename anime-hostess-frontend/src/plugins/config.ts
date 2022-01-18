import {ICustomDanmakuOptions} from "@/types"

export let userConfig = {
    get baseUrl() {
        return localStorage.getItem("animeBaseUrl") || ''
    },
    set baseUrl(value: string) {
        localStorage.setItem("animeBaseUrl", value)
    },
    get forbidDanmakuList() {
        let storage = localStorage.getItem("animeForbidDanmakuList")
        let results: RegExp[] = []
        if (storage) {
            let arr = JSON.parse(storage)
            for (let str of arr) {
                results.push(new RegExp(str))
            }
        }
        return results
    },
    set forbidDanmakuList(value: RegExp[]) {
        let arr: string[] = []
        for (let regexp of value) {
            arr.push(regexp.toString().substring(1, regexp.toString().length - 1))
        }
        localStorage.setItem('animeForbidDanmakuList', JSON.stringify(arr))
    },
    get customDanmakuOptions(): ICustomDanmakuOptions {
        let storage = localStorage.getItem("animeCustomDanmakuOptions")
        if (storage) {
            return JSON.parse(storage)
        }
        return {
            area: 0.75,
            blocked: [],
            disable: false,
            bottomUp: false,
            fontsize: 24,
            fontsizeScale: 1,
            opacity: 1,
            speed: 1,
            unlimited: false
        }
    },
    set customDanmakuOptions(value: ICustomDanmakuOptions) {
        localStorage.setItem('animeCustomDanmakuOptions', JSON.stringify(value))
    }
}
