<template>
  <div class="container">
    <div
      class="left"
      :class="props.imageSide === 'left' ? 'image-container' : 'tile'"
    >
      <slot name="left" />
    </div>
    <div
      class="right"
      :class="props.imageSide === 'right' ? 'image-container' : 'tile'"
    >
      <slot name="right" />
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps({
  imageSide: {
    type: String,
    default: ''
  },
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
    padding: 8px;
    &.image-container {
      @media screen and (max-width: 900px) {
        display: none;
      }
    }
  }
}
</style>
