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
        {{ props.categoryName }}
      </v-toolbar-title>
    </template>

    <template #content></template>
  </AppbarContent>
</template>

<script setup lang="ts">
import AppbarContent from "../../../components/AppbarContent.vue";

//globals
const emit = defineEmits(['navBack']);

//props
const props = defineProps({
  categoryName: {
    type: String,
    required: true
  }
})
</script>

<style scoped>

</style>
