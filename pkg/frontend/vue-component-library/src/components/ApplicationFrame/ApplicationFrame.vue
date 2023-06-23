<script setup lang="ts">
import {ref} from "vue";
import {i18nKeys} from "./i18n/keys";

const props = defineProps({
  hasSidebar: Boolean,
  mainTitle: String,
  documentationTarget: String,
  i18n: {
    type: Function,
    default: (k: string) => k
  }
});

// data
const sidebarOpened = ref(false);
</script>

<template>
  <v-app>
    <v-app-bar
      color="primary"
      elevation="0"
    >
      <v-app-bar-nav-icon
        :disabled="!hasSidebar"
        :title="i18n(sidebarOpened ? i18nKeys.ApplicationFrame.CloseSidebar : i18nKeys.ApplicationFrame.OpenSidebar)"
        @click="sidebarOpened = !sidebarOpened"
      />

      <v-app-bar-title role="heading" aria-level="1">
        {{ mainTitle }}
      </v-app-bar-title>

      <v-spacer />

      <a v-if="documentationTarget" target="_blank"
         :href="`https://thesortex.com/#/tutorials?target=${documentationTarget}`"
         style="color: rgb(var(--v-theme-on-primary))">
        <v-btn icon :title="i18n(i18nKeys.ApplicationFrame.Documentation)">
          <v-icon>mdi-book-open-variant</v-icon>
        </v-btn>
      </a>
    </v-app-bar>
  </v-app>
</template>

<style scoped lang="scss">

</style>