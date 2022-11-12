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

      <v-btn block color="primary" style="margin-bottom: 8px">
        {{$t(i18nDictionary.Login.Login)}}
      </v-btn>
      <v-btn block text color="primary" style="margin-bottom: 16px">
        {{$t(i18nDictionary.Login.Forgot)}}
      </v-btn>
      <h3 style="margin-bottom: 8px">{{$t(i18nDictionary.Login.NotRegistered)}}</h3>
      <v-btn block text color="primary" to="/register">
        {{$t(i18nDictionary.Common.Register)}}
      </v-btn>
      <v-btn block text color="primary">
        {{$t(i18nDictionary.Login.ContinueWithoutLogin)}}
      </v-btn>

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