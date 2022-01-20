<template>
  <div>
    <v-chip class="ml-2 mt-2" color="green" outlined v-for="(video,index) in videos"
            :key="'video-'+index" @click="navigateToWatch(video.hash)">{{ video.name }}
    </v-chip>
  </div>
</template>

<script lang="ts">
import Vue from "vue"

export default Vue.extend({
  name: "PlayList",
  props: ['seriesHash', 'videos'],
  methods: {
    navigateToWatch(hash: string) {
      this.$router.push({name: 'Watch', params: {seriesHash: this.seriesHash, hash}})
    },
  }
})

</script>

<style scoped>

</style>
