<template>
  <AppbarContent
    :level="2"
    bar-color="background"
  >
    <template #bar>
      <v-app-bar-nav-icon
        @click="emit('navBack')"
      >
        <v-icon>mdi-arrow-left</v-icon>
      </v-app-bar-nav-icon>

      <v-toolbar-title>
        {{ props.categoryName }}
      </v-toolbar-title>

      <v-spacer />

      <v-btn
        icon
        :disabled="!savePossible"
        @click="CallSaveCategory"
      >
        <v-icon>mdi-content-save</v-icon>
      </v-btn>
      <v-btn
        icon
        @click="deleteOpened = true"
      >
        <v-icon>mdi-delete</v-icon>
      </v-btn>
    </template>

    <template #content>
      <div style="max-height: calc(100vh - 112px);">
        <v-expansion-panels
          multiple
          variant="accordion"
          :model-value="[0, 1, 2]"
        >
          <v-expansion-panel>
            <v-expansion-panel-title v-ripple>
              {{ t(i18nKeys.CategoryEditor.General.Title) }}
            </v-expansion-panel-title>
            <v-expansion-panel-text>
              <ResponsiveTable
                :headers="generalHeaders"
                :rows="generalRows"
                :disable-ripple="true"
              >
                <template #0-1>
                  <v-text-field
                    v-model="categoryName"
                    color="primary"
                    variant="underlined"
                    :rules="categoryNameRules"
                  />
                </template>
                <template #1-1>
                  <v-combobox
                    v-model="citaviCategory"
                    :items="citaviCategories"
                    color="primary"
                    variant="underlined"
                  />
                </template>
                <template #2-1>
                  <v-combobox
                    v-model="citaviFilter"
                    :items="citaviAttributes"
                    color="primary"
                    variant="underlined"
                    multiple
                    hide-selected
                    :disabled="citaviCategory === ''"
                  >
                  </v-combobox>
                </template>
              </ResponsiveTable>
            </v-expansion-panel-text>
          </v-expansion-panel>
          <v-expansion-panel>
            <v-expansion-panel-title v-ripple>
              {{ t(i18nKeys.CategoryEditor.Bib.Title) }}
            </v-expansion-panel-title>
            <v-expansion-panel-text>
              <p
                class="example"
                v-html="bibModel"
              />
              <ResponsiveTable
                :headers="bibHeaders"
                :rows="bibFields"
                :disable-ripple="true"
                @btn-clicked="AddBibRow"
              >
                <template
                  v-for="(_, i) in bibFields"
                  :key="`bib-cell-${i}-0`"
                  #[getSlotName(i,0)]
                >
                  <v-text-field
                    v-model="bibValues[i].Field"
                    color="primary"
                    variant="underlined"
                    :rules="fieldNameRules(bibValues.map(v => v.Field))"
                  />
                </template>
                <template
                  v-for="(_, i) in bibFields"
                  :key="`bib-cell-${i}-1`"
                  #[getSlotName(i,1)]
                >
                  <span class="centeredCell">
                    <v-checkbox-btn
                      v-model="bibValues[i].Italic"
                      color="primary"
                      density="compact"
                    />
                  </span>
                </template>
                <template
                  v-for="(_, i) in bibFields"
                  :key="`bib-cell-${i}-2`"
                  #[getSlotName(i,2)]
                >
                  <v-text-field
                    v-model="bibValues[i].Prefix"
                    color="primary"
                    variant="underlined"
                  />
                </template>
                <template
                  v-for="(_, i) in bibFields"
                  :key="`bib-cell-${i}-3`"
                  #[getSlotName(i,3)]
                >
                  <v-text-field
                    v-model="bibValues[i].Suffix"
                    color="primary"
                    variant="underlined"
                  />
                </template>
                <template
                  v-for="(_, i) in bibFields"
                  :key="`bib-cell-${i}-4`"
                  #[getSlotName(i,4)]
                >
                  <span class="centeredCell">
                    <v-checkbox-btn
                      v-model="bibValues[i].TexValue"
                      color="primary"
                      density="compact"
                    />
                  </span>
                </template>
                <template
                  v-for="(_, i) in bibFields"
                  :key="`bib-cell-${i}-5`"
                  #[getSlotName(i,5)]
                >
                  <v-combobox
                    v-model="bibValues[i].CitaviAttributes"
                    :items="citaviAttributes"
                    color="primary"
                    variant="underlined"
                  />
                </template>
                <template
                  v-for="(_, i) in bibFields"
                  :key="`bib-cell-${i}-6`"
                  #[getSlotName(i,6)]
                >
                  <v-btn
                    flat
                    text
                    @click="rmBibField(i)"
                  >
                    <v-icon>mdi-delete</v-icon>
                  </v-btn>
                </template>
              </ResponsiveTable>
            </v-expansion-panel-text>
          </v-expansion-panel>
          <v-expansion-panel>
            <v-expansion-panel-title v-ripple>
              {{ t(i18nKeys.CategoryEditor.Cite.Title) }}
            </v-expansion-panel-title>
            <v-expansion-panel-text>
              <p
                class="example"
                v-html="citeModel"
              />
              <ResponsiveTable
                :headers="bibHeaders"
                :rows="citeFields"
                :disable-ripple="true"
                @btn-clicked="AddCiteRow"
              >
                <template
                  v-for="(_, i) in citeFields"
                  :key="`cite-cell-${i}-0`"
                  #[getSlotName(i,0)]
                >
                  <v-text-field
                    v-model="citeValues[i].Field"
                    color="primary"
                    variant="underlined"
                    :rules="fieldNameRules(citeValues.map(f => f.Field))"
                  />
                </template>
                <template
                  v-for="(_, i) in citeFields"
                  :key="`cite-cell-${i}-1`"
                  #[getSlotName(i,1)]
                >
                  <span class="centeredCell">
                    <v-checkbox-btn
                      v-model="citeValues[i].Italic"
                      color="primary"
                      density="compact"
                    />
                  </span>
                </template>
                <template
                  v-for="(_, i) in citeFields"
                  :key="`cite-cell-${i}-2`"
                  #[getSlotName(i,2)]
                >
                  <v-text-field
                    v-model="citeValues[i].Prefix"
                    color="primary"
                    variant="underlined"
                  />
                </template>
                <template
                  v-for="(_, i) in citeFields"
                  :key="`cite-cell-${i}-3`"
                  #[getSlotName(i,3)]
                >
                  <v-text-field
                    v-model="citeValues[i].Suffix"
                    color="primary"
                    variant="underlined"
                  />
                </template>
                <template
                  v-for="(_, i) in citeFields"
                  :key="`cite-cell-${i}-4`"
                  #[getSlotName(i,4)]
                >
                  <span class="centeredCell">
                    <v-checkbox-btn
                      v-model="citeValues[i].TexValue"
                      color="primary"
                      density="compact"
                    />
                  </span>
                </template>
                <template
                  v-for="(_, i) in citeFields"
                  :key="`cite-cell-${i}-5`"
                  #[getSlotName(i,5)]
                >
                  <v-combobox
                    v-model="citeValues[i].CitaviAttributes"
                    :items="citaviAttributes"
                    color="primary"
                    variant="underlined"
                  />
                </template>
                <template
                  v-for="(_, i) in citeFields"
                  :key="`cite-cell-${i}-6`"
                  #[getSlotName(i,6)]
                >
                  <v-btn
                    flat
                    text
                    @click="rmCiteField(i)"
                  >
                    <v-icon>mdi-delete</v-icon>
                  </v-btn>
                </template>
              </ResponsiveTable>
            </v-expansion-panel-text>
          </v-expansion-panel>
        </v-expansion-panels>
      </div>
    </template>
  </AppbarContent>
  <DeleteCategoryDialog
    :project="props.projectName"
    :category="props.categoryName"
    :open="deleteOpened"
    @close="deleteOpened = false"
    @success="NavBackAndRefresh"
  />
</template>

<script setup lang="ts">
import AppbarContent from "../components/AppbarContent.vue";
import ResponsiveTable, {ResponsiveTableCell, ResponsiveTableHeaderCell} from "../components/ResponsiveTable.vue";
import {computed, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";
import {useProjectDataStore} from "../stores/projectDataStore";
import CheckCategoryChanged from "../api/projectData/categories/CheckCategoryChanged";
import type {Field} from "../api/projectData/categories/BibCategory";
import GenerateModelForFields from "../api/projectData/categories/GenerateModel";
import SaveCategory from "../api/projectData/categories/SaveCategory";
import {useErrorSuccessStore} from "../stores/errorSuccessStore";
import GetProjectData from "../api/projectData/GetProjectData";
import GetCategoryNameRules from "../rules/categoryNameRules";
import GetFieldNameRules from "../rules/fieldNameRules";
import DeleteCategoryDialog from "./DeleteCategoryDialog.vue";
import {useProjectsStore} from "../stores/projectsStore";
import {useUnsaveCloseStore} from "../stores/unsaveCloseStore";

//globals
const emit = defineEmits(['navBack', 'titleChange']);

const { t } = useI18n();

const projectDataStore = useProjectDataStore();

const errorSuccessStore = useErrorSuccessStore();

const projectsStore = useProjectsStore();

const unsafeCloseStore = useUnsaveCloseStore();

//props
const props = defineProps({
  categoryName: {
    type: String,
    required: true
  },
  projectName: {
    type: String,
    required: true
  }
});

// data
const deleteOpened = ref(false);

const citaviCategories = [
  'article',
  'thesis'
];

const citaviAttributes = [
  'author',
  'doi',
];

const generalHeaders: ResponsiveTableHeaderCell[] = [
  {
    width: '40%',
    minWidth: '',
    content: t(i18nKeys.Common.Attribute),
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: '60%',
    minWidth: '',
    content: t(i18nKeys.Common.Value),
    icon: '',
    hideUnder: -1,
    event: ''
  }
];

const bibHeaders: ResponsiveTableHeaderCell[] = [
  {
    width: 'calc(25% - 56px)',
    minWidth: '',
    content: t(i18nKeys.CategoryEditor.Bib.Field),
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: '70px',
    minWidth: '',
    content: t(i18nKeys.CategoryEditor.Bib.Italic),
    icon: '',
    hideUnder: -1,
    event: '',
    maxWidth: true
  },
  {
    width: 'calc(25% - 56px)',
    minWidth: '',
    content: t(i18nKeys.CategoryEditor.Bib.Prefix),
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: 'calc(25% - 56px)',
    minWidth: '',
    content: t(i18nKeys.CategoryEditor.Bib.Suffix),
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: '100px',
    minWidth: '',
    content: t(i18nKeys.CategoryEditor.Bib.Formatted),
    icon: '',
    hideUnder: -1,
    event: '',
    maxWidth: true,
  },
  {
    width: 'calc(25% - 56px)',
    minWidth: '',
    content: t(i18nKeys.CategoryEditor.Bib.CitaviAttributes),
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: '48px',
    minWidth: '',
    content: '',
    icon: 'mdi-plus',
    hideUnder: -1,
    event: '',
    maxWidth: true
  },
];

const generalRows: ResponsiveTableCell[][] = [
  [
    {
      content: 'Name',
      icon: '',
      event: '',
      hideUnder: -1
    },
    {
      content: '',
      icon: '',
      event: '',
      hideUnder: -1,
      slot: true
    },
  ],
  [
    {
      content: 'Citavi Kategorie',
      icon: '',
      event: '',
      hideUnder: -1
    },
    {
      content: '',
      icon: '',
      event: '',
      hideUnder: -1,
      slot: true
    },
  ],
  [
    {
      content: 'Citavi Filter',
      icon: '',
      event: '',
      hideUnder: -1
    },
    {
      content: '',
      icon: '',
      event: '',
      hideUnder: -1,
      slot: true
    },
  ],
];

const categoryName = ref('');

const citaviCategory = ref('');

const citaviFilter = ref([] as string[]);

const bibValues = ref([] as Field[])

const citeValues = ref([] as Field[])

// computed

const bibFields = computed(() => {
  const rows: ResponsiveTableCell[][] = [];
  bibValues.value.forEach(() => {
    rows.push(Array(7).fill({
      content: '',
      icon: '',
      hideUnder: -1,
      event: '',
      slot: true
    }))
  });

  return rows;
})

const citeFields = computed(() => {
  const rows: ResponsiveTableCell[][] = [];
  citeValues.value.forEach(() => {
    rows.push(Array(7).fill({
      content: '',
      icon: '',
      hideUnder: -1,
      event: '',
      slot: true
    }))
  });

  return rows;
})

const initialCategory = computed(() => {
  return projectDataStore.categories.find(c => c.Name === props.categoryName)
})

const categoryChanged = computed(() => {
  if( initialCategory.value ) {
    return CheckCategoryChanged(initialCategory.value, categoryName.value, citaviCategory.value, citaviFilter.value, bibValues.value, citeValues.value)
  }

  return false

})

const bibModel = computed(() => {
  return GenerateModelForFields(bibValues.value)
});

const citeModel = computed(() => {
  return GenerateModelForFields(citeValues.value)
});

const categoryNameRules = computed(() => {
  return GetCategoryNameRules(projectDataStore.categories.map(c => c.Name), initialCategory.value ? initialCategory.value.Name : '', t)
})

const savePossible = computed(() => {
  if( !categoryChanged.value ) {
    return false
  }
  if( !categoryNameRules.value[0](categoryName.value) ) {
    return false
  }
  for( let i = 0 ; i < bibValues.value.length ; i++ ) {
    if( typeof fieldNameRules(bibValues.value.map(f => f.Field))[0](bibValues.value[i].Field) === "string" ) {
      return false
    }
  }
  for( let i = 0 ; i < citeValues.value.length ; i++ ) {
    if( typeof fieldNameRules(citeValues.value.map(f => f.Field))[0](citeValues.value[i].Field) === "string" ) {
      return false
    }
  }
  return true
})

const closeTried = computed(() => {
  return unsafeCloseStore.tried;
});

// watchers
watch(closeTried, (nV) => {
  if( nV ) {
    if( unsafeCloseStore.fromPage === 3 ) {
      if( !savePossible.value ) {
        unsafeCloseStore.promiseResolve(true);
      } else {
        unsafeCloseStore.triggered = true;
      }
    }
  }
})

// methods
function AddBibRow() {
  bibValues.value.push({
    Field: '',
    Italic: false,
    Prefix: '',
    Suffix: '',
    TexValue: false,
    CitaviAttributes: []
  });
}

function AddCiteRow() {
  citeValues.value.push({
    Field: '',
    Italic: false,
    Prefix: '',
    Suffix: '',
    TexValue: false,
    CitaviAttributes: []
  });
}

function CallSaveCategory() {
  SaveCategory(props.projectName, categoryName.value, initialCategory.value ? initialCategory.value.Name : '', citaviCategory.value, citaviFilter.value, bibValues.value, citeValues.value)
    .then(data => {
      errorSuccessStore.handleResponse(data.Ok, t(i18nKeys.Success.SaveCategory), t(i18nKeys.Errors.ErrorSaving));
      if( data.Ok ) {
        GetProjectData(props.projectName).then(data => {
          projectDataStore.setData(data);
          emit('titleChange', categoryName.value)
        });
        projectsStore.updateLastEditedOnProject(props.projectName, data.Data.LastModified);
      }
    })
}

function getSlotName(i: number, n: number) {
  return `${i}-${n}`
}

function fieldNameRules (existing: string[]) {
  return GetFieldNameRules(existing, t);
}

function NavBackAndRefresh() {
  deleteOpened.value = false;
  GetProjectData(props.projectName).then(data => {
    projectDataStore.setData(data);
    emit('navBack');
  })
}

function rmBibField(n: number) {
  bibValues.value.splice(n, 1);
}

function rmCiteField(n: number) {
  citeValues.value.splice(n, 1);
}

// onload
if( initialCategory.value ) {
  categoryName.value = initialCategory.value.Name;
  citaviCategory.value = initialCategory.value.CitaviCategory;//todo: rename to CitaviCategory
  citaviFilter.value = initialCategory.value.CitaviFilters;
  bibValues.value = initialCategory.value.Fields.map(f => {
    return {
      Field: f.Field,
      Italic: f.Italic,
      Prefix: f.Prefix,
      Suffix: f.Suffix,
      TexValue: f.TexValue,
      CitaviAttributes: f.CitaviAttributes
    }
  })
  citeValues.value = initialCategory.value.CiteFields.map(f => {
    return {
      Field: f.Field,
      Italic: f.Italic,
      Prefix: f.Prefix,
      Suffix: f.Suffix,
      TexValue: f.TexValue,
      CitaviAttributes: f.CitaviAttributes
    }
  })
}

</script>

<style scoped lang="scss">

.centeredCell {
  display: inline-block;
  text-align: center;
  width: 100%;
}

.example {
  padding: 8px;
}
</style>
