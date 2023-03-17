<template>
  <div style="min-height: 100vh; display: flex; align-items: center; box-sizing: border-box; position: relative;"
       :style="bgColor" ref="container">
    <v-container :style="padding">
      <slot />
    </v-container>
    <v-btn icon size="75" color="transparent" flat class="scroll-btn scroll-up" @click="$emit('prev')" v-if="!first">
      <v-icon size="65"
              :style="`${bg === 'gradient' ? 'color: rgba(255, 255, 255, 0.75);' : 'color: rgba(0, 0, 0, 0.5);'}`">
        mdi-arrow-up-circle-outline
      </v-icon>
    </v-btn>
    <v-btn icon size="75" color="transparent" flat class="scroll-btn scroll-down" @click="$emit('next')" v-if="!last">
      <v-icon size="65"
              :style="`${bg === 'gradient' ? 'color: rgba(255, 255, 255, 0.75);' : 'color: rgba(0, 0, 0, 0.5);'}`">
        mdi-arrow-down-circle-outline
      </v-icon>
    </v-btn>
  </div>
</template>

<script>
export default {
  name: "FullHeightContainer",
  props: ["bg", "first", "last", "top"],
  emits: ["next", "prev"],
  computed: {
    bgColor() {
      switch (this.bg) {
        case "gradient":
          return "background-image: linear-gradient(90deg, #0c8635, #259b71, #69beaf); ";
        case "gray":
          return "background-color: rgba(0,0,0,0.25);"
        default:
          return "background-color: white; "
      }
    },
    padding() {
      return this.first ? "" : "";
    },
  },
  watch: {
    top: function () {
      if (this.top) {
        this.$refs.container.scrollIntoView({behavior: "smooth"});
      }
    }
  }
}
</script>

<style scoped lang="scss">
.scroll-btn {
  position: absolute;

  &.scroll-down {
    right: 0;
    bottom: 0;
    transform: translate(-20px, -20px);
  }

  &.scroll-up {
    right: 0;
    top: 0;
    transform: translate(-20px, 20px);
  }
}
</style>