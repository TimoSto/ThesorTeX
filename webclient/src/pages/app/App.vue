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

      <AppBarBreadcrumb
        :items="['ThesorTeX', 'Test', 'T2']"
        v-on:clicked="goBackTo"
        />

    </v-app-bar>

    <v-navigation-drawer
      app
      v-model="navDrawer"
      :mini-variant="navDrawer"
      clipped
      permanent
      style="z-index: 100"
      >

      <div class="vert-flex">
        <div class="grow">

        </div>
        <div class="fix">
          <LogoSVG
              style="height: 35px"
          />
        </div>
      </div>

    </v-navigation-drawer>

  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import LogoSVG from "../../common/components/LogoSVG.vue";
import AppBarBreadcrumb from "./components/AppBarBreadcrumb.vue";

export default Vue.extend({
  name: 'App',
  components: {AppBarBreadcrumb, LogoSVG},
  data: () => ({
    navDrawer: false
  }),

  mounted() {
    document.title = 'ThesorTeX';
  },

  methods: {
    goBackTo(item: string) {
      console.log(item)
    }
  }
});
</script>

<style lang="scss" scoped>
@import '../../../node_modules/typeface-roboto/index.css';
@import '@/common/styles/theme_fixes.scss';
@import "@/common/styles/vert-flex.scss";

@include vertical-flex(64px, 56px, 960px, 50px);


</style>
