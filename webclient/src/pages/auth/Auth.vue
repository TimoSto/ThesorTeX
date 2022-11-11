<template>
  <v-app style="background-color: var(--v-background-darken2)">
    <v-card class="centered">
      <LogoSVG
          fill="rgba(0, 136, 52, 0.8)"
          style="height: 35px"
          :show-title="true"
          font-color="rgba(0,0,0,0.7)"
      />
    </v-card>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import LogoSVG from "../../commonComponents/LogoSVG.vue";

export default Vue.extend({
  name: 'App',
  components: {LogoSVG},
  data: () => ({
    //
  }),

  mounted() {
    document.title = 'ThesorTeX - Login';
  }
});
</script>

<style scoped>
@import '../../../node_modules/typeface-roboto/index.css';

.centered {
  margin: 0;
  position: absolute;
  top: 30%;
  left: 50%;
  -ms-transform: translateY(-30%) translateX(-50%);
  transform: translateY(-30%) translateX(-50%);
  width: 300px;
  background-color: var(--v-background-base);
}
</style>
