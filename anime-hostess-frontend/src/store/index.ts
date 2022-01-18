import Vue from 'vue'
import Vuex from 'vuex'
import {userConfig} from "@/plugins/config"
import {ICustomDanmakuOptions, IEngineModule} from "@/types"

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        baseUrl: userConfig.baseUrl,
        searchText: '',
        engineModule: undefined as undefined | IEngineModule,
        forbidDanmakuList: userConfig.forbidDanmakuList as RegExp[],
        customDanmakuOptions: userConfig.customDanmakuOptions
    },
    mutations: {
        setBaseUrl(state, value: string) {
            userConfig.baseUrl = value
            state.baseUrl = value
        },
        setSearchText(state, value: string) {
            state.searchText = value
        },
        setEngineModule(state, value: IEngineModule) {
            state.engineModule = value
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
