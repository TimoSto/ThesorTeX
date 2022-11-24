<template>
  <v-app>
    <v-main>
      hallo
      {{ t(i18nKeys.Welcome) }}
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {i18nKeys} from "./i18n/keys";

const { t } = useI18n()
</script>
