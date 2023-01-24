<template>
  <ToolbarAndContent>
    <template #bar>
      <v-toolbar-title>{{ categoryName }}</v-toolbar-title>
      <v-spacer />
      <v-btn
        icon
        :disabled="!changesToSave || !rulesAreMet"
        @click="save"
      >
        <v-icon>mdi-content-save</v-icon>
      </v-btn>
      <v-btn
        icon
        :disabled="categoryName === ''"
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
                      v-model="category.Name"
                      color="primary"
                      variant="underlined"
                      :rules="[categoryNameRule]"
                    />
                  </template>
                  <template #1-1>
                    <v-combobox
                      v-model="category.CitaviCategory"
                      color="primary"
                      variant="underlined"
                      :items="categories"
                      :menu-props="{maxHeight: '200px'}"
                    />
                  </template>
                  <template #2-1>
                    <v-combobox
                      v-model="category.CitaviFilter"
                      color="primary"
                      variant="underlined"
                      :items="attributes"
                      multiple
                      hide-selected
                      :menu-props="{maxHeight: '200px'}"
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
                  <p class="example" v-html="bibExample" />
                  <ResponsiveTable
                    :rows="bibRows"
                    :headers="bibHeaders"
                    disable-ripple
                  >
                    <template #h-6>
                      <v-btn
                        flat
                        text
                        @click="addBibField"
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
                        :rules="[attributeNameRule]"
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
                        v-model="category.BibFields[i-1].CitaviMapping"
                        color="primary"
                        :items="attributes"
                        multiple
                        hide-selected
                        variant="underlined"
                        :menu-props="{maxHeight: '200px'}"
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
                        @click="rmBibField(i-1)"
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
                  <p class="example" v-html="citeExample" />
                  <ResponsiveTable
                    :rows="citeRows"
                    :headers="citeHeaders"
                    disable-ripple
                  >
                    <template #h-4>
                      <v-btn
                        flat
                        text
                        @click="addCiteField"
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
                        :rules="[attributeNameRule]"
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
                        @click="rmCiteField(i-1)"
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
  <v-dialog
    v-model="deleteTriggered"
    width="450"
  >
    <v-card>
      <v-card-title>
        {{ t(i18nKeys.CategoryEditor.DeleteTitle) }}
      </v-card-title>
      <v-card-text>
        <i18n-t :keypath="i18nKeys.CategoryEditor.DeleteMessage">
          <template #name>
            <i>{{ categoryName }}</i>
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
          @click="deleteCategory"
        >
          {{ t(i18nKeys.Common.Delete) }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
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
import {attributes, categories} from "../domain/citavi/Citavi";
import GenerateEntryForCategory from "../domain/category/GenerateEntry";
import getCategoryNameRules from "../domain/category/CategoryNameRules";
import getAttributeNameRules from "../domain/category/AttributeNameRules";
import {useErrorSuccessStore} from "@thesortex/vue-component-library/src/stores/ErrorSuccessStore/ErrorSuccessStore";
import DeleteCategory from "../api/projectData/DeleteCategory";

// globals
const appStateStore = useAppStateStore();

const projectDataStore = useProjectDataStore();

const errorSuccessStore = useErrorSuccessStore();

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

const deleteTriggered = ref(false);
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
      size: SizeClasses.MaxWidth250
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

const bibExample = computed(() => {
  return GenerateEntryForCategory(category.value!.BibFields, category.value!.BibFields.map(f => f.Name));
});

const citeExample = computed(() => {
  return GenerateEntryForCategory(category.value!.CiteFields, category.value!.CiteFields.map(f => f.Name));
});

const categoryNameRule = computed(() => {
  return getCategoryNameRules(projectDataStore.categories.map(c => c.Name), categoryName.value, t);
});

const attributeNameRule = computed(() => {
  return getAttributeNameRules(t);
});

const rulesAreMet = computed(() => {
  let valid = true;
  for (let i = 0; i < category.value!.BibFields.length; i++) {
    if (categoryNameRule.value(category.value!.BibFields[i].Name) !== true) {
      valid = false;
      break;
    }
  }
  for (let i = 0; i < category.value!.CiteFields.length; i++) {
    if (categoryNameRule.value(category.value!.CiteFields[i].Name) !== true) {
      valid = false;
      break;
    }
  }
  return valid && categoryNameRule.value(category.value!.Name) === true;
});

// watchers
watch(categoryName, () => {
  if (categoryName.value != "") {
    getCategoryFromStore();
  }
});

// functions
function getCategoryFromStore() {
  category.value = JSON.parse(JSON.stringify(projectDataStore.categories.find(c => c.Name === categoryName.value)!));
}

function getSlotName(i: number, n: number) {
  return `${i}-${n}`;
}

function addBibField() {
  category.value!.BibFields.push({
    Name: "",
    Format: {
      Prefix: "",
      Suffix: "",
      Italic: false,
      Preformatted: false
    },
    CitaviMapping: []
  });
}

function rmBibField(n: number) {
  category.value!.BibFields.splice(n, 1);
}

function addCiteField() {
  category.value!.CiteFields.push({
    Name: "",
    Format: {
      Prefix: "",
      Suffix: "",
      Italic: false,
      Preformatted: false
    },
    CitaviMapping: []
  });
}

function rmCiteField(n: number) {
  category.value!.CiteFields.splice(n, 1);
}

async function save() {
  const success = await SaveCategory(categoryName.value, appStateStore.currentProject, category.value!);
  if (success) {
    projectDataStore.actualizeCategory(categoryName.value, category.value!);
    appStateStore.setItem(category.value!.Name);
    errorSuccessStore.setMessage(true, t(i18nKeys.CategoryEditor.SuccessSave));
  } else {
    errorSuccessStore.setMessage(false, t(i18nKeys.CategoryEditor.ErrorSave));
  }
}

async function deleteCategory() {
  const success = await DeleteCategory(appStateStore.currentProject, categoryName.value);
  deleteTriggered.value = false;
  if (success) {
    projectDataStore.removeCategory(categoryName.value);
    appStateStore.goBack();
    errorSuccessStore.setMessage(true, t(i18nKeys.CategoryEditor.SuccessDelete));
  } else {
    errorSuccessStore.setMessage(false, t(i18nKeys.CategoryEditor.ErrorDelete));
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
.example {
  padding: 8px;
  border: 1px solid black;
  border-radius: 4px;
  margin-bottom: 8px;
  font-family: "Times New Roman";
  font-size: 12pt;
}
</style>