<template>
  <svg :width="svg.width" :height="svg.height" :viewBox="svg.viewBox" xmlns="http://www.w3.org/2000/svg">
    <path stroke="black" stroke-width="5" d="M0,-250 l0,500 z" />
    <path stroke="black" stroke-width="5" d="M-250,0 l500,0 z" />
    <g v-for="(part, i) in svg.partials" :key="`partial-${i}`"
       :transform="generateTransform(part)">
      <path v-for="(p, j) in part.parts" :key="`partial-${i}-${j}`" :stroke="part.strokeColor"
            :stroke-width="part.strokeWidth" :fill="part.fillColor"
            :d="generatePath(p)" />
    </g>
  </svg>
</template>

<script lang="ts" setup>
import {TemplateSVG} from "./helper/SVG";
import {PropType} from "vue";
import {generatePath, generateTransform} from "./helper/generatePath";

const props = defineProps({
  svg: Object as PropType<TemplateSVG>,
});


</script>

<style scoped>

</style>