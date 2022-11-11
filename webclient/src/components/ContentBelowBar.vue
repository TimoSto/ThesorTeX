<template>
  <v-sheet
      style="overflow-y: auto;"
      :style="`
        height: calc(100% - ${barHeight}px);
        margin-top: ${barHeight}px;
        padding: ${paddingY}px ${paddingX}px;
        `"
  >
    <slot></slot>
  </v-sheet>
</template>

<script>
export default {
  name: "ContentBelowBar",
  props: [
      'barHeight',
      'paddingX',
      'paddingY'
  ]
}
</script>

<style scoped>

</style>