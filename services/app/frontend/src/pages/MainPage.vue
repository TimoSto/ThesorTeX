<template>
  <ToolbarAndContent>
    <template #bar>
      <v-toolbar-title>
        {{ t(i18nKeys.MainPage.Welcome) }}
      </v-toolbar-title>
    </template>
    <template #content>
      <div class="fullsize-card-container">
        <v-card elevation="3">
          <v-card-text>
            <ResponsiveTable
              :headers="projectHeaders"
              :rows="projectsRows"
              @row-clicked="OpenProject"
            >
              <template #h-4>
                <v-btn
                  flat
                  text
                  @click="createNewTriggered = true"
                >
                  <v-icon>mdi-plus</v-icon>
                </v-btn>
              </template>
            </ResponsiveTable>
          </v-card-text>
        </v-card>
      </div>
    </template>
  </ToolbarAndContent>

  <v-dialog
    v-model="createNewTriggered"
    width="450"
  >
    <CreateProjectCard
      :projects="existingProjects"
      @close="createNewTriggered = false"
      @confirm="triggerProjectCreation($event)"
    />
  </v-dialog>
</template>

<script lang="ts" setup>
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";
import ResponsiveTable, {ResponsiveTableHeaderCell, SizeClasses} from "../components/ResponsiveTable.vue";
import {useProjectsListStore} from "../stores/projectsList/ProjectsListStore";
import {computed, ref} from "vue";
import ProjectMetaData from "../domain/projects/ProjectMetaData";
import CreateProjectCard from "../components/CreateProjectCard.vue";
import CreateNewProject from "../api/projects/CreateNewProject";
import {useErrorSuccessStore} from "@thesortex/vue-component-library/src/stores/ErrorSuccessStore/ErrorSuccessStore";
import {pageNames, useAppStateStore} from "../stores/appState/AppStateStore";
import GetProjectsMetaData from "../api/projects/GetProjectsMetaData";

// globals
const {t} = useI18n();

const projectsStore = useProjectsListStore();

const errorSuccessStore = useErrorSuccessStore();

const appStateStore = useAppStateStore();

// data

const projectHeaders: ResponsiveTableHeaderCell[] = [
  {
    content: t(i18nKeys.MainPage.Project),
    size: SizeClasses.TwentyFivePercent
  },
  {
    content: t(i18nKeys.MainPage.ProjectCreated),
    size: SizeClasses.TwentyFivePercent
  },
  {
    content: t(i18nKeys.MainPage.ProjectLastModified),
    size: SizeClasses.TwentyFivePercent
  },
  {
    content: t(i18nKeys.MainPage.ProjectNumberOfEntries),
    size: SizeClasses.TenPercent
  },
  {
    slot: true,
    size: SizeClasses.IconBtn
  },
];

const createNewTriggered = ref(false);

// computed
const projectsRows = computed(() => {
  return projectsStore.projects.map((p: ProjectMetaData) => [
    {
      content: p.Name,
    },
    {
      content: p.Created,
    },
    {
      content: p.LastEdited,
    },
    {
      content: p.NumberOfEntries,
    },
    {
      content: ""
    }
  ])
})

const existingProjects = computed(() => {
  return projectsStore.projects.map(p => p.Name)
})

// methods
async function triggerProjectCreation(name: string) {
  const resp = await CreateNewProject(name);

  createNewTriggered.value = false;

  if (resp.Success) {
    projectsStore.addProject(resp.Data);
    errorSuccessStore.setMessage(true, t(i18nKeys.MainPage.SuccessCreation).replace("NAME", name));
  } else {
    let appendix = "";
    switch (resp.Status) {
      case 400:
        appendix = t(i18nKeys.Common.Error400);
        break;
      case 404:
        appendix = t(i18nKeys.Common.Error404);
        break;
      case 500:
        appendix = t(i18nKeys.Common.Error500);
        break;
    }
    errorSuccessStore.setMessage(false, `${t(i18nKeys.MainPage.ErrorCreation)} ${appendix}`);
  }
}

function OpenProject(n: number) {
  if (appStateStore.currentPage === pageNames[0]) {
    appStateStore.navToPage(pageNames[1]);
    appStateStore.setProject(projectsStore.projects[n].Name);
  }
}

async function syncProjectsListWithServer() {
  const resp = await GetProjectsMetaData();
  if (resp.Ok) {
    projectsStore.setProjects(resp.Projects ? resp.Projects : []);
  } else {
    // todo handle error
  }
}

// onload
syncProjectsListWithServer();
</script>

<style scoped lang="scss">
@import "../styles/common.scss";

@include fullsizeCardContainer;
</style>