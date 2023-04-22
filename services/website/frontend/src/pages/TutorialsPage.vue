<template>
  <FullHeightLayout :pages="jsonDocs ? 4 : 2" :white="true" :small-display="smallDisplay" ref="fullHeightLayout">
    <template #content-1>
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
              <v-icon>mdi-play-circle-outline</v-icon>
            </v-btn>
          </div>
        </v-col>
      </v-row>
    </template>
    <template #content-2 v-if="!jsonDocs">
      <v-progress-circular
        indeterminate
        color="primary"
        size="300"
        width="7"
        style="margin: 0 auto; display: block"
      ></v-progress-circular>
    </template>
    <template #content-2 v-if="jsonDocs">
      <h2 class="text-h3 font-weight-bold text-center pa-4">
        {{ jsonDocs?.ThesisTemplate.Title }}
      </h2>

      <p class="text-body-1 text-center pb-4"> {{ t(i18nKeys.TutorialsPage.TexKnowledge) }}</p>

      <v-expansion-panels multiple>
        <DocumentationPanel v-for="d in jsonDocs?.ThesisTemplate.Docs" :doc="d" />
      </v-expansion-panels>

    </template>
    <template #content-3 v-if="jsonDocs">
      <h2 class="text-h3 font-weight-bold text-center pa-4">
        {{ jsonDocs?.ThesisTool.Title }}
      </h2>

      <v-expansion-panels multiple>
        <DocumentationPanel v-for="d in jsonDocs?.ThesisTool.Docs" :doc="d" />
      </v-expansion-panels>

    </template>
    <template #content-4 v-if="jsonDocs">
      <h2 class="text-h3 font-weight-bold text-center pa-4">
        {{ jsonDocs?.CVTemplate.Title }}
      </h2>

      <v-expansion-panels multiple>
        <DocumentationPanel v-for="d in jsonDocs?.CVTemplate.Docs" :doc="d" />
      </v-expansion-panels>

    </template>
  </FullHeightLayout>
  <v-container>

  </v-container>

  <v-dialog v-model="presentationOpened" width="1000" height="700">
    <!--    <RevealJS :docs="thesisDocs.Docs" />-->
    <RevealJS :docs="[thesisTemplateDocsReveal, thesisToolDocsReveal, cvTemplateDocsReveal]" />
  </v-dialog>
</template>

<script lang="ts" setup>
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";
import {computed, onMounted, ref} from "vue";
import {GetDocumentationsJSON, ThesorTeXDocumentation} from "../api/GetDocumentation";
import FullHeightLayout from "../components/FullHeightLayout.vue";
import DocumentationPanel from "../components/DocumentationPanel.vue";
import {useRouter} from "vue-router";

// globals
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

const fullHeightLayout = ref(null);

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

// methods
function arrWithoutFirst(arr: any[]): any[] {
  const ac = Array.from(arr);
  ac.shift();
  return ac;
}

onMounted(async () => {
  jsonDocs.value = await GetDocumentationsJSON(i18nObject.locale.value);

  switch (router.currentRoute.value.query.target) {
    case "thesisTemplate":
      (fullHeightLayout.value! as typeof FullHeightLayout).jumpTo(2);
      break;
    case "thesisTool":
      (fullHeightLayout.value! as typeof FullHeightLayout).jumpTo(3);
      break;
    case "cvTemplate":
      (fullHeightLayout.value! as typeof FullHeightLayout).jumpTo(4);
      break;
  }
});

</script>

<style scoped lang="scss">
.fill-height-card {
  height: 100%;

  & .v-card-text {
    height: 100%;
  }
}

.presentation-trigger {
  width: 400px;
  height: 200px;
  background-color: rgba(var(--v-theme-background), 0.85);
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