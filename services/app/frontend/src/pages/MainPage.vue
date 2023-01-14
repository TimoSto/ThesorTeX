<template>
  <ToolbarAndContent>
    <template #bar>
      <v-toolbar-title>
        {{ t(i18nKeys.MainPage.Welcome) }}
      </v-toolbar-title>
    </template>
    <template #content>
      <div style="padding: 8px 16px; max-height: calc(100vh - 112px);">
        hallo
        <div style="height: 1250px; background-color: red">
          gre
        </div>
      </div>
    </template>
  </ToolbarAndContent>
</template>

<script lang="ts" setup>
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";

// globals
const { t } = useI18n();
</script>

<style scoped>
</style>