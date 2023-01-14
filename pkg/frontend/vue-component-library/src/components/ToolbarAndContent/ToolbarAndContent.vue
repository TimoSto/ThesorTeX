<template>
  <v-toolbar
      v-if="!props.hideBar"
      :color="props.barColor"
      density="compact"
      elevation="1"
      style="z-index: 1000"

  >
    <slot name="bar" />
  </v-toolbar>
  <v-sheet
      class="content"
  >
    <slot name="content" />
  </v-sheet>
</template>

<script setup lang="ts">
const props = defineProps({
  hideBar: {
    type: Boolean,
    default: false
  },
  barColor: {
    type: String,
    default: "background"
  }
})
</script>

<style scoped>
.content {
  padding: 0;
  margin: 0;
  width: 100%;
  overflow-y: auto;
}
</style>