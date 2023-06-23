<script setup lang="ts">
import {ref} from "vue";
import {i18nKeys} from "./i18n/keys";
import {accessibilityDialogKeys} from "../../index";
import {Document} from "happy-dom";
import AccessibilityDialog from "../AccessibilityDialog/AccessibilityDialog.vue";

const props = defineProps({
  hasSidebar: Boolean,
  mainTitle: String,
  documentationTarget: String,
  i18n: {
    type: Function,
    default: (k: string) => k
  },
  documentProp: {
    type: Document,
    required: false
  },
  showA11y: Boolean
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

      <AccessibilityDialog v-if="showA11y && documentProp?.readyState" :keydownTarget="documentProp as Document"
                           :i18n="i18n" v-slot="scope">
        <v-btn icon
               @click="scope.openDialog" :title="i18n(accessibilityDialogKeys.AccessibilityDialog.BtnTitle)">
          <v-icon>mdi-human</v-icon>
        </v-btn>
      </AccessibilityDialog>

    </v-app-bar>
  </v-app>
</template>

<style scoped lang="scss">

</style>