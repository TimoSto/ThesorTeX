<script setup lang="ts">
import {ref} from "vue";
import {i18nKeys} from "./i18n/keys";
import {accessibilityDialogKeys} from "../../index";
import AccessibilityDialog from "../AccessibilityDialog/AccessibilityDialog.vue";
import PageNavigator from "./PageNavigator/PageNavigator.vue";
import {useApplicationStateStore} from "../../stores/ApplicationStateStore/ApplicationStateStore";

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
  showA11y: Boolean,
  hasConfig: Boolean,
  titleAppendix: String,
  configChangesToSave: Boolean,
});

const applicationStateStore = useApplicationStateStore();

const emit = defineEmits(["saveConfig"]);

// data
const sidebarOpened = ref(false);
const configOpened = ref(false);

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

      <v-btn icon v-if="applicationStateStore.history.length > 1" @click="applicationStateStore.goBack(1)">
        <v-icon>mdi-arrow-left</v-icon>
      </v-btn>

      <v-app-bar-title role="heading" aria-level="1">
        {{ mainTitle }}{{ titleAppendix }}
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

      <v-btn v-if="hasConfig" icon :title="i18n(i18nKeys.ApplicationFrame.OpenConfig)" @click="configOpened=true">
        <v-icon>mdi-cog</v-icon>
      </v-btn>

    </v-app-bar>

    <v-navigation-drawer
      :permanent="true"
      :rail="!sidebarOpened"
      :rail-width="58"
    >
      <slot name="sidebar"></slot>
    </v-navigation-drawer>

    <v-main>
      <PageNavigator>
        <template v-for="(e,n) in applicationStateStore.history" #[n]="{openPage, goBack}">
          <slot :name="n" :openPage="openPage" :goBack="goBack"></slot>
        </template>
      </PageNavigator>
    </v-main>

    <v-dialog v-model="configOpened" width="500" height="500">
      <v-card>
        <v-card-title>{{ i18n(i18nKeys.ApplicationFrame.ConfigTitle) }}</v-card-title>
        <v-card-text>
          <slot name="config"></slot>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn color="primary" @click="configOpened = false">
            {{ i18n(i18nKeys.ApplicationFrame.CloseConfig) }}
          </v-btn>
          <v-btn color="primary" @click="emit('saveConfig')" :disabled="!configChangesToSave">
            {{ i18n(i18nKeys.ApplicationFrame.SaveConfig) }}
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-app>
</template>

<style scoped lang="scss">

</style>