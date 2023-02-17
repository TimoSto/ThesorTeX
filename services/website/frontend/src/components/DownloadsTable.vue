<template>
  <v-table>
    <thead>
      <tr>
        <th>
          {{ t(i18nKeys.DownloadPage.Version) }}
        </th>
        <th>
          {{ t(i18nKeys.DownloadPage.Date) }}
        </th>
        <th>Windows</th>
        <th>Linux</th>
        <th>MacOS</th>
        <th>MacOS (Silicon)</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="v in versions" :key="v">
        <td>
          {{ v.Name }}
        </td>
        <td>
          {{ v.Date }}
        </td>
        <td>
          <a :href="getDownloadLink(v.Name, 'windows')"
             download="ThesorTeX.zip">
            <v-btn variant="flat" style="color: rgba(var(--v-theme-on-background), 1)">
              <v-icon>mdi-download</v-icon>
            </v-btn>
          </a>
        </td>
        <td>
          <a :href="getDownloadLink(v.Name, 'linux')"
             download="ThesorTeX.zip">
            <v-btn variant="flat" style="color: rgba(var(--v-theme-on-background), 1)">
              <v-icon>mdi-download</v-icon>
            </v-btn>
          </a>
        </td>
        <td>
          <a :href="getDownloadLink(v.Name, 'mac')"
             download="ThesorTeX.zip">
            <v-btn variant="flat" style="color: rgba(var(--v-theme-on-background), 1)">
              <v-icon>mdi-download</v-icon>
            </v-btn>
          </a>
        </td>
        <td>
          <a :href="getDownloadLink(v.Name, 'mac_silicon')"
             download="ThesorTeX.zip">
            <v-btn variant="flat" style="color: rgba(var(--v-theme-on-background), 1)">
              <v-icon>mdi-download</v-icon>
            </v-btn>
          </a>
        </td>
      </tr>
    </tbody>
  </v-table>
</template>

<script lang="ts" setup>
import getDownloadLink from "../api/GetDownloadLink";
import {VersionInfo} from "../api/GetToolVersions";
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";

const {t} = useI18n();

const props = defineProps({
  versions: Array<VersionInfo>
});
</script>

<style scoped>
a {
  text-decoration: none;
}
</style>