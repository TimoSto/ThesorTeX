<script setup lang="ts">
import {WaveContainer} from "@thesortex/vue-component-library/src/components";
import {i18nKeys} from "../i18n/keys";
import {useI18n} from "vue-i18n";
import {computed, onMounted, ref, watch} from "vue";
import {GetDocumentationsJSON, ThesorTeXDocumentation} from "../api/GetDocumentation";
import DocumentationPanel from "../components/DocumentationPanel.vue";
import {useRouter} from "vue-router";

const {t} = useI18n();
const i18nObject = useI18n();
const router = useRouter();

// data
const props = defineProps({
  smallDisplay: Boolean
});
const opened = ref<number[]>([]);

const jsonDocs = ref<ThesorTeXDocumentation | undefined>(undefined);

const presentationOpened = ref(false);

const container0 = ref<InstanceType<typeof WaveContainer> | null>(null);
const container1 = ref<InstanceType<typeof WaveContainer> | null>(null);
const container2 = ref<InstanceType<typeof WaveContainer> | null>(null);
const container3 = ref<InstanceType<typeof WaveContainer> | null>(null);

const expandedThesisTemplate = ref([] as number[]);
const expandedThesisTool = ref([] as number[]);
const expandedCVTemplate = ref([] as number[]);

// computed
const thesisTemplateDocsReveal = computed(() => {
  return {
    Title: jsonDocs.value?.ThesisTemplate.Title,
    Pages: jsonDocs.value?.ThesisTemplate.Docs
  };
});

const thesisToolDocsReveal = computed(() => {
  return {
    Title: jsonDocs.value?.ThesisTool.Title,
    Pages: jsonDocs.value?.ThesisTool.Docs
  };
});

const cvTemplateDocsReveal = computed(() => {
  return {
    Title: jsonDocs.value?.CVTemplate.Title,
    Pages: jsonDocs.value?.CVTemplate.Docs
  };
});

// watchers
watch(expandedThesisTemplate, () => {
  recalculateDimensions();
});
watch(expandedThesisTool, () => {
  recalculateDimensions();
});
watch(expandedCVTemplate, () => {
  recalculateDimensions();
});

// methods
function recalculateDimensions() {
  //TODO: find a better way
  const interval = setInterval(() => {
    container0.value.recalculateDimensions();
    container1.value.recalculateDimensions();
    container2.value.recalculateDimensions();
    container3.value.recalculateDimensions();
  }, 5);
  setTimeout(() => {
    clearInterval(interval);
  }, 500);
}

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

// onMounted
onMounted(async () => {
  jsonDocs.value = await GetDocumentationsJSON(i18nObject.locale.value);
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

        <v-expansion-panels multiple v-model="expandedThesisTemplate" theme="light">
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

        <v-expansion-panels multiple v-model="expandedThesisTool" theme="light">
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

        <v-expansion-panels multiple v-model="expandedCVTemplate" theme="light">
          <DocumentationPanel v-for="d in (jsonDocs as ThesorTeXDocumentation).CVTemplate.Docs" :doc="d" />
        </v-expansion-panels>
      </v-col>
    </v-row>
  </WaveContainer>
  <v-dialog v-model="presentationOpened" width="1000" height="700">
    <!--    <RevealJS :docs="thesisDocs.Docs" />-->
    <RevealJS :docs="[thesisTemplateDocsReveal, thesisToolDocsReveal, cvTemplateDocsReveal]" />
  </v-dialog>
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