<script setup lang="ts">
import {computed} from "vue";

const props = defineProps({
  startX: Number,
  startY: Number,
  right: Boolean,
  eyePosition: Number, // -1: left, 0: middle, 1: right
});

const shape = computed(() => {
  let path = `M ${props.startX} ${props.startY + (props.right ? 2 : 0)}`;

  // bottom
  path += `q ${props.right ? "15" : "13"} -5 25 ${props.right ? "-2" : "2"}`;

  // curve
  path += `c ${props.right ? "-2 -37 -35 -45 -25" : "10 -37 -21 -51 -25"} ${props.right ? "2" : "-2"}`;

  return path;
});

const pupil = computed(() => {
  // default: left middle
  let offsetX = 12;
  let offsetY = -22;
  if (props.right) {
    offsetX = 7;
    offsetY = -21;
  }
  let path = `M ${props.startX + offsetX} ${props.startY + offsetY} `;

  path += "q 7 -1 8 6 c 1 10 -9 11 -11 3 c 3 0 6 -4 3 -9";

  return path;
});

</script>

<template>
  <g>
    <path stroke="black" stroke-width="2" fill="white" :d="shape" />
    <path stroke="black" stroke-width="2" fill="black" :d="pupil" />

  </g>
</template>

<style scoped lang="scss">

</style>