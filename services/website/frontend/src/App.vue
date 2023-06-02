<template>
  <!--TODO: Find better way-->
  <v-app style="overflow-x: clip">
    <div style="position: absolute; width: 100%; z-index: 100;">
      <v-container style="padding-top: 30px; background-color: transparent;">
        <v-row>
          <v-col style="font-size: 30px; font-weight: bold; color: white;">
            ThesorTeX
          </v-col>
          <v-col class="d-flex">
            <v-spacer />
            <v-btn variant="text" color="white" style="font-weight: bold;" v-if="currentPage !== 'Home'" to="/">
              {{ t(i18nKeys.Titles.StartPage) }}
            </v-btn>
            <v-btn variant="text" color="white" style="font-weight: bold;" v-if="currentPage !== 'Downloads'"
                   to="/downloads">
              {{ t(i18nKeys.Titles.Downloads) }}
            </v-btn>
            <v-btn variant="text" color="white" style="font-weight: bold;" v-if="currentPage !== 'Tutorials'"
                   to="/tutorials">
              {{ t(i18nKeys.Titles.Tutorials) }}
            </v-btn>
          </v-col>
        </v-row>
      </v-container>
    </div>
    <router-view :small-display="smallDisplay" />
    <AccessibilityDialog v-if="myDocument.readyState" :keydownTarget="myDocument" title="test" />
  </v-app>
</template>

<script lang="ts" setup>
import {computed, onMounted, ref} from "vue";
import {useRouter} from "vue-router";
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "./i18n/keys";

const router = useRouter();

const {t} = useI18n();

// data
const elevation = ref(0);
const smallDisplay = ref(false);
const myDocument = ref(document);

function onResize() {
  if (screen.width < 900) {
    smallDisplay.value = true;
  } else {
    smallDisplay.value = false;
  }
}

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

onMounted(() => {
  document.title = `ThesorTeX - ${t(i18nKeys.StartPage.Title)}`;
  onResize();
  window.addEventListener("resize", () => {
    onResize();
  });
});

</script>

<style scoped lang="scss">

</style>

<style>
a {
  color: rgb(var(--v-theme-primary));
}
</style>
