<template>
  <v-app style="background-image: linear-gradient(90deg, #0c8635, #259b71, #69beaf);">
    <v-app-bar :color="elevation === 0 ? 'transparent' : 'background'" :elevation="elevation" app>
      <v-container class="pa-6 fill-height">
        <v-row align="center" class="fill-height">
          <v-toolbar-title class="text-h4 font-weight-bold" :style="`color: ${elevation === 0 ? 'white' : ''}`">
            ThesorTeX {{ titleAppendix }}
          </v-toolbar-title>
          <v-btn :color="elevation === 0 ? 'white' : 'primary'" to="/" v-if="currentPage !== 'Home'">
            {{ t(i18nKeys.Titles.StartPage) }}
          </v-btn>
          <v-btn :color="elevation === 0 ? 'white' : 'primary'" to="/downloads" v-if="currentPage !== 'Downloads'">
            {{ t(i18nKeys.Titles.Downloads) }}
          </v-btn>
          <v-btn :color="elevation === 0 ? 'white' : 'primary'" to="/tutorials" v-if="currentPage !== 'Tutorials'">
            {{ t(i18nKeys.Titles.Tutorials) }}
          </v-btn>
        </v-row>
      </v-container>

    </v-app-bar>
    <v-main app>
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
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "./i18n/keys";

const router = useRouter();

const {t} = useI18n();

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

const titleAppendix = computed(() => {
  if (currentPage.value === "Tutorials") {
    return `- ${t(i18nKeys.Titles.Tutorials)}`;
  }
  if (currentPage.value === "Downloads") {
    return `- ${t(i18nKeys.Titles.Downloads)}`;
  }

  return "";
});

</script>

<style scoped lang="scss">

</style>

<style>
a {
  color: rgb(var(--v-theme-primary));
}
</style>
