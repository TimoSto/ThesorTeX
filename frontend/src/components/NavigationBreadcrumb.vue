<template>
  <v-breadcrumbs :items="props.pages">
    <template #title="{ item, index }">
      <span
        v-if="index < props.pages.length - 1"
        v-ripple
        style="cursor: pointer; font-size: 20px;"
        @click="$emit('goBackTo', index)"
      >
        {{ item.title }}
      </span>
      <span
        v-if="index === pages.length - 1"
        style="cursor: pointer; font-size: 20px;"
      >
        {{ item.title }}
      </span>
    </template>
    <template #divider>
      <v-icon icon="mdi-chevron-right" />
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

defineEmits(['goBackTo'])

</script>

<style scoped>

</style>
