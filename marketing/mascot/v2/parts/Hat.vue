<script setup lang="ts">
import {computed} from "vue";

const props = defineProps({
  startX: Number,
  startY: Number,
});

const hatPath = computed(() => {
  let path = `M ${props.startX} ${props.startY} `;

  // lower part
  path += "q 50 10 95 2";
  path += "l -1 -20";
  path += "q -50 -10 -95 -1";
  path += "l 1 19";

  // top part
  path += "m 0 -8";
  path += "l -30 -13";
  path += "l 90 -35";
  return path;
});
</script>

<template>
  <path stroke="black" stroke-width="2" fill="gray" :d="hatPath" />
</template>

<style scoped lang="scss">

</style>