<template>
  <FullHeightContainer v-for="i in pages" :bg="i === 1 ? 'gradient' : i % 2 !== 0 && !white ? 'gray' : ''"
                       :small-display="smallDisplay"
                       @scrolled-into-view="setLowestVisible(i)" ref="container">
    <slot :name="`content-${i}`" :jumpTo="jumpTo" />
  </FullHeightContainer>
  <v-btn icon size="65" color="transparent" flat class="scroll-btn scroll-up" v-if="lowestVisible > 1"
         @click="jumpTo(lowestVisible - 1)">
    <v-icon size="55" style="color: rgba(0, 0, 0, 0.5)">
      mdi-arrow-up-circle-outline
    </v-icon>
  </v-btn>
  <v-btn icon size="65" color="transparent" flat class="scroll-btn scroll-down" v-if="lowestVisible < pages"
         @click="jumpTo(lowestVisible + 1)">
    <v-icon size="55" style="color: rgba(0, 0, 0, 0.5)">
      mdi-arrow-down-circle-outline
    </v-icon>
  </v-btn>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import FullHeightContainer from "./FullHeightContainer.vue";

export default defineComponent({
  name: "FullHeightLayout",
  components: {FullHeightContainer},
  props: ["pages", "white", "smallDisplay"],
  data: () => {
    return {
      top: 1,
      lowestVisible: 1,
      scrollTriggeredByButton: false
    };
  },
  mounted() {
    window.addEventListener("scroll", () => {
      if (window.scrollY === 0) {
        this.setLowestVisible(1);
      }
    });
  },
  methods: {
    jumpTo(i: number) {
      this.scrollTriggeredByButton = true;
      (this.$refs.container as any[])[i - 1].scrollToTop();
      this.lowestVisible = i;
      setTimeout(() => {
        this.scrollTriggeredByButton = false;
      }, 750);
    },
    setLowestVisible(n: number) {
      if (n !== this.lowestVisible && !this.scrollTriggeredByButton) {
        this.lowestVisible = n;
      }
    }
  }
});
</script>

<style scoped lang="scss">
.scroll-btn {
  position: fixed;

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