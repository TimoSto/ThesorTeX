<template>
  <div id="container">
    <div class="reveal">
      <div class="slides">
        <section v-for="d in docs">
          <section>{{ d.Title }}</section>
          <section v-for="p in d.Pages" class="scrollable">
            {{ p.Title }}
            <template v-for="g in p.Groups">
              <p v-if="g.Type === 'TEXT'">
              <span v-for="e in g.Elements">
                {{ e.Content }}
              </span>
              </p>
            </template>
          </section>
        </section>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import Reveal from "reveal.js";
import "reveal.js/dist/reveal.css";
import "reveal.js/dist/theme/black.css";
import "reveal.js/plugin/highlight/monokai.css";
import {onMounted} from "vue";

const props = defineProps({
  docs: {
    type: Array<{
      Title: string,
      Pages: Array<{
        Title: string,
        Groups: { Type: "TEXT", Elements: { Content: string, Style: "PLAIN" | "BOLD" | "ITALIC" | "ITALIC_BOLD" }[] }[]
      }[]>
    }>
  },
});

onMounted(() => {
  Reveal.initialize({
    embedded: true,
    plugins: [],
    markdown: {
      breaks: true,
      gfm: true,
    },
  });
});
</script>

<style scoped lang="scss">
#container {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  height: 90vh;

  & .slide {
    overflow-y: auto !important;
  }
}

.scrollable {
  overflow-y: auto !important;
  overflow-x: hidden !important;
  height: 700px;
}
</style>