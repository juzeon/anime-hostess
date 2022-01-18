<template>
  <v-row class="mt-4">
    <v-col class="col-sm-12 col-md-10 offset-sm-0 offset-md-1">
      <slot></slot>
    </v-col>
  </v-row>
</template>

<script>
import Vue from "vue"

export default Vue.extend({
  name: "FocusArea"
})

</script>

<style scoped>

</style>
