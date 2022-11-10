<template>
  <v-app>
    <v-img
        src="../../../public/logo.svg"
        width="150"
        height="150"
        style="max-height: 150px"
        contain
    />
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'App',

  data: () => ({
    //
  }),

  mounted() {
    document.title = 'ThesorTeX - Info';
  }
});
</script>
