import Vue from 'vue'
import VueRouter, {RouteConfig} from 'vue-router'
import Index from '@/views/Index.vue'
import Settings from "@/views/Settings.vue"
import Watch from "@/views/Watch.vue"

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
    {
        path: '/',
        name: 'Index',
        component: Index
    },
    {
        path: '/settings',
        name: 'Settings',
        component: Settings
    },
    {
        path: '/watch/:seriesHash/:hash',
        name: 'Watch',
        component: Watch,
        props: true
    }
]

const router = new VueRouter({
    mode: 'history',
    routes
})

export default router
