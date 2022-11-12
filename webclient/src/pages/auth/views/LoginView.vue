<template>
  <div>
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
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import MaskedTextField from "../components/MaskedTextField.vue";

export default Vue.extend({
  name: "LoginView",
  components: {MaskedTextField}
})
</script>

<style scoped>

</style>