<template>
  <v-breadcrumbs :items="pages">
    <template v-slot:title="{ item, index }">
          <span
            v-on:click="$emit('goBackTo', index)"
            style="cursor: pointer; font-size: 20px;"
            v-if="index < pages.length - 1"
            v-ripple
          >
            {{ item.title }}
          </span>
          <span
            style="cursor: pointer; font-size: 20px;"
            v-if="index === pages.length - 1"
          >
            {{ item.title }}
          </span>
    </template>
    <template v-slot:divider>
      <v-icon icon="mdi-chevron-right"></v-icon>
    </template>
  </v-breadcrumbs>
</template>

<script setup lang="ts">

import {PropType} from "vue";

const props = defineProps({
  pages: {
    type: Array as PropType<Array<{title: string, disabled: boolean}>>,
    required: true
  }
})

</script>

<style scoped>

</style>
