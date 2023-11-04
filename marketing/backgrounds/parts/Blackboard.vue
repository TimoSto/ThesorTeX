<script setup>

import {computed} from "vue";

const border = computed(() => {
  let path = `M 50 175`;

  path += "l 250 0"

  path += "a 20 20 0 0 1 20 20"

  path += "l 0 100"

  path += "a 20 20 0 0 1 -20 20"

  path += "l -250 0"

  path += "a 20 20 0 0 1 -20 -20"

  path += "l 0 -100"

  path += "a 20 20 0 0 1 20 -20"

  return path;
})
</script>

<template>
  <g data-type="blackboard">
    <path data-type="border" stroke="#E95D26" fill="#828232" stroke-width="6" :d="border" />
  </g>
</template>

<style scoped lang="scss">

</style>