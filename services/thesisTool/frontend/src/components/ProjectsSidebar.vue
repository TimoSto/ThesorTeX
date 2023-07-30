<template>
  <ToolbarAndContent
    :hide-bar="!sidebarOpened"
  >
    <template #bar>
      <v-toolbar-title>{{ t(i18nKeys.App.ProjectsListShort) }}</v-toolbar-title>
    </template>
    <template #content>
      <v-list
        :selected="[currentProject]"
        :title="t(i18nKeys.App.ProjectsList)"
        @update:selected="currentProject = $event[0]"
      >
        <v-list-item
          v-for="(p, i) in projects"
          :key="`project-li-${i}`"
          :value="i"
          active-color="primary"
          style="text-overflow: ellipsis"
        >
          <template #prepend>
            <v-icon :title="!sidebarOpened ? p : ''">mdi-folder</v-icon>
          </template>
          {{ p }}
        </v-list-item>
      </v-list>
    </template>
  </ToolbarAndContent>
</template>

<script lang="ts" setup>
import {useProjectsListStore} from "../stores/projectsList/ProjectsListStore";
import {computed} from "vue";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";
import {
  useApplicationStateStore
} from "@thesortex/vue-component-library/src/stores/ApplicationStateStore/ApplicationStateStore";

//TODO: unit test
// globals
const projectsListStore = useProjectsListStore();

const applicationStateStore = useApplicationStateStore();

const emit = defineEmits(["switchTo"]);

const {t} = useI18n();

// props
const props = defineProps({
  project: {
    type: String,
    required: true
  },
});

// computed
const sidebarOpened = computed(() => {
  return applicationStateStore.sidebarOpened;
});

const projects = computed(() => {
  return projectsListStore.projects.map(p => p.Name);
});

const currentProject = computed({
  get() {
    return projects.value.indexOf(props.project);
  },
  set(v: number) {
    if (v !== undefined) {
      emit("switchTo", v);
    }
  }
});
</script>

<style scoped>

</style>