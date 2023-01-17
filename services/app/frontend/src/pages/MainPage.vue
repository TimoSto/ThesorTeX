<template>
  <ToolbarAndContent>
    <template #bar>
      <v-toolbar-title>
        {{ t(i18nKeys.MainPage.Welcome) }}
      </v-toolbar-title>
    </template>
    <template #content>
      <div style="padding: 8px 16px">
        <ResponsiveTable
          :headers="projectHeaders"
          :rows="projectsRows"
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
    />
  </v-dialog>
</template>

<script lang="ts" setup>
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";
import ResponsiveTable, {ResponsiveTableHeaderCell, SizeClasses} from "../components/ResponsiveTable.vue";
import {useProjectsListStore} from "../stores/appState/ProjectsListStore";
import {computed, ref} from "vue";
import ProjectMetaData from "../domain/projects/ProjectMetaData";
import CreateProjectDialog from "../components/CreateProjectCard.vue";
import CreateProjectCard from "../components/CreateProjectCard.vue";

// globals
const { t } = useI18n();

const projectsStore = useProjectsListStore();

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

// onload
projectsStore.readAllProjectsFromServer();
</script>

<style scoped>
</style>