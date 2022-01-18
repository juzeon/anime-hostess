<template>
  <div class="d-flex justify-center align-center">
    <v-text-field class="d-inline-block flex-grow-1" :value="value"
                  @input="$emit('input', $event)"
                  @keydown.enter.prevent="$emit('search')"></v-text-field>
    <v-btn icon large @click="$emit('search')">
      <v-icon>mdi-magnify</v-icon>
    </v-btn>
  </div>
</template>

<script lang="ts">
import Vue from "vue"

export default Vue.extend({
  name: 'SearchBar',
  props: ['value'],
  data() {
    return {}
  },
})
</script>

<style scoped>

</style>
