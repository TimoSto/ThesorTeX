<script setup>
import {computed} from "vue";

const props = defineProps({
  startX: Number,
  startY: Number,
})

const shape = computed(() => {
  let path = `M ${props.startX} ${props.startY}`;

  path += "l 10 -10 l 7 0 q 1 15 0 30 l -7 0 l 0 -20 l -5 5 q -3 -2 -5 -5";

  return path;
})

</script>

<template>
  <g data-type="two">
    <path :d="shape" stroke="white" fill="white" />
  </g>
</template>

<style scoped lang="scss">

</style>