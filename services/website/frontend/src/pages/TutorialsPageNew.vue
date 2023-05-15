<script setup lang="ts">
import {WaveContainer} from "@thesortex/vue-component-library/src/components";
import {i18nKeys} from "../i18n/keys";
import {useI18n} from "vue-i18n";
import {onMounted, ref} from "vue";
import {GetDocumentationsJSON, ThesorTeXDocumentation} from "../api/GetDocumentation";
import DocumentationPanel from "../components/DocumentationPanel.vue";

const {t} = useI18n();
const i18nObject = useI18n();

// data
const props = defineProps({
  smallDisplay: Boolean
});
const opened = ref<number[]>([]);

const jsonDocs = ref<ThesorTeXDocumentation | undefined>(undefined);

const presentationOpened = ref(false);

// onMounted
onMounted(async () => {
  jsonDocs.value = await GetDocumentationsJSON(i18nObject.locale.value);
});
</script>

<template>
  <WaveContainer :wave-func="2" :small-display="smallDisplay" ref="container0" bg-color="#001220" wave-color="#008833">
    <v-row class="d-flex flex-row">
      <v-col cols="12">
        <h2 class="text-h3 font-weight-bold text-center text-white pa-4">
          {{ t(i18nKeys.TutorialsPage.Title) }}
        </h2>

        <p class="text-h5 text-center text-white mb-4">
          <i18n-t :keypath="i18nKeys.TutorialsPage.SubTitle">
            <template #pdf>
              <a :href="`/documentation?lang=${i18nObject.locale.value}&format=pdf`" style="color: white"
                 download="documentation.pdf">PDF</a>
            </template>
            <template #example>
              <a :href="`/documentation?lang=${i18nObject.locale.value}&format=zip`" style="color: white"
                 download>{{ t(i18nKeys.TutorialsPage.ExampleProject) }}</a>
            </template>
          </i18n-t>
        </p>
        <div class="presentation-trigger">
          <v-btn icon flat color="transparent" class="play-btn" width="100" height="100" title="Slide-Präsi"
                 @click="presentationOpened = true">
            <v-icon color="black">mdi-play-circle-outline</v-icon>
          </v-btn>
        </div>
      </v-col>
    </v-row>
  </WaveContainer>
  <WaveContainer :wave-func="1" :small-display="smallDisplay" ref="container1" bg-color="#008833"
                 wave-color="#001220">
    <v-progress-circular
      indeterminate
      color="#008833"
      size="300"
      width="7"
      style="margin: 0 auto; display: block"
      v-if="!jsonDocs"
    ></v-progress-circular>
    <v-row v-if="jsonDocs">
      <v-col>
        <h2 class="text-h3 font-weight-bold text-center pa-4">
          {{ (jsonDocs as ThesorTeXDocumentation).ThesisTemplate.Title }}
        </h2>

        <p class="text-body-1 text-center pb-4"> {{ t(i18nKeys.TutorialsPage.TexKnowledge) }}</p>

        <v-expansion-panels multiple>
          <DocumentationPanel v-for="d in (jsonDocs as ThesorTeXDocumentation).ThesisTemplate.Docs" :doc="d" />
        </v-expansion-panels>
      </v-col>
    </v-row>
  </WaveContainer>
  <WaveContainer :wave-func="2" :small-display="smallDisplay" ref="container2" bg-color="#001220"
                 wave-color="#008833">
    <v-progress-circular
      indeterminate
      color="#008833"
      size="300"
      width="7"
      style="margin: 0 auto; display: block"
      v-if="!jsonDocs"
    ></v-progress-circular>
    <v-row v-if="jsonDocs">
      <v-col>
        <h2 class="text-h3 font-weight-bold text-center pa-4">
          {{ (jsonDocs as ThesorTeXDocumentation).ThesisTool.Title }}
        </h2>

        <v-expansion-panels multiple>
          <DocumentationPanel v-for="d in (jsonDocs as ThesorTeXDocumentation).ThesisTool.Docs" :doc="d" />
        </v-expansion-panels>
      </v-col>
    </v-row>
  </WaveContainer>
  <WaveContainer :wave-func="3" :small-display="smallDisplay" ref="container3" bg-color="#008833"
                 wave-color="#001220">
    <v-progress-circular
      indeterminate
      color="#008833"
      size="300"
      width="7"
      style="margin: 0 auto; display: block"
      v-if="!jsonDocs"
    ></v-progress-circular>
    <v-row v-if="jsonDocs">
      <v-col>
        <h2 class="text-h3 font-weight-bold text-center pa-4">
          {{ (jsonDocs as ThesorTeXDocumentation).CVTemplate.Title }}
        </h2>

        <v-expansion-panels multiple>
          <DocumentationPanel v-for="d in (jsonDocs as ThesorTeXDocumentation).CVTemplate.Docs" :doc="d" />
        </v-expansion-panels>
      </v-col>
    </v-row>
  </WaveContainer>
</template>

<style scoped lang="scss">
@import "../styles/fillHeightCard";

.presentation-trigger {
  width: 400px;
  height: 200px;
  background-color: rgba(255, 255, 255, 0.85);
  border-radius: 8px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  justify-content: center;

  & .play-btn {
    font-size: 65px;
  }
}
</style>