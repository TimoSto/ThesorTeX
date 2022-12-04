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
              <tbody>
                <tr
                  v-for="e in projectDataStore.entries"
                  :key="e.Key"
                >
                  <td>{{ e.Key }}</td>
                  <td>{{ e.Category }}</td>
                </tr>
              </tbody>
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
              <tbody>
                <tr
                  v-for="c in projectDataStore.categories"
                  :key="c.Name"
                >
                  <td>{{ c.Name }}</td>
                  <td v-html="c.Example" />
                  <td></td>
                </tr>
              </tbody>
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
import {useProjectDataStore} from "../stores/projectDataStore";
import GetProjectData from "../api/projectData/GetProjectData";
import {ProjectData} from "../api/projectData/ProjectData";

const props = defineProps({
  projectName: {
    type: String,
    required: true
  }
});

const { t } = useI18n();

const projectDataStore = useProjectDataStore();

GetProjectData(props.projectName).then((data: ProjectData) => {
  projectDataStore.setData(data);
})
</script>

<style scoped>

</style>
