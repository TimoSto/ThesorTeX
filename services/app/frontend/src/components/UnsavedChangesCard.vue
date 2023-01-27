<template>
  <v-card>
    <v-card-title>
      {{ t(i18nKeys.UnsavedChanges.Title) }}
    </v-card-title>
    <v-card-text>
      {{ t(i18nKeys.UnsavedChanges.Message) }}
    </v-card-text>
    <v-card-actions>
      <v-spacer />
      <v-btn
        color="primary"
        @click="emit('resolve', false)"
      >
        {{ t(i18nKeys.Common.Abort) }}
      </v-btn>
      <v-btn
        color="primary"
        @click="emit('resolve', true)"
      >
        {{ t(i18nKeys.Common.Continue) }}
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts" setup>
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";

const {t} = useI18n();

const emit = defineEmits(["resolve"]);
</script>

<style scoped>

</style>