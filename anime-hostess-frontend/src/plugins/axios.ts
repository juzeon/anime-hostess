import axios from "axios"
import {userConfig} from "@/plugins/config"
import {swal} from "@/plugins/swal"

export let appAxios = axios.create({
    baseURL: userConfig.baseUrl,
    transformResponse: data => {
        try {
            return JSON.parse(data)
        } catch (e) {
            swal.error("从伺服器传回的资讯无效：" + data)
        }
    }
})
appAxios.interceptors.response.use(function (response) {
    // Do something with response data
    return response
}, function (error) {
    // Do something with response error
    swal.error('与伺服器连线时发生错误')
    return Promise.reject(error)
})
