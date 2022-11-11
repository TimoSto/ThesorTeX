<template>
  <v-app style="background-color: var(--v-background-darken2)">
    <v-card class="centered">
      <LogoSVG
          fill="rgba(0, 136, 52, 0.8)"
          style="height: 45px; margin: 0 auto; display: block"
          font-color="rgba(0,0,0,0.7)"
      />
      <v-card-title
        style="display: block; text-align: center"
      >
        ThesorTeX - Anmeldung
      </v-card-title>
      <v-card-text>
        <v-text-field
          label="Benutzername"
        />

        <MaskedTextField
          label="Passwort"
        />

      </v-card-text>
    </v-card>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import LogoSVG from "../../commonComponents/LogoSVG.vue";
import MaskedTextField from "./components/MaskedTextField.vue";

export default Vue.extend({
  name: 'App',
  components: {MaskedTextField, LogoSVG},
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
