<template>
  <AppbarContent
    :level="2"
    bar-color="background"
  >
    <template #bar>
      <v-app-bar-nav-icon
        @click="emit('navBack')"
      >
        <v-icon>mdi-arrow-left</v-icon>
      </v-app-bar-nav-icon>

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

// globals
const emit = defineEmits(['navBack']);

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
