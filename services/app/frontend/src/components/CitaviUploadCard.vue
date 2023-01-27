<template>
  <v-card>
    <v-card-title>Einträge hinzufügen</v-card-title>
    <v-card-text>
      <span v-if="entries.length > 0">
        Folgende Einträge werden hochgeladen und werden bestehende Einsträge mit demselben Schlüssel überschreiben:
      </span>
      <v-list
        v-if="entries.length > 0"
        lines="two"
      >
        <v-list-item
          v-for="i in entries"
          :key="i"
          item
        >
          <v-list-item-title>{{ i.Key }}</v-list-item-title>
          <v-list-item-subtitle>{{ i.Category }}</v-list-item-subtitle>
          <template #append>
            <v-btn
              color="grey-lighten-1"
              icon="mdi-close"
              variant="text"
            />
          </template>
        </v-list-item>
      </v-list>
      <span v-if="unknowns.length > 0">
        Folgende Einträge konnten nicht zugeordnet werden:
      </span>
      <v-list
        v-if="unknowns.length > 0"
        lines="two"
      >
        <v-list-item
          v-for="i in unknowns"
          :key="i"
          item
        >
          <v-list-item-title>{{ i.Key }}</v-list-item-title>
          <v-list-item-subtitle>{{ i.Category }}</v-list-item-subtitle>
        </v-list-item>
      </v-list>
    </v-card-text>
    <v-card-actions>
      <v-spacer />
      <v-btn
        color="primary"
      >
        Abbrechen
      </v-btn>
      <v-btn
        color="primary"
      >
        Hochladen
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts" setup>

import {Entry} from "../domain/entry/Entry";
import {Unknown} from "../domain/citavi/BibAnalytics";

const props = defineProps({
  entries: {
    default: [],
    type: Array<Entry>
  },
  unknowns: {
    default: [],
    type: Array<Unknown>
  },
});
</script>

<style scoped>

</style>