<template>
  <ToolbarAndContent>
    <template #bar>
      <v-toolbar-title>
        {{ projectName }}
      </v-toolbar-title>
      <v-spacer />
      <v-btn
        icon
        @click="deleteTriggered = true"
      >
        <v-icon>mdi-delete</v-icon>
      </v-btn>
    </template>
    <template #content>
      <div class="fullsize-card-container">
        <v-card elevation="3">
          <div>
            <v-expansion-panels :model-value="0">
              <v-expansion-panel>
                <v-expansion-panel-title>
                  {{ t(i18nKeys.ProjectPage.Entries) }}
                </v-expansion-panel-title>
                <v-expansion-panel-text>
                  <CitaviDragNDrop
                    :title="t(i18nKeys.ProjectPage.DragNDrop)"
                    :categories="categories"
                    style="margin-bottom: 8px;"
                    @entries-uploaded="handleUpload"
                  />
                  <ResponsiveTable
                    :rows="entriesRows"
                    :headers="entriesHeaders"
                    @row-clicked="openEntryEditor"
                  >
                    <template #h-3>
                      <v-btn
                        flat
                        text
                        @click="openEntryEditor(-1)"
                      >
                        <v-icon>mdi-plus</v-icon>
                      </v-btn>
                    </template>
                  </ResponsiveTable>
                </v-expansion-panel-text>
              </v-expansion-panel>
            </v-expansion-panels>
          </div>
        </v-card>
      </div>
      <div class="fullsize-card-container">
        <v-card elevation="3">
          <div>
            <v-expansion-panels :model-value="0">
              <v-expansion-panel>
                <v-expansion-panel-title>
                  {{ t(i18nKeys.ProjectPage.Categories) }}
                </v-expansion-panel-title>
                <v-expansion-panel-text>
                  <ResponsiveTable
                    :rows="categoriesRows"
                    :headers="categoriesHeaders"
                    @row-clicked="openCategoryEditor"
                  >
                    <template #h-2>
                      <v-btn
                        flat
                        text
                        @click="openCategoryEditor(-1)"
                      >
                        <v-icon>mdi-plus</v-icon>
                      </v-btn>
                    </template>
                  </ResponsiveTable>
                </v-expansion-panel-text>
              </v-expansion-panel>
            </v-expansion-panels>
          </div>
        </v-card>
      </div>
    </template>
  </ToolbarAndContent>
  <v-dialog
    v-model="deleteTriggered"
    width="400"
  >
    <v-card>
      <v-card-title>{{ t(i18nKeys.ProjectPage.DeleteTitle) }}</v-card-title>
      <v-card-text>
        <i18n-t :keypath="i18nKeys.ProjectPage.DeleteMessage">
          <template #project>
            <i>{{ projectName }}</i>
          </template>
        </i18n-t>
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn
          color="primary"
          @click="deleteTriggered=false"
        >
          {{ t(i18nKeys.Common.Abort) }}
        </v-btn>
        <v-btn
          color="primary"
          @click="deleteProject"
        >
          {{ t(i18nKeys.Common.Delete) }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
  <v-dialog
    v-model="uploadTriggered"
    width="400"
  >
    <CitaviUploadCard
      :entries="uploadedEntries"
      :unknowns="uploadedUnknowns"
      @rm-entry="rmUploadedEntry"
      @close="uploadTriggered = false"
      @upload="uploadDroppedEntries"
    />
  </v-dialog>
</template>

<script lang="ts" setup>
import {pageNames, useAppStateStore} from "../stores/appState/AppStateStore";
import {computed, ref, watch} from "vue";
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
import DeleteProject from "../api/projects/DeleteProject";
import {useProjectsListStore} from "../stores/projectsList/ProjectsListStore";
import CitaviDragNDrop from "../components/CitaviDragNDrop.vue";
import {AnalyseResult, Unknown} from "../domain/citavi/BibAnalytics";
import CitaviUploadCard from "../components/CitaviUploadCard.vue";
import UploadEntries from "../api/projectData/uploadEntries";

// globals
const appStateStore = useAppStateStore();

const projectDataStore = useProjectDataStore();

const errorSuccessStore = useErrorSuccessStore();

const projectsListStore = useProjectsListStore();

const {t} = useI18n();

// data
const deleteTriggered = ref(false);

const uploadTriggered = ref(false);

const uploadedEntries = ref([] as Entry[]);

const uploadedUnknowns = ref([] as Unknown[]);

// computed
const projectName = computed(() => {
  return appStateStore.currentProject;
});

const entriesHeaders = computed((): ResponsiveTableHeaderCell[] => {
  return [
    {
      content: t(i18nKeys.ProjectPage.EntryKey),
      size: SizeClasses.MaxWidth200
    },
    {
      content: t(i18nKeys.ProjectPage.EntryCategory),
      size: SizeClasses.MaxWidth150
    },
    {
      content: t(i18nKeys.ProjectPage.Entry),
      size: SizeClasses.ThirtyPercent
    },
    {
      slot: true,
      size: SizeClasses.IconBtn
    },
  ];
});

const entriesRows = computed(() => {
  const rows = [] as ResponsiveTableCell[][];
  projectDataStore.entries.forEach((e: Entry) => {
    const category = projectDataStore.categories.find(c => c.Name === e.Category);
    rows.push([
      {
        content: e.Key
      },
      {
        content: e.Category
      },
      {
        content: category ? GenerateEntryForCategory(category.BibFields, e.Fields) : t(i18nKeys.ProjectPage.UnknownCategory),
        colSpan: 2
      }
    ]);
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
        content: GenerateEntryForCategory(c.BibFields, c.BibFields.map((f: Field) => f.Name)),
        colSpan: 2
      }
    ]);
  });

  return rows;
});

const categoriesHeaders = computed((): ResponsiveTableHeaderCell[] => {
  return [
    {
      content: t(i18nKeys.ProjectPage.CategoryName),
      size: SizeClasses.MaxWidth200
    },
    {
      content: t(i18nKeys.ProjectPage.CategoryExample),
      size: ""
    },
    {
      slot: true,
      size: SizeClasses.IconBtn
    },
  ];
});

const categories = computed(() => {
  return projectDataStore.categories;
});

async function syncProjectData() {
  const resp = await GetProjectData(projectName.value);
  if (resp.Ok) {
    projectDataStore.setProjectData(resp.Data.Entries, resp.Data.Categories);
  } else {
    errorSuccessStore.setMessage(false, t(i18nKeys.ProjectPage.ErrorReadingData));
  }
}

async function deleteProject() {
  deleteTriggered.value = false;
  const success = await DeleteProject(projectName.value);
  if (success) {
    errorSuccessStore.setMessage(true, t(i18nKeys.ProjectPage.SuccessDelete).replace("PROJECTNAME", projectName.value));
    projectsListStore.removeProject(projectName.value);
    appStateStore.goBack();
  } else {
    errorSuccessStore.setMessage(false, t(i18nKeys.ProjectPage.ErrorDelete));
  }
}

// watchers
watch(projectName, async () => {//TODO: reload on close and open again
  if (projectName.value !== "") {
    await syncProjectData();
  }
});

watch(uploadTriggered, () => {
  if (!uploadTriggered.value) {
    uploadedEntries.value = [];
    uploadedUnknowns.value = [];
  }
});

// methods
function openCategoryEditor(n: number) {
  appStateStore.navToPage(pageNames[2]);
  appStateStore.setItem(n > -1 ? projectDataStore.categories[n].Name : "");
}

function openEntryEditor(n: number) {
  appStateStore.navToPage(pageNames[3]);
  appStateStore.setItem(n > -1 ? projectDataStore.entries[n].Key : "");
}

function handleUpload(result: AnalyseResult) {
  uploadedEntries.value = result.Entries;
  uploadedUnknowns.value = result.Unknown;
  uploadTriggered.value = true;
}

function rmUploadedEntry(n: number) {
  uploadedEntries.value.splice(n, 1);
}

async function uploadDroppedEntries() {
  const success = await UploadEntries(projectName.value, uploadedEntries.value);
  uploadTriggered.value = false;
  if (success) {
    errorSuccessStore.setMessage(true, t(i18nKeys.ProjectPage.UploadSuccess));
    await syncProjectData();
  } else {
    errorSuccessStore.setMessage(false, t(i18nKeys.ProjectPage.UploadError));
  }
}

// onload
syncProjectData();
</script>

<style lang="scss">
@import "src/styles/common";

@include disableOverlayOnExpansionPanel;

@include fullsizeCardContainer;
</style>