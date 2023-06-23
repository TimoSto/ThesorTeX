<script setup lang="ts">
const props = defineProps({
  hasSidebar: Boolean,
  mainTitle: String,
});
</script>

<template>
  <v-app>
    <v-app-bar
      color="primary"
      elevation="0"
    >
      <v-app-bar-nav-icon
        :disabled="!hasSidebar"
      />

      <v-app-bar-title role="heading" aria-level="1">
        {{ mainTitle }}
      </v-app-bar-title>
    </v-app-bar>
  </v-app>
</template>

<style scoped lang="scss">

</style>