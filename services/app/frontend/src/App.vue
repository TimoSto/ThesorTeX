<template>
  <v-app>
    <v-app-bar
      color="primary"
      elevation="0"
    >
      <v-app-bar-nav-icon
        :disabled="sidebarDisabled"
        @click.stop="sidebarOpened = !sidebarOpened"
      />

      <v-app-bar-nav-icon
        v-if="pagesCount > 1"
        @click="navBack"
      >
        <v-icon>mdi-arrow-left</v-icon>
      </v-app-bar-nav-icon>

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
    <v-main>
      <PageNavigator
        :pages="pagesCount"
        :instant-switch="false"
        :navigating-back="navigatingBack"
        @nav-back-finish="finishNavBack"
      >
        <template
          v-for="i in pagesCount"
          #[i]
        >
          <MainPage
            v-if="i === 1"
            :key="`page-${i}`"
          />

          <ProjectPage
            v-if="i === 2"
            :key="`page-${i}`"
          />

          <CategoryEditor
            v-if="i === 3"
            :key="`page-${i}`"
          />
        </template>
      </PageNavigator>
    </v-main>
    <ErrorSuccessDisplay
      :valid="success"
      :message="message"
      :error-title="t(i18nKeys.Common.Error)"
      :close="t(i18nKeys.Common.Close)"
      @close="message = ''"
    >
      <template #suffix>
        <i18n-t :keypath="i18nKeys.Common.ContactBug">
          <template #link>
            <a
              href="https://github.com/TimoSto/ThesorTeX/issues"
              class="text-primary"
              target="_blank"
            >
              https://github.com/TimoSto/ThesorTeX/issues
            </a>
          </template>
        </i18n-t>
      </template>
    </ErrorSuccessDisplay>
  </v-app>
</template>

<script lang="ts" setup>
import {pageNames, useAppStateStore} from "./stores/appState/AppStateStore";
import {computed,} from "vue";
import PageNavigator from "./components/PageNavigator.vue";
import {ErrorSuccessDisplay} from "@thesortex/vue-component-library/src/components";
import MainPage from "./pages/MainPage.vue";
import {useErrorSuccessStore} from "@thesortex/vue-component-library/src/stores/ErrorSuccessStore/ErrorSuccessStore";
import {i18nKeys} from "./i18n/keys";
import {useI18n} from "@thesortex/vue-i18n-plugin"
import ProjectPage from "./pages/ProjectPage.vue";
import CategoryEditor from "./pages/CategoryEditor.vue";

//globals
const appStateStore = useAppStateStore();

const errorSuccessStore = useErrorSuccessStore();

const {t} = useI18n();

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
  return appStateStore.history.length === 1;
})

const pagesCount = computed(() => {
  console.log(appStateStore.history.length)
  return appStateStore.history.length;
})

const titleAppendix = computed(() => {
  let appendix: string;
  switch (appStateStore.currentPage) {
    case pageNames[0]:
      appendix = "";
      break;
    case pageNames[1]:
      appendix = t(i18nKeys.ProjectPage.Title);
      break;
    case pageNames[2]:
      appendix = t(i18nKeys.CategoryEditor.Title);
      break;
    default:
      appendix = "";
  }
  return appendix
});

const success = computed(() => {
  return errorSuccessStore.valid;
})

const message = computed({
  get(): string {
    return errorSuccessStore.message;
  },
  set(v: string) {
    if (v === "") {
      errorSuccessStore.clear();
    }
  }
});

const navigatingBack = computed(() => {
  return appStateStore.navigatingBack;
})

// methods
function navBack() {
  appStateStore.goBack();
}

function finishNavBack() {
  appStateStore.finishGoBack();
}

</script>

<style scoped>

</style>