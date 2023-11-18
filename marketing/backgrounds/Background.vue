<script setup>
import Blackboard from "./parts/Blackboard.vue";
</script>

<template>
  <svg xmlns="http://www.w3.org/2000/svg" width="850" height="400"
       style="shape-rendering: crispEdges; background-color: #ffffff; border: 1px solid black"
       viewBox="0 0 850 400">
    <Blackboard :start-x="290" :start-y="75" />
  </svg>
</template>

<style scoped lang="scss">

</style>