<template>
  <svg width="500" height="500" viewBox="-100 -100 500 500">
    <path stroke="red" stroke-width="10" fill="white" :d="path" />
  </svg>
</template>

<script lang="ts" setup>//TODO: turn into general svg component
import turnVector, {Vector} from "./helper/turnVector";
import {computed, PropType} from "vue";

const props = defineProps({
  turn: Number,
  startPoint: {
    type: Object as PropType<Vector>,
    default: {
      x: 0,
      y: 0
    }
  },
  points: Array<Vector>
});

// computed
const path = computed(() => {
  const startPoint: Vector = {
    x: props.startPoint.x,
    y: props.startPoint.y
  };

  let p = `M${startPoint.x},${startPoint.y}`;

  let prevVec = {
    x: startPoint.x,
    y: startPoint.y
  };

  let prevVecTurned = {
    x: startPoint.x,
    y: startPoint.y
  };

  const vectors = props.points.map(v => {
    const relVec = {
      x: v.x - prevVec.x,
      y: v.y - prevVec.y
    };

    const turned = turnVector(relVec, props.turn);

    const result = {
      x: prevVecTurned.x + turned.x,
      y: prevVecTurned.y + turned.y
    };

    prevVec.x = v.x;
    prevVec.y = v.y;

    prevVecTurned.x = result.x;
    prevVecTurned.y = result.y;

    return result;
  });

  vectors.forEach(v => {
    p += `L${v.x},${v.y} `;
  });

  p += "z";

  return p;
});

</script>

<style scoped>

</style>