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
                {{ t(i18nKeys.CategoryEditor.General) }}
              </v-expansion-panel-title>
              <v-expansion-panel-text>
                <ResponsiveTable
                  :rows="[]"
                  :headers="generalHeaders"
                >
                </ResponsiveTable>
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
                  {{ t(i18nKeys.CategoryEditor.BibEntry) }}
                </v-expansion-panel-title>
                <v-expansion-panel-text>
                  <ResponsiveTable
                    :rows="[]"
                    :headers="entryHeaders"
                  >
                  </ResponsiveTable>
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
                  {{ t(i18nKeys.CategoryEditor.Cites) }}
                </v-expansion-panel-title>
                <v-expansion-panel-text>
                  <ResponsiveTable
                    :rows="[]"
                    :headers="entryHeaders"
                  >
                  </ResponsiveTable>
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
import {useAppStateStore} from "../stores/appState/AppStateStore";
import {computed} from "vue";
import ResponsiveTable, {ResponsiveTableHeaderCell, SizeClasses} from "../components/ResponsiveTable.vue";
import {i18nKeys} from "../i18n/keys";
import {useI18n} from "@thesortex/vue-i18n-plugin"

// globals
const appStateStore = useAppStateStore();

const {t} = useI18n();

// computed
const categoryName = computed(() => {
  return appStateStore.currentItem;
});

const generalHeaders = computed((): ResponsiveTableHeaderCell[] => {
  return [
    {
      content: t(i18nKeys.Common.Attribute),
      size: SizeClasses.MaxWidth250
    },
    {
      content: t(i18nKeys.Common.Value),
      size: ""
    }
  ]
});

const entryHeaders = computed((): ResponsiveTableHeaderCell[] => {
  return [
    {
      content: t(i18nKeys.Common.Attribute),
      size: ""
    },
    {
      content: t(i18nKeys.CategoryEditor.Prefix),
      size: ""
    },
    {
      content: t(i18nKeys.CategoryEditor.Suffix),
      size: ""
    },
    {
      content: t(i18nKeys.CategoryEditor.Italic),
      size: SizeClasses.MaxWidth100
    },
    {
      content: t(i18nKeys.CategoryEditor.Preformatted),
      size: SizeClasses.MaxWidth100
    },
    {
      slot: true,
      size: SizeClasses.IconBtn
    }
  ]
})
</script>

<style scoped>

</style>