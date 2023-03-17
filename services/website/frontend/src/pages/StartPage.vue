<template>
  <FullHeightContainer bg="gradient" :first="true">
    <v-row>
      <v-col cols="6">
        <h2 class="text-h3 font-weight-bold text-white pa-4 pt-15">{{ t(i18nKeys.StartPage.Title) }}</h2>
        <p class="text-h5 text-white pa-4 pb-15">{{ t(i18nKeys.StartPage.Subtitle) }}</p>
        <v-list class="text-body-1"
                style="background-color: transparent; color: white; font-weight: bold; cursor: pointer;">
          <v-list-item v-ripple>
            <template #prepend>
              <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
            </template>
            Vorlage für eine Haus-/Abschlussarbeit
          </v-list-item>
          <v-list-item v-ripple>
            <template #prepend>
              <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
            </template>
            Tool zum Literaturmanagement
          </v-list-item>
          <v-list-item v-ripple>
            <template #prepend>
              <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
            </template>
            Vorlage für einen Lebenslauf
          </v-list-item>
        </v-list>
      </v-col>
      <v-col cols="6" style="position: relative">
        <SVGTemplate :svg="thesisPaths" class="smallIcon top" />
        <SVGTemplate :svg="laptopWithThesis" class="smallIcon left" />
        <SVGTemplate :svg="cvSVG" class="smallIcon right" />
      </v-col>
    </v-row>
  </FullHeightContainer>
  <FullHeightContainer bg="white">
    <v-row>
      <v-col :cols="smallDisplay ? 12 : 6">
        <h2 class="text-h3 font-weight-bold pt-6 pb-6">{{ t(i18nKeys.StartPage.ThesisTemplateTitle) }}</h2>
        <p class="text-h6 pb-6">
          <i18n-t :keypath="i18nKeys.StartPage.ThesisTemplateSubtitle" scope="global">
            <template #link>
              <a :href="getThesisTemplateDownloadLink('latest')" download>
                {{ t(i18nKeys.StartPage.ThesisTemplateLink) }}
              </a>
            </template>
          </i18n-t>
        </p>
        <v-btn color="primary" to="/tutorials">
          {{ t(i18nKeys.Common.LearnMore) }}
        </v-btn>
      </v-col>
      <v-col v-if="!smallDisplay" cols="6" class="svg-container">
        <SVGTemplate :svg="thesisPaths" style="max-height: 300px; display: block; margin: 0 auto;" />
      </v-col>
    </v-row>
  </FullHeightContainer>
  <FullHeightContainer>
    <v-row>
      <v-col v-if="!smallDisplay" cols="6" class="svg-container">
        <SVGTemplate :svg="laptopWithThesis" style="max-height: 300px; display: block; margin: 0 auto;" />
      </v-col>
      <v-col :cols="smallDisplay ? 12 : 6">
        <h2 class="text-h3 font-weight-bold pt-6 pb-6">{{ t(i18nKeys.StartPage.ThesisToolTitle) }}</h2>
        <p class="text-h6 pb-6">
          {{ t(i18nKeys.StartPage.ThesisToolSubtitle) }}
        </p>
        <v-btn color="primary" to="/tutorials">
          {{ t(i18nKeys.Common.LearnMore) }}
        </v-btn>
      </v-col>
    </v-row>
  </FullHeightContainer>
  <FullHeightContainer>
    <v-row>
      <v-col :cols="smallDisplay ? 12 : 6">
        <h2 class="text-h3 font-weight-bold pt-6 pb-6">{{ t(i18nKeys.StartPage.CVTemplateTitle) }}</h2>
        <p class="text-h6 pb-6">
          <i18n-t :keypath="i18nKeys.StartPage.CVTemplateSubtitle" scope="global">
            <template #link>
              <a :href="getCVTemplateDownloadLink('latest')" download>
                {{ t(i18nKeys.StartPage.CVTemplateLink) }}
              </a>
            </template>
          </i18n-t>
        </p>
        <v-btn color="primary" to="/tutorials">
          {{ t(i18nKeys.Common.LearnMore) }}
        </v-btn>
      </v-col>
      <v-col v-if="!smallDisplay" cols="6" class="svg-container">
        <SVGTemplate :svg="cvSVG" style="max-height: 300px; display: block; margin: 0 auto;" />
      </v-col>
    </v-row>
  </FullHeightContainer>
</template>

<script lang="ts" setup>
import {computed} from "vue";
import {ThesisSVG} from "../components/svgs/ThesisSVG";
import {LaptopSVG} from "../components/svgs/LaptopSVG";
import {CVSVG} from "../components/svgs/CVSVG";
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";
import {getCVTemplateDownloadLink, getThesisTemplateDownloadLink} from "../api/GetToolDownloadLink";
import FullHeightContainer from "../components/FullHeightContainer.vue";

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

<style scoped lang="scss">
.svg-container {
  display: flex;
  align-items: center;
}

.smallIcon {
  max-height: 200px;
  max-width: 200px;
  position: absolute;
  left: 50%;
  top: 50%;

  &.top {
    transform: translate(-50%, calc(-50% - 90px));
  }

  &.left {
    transform: translate(calc(-50% - 120px), calc(-50% + 100px));
  }

  &.right {
    transform: translate(calc(-50% + 120px), calc(-50% + 100px));
  }
}
</style>