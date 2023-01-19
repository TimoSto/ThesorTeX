<template>
  <ToolbarAndContent>
    <template #bar>
      <v-toolbar-title>
        {{ projectName }}
      </v-toolbar-title>
    </template>
    <template #content>
      <v-card elevation="0">
        <v-card-text>
          <v-expansion-panels>
            <v-expansion-panel>
              <v-expansion-panel-title>
                Einträge im Literaturverzeichnis
              </v-expansion-panel-title>
              <v-expansion-panel-text>
                <div style="width: 200px; height: 1000px; border: 1px solid black"></div>
              </v-expansion-panel-text>
            </v-expansion-panel>
          </v-expansion-panels>
        </v-card-text>
      </v-card>
      <v-card elevation="0">
        <v-card-text>
          <v-expansion-panels>
            <v-expansion-panel>
              <v-expansion-panel-title>
                Einträge im Literaturverzeichnis
              </v-expansion-panel-title>
              <v-expansion-panel-text>
                <div style="width: 200px; height: 1000px; border: 1px solid black"></div>
              </v-expansion-panel-text>
            </v-expansion-panel>
          </v-expansion-panels>
        </v-card-text>
      </v-card>
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