<template>

  <v-dialog
    v-model="opened"
    width="600"
    >

    <v-card>
      <v-card-title>Konfiguration</v-card-title>
      <v-card-text>

        <v-text-field
          label="Port"
          v-model="config.Port"
          ></v-text-field>

        <v-text-field
            label="Speicherort für Projekte"
            v-model="config.ProjectsDir"
        ></v-text-field>

        <v-row style="padding: 16px 8px 4px 0">
          <v-spacer />

          <v-btn
              text
              color="primary"
              @click="opened = false"
          >
            Schließen
          </v-btn>
          <v-btn
              color="primary"
              :disabled="JSON.stringify(config) === JSON.stringify(initialConfig)"
          >
            Speichern
          </v-btn>
        </v-row>

      </v-card-text>
    </v-card>

  </v-dialog>

</template>

<script lang="ts">
import Vue from "vue";
import GetConfig, { AppConfiguration } from "./api/readConfig";

export default Vue.extend({
  name: "ConfigDialog",
  props: [
      "open"
  ],
  data() {
    return {
      initialConfig: {} as AppConfiguration,
      config: {} as AppConfiguration
    }
  },
  async mounted() {
    this.initialConfig = await GetConfig();
    this.config = {
      Port: this.initialConfig.Port,
      ProjectsDir: this.initialConfig.ProjectsDir,
      Error: this.initialConfig.Error
    };
  },
  computed: {
    opened: {
      get(): boolean { return this.open},
      set(v:boolean) { this.$emit("stateChange", v)}
    }
  }
});
</script>

<style scoped>

</style>