<template>
  <v-app>
    <v-main>
      <v-app-bar
        color="primary"
        elevation="0"
      >
        <v-app-bar-nav-icon
          :disabled="sidebarDisabled"
          @click.stop="sidebarOpened = !sidebarOpened"
        />

        <v-app-bar-title>
          ThesorTeX
          {{ titleAppendix }}
        </v-app-bar-title>

        <v-spacer />
      </v-app-bar>

      <v-navigation-drawer
        permanent
        :rail="!sidebarOpened"
        :rail-width="68"
      >
        <!--Sidebar content-->
      </v-navigation-drawer>
    </v-main>
  </v-app>
</template>

<script lang="ts" setup>
import {pageNames, useAppStateStore} from "./stores/appState/AppStateStore";
import {computed} from "vue";

//globals

const appStateStore = useAppStateStore();

// data

// computed
const sidebarOpened = computed({
  get() {
    return appStateStore.sidebarOpen;
  },
  set(v: boolean) {
    appStateStore.setSidebarOpened(v);
  }
});

const sidebarDisabled = computed(() => {
  return appStateStore.currentPage === pageNames[0];
})

const titleAppendix = computed(() => {
  let appendix: string;
  switch (appStateStore.currentPage) {
    case pageNames[0]:
      appendix = "";
      break;
    default: appendix = "";
  }
  return appendix
})

</script>

<style scoped>

</style>