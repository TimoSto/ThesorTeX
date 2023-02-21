<template>
  <div style="background-image: linear-gradient(90deg, #0c8635, #259b71, #69beaf);">
    <v-container class="bg-transparent pa-16">
      <v-row class="d-flex flex-row">
        <v-col cols="12">
          <h2 class="text-h3 font-weight-bold text-center text-white pa-4">
            {{ t(i18nKeys.TutorialsPage.Title) }}
          </h2>
        </v-col>
      </v-row>
    </v-container>
  </div>
  <v-container>
    <v-expansion-panels multiple v-model="opened">
      <v-expansion-panel>
        <v-expansion-panel-title>
          {{ t(i18nKeys.TutorialsPage.ThesisTemplate) }}
        </v-expansion-panel-title>
        <v-expansion-panel-text>
          <v-progress-circular
            indeterminate
            color="primary"
            size="150"
            style="width: 100%"
            v-if="thesisDocs === undefined"
          ></v-progress-circular>
          <MarkdownToVuetify :file="thesisDocs.Main" v-if="thesisDocs !== undefined" />
          <v-expansion-panels multiple v-if="thesisDocs !== undefined">
            <v-expansion-panel>
              <v-expansion-panel-title>
                Wie kann ich die Nummerierung der Kapitel ändern?
              </v-expansion-panel-title>
              <v-expansion-panel-text>
                <MarkdownToVuetify :file="thesisDocs.ChapterNumbering" />
              </v-expansion-panel-text>
            </v-expansion-panel>
            <v-expansion-panel>
              <v-expansion-panel-title>
                Wie kann den Inhalt der Kopf- und Fußzeile ändern?
              </v-expansion-panel-title>
              <v-expansion-panel-text>
                <MarkdownToVuetify :file="thesisDocs.HeaderFooter" />
              </v-expansion-panel-text>
            </v-expansion-panel>
            <v-expansion-panel>
              <v-expansion-panel-title>
                Wie kann ich Abkürzungen im Abkürzungsverzeichnis auflisten?
              </v-expansion-panel-title>
              <v-expansion-panel-text>
                <MarkdownToVuetify :file="thesisDocs.Abbreviations" />
              </v-expansion-panel-text>
            </v-expansion-panel>
            <v-expansion-panel>
              <v-expansion-panel-title>
                Wie kann ich meine Anhänge strukturieren?
              </v-expansion-panel-title>
              <v-expansion-panel-text>
                <MarkdownToVuetify :file="thesisDocs.Appendix" />
              </v-expansion-panel-text>
            </v-expansion-panel>
          </v-expansion-panels>
        </v-expansion-panel-text>
      </v-expansion-panel>
      <!--      <v-expansion-panel>-->
      <!--        <v-expansion-panel-title>-->
      <!--          {{ t(i18nKeys.TutorialsPage.ThesisTemplateTech) }}-->
      <!--        </v-expansion-panel-title>-->
      <!--      </v-expansion-panel>-->
    </v-expansion-panels>
  </v-container>
</template>

<script lang="ts" setup>
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";
import {ref, watch} from "vue";
import {GetThesisDocumentation, ThesisDoc} from "../api/GetDocumentation";

// globals
const {t} = useI18n();
const i18nObject = useI18n();

// data
const props = defineProps({
  smallDisplay: Boolean
});
const opened = ref<number[]>([]);

const thesisDocs = ref<ThesisDoc | undefined>(undefined);

// watchers
watch(opened, async () => {
  if (opened.value.indexOf(0) >= 0 && !thesisDocs.value) {
    thesisDocs.value = await GetThesisDocumentation(i18nObject.locale.value);
  }
});

</script>

<style scoped lang="scss">
.fill-height-card {
  height: 100%;

  & .v-card-text {
    height: 100%;
  }
}

</style>