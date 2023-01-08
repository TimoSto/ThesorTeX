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
    >
    {{ t(i18nKeys.Project.Import) }}
  </div>
</template>

<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";
import {ref} from "vue";

const {t} = useI18n();

const fileupload = ref(null);

// methods
function handleDrop(evt: DragEvent) {
  console.log('drop', evt)
}

function handleClick() {
  if( fileupload.value ) {
    (fileupload.value as HTMLElement).click();
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