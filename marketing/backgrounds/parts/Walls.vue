<script setup>
import {computed} from "vue";

const props = defineProps({
  width: Number,
  height: Number,
})

const floor = computed(() => {
  let path = `M 0 ${props.height}`;

  path += `l ${props.width / 2 - 220} -${props.height / 2 - 50}`;

  path += "l 440 0";

  path += `l ${props.width / 2 - 220} ${props.height / 2 - 50}`;

  return path;
})

const wallMiddle = computed(() => {
  let path = `M ${props.width / 2 - 220} ${props.height / 2 + 50}`;

  path += `l 0 -${props.height / 2 + 50}`;

  path += "l 440 0";

  path += `l 0 ${props.height / 2 + 50} z`;

  return path;
})
</script>

<template>
  <g data-type="walls-n-floor">
    <path stroke="black" stroke-width="1" :d="floor" fill="#A0A1A6" />
    <path stroke="black" stroke-width="1" :d="wallMiddle" fill="#f4f3ef" />
  </g>
</template>

<style scoped lang="scss">

</style>