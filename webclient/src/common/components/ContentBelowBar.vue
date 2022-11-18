<template>
  <v-sheet
      style="overflow-y: auto" :id="scrollId"
      :style="`
        max-height: calc(100vh - ${barHeight}px);
        margin-top: ${marginTop}px;
        padding: ${paddingY}px ${paddingX}px;
        `"
  >
    <div>
      <slot></slot>
    </div>
  </v-sheet>
</template>

<script>
export default {
  name: "ContentBelowBar",
  props: [
      'barHeight',
      'marginTop',
      'paddingX',
      'paddingY',
      'scrollId'
  ]
}
</script>

<style scoped>

</style>