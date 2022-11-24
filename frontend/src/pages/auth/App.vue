<template>
  <v-app>
    <v-main>
      <HelloWorld />
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import {useI18n} from "vue-i18n";
import HelloWorld from "../../components/HelloWorld.vue";

const { t } = useI18n()
</script>
