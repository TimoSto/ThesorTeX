<template>
  <div class="bg-container" style="background-image: linear-gradient(90deg, #0c8635, #259b71, #69beaf);">
    <v-container class="bg-transparent pa-16">
      <v-row>
        <v-col cols="12">
          <h2 class="text-h3 font-weight-bold text-center text-white pa-4 pt-15">{{ t(i18nKeys.StartPage.Title) }}</h2>
          <p class="text-h5 text-center text-white pa-4 pb-15">{{ t(i18nKeys.StartPage.Subtitle) }}</p>
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
        <v-col v-if="!smallDisplay" cols="6" class="svg-container">
          <SVGTemplate :svg="thesisPaths" style="max-height: 300px; display: block; margin: 0 auto;" />
        </v-col>
      </v-row>
    </v-container>
  </div>
  <div class="bg-container" style="background-color: #e8e8e8">
    <v-container class="bg-transparent pb-6 pr-12">
      <v-row>
        <v-col v-if="!smallDisplay" cols="6" class="svg-container">
          <SVGTemplate :svg="laptopWithThesis" style="max-height: 300px; display: block; margin: 0 auto;" />
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
        <v-col v-if="!smallDisplay" cols="6" class="svg-container">
          <SVGTemplate :svg="cvSVG" style="max-height: 300px; display: block; margin: 0 auto;" />
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script lang="ts" setup>
import {computed} from "vue";
import {ThesisSVG} from "../components/svgs/ThesisSVG";
import {LaptopSVG} from "../components/svgs/LaptopSVG";
import {CVSVG} from "../components/svgs/CVSVG";
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";

// globals
const {t} = useI18n();

// data
const props = defineProps({
  smallDisplay: Boolean
});

// computed
const thesisPaths = computed(() => {
  //TODO: find a better way to loose reactivity
  const svg = JSON.parse(JSON.stringify(ThesisSVG));

  return svg;
});

const laptopWithThesis = computed(() => {
  const svg = JSON.parse(JSON.stringify(LaptopSVG));

  return svg;
});

const cvSVG = computed(() => {
  const svg = JSON.parse(JSON.stringify(CVSVG));

  return svg;
});

</script>

<style scoped>
.svg-container {
  display: flex;
  align-items: center;
}
</style>