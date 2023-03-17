<template>
  <FullHeightContainer v-for="i in pages" :bg="i === 1 ? 'gradient' : i % 2 !== 0 ? 'gray' : ''" :top="i === top"
                       :first="i === 1"
                       :last="i === pages"
                       @next="top = i + 1" @prev="top = i - 1">
    <slot :name="`content-${i}`" :jumpTo="jumpTo" />
  </FullHeightContainer>
</template>

<script>
import FullHeightContainer from "./FullHeightContainer.vue";

export default {
  name: "FullHeightLayout",
  components: {FullHeightContainer},
  props: ["pages"],
  data: () => {
    return {
      top: 1
    }
  },
  mounted() {
    window.addEventListener("scroll", () => {
      this.top = -1;
    });
  },
  methods: {
    jumpTo(i) {
      this.top = i;
    }
  }
}
</script>

<style scoped>

</style>