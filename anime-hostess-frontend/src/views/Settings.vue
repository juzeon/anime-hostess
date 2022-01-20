<template>
  <focus-area>
    <p class="text-h3">设定</p>
    <v-row class="d-flex align-center justify-center my-3">
      <div style="width: 60%">
        <v-text-field label="API地址" v-model="baseUrl" :rules="[validateBaseUrl]"></v-text-field>
        <v-btn block :disabled="!inputValid" @click="applySettings" :loading="btnLoading">更新</v-btn>
        <p class="text-h6 mt-3">屏蔽弹幕</p>
        <forbid-danmaku-card></forbid-danmaku-card>
      </div>
    </v-row>
  </focus-area>
</template>

<script lang="ts">
import Vue from 'vue'
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
      btnLoading: false
    }
  },
  methods: {
    applySettings() {
      this.btnLoading = true
      this.$store.commit('setBaseUrl', this.baseUrl)
      this.$swal.success('设定已更新').then(() => {
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
