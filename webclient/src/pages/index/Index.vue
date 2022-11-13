<template>
  <v-app>
    <v-app-bar
      fixed
      elevate-on-scroll
      scroll-target="#index-content"
      color="background"
    >
      <LogoSVG
          fill="rgba(0, 136, 52, 0.8)"
          style="height: 35px"
          :show-title="true"
          font-color="rgba(0,0,0,0.7)"
      />

      <v-spacer />

      <v-btn text color="primary" @click="toLogin">
        {{$t(i18nDictionary.Common.Login)}}
      </v-btn>

    </v-app-bar>
    <ContentBelowBar
      bar-height="64"
      padding-x="16"
      padding-y="0"
      scroll-id="index-content"
    >

      <TextAndImage>
        <template v-slot:left>
          <div class="tile">
            <p class="tile_heading">{{$t(i18nDictionary.Template.Title)}}</p>
            <p  class="tile_text" v-html="$t(i18nDictionary.Template.Text)"></p>
            <a href="/static/downloads/projectTemplate.zip" download="projectTemplate.zip">
              <v-btn block color="primary">{{$t(i18nDictionary.Common.Download)}}</v-btn>
            </a>

          </div>
        </template>
        <template v-slot:right>
          <!--TODO: besseres Bild oder nur eine Grafik-->
          <TempalteIcon style="max-height: 250px; display: block; margin: 0 auto;"/>
        </template>
      </TextAndImage>

      <TextAndImage>
        <template v-slot:left>
          <TempalteIcon style="max-height: 250px; display: block; margin: 0 auto;" :show-icon="true"/>
        </template>
        <template v-slot:right>
          <div class="tile">
            <p class="tile_heading">{{$t(i18nDictionary.Bib.Title)}}</p>
            <p  class="tile_text" v-html="$t(i18nDictionary.Bib.Text)"></p>
            <v-btn block color="primary" style="margin-bottom: 8px" @click="toLogin">{{$t(i18nDictionary.Common.Login)}}</v-btn>
            <v-btn block text color="primary">{{$t(i18nDictionary.Common.ContinueWithoutLogin)}}</v-btn>
          </div>
        </template>
      </TextAndImage>

      <TextAndImage>
        <template v-slot:left>
          <div class="tile">
            <p class="tile_heading">{{$t(i18nDictionary.CV.Title)}}</p>
            <p  class="tile_text" v-html="$t(i18nDictionary.CV.Text)"></p>
            <a href="/static/downloads/cvTemplate.zip" download="cvTemplate.zip">
              <v-btn block color="primary">{{$t(i18nDictionary.Common.Download)}}</v-btn>
            </a>
          </div>
        </template>
        <template v-slot:right>
          <TempalteIcon style="max-height: 250px; display: block; margin: 0 auto;"/>
        </template>
      </TextAndImage>



      <v-footer padless color="background">
        <v-col cols="12" class="text-center">
          <v-icon>mdi-copyright</v-icon>
          ThesorTeX
        </v-col>
      </v-footer>

    </ContentBelowBar>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import LogoSVG from "../../common/components/LogoSVG.vue";
import ContentBelowBar from "../../common/components/ContentBelowBar.vue";
import TextAndImage from "./components/TextAndImage.vue";
import {i18nDictionary} from "./i18n/Keys";
import TempalteIcon from "./components/TempalteIcon.vue";

//TODO: Präsentationen bereich
export default Vue.extend({
  name: 'App',
  components: {TempalteIcon, TextAndImage, ContentBelowBar, LogoSVG},
  data: () => ({
    i18nDictionary: i18nDictionary
  }),

  mounted() {
    document.title = 'ThesorTeX - Info';
  },

  methods: {
    toLogin() {
      window.location.href = '/auth.html/#/';
    }
  }
});
</script>

<style lang="scss">
@import '../../../node_modules/typeface-roboto/index.css';

.tile {
  //text-align: center;
  padding: 8px;
  & .tile_heading {
    font-size: 30px;
  }
  & .tile_text {
    font-size: 20px;
    color: var(--v-accent-lighten4);
  }
}
</style>
