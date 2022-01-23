import Vue from 'vue'
import Vuex from 'vuex'
import {userConfig} from "@/plugins/config"
import {ICustomDanmakuOptions} from "@/types"

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        baseUrl: userConfig.baseUrl,
        userKey: userConfig.userKey,
        forbidDanmakuList: userConfig.forbidDanmakuList as RegExp[],
        customDanmakuOptions: userConfig.customDanmakuOptions
    },
    mutations: {
        setBaseUrl(state, value: string) {
            userConfig.baseUrl = value
            state.baseUrl = value
        },
        setUserKey(state, value: string) {
            userConfig.userKey = value
            state.userKey = value
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
