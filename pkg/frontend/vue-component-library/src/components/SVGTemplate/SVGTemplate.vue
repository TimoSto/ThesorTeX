<template>
  <svg :width="svgEl.width" :height="svgEl.height" :viewBox="svgEl.viewBox" xmlns="http://www.w3.org/2000/svg"
       :style="`max-height: ${maxHeight}; max-width: ${maxWidth}`">
    <!--    <path stroke="black" stroke-width="5" d="M0,-250 l0,500 z" />-->
    <!--    <path stroke="black" stroke-width="5" d="M-250,0 l500,0 z" />-->
    <g v-for="(part, i) in svgEl.partials" :key="`partial-${i}`"
       :transform="generateTransform(part)">
      <path v-for="(p, j) in part.parts" :key="`partial-${i}-${j}`" :stroke="part.strokeColor"
            :stroke-width="part.strokeWidth" :fill="part.fillColor"
            :d="generatePath(p)" />
    </g>
  </svg>
</template>

<script lang="ts" setup>
import {TemplateSVG} from "./helper/SVG";
import {computed, PropType} from "vue";
import {generatePath, generateTransform} from "./helper/generatePath";

const props = defineProps({
  maxHeight: String,
  maxWidth: String,
  svg: Object as PropType<TemplateSVG>,
});

const svgEl = computed(() => {
  return JSON.parse(JSON.stringify(props.svg));
});

</script>

<style scoped>

</style>