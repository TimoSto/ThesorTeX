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
</template>

<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";
import {ref} from "vue";
import AnalyseBibFile from "../bibanalytics/BibAnalytics";
import {useProjectDataStore} from "../stores/projectDataStore";

const {t} = useI18n();

const projectDataStore = useProjectDataStore();

const fileupload = ref(null);

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
        const result = AnalyseBibFile(content, projectDataStore.categories);
        console.log(result);
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
}
</style>