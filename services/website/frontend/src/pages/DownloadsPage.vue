<template>
  <div style="background-image: linear-gradient(90deg, #0c8635, #259b71, #69beaf);">
    <v-container class="bg-transparent pa-16 pt-12">
      <v-row class="d-flex flex-row">
        <v-col cols="4">
          <v-card class="fill-height-card">
            <v-card-text>
              <div class="d-flex flex-column" style="height: 100%">
              <span class="text-h5 text-center font-weight-bold"
                    style="display: inline-block; width: 100%;">{{ t(i18nKeys.StartPage.ThesisTemplateTitle) }}</span>
                <div class="mt-auto">
                  <SVGTemplate :svg="thesisPaths"
                               style="max-width: 100%; max-height: 250px; display: block; margin: 0 auto;" />
                  <v-btn variant="text" color="primary" style="width: 100%;" class="mb-2">
                    {{ t(i18nKeys.Common.LearnMore) }}
                  </v-btn>
                  <v-btn color="primary" style="width: 100%;">
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
                  <SVGTemplate :svg="laptopWithThesis"
                               style="max-width: 100%; max-height: 250px; display: block; margin: 0 auto;" />
                  <v-btn variant="text" color="primary" style="width: 100%;" class="mb-2">
                    {{ t(i18nKeys.Common.LearnMore) }}
                  </v-btn>
                  <v-btn color="primary" style="width: 100%;" @click="scrollToToolDownload">
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
                  <SVGTemplate :svg="cvSVG"
                               style="max-width: 100%; max-height: 250px; display: block; margin: 0 auto;" />
                  <v-btn variant="text" color="primary" style="width: 100%;" class="mb-2">
                    {{ t(i18nKeys.Common.LearnMore) }}
                  </v-btn>
                  <v-btn color="primary" style="width: 100%;">
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
    </v-container>
  </div>
  <div>
    <v-container class="bg-transparent pa-16 pt-4" ref="toolDownload">
      <h2 class="text-h5 font-weight-bold">{{ t(i18nKeys.StartPage.ThesisToolTitle) }}</h2>
      <p class="text-body-1">
        {{ t(i18nKeys.DownloadPage.InfoText) }}
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
                      style="display: inline-block; width: 100%;">MacOS (Intel)</span>
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
                      style="display: inline-block; width: 100%;">MacOS (Silicon)</span>
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
      <DownloadsTable :versions="versions" :per-os="true" :download-func="getToolDownloadLink" />
    </v-container>
  </div>
  <div>
    <v-container class="bg-transparent pa-16 pt-4" ref="toolDownload">
      <h2 class="text-h5 font-weight-bold">{{ t(i18nKeys.StartPage.ThesisTemplateTitle) }}</h2>
      <DownloadsTable :versions="versions" :download-func="getThesisTemplateDownloadLink" />
    </v-container>
  </div>
  <div>
    <v-container class="bg-transparent pa-16 pt-4" ref="toolDownload">
      <h2 class="text-h5 font-weight-bold">{{ t(i18nKeys.StartPage.CVTemplateTitle) }}</h2>
      <DownloadsTable :versions="versions" :download-func="getCVTemplateDownloadLink" />
    </v-container>
  </div>
</template>

<script lang="ts" setup>
import DownloadsTable from "../components/DownloadsTable.vue";
import WindowsIcon from "../components/WindowsIcon.vue";
import LinuxIcon from "../components/LinuxIcon.vue";
import MacIcon from "../components/MacIcon.vue";
import {computed, ref} from "vue";
import GetToolVersions, {VersionInfo} from "../api/GetToolVersions";
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

// globals

const {t} = useI18n();

// data
const props = defineProps({
  smallDisplay: Boolean
});

const toolDownload = ref(null);

const versions = ref([] as VersionInfo[]);

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

// methods
function scrollToToolDownload() {
  if (toolDownload.value) {
    toolDownload.value.$el.scrollIntoView({
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