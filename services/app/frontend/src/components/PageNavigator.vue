<template>
  <div
    class="container"
    :class="animationClass + (props.instantSwitch ? ' disable-animations' : '')"
  >
    <div class="pages">
      <div
        v-for="i in pages"
        :id="`page-${i}`"
        :key="`page-${i}`"
        class="page"
        :class="`${i === pages && !navigatingBack || i === pages -1 && navigatingBack ? 'opened' : ''}`"
      >
        <div class="page--container">
          <slot :name="i" />
        </div>
      </div>
      <div
        v-if="animationClass === 'nav-back'"
        class="page opened nav-back-area"
      />
    </div>
  </div>
</template>
<script setup lang="ts">

// data
import {watch} from "vue";

const props = defineProps({
  instantSwitch: {
    type: Boolean,
    required: true
  },
  pages: {
    type: Number,
    required: true,
    default: 0
  },
  navigatingBack: {
    type: Boolean,
    default: false
  }
})

let animationClass = "";

// watchers
watch(() => props.pages, (newValue: number, oldValue: number) => {
  if( newValue < oldValue ) {
    animationClass = 'nav-back';
    setTimeout(() => {
      animationClass = '';
    }, 750);
  }
})
</script>