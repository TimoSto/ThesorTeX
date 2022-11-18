<template>
  <div>

    <v-app-bar
      color="background"
      dense
      elevate-on-scroll
      scroll-target="#index-content"
      absolute
      >
      <v-toolbar-title>Willkommen bei ThesorTeX!</v-toolbar-title>

      <v-spacer />

      <v-btn icon>
        <v-icon>mdi-help-circle-outline</v-icon>
      </v-btn>

    </v-app-bar>

    <ContentBelowBar
        bar-height="112"
        margin-top="48"
        padding-x="16"
        padding-y="0"
        scroll-id="index-content"
    >

      <v-simple-table>
        <template v-slot:default>
          <thead>
            <tr>
              <th >Projekt</th>
              <th >Erstellt</th>
              <th >Zuletzt bearbeitet</th>
              <th >Anzahl Einträge</th>
              <th style="min-width: 48px; max-width: 48px;">
                <v-btn icon>
                  <v-icon>mdi-plus</v-icon>
                </v-btn>
              </th>
            </tr>
          </thead>
        </template>
      </v-simple-table>
    </ContentBelowBar>

  </div>
</template>

<script lang="ts">
import Vue from "vue";
import ContentBelowBar from "../../../common/components/ContentBelowBar.vue";

export default Vue.extend({
  name: "StartPage",
  components: {ContentBelowBar}
})
</script>

<style scoped>

</style>