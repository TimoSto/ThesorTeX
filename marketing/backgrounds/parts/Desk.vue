<script setup>

import {computed} from "vue";

const props = defineProps({
  startX: Number,
  startY: Number,
})

const deskShape = computed(() => {
  let path = `M ${props.startX} ${props.startY}`

  path += "l 125 0 l 0 50 l -125 0 l 0 -50"; // front

  path += "l 15 -10 l 125 0 l -15 10"; // top

  path += "l 0 50 l 15 -10 l 0 -50"; // right side

  path += "l -15 10"; // to fill correctly

  return path
})

</script>

<template>
  <g data-type="desk">
    <path :d="deskShape" fill="#90684A" stroke="black" />
  </g>
</template>

<style scoped lang="scss">

</style>