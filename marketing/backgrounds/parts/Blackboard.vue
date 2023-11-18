<script setup>

import {computed} from "vue";
import Two from "./Two.vue";
import PlusSign from "./PlusSign.vue";
import EqualSign from "./EqualSign.vue";
import Four from "./Four.vue";
import MinusSign from "./MinusSign.vue";
import One from "./One.vue";
import Three from "./Three.vue";

const props = defineProps({
  startX: Number,
  startY: Number,
})

// width: 270
const border = computed(() => {
  let path = `M ${props.startX} ${props.startY}`;

  path += "l 230 0"

  path += "a 20 20 0 0 1 20 20"

  path += "l 0 100"

  path += "a 20 20 0 0 1 -20 20"

  path += "l -230 0"

  path += "a 20 20 0 0 1 -20 -20"

  path += "l 0 -100"

  path += "a 20 20 0 0 1 20 -20"

  return path;
})
</script>

<template>
  <g data-type="blackboard">
    <path data-type="border" stroke="#4F3A2B" fill="#3C5E44" stroke-width="8" :d="border" />
    <Two :start-x="props.startX" :start-y="props.startY + 25" />
    <PlusSign :start-x="props.startX + 45" :start-y="props.startY + 35" />
    <Two :start-x="props.startX + 65" :start-y="props.startY + 25" />
    <EqualSign :start-x="props.startX + 100" :start-y="props.startY + 29" />
    <Four :start-x="props.startX +145" :start-y="props.startY + 50" />
    <MinusSign :start-x="props.startX +105" :start-y="props.startY + 80" />
    <One :start-x="props.startX + 132" :start-y="props.startY + 75" />
    <EqualSign :start-x="props.startX + 162" :start-y="props.startY + 74" />
    <Three :start-x="props.startX + 190" :start-y="props.startY + 70" />
  </g>
</template>

<style scoped lang="scss">

</style>