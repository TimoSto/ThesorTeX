<template>
  <div class="markdown-container">
    <div v-for="(c,i) in components" :key="i">
      <h1 class="text-h4" v-if="c.Tag === 'H1'">{{ c.Content }}</h1>
      <h2 class="text-h5" v-if="c.Tag === 'H2'">{{ c.Content }}</h2>
      <h3 class="text-h6" v-if="c.Tag === 'H3'">{{ c.Content }}</h3>
      <p class="text-body-1" v-if="c.Tag === 'TEXT'" v-html="c.Content"></p>
      <v-list v-if="c.Tag === 'LIST'" density="compact">
        <v-list-item v-for="li in c.Content" :key="li" v-html="li" />
      </v-list>
      <div class="code" v-if="c.Tag === 'CODE'">
        <span class="code-line" v-for="(l,i) in c.Content" :key="i">
          {{ l }}
        </span>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {computed} from "vue";
import {ConvertMarkdownToVuetify} from "./helper/convertMarkdownToVuetify";

const props = defineProps({
  file: String,
});

const components = computed(() => ConvertMarkdownToVuetify(props.file!));

</script>

<style lang="scss">
.indented-list {
  padding-left: 24px;
}

.code {
  border: 3px solid rgba(var(--v-theme-on-background), 0.5);
  padding: 8px;
  border-radius: 4px;

  & .code-line {
    display: block;
    white-space: nowrap;
  }
}
</style>
