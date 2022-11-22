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
        :items="pages"
        v-on:clicked="goBackTo"
        />

      <v-spacer />

      <v-btn icon
        @click="configOpened = true"
      >
        <v-icon>mdi-cog</v-icon>
      </v-btn>

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

    <v-main>
      <NavArea
        ref="navArea"
        v-on:clicked="goBackTo"
        :pages="pages"
        >
        <template v-for="p in pages" v-slot:[p] >
          <!--todo: min-width is overflowed and hidden in .page, so it is still not shown, but maybe this can be done better-->
          <div
              :key="`slot_${p}`"
              style="min-width: 150px"
          >

            <StartPage v-if="p === 'ThesorTeX'" />

          </div>
        </template>
      </NavArea>
    </v-main>

    <ConfigDialog
      :open="configOpened"
      v-on:stateChange="configOpened = $event"
    />

  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import LogoSVG from "../../common/components/LogoSVG.vue";
import AppBarBreadcrumb, { Item } from "./components/AppBarBreadcrumb.vue";
import NavArea from "./components/NavArea.vue";
import StartPage from "./views/StartPage.vue";
import ConfigDialog from "./views/ConfigDialog.vue";

export default Vue.extend({
  name: 'App',
  components: {ConfigDialog, StartPage, NavArea, AppBarBreadcrumb, LogoSVG},
  data: () => ({
    navDrawer: false,
    pages: ["ThesorTeX"] as string[],
    configOpened: false
  }),

  mounted() {
    document.title = 'ThesorTeX';
  },

  methods: {
    goBackTo(item: Item) {
      const i = this.pages.indexOf(item.text);
      if( i > -1 ) {
        this.pages = this.pages.slice(0, i+1);
      }
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
