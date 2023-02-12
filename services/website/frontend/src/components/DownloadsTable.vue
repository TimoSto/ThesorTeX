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
          {{ v }}
        </td>
        <td>

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

const props = defineProps({
  versions: Array<String>
});
</script>

<style scoped>

</style>