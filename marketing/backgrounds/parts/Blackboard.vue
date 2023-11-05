<script setup>

import {computed} from "vue";
import Two from "./Two.vue";
import PlusSign from "./PlusSign.vue";
import EqualSign from "./EqualSign.vue";
import Four from "./Four.vue";
import MinusSign from "./MinusSign.vue";
import One from "./One.vue";
import Three from "./Three.vue";

const border = computed(() => {
  let path = `M 50 175`;

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
    <path data-type="border" stroke="#E95D26" fill="#828232" stroke-width="6" :d="border" />
    <Two :start-x="50" :start-y="200" />
    <PlusSign :start-x="95" :start-y="210" />
    <Two :start-x="115" :start-y="200" />
    <EqualSign :start-x="150" :start-y="204" />
    <Four :start-x="195" :start-y="225" />
    <MinusSign :start-x="155" :start-y="255" />
    <One :start-x="182" :start-y="250" />
    <EqualSign :start-x="212" :start-y="249" />
    <Three :start-x="240" :start-y="245" />
  </g>
</template>

<style scoped lang="scss">

</style>