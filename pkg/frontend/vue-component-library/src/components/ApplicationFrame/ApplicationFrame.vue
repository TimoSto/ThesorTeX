<script setup lang="ts">
const props = defineProps({
  hasSidebar: Boolean,
  mainTitle: String,
  documentationTarget: String
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

      <v-spacer />

      <a v-if="documentationTarget" target="_blank"
         :href="`https://thesortex.com/#/tutorials?target=${documentationTarget}`"
         style="color: rgb(var(--v-theme-on-primary))">
        <v-btn icon>
          <v-icon>mdi-book-open-variant</v-icon>
        </v-btn>
      </a>
    </v-app-bar>
  </v-app>
</template>

<style scoped lang="scss">

</style>