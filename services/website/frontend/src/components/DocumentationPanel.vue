<template>
  <v-expansion-panel>
    <v-expansion-panel-title>
      {{ doc.Title }}
    </v-expansion-panel-title>
    <v-expansion-panel-text>
      <template v-for="g in doc.Groups">
        <p v-if="g.Type === 'TEXT'" class="text-body-1 ma-2">
          <span v-for="e in g.Elements" :class="getClass(e.Style)">{{ e.Content }}</span>
        </p>
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

<style scoped>

</style>