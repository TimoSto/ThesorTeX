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
        :disabled="!categoryChanged"
      >
        <v-icon>mdi-content-save</v-icon>
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
            <v-expansion-panel-title>
              {{ t(i18nKeys.CategoryEditor.General.Title) }}
            </v-expansion-panel-title>
            <v-expansion-panel-text>
              <ResponsiveTable
                :headers="generalHeaders"
                :rows="generalRows"
              >
                <template #0-1>
                  <v-text-field
                    v-model="categoryName"
                    color="primary"
                    variant="underlined"
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
            <v-expansion-panel-title>
              {{ t(i18nKeys.CategoryEditor.Bib.Title) }}
            </v-expansion-panel-title>
            <v-expansion-panel-text>
              <ResponsiveTable
                :headers="bibHeaders"
                :rows="bibFields"
                @btn-clicked="AddBibRow"
              >
                <template
                  v-for="(r, i) in bibFields"
                  :key="`bib-cell-${i}-0`"
                  #[getSlotName(i,0)]
                >
                  <v-text-field
                    v-model="bibValues[i].Field"
                    color="primary"
                    variant="underlined"
                  />
                </template>
                <template
                  v-for="(r, i) in bibFields"
                  :key="`bib-cell-${i}-1`"
                  #[getSlotName(i,1)]
                >
                  <v-text-field
                    v-model="bibValues[i].Style"
                    color="primary"
                    variant="underlined"
                  />
                </template>
                <template
                  v-for="(r, i) in bibFields"
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
                  v-for="(r, i) in bibFields"
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
                  v-for="(r, i) in bibFields"
                  :key="`bib-cell-${i}-4`"
                  #[getSlotName(i,4)]
                >
                  <v-checkbox-btn
                    v-model="bibValues[i].TexValue"
                    color="primary"
                    density="compact"
                  />
                </template>
                <template
                  v-for="(r, i) in bibFields"
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
                  v-for="(r, i) in bibFields"
                  :key="`bib-cell-${i}-6`"
                  #[getSlotName(i,6)]
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
          <v-expansion-panel>
            <v-expansion-panel-title>
              {{ t(i18nKeys.CategoryEditor.Cite.Title) }}
            </v-expansion-panel-title>
            <v-expansion-panel-text>
              <ResponsiveTable
                :headers="bibHeaders"
                :rows="citeFields"
                @btn-clicked="AddCiteRow"
              >
                <template
                  v-for="(r, i) in citeFields"
                  :key="`cite-cell-${i}-0`"
                  #[getSlotName(i,0)]
                >
                  <v-text-field
                    v-model="citeValues[i].Field"
                    color="primary"
                    variant="underlined"
                  />
                </template>
                <template
                  v-for="(r, i) in citeFields"
                  :key="`cite-cell-${i}-1`"
                  #[getSlotName(i,1)]
                >
                  <v-text-field
                    v-model="citeValues[i].Style"
                    color="primary"
                    variant="underlined"
                  />
                </template>
                <template
                  v-for="(r, i) in citeFields"
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
                  v-for="(r, i) in citeFields"
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
                  v-for="(r, i) in citeFields"
                  :key="`cite-cell-${i}-4`"
                  #[getSlotName(i,4)]
                >
                  <v-checkbox-btn
                    v-model="citeValues[i].TexValue"
                    color="primary"
                    density="compact"
                  />
                </template>
                <template
                  v-for="(r, i) in citeFields"
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
                  v-for="(r, i) in citeFields"
                  :key="`cite-cell-${i}-6`"
                  #[getSlotName(i,6)]
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
      </div>
    </template>
  </AppbarContent>
</template>

<script setup lang="ts">
import AppbarContent from "../../../components/AppbarContent.vue";
import ResponsiveTable, {ResponsiveTableCell, ResponsiveTableHeaderCell} from "../../../components/ResponsiveTable.vue";
import {computed, ref} from "vue";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";
import {useProjectDataStore} from "../stores/projectDataStore";
import CheckCategoryChanged from "../api/projectData/categories/CheckCategoryChanged";
import type {Field} from "../api/projectData/categories/BibCategory";

//globals
const emit = defineEmits(['navBack']);

const { t } = useI18n();

const projectDataStore = useProjectDataStore();

//props
const props = defineProps({
  categoryName: {
    type: String,
    required: true
  }
});

// data
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
    width: '16%',
    minWidth: '',
    content: t(i18nKeys.CategoryEditor.Bib.Field),
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: '16%',
    minWidth: '',
    content: t(i18nKeys.CategoryEditor.Bib.Style),
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: '16%',
    minWidth: '',
    content: t(i18nKeys.CategoryEditor.Bib.Prefix),
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: '16%',
    minWidth: '',
    content: t(i18nKeys.CategoryEditor.Bib.Suffix),
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: '16%',
    minWidth: '',
    content: t(i18nKeys.CategoryEditor.Bib.Formatted),
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: '16%',
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
    event: ''
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

// methods
function AddBibRow() {
  bibValues.value.push({
    Field: '',
    Style: '',
    Prefix: '',
    Suffix: '',
    TexValue: false,
    CitaviAttributes: []
  });
}

function AddCiteRow() {
  citeValues.value.push({
    Field: '',
    Style: '',
    Prefix: '',
    Suffix: '',
    TexValue: false,
    CitaviAttributes: []
  });
}

function getSlotName(i: number, n: number) {
  return `${i}-${n}`
}

// onload
if( initialCategory.value ) {
  categoryName.value = initialCategory.value.Name;
  citaviCategory.value = initialCategory.value.CitaviType;//todo: rename to CitaviCategory
  citaviFilter.value = initialCategory.value.CitaviNecessaryFields;
  bibValues.value = initialCategory.value.Fields.map(f => {
    return {
      Field: f.Field,
      Style: f.Style,
      Prefix: f.Prefix,
      Suffix: f.Suffix,
      TexValue: f.TexValue,
      CitaviAttributes: f.CitaviAttributes
    }
  })
  citeValues.value = initialCategory.value.CiteFields.map(f => {
    return {
      Field: f.Field,
      Style: f.Style,
      Prefix: f.Prefix,
      Suffix: f.Suffix,
      TexValue: f.TexValue,
      CitaviAttributes: f.CitaviAttributes
    }
  })
}

</script>

<style scoped>

</style>
