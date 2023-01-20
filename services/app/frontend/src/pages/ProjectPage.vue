<template>
  <ToolbarAndContent>
    <template #bar>
      <v-toolbar-title>
        {{ projectName }}
      </v-toolbar-title>
    </template>
    <template #content>
      <v-card elevation="1">
        <v-card-text style="padding-bottom: 8px;">
          <v-expansion-panels :model-value="0">
            <v-expansion-panel>
              <v-expansion-panel-title>
                {{ t(i18nKeys.ProjectPage.Entries) }}
              </v-expansion-panel-title>
              <v-expansion-panel-text>
                <ResponsiveTable
                  :rows="entriesRows"
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
                {{ t(i18nKeys.ProjectPage.Categories) }}
              </v-expansion-panel-title>
              <v-expansion-panel-text>
                <ResponsiveTable
                  :rows="categoriesRows"
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
import ResponsiveTable, {
  ResponsiveTableCell,
  ResponsiveTableHeaderCell,
  SizeClasses
} from "../components/ResponsiveTable.vue";
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";
import GetProjectData from "../api/projectData/GetProjectData";
import {useProjectDataStore} from "../stores/projectData/ProjectDataStore";
import {Entry} from "../domain/entry/Entry";
import {Category, Field} from "../domain/category/Category";
import GenerateEntryForCategory from "../domain/category/GenerateEntry";
import {useErrorSuccessStore} from "@thesortex/vue-component-library/src/stores/ErrorSuccessStore/ErrorSuccessStore";

// globals
const appStateStore = useAppStateStore();

const projectDataStore = useProjectDataStore();

const errorSuccessStore = useErrorSuccessStore();

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

const entriesRows = computed(() => {
  const rows = [] as ResponsiveTableCell[][];
  projectDataStore.entries.forEach((e: Entry) => {
    rows.push([
      {
        content: e.Key
      },
      {
        content: e.Category
      },
      {
        content: GenerateEntryForCategory(projectDataStore.categories.find(c => c.Name === e.Category), e.Fields)
      }
    ])
  });

  return rows;
});

const categoriesRows = computed(() => {
  const rows = [] as ResponsiveTableCell[][];
  projectDataStore.categories.forEach((c: Category) => {
    rows.push([
      {
        content: c.Name
      },
      {
        content: GenerateEntryForCategory(c, c.BibFields.map((f: Field) => f.Name))
      }
    ])
  });

  return rows;
})

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
  if (resp.Ok) {
    projectDataStore.setProjectData(resp.Data.Entries, resp.Data.Categories);
  } else {
    errorSuccessStore.setMessage(false, t(i18nKeys.ProjectPage.ErrorReadingData))
  }
}

// watchers
watch(projectName, async () => {//TODO: reload on close and open again
  await syncProjectData();
})

// onload
syncProjectData();
</script>

<style lang="scss">
@import "src/styles/common";

@include disableOverlayOnExpansionPanel;
</style>