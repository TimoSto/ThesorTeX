<template>
  <ToolbarAndContent>
    <template #bar>
      <v-toolbar-title>
        {{ projectName }}
      </v-toolbar-title>
    </template>
    <template #content>
      <v-card elevation="0">
        <v-card-text style="padding-bottom: 8px;">
          <v-expansion-panels :model-value="0">
            <v-expansion-panel>
              <v-expansion-panel-title>
                {{ t(i18nKeys.ProjectPage.Entries) }}
              </v-expansion-panel-title>
              <v-expansion-panel-text>
                <ResponsiveTable
                  :rows="[]"
                  :headers="entriesHeaders"
                >
                  <template #h-3>
                    <v-btn
                      flat
                      text
                    >
                      <v-icon>mdi-plus</v-icon>
                    </v-btn>
                  </template>
                </ResponsiveTable>
              </v-expansion-panel-text>
            </v-expansion-panel>
          </v-expansion-panels>
        </v-card-text>
      </v-card>
      <v-card elevation="0">
        <v-card-text style="padding-top: 8px;">
          <v-expansion-panels :model-value="0">
            <v-expansion-panel>
              <v-expansion-panel-title>
                {{ t(i18nKeys.ProjectPage.EntryCategory) }}
              </v-expansion-panel-title>
              <v-expansion-panel-text>
                <ResponsiveTable
                  :rows="[]"
                  :headers="categoriesHeaders"
                >
                  <template #h-2>
                    <v-btn
                      flat
                      text
                    >
                      <v-icon>mdi-plus</v-icon>
                    </v-btn>
                  </template>
                </ResponsiveTable>
              </v-expansion-panel-text>
            </v-expansion-panel>
          </v-expansion-panels>
        </v-card-text>
      </v-card>
    </template>
  </ToolbarAndContent>

</template>

<script lang="ts" setup>
import {useAppStateStore} from "../stores/appState/AppStateStore";
import {computed, watch} from "vue";
import ResponsiveTable, {ResponsiveTableHeaderCell, SizeClasses} from "../components/ResponsiveTable.vue";
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";
import GetProjectData from "../api/projectData/GetProjectData";
import {useProjectDataStore} from "../stores/projectData/ProjectDataStore";

// globals
const appStateStore = useAppStateStore();

const projectDataStore = useProjectDataStore();

const {t} = useI18n();

// computed
const projectName = computed(() => {
  return appStateStore.currentProject;
});

const entriesHeaders = computed((): ResponsiveTableHeaderCell[] => {
  return [
    {
      content: t(i18nKeys.ProjectPage.EntryKey),
      size: SizeClasses.MaxWidth250
    },
    {
      content: t(i18nKeys.ProjectPage.EntryCategory),
      size: SizeClasses.MaxWidth250
    },
    {
      content: t(i18nKeys.ProjectPage.Entry),
      size: SizeClasses.ThirtyPercent
    },
    {
      slot: true,
      size: SizeClasses.IconBtn
    },
  ]
});

const categoriesHeaders = computed((): ResponsiveTableHeaderCell[] => {
  return [
    {
      content: t(i18nKeys.ProjectPage.CategoryName),
      size: SizeClasses.MaxWidth250
    },
    {
      content: t(i18nKeys.ProjectPage.CategoryExample),
      size: ""
    },
    {
      slot: true,
      size: SizeClasses.IconBtn
    },
  ]
});

async function syncProjectData() {
  const resp = await GetProjectData(projectName.value);
  if( resp.Ok ) {
    console.log(resp.Data.Entries)
    projectDataStore.setProjectData(resp.Data.Entries);
  }
}

// watchers
watch(projectName, async () => {
  await syncProjectData();
})

// onload
syncProjectData();
</script>

<style scoped>

</style>