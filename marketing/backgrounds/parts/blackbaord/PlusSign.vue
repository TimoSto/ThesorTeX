<script setup>
import {computed} from "vue";

const props = defineProps({
  startX: Number,
  startY: Number,
})

const shape = computed(() => {
  let path = `M ${props.startX} ${props.startY}`;

  path += "l 0 -10 l 5 0 l 0 10 l 10 0 l 0 5 l -10 0 l 0 10 l -5 0 l 0 -10 l -10 0 l 0 -5 l 10 0";

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