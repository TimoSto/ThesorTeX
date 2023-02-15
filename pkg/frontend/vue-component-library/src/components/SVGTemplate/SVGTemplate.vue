<template>
  <svg :width="svg.width" :height="svg.height" :viewBox="svg.viewBox" xmlns="http://www.w3.org/2000/svg">
    <g v-for="(part, i) in svg.partials" :key="`partial-${i}`"
       :transform="generateTransform(part)">
      <text>{{ generateTransform(part) }}</text>
      <path :stroke="part.strokeColor" :stroke-width="part.strokeWidth" :fill="part.fillColor"
            :d="generatePath(part.parts)" />
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