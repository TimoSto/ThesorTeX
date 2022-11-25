<template>
  <div class="container">
    <div class="left" :class="imageSide === 'left' ? 'image-container' : 'tile'">
      <slot name="left"></slot>
    </div>
    <div class="right" :class="imageSide === 'right' ? 'image-container' : 'tile'">
      <slot name="right"></slot>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps({
  imageSide: String,
})
</script>

<style lang="scss" scoped>
@use "@/styles/tiles.scss";

.container{
  width: calc(100% - 64px);
  margin: 16px auto;
  display: flex;
  & div {
    height: 100%;
    flex: 1;
    &.image-container {
      @media screen and (max-width: 900px) {
        display: none;
      }
    }
  }
}
</style>
