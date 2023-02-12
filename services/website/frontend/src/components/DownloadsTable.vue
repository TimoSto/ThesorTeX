<template>
  <v-table>
    <thead>
      <tr>
        <th>Version</th>
        <th>Datum</th>
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
          <v-btn variant="flat">
            <v-icon>mdi-download</v-icon>
          </v-btn>
        </td>
        <td>
          <v-btn variant="flat">
            <v-icon>mdi-download</v-icon>
          </v-btn>
        </td>
        <td>
          <v-btn variant="flat">
            <v-icon>mdi-download</v-icon>
          </v-btn>
        </td>
        <td>
          <v-btn variant="flat">
            <v-icon>mdi-download</v-icon>
          </v-btn>
        </td>
      </tr>
    </tbody>
  </v-table>
</template>

<script lang="ts" setup>

import {VersionInfo} from "../api/GetToolVersions";

const props = defineProps({
  versions: Array<VersionInfo>
});
</script>

<style scoped>

</style>