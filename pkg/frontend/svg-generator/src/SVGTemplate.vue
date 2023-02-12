<template>
  <svg width="500" height="500" viewBox="-100 -100 500 500">
    <path stroke="red" stroke-width="10" fill="white" :d="path" />
  </svg>
</template>

<script lang="ts" setup>//TODO: turn into general svg component
import turnVector, {Vector} from "./helper/turnVector";
import {computed} from "vue";

const props = defineProps({
  turn: Number,
});

const startPoint: Vector = {
  x: 0,
  y: 50,
};

// this is the collection of points, the vectors have to be calculated based on the turning
const basePath: Vector[] = [
  {
    x: 100,
    y: 50,
  },
  {
    x: 100,
    y: 100
  },
  {
    x: 0,
    y: 100
  },
  {
    x: 0,
    y: 50
  }
];

// computed
const path = computed(() => {
  let p = `M${startPoint.x},${startPoint.y}`;

  let prevVec = {
    x: startPoint.x,
    y: startPoint.y
  };

  let prevVecTurned = {
    x: startPoint.x,
    y: startPoint.y
  };

  const vectors = basePath.map(v => {
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