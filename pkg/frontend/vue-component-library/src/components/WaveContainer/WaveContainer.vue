<script setup lang="ts">
import {computed, onBeforeUnmount, onMounted, ref} from "vue";
import generateWaveSVG from "./waves/generateWaveSVG";

const props = defineProps({
  waveFunc: Number,
  bgColor: String,
  waveColor: String,
  smallDisplay: Boolean
});

// data
const width = ref(0);
const height = ref(0);

const content = ref(null);

// computed
const bgSvg = computed(() => {
  const svg = generateWaveSVG(width.value, height.value, props.bgColor!, props.waveColor!, waveFunc, props.waveFunc === -1);

  return window.btoa(svg);
});

// methods
function waveFunc(x: number): number {
  switch (props.waveFunc) {
    case 2:
      return 20 * Math.sin(x / 100) + 50 * Math.cos((x + 150) / 200) - 30 * Math.sin(x / 750) + 0.4 * height.value;
    case 3:
      return 40 * Math.cos(x / 100) + 50 * Math.sin((x + 150) / 200) - 10 * Math.sin(x / 500) + 0.35 * height.value;
    default:
      return 50 * Math.sin(x / 200) + 30 * Math.cos((x + 150) / 600) + 0.35 * height.value;
  }
}

function recalculateDimensions() {
  width.value = (content.value! as HTMLElement).clientWidth;
  height.value = (content.value! as HTMLElement).clientHeight;
}

defineExpose({
  recalculateDimensions
});

//onmounted
function resizeHandler() {
  width.value = (content.value! as HTMLElement).clientWidth;
  height.value = (content.value! as HTMLElement).clientHeight;
}

onMounted(() => {
  width.value = (content.value! as HTMLElement).clientWidth;
  height.value = (content.value! as HTMLElement).clientHeight;

  window.addEventListener("resize", resizeHandler);
});

onBeforeUnmount(() => {
  window.removeEventListener("resize", resizeHandler);
});

</script>

<template>
  <div :style="`background-image: url(data:image/svg+xml;base64,${bgSvg})`" class="wave-container"
       :class="smallDisplay ? 'small': 'large'" ref="content">
    <v-container>
      <slot></slot>
    </v-container>
  </div>
</template>

<style scoped lang="scss">
.wave-container {
  display: flex;
  align-items: center;
  box-sizing: border-box;
  position: relative;

  &.large {
    min-height: 100vh;
  }

  &.small {
    min-height: 33vh;
    padding-bottom: 75px;
    padding-top: 75px;
  }
}
</style>