<template>
  <v-app>
    <v-sheet>
      <v-app-bar color="background" :elevation="elevation" app>
        <v-container class="pa-6 fill-height">
          <v-row align="center" class="fill-height">
            <v-toolbar-title class="text-h4 font-weight-bold">ThesorTeX</v-toolbar-title>
            <v-spacer />
            <v-btn color="primary" to="/" v-if="currentPage !== 'Home'">
              Übersicht
            </v-btn>
            <v-btn color="primary" to="/downloads" v-if="currentPage !== 'Downloads'">
              Downloads
            </v-btn>
            <v-btn color="primary" to="/tutorials" v-if="currentPage !== 'Tutorials'">
              Tutorials
            </v-btn>
          </v-row>
        </v-container>

      </v-app-bar>
    </v-sheet>
    <v-main>
      <ToolbarAndContent :hide-bar="true" @scroll="elevation=1" @no-scroll="elevation=0" :page="currentPage">
        <template #content>
          <router-view :small-display="smallDisplay" />
        </template>
      </ToolbarAndContent>
    </v-main>
  </v-app>
</template>

<script lang="ts" setup>
import {computed, ref} from "vue";
import {useRouter} from "vue-router";

const router = useRouter();

// data
const elevation = ref(0);
const smallDisplay = ref(false);

// onload
window.addEventListener("resize", () => {
  if (window.innerWidth < 750) {
    smallDisplay.value = true;
  } else {
    smallDisplay.value = false;
  }
});

const currentPage = computed(() => {
  return router.currentRoute.value.name;
});

</script>

<style scoped lang="scss">

</style>

<style>
a {
  color: rgb(var(--v-theme-primary));
}
</style>
