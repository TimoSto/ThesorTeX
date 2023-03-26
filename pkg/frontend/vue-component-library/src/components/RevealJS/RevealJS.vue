<template>
  <div id="container">
    <div class="reveal">
      <div class="slides">
        <section data-markdown>
          <textarea data-template>
            # Slide 1
            > A paragraph with some text and a [link](http://hakim.se).
          </textarea>
        </section>
        <section data-markdown>
          <textarea data-template>
            ```js [1-2|3|4]
            let a = 1;
            let b = 2;
            let c = x => 1 + 2 + x;
            c(3);
            ```
          </textarea>
        </section>
        <section>Slide 1</section>
        <section>Slide 2</section>
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

<style scoped>
#container {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  height: 90vh;
}
</style>