<template>
  <ToolbarAndContent>
    <template #bar>
      <v-toolbar-title>
        {{ projectName }}
      </v-toolbar-title>
    </template>
  </ToolbarAndContent>

</template>

<script lang="ts" setup>
import {useAppStateStore} from "../stores/appState/AppStateStore";
import {computed} from "vue";

// globals
const appStateStore = useAppStateStore();

// computed
const projectName = computed(() => {
  return appStateStore.currentProject;
})
</script>

<style scoped>

</style>