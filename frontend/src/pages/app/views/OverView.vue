<template>
  <AppbarContent
    :level="2"
    bar-color="background"
  >
    <template #bar>
      <v-toolbar-title>{{ t(i18nKeys.Overview.Welcome) }}</v-toolbar-title>

      <v-spacer />
    </template>

    <template #content>
      <div style="padding: 8px 16px; max-height: calc(100vh - 112px);">
        <ResponsiveTable
          :headers="projectHeaders"
          :rows="projectRows"
          @row-clicked="emit('openProject', projects[parseInt($event)].Name)"
          @btn-clicked="handleEvent"
        />
      </div>
    </template>
  </AppbarContent>

  <CreateProjectDialog
    :open="open"
    :projects="projects.map(p => p.Name)"
    @close="open = false"
    @success="AddProjectToList"
  />
</template>

<script setup lang="ts">
import AppbarContent from "../../../components/AppbarContent";
import CreateProjectDialog from "./CreateProjectDialog.vue";
import {computed, ref} from "vue";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";
import {GetProjects} from "../api/projects/GetProjects";
import ProjectMetaData from "../api/projects/ProjectMetaData";
import {useErrorSuccessStore} from "../stores/errorSuccessStore";
import ResponsiveTable, {ResponsiveTableCell, ResponsiveTableHeaderCell} from "../../../components/ResponsiveTable.vue";
import {useProjectsStore} from "../stores/projectsStore";

// globals
const { t } = useI18n();

const emit = defineEmits(['openProject'])

const projectsStore = useProjectsStore();

// data
const open = ref(false);

const projectHeaders: ResponsiveTableHeaderCell[] = [
  {
    width: 'auto',
    minWidth: '',
    content: t(i18nKeys.Overview.Project),
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: 'auto',
    minWidth: '',
    content: t(i18nKeys.Overview.Created),
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: 'auto',
    minWidth: '',
    content: t(i18nKeys.Overview.LastModified),
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: 'auto',
    minWidth: '',
    content: t(i18nKeys.Overview.NumberOfEntries),
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
    event: 'addProject'
  },
];

// computed
const projects = computed(() => {
  return projectsStore.projects;
});

const projectRows = computed(() => {
  const r: ResponsiveTableCell[][] = [];
  projects.value.forEach(p => {
    r.push([
      {
        content: p.Name,
        icon: '',
        event: '',
        hideUnder: -1
      },
      {
        content: p.Created,
        icon: '',
        event: '',
        hideUnder: -1
      },
      {
        content: p.LastModified,
        icon: '',
        event: '',
        hideUnder: -1
      },
      {
        content: '' + p.NumberOfEntries,
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
  return r;
})

// onload
GetProjects().then((p: ProjectMetaData[]) => {
  const v = !p ? [] : p;
  projectsStore.setProjects(v);
})

function AddProjectToList(project: ProjectMetaData) {
  projectsStore.addProject(project);
}

function handleEvent(evt: string) {
  if( evt === 'addProject' ) {
    open.value = true;
  }
}

</script>

<style scoped>

</style>
