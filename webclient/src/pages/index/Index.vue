<template>
  <v-app>
    <v-app-bar
      fixed
      elevation="0"
      color="background"
    >
      <LogoSVG
          fill="rgba(0, 136, 52, 0.8)"
          style="height: 35px"
          :show-title="true"
          font-color="rgba(0,0,0,0.7)"
      />

      <v-spacer />

      <v-btn text color="primary">
        Anmelden
      </v-btn>

    </v-app-bar>
    <ContentBelowBar
      bar-height="64"
      padding-x="16"
      padding-y="0"
    >
      hallo
    </ContentBelowBar>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import LogoSVG from "../../commonComponents/LogoSVG.vue";
import ContentBelowBar from "../../commonComponents/ContentBelowBar.vue";

export default Vue.extend({
  name: 'App',
  components: {ContentBelowBar, LogoSVG},
  data: () => ({
    //
  }),

  mounted() {
    document.title = 'ThesorTeX - Info';
  }
});
</script>
