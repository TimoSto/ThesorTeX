<template>
  <ToolbarAndContent>
    <template #bar>
      <v-toolbar-title>
        {{ projectName }}
      </v-toolbar-title>
    </template>
    <template #content>
      <v-card elevation="0">
        <v-card-text style="padding-bottom: 8px;">
          <v-expansion-panels :model-value="0">
            <v-expansion-panel>
              <v-expansion-panel-title>
                Einträge im Literaturverzeichnis
              </v-expansion-panel-title>
              <v-expansion-panel-text>
                <ResponsiveTable
                  :rows="[]"
                  :headers="entriesHeaders"
                >
                  <template #h-3>
                    <v-btn
                      flat
                      text
                    >
                      <v-icon>mdi-plus</v-icon>
                    </v-btn>
                  </template>
                </ResponsiveTable>
              </v-expansion-panel-text>
            </v-expansion-panel>
          </v-expansion-panels>
        </v-card-text>
      </v-card>
      <v-card elevation="0">
        <v-card-text style="padding-top: 8px;">
          <v-expansion-panels :model-value="0">
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
import ResponsiveTable, {ResponsiveTableHeaderCell, SizeClasses} from "../components/ResponsiveTable.vue";

// globals
const appStateStore = useAppStateStore();

// computed
const projectName = computed(() => {
  return appStateStore.currentProject;
});

const entriesHeaders = computed((): ResponsiveTableHeaderCell[] => {
  return [
    {
      content: "Schlüsel",
      size: SizeClasses.MaxWidth250
    },
    {
      content: "Kategorie",
      size: SizeClasses.MaxWidth250
    },
    {
      content: "Eintrag",
      size: SizeClasses.ThirtyPercent
    },
    {
      slot: true,
      size: SizeClasses.IconBtn
    },
  ]
})
</script>

<style scoped>

</style>