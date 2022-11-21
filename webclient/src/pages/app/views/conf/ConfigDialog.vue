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
          :rules="[checkPort]"
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
              :disabled="JSON.stringify(config) === JSON.stringify(initialConfig) || !rulesAreMet"
              @click="save"
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
import {AppConfiguration, GetConfig, SaveConfig} from "../../api_calls/config";
import CheckPortValid from "../../rules/portRules";

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
      Error: ''
    };
  },
  computed: {
    opened: {
      get(): boolean { return this.open},
      set(v:boolean) { this.$emit("stateChange", v)}
    },
    rulesAreMet(): boolean {
      return this.checkPort(this.config.Port) === true
    }
  },
  methods: {
    async save() {
      const ok = await SaveConfig(this.config);
      if( ok ) {
        this.initialConfig = {
          Port: this.config.Port,
          ProjectsDir: this.config.ProjectsDir,
          Error: ''
        }
      }
    },
    checkPort(v: string): boolean|string {
      return CheckPortValid(v)
    }
  }
});
</script>

<style scoped>

</style>