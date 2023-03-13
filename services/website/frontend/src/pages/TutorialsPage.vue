<template>
  <div style="background-image: linear-gradient(90deg, #0c8635, #259b71, #69beaf);">
    <v-container class="bg-transparent pa-16 pt-12">
      <v-row class="d-flex flex-row">
        <v-col cols="12">
          <h2 class="text-h3 font-weight-bold text-center text-white pa-4">
            {{ t(i18nKeys.TutorialsPage.Title) }}
          </h2>
          <div class="presentation-trigger">
            <v-btn icon flat color="transparent" class="play-btn" width="100" height="100" title="Slide-Präsi"
                   @click="presentationOpened = true">
              <v-icon>mdi-play-circle-outline</v-icon>
            </v-btn>
          </div>
        </v-col>
      </v-row>
    </v-container>
  </div>
  <v-container>
    <p class="text-body-1 pb-3">
      <i18n-t :keypath="i18nKeys.TutorialsPage.ExampleDownload">
        <template #example>
          <a :href="`/example?lang=${i18nObject.locale.value}`" download>
            {{ t(i18nKeys.TutorialsPage.Example) }}
          </a>
        </template>
      </i18n-t>
    </p>
    <v-expansion-panels multiple v-model="opened">
      <v-expansion-panel>
        <v-expansion-panel-title>
          {{ t(i18nKeys.TutorialsPage.ThesisTemplate) }}
        </v-expansion-panel-title>
        <v-expansion-panel-text>
          <v-progress-circular
            indeterminate
            color="primary"
            size="150"
            style="width: 100%"
            v-if="thesisDocs === undefined"
          ></v-progress-circular>
          <MarkdownToVuetify :file="thesisDocs.Docs[0].Content" v-if="thesisDocs !== undefined" />
          <v-expansion-panels multiple v-if="thesisDocs !== undefined">
            <v-expansion-panel v-for="(e, i) in arrWithoutFirst(thesisDocs.Docs)">
              <v-expansion-panel-title>
                {{ e.Title }}
              </v-expansion-panel-title>
              <v-expansion-panel-text>
                <MarkdownToVuetify :file="e.Content" />
              </v-expansion-panel-text>
            </v-expansion-panel>
          </v-expansion-panels>
        </v-expansion-panel-text>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-title>
          {{ t(i18nKeys.TutorialsPage.ThesisTool) }}
        </v-expansion-panel-title>
        <v-expansion-panel-text>
          <MarkdownToVuetify :file="thesisToolDocs.Docs[0].Content" v-if="thesisToolDocs !== undefined" />
          <v-expansion-panels multiple v-if="thesisToolDocs !== undefined">
            <v-expansion-panel v-for="(e, i) in arrWithoutFirst(thesisToolDocs.Docs)">
              <v-expansion-panel-title>
                {{ e.Title }}
              </v-expansion-panel-title>
              <v-expansion-panel-text>
                <MarkdownToVuetify :file="e.Content" />
              </v-expansion-panel-text>
            </v-expansion-panel>
          </v-expansion-panels>
        </v-expansion-panel-text>
      </v-expansion-panel>
    </v-expansion-panels>
  </v-container>

  <v-dialog v-model="presentationOpened" width="600" height="400">

  </v-dialog>
</template>

<script lang="ts" setup>
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";
import {ref, watch} from "vue";
import {Documentation, GetThesisDocumentation, GetThesisToolDocumentation} from "../api/GetDocumentation";

// globals
const {t} = useI18n();
const i18nObject = useI18n();

// data
const props = defineProps({
  smallDisplay: Boolean
});
const opened = ref<number[]>([]);

const thesisDocs = ref<Documentation | undefined>(undefined);

const thesisToolDocs = ref<Documentation | undefined>(undefined);

const presentationOpened = ref(false);

// methods
function arrWithoutFirst(arr: any[]): any[] {
  const ac = Array.from(arr);
  ac.shift();
  return ac;
}

// watchers
watch(opened, async () => {
  if (opened.value.indexOf(0) >= 0 && !thesisDocs.value) {
    thesisDocs.value = await GetThesisDocumentation(i18nObject.locale.value);
  } else if (opened.value.indexOf(1) >= 0 && !thesisToolDocs.value) {
    thesisToolDocs.value = await GetThesisToolDocumentation(i18nObject.locale.value);
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