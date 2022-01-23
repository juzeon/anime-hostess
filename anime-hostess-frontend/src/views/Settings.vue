<template>
  <focus-area>
    <p class="text-h3">设定</p>
    <v-row class="d-flex align-center justify-center my-3">
      <div style="width: 60%">
        <v-text-field label="API地址" v-model="baseUrl" :rules="[validateBaseUrl]"></v-text-field>
        <div class="d-flex align-baseline">
          <v-text-field label="使用者识别码" v-model="userKey"></v-text-field>
          <v-btn class="ml-3" @click="generateUserKey" :loading="genUserKeyLoading">获取新的</v-btn>
        </div>
        <p class="caption">使用者识别码用于在伺服器上储存浏览记录、播放记录等资讯，以便多装置同步。请妥善保存以免记录丢失。</p>
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
      userKey: this.$store.state.userKey,
      btnLoading: false,
      genUserKeyLoading: false,
    }
  },
  methods: {
    applySettings() {
      this.btnLoading = true
      this.$store.commit('setBaseUrl', this.baseUrl)
      this.$store.commit('setUserKey', this.userKey)
      this.$axios.defaults.baseURL = this.baseUrl
      this.$axios.defaults.headers['Authorization'] = this.userKey
      this.$swal.success('设定已更新').then(() => {
        this.btnLoading = false
      })
    },
    generateUserKey() {
      this.$swal.confirm('是否获取新的识别码并覆盖旧的？').then(res => {
        if (res.isConfirmed) {
          this.genUserKeyLoading = true
          this.$axios.post('user/generate').then(res => {
            this.userKey = res.data.data.token
            this.genUserKeyLoading = false
          })
        }
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
