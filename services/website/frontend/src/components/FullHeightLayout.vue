<template>
  <FullHeightContainer v-for="i in pages" :bg="i === 1 ? 'gradient' : ''" :top="i === top" @next="top = i + 1">
    <slot :name="`content-${i}`" />
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
  }
}
</script>

<style scoped>

</style>