<template>
  <focus-area>
    <v-card :loading="loading">
      <v-row>
        <v-col cols="12">
          <p class="text-h5 ml-3 mt-3">搜寻「{{ searchText }}」的结果</p>
        </v-col>
        <v-col cols="10" offset="1" class="">
          <search-bar v-model="inputText" @search="applySearchText"></search-bar>
        </v-col>
        <v-col cols="12">
          <anime-item class="ma-4 my-4" v-for="(item,index) in animeArr" :key="'animeArr-'+index"
                      :cover="item.cover_url" :category="item.category"
                      :description="item.description" :engine="item.engine"
                      :title="item.title" :url="item.url" :score="item.score"
                      @click.native="navigateToAnime(item.url)"></anime-item>
        </v-col>
      </v-row>
    </v-card>
  </focus-area>
</template>

<script lang="ts">
import Vue from "vue"
import {MetaInfo} from "vue-meta"
import SearchBar from "@/components/SearchBar.vue"
import FocusArea from "@/components/FocusArea.vue"
import * as vuex from 'vuex'
import {IAnimeBySearch} from "@/types"
import AnimeItem from "@/components/AnimeItem.vue"

export default Vue.extend({
  name: "Search",
  components: {AnimeItem, FocusArea, SearchBar},
  props: ['searchTextPassed'],
  metaInfo(): MetaInfo {
    return {
      title: this.searchText
    }
  },
  computed: {
    ...vuex.mapState(['searchText'])
  },
  watch: {
    searchTextPassed: {
      handler() {
        this.inputText = this.searchTextPassed
        this.$store.commit('setSearchText', this.searchTextPassed)
        this.startSearch()
      },
      immediate: true
    }
  },
  data() {
    return {
      inputText: '',
      ws: undefined as undefined | WebSocket,
      animeArr: [] as IAnimeBySearch[],
      loading: false
    }
  },
  deactivated() {
    if (this.ws !== undefined) {
      this.ws.close()
    }
  },
  methods: {
    navigateToAnime(url: string) {
      this.$router.push({name: 'Anime', params: {token: this.$helper.getAnimeTokenFromUrl(url)}})
    },
    applySearchText() {
      this.$store.commit('setSearchText', this.inputText)
      this.startSearch()
    },
    startSearch() {
      this.loading = true
      console.log('start searching by ' + this.searchText)
      this.animeArr = []
      if (this.ws !== undefined) {
        this.ws.close()
      }
      this.ws = new WebSocket(this.$helper.getWebsocketBaseUrl() + '/anime/search')
      this.ws.onopen = () => {
        console.log('websocket open')
        this.ws?.send(this.searchText)
      }
      this.ws.onmessage = (event) => {
        let data: IAnimeBySearch = JSON.parse(event.data)
        console.log('Received: ' + JSON.stringify(data))
        this.animeArr.push(data)
      }
      this.ws.onclose = () => {
        console.log('websocket close')
        this.ws = undefined
        this.loading = false
      }
    }
  }
})
</script>

<style scoped>

</style>
