<template>
  <ToolbarAndContent>
    <template #bar>
      <v-toolbar-title>sdf</v-toolbar-title>
      <v-spacer />
      <v-btn
        icon
        :disabled="!changesToSave"
        @click="save"
      >
        <v-icon>mdi-content-save</v-icon>
      </v-btn>
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
          <v-expansion-panels :model-value="0">
            <v-expansion-panel>
              <v-expansion-panel-title>
                {{ t(i18nKeys.Common.General) }}
              </v-expansion-panel-title>
              <v-expansion-panel-text>
                <ResponsiveTable
                  :rows="generalRows"
                  :headers="generalHeaders"
                  disable-ripple
                >
                  <template #0-1>
                    <v-text-field
                      v-model="entry.Key"
                      color="primary"
                      variant="underlined"
                      :rules="[keyRules]"
                    />
                  </template>
                  <template #1-1>
                    <v-select
                      v-model="entry.Category"
                      color="primary"
                      variant="underlined"
                      :items="categories"
                    />
                  </template>
                </ResponsiveTable>
              </v-expansion-panel-text>
            </v-expansion-panel>
          </v-expansion-panels>
        </v-card>
      </div>
      <div class="fullsize-card-container">
        <v-card elevation="3">
          <v-expansion-panels :model-value="0">
            <v-expansion-panel>
              <v-expansion-panel-title>
                {{ t(i18nKeys.EntryEditor.Fields) }}
              </v-expansion-panel-title>
              <v-expansion-panel-text>
                <ResponsiveTable
                  :rows="fieldsRows"
                  :headers="fieldsHeaders"
                  disable-ripple
                >
                  <template
                    v-for="i in fieldsRows.length"
                    :key="`field-${i}`"
                    #[getSlotName(i-1,1)]
                  >
                    <v-text-field
                      v-model="entry.Fields[i-1]"
                      color="primary"
                      variant="underlined"
                    />
                  </template>
                </ResponsiveTable>
              </v-expansion-panel-text>
            </v-expansion-panel>
          </v-expansion-panels>
        </v-card>
      </div>
    </template>
  </ToolbarAndContent>
  <v-dialog
    v-model="deleteTriggered"
    width="450"
  >
    <v-card>
      <v-card-title>
        {{ t(i18nKeys.EntryEditor.DeleteTitle) }}
      </v-card-title>
      <v-card-text>
        <i18n-t :keypath="i18nKeys.EntryEditor.DeleteMessage">
          <template #key>
            <i>{{ entryKey }}</i>
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
          @click="deleteEntry"
        >
          {{ t(i18nKeys.Common.Delete) }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts" setup>
import ResponsiveTable, {ResponsiveTableCell, SizeClasses} from "../components/ResponsiveTable.vue";
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";
import {computed, ref} from "vue";
import {useProjectDataStore} from "../stores/projectData/ProjectDataStore";
import JoinFields from "../domain/entry/JoinFields";
import {Entry} from "../domain/entry/Entry";
import {useAppStateStore} from "../stores/appState/AppStateStore";
import getEntryKeyRules from "../domain/entry/EntryKeyRules";
import SaveEntry from "../api/projectData/SaveEntry";
import {useErrorSuccessStore} from "@thesortex/vue-component-library/src/stores/ErrorSuccessStore/ErrorSuccessStore";
import DeleteEntry from "../api/projectData/DeleteEntry";

// globals
const {t} = useI18n();

const projectDataStore = useProjectDataStore();

const appStateStore = useAppStateStore();

const errorSuccessStore = useErrorSuccessStore();

// data
const entry = ref(undefined as Entry | undefined);

const deleteTriggered = ref(false);

// computed
const entryKey = computed(() => {
  return appStateStore.currentItem;
});

const generalHeaders = computed(() => {
  return [
    {
      content: t(i18nKeys.Common.Attribute),
      size: SizeClasses.MaxWidth250
    },
    {
      content: t(i18nKeys.Common.Value),
      size: ""
    }
  ];
});

const generalRows = computed((): ResponsiveTableCell[][] => {
  return [
    [
      {
        content: t(i18nKeys.EntryEditor.Key),
      },
      {
        slot: true
      }
    ],
    [
      {
        content: t(i18nKeys.EntryEditor.Category),
      },
      {
        slot: true
      }
    ]
  ];
});

const categories = computed(() => {
  return projectDataStore.categories.map(c => c.Name);
});

const fields = computed(() => {
  const i = categories.value.indexOf(entry.value!.Category);
  if (i >= 0) {
    return JoinFields(projectDataStore.categories[i].BibFields, projectDataStore.categories[i].CiteFields);
  }
  return [];
});

const fieldsHeaders = computed(() => {
  return [
    {
      content: t(i18nKeys.EntryEditor.Field),
      size: SizeClasses.MaxWidth250
    },
    {
      content: t(i18nKeys.Common.Value),
      size: ""
    }
  ];
});

const fieldsRows = computed(() => {
  return fields.value.map(f => [
    {
      content: f.Name
    },
    {
      slot: true
    }
  ]);
});

const changesToSave = computed(() => {
  return JSON.stringify(entry.value) !== JSON.stringify(projectDataStore.entries.find(c => c.Key === entryKey.value));
});

const keyRules = computed(() => {
  return getEntryKeyRules(projectDataStore.entries.map(e => e.Key), entryKey.value, t);
});

// methods
function getSlotName(i: number, n: number) {
  return `${i}-${n}`;
}

function getEntryFromStore() {
  entry.value = JSON.parse(JSON.stringify(projectDataStore.entries.find(c => c.Key === entryKey.value)!));
}

async function save() {
  const success = await SaveEntry(appStateStore.currentProject, entryKey.value, entry.value!);
  if (success) {
    projectDataStore.actualizeEntry(entryKey.value, entry.value!);
    appStateStore.setItem(entry.value!.Key);
    errorSuccessStore.setMessage(true, t(i18nKeys.EntryEditor.SuccessSave));
  } else {
    errorSuccessStore.setMessage(false, t(i18nKeys.EntryEditor.ErrorSave));
  }
}

async function deleteEntry() {
  const success = await DeleteEntry(appStateStore.currentProject, entryKey.value);
  deleteTriggered.value = false;
  if (success) {
    projectDataStore.removeEntry(entryKey.value);
    appStateStore.goBack();
    errorSuccessStore.setMessage(true, t(i18nKeys.EntryEditor.SuccessDelete));
  } else {
    errorSuccessStore.setMessage(false, t(i18nKeys.EntryEditor.ErrorDelete));
  }
}

// onload
getEntryFromStore();
</script>

<style scoped>

</style>