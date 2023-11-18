<script setup>
import {computed} from "vue";

const props = defineProps({
  startX: Number,
  startY: Number,
})

const shape = computed(() => {
  let path = `M ${props.startX} ${props.startY}`;

  path += "q 2 -17 1 -28 q -5 6 -7 12 q 8 -1 18 -1 l 0 6 q -14 1 -26 0 q 6 -13 13 -23 q 4 -1 8 0 q 1 16 0 34 l -7 0";

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