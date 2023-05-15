<script setup lang="ts">
import {WaveContainer} from "@thesortex/vue-component-library/src/components";
import DownloadsTable from "../components/DownloadsTable.vue";
import {i18nKeys} from "../i18n/keys";
import {computed, onMounted, ref} from "vue";
import {GetThesisSVG} from "../components/svgs/ThesisSVG";
import {GetLaptopSVG} from "../components/svgs/LaptopSVG";
import {GetCvSVG} from "../components/svgs/CVSVG";
import {useI18n} from "vue-i18n";
import {
  getCVTemplateDownloadLink,
  getThesisTemplateDownloadLink,
  getToolDownloadLink
} from "../api/GetToolDownloadLink";
import GetToolVersions, {VersionData} from "../api/GetToolVersions";
import LinuxIcon from "../components/LinuxIcon.vue";
import MacIcon from "../components/MacIcon.vue";
import WindowsIcon from "../components/WindowsIcon.vue";
import {useRouter} from "vue-router";

defineProps({
  smallDisplay: Boolean
});

const {t} = useI18n();

const router = useRouter();

const container0 = ref<InstanceType<typeof WaveContainer> | null>(null);
const container1 = ref<InstanceType<typeof WaveContainer> | null>(null);
const container2 = ref<InstanceType<typeof WaveContainer> | null>(null);
const container3 = ref<InstanceType<typeof WaveContainer> | null>(null);

// data
const versions = ref<VersionData | undefined>(undefined);

const openedReleaseNotes = ref("");

// computed
const thesisSVG = computed(() => {
  //TODO: find a better way to loose reactivity
  const svg = GetThesisSVG("rgba(var(--v-theme-on-background))", "rgba(var(--v-theme-background))");

  return svg;
});

const laptopWithThesisSVG = computed(() => {
  const svg = GetLaptopSVG("rgba(var(--v-theme-on-background))", "rgba(var(--v-theme-background))");

  return svg;
});

const cvSVG = computed(() => {
  const svg = GetCvSVG("rgba(var(--v-theme-on-background))", "rgba(var(--v-theme-background))");

  return svg;
});

// methods
function jumpTo(n: number) {
  switch (n) {
    case 0:
      (container0.value!.$el as HTMLElement).scrollIntoView({behavior: "smooth"});
      break;
    case 1:
      (container1.value!.$el as HTMLElement).scrollIntoView({behavior: "smooth"});
      break;
    case 2:
      (container2.value!.$el as HTMLElement).scrollIntoView({behavior: "smooth"});
      break;
    case 3:
      (container3.value!.$el as HTMLElement).scrollIntoView({behavior: "smooth"});
      break;
  }
}

// onmounted
GetToolVersions().then(res => {
  versions.value = res;
});
onMounted(() => {
  setTimeout(() => {
    switch (router.currentRoute.value.query.target) {
      case "thesisTemplate":
        jumpTo(1);
        break;
      case "thesisTool":
        jumpTo(2);
        break;
      case "cvTemplate":
        jumpTo(3);
        break;
    }
  });
});
</script>

<template>
  <WaveContainer :wave-func="3" :small-display="smallDisplay" ref="container0" bg-color="#008833" wave-color="#001220">
    <v-row class="d-flex flex-row">
      <v-col cols="4">
        <v-card class="fill-height-card">
          <v-card-text>
            <div class="d-flex flex-column" style="height: 100%">
              <span class="text-h5 text-center font-weight-bold"
                    style="display: inline-block; width: 100%;">{{ t(i18nKeys.StartPage.ThesisTemplateTitle) }}</span>
              <div class="mt-auto">
                <SVGTemplate :svg="thesisSVG" max-height="200px" max-width="100%" />
              </div>
              <div class="mt-auto">
                <v-btn variant="text" color="primary" style="width: 100%;" class="mb-2"
                       to="/tutorials?target=thesisTemplate">
                  {{ t(i18nKeys.Common.LearnMore) }}
                </v-btn>
                <v-btn color="primary" style="width: 100%;" @click="jumpTo(1)">
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
                      style="display: inline-block; width: 100%;">{{
                    t(i18nKeys.StartPage.ThesisToolTitle, {hyphen: "&shy;"})
                  }}</span>
              <div class="mt-auto">
                <SVGTemplate :svg="laptopWithThesisSVG" max-height="200px" max-width="100%" />
              </div>
              <div class="mt-auto">
                <v-btn variant="text" color="primary" style="width: 100%;" class="mb-2"
                       to="/tutorials?target=thesisTool">
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
                      style="display: inline-block; width: 100%;">{{ t(i18nKeys.StartPage.CVTemplateTitle) }}</span>
              <div class="mt-auto">
                <SVGTemplate :svg="cvSVG" max-height="200px" max-width="100%" />
              </div>
              <div class="mt-auto">
                <v-btn variant="text" color="primary" style="width: 100%;" class="mb-2"
                       to="/tutorials?target=cvTemplate">
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
    </v-row>
  </WaveContainer>

  <WaveContainer :wave-func="1" :small-display="smallDisplay" ref="container1" bg-color="#001220" wave-color="#008833">
    <v-row>
      <v-col>
        <h2 class="text-h3 font-weight-bold">{{ t(i18nKeys.StartPage.ThesisTemplateTitle) }}</h2>
        <p class="text-body-1">
          {{ t(i18nKeys.DownloadPage.ThesisInfoText) }}
        </p>
        <v-list class="text-h6"
                style="background-color: transparent; font-weight: bold; cursor: pointer;">
          <v-list-item v-ripple to="/tutorials?target=thesisTemplate">
            <template #prepend>
              <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
            </template>
            {{ t(i18nKeys.Common.LearnMore) }}
          </v-list-item>
        </v-list>
        <DownloadsTable :versions="versions ? (versions as VersionData).ThesisTemplate : []"
                        :download-func="getThesisTemplateDownloadLink"
                        :max-width="500"
                        @open-notes="openedReleaseNotes = `thesisTemplate - ${$event}`" />
      </v-col>
    </v-row>
  </WaveContainer>

  <WaveContainer :wave-func="2" :small-display="smallDisplay" ref="container2" bg-color="#008833" wave-color="#001220">
    <v-row>
      <v-col>
        <h2 class="text-h3 font-weight-bold">{{ t(i18nKeys.StartPage.ThesisToolTitle) }}</h2>
        <p class="text-body-1">
          {{ t(i18nKeys.DownloadPage.ToolInfoText) }}
        </p>
        <v-list class="text-h6 text-white"
                style="background-color: transparent; font-weight: bold; cursor: pointer;">
          <v-list-item v-ripple to="/tutorials?target=thesisTool">
            <template #prepend>
              <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
            </template>
            {{ t(i18nKeys.Common.LearnMore) }}
          </v-list-item>
        </v-list>
        <v-row class="pa-4 d-flex flex-row mb-1">
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
                    <a :href="getToolDownloadLink('latest', 'mac_arm')">
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
        <DownloadsTable :versions="versions ? (versions as VersionData).Tool : []" :per-os="true"
                        :download-func="getToolDownloadLink"
                        :max-width="800" @open-notes="openedReleaseNotes = `thesisTool - ${$event}`" />
      </v-col>
    </v-row>
  </WaveContainer>

  <WaveContainer :wave-func="-1" :small-display="smallDisplay" ref="container3" bg-color="#001220" wave-color="#008833">
    <v-row>
      <v-col>
        <h2 class="text-h3 font-weight-bold">{{ t(i18nKeys.StartPage.CVTemplateTitle) }}</h2>
        <p class="text-body-1">
          {{ t(i18nKeys.DownloadPage.CVInfoText) }}
        </p>
        <v-list class="text-h6"
                style="background-color: transparent; font-weight: bold; cursor: pointer;">
          <v-list-item v-ripple to="/tutorials?target=cvTemplate">
            <template #prepend>
              <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
            </template>
            {{ t(i18nKeys.Common.LearnMore) }}
          </v-list-item>
        </v-list>
        <DownloadsTable :versions="versions ? (versions as VersionData).CvTemplate : []"
                        :download-func="getCVTemplateDownloadLink"
                        :max-width="500" @open-notes="openedReleaseNotes = `cvTemplate - ${$event}`" />
      </v-col>
    </v-row>
  </WaveContainer>
</template>

<style scoped lang="scss">

</style>