<template>
  <div style="background-image: linear-gradient(90deg, #0c8635, #259b71, #69beaf);">
    <v-container class="bg-transparent pa-16">
      <v-row>
        <v-col cols="4">
          <v-card>
            <v-card-text>
              <span class="text-h5 text-center font-weight-bold"
                    style="display: inline-block; width: 100%;">Vorlage - Wissenschaftliche Arbeiten</span>
              <TemplateIcon :hide-icon="true" style="display: block; margin: 0 auto; transform: scale(0.9)" />
              <v-btn variant="text" color="primary" style="width: 100%;" class="mb-2">
                Weitere Informationen
              </v-btn>
              <a href="/templates/thesis" download>
                <v-btn color="primary" style="width: 100%;">
                <span style="white-space: normal;">
                  Herunterladen
                </span>
                </v-btn>
              </a>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="4">
          <v-card>
            <v-card-text>
              <span class="text-h5 text-center font-weight-bold"
                    style="display: inline-block; width: 100%;">Tool - Literaturmanagement</span>
              <TemplateIcon :hide-icon="false" style="display: block; margin: 0 auto; transform: scale(0.9)" />
              <v-btn variant="text" color="primary" style="width: 100%;" class="mb-2">
                Weitere Informationen
              </v-btn>
              <v-btn color="primary" style="width: 100%;" @click="scrollToToolDownload">
                <span style="white-space: normal;">
                  Herunterladen
                </span>
              </v-btn>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="4">
          <v-card>
            <v-card-text>
              <span class="text-h5 text-center font-weight-bold"
                    style="display: inline-block; width: 100%;">Vorlage - Lebenslauf</span>
              <CVIcon style="display: block; margin: 0 auto; transform: scale(0.7)" />
              <v-btn variant="text" color="primary" style="width: 100%;" class="mb-2">
                Weitere Informationen
              </v-btn>
              <a href="/templates/cv" download>
                <v-btn color="primary" style="width: 100%;">
                  <span style="white-space: normal;">
                    Herunterladen
                  </span>
                </v-btn>
              </a>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
  <div>
    <v-container class="bg-transparent pa-16 pt-4" ref="toolDownload">
      <h2 class="text-h5 font-weight-bold">Tool - Literaturmanagement</h2>
      <p class="text-body-1">Hier kannst du die aktuellst Version für dein Betriebssystem herunterladen. Unten kannst du
        auch eine der älteren Versionen herunterladen. Wenn du die heruntergeladene ZIP-Datei entpackst, kopiere die
        Datei darin an einen Ort deiner Wahl und starte sie durch einen Doppelklick.</p>
      <v-row class="pa-4">
        <v-col cols="3">
          <v-card elevation="6">
            <v-card-text>
              <span class="text-center font-weight-bold text-h6"
                    style="display: inline-block; width: 100%;">Windows</span>
              <WindowsIcon style="display: block; margin: 0 auto; height: 50px;" class="mb-4" />
              <a :href="getDownloadLink('latest', 'windows')">
                <v-btn color="primary" style="width: 100%;">
                  <v-icon>mdi-download</v-icon>
                </v-btn>
              </a>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="3">
          <v-card elevation="6">
            <v-card-text>
              <span class="text-center font-weight-bold text-h6"
                    style="display: inline-block; width: 100%;">Linux</span>
              <LinuxIcon style="display: block; margin: 0 auto; height: 50px; scale: 2" class="mb-4" />
              <a :href="getDownloadLink('latest', 'linux')">
                <v-btn color="primary" style="width: 100%;">
                  <v-icon>mdi-download</v-icon>
                </v-btn>
              </a>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="3">
          <v-card elevation="6">
            <v-card-text>
              <span class="text-center font-weight-bold text-h6"
                    style="display: inline-block; width: 100%;">MacOS (Intel)</span>
              <MacIcon style="display: block; margin: 0 auto;height: 50px;" class="mb-4" />
              <a :href="getDownloadLink('latest', 'mac')">
                <v-btn color="primary" style="width: 100%;">
                  <v-icon>mdi-download</v-icon>
                </v-btn>
              </a>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="3">
          <v-card elevation="6">
            <v-card-text>
              <span class="text-center font-weight-bold text-h6"
                    style="display: inline-block; width: 100%;">MacOS (Silicon)</span>
              <MacIcon style="display: block; margin: 0 auto; height: 50px;" class="mb-4" />
              <a :href="getDownloadLink('latest', 'mac_silicon')">
                <v-btn color="primary" style="width: 100%;">
                  <v-icon>mdi-download</v-icon>
                </v-btn>
              </a>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
      <DownloadsTable :versions="versions" />
    </v-container>
  </div>
</template>

<script lang="ts" setup>
import TemplateIcon from "../components/TemplateIcon.vue";
import CVIcon from "../components/CVIcon.vue";
import DownloadsTable from "../components/DownloadsTable.vue";
import WindowsIcon from "../components/WindowsIcon.vue";
import LinuxIcon from "../components/LinuxIcon.vue";
import MacIcon from "../components/MacIcon.vue";
import {ref} from "vue";
import GetToolVersions, {VersionInfo} from "../api/GetToolVersions";
import getDownloadLink from "../api/GetDownloadLink";

// data
const toolDownload = ref(null);

const versions = ref([] as VersionInfo[]);

// methods
function scrollToToolDownload() {
  if (toolDownload.value) {
    toolDownload.value.$el.scrollIntoView(true);
  }
}

// onload
GetToolVersions().then(res => {
  versions.value = res;
  console.log(res);
});

</script>

<style scoped>

</style>