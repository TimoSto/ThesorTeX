<template>
  <div id="container">
    <div class="reveal">
      <div class="slides">
        <section data-markdown v-for="d in docs">
          <textarea data-template class="slide">
            # {{d.Title}}
            {{ d.Content }}
          </textarea>
        </section>
      </div>
    </div>
  </div>
</template>

<script>
import Reveal from "reveal.js";
import Markdown from "reveal.js/plugin/markdown/markdown.js";
import Highlight from "reveal.js/plugin/highlight/highlight.js";
import "reveal.js/dist/reveal.css";
import "reveal.js/dist/theme/black.css";
import "reveal.js/plugin/highlight/monokai.css";

export default {
  name: "RevealJS",
  props: ["docs"],
  mounted() {
    Reveal.initialize({
      embedded: true,
      plugins: [Markdown, Highlight],
      markdown: {
        breaks: true,
        gfm: true,
      },
    });
  },
};
</script>

<style scoped lang="scss">
#container {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  height: 90vh;

  & .slide {
    overflow-y: auto !important;
  }
}
</style>