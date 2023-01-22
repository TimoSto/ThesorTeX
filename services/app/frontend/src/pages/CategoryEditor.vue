<template>
  <ToolbarAndContent>
    <template #bar>
      <v-toolbar-title>{{ categoryName }}</v-toolbar-title>
      <v-spacer />
      <v-btn
        icon
        :disabled="!changesToSave"
        @click="save"
      >
        <v-icon>mdi-content-save</v-icon>
      </v-btn>
    </template>
    <template #content>
      <div class="fullsize-card-container">
        <v-card elevation="3">
          <v-expansion-panels :model-value="0">
            <v-expansion-panel>
              <v-expansion-panel-title>
                {{ t(i18nKeys.CategoryEditor.General) }}
              </v-expansion-panel-title>
              <v-expansion-panel-text>
                <ResponsiveTable
                  :rows="generalRows"
                  :headers="generalHeaders"
                >
                  <template #0-1>
                    <v-text-field
                      v-model="category.Name"
                      color="primary"
                      variant="underlined"
                    />
                  </template>
                  <template #1-1>
                    <v-text-field
                      v-model="category.CitaviCategory"
                      color="primary"
                      variant="underlined"
                    />
                  </template>
                  <template #2-1>
                    <v-combobox
                      v-model="category.CitaviFilter"
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
      <div class="fullsize-card-container">
        <v-card elevation="3">
          <v-card elevation="3">
            <v-expansion-panels :model-value="0">
              <v-expansion-panel>
                <v-expansion-panel-title>
                  {{ t(i18nKeys.CategoryEditor.BibEntry) }}
                </v-expansion-panel-title>
                <v-expansion-panel-text>
                  <ResponsiveTable
                    :rows="bibRows"
                    :headers="bibHeaders"
                  >
                    <template #h-6>
                      <v-btn
                        flat
                        text
                      >
                        <v-icon>mdi-plus</v-icon>
                      </v-btn>
                    </template>
                    <template
                      v-for="i in bibRows.length"
                      #[getSlotName(i-1,0)]
                      :key="`bib-cell-${i}-0`"
                    >
                      <v-text-field
                        v-model="category.BibFields[i-1].Name"
                        color="primary"
                        variant="underlined"
                      />
                    </template>
                    <template
                      v-for="i in bibRows.length"
                      #[getSlotName(i-1,1)]
                      :key="`bib-cell-${i}-1`"
                    >
                      <v-text-field
                        v-model="category.BibFields[i-1].Format.Prefix"
                        color="primary"
                        variant="underlined"
                      />
                    </template>
                    <template
                      v-for="i in bibRows.length"
                      #[getSlotName(i-1,2)]
                      :key="`bib-cell-${i}-2`"
                    >
                      <v-text-field
                        v-model="category.BibFields[i-1].Format.Suffix"
                        color="primary"
                        variant="underlined"
                      />
                    </template>
                    <template
                      v-for="i in bibRows.length"
                      #[getSlotName(i-1,3)]
                      :key="`bib-cell-${i}-3`"
                    >
                      <v-checkbox-btn
                        v-model="category.BibFields[i-1].Format.Italic"
                        color="primary"
                      />
                    </template>
                    <template
                      v-for="i in bibRows.length"
                      #[getSlotName(i-1,4)]
                      :key="`bib-cell-${i}-4`"
                    >
                      <v-checkbox-btn
                        v-model="category.BibFields[i-1].Format.Preformatted"
                        color="primary"
                      />
                    </template>
                    <template
                      v-for="i in bibRows.length"
                      #[getSlotName(i-1,5)]
                      :key="`bib-cell-${i}-5`"
                    >
                      <v-combobox
                        color="primary"
                        v-model="category.BibFields[i-1].CitaviMapping"
                        :items="['test', 'ts']"
                        multiple
                        variant="underlined"
                      />
                    </template>
                    <template
                      v-for="i in bibRows.length"
                      #[getSlotName(i-1,6)]
                      :key="`bib-cell-${i}-6`"
                    >
                      <v-btn
                        flat
                        text
                      >
                        <v-icon>mdi-delete</v-icon>
                      </v-btn>
                    </template>
                  </ResponsiveTable>
                </v-expansion-panel-text>
              </v-expansion-panel>
            </v-expansion-panels>
          </v-card>
        </v-card>
      </div>
      <div class="fullsize-card-container">
        <v-card elevation="3">
          <v-card elevation="3">
            <v-expansion-panels :model-value="0">
              <v-expansion-panel>
                <v-expansion-panel-title>
                  {{ t(i18nKeys.CategoryEditor.Cites) }}
                </v-expansion-panel-title>
                <v-expansion-panel-text>
                  <ResponsiveTable
                    :rows="citeRows"
                    :headers="citeHeaders"
                  >
                    <template #h-4>
                      <v-btn
                        flat
                        text
                      >
                        <v-icon>mdi-plus</v-icon>
                      </v-btn>
                    </template>
                    <template
                      v-for="i in citeRows.length"
                      #[getSlotName(i-1,0)]
                      :key="`cite-cell-${i}-0`"
                    >
                      <v-text-field
                        v-model="category.CiteFields[i-1].Name"
                        color="primary"
                        variant="underlined"
                      />
                    </template>
                    <template
                      v-for="i in citeRows.length"
                      #[getSlotName(i-1,1)]
                      :key="`cite-cell-${i}-1`"
                    >
                      <v-text-field
                        v-model="category.CiteFields[i-1].Format.Prefix"
                        color="primary"
                        variant="underlined"
                      />
                    </template>
                    <template
                      v-for="i in citeRows.length"
                      #[getSlotName(i-1,2)]
                      :key="`cite-cell-${i}-2`"
                    >
                      <v-text-field
                        v-model="category.CiteFields[i-1].Format.Suffix"
                        color="primary"
                        variant="underlined"
                      />
                    </template>
                    <template
                      v-for="i in citeRows.length"
                      #[getSlotName(i-1,3)]
                      :key="`cite-cell-${i}-3`"
                    >
                      <v-checkbox-btn
                        v-model="category.CiteFields[i-1].Format.Italic"
                        color="primary"
                      />
                    </template>
                    <template
                      v-for="i in bibRows.length"
                      #[getSlotName(i-1,4)]
                      :key="`bib-cell-${i}-4`"
                    >
                      <v-btn
                        flat
                        text
                      >
                        <v-icon>mdi-delete</v-icon>
                      </v-btn>
                    </template>
                  </ResponsiveTable>
                </v-expansion-panel-text>
              </v-expansion-panel>
            </v-expansion-panels>
          </v-card>
        </v-card>
      </div>
    </template>
  </ToolbarAndContent>
</template>

<script lang="ts" setup>
import {useAppStateStore} from "../stores/appState/AppStateStore";
import {computed, ref, watch} from "vue";
import ResponsiveTable, {
  ResponsiveTableCell,
  ResponsiveTableHeaderCell,
  SizeClasses
} from "../components/ResponsiveTable.vue";
import {i18nKeys} from "../i18n/keys";
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {useProjectDataStore} from "../stores/projectData/ProjectDataStore";
import {Category} from "../domain/category/Category";
import SaveCategory from "../api/projectData/SaveCategory";

// globals
const appStateStore = useAppStateStore();

const projectDataStore = useProjectDataStore();

const {t} = useI18n();

// data
const category = ref(undefined as Category | undefined);

const fieldRow = [
  {
    slot: true
  },
  {
    slot: true
  },
  {
    slot: true
  },
  {
    slot: true,
    centered: true
  },
  {
    slot: true,
    centered: true
  },
  {
    slot: true
  },
  {
    slot: true,
  },
];

// computed
const categoryName = computed(() => {
  return appStateStore.currentItem;
});

const generalHeaders = computed((): ResponsiveTableHeaderCell[] => {
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
        content: t(i18nKeys.CategoryEditor.Name)
      },
      {
        slot: true
      }
    ],
    [
      {
        content: t(i18nKeys.CategoryEditor.CitaviCategory)
      },
      {
        slot: true
      }
    ],
    [
      {
        content: t(i18nKeys.CategoryEditor.CitaviFilter)
      },
      {
        slot: true
      }
    ]
  ];
});

const bibHeaders = computed((): ResponsiveTableHeaderCell[] => {
  return [
    {
      content: t(i18nKeys.Common.Attribute),
      size: ""
    },
    {
      content: t(i18nKeys.CategoryEditor.Prefix),
      size: ""
    },
    {
      content: t(i18nKeys.CategoryEditor.Suffix),
      size: ""
    },
    {
      content: t(i18nKeys.CategoryEditor.Italic),
      size: SizeClasses.MaxWidth100,
      centered: true
    },
    {
      content: t(i18nKeys.CategoryEditor.Preformatted),
      size: SizeClasses.MaxWidth100,
      centered: true
    },
    {
      content: t(i18nKeys.CategoryEditor.CitaviMapping),
      size: ""
    },
    {
      slot: true,
      size: SizeClasses.IconBtn
    }
  ];
});

const citeHeaders = computed((): ResponsiveTableHeaderCell[] => {
  let h = bibHeaders.value;
  h = h.slice(0, 4);
  h.push(bibHeaders.value[6]);
  return h;
});

const bibRows = computed((): ResponsiveTableCell[][] => {
  return Array(category.value ? category.value.BibFields.length : 0).fill(fieldRow);
});

const citeRows = computed((): ResponsiveTableCell[][] => {
  return Array(category.value ? category.value.CiteFields.length : 0).fill(fieldRow.slice(0, 5));
});

const changesToSave = computed(() => {
  return JSON.stringify(category.value) !== JSON.stringify(projectDataStore.categories.find(c => c.Name === categoryName.value));
});

// watchers
watch(categoryName, () => {
  if (categoryName.value != "") {
    getCategoryFromStore();
  }
});

// functions
function getCategoryFromStore() {
  console.log(categoryName.value);
  category.value = JSON.parse(JSON.stringify(projectDataStore.categories.find(c => c.Name === categoryName.value)!));
}

function getSlotName(i: number, n: number) {
  return `${i}-${n}`;
}

async function save() {
  const success = await SaveCategory(categoryName.value, appStateStore.currentProject, category.value!);
  if (success) {
    projectDataStore.actualizeCategory(categoryName.value, category.value!);
    appStateStore.setItem(category.value!.Name);
  }
}

// onload
if (categoryName.value != "") {
  getCategoryFromStore();
} else {
  category.value = {
    Name: "",
    CitaviCategory: "",
    BibFields: [],
    CiteFields: [],
    CitaviFilter: [],
  };
}
</script>

<style scoped>

</style>