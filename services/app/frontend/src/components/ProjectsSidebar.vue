<template>
  <ToolbarAndContent
    :hide-bar="!open"
  >
    <template #bar>
      <v-toolbar-title>Projects</v-toolbar-title>
    </template>
    <template #content>
      <v-list
        :selected="[currentProject]"
        @update:selected="currentProject = $event[0]"
      >
        <v-list-item
          v-for="(p, i) in projects"
          :key="`project-li-${i}`"
          :value="i"
          active-color="primary"
          style="text-overflow: ellipsis"
        >
          {{ p }}
        </v-list-item>
      </v-list>
    </template>
  </ToolbarAndContent>
</template>

<script lang="ts" setup>
import {useProjectsListStore} from "../stores/projectsList/ProjectsListStore";
import {computed} from "vue";

// globals
const projectsListStore = useProjectsListStore();

const emit = defineEmits(["switchTo"]);

// props
const props = defineProps({
  project: {
    type: String,
    required: true
  },
  open: {
    type: Boolean,
    required: true
  }
});

// computed
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