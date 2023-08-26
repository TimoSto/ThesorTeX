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
    <v-card width="350">
      <v-card-title>{{ i18n(i18nKeys.FeedbackForm.Title) }}</v-card-title>
      <v-card-text>
        <v-text-field width="200"></v-text-field>
        <v-textarea />
      </v-card-text>
    </v-card>
  </v-menu>

</template>

<style scoped lang="scss">

</style>