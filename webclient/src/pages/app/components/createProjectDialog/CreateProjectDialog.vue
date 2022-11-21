<template>
  <v-dialog
    width="350"
    v-model="opened"
    >
    <v-card>
      <v-card-title>
        Projekt erstellen
      </v-card-title>
      <v-card-text>
        <v-text-field
          color="primary"
          label="Name des Projektes"
          v-model="name"
        />
        <v-row style="padding: 16px 8px 4px 0">
          <v-spacer />
          <v-btn
              color="primary"
              text
              @click="$emit('stateChange', false)"
          >
            Abbrechen
          </v-btn>
          <v-btn
            color="primary"
            >Erstellen</v-btn>
        </v-row>
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  name: "CreateProjectDialog",
  props: [
      'open',
      'existingProjects'
  ],
  data() {
    return {
      name: ''
    }
  },
  computed: {
    opened: {
      get(): boolean { return this.open},
      set(v:boolean) { this.$emit("stateChange", v)}
    },
  }
})
</script>

<style scoped>

</style>