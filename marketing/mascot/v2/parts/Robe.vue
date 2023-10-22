<script setup lang="ts">
import {computed} from "vue";

const props = defineProps({
  startX: Number,
  startY: Number,
});

const coloredTopShape = computed(() => {
  let path = `M ${props.startX} ${props.startY} `;

  path += "l 39 44 q 15 -10 46 -40 q 16 0 23 10 q -44 31 -66 59 q -33 -31 -63 -66 q 6 -5 21 -7";

  return path;
});

const robeLower = computed(() => {
  let path = `M ${props.startX - 18} ${props.startY + 6} `;

  path += "q -16 3 -20 75 q 90 14 179 2 q -3 -29 -33 -69 l -13 -9 l -95 -5 l -18 6";

  // left arm
  path += "m 11 79 q -1 -12 3 -39 q -4 28 -3 39";

  // right arm
  path += "m 118 1 q 1 -11 -11 -37 q 12 26 11 37";

  return path;
});
</script>

<template>
  <g>
    <path data-type="top-part" stroke="black" stroke-width="2" fill="#3A373B" :d="robeLower" />
    <path data-type="colored-part" stroke="black" stroke-width="2" fill="#008833" :d="coloredTopShape" />
  </g>
</template>

<style scoped lang="scss">

</style>