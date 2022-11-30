<template>
  <AppbarContent
    :level="2"
    barColor="background">

    <template v-slot:bar>
      <v-toolbar-title>{{ t(i18nKeys.Overview.Welcome) }}</v-toolbar-title>

      <v-spacer />

      <v-btn icon>
        <v-icon>mdi-help-circle-outline</v-icon>
      </v-btn>
    </template>

    <template v-slot:content>

      <div style="padding: 8px">
        <v-table>
          <thead>
          <tr>
            <th>{{ t(i18nKeys.Overview.Project) }}</th>
            <th>{{ t(i18nKeys.Overview.Created) }}</th>
            <th>{{ t(i18nKeys.Overview.LastModified) }}</th>
            <th>{{ t(i18nKeys.Overview.NumberOfEntries) }}</th>
            <th style="min-width: 48px; max-width: 48px;">
              <v-btn text flat @click="open=true" :title="t(i18nKeys.Overview.CreateProject)">
                <v-icon>mdi-plus</v-icon>
              </v-btn>
            </th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="p in projects">
            <td>{{p.Name}}</td>
            <td>{{p.Created}}</td>
            <td>{{ p.LastModified }}</td>
            <td>{{p.NumberOfEntries}}</td>
            <td>
              <v-btn text flat @click="projectToDelete = p.Name" :title="t(i18nKeys.Common.Delete)">
                <v-icon style="color: rgba(var(--v-theme-on-background), 0.45)">mdi-delete</v-icon>
              </v-btn>
            </td>
          </tr>
          </tbody>
        </v-table>
      </div>

    </template>

  </AppbarContent>

  <CreateProjectDialog
    :open="open"
    :projects="projects.map(p => p.Name)"
    v-on:close="open = false"
    v-on:success="AddProjectToList"
  />

  <DeleteProjectDialog
    :open="projectToDelete !== ''"
    :project="projectToDelete"
    v-on:close="projectToDelete = ''"
    v-on:success="RemoveProjectFromList"
  />

  <SuccessErrorDisplay
    :error="errorStore.errorMessage"
    :success="errorStore.successMessage"
    :hint="t(i18nKeys.Errors.ErrorHint)"
    :title="t(i18nKeys.Errors.Title)"
    v-on:errorClosed="errorStore.clean"
    v-on:successClosed="errorStore.clean"
  />
</template>

<script setup lang="ts">
import AppbarContent from "@/components/AppbarContent";
import CreateProjectDialog from "./CreateProjectDialog.vue";
import {ref} from "vue";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";
import {GetProjects} from "../api/projects/GetProjects";
import ProjectData from "../api/projects/ProjectData";
import DeleteProjectDialog from "./DeleteProjectDialog.vue";
import {useErrorSuccessStore} from "../stores/errorSuccessStore";
import SuccessErrorDisplay from "../../../components/SuccessErrorDisplay.vue";

const errorStore = useErrorSuccessStore();

const open = ref(false);

const { t } = useI18n();

const projects = ref([] as ProjectData[])

GetProjects().then((p: ProjectData[]) => {
  projects.value = p;
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

function AddProjectToList(project: ProjectData) {
  projects.value.push(project);

  projects.value.sort((p1, p2) => {
    return p1.Name.toUpperCase() > p2.Name.toUpperCase() ? 1 : -1;
  })
}

</script>

<style scoped>

</style>
