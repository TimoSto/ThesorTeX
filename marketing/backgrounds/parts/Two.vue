<script setup>
import {computed} from "vue";

const props = defineProps({
  startX: Number,
  startY: Number,
})

const shape = computed(() => {
  let path = `M ${props.startX} ${props.startY}`;

  path += "q 3 -8 10 -10 q 5 0 10 4 q 7 9 2 15 q -3 6 -8 11 q 5 0 9 3 q 4 3 0 5 q -11 2 -19 1 q -9 -2 0 -9 q 10 -11 9 -18 q -3 -7 -9 2 q -4 -2 -4 -4";

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