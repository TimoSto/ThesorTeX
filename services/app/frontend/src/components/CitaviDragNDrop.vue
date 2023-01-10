<template>
  <div
    v-ripple
    class="drag_n_drop--container text-center"
    @drop.prevent="handleDrop"
    @dragover.prevent
    @click="handleClick"
  >
    <input
      ref="fileupload"
      type="file"
      style="display:none"
      @change="handleSelect"
    >
    {{ t(i18nKeys.Project.Import) }}
  </div>
  <v-dialog
    v-model="dialogOpen"
    width="450"
  >
    <v-card>
      <v-card-title>Einträge hochladen</v-card-title>
      <v-card-text>
        Folgende Einträge werden hochgeladen und werden bestehende Einträge überschreiben:
        <v-list lines="two">
          <v-list-item
            v-for="(e,i) in result.Entries"
            :key="`e-${e.Key}`"
            :title="e.Key"
            :subtitle="e.Category"
          >
            <template #append>
              <v-checkbox-btn
                v-model="uploaded[i]"
                color="primary"
              />
            </template>
          </v-list-item>
        </v-list>
        <div v-if="result.Unknowns.length > 0">
          Folgende Einträge konnten keiner Kategorie zugeordnet werden:
          <v-list lines="two">
            <v-list-item
              v-for="(e) in result.Unknowns"
              :key="`e-${e.Key}`"
              :title="e.Key"
              :subtitle="e.Category"
            />
          </v-list>
        </div>
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn
          color="primary"
          @click="dialogOpen=false"
        >
          Abbrechen
        </v-btn>
        <v-btn
          color="primary"
          :disabled="!uploadPossible"
        >
          Hochladen
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";
import {computed, ref} from "vue";
import AnalyseBibFile, {BibAnalyticsResult} from "../bibanalytics/BibAnalytics";
import {useProjectDataStore} from "../stores/projectDataStore";

// globals
const {t} = useI18n();

const projectDataStore = useProjectDataStore();

// data
const fileupload = ref(null);

const triggered = ref(false);

const result = ref({} as BibAnalyticsResult);

const uploaded = ref([] as boolean[]);

// computed
const dialogOpen = computed({
  get(): boolean {
    return triggered.value
  },
  set(v: boolean) {
    triggered.value = v;
  }
})

const uploadPossible = computed(() => {
  return uploaded.value.filter(e => e === true).length > 0;
})

// methods
function handleDrop(evt: DragEvent) {
  evt.preventDefault();
  if( evt.dataTransfer ) {
    if( evt.dataTransfer.items ) {
      const file = evt.dataTransfer.items[0].getAsFile();
      if(file) {
        processFile(file)
      }
    }
  }
}

function handleSelect(evt: Event) {
  const target = evt.target as HTMLInputElement
  if( evt.target && target.files ) {
    const file = target.files[0];

    if( file ) {
      processFile(file);
    }
  }
}

function handleClick() {
  if( fileupload.value ) {
    (fileupload.value as HTMLElement).click();
  }
}

function processFile(file: File) {
  const extension = file.name.substring(file.name.lastIndexOf('.'))
  if( extension !== '.bib' ) {
    console.error(`wrong file type: ${extension}`);
    return
  }
  let reader = new FileReader();
  reader.readAsText(file,'UTF-8');

  // here we tell the reader what to do when it's done reading...
  reader.onload = readerEvent => {
    if( readerEvent.target ) {
      console.debug(`got file with name ${file.name}`)
      let content = readerEvent.target.result; // this is the content!
      if( content ) {
        if( content instanceof ArrayBuffer ) {
          const enc = new TextDecoder("utf-8");
          content = enc.decode(content);
        }
        result.value = AnalyseBibFile(content, projectDataStore.categories);
        uploaded.value = Array.from({length: result.value.Entries.length}, () => true);
        triggered.value = true;
      }
    }
  }
}
</script>

<style scoped lang="scss">
.drag_n_drop--container {
  width: calc(100% - 48px);
  padding: 8px;
  margin: 8px auto;
  border: 3px dashed rgba(var(--v-theme-on-background), .75);
  border-radius: 8px;
  cursor: pointer;
  min-height: 75px;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>