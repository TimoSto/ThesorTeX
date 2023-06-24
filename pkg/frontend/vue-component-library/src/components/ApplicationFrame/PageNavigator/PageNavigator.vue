<script setup lang="ts">
import {computed} from "vue";

const props = defineProps({
  numberOfPages: {
    type: Number,
    required: true
  },
  nextNumberOfPages: {
    type: Number,
    required: true
  }
});

const NavigatingStates = {
  forward: 1,
  backward: -1,
  backwardMultiple: -2,
  none: 0
};

const navigatingState = computed(() => {
  const diff = props.nextNumberOfPages - props.numberOfPages;
  if (diff === 0) {
    return NavigatingStates.none;
  }
  if (diff === 1) {
    return NavigatingStates.forward;
  }
  if (diff === -1) {
    return NavigatingStates.backward;
  }

  return NavigatingStates.backwardMultiple;
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