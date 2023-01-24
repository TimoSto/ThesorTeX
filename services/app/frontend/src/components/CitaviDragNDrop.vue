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

// globals
defineProps({
  title: String
});

// data
const fileupload = ref(null);

// methods
function handleDrop(evt: DragEvent) {
  evt.preventDefault();
  if (evt.dataTransfer) {
    if (evt.dataTransfer.items) {
      const file = evt.dataTransfer.items[0].getAsFile();
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