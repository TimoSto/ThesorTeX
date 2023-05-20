<script setup lang="ts">
import {i18nKeys} from "../i18n/keys";
import {useI18n} from "vue-i18n";
import {computed, ref} from "vue";
import {GetThesisSVG} from "../components/svgs/ThesisSVG";
import {GetLaptopSVG} from "../components/svgs/LaptopSVG";
import {GetCvSVG} from "../components/svgs/CVSVG";
import {GetBugSVG} from "../components/svgs/BugSVG";
import GermanTemplateImg from "../assets/thesis_template_de.png";
import ThesisToolImg from "../assets/thesis_tool_de.png";
import CvTemplateImg from "../assets/cv_template_de.png";
import {WaveContainer} from "@thesortex/vue-component-library/src/components";

const {t} = useI18n();

defineProps({
  smallDisplay: Boolean
});

const container0 = ref<InstanceType<typeof WaveContainer> | null>(null);
const container1 = ref<InstanceType<typeof WaveContainer> | null>(null);
const container2 = ref<InstanceType<typeof WaveContainer> | null>(null);
const container3 = ref<InstanceType<typeof WaveContainer> | null>(null);
const container4 = ref<InstanceType<typeof WaveContainer> | null>(null);
const container5 = ref<InstanceType<typeof WaveContainer> | null>(null);

// computed

const thesisSVG = computed(() => {
  //TODO: find a better way to loose reactivity
  const svg = GetThesisSVG("rgba(var(--v-theme-background))", "white");

  return svg;
});

const laptopWithThesisSVG = computed(() => {
  const svg = GetLaptopSVG("rgba(var(--v-theme-background))", "white");

  return svg;
});

const cvSVG = computed(() => {
  const svg = GetCvSVG("rgba(var(--v-theme-background))", "white");

  return svg;
});

const bugSVG = computed(() => {
  const svg = GetBugSVG("rgba(var(--v-theme-background))", "white");

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
    case 4:
      (container4.value!.$el as HTMLElement).scrollIntoView({behavior: "smooth"});
      break;
  }
}
</script>

<template>
  <WaveContainer :wave-func="1" :small-display="smallDisplay" ref="container0" bg-color="#001220" wave-color="#008833">
    <v-row>
      <v-col :cols="smallDisplay ? 12: 6">
        <h2 class="text-h3 font-weight-bold text-white pa-4 pt-15" ref="container0">{{
            t(i18nKeys.StartPage.Title)
          }}</h2>
        <p class="text-h5 text-white pa-4 pb-5">{{ t(i18nKeys.StartPage.Subtitle) }}</p>
        <v-list class="text-h6"
                style="background-color: transparent; color: white; font-weight: bold; cursor: pointer;">
          <v-list-item v-ripple @click="jumpTo(1)">
            <template #prepend>
              <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
            </template>
            {{ t(i18nKeys.StartPage.ThesisTemplateTitle) }}
          </v-list-item>
          <v-list-item v-ripple @click="jumpTo(2)">
            <template #prepend>
              <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
            </template>
            {{ t(i18nKeys.StartPage.ThesisToolTitle) }}
          </v-list-item>
          <v-list-item v-ripple @click="jumpTo(3)">
            <template #prepend>
              <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
            </template>
            {{ t(i18nKeys.StartPage.CVTemplateTitle) }}
          </v-list-item>
        </v-list>
      </v-col>
      <v-col cols="6" style="position: relative" v-if="!smallDisplay">
        <SVGTemplate :svg="thesisSVG" class="smallIcon top" />
        <SVGTemplate :svg="laptopWithThesisSVG" class="smallIcon left" />
        <SVGTemplate :svg="cvSVG" class="smallIcon right" />
      </v-col>
    </v-row>
  </WaveContainer>
  <WaveContainer :wave-func="2" :small-display="smallDisplay" ref="container1" bg-color="#008833" wave-color="#001220">
    <v-row>
      <v-col v-if="!smallDisplay" cols="6" class="d-flex" style="justify-content: center; align-items: center;">
        <ImageViewer :image="GermanTemplateImg" :image-title="t(i18nKeys.StartPage.CVTemplateTitle)" />
        <!--          <img :src="GermanTemplateImg" style="border: 1px solid black; max-height: 500px;" />-->
        <!--TODO: Update doc image-->
      </v-col>
      <v-col :cols="smallDisplay ? 12 : 6" class="d-flex" style="justify-content: center; align-items: center;">
        <div style="display: block">
          <h2 class="text-h3 text-white font-weight-bold pt-6 pb-6">
            {{ t(i18nKeys.StartPage.ThesisTemplateTitle) }}
          </h2>
          <p class="text-h6 pb-6 text-white">
            {{ t(i18nKeys.StartPage.ThesisTemplateSubtitle) }}
          </p>
          <v-list class="text-h6 text-white"
                  style="background-color: transparent; font-weight: bold; cursor: pointer;">
            <v-list-item v-ripple to="/downloads?target=thesisTemplate">
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
            <v-list-item v-ripple to="/tutorials?target=thesisTemplate">
              <template #prepend>
                <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
              </template>
              {{ t(i18nKeys.Common.Tutorial) }}
            </v-list-item>
          </v-list>
        </div>
      </v-col>
    </v-row>
  </WaveContainer>
  <WaveContainer :wave-func="3" :small-display="smallDisplay" ref="container2" bg-color="#001220" wave-color="#008833">
    <v-row>
      <v-col :cols="smallDisplay ? 12 : 6">
        <h2 class="text-h3 text-white font-weight-bold pt-6 pb-6">{{ t(i18nKeys.StartPage.ThesisToolTitle) }}</h2>
        <p class="text-h6 pb-6 text-white">
          {{ t(i18nKeys.StartPage.ThesisToolSubtitle) }}
        </p>
        <v-list class="text-h6 text-white"
                style="background-color: transparent; font-weight: bold; cursor: pointer;">
          <v-list-item v-ripple to="/downloads?target=thesisTool">
            <template #prepend>
              <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
            </template>
            {{ t(i18nKeys.Common.ToDownloads) }}
          </v-list-item>
          <v-list-item v-ripple to="/tutorials?target=thesisTool">
            <template #prepend>
              <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
            </template>
            {{ t(i18nKeys.Common.Tutorial) }}
          </v-list-item>
        </v-list>
      </v-col>
      <v-col v-if="!smallDisplay" cols="6" class="d-flex"
             style="justify-content: center; align-items: center; position: relative;">
        <ImageViewer :image="ThesisToolImg" :image-title="t(i18nKeys.StartPage.ThesisToolTitle)" />
      </v-col>
    </v-row>
  </WaveContainer>
  <WaveContainer :wave-func="2" :small-display="smallDisplay" ref="container3" bg-color="#008833" wave-color="#001220">
    <v-row>
      <v-col v-if="!smallDisplay" cols="6" class="d-flex" style="justify-content: center; align-items: center;">
        <ImageViewer :image="CvTemplateImg" :image-title="t(i18nKeys.StartPage.CVTemplateTitle)" />
      </v-col>
      <v-col :cols="smallDisplay ? 12 : 6">
        <h2 class="text-h3 text-white font-weight-bold pt-6 pb-6">{{ t(i18nKeys.StartPage.CVTemplateTitle) }}</h2>
        <p class="text-h6 text-white pb-6">
          {{ t(i18nKeys.StartPage.CVTemplateSubtitle) }}
        </p>
        <v-list class="text-h6 text-white"
                style="background-color: transparent; font-weight: bold; cursor: pointer;">
          <v-list-item v-ripple to="/downloads?target=cvTemplate">
            <template #prepend>
              <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
            </template>
            {{ t(i18nKeys.Common.ToDownloads) }}
          </v-list-item>
          <v-list-item v-ripple to="/tutorials?target=cvTemplate">
            <template #prepend>
              <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
            </template>
            {{ t(i18nKeys.Common.Tutorial) }}
          </v-list-item>
        </v-list>
      </v-col>
    </v-row>
  </WaveContainer>
  <WaveContainer :wave-func="-1" :small-display="smallDisplay" ref="container4" bg-color="#001220" wave-color="#008833">
    <v-row>
      <v-col :cols="smallDisplay ? 12 : 6">
        <h2 class="text-h3 text-white font-weight-bold pt-6 pb-6">{{ t(i18nKeys.StartPage.KnownIssuesTitle) }}</h2>
        <p class="text-h6 pb-6 text-white">
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
            <template #here2>
              <a href="https://github.com/TimoSto/ThesorTeX/labels/new%20feature" target="_blank">{{
                  t(i18nKeys.Common.Here)
                }}</a>
            </template>
            <template #feature>
                <span
                  style="color: green; border: 2px solid green; border-radius: 16px; padding: 0 8px; background-color: rgba(0, 255, 0, 0.25)">new feature</span>
            </template>
          </i18n-t>
        </p>
      </v-col>
      <v-col v-if="!smallDisplay" cols="6" class="d-flex" style="justify-content: center; align-items: center;">
        <SVGTemplate :svg="bugSVG" />
      </v-col>
    </v-row>
  </WaveContainer>
</template>

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