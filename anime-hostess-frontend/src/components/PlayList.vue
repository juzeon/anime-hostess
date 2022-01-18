<template>
  <div>
    <p class="text-h6">{{ name }}</p>
    <v-chip class="ml-2 mt-2" color="green" outlined v-for="(videoList,index) in videoList"
            :key="'videolist-'+name+'-'+index" @click="navigateToWatch(videoList)">{{ videoList.name }}
    </v-chip>
  </div>
</template>

<script lang="ts">
import Vue from "vue"
import {IVideoListOfAnimeDetail} from "@/types"

export default Vue.extend({
  name: "PlayList",
  props: ['name', 'videoList'],
  methods: {
    navigateToWatch(videoList: IVideoListOfAnimeDetail) {
      let videoExploded = this.$helper.getVideoExplodedFromUrl(videoList.info)
      this.$router.push({
        name: 'Watch',
        params: {token: videoExploded.token, playlist: videoExploded.playlist + '', episode: videoExploded.episode + ''}
      })
    },
  }
})

</script>

<style scoped>

</style>
