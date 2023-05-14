<script setup lang="ts">
import {i18nKeys} from "../i18n/keys";
import {useI18n} from "vue-i18n";
import {computed} from "vue";
import {ThesisSVG} from "../components/svgs/ThesisSVG";
import {LaptopSVG} from "../components/svgs/LaptopSVG";
import {CVSVG} from "../components/svgs/CVSVG";
import {BugSVG} from "../components/svgs/BugSVG";

const {t} = useI18n();

defineProps({
  smallDisplay: Boolean
});

// computed

const thesisSVG = computed(() => {
  //TODO: find a better way to loose reactivity
  const svg = JSON.parse(JSON.stringify(ThesisSVG));

  return svg;
});

const laptopWithThesisSVG = computed(() => {
  const svg = JSON.parse(JSON.stringify(LaptopSVG));

  return svg;
});

const cvSVG = computed(() => {
  const svg = JSON.parse(JSON.stringify(CVSVG));

  return svg;
});

const bugSVG = computed(() => {
  const svg = JSON.parse(JSON.stringify(BugSVG));

  return svg;
});
</script>

<template>
  <WaveContainer :wave-func="1" bg-color="#001220" wave-color="#008833">
    <v-row>
      <v-col :cols="smallDisplay ? 12: 6">
        <h2 class="text-h3 font-weight-bold text-white pa-4 pt-15">{{ t(i18nKeys.StartPage.Title) }}</h2>
        <p class="text-h5 text-white pa-4 pb-5">{{ t(i18nKeys.StartPage.Subtitle) }}</p>
        <v-list class="text-h6"
                style="background-color: transparent; color: white; font-weight: bold; cursor: pointer;">
          <v-list-item v-ripple>
            <template #prepend>
              <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
            </template>
            {{ t(i18nKeys.StartPage.ThesisTemplateTitle) }}
          </v-list-item>
          <v-list-item v-ripple>
            <template #prepend>
              <v-icon style="opacity: 1">mdi-arrow-right-circle-outline</v-icon>
            </template>
            {{ t(i18nKeys.StartPage.ThesisToolTitle) }}
          </v-list-item>
          <v-list-item v-ripple>
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
  <WaveContainer :wave-func="2" bg-color="#008833" wave-color="#001220">
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