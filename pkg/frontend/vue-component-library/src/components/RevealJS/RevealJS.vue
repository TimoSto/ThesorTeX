<template>
  <div id="container">
    <span id="beta-label">
      <span id="pre">

      </span>
    </span>
    <div class="reveal">
      <div class="slides">
        <section v-for="d in docs">
          <section>{{ d.Title }}</section>
          <section v-for="p in d.Pages" class="scrollable">
            {{ p.Title }}
            <template v-for="g in p.Groups">
              <p v-if="g.Type === 'TEXT'" class="small">
                <span v-for="e in g.Elements">
                  {{ e.Content }}
                </span>
              </p>
              <pre v-if="g.Type === 'CODE'" class="small">
                <code data-trim data-noescape>
                  <template v-for="e in g.Elements">
                    {{ e.Content }}
                  </template>
                </code>
              </pre>
              <ul v-if="g.Type === 'LIST'" class="small">
                <li v-for="e in g.Elements">
                  {{ e.Content }}
                </li>
              </ul>
            </template>
          </section>
        </section>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import Reveal from "reveal.js";
import Hightlight from "reveal.js/plugin/highlight/highlight";
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
    plugins: [Hightlight],
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
  position: relative;

  & .slide {
    overflow-y: auto !important;
  }
}

.scrollable {
  overflow-y: auto !important;
  overflow-x: hidden !important;
  height: 700px;
}

.small {
  font-size: 25px;
}

#beta-label {
  position: absolute;
  z-index: 1;
  right: 0;
  width: 0;
  height: 0;
  border-style: solid;
  border-width: 0 75px 75px 0;
  border-color: transparent rgb(var(--v-theme-primary)) transparent transparent;

  &:after {
    content: "beta";
    position: fixed;
    background: transparent;
    top: 7px;
    right: 20px;
    text-align: center;
    font-size: 13px;
    font-family: sans-serif;
    text-transform: uppercase;
    font-weight: bold;
    color: #fff;
    line-height: 35px;
    transform: rotate(45deg) translate(8px, -3px);
  }

  & #pre {
    position: absolute;
    z-index: 2;
    right: -75px;
    width: 0;
    height: 0;
    border-style: solid;
    border-width: 0 35px 35px 0;
    border-color: transparent black transparent transparent;
  }
}

//#container:after {
//  content: "beta";
//  position: fixed;
//  width: 80px;
//  height: 25px;
//  color: rgb(var(--v-theme-primary));
//  top: 7px;
//  left: -20px;
//  text-align: center;
//  font-size: 13px;
//  font-family: sans-serif;
//  text-transform: uppercase;
//  font-weight: bold;
//  line-height: 27px;
//  transform: rotate(-45deg);
//}
</style>