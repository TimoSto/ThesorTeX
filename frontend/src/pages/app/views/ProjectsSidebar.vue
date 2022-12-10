<template>
  <AppbarContent
    :level="2"
    bar-color="background"
  >
    <template #bar>
      <v-toolbar-title>
        {{ t(i18nKeys.Sidebar.Projects) }}
      </v-toolbar-title>
    </template>
    <template #content>
      <v-list
        density="compact"
        :selected="[currentProject]"
        @update:selected="currentProject = $event[0]"
      >
        <v-list-item
          v-for="(p, i) in projects"
          :key="`project-li-${i}`"
          :value="i"
          active-color="primary"
        >
          {{ p.Name }}
        </v-list-item>
      </v-list>
    </template>
  </AppbarContent>
</template>

<script setup lang="ts">
import AppbarContent from "../../../components/AppbarContent.vue";
import {useProjectsStore} from "../stores/projectsStore";
import {computed} from "vue";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";

// globals
const projectsStore = useProjectsStore();

const emit = defineEmits(['switchTo']);

const { t } = useI18n();

// props
const props = defineProps({
  project: {
    type: String,
    required: true
  }
})

// computed
const projects = computed(() => {
  return projectsStore.projects
});

const currentProject = computed({
  get() {
    return projects.value.map(p => p.Name).indexOf(props.project)
  },
  set(v: number) {
    emit('switchTo', v);
  }
})

</script>

<style scoped>

</style>
