import Vue from 'vue'
import Vuex from 'vuex'
import {userConfig} from "@/plugins/config"
import {ICustomDanmakuOptions} from "@/types"

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        baseUrl: userConfig.baseUrl,
        searchText: '',
        forbidDanmakuList: userConfig.forbidDanmakuList as RegExp[],
        customDanmakuOptions: userConfig.customDanmakuOptions
    },
    mutations: {
        setBaseUrl(state, value: string) {
            userConfig.baseUrl = value
            state.baseUrl = value
        },
        setForbidDanmakuList(state, value: RegExp[]) {
            userConfig.forbidDanmakuList = value
            state.forbidDanmakuList = value
        },
        setCustomDanmakuOptions(state, value: ICustomDanmakuOptions) {
            userConfig.customDanmakuOptions = value
            state.customDanmakuOptions = value
        }
    },
    actions: {},
    modules: {}
})
