<template>
  <focus-area>
    <v-progress-linear indeterminate v-show="!animeDetail"></v-progress-linear>
    <v-btn icon class="d-block my-4" @click="navigateBack">
      <v-icon>mdi-arrow-left-circle-outline</v-icon>
    </v-btn>
    <div v-if="animeDetail">
      <anime-item :description="animeDetail.description" :category="animeDetail.category"
                  :cover="animeDetail.cover_url" :engine="animeDetail.module"
                  :title="animeDetail.title" :disable-hover="true"></anime-item>
      <v-divider></v-divider>
      <v-card class="mt-5" elevation="0" v-for="(playList,index) in animeDetail.play_lists" :key="'playlist-'+index">
        <play-list :name="playList.name" :video-list="playList.video_list"></play-list>
      </v-card>
    </div>
  </focus-area>
</template>

<script lang="ts">
import Vue from "vue"
import FocusArea from "@/components/FocusArea.vue"
import {IAnimeDetail, IVideoListOfAnimeDetail} from "@/types"
import AnimeItem from "@/components/AnimeItem.vue"
import {MetaInfo} from "vue-meta"
import PlayList from "@/components/PlayList.vue"

export default Vue.extend({
  name: "Anime",
  components: {PlayList, AnimeItem, FocusArea},
  props: ['token'],
  metaInfo(): MetaInfo {
    return {
      title: this.animeDetail?.title || '动画详情'
    }
  },
  data() {
    return {
      animeDetail: undefined as IAnimeDetail | undefined
    }
  },
  watch: {
    token: {
      immediate: true,
      handler() {
        this.getAnimeDetail()
      }
    }
  },
  methods: {
    navigateBack() {
      this.$router.go(-1)
    },
    getAnimeDetail() {
      this.animeDetail = undefined
      this.$axios.get("anime/" + this.token).then(res => {
        this.animeDetail = res.data
        this.animeDetail!.cover_url = this.animeDetail!.cover_url.replace('http:http', 'http')
      }).catch(err => {
        this.$router.go(-1)
      })
    }
  }
})

</script>

<style scoped>

</style>
