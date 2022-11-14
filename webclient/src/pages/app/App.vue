<template>
  <v-app>

    <v-app-bar
      color="primary"
      fixed
      clipped-left
    >

      <v-app-bar-nav-icon />

      <v-toolbar-title>
        ThesorTeX
      </v-toolbar-title>

    </v-app-bar>

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
    document.title = 'ThesorTeX';
  }
});
</script>

<style lang="scss" scoped>
@import '../../../node_modules/typeface-roboto/index.css';
@import '@/common/styles/theme_fixes.scss';

</style>
