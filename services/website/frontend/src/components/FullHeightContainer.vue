<template>
  <div class="fullHeightContainer" :class="smallDisplay ? 'small': 'large'"
       :style="bgColor" ref="container">
    <v-container>
      <slot />
    </v-container>
  </div>
</template>

<script>
export default {
  name: "FullHeightContainer",
  props: ["bg", "first", "last", "top", "smallDisplay"],
  emits: ["scrolledIntoView"],
  computed: {
    bgColor() {
      switch (this.bg) {
        case "gradient":
          return "background-image: linear-gradient(90deg, #0c8635, #259b71, #69beaf); ";
        case "gray":
          return "background-color: rgba(0,0,0,0.15);"
        default:
          return "background-color: white; "
      }
    },
  },
  methods: {
    scrollToTop() {
      this.$refs.container.scrollIntoView({behavior: "smooth"});
    },
    onScroll() {
      let docViewTop = document.documentElement.scrollTop;
      let docViewBottom = docViewTop + window.innerHeight - 200;

      let elemTop = this.$refs.container.offsetTop;

      if ((elemTop < docViewBottom) && (elemTop > docViewTop)) {
        this.$emit("scrolledIntoView");
      }
    }
  },
  mounted() {
    window.addEventListener("scroll", this.onScroll, true)
  },
  unmounted() {
    window.removeEventListener("scroll", this.onScroll, true)
  }
}
</script>

<style scoped lang="scss">
.fullHeightContainer {
  display: flex;
  align-items: center;
  box-sizing: border-box;
  position: relative;

  &.large {
    min-height: 100vh;
  }

  &.small {
    min-height: 33vh;
    padding-bottom: 75px;
    padding-top: 75px;
  }
}
</style>