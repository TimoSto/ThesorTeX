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
        {{ props.projectName }}
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
        :model-value="[0]"
      >
        <v-expansion-panel>
          <v-expansion-panel-title>
            {{ t(i18nKeys.Project.Entry.Heading) }}
          </v-expansion-panel-title>
          <v-expansion-panel-text>
            <ResponsiveTable
              :headers="entriesTableHeaders"
              :rows="entriesTableRows"
              @row-clicked="emit('openEntry', entries[parseInt($event)].Key)"
            />
          </v-expansion-panel-text>
        </v-expansion-panel>

        <v-expansion-panel>
          <v-expansion-panel-title>
            {{ t(i18nKeys.Project.Category.Heading) }}
          </v-expansion-panel-title>
          <v-expansion-panel-text>
            <ResponsiveTable
              :headers="categoriesTableHeaders"
              :rows="categoriesTableRows"
            />
          </v-expansion-panel-text>
        </v-expansion-panel>
      </v-expansion-panels>
    </template>
  </AppbarContent>
</template>

<script setup lang="ts">
import AppbarContent from "@/components/AppbarContent";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";
import {useProjectDataStore} from "../stores/projectDataStore";
import GetProjectData from "../api/projectData/GetProjectData";
import {ProjectData} from "../api/projectData/ProjectData";
import {computed} from "vue";
import ResponsiveTable, {ResponsiveTableHeaderCell, ResponsiveTableCell} from "../../../components/ResponsiveTable.vue";

//global stuff
const { t } = useI18n();

const projectDataStore = useProjectDataStore();

const emit = defineEmits(['navBack', 'openEntry'])

//props
const props = defineProps({
  projectName: {
    type: String,
    required: true
  }
});

// data
const entriesTableHeaders: ResponsiveTableHeaderCell[] = [
  {
    width: '200px',
    minWidth: '',
    content: t(i18nKeys.Project.Entry.Key),
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: '200px',
    minWidth: '',
    content: t(i18nKeys.Project.Category.Category),
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: 'auto',
    minWidth: '',
    content: t(i18nKeys.Project.Entry.Entry),
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
  }
];

const categoriesTableHeaders: ResponsiveTableHeaderCell[] = [
  {
    width: '200px',
    minWidth: '',
    content: t(i18nKeys.Project.Category.Category),
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: 'auto',
    minWidth: '',
    content: t(i18nKeys.Project.Category.Example),
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
  }
];

// computed
const entries = computed(() => {
  return projectDataStore.entries;
})

const entriesTableRows = computed(() => {
  const r: ResponsiveTableCell[][] = [];
  projectDataStore.entries.forEach(e => {
    r.push([
      {
        content: e.Key,
        icon: '',
        event: '',
        hideUnder: -1
      },
      {
        content: e.Category,
        icon: '',
        event: '',
        hideUnder: -1
      },
      {
        content: e.Example,
        icon: '',
        event: '',
        hideUnder: -1
      },
      {
        content: '',
        icon: '',
        event: '',
        hideUnder: -1
      },
    ])
  })

  return r
});

const categoriesTableRows = computed(() => {
  const r: ResponsiveTableCell[][] = [];
  projectDataStore.categories.forEach(c => {
    r.push([
      {
        content: c.Name,
        icon: '',
        event: '',
        hideUnder: -1
      },
      {
        content: c.Example,
        icon: '',
        event: '',
        hideUnder: -1
      },
      {
        content: '',
        icon: '',
        event: '',
        hideUnder: -1
      },
    ])
  });
  return r;
})

// onload
GetProjectData(props.projectName).then((data: ProjectData) => {
  projectDataStore.setData(data);
});
</script>

<style scoped>

</style>
