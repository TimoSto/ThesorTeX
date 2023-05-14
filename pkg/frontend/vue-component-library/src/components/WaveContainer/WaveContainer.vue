<script setup lang="ts">
import {computed, onMounted, ref} from "vue";
import generateWaveSVG from "./waves/generateWaveSVG";

const props = defineProps({
  waveFunc: Number,
  bgColor: String,
  waveColor: String,
});

// data
const width = ref(0);
const height = ref(0);

const content = ref(null);

// computed
const bgSvg = computed(() => {
  const fn = waveFunc1;
  const svg = generateWaveSVG(width.value, height.value, props.bgColor!, props.waveColor!, fn);

  return window.btoa(svg);
});

// methods
function waveFunc1(x: number): number {
  return 50 * Math.sin(x / 100) + 30 * Math.cos((x + 150) / 200) + 0.35 * height.value;
}


//onmounted
onMounted(() => {
  width.value = (content.value! as HTMLElement).clientWidth;
  height.value = (content.value! as HTMLElement).clientHeight;

  window.onresize = () => {
    width.value = (content.value! as HTMLElement).clientWidth;
    height.value = (content.value! as HTMLElement).clientHeight;
  };
});

</script>

<template>
  <div :style="`background-image: url(data:image/svg+xml;base64,${bgSvg})`" class="container" ref="content">
    <slot></slot>
  </div>
</template>

<style scoped lang="scss">
.container {
  width: 100vw;
  min-height: 100vh;
}
</style>