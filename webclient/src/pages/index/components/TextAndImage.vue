<template>
  <div class="container">
    <div class="left" :class="imageSide === 'left' ? 'image-container' : ''">
      <slot name="left"></slot>
    </div>
    <div class="right" :class="imageSide === 'right' ? 'image-container' : ''">
      <slot name="right"></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: "TextAndImage",
  props: [
      "imageSide"
  ]
}
</script>

<style scoped lang="scss">
.container{
  width: calc(100% - 32px);
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