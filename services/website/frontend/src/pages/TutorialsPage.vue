<template>
  <div style="background-image: linear-gradient(90deg, #0c8635, #259b71, #69beaf);">
    <v-container class="bg-transparent pa-16">
      <v-row class="d-flex flex-row">
        <v-col cols="12">
          <h2 class="text-h3 font-weight-bold text-center text-white pa-4">
            {{ t(i18nKeys.TutorialsPage.Title) }}
          </h2>
        </v-col>
      </v-row>
    </v-container>
  </div>
  <v-container>
    <v-expansion-panels>
      <v-expansion-panel>
        <v-expansion-panel-title>
          {{ t(i18nKeys.TutorialsPage.ThesisTemplate) }}
        </v-expansion-panel-title>
        <v-expansion-panel-text>
          <MarkdownToVuetify :file="''" />
        </v-expansion-panel-text>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-title>
          {{ t(i18nKeys.TutorialsPage.ThesisTemplateTech) }}
        </v-expansion-panel-title>
      </v-expansion-panel>
    </v-expansion-panels>
  </v-container>
</template>

<script lang="ts" setup>
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";

// globals
const {t} = useI18n();

</script>

<style scoped lang="scss">
.fill-height-card {
  height: 100%;

  & .v-card-text {
    height: 100%;
  }
}

</style>