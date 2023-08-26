<script setup lang="ts">
import {ref} from "vue";
import {i18nKeys} from "./i18n/keys";

const props = defineProps({
  i18n: {
    type: Function,
    default: (k: string) => k
  }
});

const opened = ref(false);
</script>

<template>

  <v-menu :close-on-content-click="false">
    <template v-slot:activator="{ props }">
      <v-btn icon color="primary" v-bind="props">
        <v-icon>mdi-delete</v-icon>
      </v-btn>
    </template>
    <v-card width="550">
      <v-card-title>{{ i18n(i18nKeys.FeedbackForm.Title) }}</v-card-title>
      <v-alert type="warning" density="compact" variant="tonal" :closable="true">
        <v-alert-title>{{ i18n(i18nKeys.FeedbackForm.NoPersonalData) }}</v-alert-title>
      </v-alert>
      <v-card-text>
        <v-text-field width="200" color="primary" :label="i18n(i18nKeys.FeedbackForm.Subject)"></v-text-field>
        <v-textarea color="primary" :label="i18n(i18nKeys.FeedbackForm.Message)" />
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn color="primary">
          {{ i18n(i18nKeys.FeedbackForm.Discard) }}
        </v-btn>
        <v-btn color="primary">
          {{ i18n(i18nKeys.FeedbackForm.Send) }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-menu>

</template>

<style scoped lang="scss">

</style>