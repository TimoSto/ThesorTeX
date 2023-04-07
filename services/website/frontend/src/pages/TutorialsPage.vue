<template>
  <FullHeightLayout :pages="3" :white="true">
    <template #content-1>
      <v-row class="d-flex flex-row">
        <v-col cols="12">
          <h2 class="text-h3 font-weight-bold text-center text-white pa-4">
            {{ t(i18nKeys.TutorialsPage.Title) }}
          </h2>

          <p class="text-h5 text-center text-white mb-4">
            <i18n-t :keypath="i18nKeys.TutorialsPage.SubTitle">
              <template #pdf>
                <a :href="`/example?lang=${i18nObject.locale.value}&format=pdf`" style="color: white" download>PDF</a>
              </template>
              <template #example>
                <a :href="`/example?lang=${i18nObject.locale.value}`" style="color: white"
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
    <template #content-2>
      <h2 class="text-h3 font-weight-bold text-center pa-4">
        {{ t(i18nKeys.TutorialsPage.ThesisTemplate) }}
      </h2>

      <p class="text-body-1 text-center pb-4"> {{ t(i18nKeys.TutorialsPage.TexKnowledge) }}</p>

      <v-expansion-panels multiple>
        <DocumentationPanel v-for="d in thesisTemplateDocs?.Docs" :doc="d" />
      </v-expansion-panels>

    </template>
    <template #content-3>
      <h2 class="text-h3 font-weight-bold text-center pa-4">
        {{ t(i18nKeys.TutorialsPage.ThesisTool) }}
      </h2>

    </template>
  </FullHeightLayout>
  <v-container>

  </v-container>

  <v-dialog v-model="presentationOpened" width="1000" height="700">
    <!--    <RevealJS :docs="thesisDocs.Docs" />-->
    <RevealJS :docs="[thesisTemplateDocsReveal]" />
  </v-dialog>
</template>

<script lang="ts" setup>
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";
import {computed, onMounted, ref} from "vue";
import {Documentation, DocumentationPack, GetThesisTemplateDocumentation} from "../api/GetDocumentation";
import FullHeightLayout from "../components/FullHeightLayout.vue";
import DocumentationPanel from "../components/DocumentationPanel.vue";

// globals
const {t} = useI18n();
const i18nObject = useI18n();

// data
const props = defineProps({
  smallDisplay: Boolean
});
const opened = ref<number[]>([]);

const thesisTemplateDocs = ref<DocumentationPack | undefined>(undefined);

const thesisToolDocs = ref<Documentation | undefined>(undefined);

const presentationOpened = ref(false);

// computed
const thesisTemplateDocsReveal = computed(() => {
  return {
    Title: t(i18nKeys.TutorialsPage.ThesisTemplate),
    Pages: thesisTemplateDocs.value
  };
});

// methods
function arrWithoutFirst(arr: any[]): any[] {
  const ac = Array.from(arr);
  ac.shift();
  return ac;
}

onMounted(async () => {
  thesisTemplateDocs.value = await GetThesisTemplateDocumentation(i18nObject.locale.value);
  console.log(thesisTemplateDocs.value);
  // thesisToolDocs.value = await GetThesisToolDocumentation(i18nObject.locale.value);

  console.log(thesisTemplateDocsReveal.value, thesisTemplateDocs.value);
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