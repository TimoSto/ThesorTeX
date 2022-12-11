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
              Allgemein
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
                    :items="citaviCategories"
                    v-model="citaviCategory"
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
                  >
                  </v-combobox>
                </template>
              </ResponsiveTable>
            </v-expansion-panel-text>
          </v-expansion-panel>
          <v-expansion-panel>
            <v-expansion-panel-title>
              Literaturverzeichnis
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
                    v-model="bibValues[i][0]"
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
                    v-model="bibValues[i][1]"
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
                    v-model="bibValues[i][2]"
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
                    v-model="bibValues[i][3]"
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
                    v-model="bibValues[i][4]"
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
                    v-model="bibValues[i][5]"
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
              Zitate
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
                    v-model="citeValues[i][0]"
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
                    v-model="citeValues[i][1]"
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
                    v-model="citeValues[i][2]"
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
                    v-model="citeValues[i][3]"
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
                    v-model="citeValues[i][4]"
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
                    v-model="citeValues[i][5]"
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
import {ref} from "vue";

//globals
const emit = defineEmits(['navBack']);

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
    content: 'Attribut',
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: '60%',
    minWidth: '',
    content: 'Wert',
    icon: '',
    hideUnder: -1,
    event: ''
  }
];

const bibHeaders: ResponsiveTableHeaderCell[] = [
  {
    width: '16%',
    minWidth: '',
    content: 'Feld',
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: '16%',
    minWidth: '',
    content: 'Stil',
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: '16%',
    minWidth: '',
    content: 'Prefix',
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: '16%',
    minWidth: '',
    content: 'Suffix',
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: '16%',
    minWidth: '',
    content: 'Formatiert',
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: '16%',
    minWidth: '',
    content: 'Citavi Attribute',
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

const bibFields = ref([] as ResponsiveTableCell[][]);

const bibValues = ref([] as any[][])

const citeFields = ref([] as ResponsiveTableCell[][]);

const citeValues = ref([] as any[][])

// methods
function AddBibRow() {
  bibValues.value.push([
    '',
    '',
    '',
    '',
    false,
    []
  ])
  bibFields.value.push(Array(7).fill({
    content: '',
    icon: '',
    hideUnder: -1,
    event: '',
    slot: true
  }));
}

function AddCiteRow() {
  citeValues.value.push([
    '',
    '',
    '',
    '',
    false,
    []
  ])
  citeFields.value.push(Array(7).fill({
    content: '',
    icon: '',
    hideUnder: -1,
    event: '',
    slot: true
  }));
}

function getSlotName(i: number, n: number) {
  return `${i}-${n}`
}

</script>

<style scoped>

</style>
