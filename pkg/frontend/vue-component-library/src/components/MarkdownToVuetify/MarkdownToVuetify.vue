<template>
  <div class="markdown-container" v-html="parsed">

  </div>
</template>

<script lang="ts" setup>
import {computed} from "vue";
import {convertMarkdownToHTML} from "./helper/convertMarkdownToHTML";

const props = defineProps({
  file: String,
});

const parsed = computed(() => convertMarkdownToHTML(props.file!));
</script>

<style lang="scss">
.indented-list {
  padding-left: 24px;
}

.code {
  border: 3px solid rgba(var(--v-theme-on-background), 0.5);
  padding: 8px;
  border-radius: 4px;
}
</style>