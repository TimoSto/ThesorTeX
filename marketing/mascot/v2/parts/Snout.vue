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
  path += "q -11 -4 -40 3";

  // part covered by nose
  path += "l -11 7 ";

  // lower shaft
  path += "q 10 40 70 42";

  // mouth
  path += "q 4 0 15 -21";

  // smile wrinkle
  path += "m -7 1 q 9 -2 13 0";

  return path;
});

</script>

<template>
  <g data-type="snout">
    <path stroke="black" stroke-width="2" fill="none" :d="shape" />
  </g>
</template>

<style scoped lang="scss">

</style>