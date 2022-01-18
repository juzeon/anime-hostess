<template>
  <div>
    <v-row class="d-flex justify-center align-center mx-4">
      <v-text-field class="flex-grow-1 d-inline-block" label="添加屏蔽弹幕" v-model="inputText"></v-text-field>
      <v-btn icon @click="addDanmaku">
        <v-icon>mdi-plus-circle-outline</v-icon>
      </v-btn>
    </v-row>
    <v-chip v-for="(danmaku,index) in forbidDanmakuList" :key="'forbidDanmaku-'+index"
            class="my-3 ml-3" close @click:close="removeDanmaku(danmaku)">
      {{ danmaku.toString().substring(1, danmaku.toString().length - 1) }}
    </v-chip>
  </div>
</template>

<script lang="ts">
import Vue from "vue"

export default Vue.extend({
  name: "ForbidDanmakuCard",
  data() {
    return {
      inputText: ''
    }
  },
  computed: {
    forbidDanmakuList: {
      set(value: RegExp[]) {
        this.$store.commit('setForbidDanmakuList', value)
      },
      get(): RegExp[] {
        return this.$store.state.forbidDanmakuList
      }
    }
  },
  methods: {
    addDanmaku() {
      for (let danmaku of this.forbidDanmakuList) {
        if (danmaku.toString() === '/' + this.inputText + '/') {
          this.$swal.error('已存在此屏蔽词')
          return
        }
      }
      let regexp: RegExp
      try {
        regexp = new RegExp(this.inputText)
      } catch (_e) {
        this.$swal.error('正则表达式不合法')
        return
      }
      let arr: RegExp[] = this.forbidDanmakuList.slice(0)
      arr.push(regexp)
      this.forbidDanmakuList = arr
      this.inputText = ''
      this.$emit('modify')
    },
    removeDanmaku(regExp: RegExp) {
      let arr: RegExp[] = this.forbidDanmakuList.slice(0)
      arr = arr.filter(value => value.toString() !== regExp.toString())
      this.forbidDanmakuList = arr
      this.$emit('modify')
    }
  }
})
</script>

<style scoped>

</style>
