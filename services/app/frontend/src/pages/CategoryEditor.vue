<template>
  <ToolbarAndContent>
    <template #bar>
      <v-toolbar-title>{{ categoryName }}</v-toolbar-title>
    </template>
  </ToolbarAndContent>
</template>

<script lang="ts" setup>

// globals
import {useAppStateStore} from "../stores/appState/AppStateStore";
import {computed} from "vue";

const appStateStore = useAppStateStore();

// computed
const categoryName = computed(() => {
  return appStateStore.currentItem;
})
</script>

<style scoped>

</style>