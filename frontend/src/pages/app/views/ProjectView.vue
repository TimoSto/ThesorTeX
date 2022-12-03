<template>
  <AppbarContent
    :level="2"
    bar-color="background"
  >
    <template #bar>
      <v-toolbar-title>
        {{ props.projectName }}
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
        :model-value="[0]"
      >
        <v-expansion-panel>
          <v-expansion-panel-title>
            {{ t(i18nKeys.Project.Entry.Heading) }}
          </v-expansion-panel-title>
          <v-expansion-panel-text>
            <v-table>
              <thead>
                <tr>
                  <th style="width: 200px;">
                    {{ t(i18nKeys.Project.Entry.Key) }}
                  </th>
                  <th style="width: 200px;">
                    {{ t(i18nKeys.Project.Category.Category) }}
                  </th>
                  <th>
                    {{ t(i18nKeys.Project.Entry.Entry) }}
                  </th>
                  <th style="width: 48px">
                    <v-btn
                      text
                      flat
                    >
                      <v-icon>mdi-plus</v-icon>
                    </v-btn>
                  </th>
                </tr>
              </thead>
            </v-table>
          </v-expansion-panel-text>
        </v-expansion-panel>

        <v-expansion-panel>
          <v-expansion-panel-title>
            {{ t(i18nKeys.Project.Category.Heading) }}
          </v-expansion-panel-title>
          <v-expansion-panel-text>
            <v-table>
              <thead>
                <tr>
                  <th style="width: 200px;">
                    {{ t(i18nKeys.Project.Category.Category) }}
                  </th>
                  <th>
                    {{ t(i18nKeys.Project.Category.Example) }}
                  </th>
                  <th style="width: 48px">
                    <v-btn
                      text
                      flat
                    >
                      <v-icon>mdi-plus</v-icon>
                    </v-btn>
                  </th>
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
import AppbarContent from "@/components/AppbarContent";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";

const props = defineProps({
  projectName: {
    type: String,
    required: true
  }
});

const { t } = useI18n()
</script>

<style scoped>

</style>
