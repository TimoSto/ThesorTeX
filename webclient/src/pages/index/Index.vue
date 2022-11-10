<template>
  <v-app>
    <LogoSVG
        fill="black"
        style="width: 150px"
    />
    <h2>ThesorTeX</h2>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import LogoSVG from "../../components/LogoSVG.vue";

export default Vue.extend({
  name: 'App',
  components: {LogoSVG},
  data: () => ({
    //
  }),

  mounted() {
    document.title = 'ThesorTeX - Info';
  }
});
</script>
