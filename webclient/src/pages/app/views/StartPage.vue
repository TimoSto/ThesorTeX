<template>
  <div>

    <v-app-bar
      color="background"
      dense
      elevate-on-scroll
      scroll-target="#index-content"
      absolute
      >
      <v-toolbar-title>Willkommen bei ThesorTeX!</v-toolbar-title>
    </v-app-bar>

    <ContentBelowBar
        bar-height="112"
        margin-top="48"
        padding-x="16"
        padding-y="0"
        scroll-id="index-content"
    >
      <div style="width: 100px; height:2000px; background-color: red; margin: 8px"></div>
    </ContentBelowBar>

  </div>
</template>

<script lang="ts">
import Vue from "vue";
import ContentBelowBar from "../../../common/components/ContentBelowBar.vue";

export default Vue.extend({
  name: "StartPage",
  components: {ContentBelowBar}
})
</script>

<style scoped>

</style>