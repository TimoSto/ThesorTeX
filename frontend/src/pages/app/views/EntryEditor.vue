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

      <v-btn icon>
        <v-icon>mdi-delete</v-icon>
      </v-btn>
    </template>

    <template #content>
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
              <template #0>
                <v-text-field
                  color="primary"
                  variant="underlined"
                  :model-value="entryKey"
                />
              </template>
              <template #1>
                <v-select
                  color="primary"
                  variant="underlined"
                  :items="availableCategories"
                  :model-value="entryCategory"
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
                #[i]
              >
                <v-text-field
                  color="primary"
                  variant="underlined"
                  :model-value="entryFields[i]"
                />
              </template>
            </ResponsiveTable>
          </v-expansion-panel-text>
        </v-expansion-panel>
      </v-expansion-panels>
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

// globals
const emit = defineEmits(['navBack']);

const { t } = useI18n();

const projectDataStore = useProjectDataStore();

// props
const props = defineProps({
  entryKey: {
    type: String,
    required: true
  }
});

// data
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
]

const generalRows: ResponsiveTableCell[][] = [
  [
    {
      content: 'Schlüssel',
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
      content: 'Kategorie',
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
    content: 'Feld',
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

// watchers
watch(initialEntry, () => {
  if( initialEntry.value ) {
    entryKey.value = initialEntry.value.Key;
    entryCategory.value = initialEntry.value.Category;
    entryFields.value = initialEntry.value.Fields;
  }
});

// onload
if( initialEntry.value ) {
  entryKey.value = initialEntry.value.Key;
  entryCategory.value = initialEntry.value.Category;
  entryFields.value = initialEntry.value.Fields;
}
</script>

<style scoped>

</style>
