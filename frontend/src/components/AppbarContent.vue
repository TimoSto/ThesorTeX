<template>
  <v-app-bar
    :color="barColor"
    :density="dense ? 'compact' : 'default'"
    :style="fixed ? 'position: fixed' : ''"
  >
    <slot name="bar"></slot>
  </v-app-bar>
  <v-main :class="level === 1 ? 'top-level' : 'inner-level'">
    <slot name="content"></slot>
  </v-main>
</template>

<script setup lang="ts">

const props = defineProps({
  barColor: String,
  dense: Boolean,
  fixed: Boolean,
  level: Number,
})
</script>

<style scoped>

</style>
