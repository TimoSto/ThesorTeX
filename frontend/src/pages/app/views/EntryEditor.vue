<template>
  <AppbarContent
    :level="2"
    bar-color="background"
  >
    <template #bar>
      <v-toolbar-title>
        {{ props.entryKey }}
      </v-toolbar-title>

      <v-spacer />

      <v-btn icon>
        <v-icon>mdi-delete</v-icon>
      </v-btn>
    </template>

    <template #content>
    entry
    </template>
  </AppbarContent>
</template>

<script setup lang="ts">

import AppbarContent from "../../../components/AppbarContent.vue";

// props
const props = defineProps({
  entryKey: {
    type: String,
    required: true
  }
})
</script>

<style scoped>

</style>
