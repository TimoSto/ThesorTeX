<template>
  <v-expansion-panel>
    <v-expansion-panel-title>
      {{ doc.Title }}
    </v-expansion-panel-title>
    <v-expansion-panel-text>
      <template v-for="(g, i) in doc.Groups">
        <p v-if="g.Type === 'TEXT'" class="text-body-1 ma-2">
          <template v-for="(e, j) in g.Elements">
            <span v-if="e.Style != 'LINK-HREF' && e.Style != 'LINK-TITLE'"
                  :class="getClass(e.Style)">{{ e.Content }}</span>
            <a v-if="e.Style == 'LINK-HREF'" :href="e.Content" target="_blank">{{ g.Elements[j - 1].Content }}</a>
          </template>
        </p>
        <div v-if="g.Type === 'CODE'" class="code-container pa-2 ml-2 mr-2">
          <p v-for="e in g.Elements" class="text-body-1"> {{ e.Content }} </p>
        </div>
        <v-list v-if="g.Type === 'LIST'" density="compact">
          <v-list-item v-for="e in g.Elements">
            {{ e.Content }}
          </v-list-item>
        </v-list>
        <ImageViewer v-if="g.Type === 'IMAGE'" :image="`/documentation/images/${g.Elements[0].Content}`"
                     image-styles="border: 1px solid rgb(var(--v-theme-on-background)); width: calc(100% - 40px); margin: 0 auto;">
        </ImageViewer>
      </template>
    </v-expansion-panel-text>
  </v-expansion-panel>
</template>

<script lang="ts" setup>
import {Documentation} from "../api/GetDocumentation";
import {PropType} from "vue";

const props = defineProps({
  doc: {
    type: Object as PropType<Documentation>,
    required: true
  }
});

function getClass(s: string): string {
  let className = "";

  if (s === "ITALIC" || s === "ITALIC-BOLD") {
    className = "font-italic";
  }

  if (s === "BOLD" || s === "ITALIC-BOLD") {
    className += " font-weight-bold";
  }

  return className.trimStart();
}
</script>

<style scoped lang="scss">
.code-container {
  width: calc(100% - 20px);
  margin: 0 auto;
  background-color: rgba(0, 0, 0, 0.71);
  border-radius: 8px;

  & p {
    color: white;
    opacity: 0.9;
  }
}
</style>