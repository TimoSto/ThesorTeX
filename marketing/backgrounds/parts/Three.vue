<script setup>
import {computed} from "vue";

const props = defineProps({
  startX: Number,
  startY: Number,
})

const shape = computed(() => {
  let path = `M ${props.startX} ${props.startY}`;

  path += "q 6 -9 15 -9 q 12 4.5 3 16.5 q 12 10.5 3 18 q -9 3 -18 -1.5 q 1.5 -3.75 3 -6 q 10.5 3 10.5 -1.5 q -1.5 -4.5 -9 -4.5 l 0 -4.5 q 7.5 -6 6 -7.5 q -3 -3 -10.5 4.5 l -3 -4.5";

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