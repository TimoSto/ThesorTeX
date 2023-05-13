<script setup lang="ts">
// computed
import {computed, onMounted, ref} from "vue";
import generateWaveSVG from "./waves/generateWaveSVG";

// data
const width = ref(0);
const height = ref(0);

// computed
const bgSvg = computed(() => {
  const svg = generateWaveSVG(width.value, height.value, (x: number) => 100 * Math.sin(x / 100) + 0.25 * height.value);

  return window.btoa(svg);
});


//onmounted
onMounted(() => {
  width.value = window.innerWidth;
  height.value = window.innerHeight;

  window.onresize = () => {
    width.value = window.innerWidth;
    height.value = window.innerHeight;
  };
});

</script>

<template>
  <div :style="`background-image: url(data:image/svg+xml;base64,${bgSvg})`" class="container"></div>
</template>

<style scoped lang="scss">
.container {
  width: 100vw;
  height: 100vh;
}
</style>