<script setup>
import {computed} from "vue";

const props = defineProps({
  startX: Number,
  startY: Number,
})

const shape = computed(() => {
  let path = `M ${props.startX} ${props.startY}`;

  path += "l 21 0 l 0 5 l -21 0 l 0 -5";

  return path;
})

const shape2 = computed(() => {
  let path = `M ${props.startX} ${props.startY + 11}`;

  path += "l 21 0 l 0 5 l -21 0 l 0 -5";

  return path;
})

</script>

<template>
  <g data-type="two">
    <path :d="shape" stroke="white" fill="white" />
    <path :d="shape2" stroke="white" fill="white" />
  </g>
</template>

<style scoped lang="scss">

</style>