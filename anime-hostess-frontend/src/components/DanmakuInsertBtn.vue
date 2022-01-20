<template>
  <div>
    <v-btn icon @click.stop="openDialog">
      <v-icon>mdi-plus</v-icon>
    </v-btn>
    <v-dialog v-model="dialog">
      <v-card :loading="loading">
        <v-card-title class="text-h5">
          {{ title }}
        </v-card-title>
        <v-card-text>
          <v-chip color="pink" outlined class="ml-3 mt-3" v-for="(danmakuEpisode,index) in danmakuEpisodeList"
                  :key="index" @click="insertDanmaku(danmakuEpisode)">
            {{ danmakuEpisode.title }}
          </v-chip>
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn text @click="dialog = false">
            关闭
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script lang="ts">
import Vue from "vue"
import {IBulletRaw, IDanmakuEpisode} from "@/types"
import {BulletOption} from "@nplayer/danmaku/dist/src/ts/danmaku/bullet"

export default Vue.extend({
  name: "DanmakuInsertBtn",
  props: ['seasonID', 'title'],
  data() {
    return {
      dialog: false,
      loading: true,
      danmakuEpisodeList: [] as IDanmakuEpisode[],
    }
  },
  methods: {
    openDialog() {
      this.dialog = true
      if (!this.danmakuEpisodeList.length) {
        this.getDanmakuEpisodeList()
      }
    },
    getDanmakuEpisodeList() {
      this.loading = true
      this.$axios.get('bullet/anime/'+this.seasonID).then(res => {
        this.danmakuEpisodeList = res.data.data
        this.loading = false
      })
    },
    insertDanmaku(danmakuEpisode: IDanmakuEpisode) {
      this.loading = true

      interface IPosMap {
        [key: number]: "scroll" | "top" | "bottom"
      }

      let posMap={
        1:'scroll',
        2:'scroll',
        3:'scroll',
        4:'bottom',
        5:'top',
        6:'top',
        7:'top',
        8:'scroll',
        9:'scroll',
      } as IPosMap
      this.$axios.get('bullet/bullet/'+danmakuEpisode.cid).then(res => {
        let arr = res.data.data as IBulletRaw[]
        let bulletsArr = [] as BulletOption[]
        for (let single of arr) {
          let time = single.time
          let type = posMap[single.type]
          let color = '#' + single.color.toString(16)
          let text = single.text
          bulletsArr.push({
            color,
            text,
            time,
            type
          })
        }
        bulletsArr.sort((a, b) => a.time - b.time)
        this.$emit('insertDanmaku', bulletsArr)
        this.loading = false
        this.dialog = false
      })
    }
  }
})

</script>

<style scoped>

</style>
