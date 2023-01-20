<template>
  <ToolbarAndContent>
    <template #bar>
      <v-toolbar-title>{{ categoryName }}</v-toolbar-title>
    </template>
    <template #content>
      <div class="fullsize-card-container">
        <v-card elevation="3">
          <v-expansion-panels :model-value="0">
            <v-expansion-panel>
              <v-expansion-panel-title>
                Allgemeines
              </v-expansion-panel-title>
              <v-expansion-panel-text>
              </v-expansion-panel-text>
            </v-expansion-panel>
          </v-expansion-panels>
        </v-card>
      </div>
      <div class="fullsize-card-container">
        <v-card elevation="3">
          <v-card elevation="3">
            <v-expansion-panels :model-value="0">
              <v-expansion-panel>
                <v-expansion-panel-title>
                  Literatureintrag
                </v-expansion-panel-title>
                <v-expansion-panel-text>
                </v-expansion-panel-text>
              </v-expansion-panel>
            </v-expansion-panels>
          </v-card>
        </v-card>
      </div>
      <div class="fullsize-card-container">
        <v-card elevation="3">
          <v-card elevation="3">
            <v-expansion-panels :model-value="0">
              <v-expansion-panel>
                <v-expansion-panel-title>
                  Zitate
                </v-expansion-panel-title>
                <v-expansion-panel-text>
                </v-expansion-panel-text>
              </v-expansion-panel>
            </v-expansion-panels>
          </v-card>
        </v-card>
      </div>
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