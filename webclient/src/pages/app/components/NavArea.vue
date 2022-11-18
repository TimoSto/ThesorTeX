<template>
  <div class="container" :class="animationClass">
    <div class="pages">
      <div
          v-for="(n, i) in pages"
          :key="`page-${i}`"
          class="page"
          :class="`${i === pages.length - 1 ? 'opened' : ''}`"
      >
        <slot :name="n"></slot>
      </div>
      <div
          class="page opened nav-back-area"
          v-if="animationClass === 'nav-back'"
      ></div>
    </div>

  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  name: "NavArea",
  props: [
      "pages"
  ],
  data() {
    return {
      animationClass: ''
    }
  },
  watch: {
    pages(n: string[], o: string[]) {
      if( n.length < o.length ) {
        this.animationClass = 'nav-back';
        setTimeout(() => {
          this.animationClass = '';
        }, 750);
      }
    }
  }
})
</script>

<style scoped lang="scss">
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