<script setup>
import {computed} from "vue";

const props = defineProps({
  startX: Number,
  startY: Number,
});

const shape = computed(() => {
  let path = `M ${props.startX} ${props.startY} `;

  // forehead
  path += "q 50 -12 90 0";

  // right side
  path += "q 15 5 20 55";

  // right beard
  path += "q 13 15 13 20";
  path += "q 2 5 -10 -5";

  path += "q 14 16 15 20";
  path += "q 2 5 -9 -2";

  path += "q 14 12 18 19";
  path += "q -4 2 -18 -4";

  path += "q 6 10 10 14";
  path += "q 2 5 -10 -3";

  // chin
  path += "q -10 5 -20 15";

  path += "q -45 15 -90 0";

  path += "q -10 -10 -20 -15";

  // left beard
  path += "q -12 8 -10 3 ";
  path += "q 4 -4 10 -14 ";

  path += "q -14 6 -18 4 ";
  path += "q 4 -7 18 -19 ";

  path += "q -11 7 -9 2 ";
  path += "q 1 -4 15 -20 ";

  path += "q -12 10 -10 5 ";
  path += "q 0 -5 8 -20 ";

  // path += "q -2 5 -10 5";
  // path += "q -13 15 9 -12";

  // left side
  path += "q -2 -55 7 -55";

  return path;
});

const leftEar = computed(() => {
  let path = `M ${props.startX + 15} ${props.startY} `;

  path += "q -10 -4 -25.5 -3.5 q -18 5 -24 15 q -8 9 -8 21 q 10 -3 23 -1 q 13 3 32 14";

  return path;
});

const rightEar = computed(() => {
  let path = `M ${props.startX + 85} ${props.startY} `;

  path += "q 10 -4 31 -5 q 18 5 30 16 q 5 5 5 15 q -10 -3 -21 -2 q -13 3 -30 14";

  return path;
});

</script>

<template>
  <g data-type="head">
    <path data-type="left-ear" stroke="black" stroke-width="2" fill="#B88058" :d="leftEar" />
    <path data-type="right-ear" stroke="black" stroke-width="2" fill="#B88058" :d="rightEar" />
    <path data-type="head-shape" stroke="black" stroke-width="2" fill="#B88058" :d="shape" />
  </g>

</template>

<style scoped lang="scss">

</style>