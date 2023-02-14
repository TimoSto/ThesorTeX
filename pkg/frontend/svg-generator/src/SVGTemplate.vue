<template>
  <svg width="500" height="500" :viewBox="viewBox">
    <!--    <path stroke="black" stroke-width="5" d="M0,-250 l0,500 z" />-->
    <!--    <path stroke="black" stroke-width="5" d="M-250,0 l500,0 z" />-->
    <path :stroke="paths[i].strokeColor" :stroke-width="paths[i].strokeWidth" :fill="paths[i].fillColor"
          v-for="(p, i) in rotatedPaths" :d="p" />
  </svg>
</template>

<script lang="ts" setup>//TODO: turn into general svg component
import {computed} from "vue";
import generatePath, {PathSegment} from "./helper/generatePath";

const props = defineProps({
  viewBox: String,
  paths: Array<Array<PathSegment>>
});

// computed
const rotatedPaths = computed(() => {
  return props.paths.map(p => generatePath(p.parts, p.angle));
});

</script>

<style scoped>

</style>