<template>
  <FullHeightLayout :pages="4">
    <template #content-1="{ jumpTo }">
      <v-row class="d-flex flex-row">
        <v-col cols="4">
          <v-card class="fill-height-card">
            <v-card-text>
              <div class="d-flex flex-column" style="height: 100%">
              <span class="text-h5 text-center font-weight-bold"
                    style="display: inline-block; width: 100%;">{{ t(i18nKeys.StartPage.ThesisTemplateTitle) }}</span>
                <div class="mt-auto">
                  <v-btn variant="text" color="primary" style="width: 100%;" class="mb-2">
                    {{ t(i18nKeys.Common.LearnMore) }}
                  </v-btn>
                  <v-btn color="primary" style="width: 100%;" @click="jumpTo(2)">
                    <span style="white-space: normal;">
                      {{ t(i18nKeys.Common.Download) }}
                    </span>
                  </v-btn>
                </div>
              </div>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="4">
          <v-card class="fill-height-card">
            <v-card-text>
              <div class="d-flex flex-column" style="height: 100%">
                <span class="text-h5 text-center font-weight-bold"
                      style="display: inline-block; width: 100%;">{{ t(i18nKeys.StartPage.ThesisToolTitle) }}</span>
                <div class="mt-auto">
                  <v-btn variant="text" color="primary" style="width: 100%;" class="mb-2">
                    {{ t(i18nKeys.Common.LearnMore) }}
                  </v-btn>
                  <v-btn color="primary" style="width: 100%;" @click="jumpTo(3)">
                    <span style="white-space: normal;">
                      {{ t(i18nKeys.Common.Download) }}
                    </span>
                  </v-btn>
                </div>
              </div>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="4">
          <v-card class="fill-height-card">
            <v-card-text>
              <div class="d-flex flex-column" style="height: 100%">
                <span class="text-h5 text-center font-weight-bold"
                      style="display: inline-block; width: 100%;">{{ t(i18nKeys.StartPage.CVTemplateTitle) }}</span>
                <div class="mt-auto">
                  <v-btn variant="text" color="primary" style="width: 100%;" class="mb-2">
                    {{ t(i18nKeys.Common.LearnMore) }}
                  </v-btn>
                  <v-btn color="primary" style="width: 100%;" @click="jumpTo(4)">
                      <span style="white-space: normal;">
                        {{ t(i18nKeys.Common.Download) }}
                      </span>
                  </v-btn>
                </div>
              </div>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </template>
    <template #content-2>
      <v-row>
        <v-col>
          <h2 class="text-h5 font-weight-bold">{{ t(i18nKeys.StartPage.ThesisTemplateTitle) }}</h2>
          <p class="text-body-1 mb-4">
            {{ t(i18nKeys.DownloadPage.ThesisInfoText) }}
          </p>
          <DownloadsTable :versions="versions ? versions.ThesisTemplate : []"
                          :download-func="getThesisTemplateDownloadLink"
                          :max-width="500"
                          @open-notes="openedReleaseNotes = `thesisTemplate - ${$event}`" />
        </v-col>
      </v-row>
    </template>
    <template #content-3>
      <v-row>
        <v-col>
          <h2 class="text-h5 font-weight-bold">{{ t(i18nKeys.StartPage.ThesisToolTitle) }}</h2>
          <p class="text-body-1">
            {{ t(i18nKeys.DownloadPage.ToolInfoText) }}
          </p>
          <v-row class="pa-4 d-flex flex-row mt-1 mb-1">
            <v-col cols="3">
              <v-card elevation="6" style="height: 100%">
                <v-card-text style="height: 100%">
                  <div class="d-flex flex-column" style="height: 100%">
                          <span class="text-center font-weight-bold text-h6"
                                style="display: inline-block; width: 100%;">Windows</span>
                    <div class="mt-auto">
                      <WindowsIcon style="display: block; margin: 0 auto; height: 50px;" class="mb-4" />
                      <a :href="getToolDownloadLink('latest', 'windows')">
                        <v-btn color="primary" style="width: 100%;">
                          <v-icon>mdi-download</v-icon>
                        </v-btn>
                      </a>
                    </div>
                  </div>
                </v-card-text>
              </v-card>
            </v-col>
            <v-col cols="3">
              <v-card elevation="6" style="height: 100%">
                <v-card-text style="height: 100%">
                  <div class="d-flex flex-column" style="height: 100%">
                          <span class="text-center font-weight-bold text-h6"
                                style="display: inline-block; width: 100%;">Linux</span>
                    <div class="mt-auto">
                      <LinuxIcon style="display: block; margin: 0 auto; height: 50px; scale: 2" class="mb-4" />
                      <a :href="getToolDownloadLink('latest', 'linux')">
                        <v-btn color="primary" style="width: 100%;">
                          <v-icon>mdi-download</v-icon>
                        </v-btn>
                      </a>
                    </div>
                  </div>
                </v-card-text>
              </v-card>
            </v-col>
            <v-col cols="3">
              <v-card elevation="6" style="height: 100%">
                <v-card-text style="height: 100%">
                  <div class="d-flex flex-column" style="height: 100%">
                          <span class="text-center font-weight-bold text-h6"
                                style="display: inline-block; width: 100%;">MacOS (AMD)</span>
                    <div class="mt-auto">
                      <MacIcon style="display: block; margin: 0 auto;height: 50px;" class="mb-4" />
                      <a :href="getToolDownloadLink('latest', 'mac')">
                        <v-btn color="primary" style="width: 100%;">
                          <v-icon>mdi-download</v-icon>
                        </v-btn>
                      </a>
                    </div>
                  </div>
                </v-card-text>
              </v-card>
            </v-col>
            <v-col cols="3">
              <v-card elevation="6" style="height: 100%">
                <v-card-text style="height: 100%">
                  <div class="d-flex flex-column" style="height: 100%">
                          <span class="text-center font-weight-bold text-h6"
                                style="display: inline-block; width: 100%;">MacOS (ARM)</span>
                    <div class="mt-auto">
                      <MacIcon style="display: block; margin: 0 auto;height: 50px;" class="mb-4" />
                      <a :href="getToolDownloadLink('latest', 'mac_silicon')">
                        <v-btn color="primary" style="width: 100%;">
                          <v-icon>mdi-download</v-icon>
                        </v-btn>
                      </a>
                    </div>
                  </div>
                </v-card-text>
              </v-card>
            </v-col>
          </v-row>
          <DownloadsTable :versions="versions ? versions.Tool : []" :per-os="true"
                          :download-func="getToolDownloadLink"
                          :max-width="800" @open-notes="openedReleaseNotes = `thesisTool - ${$event}`" />
        </v-col>
      </v-row>
    </template>
    <template #content-4>
      <h2 class="text-h5 font-weight-bold">{{ t(i18nKeys.StartPage.CVTemplateTitle) }}</h2>
      <p class="text-body-1 mb-4">
        {{ t(i18nKeys.DownloadPage.CVInfoText) }}
      </p>
      <DownloadsTable :versions="versions ? versions.CvTemplate : []" :download-func="getCVTemplateDownloadLink"
                      :max-width="500" @open-notes="openedReleaseNotes = `cvTemplate - ${$event}`" />
    </template>
  </FullHeightLayout>
  <v-dialog v-model="releaseNotesOpen" width="600">
    <v-card>
      <v-card-title class="text-h5">
        {{ releaseNotesTitle }}
      </v-card-title>
      <v-card-text style="padding-top: 0;">
        <MarkdownToVuetify :file="releaseNotes" v-if="releaseNotes !== undefined" />
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script lang="ts" setup>
import DownloadsTable from "../components/DownloadsTable.vue";
import WindowsIcon from "../components/WindowsIcon.vue";
import LinuxIcon from "../components/LinuxIcon.vue";
import MacIcon from "../components/MacIcon.vue";
import {computed, ref, watch} from "vue";
import GetToolVersions, {VersionData} from "../api/GetToolVersions";
import {
  getCVTemplateDownloadLink,
  getThesisTemplateDownloadLink,
  getToolDownloadLink
} from "../api/GetToolDownloadLink";
import {ThesisSVG} from "../components/svgs/ThesisSVG";
import {LaptopSVG} from "../components/svgs/LaptopSVG";
import {CVSVG} from "../components/svgs/CVSVG";
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";
import GetReleaseNotes from "../api/GetReleaseNotes";
import FullHeightLayout from "../components/FullHeightLayout.vue";

// globals

const {t} = useI18n();

// data
const props = defineProps({
  smallDisplay: Boolean
});

const thesisDownload = ref(null);
const toolDownload = ref(null);
const cvDownload = ref(null);

const versions = ref<VersionData | undefined>(undefined);

const openedReleaseNotes = ref("");

const releaseNotes = ref<string | undefined>(undefined);

// computed
const releaseNotesOpen = computed({
  get(): boolean {
    return openedReleaseNotes.value !== "";
  },
  set(v: boolean) {
    return openedReleaseNotes.value = "";
  }
});

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

const releaseNotesTitle = computed(() => {
  const parts = openedReleaseNotes.value.split(" - ");
  let title = "";
  switch (parts[0]) {
    case "thesisTool":
      title = "Tool für Literaturmanagement";
      break;
    case "thesisTemplate":
      title = "Vorlage für wissenschaftliche Arbeiten";
      break;
    case "cvTemplate":
      title = "Vorlage für einen Lebenslauf";
      break;
  }
  title += " - " + parts[1];

  return title;
});

// watchers
watch(releaseNotesOpen, async () => {
  if (releaseNotesOpen.value) {
    releaseNotes.value = await GetReleaseNotes(openedReleaseNotes.value);
  }
});

// methods
function scrollToThesisDownload() {
  if (thesisDownload.value) {
    thesisDownload.value.$el.scrollIntoView({
      alignToTop: true,
      behavior: "smooth"
    });
  }
}

function scrollToToolDownload() {
  if (toolDownload.value) {
    toolDownload.value.$el.scrollIntoView({
      alignToTop: true,
      behavior: "smooth"
    });
  }
}

function scrollToCVDownload() {
  if (cvDownload.value) {
    cvDownload.value.$el.scrollIntoView({
      alignToTop: true,
      behavior: "smooth"
    });
  }
}

// onload
GetToolVersions().then(res => {
  versions.value = res;
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