<script setup>

import {computed} from "vue";

const props = defineProps({
  x: Number,
  y: Number,
})

const path = computed(() => {
  let path = `M ${props.x} ${props.y}`

  path += "l 40 0 l 0 50 l -40 0 l 0 -50"

  return path;
});
</script>

<template>
  <path stroke="black" stroke-width="1" fill="#90684A" :d="path" />
</template>

<style scoped lang="scss">

</style>