<script setup lang="ts">
import {computed} from "vue";

const props = defineProps({
  startX: Number,
  startY: Number,
});

const coloredTopShape = computed(() => {
  let path = `M ${props.startX} ${props.startY - 0} `;

  path += "l 39 44 q 15 -10 46 -40 q 16 0 23 10 q -44 31 -66 59 q -33 -31 -63 -66 q 6 -5 21 -7";

  return path;
});

const robeLower = computed(() => {
  let path = `M ${props.startX - 18} ${props.startY + 6} `;

  path += "q -16 3 -22 97 q 92 -8 184 2 q -6 -51 -36 -91 l -13 -9 l -95 -5 l -18 6";

  // left arm
  path += "m 11 95 q -1 -28 3 -55 q -4 28 -3 55";

  // right arm
  path += "m 120 1 q -1 -27 -13 -53 q 12 26 13 53";

  return path;
});
</script>

<template>
  <g>
    <path data-type="top-part" stroke="black" stroke-width="2" fill="#3A373B" :d="robeLower" />
    <path data-type="top-part" stroke="black" stroke-width="2" fill="#008833" :d="coloredTopShape" />
  </g>
</template>

<style scoped lang="scss">

</style>