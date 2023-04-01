<template>
  <div id="container">
    <div class="reveal">
      <div class="slides">
        <section>
          <section>slide 1</section>
          <section>slide 1.2</section>
        </section>
        <section>
          <section>slide 2</section>
          <section>slide 2.2</section>
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
  height: 90vh;

  & .slide {
    overflow-y: auto !important;
  }
}
</style>