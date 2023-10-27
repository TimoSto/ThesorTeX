<script setup lang="ts">
import {computed} from "vue";

const props = defineProps({
  startX: Number,
  startY: Number,
  right: Boolean,
});

const coloredTopShape = computed(() => {
  let path = `M ${props.startX} ${props.startY} `;

  path += `l ${props.right ? "45" : "39"} 44 q 15 -10 ${props.right ? "40" : "46"} -40 q 16 0 23 10 q ${props.right ? "-38" : "-44"} 31 ${props.right ? "-60" : "-66"} 59 q -33 -31 ${props.right ? "-69" : "-63"} -66 q 6 -5 21 -7`;

  return path;
});

const robeLower = computed(() => {
  let path = `M ${props.startX - 18} ${props.startY + 6} `;

  path += "q -16 3 -20 75 q 90 14 179 2 q -3 -29 -33 -69 l -13 -9 l -95 -5 l -18 6";

  // left arm
  path += props.right ? "m 11 79 q -1 -12 6 -39 q -7 28 -6 39" : "m 11 79 q -1 -12 4 -37 q -5 28 -4 37";

  // right arm
  path += props.right ? "m 118 1 q 1 -11 -9 -38 q 10 27 9 39" : "m 118 1 q 1 -11 -13 -41 q 13 27 13 41";

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