<template>
  <div
    v-ripple
    class="drag-d-drop--container"
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
    {{ title }}
  </div>
</template>

<script lang="ts" setup>

import {ref} from "vue";
import {useErrorSuccessStore} from "@thesortex/vue-component-library/src/stores/ErrorSuccessStore/ErrorSuccessStore";
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {i18nKeys} from "../i18n/keys";
import {AnalyseBibFile} from "../domain/citavi/BibAnalytics";
import {Category} from "../domain/category/Category";

// globals
const props = defineProps({
  title: String,
  categories: Array<Category>
});

const errorSuccessStore = useErrorSuccessStore();

const {t} = useI18n();

// data
const fileupload = ref(null);

// methods
function handleDrop(evt: DragEvent) {
  evt.preventDefault();
  if (evt.dataTransfer) {
    if (evt.dataTransfer.items) {
      const file = evt.dataTransfer.items[0].getAsFile();
      // TODO: why null?
      if (file) {
        processFile(file);
      }
    }
  }
}

function handleSelect(evt: Event) {
  const target = evt.target as HTMLInputElement;
  if (evt.target && target.files) {
    const file = target.files[0];

    if (file) {
      processFile(file);
    }
  }
}

function handleClick() {
  if (fileupload.value) {
    (fileupload.value as HTMLElement).click();
  }
}

function processFile(file: File) {
  const extension = file.name.substring(file.name.lastIndexOf("."));
  if (extension !== ".bib") {
    console.error(`wrong file type: ${extension}`);
    errorSuccessStore.setMessage(false, t(i18nKeys.ProjectPage.WrongFileExt));
    return;
  }
  let reader = new FileReader();
  reader.readAsText(file, "UTF-8");

  // here we tell the reader what to do when it's done reading...
  reader.onload = readerEvent => {
    if (readerEvent.target) {
      console.debug(`got file with name ${file.name}`);
      let content = readerEvent.target.result; // this is the content!
      if (content) {
        if (content instanceof ArrayBuffer) {
          const enc = new TextDecoder("utf-8");
          content = enc.decode(content);
        }
        const result = AnalyseBibFile(content, props.categories!);
        console.log(result);
      }
    }
  };
}
</script>

<style lang="scss" scoped>
.drag-d-drop--container {
  width: calc(100% - 32px);
  margin: 0 auto;
  border: 3px dashed rgba(var(--v-theme-on-background), 0.75);
  border-radius: 8px;
  min-height: 50px;
  cursor: pointer;
  display: flex;
  justify-content: center;
  align-items: center;

}
</style>