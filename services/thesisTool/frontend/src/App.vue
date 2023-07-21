<template>
  <ApplicationFrame
    mainTitle="ThesorTeX"
    :titleAppendix="titleAppendix"
    :hasSidebar="!sidebarDisabled"
  >
    <template #sidebar>
      <ProjectsSidebar
        v-if="pagesCount > 1"
        :project="appStateStore.currentProject"
        :open="sidebarOpened"
        @switch-to="switchToProject($event)"
      />
    </template>
    <template
      v-for="i in pagesCount"
      #[i-1]="{openPage, goBack}"
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
        v-if="i === 3 && currentPage === pageNames[2]"
        :key="`page-${i}`"
      />

      <EntryEditor
        v-if="i === 3 && currentPage === pageNames[3]"
        :key="`page-${i}`"
      />
    </template>
  </ApplicationFrame>
  <!--  <v-app>-->
  <!--    <v-app-bar-->
  <!--      color="primary"-->
  <!--      elevation="0"-->
  <!--    >-->
  <!--      <v-app-bar-nav-icon-->
  <!--        :disabled="sidebarDisabled"-->
  <!--        :title="sidebarOpened ? t(i18nKeys.App.CloseSidebar) : t(i18nKeys.App.OpenSidebar)"-->
  <!--        @click.stop="sidebarOpened = !sidebarOpened"-->
  <!--      />-->

  <!--      <v-app-bar-nav-icon-->
  <!--        v-if="pagesCount > 1"-->
  <!--        :title="t(i18nKeys.Common.Back)"-->
  <!--        @click="navBack"-->
  <!--      >-->
  <!--        <v-icon>mdi-arrow-left</v-icon>-->
  <!--      </v-app-bar-nav-icon>-->

  <!--      <v-app-bar-title role="heading" aria-level="1">-->
  <!--        ThesorTeX-->
  <!--        {{ titleAppendix }}-->
  <!--      </v-app-bar-title>-->

  <!--      <v-spacer />-->

  <!--      <a target="_blank" href="https://thesortex.com/#/tutorials?target=ThesisTool"-->
  <!--         style="color: rgb(var(&#45;&#45;v-theme-on-primary))">-->
  <!--        <v-btn icon :title="t(i18nKeys.App.GoToDocs)">-->
  <!--          <v-icon>mdi-book-open-variant</v-icon>-->
  <!--        </v-btn>-->
  <!--      </a>-->
  <!--      <AccessibilityDialog v-if="myDocument.readyState" :keydownTarget="myDocument" :i18n="t" v-slot="scope">-->
  <!--        <v-btn icon-->
  <!--               @click="scope.openDialog" :title="t(accessibilityDialogKeys.AccessibilityDialog.BtnTitle)">-->
  <!--          <v-icon>mdi-human</v-icon>-->
  <!--        </v-btn>-->
  <!--      </AccessibilityDialog>-->
  <!--      <v-btn icon :title="t(i18nKeys.Config.OpenConfig)" @click="configOpened=true">-->
  <!--        <v-icon>mdi-cog</v-icon>-->
  <!--      </v-btn>-->
  <!--    </v-app-bar>-->

  <!--    <v-navigation-drawer-->
  <!--      permanent-->
  <!--      :rail="!sidebarOpened"-->
  <!--      :rail-width="56"-->
  <!--    >-->
  <!--      -->
  <!--    </v-navigation-drawer>-->
  <!--    <v-main>-->
  <!--      <PageNavigator-->
  <!--        :pages="pagesCount"-->
  <!--        :instant-switch="instantSwitch"-->
  <!--        :navigating-back="navigatingBack"-->
  <!--        @nav-back-finish="finishNavBack"-->
  <!--      >-->
  <!--        -->
  <!--      </PageNavigator>-->
  <!--    </v-main>-->
  <!--    <ErrorSuccessDisplay-->
  <!--      :valid="success"-->
  <!--      :message="message"-->
  <!--      :error-title="t(i18nKeys.Common.Error)"-->
  <!--      :close="t(i18nKeys.Common.Close)"-->
  <!--      @close="message = ''"-->
  <!--    >-->
  <!--      <template #suffix>-->
  <!--        <i18n-t :keypath="i18nKeys.Common.ContactBug">-->
  <!--          <template #link>-->
  <!--            <a-->
  <!--              href="https://github.com/TimoSto/ThesorTeX/issues"-->
  <!--              class="text-primary"-->
  <!--              target="_blank"-->
  <!--            >-->
  <!--              https://github.com/TimoSto/ThesorTeX/issues-->
  <!--            </a>-->
  <!--          </template>-->
  <!--        </i18n-t>-->
  <!--      </template>-->
  <!--    </ErrorSuccessDisplay>-->
  <!--    <v-dialog-->
  <!--      v-model="unsaveDialogOpened"-->
  <!--      width="450"-->
  <!--    >-->
  <!--      <UnsavedChangesDialog-->
  <!--        @resolve="appStateStore.resolveCallback($event)"-->
  <!--      />-->
  <!--    </v-dialog>-->
  <!--    <ConfigDialog-->
  <!--      :open="configOpened"-->
  <!--      @close="configOpened=false"-->
  <!--    />-->
  <!--  </v-app>-->
</template>

<script lang="ts" setup>
import {pageNames, useAppStateStore} from "./stores/appState/AppStateStore";
import {computed, onMounted, ref, watch,} from "vue";
import PageNavigator from "./components/PageNavigator.vue";
import {ErrorSuccessDisplay} from "@thesortex/vue-component-library/src/components";
import MainPage from "./pages/MainPage.vue";
import {useErrorSuccessStore} from "@thesortex/vue-component-library/src/stores/ErrorSuccessStore/ErrorSuccessStore";
import {i18nKeys} from "./i18n/keys";
import {useI18n} from "@thesortex/vue-i18n-plugin";
import ProjectPage from "./pages/ProjectPage.vue";
import CategoryEditor from "./pages/CategoryEditor.vue";
import EntryEditor from "./pages/EntryEditor.vue";
import ProjectsSidebar from "./components/ProjectsSidebar.vue";
import {useProjectsListStore} from "./stores/projectsList/ProjectsListStore";
import UnsavedChangesDialog from "./components/UnsavedChangesCard.vue";
import ConfigDialog from "./pages/ConfigDialog.vue";
import {accessibilityDialogKeys} from "@thesortex/vue-component-library";
import {
  useApplicationStateStore
} from "@thesortex/vue-component-library/src/stores/ApplicationStateStore/ApplicationStateStore";

//globals
const appStateStore = useAppStateStore();

const errorSuccessStore = useErrorSuccessStore();

const projectsListStore = useProjectsListStore();

const applicationStateStore = useApplicationStateStore();

const {t} = useI18n();

// data
const instantSwitch = ref(false);

const configOpened = ref(false);

const myDocument = ref(document);

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
  return applicationStateStore.history.length === 1;
});

const pagesCount = computed(() => {
  return applicationStateStore.history.length;
});

const currentPage = computed(() => applicationStateStore.currentPage);

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
    case pageNames[3]:
      appendix = t(i18nKeys.EntryEditor.Title);
      break;
    default:
      appendix = "";
  }
  return appendix;
});

const success = computed(() => {
  return errorSuccessStore.valid;
});

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
});

const unsaveDialogOpened = computed({
  get(): boolean {
    return appStateStore.unsavedDialogTriggered;
  },
  set(v: boolean) {
    if (!v) {
      appStateStore.resolveCallback(false);
    }
  }
});

// watcher
watch(() => appStateStore.unsavedDialogTriggered, () => {
  if (!appStateStore.unsavedDialogTriggered) {
    setTimeout(() => {
      instantSwitch.value = false;
    }, 0);
  }
});

// methods
function navBack() {
  appStateStore.goBack();
}

function finishNavBack() {
  appStateStore.finishGoBack();
}

function switchToProject(n: number) {
  instantSwitch.value = true;
  appStateStore.switchToProject(projectsListStore.projects[n].Name);
  if (!appStateStore.unsavedChanges) {
    setTimeout(() => {
      instantSwitch.value = false;
    }, 0);
  }
}

onMounted(() => {
  console.log(appStateStore.history.length);
});

</script>

<style scoped>

</style>