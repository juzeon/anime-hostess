<template>
  <focus-area>
    <v-progress-linear indeterminate v-show="!animeDetail || !watchInfo"></v-progress-linear>
    <p class="text-h6">
      <v-btn icon @click="navigateToAnime">
        <v-icon>mdi-arrow-left-circle-outline</v-icon>
      </v-btn>
      {{ (this.animeDetail) ? '第' + (parseInt(this.episode) + 1) + '集 - ' + this.animeDetail.title : '观看影片' }}
      &nbsp;
      {{ '（' + ((this.watchInfo) ? this.watchInfo.format : '加载中') + (this.useProxy ? ', proxy' : '') + '）' }}
    </p>
    <div v-show="this.watchInfo">
      <v-row>
        <v-col cols="10" offset="1">
          <NPlayer :set="setPlayer" :options="playerDefaultOptions"></NPlayer>
        </v-col>
      </v-row>
      <v-card class="mt-5 pa-4" v-if="playList">
        <p class="text-h6 ml-3 mt-3">选集</p>
        <play-list :name="playList.name" :video-list="playList.video_list"></play-list>
      </v-card>
      <v-card class="mt-5" :loading="danmakuLoading">
        <p class="text-h6 ml-3 mt-3">添加弹幕</p>
        <search-bar class="mx-3" v-model="danmakuSearchInput" @search="getDanmakuSourceList"></search-bar>
        <v-list-item v-for="(danmakuSource,index) in danmakuSourceList" :key="'danmakuSource-'+index" two-line>
          <v-list-item-content>
            <v-list-item-title>{{ danmakuSource.title }}</v-list-item-title>
            <v-list-item-subtitle>
              {{ $helper.translateDanmakuEngine(danmakuSource.module) }}（含{{ danmakuSource.num }}集）
            </v-list-item-subtitle>
          </v-list-item-content>
          <v-list-item-action>
            <danmaku-insert-btn :title="danmakuSource.title" :url="danmakuSource.url"
                                @insertDanmaku="insertDanmaku"></danmaku-insert-btn>
          </v-list-item-action>
        </v-list-item>
      </v-card>
      <v-card class="mt-5">
        <p class="text-h6 ml-3 mt-3">屏蔽弹幕</p>
        <forbid-danmaku-card @modify="resetDanmaku"></forbid-danmaku-card>
      </v-card>
    </div>
    <v-sheet class="py-16"></v-sheet>
  </focus-area>
</template>

<script lang="ts">
import Vue from "vue"
import FocusArea from "../components/FocusArea.vue"
import {IAnimeDetail, IDanmakuSource, IPlayListOfAnimeDetail, IWatchInfo} from "@/types"
import {MetaInfo} from "vue-meta"
import {EVENT, Player, PlayerOptions} from "nplayer"
import Hls from "hls.js"
import Danmaku from '@nplayer/danmaku'
import {BulletOption} from "@nplayer/danmaku/dist/src/ts/danmaku/bullet"
import SearchBar from "@/components/SearchBar.vue"
import DanmakuInsertBtn from "@/components/DanmakuInsertBtn.vue"
import axios from "axios"
import ForbidDanmakuCard from "@/components/ForbidDanmakuCard.vue"
import * as vuex from 'vuex'
import PlayList from "@/components/PlayList.vue"

export default Vue.extend({
  name: "Watch",
  components: {PlayList, ForbidDanmakuCard, DanmakuInsertBtn, SearchBar, FocusArea},
  props: ['token', 'playlist', 'episode'],
  metaInfo(): MetaInfo {
    return {
      title: (this.animeDetail) ? '第' + (parseInt(this.episode) + 1) + '集 - ' + this.animeDetail.title : '观看影片'
    }
  },
  computed: {
    routerProps(): string {
      return this.token + '-' + this.playlist + '-' + this.episode
    },
    playList(): IPlayListOfAnimeDetail | null {
      if (!this.animeDetail) {
        return null
      }
      return this.animeDetail.play_lists[parseInt(this.playlist)]
    },
    ...vuex.mapState(['forbidDanmakuList', 'customDanmakuOptions'])
  },
  data() {
    return {
      watchInfo: undefined as IWatchInfo | undefined,
      animeDetail: undefined as IAnimeDetail | undefined,
      player: undefined as Player | undefined,
      playerDefaultOptions: {
        controls: [
          ['play', 'volume', 'time', 'spacer', 'airplay', 'danmaku-settings', 'settings', 'web-fullscreen', 'fullscreen'],
          ['progress']
        ],
        plugins: [
          new Danmaku({
            autoInsert: false,
            ...this.$store.state.customDanmakuOptions
          })
        ]
      } as PlayerOptions,
      danmakuSourceList: [] as IDanmakuSource[],
      danmakuListUnfiltered: [] as BulletOption[],
      danmakuListFiltered: [] as BulletOption[],
      danmakuSearchInput: '',
      danmakuLoading: false,
      useProxy: false,
      historyTimeLogger: undefined as number | undefined,
      firstPlay: true
    }
  },
  watch: {
    routerProps: {
      immediate: true,
      handler() {
        this.initWatch()
      }
    }
  },
  deactivated() {
    clearInterval(this.historyTimeLogger)
    this.$destroy()
  },
  methods: {
    navigateToAnime() {
      this.$router.push({name: 'Anime', params: {token: this.token}})
    },
    setPlayer(player: Player) {
      this.player = player
    },
    loadVideo() {
      console.log('load video type: ' + this.watchInfo?.format + ', use proxy: ' + this.useProxy)
      let videoUrl = this.useProxy ? this.watchInfo!.proxy_url : this.watchInfo!.raw_url
      if (this.watchInfo?.format === 'hls') {
        let hls = new Hls()
        hls.attachMedia(this.player!.video)
        hls.on(Hls.Events.MEDIA_ATTACHED, () => {
          hls.loadSource(videoUrl)
        })
      } else {
        this.player?.updateOptions({src: videoUrl})
      }
    },
    initWatch() {
      this.animeDetail = undefined
      this.watchInfo = undefined
      this.danmakuListFiltered = []
      this.danmakuListUnfiltered = []
      clearInterval(this.historyTimeLogger)
      this.firstPlay = true
      this.player?.danmaku.resetItems([])
      this.useProxy = false

      this.$axios.get("anime/" + this.token).then(res => {
        this.animeDetail = res.data
        this.danmakuSearchInput = this.animeDetail!.title
        this.getDanmakuSourceList()
      })
      this.$axios.get('anime/' + this.token + '/' + this.playlist + '/' + this.episode).then(res => {
        this.watchInfo = res.data
        this.player?.on(EVENT.ERROR, () => {
          this.loadVideo()
        })
        this.player?.on(this.player?.danmaku.EVENT.DANMAKU_UPDATE_OPTIONS, () => {
          this.$store.commit('setCustomDanmakuOptions', this.player?.danmaku.opts)
        })
        this.player?.on(EVENT.PLAY, () => {
          if (!this.firstPlay) {
            return
          }
          let historyTimeKey = this.token + '_' + this.episode
          let historyTime = localStorage.getItem(historyTimeKey)
          console.log('history time: ' + historyTime)
          if (historyTime) {
            this.player?.seek(parseInt(historyTime))
          }
          this.historyTimeLogger = setInterval(() => {
            // console.log('set time: ' + this.player?.currentTime)
            localStorage.setItem(historyTimeKey, this.player!.currentTime + '')
          }, 1000)
          this.firstPlay = false
        })
        return axios.head(this.watchInfo!.raw_url)
      }).then(res => {
        this.loadVideo()
      }).catch(err => {
        this.useProxy = true
        this.loadVideo()
      })
    },
    getDanmakuSourceList() {
      this.danmakuLoading = true
      this.$axios.get('danmaku/search/' + encodeURI(this.danmakuSearchInput)).then(res => {
        this.danmakuSourceList = res.data
        this.danmakuLoading = false
      })
    },
    insertDanmaku(danmakuList: BulletOption[]) {
      this.danmakuListUnfiltered = [...this.danmakuListUnfiltered, ...danmakuList].sort((a, b) => a.time - b.time)
      this.resetDanmaku()
    },
    resetDanmaku() {
      let danmakuList = this.danmakuListUnfiltered.filter(value => {
        for (let regexp of this.forbidDanmakuList) {
          if (regexp.test(value.text)) {
            return false
          }
        }
        return true
      })
      this.danmakuListFiltered = danmakuList
      this.player?.danmaku.resetItems(danmakuList)
    }
  }
})
</script>

<style scoped>

</style>
