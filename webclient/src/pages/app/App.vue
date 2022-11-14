<template>
  <v-app>

    <v-app-bar
      color="primary"
      fixed
      clipped-left
      app
      elevation="0"
    >

      <v-app-bar-nav-icon
        @click="navDrawer = !navDrawer"
      />

      <v-toolbar-title>
        ThesorTeX
      </v-toolbar-title>

    </v-app-bar>

    <v-navigation-drawer
      app
      v-model="navDrawer"
      :mini-variant="navDrawer"
      clipped
      permanent
      style="z-index: 100"
      >

    </v-navigation-drawer>

  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'App',

  data: () => ({
    navDrawer: false
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
