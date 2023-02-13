<template>
  <svg width="500" height="500" :viewBox="viewBox">
    <path stroke="black" stroke-width="5" d="M0,-250 l0,500 z" />
    <path stroke="black" stroke-width="5" d="M-250,0 l500,0 z" />
    <path stroke="red" stroke-width="10" fill="white" :d="path" />
  </svg>
</template>

<script lang="ts" setup>//TODO: turn into general svg component
import {computed} from "vue";
import generatePath, {PathPart} from "./helper/generatePath";

const props = defineProps({
  angle: Number,
  viewBox: String,
  points: Array<PathPart>
});

// computed
const path = computed(() => {
  return generatePath(props.points, props.angle);
});

</script>

<style scoped>

</style>