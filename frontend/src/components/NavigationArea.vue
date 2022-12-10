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

import {watch} from "vue";

// props
const props = defineProps({
  pages: {
    type: Number,
    required: true,
    default: 0
  },
  instantSwitch: {
    type: Boolean,
    default: false
  },
  navigatingBack: {
    type: Boolean,
    default: false
  }
})

// data
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

<style scoped lang="scss">
.disable-animations {
  & * {
    transition: none!important;
  }
}

.container {
  display: flex;
  height: 100%;
  width: 100%;
  padding: 0;
  max-width: unset;
  & .pages {
    width: 100%;
    display: flex;
    flex: 1 1 auto;
    box-sizing: border-box;
    transition: width .5s ease-in-out;
    overflow-x: hidden;
    & .page {
      flex: 1 0 auto;
      height: 100%;
      max-height: calc(100vh - 64px);
      width: 0;
      overflow-x: hidden;
      transition: width .75s ease-in-out;
      &.opened {
        width: 100%;
      }
      & .page--container {
        width: 100%;
        min-height: 100%;
        min-width: 150px;
      }
    }
  }

  &.nav-back {
    & .nav-back-area {
      width: 0;
    }
  }

  &.two-thirds {
    & .pages {
      width: 30%;
    }
    & .edit-area {
      width: 70%;
    }
  }

  &.half-screen {
    & .pages {
      width: 50%;
    }
    & .edit-area {
      width: 50%;
    }
  }
}
</style>
