<template>
  <v-card :hover="!disableHover">
    <v-row class="d-flex align-center">
      <v-col cols="2">
        <v-img :src="cover" class="ml-3"></v-img>
      </v-col>
      <v-col cols="10">
        <div class="ma-3">
          <p class="text-h5">{{ title }}</p>
          <v-chip color="blue" outlined>{{ $helper.translateAnimeEngine(engine) }}</v-chip>
          <p class="body-1 mt-2">{{ description }}</p>
        </div>
      </v-col>
    </v-row>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue"

export default Vue.extend({
  name: "AnimeItem",
  props: ['cover', 'title', 'category', 'description', 'url', 'engine', 'score', 'disableHover']
})
</script>

<style scoped>

</style>
