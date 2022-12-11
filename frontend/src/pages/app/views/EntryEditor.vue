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
        {{ props.entryKey }}
      </v-toolbar-title>

      <v-spacer />

      <v-btn
        icon
        :disabled="!entryChanged"
        @click="CallSaveEntry"
      >
        <v-icon>mdi-content-save</v-icon>
      </v-btn>
      <v-btn icon>
        <v-icon>mdi-delete</v-icon>
      </v-btn>
    </template>

    <template #content>
      <div style="max-height: calc(100vh - 112px);">
        <v-expansion-panels
          multiple
          variant="accordion"
          :model-value="[0,1]"
        >
          <v-expansion-panel>
            <v-expansion-panel-title>
              {{ t(i18nKeys.EntryEditor.General.Title) }}
            </v-expansion-panel-title>
            <v-expansion-panel-text>
              <ResponsiveTable
                :headers="generalHeaders"
                :rows="generalRows"
                :disable-ripple="true"
              >
                <template #0-1>
                  <v-text-field
                    v-model="entryKey"
                    color="primary"
                    variant="underlined"
                  />
                </template>
                <template #1-1>
                  <v-select
                    v-model="entryCategory"
                    color="primary"
                    variant="underlined"
                    :items="availableCategories"
                  />
                </template>
              </ResponsiveTable>
            </v-expansion-panel-text>
          </v-expansion-panel>
          <v-expansion-panel>
            <v-expansion-panel-title>
              {{ t(i18nKeys.EntryEditor.Fields.Title) }}
            </v-expansion-panel-title>
            <v-expansion-panel-text>
              <ResponsiveTable
                :headers="fieldsHeaders"
                :rows="fieldRows"
                :disable-ripple="true"
              >
                <template
                  v-for="(_,i) in fieldRows"
                  :key="`field-${i}`"
                  #[getSlotName(i)]
                >
                  <v-text-field
                    v-model="entryFields[i]"
                    color="primary"
                    variant="underlined"
                  />
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
import {i18nKeys} from "../i18n/keys";
import {useI18n} from "vue-i18n";
import ResponsiveTable, {ResponsiveTableCell, ResponsiveTableHeaderCell} from "../../../components/ResponsiveTable.vue";
import {computed, ref, watch} from "vue";
import {useProjectDataStore} from "../stores/projectDataStore";
import SaveEntry from "../api/projectData/entries/SaveEntry";
import {BibEntry} from "../api/projectData/entries/BibEntry";

// globals
const emit = defineEmits(['navBack']);

const { t } = useI18n();

const projectDataStore = useProjectDataStore();

// props
const props = defineProps({
  entryKey: {
    type: String,
    required: true
  },
  projectName: {
    type: String,
    required: true
  }
});

// data
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
]

const generalRows: ResponsiveTableCell[][] = [
  [
    {
      content: t(i18nKeys.EntryEditor.General.Key),
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
      content: t(i18nKeys.EntryEditor.General.Category),
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

const fieldsHeaders: ResponsiveTableHeaderCell[] = [
  {
    width: '40%',
    minWidth: '',
    content: t( i18nKeys.Common.Attribute ),
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: '60%',
    minWidth: '',
    content: t( i18nKeys.Common.Value ),
    icon: '',
    hideUnder: -1,
    event: ''
  }
]

const entryKey = ref('');

const entryCategory = ref('');

const entryFields = ref([] as string[])

// computed
const availableCategories = computed((): string[] => {
  return projectDataStore.categories.map(c => c.Name);
});

const initialEntry = computed(() => {
  return projectDataStore.entries.find(e => e.Key === props.entryKey)
});

const fieldRows = computed(() => {
  const rows: ResponsiveTableCell[][] = [];
  const category = projectDataStore.categories.find(c => c.Name === entryCategory.value);
  if( !category ) {
    return [];
  }
  category.Fields.forEach(f => {
    rows.push([
      {
        content: f.Field,
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
    ])
  })
  return rows;
})

const entryChanged = computed(() => {
  if( initialEntry.value ) {
    return initialEntry.value.Key !== entryKey.value ||
      initialEntry.value.Category !== entryCategory.value ||
      initialEntry.value.Fields.length !== entryFields.value.length ||
      !initialEntry.value.Fields.every((v, i) => v === entryFields.value[i]);
  }
  return false
})

// watchers
watch(initialEntry, () => {
  if( initialEntry.value ) {
    entryKey.value = initialEntry.value.Key;
    entryCategory.value = initialEntry.value.Category;
    entryFields.value = initialEntry.value.Fields;
  }
});

// methods
function CallSaveEntry() {
  const entry: BibEntry = {
    Key: entryKey.value,
    Category: entryCategory.value,
    Fields: entryFields.value,
    Comment: '',
    CiteNumber: 0,
    Example: ''
  }
  SaveEntry(entry, initialEntry.value ? initialEntry.value.Key : '', props.projectName)
}

function getSlotName(i: number) {
  return `${i}-1`
}

// onload
if( initialEntry.value ) {
  entryKey.value = initialEntry.value.Key;
  entryCategory.value = initialEntry.value.Category;
  entryFields.value = JSON.parse(JSON.stringify(initialEntry.value.Fields));
}
</script>

<style scoped>

</style>
