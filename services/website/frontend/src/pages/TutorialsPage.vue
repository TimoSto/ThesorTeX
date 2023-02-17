<template>
  <div style="background-image: linear-gradient(90deg, #0c8635, #259b71, #69beaf);">
    <v-container class="bg-transparent pa-16">
      <v-row class="d-flex flex-row">
        <v-col cols="12">
          <h2 class="text-h3 font-weight-bold text-center text-white pa-4 pt-15">Wie kann ich die Vorlagen und Tools
            nutzen?</h2>
          <p class="text-h5 text-center text-white pa-4 pb-15">Hier wird ihre Verwendung und Anpassung beschrieben.</p>
        </v-col>
      </v-row>
    </v-container>
  </div>
  <v-container>
    <v-expansion-panels>
      <v-expansion-panel>
        <v-expansion-panel-title>Verwendung der Vorlage für wissenschaftliches Arbeiten</v-expansion-panel-title>
        <v-expansion-panel-text>

        </v-expansion-panel-text>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-title>Beschreibung der Vorlage für wissenschaftliches Arbeiten</v-expansion-panel-title>
      </v-expansion-panel>
    </v-expansion-panels>
  </v-container>
</template>

<script>
export default {
  name: "TutorialsPage"
}
</script>

<style scoped lang="scss">
.fill-height-card {
  height: 100%;

  & .v-card-text {
    height: 100%;
  }
}

</style>