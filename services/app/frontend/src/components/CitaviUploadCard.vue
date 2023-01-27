<template>
  <v-card>
    <v-card-title>
      {{ t(i18nKeys.ProjectPage.UploadTitle) }}
    </v-card-title>
    <v-card-text>
      <span v-if="entries.length === 0 && unknowns.length === 0">
        {{ t(i18nKeys.ProjectPage.UploadNoEntriesFound) }}
      </span>
      <span v-if="entries.length > 0">
        {{ t(i18nKeys.ProjectPage.UploadEntries) }}
      </span>
      <v-list
        v-if="entries.length > 0"
        lines="two"
      >
        <v-list-item
          v-for="(e, i) in entries"
          :key="`entry-${i}`"
          item
        >
          <v-list-item-title>{{ e.Key }}</v-list-item-title>
          <v-list-item-subtitle>{{ e.Category }}</v-list-item-subtitle>
          <template #append>
            <v-btn
              color="grey-lighten-1"
              icon="mdi-close"
              variant="text"
              @click="emit('rmEntry', i)"
            />
          </template>
        </v-list-item>
      </v-list>
      <span v-if="unknowns.length > 0">
        {{ t(i18nKeys.ProjectPage.UploadUnknowns) }}
      </span>
      <v-list
        v-if="unknowns.length > 0"
        lines="two"
      >
        <v-list-item
          v-for="(e,i) in unknowns"
          :key="`unknown${i}`"
          item
        >
          <v-list-item-title>{{ e.Key }}</v-list-item-title>
          <v-list-item-subtitle>{{ e.Category }}</v-list-item-subtitle>
        </v-list-item>
      </v-list>
    </v-card-text>
    <v-card-actions>
      <v-spacer />
      <v-btn
        color="primary"
        @click="emit('close')"
      >
        {{ entries.length === 0 && unknowns.length === 0 ? t(i18nKeys.Common.Close) : t(i18nKeys.Common.Abort) }}
      </v-btn>
      <v-btn
        color="primary"
        :disabled="entries.length > 0"
        @click="emit('upload')"
      >
        {{ t(i18nKeys.ProjectPage.Add) }}
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts" setup>
//TODO: unit tests
import {Entry} from "../domain/entry/Entry";
import {Unknown} from "../domain/citavi/BibAnalytics";
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";

defineProps({
  entries: {
    default: [],
    type: Array<Entry>
  },
  unknowns: {
    default: [],
    type: Array<Unknown>
  },
});

const {t} = useI18n();

const emit = defineEmits(["rmEntry", "close", "upload"]);
</script>

<style scoped>

</style>