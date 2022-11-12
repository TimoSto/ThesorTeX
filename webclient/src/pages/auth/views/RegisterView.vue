<template>
  <div>
    <v-card-title
        style="display: block; text-align: center"
    >
      {{$t(i18nDictionary.Register.Title)}}
    </v-card-title>
    <v-card-text>
      <v-text-field
          :label="$t(i18nDictionary.Common.EMail)"
      />

      <MaskedTextField
          :label="$t(i18nDictionary.Common.Password)"
      />

      <MaskedTextField
          :label="$t(i18nDictionary.Register.RepeatPassword)"
      />

      <v-btn block color="primary" style="margin-bottom: 8px">
        {{$t(i18nDictionary.Common.Register)}}
      </v-btn>
      <v-btn block text color="primary">
        {{$t(i18nDictionary.Login.Login)}}
      </v-btn>

    </v-card-text>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import {i18nDictionary} from "../i18n/Keys";
import MaskedTextField from "../components/MaskedTextField.vue";

export default Vue.extend({
  name: "RegisterView",
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