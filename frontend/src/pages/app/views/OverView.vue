<template>
  <AppbarContent
    :level="2"
    bar-color="background"
  >
    <template #bar>
      <v-toolbar-title>{{ t(i18nKeys.Overview.Welcome) }}</v-toolbar-title>

      <v-spacer />

      <v-btn icon>
        <v-icon>mdi-help-circle-outline</v-icon>
      </v-btn>
    </template>

    <template #content>
      <div style="padding: 8px 16px">
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

  <DeleteProjectDialog
    :open="projectToDelete !== ''"
    :project="projectToDelete"
    @close="projectToDelete = ''"
    @success="RemoveProjectFromList"
  />

  <!--todo: move up -->
  <SuccessErrorDisplay
    :error="errorStore.errorMessage"
    :success="errorStore.successMessage"
    :hint="t(i18nKeys.Errors.ErrorHint)"
    :title="t(i18nKeys.Errors.Title)"
    @error-closed="errorStore.clean"
    @success-closed="errorStore.clean"
  />
</template>

<script setup lang="ts">
import AppbarContent from "@/components/AppbarContent";
import CreateProjectDialog from "./CreateProjectDialog.vue";
import {computed, ref} from "vue";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";
import {GetProjects} from "../api/projects/GetProjects";
import ProjectOverviewData from "../api/projects/ProjectOverviewData";
import DeleteProjectDialog from "./DeleteProjectDialog.vue";
import {useErrorSuccessStore} from "../stores/errorSuccessStore";
import SuccessErrorDisplay from "../../../components/SuccessErrorDisplay.vue";
import ResponsiveTable, {ResponsiveTableCell, ResponsiveTableHeaderCell} from "../../../components/ResponsiveTable.vue";

// globals
const { t } = useI18n();

const emit = defineEmits(['openProject'])

const errorStore = useErrorSuccessStore();

// data
const open = ref(false);

const projects = ref([] as ProjectOverviewData[])

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
        icon: 'mdi-delete',
        event: 'deleteProject',
        hideUnder: -1
      },
    ])
  })
  return r;
})

// onload
GetProjects().then((p: ProjectOverviewData[]) => {
  projects.value = !p ? [] : p;
})

const projectToDelete = ref('');

function RemoveProjectFromList() {
  const index = projects.value.map(p => p.Name).indexOf(projectToDelete.value);
  console.log(index)
  if( index > -1 ) {
    projects.value.splice(index, 1);
    projectToDelete.value = '';
  }
}

function AddProjectToList(project: ProjectOverviewData) {
  projects.value.push(project);

  projects.value.sort((p1, p2) => {
    return p1.Name.toUpperCase() > p2.Name.toUpperCase() ? 1 : -1;
  })
}

function handleEvent(evt: string) {
  if( evt === 'addProject' ) {
    open.value = true;
  }
}

</script>

<style scoped>

</style>
