<template>
  <AppbarContent
    level="2"
    barColor="background">

    <template v-slot:bar>
      <v-toolbar-title>Willkommen bei ThesorTeX!</v-toolbar-title>

      <v-spacer />

      <v-btn icon>
        <v-icon>mdi-help-circle-outline</v-icon>
      </v-btn>
    </template>

    <template v-slot:content>

      <div style="padding: 8px">
        <v-table>
          <thead>
          <tr>
            <th>Projekt</th>
            <th>Erstellt</th>
            <th>Zuletzt bearbeitet</th>
            <th>Anzahl Einträge</th>
            <th style="min-width: 48px; max-width: 48px;">
              <v-btn text flat>
                <v-icon>mdi-plus</v-icon>
              </v-btn>
            </th>
          </tr>
          </thead>
        </v-table>
      </div>

    </template>

  </AppbarContent>
</template>

<script setup lang="ts">
import AppbarContent from "@/components/AppbarContent";

</script>

<style scoped>

</style>
