<template>
  <div>
    <v-card-title
        style="display: block; text-align: center"
    >
      {{$t(i18nDictionary.Login.Title)}}
    </v-card-title>
    <v-card-text>
      <v-text-field
          :label="$t(i18nDictionary.Common.EMail)"
      />

      <MaskedTextField
          :label="$t(i18nDictionary.Common.Password)"
      />

    </v-card-text>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import MaskedTextField from "../components/MaskedTextField.vue";
import {i18nDictionary} from "../i18n/Keys";

export default Vue.extend({
  name: "LoginView",
  components: {MaskedTextField},
  data() {
    return {
      i18nDictionary: i18nDictionary
    }
  }
})
</script>

<style scoped>

</style>