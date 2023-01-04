<template>
  <v-app>
    <div class="outer">
      <div class="inner">
        <v-btn>{{ t(i18nKeys.Common.Continue) }}</v-btn>
      </div>
    </div>
  </v-app>
</template>

<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {i18nKeys} from "./i18n/keys";

const { t } = useI18n();
</script>

<style scoped lang="scss">
.outer{
  background-color: red;
  padding: 10px;
  & .inner {
    background-color: green;
    padding: 100px;
  }
}
</style>
