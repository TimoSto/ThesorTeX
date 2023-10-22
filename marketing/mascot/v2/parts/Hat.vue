<script setup lang="ts">
import {computed} from "vue";

const props = defineProps({
  startX: Number,
  startY: Number,
});

const hatPath = computed(() => {
  let path = `M ${props.startX} ${props.startY + 1} `;

  // lower part
  path += "q 50 10 75 1";
  path += "l -1 -18";
  path += "q -27 -15 -75 -3";
  path += "l 1 20";

  return path;
});

const hatPathTop = computed(() => {
  let path = `M ${props.startX} ${props.startY - 8} `;

  path += "l -28 -10";
  path += "q 25 -15 70 -32";
  path += "q 25 9 63 35";
  path += "l -30 7";
  return path;
});
</script>

<template>
  <path stroke="black" stroke-width="2" fill="#3A373B" :d="hatPathTop" />
  <path stroke="black" stroke-width="2" fill="#3A373B" :d="hatPath" />
</template>

<style scoped lang="scss">

</style>