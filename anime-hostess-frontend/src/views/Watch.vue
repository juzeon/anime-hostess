<template>
  <focus-area>
    <v-progress-linear indeterminate v-show="!episodeDetail"></v-progress-linear>
    <p class="text-h6">
      <v-btn icon @click="navigateToIndex">
        <v-icon>mdi-arrow-left-circle-outline</v-icon>
      </v-btn>
      {{ (episodeDetail) ? episodeDetail.name + ' - ' + episodeDetail.seriesName : '观看影片' }}
    </p>
    <div v-show="episodeDetail">
      <v-row>
        <v-col cols="10" offset="1">
          <NPlayer :set="setPlayer" :options="playerDefaultOptions"></NPlayer>
        </v-col>
      </v-row>
      <v-card class="mt-5 pa-4">
        <p class="text-h6 ml-3 mt-3">选集</p>
        <play-list v-if="animeDetail" :series-hash="animeDetail.hash" :videos="animeDetail.videos"></play-list>
      </v-card>
      <v-card class="mt-5" :loading="danmakuLoading">
        <p class="text-h6 ml-3 mt-3">添加弹幕</p>
        <search-bar class="mx-3" v-model="danmakuSearchInput" @search="userEmitDanmakuSearch"></search-bar>
        <v-progress-linear indeterminate v-show="danmakuSearchInputLoading"></v-progress-linear>
        <v-list-item v-for="(danmakuSource,index) in danmakuSourceList" :key="'danmakuSource-'+index" two-line>
          <v-list-item-content>
            <v-list-item-title>{{ danmakuSource.title }}</v-list-item-title>
          </v-list-item-content>
          <v-list-item-action>
            <danmaku-insert-btn :title="danmakuSource.title" :seasonID="danmakuSource.seasonID"
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
import {IAnimeDetail, IBulletSeries, IEpisode} from "@/types"
import {MetaInfo} from "vue-meta"
import {EVENT, Player, PlayerOptions} from "nplayer"
import Danmaku from '@nplayer/danmaku'
import {BulletOption} from "@nplayer/danmaku/dist/src/ts/danmaku/bullet"
import SearchBar from "@/components/SearchBar.vue"
import DanmakuInsertBtn from "@/components/DanmakuInsertBtn.vue"
import ForbidDanmakuCard from "@/components/ForbidDanmakuCard.vue"
import * as vuex from 'vuex'
import PlayList from "@/components/PlayList.vue"
import * as QueryString from "querystring"

export default Vue.extend({
  name: "Watch",
  components: {PlayList, ForbidDanmakuCard, DanmakuInsertBtn, SearchBar, FocusArea},
  props: ['hash', 'seriesHash'],
  metaInfo(): MetaInfo {
    return {
      title: (this.episodeDetail) ? this.episodeDetail.name + ' - ' + this.episodeDetail.seriesName : '观看影片'
    }
  },
  computed: {
    routerProps(): string {
      return this.hash + '-' + this.seriesHash
    },
    ...vuex.mapState(['forbidDanmakuList', 'customDanmakuOptions', 'baseUrl'])
  },
  data() {
    return {
      initWatchPending: false,
      episodeDetail: undefined as IEpisode | undefined,
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
      danmakuSourceList: [] as IBulletSeries[],
      danmakuListUnfiltered: [] as BulletOption[],
      danmakuListFiltered: [] as BulletOption[],
      danmakuSearchInput: '',
      danmakuSearchInputLoading: false,
      danmakuLoading: false,
      historyTimeLogger: undefined as number | undefined,
      firstPlay: true
    }
  },
  watch: {
    routerProps: {
      immediate: true,
      handler() {
        if (!this.player) {
          this.initWatchPending = true
        } else {
          this.initWatch()
        }
      }
    }
  },
  deactivated() {
    clearInterval(this.historyTimeLogger)
    this.$destroy()
  },
  methods: {
    navigateToIndex() {
      this.$router.push({name: 'Index'})
    },
    setPlayer(player: Player) {
      this.player = player
      if (this.initWatchPending) {
        this.initWatch()
        this.initWatchPending = false
      }
    },
    loadVideo() {
      let videoUrl = this.baseUrl + 'video/stream/' + this.episodeDetail?.hash
      this.player!.updateOptions({src: videoUrl})
    },
    initWatch() {
      clearInterval(this.historyTimeLogger)
      this.animeDetail = undefined
      this.episodeDetail = undefined
      this.danmakuListFiltered = []
      this.danmakuListUnfiltered = []
      this.firstPlay = true
      this.player!.danmaku.resetItems([])
      this.danmakuSearchInputLoading = true
      this.$axios.get('user/searchText/' + this.seriesHash).then(res => {
        this.danmakuSearchInput = res.data.data
        this.danmakuSearchInputLoading = false
      })
      this.$axios.get("video/detail/" + this.seriesHash).then(res => {
        this.animeDetail = res.data.data
        this.getDanmakuSourceList()
        this.episodeDetail = this.animeDetail!.videos.find(video => video.hash === this.hash)
        this.loadVideo()
        this.player!.on(EVENT.ERROR, () => {
          this.loadVideo()
        })
        this.player!.on(this.player!.danmaku.EVENT.DANMAKU_UPDATE_OPTIONS, () => {
          this.$store.commit('setCustomDanmakuOptions', this.player!.danmaku.opts)
        })
        this.player!.on(EVENT.PLAY, () => {
          if (!this.firstPlay) {
            return
          }
          this.firstPlay = false
          this.$axios.get('user/progress/' + this.hash).then(res => {
            let progress = parseInt(res.data.data)
            console.log('progress: ' + progress)
            if (progress) {
              this.player!.seek(progress)
            }
            this.historyTimeLogger = setInterval(() => {
              this.$axios.post('user/progress', QueryString.stringify({
                time: this.player!.currentTime,
                hash: this.hash,
              }))
            }, 5000)
          })
        })
      })
    },
    userEmitDanmakuSearch() {
      if (this.$store.state.userKey.length) {
        this.$axios.post('user/searchText', QueryString.stringify({
          hash: this.seriesHash,
          searchText: this.danmakuSearchInput,
        }))
      }
      this.getDanmakuSourceList()
    },
    getDanmakuSourceList() {
      this.danmakuLoading = true
      this.$axios.get('bullet/search/' + encodeURI(this.danmakuSearchInput)).then(res => {
        this.danmakuSourceList = res.data.data
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
