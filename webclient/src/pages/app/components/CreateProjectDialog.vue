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
          :rules="[checkName]"
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
            :disabled="name === '' || typeof checkName === 'string'"
            @click="createProject"
            >Erstellen</v-btn>
        </v-row>
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import Vue from "vue";
import { CheckProjectName } from "../rules/projectNameRules";
import {ActionTypes} from "../store/action-types";
import {ProjectData} from "../store/state";

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
  mounted() {
    this.$store.dispatch(ActionTypes.GET_PROJECTS);
  },
  computed: {
    opened: {
      get(): boolean { return this.open},
      set(v:boolean) { this.$emit("stateChange", v)}
    },
    //TODO: global rule-type for these computed funcs
    checkName(): boolean|string {
      return CheckProjectName(this.name, this.$store.state.projects.map((p: ProjectData) => p.Name))
    }
  },
  methods: {
    createProject() {
      this.$store.dispatch(ActionTypes.ADD_PROJECT, this.name);
    }
  }
})
</script>

<style scoped>

</style>