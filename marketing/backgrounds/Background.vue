<script setup>
import Blackboard from "./parts/blackbaord/Blackboard.vue";
import Walls from "./parts/Walls.vue";
</script>

<template>
  <svg xmlns="http://www.w3.org/2000/svg" width="850" height="400"
       style="shape-rendering: crispEdges; background-color: #f8f4ea; border: 1px solid black"
       viewBox="0 0 850 400">
    <Walls :width="850" :height="400" />
    <Blackboard :start-x="315" :start-y="50" />
    <!--    <Desk :start-x="440" :start-y="250" />-->
    <!--    <Chair :x="100" :y="180" />-->
  </svg>
</template>

<style scoped lang="scss">

</style>