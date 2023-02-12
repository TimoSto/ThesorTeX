<template>
  <div class="bg-container" style="background-image: linear-gradient(90deg, #0c8635, #259b71, #69beaf);">
    <v-container class="bg-transparent pa-16">
      <v-row>
        <v-col cols="12">
          <h2 class="text-h3 font-weight-bold text-center text-white pa-4 pt-15">LaTeX bequemer nutzen</h2>
          <p class="text-h5 text-center text-white pa-4 pb-15">Wenn du LaTeX im akademischen oder beruflichen
            Kontext
            nutzt, findest
            du hier
            verschiedene Vorlagen und Tools, die dich dabei unterstützen können.</p>
        </v-col>
      </v-row>
    </v-container>
  </div>
  <div class="bg-container">
    <v-container class="bg-transparent pb-6 pl-12">
      <v-row>
        <v-col :cols="smallDisplay ? 12 : 6">
          <h2 class="text-h3 font-weight-bold pt-6 pb-6">Vorlage für wissenschaftliche Arbeiten</h2>
          <p class="text-h6 pb-6">
            Möchtest du eine Abschluss- oder Hausarbeit mit LaTeX schreiben? Dann schau dir doch mal <a
            href="/templates/thesis" download>diese
            Vorlage</a> an.</p>
          <v-btn color="primary" to="/docs/template">
            Mehr erfahren
          </v-btn>
        </v-col>
        <v-col v-if="!smallDisplay" cols="6">
          <SVGTemplate :turn="10" :start-point="{x:0, y:50}" :points="thesisPath"
                       style="display: block; margin: 0 auto; max-height: 350px" />
        </v-col>
      </v-row>
    </v-container>
  </div>
  <div class="bg-container" style="background-color: #e8e8e8">
    <v-container class="bg-transparent pb-6 pr-12">
      <v-row>
        <v-col v-if="!smallDisplay" cols="6">
          <TemplateIcon :hide-icon="false" style="display: block; margin: 0 auto; max-height: 350px" />
        </v-col>
        <v-col :cols="smallDisplay ? 12 : 6">
          <h2 class="text-h3 font-weight-bold pt-6 pb-6">Literatur-Management</h2>
          <p class="text-h6 pb-6">Mit diesem Tool kannst du dein Literaturverzeichnis bequem und individuell
            zusammenbauen. Formatiere deine Literatur nach deinen Wünschen, lege Einträge händisch an oder
            importiere deine bibliothek aus <i>Citavi</i>.</p>
          <v-btn color="primary">
            Mehr erfahren
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
  </div>
  <div class="bg-container">
    <v-container class="bg-transparent pb-6 pl-12 mb-6">
      <v-row>
        <v-col :cols="smallDisplay ? 12 : 6">
          <h2 class="text-h3 font-weight-bold pt-6 pb-6">Vorlage für einen Lebenslauf</h2>
          <p class="text-h6 pb-6">
            Wenn du mithilfe der hier angebotenen Unterstützungen deinen Abschluss geschafft hast, kannst du
            dich
            mit <a href="/templates/cv" download>diese
            Vorlage</a> für deien Lebenslauf auf deinen nächsten Job bewerben.</p>
          <v-btn color="primary">
            Mehr erfahren
          </v-btn>
        </v-col>
        <v-col v-if="!smallDisplay" cols="6">
          <CVIcon style="display: block; margin: 0 auto; max-height: 350px" />
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script lang="ts" setup>
import TemplateIcon from "../components/TemplateIcon.vue";
import CVIcon from "../components/CVIcon.vue";

const props = defineProps({
  smallDisplay: Boolean
});

const thesisPath = [
  {
    x: 100,
    y: 50,
  },
  {
    x: 100,
    y: 100
  },
  {
    x: 0,
    y: 100
  },
  {
    x: 0,
    y: 50
  }
];
</script>

<style scoped>

</style>