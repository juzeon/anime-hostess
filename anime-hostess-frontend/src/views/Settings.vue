<template>
  <focus-area>
    <p class="text-h3">设定</p>
    <v-row class="d-flex align-center justify-center my-3">
      <div style="width: 60%">
        <v-text-field label="API地址" v-model="baseUrl" :rules="[validateBaseUrl]"></v-text-field>
        <div v-if="engineModuleInput" class="mt-4">
          <p class="text-h6">番剧源</p>
          <v-list-item v-for="(module,index) in engineModuleInput.anime" :key="'anime-'+index" two-line>
            <v-list-item-content>
              <v-list-item-title>{{ module.name }}</v-list-item-title>
              <v-list-item-subtitle>{{ module.notes }}</v-list-item-subtitle>
            </v-list-item-content>
            <v-list-item-action>
              <v-switch v-model="module.enable"></v-switch>
            </v-list-item-action>
          </v-list-item>
          <p class="text-h6 mt-3">弹幕源</p>
          <v-list-item v-for="(module,index) in engineModuleInput.danmaku" :key="'danmaku-'+index" two-line>
            <v-list-item-content>
              <v-list-item-title>{{ module.name }}</v-list-item-title>
              <v-list-item-subtitle>{{ module.notes }}</v-list-item-subtitle>
            </v-list-item-content>
            <v-list-item-action>
              <v-switch v-model="module.enable"></v-switch>
            </v-list-item-action>
          </v-list-item>
        </div>
        <v-btn block :disabled="!inputValid" @click="applySettings" :loading="btnLoading">更新</v-btn>
        <p class="text-h6 mt-3">屏蔽弹幕</p>
        <forbid-danmaku-card></forbid-danmaku-card>
      </div>
    </v-row>
  </focus-area>
</template>

<script lang="ts">
import Vue from 'vue'
import {IEngineModule} from "@/types"
import * as vuex from 'vuex'
import ForbidDanmakuCard from "@/components/ForbidDanmakuCard.vue"
import FocusArea from "@/components/FocusArea.vue"

export default Vue.extend({
  name: "Settings",
  components: {FocusArea, ForbidDanmakuCard},
  metaInfo: {
    title: '设定'
  },
  data() {
    return {
      inputValid: false,
      baseUrl: this.$store.state.baseUrl,
      engineModuleInput: undefined as IEngineModule | undefined,
      btnLoading: false
    }
  },
  computed: {
    ...vuex.mapState(['engineModule'])
  },
  watch: {
    engineModule: {
      immediate: true,
      handler() {
        if (this.engineModule) {
          this.engineModuleInput = JSON.parse(JSON.stringify(this.engineModule)) as IEngineModule
        }
      }
    }
  },
  methods: {
    applySettings() {
      this.btnLoading = true
      this.$store.commit('setBaseUrl', this.baseUrl)
      let promises = [] as Promise<any>[]
      if (this.engineModule) {
        let types = ['anime', 'danmaku']

        interface IModulePayload {
          module: string,
          enable: boolean
        }

        let payload = [] as IModulePayload[]
        for (let type of types) {
          for (let [index, module] of this.engineModuleInput![type as keyof IEngineModule].entries()) {
            if (module.enable !== this.engineModule[type][index].enable) {
              payload.push({
                module: module.module,
                enable: module.enable
              })
            }
          }
        }
        promises.push(this.$axios.post('system/modules', payload))
      }

      Promise.all(promises).then(() => {
        return this.$swal.success('设定已更新')
      }).then(() => {
        window.location.reload()
      })
    },
    validateBaseUrl() {
      if (!this.baseUrl.trim().length) {
        this.inputValid = false
        return '地址不可为空'
      }
      if (!this.isValidHttpUrl(this.baseUrl)) {
        this.inputValid = false
        return 'URL不合法'
      }
      if (!this.baseUrl.endsWith('/')) {
        this.inputValid = false
        return '地址需要以「/」结尾'
      }
      this.inputValid = true
      return true
    },
    isValidHttpUrl(string: string) {
      let url
      try {
        url = new URL(string)
      } catch (_) {
        return false
      }
      return url.protocol === "http:" || url.protocol === "https:"
    }
  }
})
</script>

<style scoped>

</style>
