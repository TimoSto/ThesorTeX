<template>
  <FullHeightLayout :pages="5" :small-display="smallDisplay">
    <template #content-1="{ jumpTo }">
      <v-row>
        <v-col :cols="smallDisplay ? 12: 6">
          <h2 class="text-h3 font-weight-bold text-white pa-4 pt-15">{{ t(i18nKeys.StartPage.Title) }}</h2>
          <p class="text-h5 text-white pa-4 pb-5">{{ t(i18nKeys.StartPage.Subtitle) }}</p>
          <v-list class="text-h6"
                  style="background-color: transparent; color: white; font-weight: bold; cursor: pointer;">
            <v-list-item v-ripple @click="jumpTo(2)">
              <template #prepend>
                <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
              </template>
              {{ t(i18nKeys.StartPage.ThesisTemplateTitle) }}
            </v-list-item>
            <v-list-item v-ripple @click="jumpTo(3)">
              <template #prepend>
                <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
              </template>
              {{ t(i18nKeys.StartPage.ThesisToolTitle) }}
            </v-list-item>
            <v-list-item v-ripple @click="jumpTo(4)">
              <template #prepend>
                <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
              </template>
              {{ t(i18nKeys.StartPage.CVTemplateTitle) }}
            </v-list-item>
          </v-list>
        </v-col>
        <v-col cols="6" style="position: relative" v-if="!smallDisplay">
          <SVGTemplate :svg="thesisPaths" class="smallIcon top" />
          <SVGTemplate :svg="laptopWithThesis" class="smallIcon left" />
          <SVGTemplate :svg="cvSVG" class="smallIcon right" />
        </v-col>
      </v-row>
    </template>
    <template #content-2>
      <v-row>
        <v-col v-if="!smallDisplay" cols="6" class="d-flex" style="justify-content: center; align-items: center;">
          <img :src="GermanTemplateImg" style="border: 1px solid black; max-height: 500px;" />
        </v-col>
        <v-col :cols="smallDisplay ? 12 : 6" class="d-flex" style="justify-content: center; align-items: center;">
          <div style="display: block">
            <h2 class="text-h3 font-weight-bold pt-6 pb-6">
              {{ t(i18nKeys.StartPage.ThesisTemplateTitle) }}
            </h2>
            <p class="text-h6 pb-6">
              {{ t(i18nKeys.StartPage.ThesisTemplateSubtitle) }}
            </p>
            <v-list class="text-h6"
                    style="background-color: transparent; font-weight: bold; cursor: pointer;">
              <v-list-item v-ripple to="/downloads">
                <template #prepend>
                  <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
                </template>
                {{ t(i18nKeys.Common.ToDownloads) }}
              </v-list-item>
              <v-list-item v-ripple href="/documentation?lang=de&format=zip" target="_blank">
                <template #prepend>
                  <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
                </template>
                {{ t(i18nKeys.StartPage.ThesisTemplateExample) }}
              </v-list-item>
              <v-list-item v-ripple to="/tutorials">
                <template #prepend>
                  <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
                </template>
                {{ t(i18nKeys.Common.Tutorial) }}
              </v-list-item>
            </v-list>
          </div>
        </v-col>
      </v-row>
    </template>
    <template #content-3>
      <v-row>
        <v-col :cols="smallDisplay ? 12 : 6">
          <h2 class="text-h3 font-weight-bold pt-6 pb-6">{{ t(i18nKeys.StartPage.ThesisToolTitle) }}</h2>
          <p class="text-h6 pb-6">
            {{ t(i18nKeys.StartPage.ThesisToolSubtitle) }}
          </p>
          <v-list class="text-h6"
                  style="background-color: transparent; font-weight: bold; cursor: pointer;">
            <v-list-item v-ripple to="/downloads">
              <template #prepend>
                <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
              </template>
              {{ t(i18nKeys.Common.ToDownloads) }}
            </v-list-item>
            <v-list-item v-ripple to="/tutorials">
              <template #prepend>
                <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
              </template>
              {{ t(i18nKeys.Common.Tutorial) }}
            </v-list-item>
          </v-list>
        </v-col>
        <v-col v-if="!smallDisplay" cols="6" class="d-flex"
               style="justify-content: center; align-items: center; position: relative;">
          <img :src="ThesisTool1Img"
               style="border: 1px solid black; max-width: 80%; transform: translate(75px, -50px); position: absolute;" />
          <img :src="ThesisTool2Img"
               style="border: 1px solid black; max-width: 80%; transform: translate(125px, 50px); position: absolute;" />
        </v-col>
      </v-row>
    </template>
    <template #content-4>
      <v-row>
        <v-col v-if="!smallDisplay" cols="6" class="d-flex" style="justify-content: center; align-items: center;">
          <img :src="CvTemplateImg" style="border: 1px solid black; max-height: 500px;" />
        </v-col>
        <v-col :cols="smallDisplay ? 12 : 6">
          <h2 class="text-h3 font-weight-bold pt-6 pb-6">{{ t(i18nKeys.StartPage.CVTemplateTitle) }}</h2>
          <p class="text-h6 pb-6">
            {{ t(i18nKeys.StartPage.ThesisTemplateSubtitle) }}
          </p>
          <v-list class="text-h6"
                  style="background-color: transparent; font-weight: bold; cursor: pointer;">
            <v-list-item v-ripple to="/downloads">
              <template #prepend>
                <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
              </template>
              {{ t(i18nKeys.Common.ToDownloads) }}
            </v-list-item>
            <v-list-item v-ripple to="/tutorials">
              <template #prepend>
                <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
              </template>
              {{ t(i18nKeys.Common.Tutorial) }}
            </v-list-item>
          </v-list>
        </v-col>
      </v-row>
    </template>
    <template #content-5>
      <v-row>
        <v-col :cols="smallDisplay ? 12 : 6">
          <h2 class="text-h3 font-weight-bold pt-6 pb-6">{{ t(i18nKeys.StartPage.KnownIssuesTitle) }}</h2>
          <p class="text-h6 pb-6">
            <i18n-t :keypath="i18nKeys.StartPage.KnownIssuesText">
              <template #here>
                <a href="https://github.com/TimoSto/ThesorTeX/labels/bug" target="_blank">{{
                    t(i18nKeys.Common.Here)
                  }}</a>
              </template>
              <template #bug>
                <span
                  style="color: red; border: 2px solid red; border-radius: 16px; padding: 0 8px; background-color: rgba(255, 0, 0, 0.25)">bug</span>
              </template>
            </i18n-t>
          </p>
        </v-col>
        <v-col v-if="!smallDisplay" cols="6" class="d-flex" style="justify-content: center; align-items: center;">
          
        </v-col>
      </v-row>
    </template>
  </FullHeightLayout>
</template>

<script lang="ts" setup>
import {computed} from "vue";
import {ThesisSVG} from "../components/svgs/ThesisSVG";
import {LaptopSVG} from "../components/svgs/LaptopSVG";
import {CVSVG} from "../components/svgs/CVSVG";
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";
import FullHeightLayout from "../components/FullHeightLayout.vue";
import GermanTemplateImg from "../assets/thesis_template_de.png";
import CvTemplateImg from "../assets/cv_template_de.png";
import ThesisTool1Img from "../assets/thesis_tool_1_de.png";
import ThesisTool2Img from "../assets/thesis_tool_2_de.png";

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