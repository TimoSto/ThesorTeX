<template>
  <div style="min-height: 100vh; display: flex; align-items: center; box-sizing: border-box" :style="bgColor">
    <v-container :style="padding">
      <slot />
    </v-container>
  </div>
</template>

<script>
export default {
  name: "FullHeightContainer",
  props: ["bg", "first"],
  computed: {
    bgColor() {
      return this.bg === "gradient" ? "background-image: linear-gradient(90deg, #0c8635, #259b71, #69beaf);" : "background-color: white";
    },
    padding() {
      return this.first ? "padding-top: 100px;" : "";
    }
  }
}
</script>

<style scoped>

</style>