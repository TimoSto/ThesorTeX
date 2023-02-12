<template>
  <svg width="500" height="500" viewBox="-100 -100 500 500">
    <path stroke="red" stroke-width="10" fill="white" :d="path" />
  </svg>
</template>

<script lang="ts" setup>//TODO: turn into general svg component
import {Vector} from "./helper/turnVector";
import {computed} from "vue";
import generatePath from "./helper/generatePath";

const props = defineProps({
  angle: Number,
  points: Array<Vector>
});

// computed
const path = computed(() => {
  return generatePath(props.points, props.angle);
});

</script>

<style scoped>

</style>