<template>
  <ToolbarAndContent>
    <template #bar>
      <v-toolbar-title>sdf</v-toolbar-title>
      <v-spacer />
      <v-btn
        icon
      >
        <v-icon>mdi-content-save</v-icon>
      </v-btn>
      <v-btn
        icon
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
                      color="primary"
                      variant="underlined"
                    />
                  </template>
                  <template #1-1>
                    <v-select
                      v-model="category"
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
                      v-model="fieldsValues[i]"
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
</template>

<script lang="ts" setup>
import ResponsiveTable, {ResponsiveTableCell, SizeClasses} from "../components/ResponsiveTable.vue";
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";
import {computed, ref} from "vue";
import {useProjectDataStore} from "../stores/projectData/ProjectDataStore";
import JoinFields from "../domain/entry/JoinFields";

// globals
const {t} = useI18n();

const projectDataStore = useProjectDataStore();

// data
const category = ref("");

const fieldsValues = ref([] as string[]);

// computed
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
  const i = categories.value.indexOf(category.value);
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

// methods
function getSlotName(i: number, n: number) {
  return `${i}-${n}`;
}
</script>

<style scoped>

</style>