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
        >
        <template v-for="p in pages" v-slot:[p] >
          Hallo
        </template>
      </NavArea>
    </v-main>

  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import LogoSVG from "../../common/components/LogoSVG.vue";
import AppBarBreadcrumb from "./components/AppBarBreadcrumb.vue";
import NavArea, {NavAreaMethods} from "./components/NavArea.vue";

export default Vue.extend({
  name: 'App',
  components: {NavArea, AppBarBreadcrumb, LogoSVG},
  data: () => ({
    navDrawer: false,
    pages: [] as string[]
  }),

  mounted() {
    document.title = 'ThesorTeX';
    this.$nextTick(() => {
      if( this.$refs.navArea ) {//the ref should act as a proxy, like the dapi
        const area = this.$refs.navArea as NavAreaMethods;
        area.addPage('ThesorTeX');
        this.pages.push('ThesorTeX');
      }
    })
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
