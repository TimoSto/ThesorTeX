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
      <v-expansion-panels
        multiple
        variant="accordion"
        :model-value="[0,1]"
      >
        <v-expansion-panel>
          <v-expansion-panel-title>
            {{ t(i18nKeys.EntryEditor.General.Title) }}
          </v-expansion-panel-title>
          <v-expansion-panel-text>
            <v-table>
              <thead>
                <tr>
                  <th>Attribut</th>
                  <th>Wert</th>
                </tr>
              </thead>
            </v-table>
          </v-expansion-panel-text>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-title>
            {{ t(i18nKeys.EntryEditor.Fields.Title) }}
          </v-expansion-panel-title>
          <v-expansion-panel-text>
            <v-table>
              <thead>
                <tr>
                  <th>Attribut</th>
                  <th>Wert</th>
                </tr>
              </thead>
            </v-table>
          </v-expansion-panel-text>
        </v-expansion-panel>
      </v-expansion-panels>
    </template>
  </AppbarContent>
</template>

<script setup lang="ts">
import AppbarContent from "../../../components/AppbarContent.vue";
import {i18nKeys} from "../i18n/keys";
import {useI18n} from "vue-i18n";

// globals
const emit = defineEmits(['navBack']);

const { t } = useI18n();

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
