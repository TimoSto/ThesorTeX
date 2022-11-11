<template>
  <div class="container">
    <div class="left">
      <slot name="left"></slot>
    </div>
    <div class="right">
      <slot name="right"></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: "TextAndImage"
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
  }
}
</style>