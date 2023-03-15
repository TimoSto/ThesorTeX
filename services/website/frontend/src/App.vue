<template>
  <v-app>
    <router-view :small-display="smallDisplay" />
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
