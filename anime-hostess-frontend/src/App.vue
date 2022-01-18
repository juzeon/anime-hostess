<template>
  <v-app>
    <v-app-bar app>
      <div class="d-flex align-center" style="cursor: pointer" @click="$router.push({name:'Index'})">
        <v-img
            class="shrink mr-2"
            src="https://img14.360buyimg.com/ddimg/jfs/t1/208970/36/2976/2479/6156a3fcEe432683d/f6cf924b25e69f6b.png"
            width="40"
        />
        <span class="text-h6 ml-2">Anime</span>
      </div>
      <v-spacer></v-spacer>
      <v-btn icon @click="navigateToSettings"><v-icon>mdi-cog</v-icon></v-btn>
    </v-app-bar>
    <v-main>
      <keep-alive>
        <router-view/>
      </keep-alive>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'App',
  metaInfo: {
    titleTemplate: '%s | Anime'
  },
  data() {
    return {}
  },
  mounted() {
    if (!this.$store.state.baseUrl.length) {
      this.navigateToSettings()
      return
    }
    this.$axios.get('system/modules').then(res => {
      this.$store.commit('setEngineModule', res.data)
    })
  },
  methods:{
    navigateToSettings(){
      this.$router.push({name: 'Settings'})
    }
  }
})
</script>
