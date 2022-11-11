<template>
  <div class="container">
    <div class="left">

    </div>
    <div class="right">

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
  width: 100%;
  display: flex;
  & div {
    height: 100%;
    flex: 1;
  }
  & .left {
    background-color: red;
  }
  & .right {
    background-color: green;
  }
}
</style>