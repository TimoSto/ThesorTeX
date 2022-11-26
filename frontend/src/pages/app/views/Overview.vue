<template>
  <AppbarContent
    :level="2"
    barColor="background">

    <template v-slot:bar>
      <v-toolbar-title>{{ t(i18nKeys.Overview.Welcome) }}</v-toolbar-title>

      <v-spacer />

      <v-btn icon>
        <v-icon>mdi-help-circle-outline</v-icon>
      </v-btn>
    </template>

    <template v-slot:content>

      <div style="padding: 8px">
        <v-table>
          <thead>
          <tr>
            <th>{{ t(i18nKeys.Overview.Project) }}</th>
            <th>{{ t(i18nKeys.Overview.Created) }}</th>
            <th>{{ t(i18nKeys.Overview.LastEdited) }}</th>
            <th>{{ t(i18nKeys.Overview.NumberOfEntries) }}</th>
            <th style="min-width: 48px; max-width: 48px;">
              <v-btn text flat @click="open=true" :title="t(i18nKeys.Overview.CreateProject)">
                <v-icon>mdi-plus</v-icon>
              </v-btn>
            </th>
          </tr>
          </thead>
        </v-table>
      </div>

    </template>

  </AppbarContent>

  <CreateProjectDialog
    :open="open"
    v-on:close="open = false"
  />
</template>

<script setup lang="ts">
import AppbarContent from "@/components/AppbarContent";
import CreateProjectDialog from "./CreateProjectDialog.vue";
import {ref} from "vue";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";

const open = ref(false);

const { t } = useI18n();

</script>

<style scoped>

</style>
