<template>
  <svg></svg>
</template>

<script lang="ts" setup>
import {TemplateSVG} from "./helper/SVG";
import {PropType} from "vue";

const props = defineProps({
  svg: Object as PropType<TemplateSVG>,
});
</script>

<style scoped>

</style>