<template>
  <focus-area>
    <v-card :loading="loading">
      <v-card-title class="d-flex align-center">
        <p>伺服器上的影片资料夹</p>
        <v-spacer></v-spacer>
        <p><v-btn icon @click="getSeriesList"><v-icon>mdi-refresh</v-icon></v-btn></p>
      </v-card-title>
      <search-bar class="mx-7" v-model="searchText"></search-bar>
      <v-expansion-panels>
        <v-expansion-panel v-for="(item,index) in filteredSeries" :key="'series-'+index">
          <v-expansion-panel-header>
            {{ item.name }}
          </v-expansion-panel-header>
          <v-expansion-panel-content>
            <v-list>
              <v-list-item v-for="(episode,index2) in item.videos" :key="'episode-'+index+'-'+index2">
                <v-list-item-title>
                  <v-btn text @click="navigateToWatch(item.hash,episode.hash)">{{ episode.name }}</v-btn>
                </v-list-item-title>
              </v-list-item>
            </v-list>
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-card>
  </focus-area>
</template>

<script lang="ts">
import Vue from 'vue'
import FocusArea from "@/components/FocusArea.vue"
import {ISeries} from "@/types"
import SearchBar from "@/components/SearchBar.vue"

export default Vue.extend({
  name: 'Index',
  components: {SearchBar, FocusArea},
  metaInfo: {
    title: '主页'
  },
  activated() {
    if (!this.$store.state.baseUrl.length) {
      return
    }
    if (!this.pageInit) {
      this.pageInit = true
      this.initPage()
    }
  },
  data() {
    return {
      pageInit: false,
      searchText: '',
      series: [] as ISeries[],
      loading:false,
    }
  },
  computed: {
    filteredSeries(): ISeries[] {
      if (this.searchText.length === 0) {
        return this.series
      }
      return this.series.filter(item => {
        return item.name.includes(this.searchText)
      })
    }
  },
  methods: {
    initPage() {
      this.getSeriesList()
    },
    navigateToWatch(seriesHash: string, hash: string) {
      this.$router.push({name: 'Watch', params: {seriesHash, hash}})
    },
    getSeriesList() {
      this.loading=true
      this.$axios.get('video/list').then(res => {
        this.series = res.data.data
        this.loading=false
      })
    }
  }
})
</script>
