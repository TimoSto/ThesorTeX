<template>
  <v-table style="border: 1px solid rgba(var(--v-theme-on-background), 0.25); border-radius: 4px; margin: 0 auto;"
           :style="`max-width: ${maxWidth}px;`" theme="light">
    <thead>
      <tr>
        <th class="keepAlign">
          {{ t(i18nKeys.DownloadPage.Version) }}
        </th>
        <th class="keepAlign" style="min-width: 115px">
          {{ t(i18nKeys.DownloadPage.Date) }}
        </th>
        <th v-if="perOs">Windows</th>
        <th v-if="perOs">Linux</th>
        <th v-if="perOs">MacOS (AMD)</th>
        <th v-if="perOs">MacOS (ARM)</th>
        <th v-if="!perOs">Artefakt</th>
        <th>
          Release-Notes
        </th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="v in versions" :key="v">
        <td class="keepAlign">
          {{ v.Name }}
        </td>
        <td class="keepAlign">
          {{ v.Date }}
        </td>
        <td v-if="perOs">
          <v-btn variant="flat" :href="downloadFunc(v.Name, 'windows')"
                 download="ThesorTeX.zip" style="color: rgba(var(--v-theme-on-background), 1)">
            <v-icon color="black">mdi-download</v-icon>
          </v-btn>
        </td>
        <td v-if="perOs">
          <v-btn variant="flat" :href="downloadFunc(v.Name, 'linux')"
                 download="ThesorTeX.zip" style="color: rgba(var(--v-theme-on-background), 1)">
            <v-icon color="black">mdi-download</v-icon>
          </v-btn>
        </td>
        <td v-if="perOs">
          <v-btn variant="flat" :href="downloadFunc(v.Name, 'mac')"
                 download="ThesorTeX.zip" style="color: rgba(var(--v-theme-on-background), 1)">
            <v-icon color="black">mdi-download</v-icon>
          </v-btn>
        </td>
        <td v-if="perOs">
          <v-btn variant="flat" :href="downloadFunc(v.Name, 'mac_arm')"
                 download="ThesorTeX.zip" style="color: rgba(var(--v-theme-on-background), 1)">
            <v-icon color="black">mdi-download</v-icon>
          </v-btn>
        </td>
        <td v-if="!perOs">
          <v-btn variant="flat" :href="downloadFunc(v.Name)"
                 download="ThesorTeX.zip" style="color: rgba(var(--v-theme-on-background), 1)">
            <v-icon color="black">mdi-download</v-icon>
          </v-btn>
        </td>
        <td>
          <v-btn variant="flat" style="color: rgba(var(--v-theme-on-background), 1)" @click="emit('openNotes', v.Name)">
            <v-icon color="black">mdi-file-outline</v-icon>
          </v-btn>
        </td>
      </tr>
    </tbody>
  </v-table>
</template>

<script lang="ts" setup>
import {VersionInfo} from "../api/GetToolVersions";
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";

const {t} = useI18n();

const props = defineProps({
  perOs: Boolean,
  downloadFunc: {
    type: Function,
    required: true
  },
  maxWidth: Number,
  versions: Array<VersionInfo>
});

const emit = defineEmits(["openNotes"]);
</script>

<style scoped lang="scss">
a {
  text-decoration: none;
}

td:not(.keepAlign), th:not(.keepAlign) {
  text-align: center !important;
}
</style>