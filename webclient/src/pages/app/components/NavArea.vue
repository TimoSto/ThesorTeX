<template>
  <div class="container">
    <div class="pages">
      <div
          v-for="(n, i) in pages"
          :key="`page-${i}`"
          class="page"
          :class="i === pages.length - 1 ? 'opened' : ''"
      >
        <slot :name="n"></slot>
      </div>
    </div>

  </div>
</template>

<script lang="ts">
import Vue from "vue";

export interface NavAreaMethods extends Vue {
  addPage(name: string): void
  goBackTo(name: string): void
}

export default Vue.extend({
  name: "NavArea",
  data() {
    return {
      pages: [] as string[]
    }
  },
  methods: {
    addPage(name: string) {
      this.pages.push(name)
    },
    goBackTo(name: string) {
      const i = this.pages.indexOf(name);
      console.log(i)
      if( i > -1 ) {
        this.pages = this.pages.slice(0, i+1);
      }
      console.log(this.pages)
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
      transition: width .5s ease-in-out;
      &.opened {
        width: 100%;
      }
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