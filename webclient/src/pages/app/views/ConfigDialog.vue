<template>

  <v-dialog
    v-model="opened"
    width="400"
    >

    <v-card>
      <v-card-title>Konfiguration</v-card-title>
      <v-card-text>

        <v-text-field
          label="Port"
          ></v-text-field>

        <v-text-field
            label="Speicherort für Projekte"
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

export default Vue.extend({
  name: "ConfigDialog",
  props: [
      "open"
  ],
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