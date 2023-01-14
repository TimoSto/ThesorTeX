<template>
  <ToolbarAndContent>
    <template #bar>
      <v-toolbar-title>
        {{ t(i18nKeys.MainPage.Welcome) }}
      </v-toolbar-title>
    </template>
    <template #content>
      <div style="padding: 8px 16px">
        <ResponsiveTable
          :headers="projectHeaders"
          :rows="[]"
        >
          <template #h-4>
            <v-btn
              flat
              text
            >
              <v-icon>mdi-plus</v-icon>
            </v-btn>
          </template>
        </ResponsiveTable>
      </div>
    </template>
  </ToolbarAndContent>
</template>

<script lang="ts" setup>
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";
import ResponsiveTable, {ResponsiveTableHeaderCell, SizeClasses} from "../components/ResponsiveTable.vue";

// globals
const { t } = useI18n();

// data

const projectHeaders: ResponsiveTableHeaderCell[] = [
  {
    content: t(i18nKeys.MainPage.Project),
    size: SizeClasses.TwentyFivePercent
  },
  {
    content: t(i18nKeys.MainPage.ProjectCreated),
    size: SizeClasses.TwentyFivePercent
  },
  {
    content: t(i18nKeys.MainPage.ProjectLastModified),
    size: SizeClasses.TwentyFivePercent
  },
  {
    content: t(i18nKeys.MainPage.ProjectNumberOfEntries),
    size: SizeClasses.TenPercent
  },
  {
    slot: true,
    size: SizeClasses.IconBtn
  },
];
</script>

<style scoped>
</style>