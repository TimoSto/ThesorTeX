<script setup lang="ts">
import {computed} from "vue";

const props = defineProps({
  startX: Number,
  startY: Number,
  right: Boolean,
});

const shape = computed(() => {
  let path = `M ${props.startX} ${props.startY}`;

  // upper shaft
  path += `q ${props.right ? "" : "-"}11 -4 ${props.right ? "" : "-"}40 3`;

  // part covered by nose
  path += `l ${props.right ? "" : "-"}11 7 `;

  // lower shaft
  path += `q ${props.right ? "-" : ""}10 40 ${props.right ? "-" : ""}70 42`;

  // mouth
  path += `q ${props.right ? "-" : ""}4 0 ${props.right ? "-" : ""}15 -21`;

  // smile wrinkle
  path += `m ${props.right ? "" : "-"}7 1 q ${props.right ? "-" : ""}9 -2 ${props.right ? "-" : ""}13 0`;

  return path;
});

const nose = computed(() => {
  let path = `M ${props.right ? props.startX + 55 : props.startX - 40} ${props.startY + 3}`;

  // upper
  path += "q -11 -2 -15 0";

  // left
  path += "q -5 7 7 15";

  // right
  path += "q 10 -6 7 -15";

  return path;
});

const smileWrinkle = computed(() => {
  let path = `M ${props.right ? props.startX - 12 : props.startX + 15} ${props.startY + 57}`;

  path += `q ${props.right ? "-" : ""}10 6 ${props.right ? "-" : ""}15 -4`;

  return path;
});

// TODO: maybe add whiskers

</script>

<template>
  <g data-type="snout">
    <path data-type="snout-shaft" stroke="black" stroke-width="2" fill="#B88058" :d="shape" />
    <path data-type="smile-wrinkle" stroke="black" stroke-width="2" fill="none" :d="smileWrinkle" />
    <path data-type="nose" stroke="black" fill="#58676A" stroke-width="2" :d="nose" />
  </g>
</template>

<style scoped lang="scss">

</style>