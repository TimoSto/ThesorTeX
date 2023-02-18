<template>
  <div class="markdown-container" v-html="parsed">

  </div>
</template>

<script lang="ts" setup>
import {computed} from "vue";
import {convertMarkdownToVuetify} from "./helper/convertMarkdownToVuetify";

const props = defineProps({
  file: String,
});

const parsed = computed(() => convertMarkdownToVuetify(props.file!));
</script>

<style scoped>

</style>