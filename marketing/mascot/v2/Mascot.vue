<script setup lang="ts">
import Eye from "./parts/Eye.vue";
import Snout from "./parts/Snout.vue";
import Head from "./parts/Head.vue";
import Hat from "./parts/Hat.vue";
import Robe from "./parts/Robe.vue";

const props = defineProps({
  right: Boolean,
});

</script>

<template>
  <svg xmlns="http://www.w3.org/2000/svg" width="300" height="300"
       viewBox="0 0 300 300">
    <Robe :start-x="100" :start-y="175" :right="props.right" />
    <g id="head">
      <Head :start-x="95" :start-y="50" />
      <Eye :start-x="props.right ? 110 : 103" :start-y="95" />
      <Eye :start-x="props.right ? 160 : 153" :start-y="95" right />
      <Snout :start-x="props.right ? 155 : 128" :start-y="112" :right="props.right" />
      <Hat :start-x="103" :start-y="52" />
    </g>
  </svg>
</template>

<style scoped lang="scss">

</style>