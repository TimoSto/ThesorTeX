<template>
  <v-app>
    <v-app-bar
      fixed
      elevation="0"
      color="background"
    >
      <LogoSVG
          fill="rgba(0, 136, 52, 0.8)"
          style="height: 35px"
          :show-title="true"
          font-color="rgba(0,0,0,0.7)"
      />

      <v-spacer />

      <v-btn text color="primary">
        {{$t(i18nDictionary.Login)}}
      </v-btn>

    </v-app-bar>
    <ContentBelowBar
      bar-height="64"
      padding-x="16"
      padding-y="0"
    >

      <TextAndImage
        style="height: 400px"
      >
        <template v-slot:left>
          <div class="tile">
            <p class="tile_heading">{{$t(i18nDictionary.Template.Title)}}</p>
            <p  class="tile_text" v-html="$t(i18nDictionary.Template.Text)"></p>
            <v-btn block color="primary">{{$t(i18nDictionary.Common.Download)}}</v-btn>
          </div>
        </template>
        <template v-slot:right>
          <!--TODO: besseres Bild oder nur eine Grafik-->
          <v-img
            src="/static/images/index/toc.png"
            contain
            height="300"
            width="200"
            style="margin: 0 auto;"
          ></v-img>
        </template>
      </TextAndImage>

      <TextAndImage>
        <template v-slot:left>

        </template>
        <template v-slot:right>
          <div class="tile">
            <p class="tile_heading">{{$t(i18nDictionary.CV.Title)}}</p>
            <p  class="tile_text" v-html="$t(i18nDictionary.CV.Text)"></p>
            <v-btn block color="primary">{{$t(i18nDictionary.Common.Download)}}</v-btn>
          </div>
        </template>
      </TextAndImage>

    </ContentBelowBar>

    <v-footer padless color="background">
      <v-col class="text-center" cols="12">
        <v-icon>mdi-copyright</v-icon>
        ThesorTeX 2022
      </v-col>
    </v-footer>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import LogoSVG from "../../commonComponents/LogoSVG.vue";
import ContentBelowBar from "../../commonComponents/ContentBelowBar.vue";
import TextAndImage from "./components/TextAndImage.vue";
import {i18nDictionary} from "./i18n/Keys";

//TODO: Präsentationen bereich
export default Vue.extend({
  name: 'App',
  components: {TextAndImage, ContentBelowBar, LogoSVG},
  data: () => ({
    i18nDictionary: i18nDictionary
  }),

  mounted() {
    document.title = 'ThesorTeX - Info';
  }
});
</script>

<style lang="scss">
.tile {
  //text-align: center;
  & .tile_heading {
    font-size: 30px;
  }
  & .tile_text {
    font-size: 20px;
    color: var(--v-accent-lighten4);
  }
}
</style>
