<script setup lang="ts">
defineProps({
  numberOfPages: Number
});
</script>

<template>
  <div class="container">
    <div class="page" v-for="n in numberOfPages" :class="`${n === numberOfPages ? 'opened' : ''}`">
      <slot :name="n" />
    </div>
  </div>
</template>

<style scoped lang="scss">
.container {
  display: flex;
  height: 100%;
  width: 100%;
  padding: 0;

  & .page {
    flex: 1 0 auto;
    height: 100%;
    width: 0;
    overflow-x: hidden;
    transition: width .75s ease-in-out;

    &.opened {
      width: 100%;
    }
  }
}
</style>