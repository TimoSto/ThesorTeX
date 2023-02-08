<template>
  <v-app>
    <v-app-bar color="primary" :elevation="elevation" style="z-index: 1000">
      <v-toolbar-title>ThesorTeX</v-toolbar-title>
    </v-app-bar>
    <v-main>
      <ToolbarAndContent :hide-bar="true" @scroll="elevation=1" @no-scroll="elevation=0">
        <template #content>
          <div style="width: 200px; height: 1234px; background-color: #0d47a1"></div>
        </template>
      </ToolbarAndContent>
    </v-main>
  </v-app>
</template>

<script lang="ts" setup>
import {ref} from "vue";

// data
const elevation = ref(0);

</script>

<style scoped>

</style>