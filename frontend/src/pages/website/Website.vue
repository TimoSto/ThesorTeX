<template>
  <v-app>
    <v-layout>
      <!-- <v-system-bar color="deep-purple darken-3"></v-system-bar> -->

      <AppbarContent
        bar-color="background"
        :level="1"
      >
        <template v-slot:bar>
          <LogoSVG
            fill="rgba(0, 136, 52, 0.8)"
            style="height: 35px; margin-left: 20px;"
            :show-title="true"
            font-color="rgba(0,0,0,0.7)"
          />

          <v-spacer></v-spacer>

          <v-btn
            variant="text"
            color="primary"
          >Download</v-btn>
        </template>
        <template v-slot:content>

          <ContentAndImage
            image-side="right">
            <template v-slot:left>
              <p class="tile_heading">{{t(i18nKeys.Template.Title)}}</p>
              <i18n-t :keypath="i18nKeys.Template.Text" tag="p" class="tile_text">
                <template v-slot:documentation>
                  <a href="/static/documentation.pdf" target="_blank" style="text-decoration: none">
                    {{ t(i18nKeys.Common.Documentation) }}
                  </a>
                </template>
              </i18n-t>
              <div style="padding:8px 0">
                <a href="/static/downloads/cvTemplate.zip" download="cvTemplate.zip" style="text-decoration: none">
                  <v-btn block color="primary">{{t(i18nKeys.Common.Download)}}</v-btn>
                </a>
              </div>
            </template>
            <template v-slot:right>
              <p>TODO: Image</p>
            </template>
          </ContentAndImage>

          <ContentAndImage
            image-side="left">
            <template v-slot:left>
              image
            </template>
            <template v-slot:right>
              <p class="tile_heading">{{t(i18nKeys.App.Title)}}</p>
              <i18n-t :keypath="i18nKeys.App.Text" tag="p" class="tile_text">
                <template v-slot:documentation>
                  <a href="/static/documentation.pdf" target="_blank" style="text-decoration: none">
                    {{ t(i18nKeys.Common.Documentation) }}
                  </a>
                </template>
              </i18n-t>
              <div style="padding:8px 0">

              </div>
            </template>
          </ContentAndImage>

          <ContentAndImage
            image-side="right">
            <template v-slot:left>
              <p class="tile_heading">{{t(i18nKeys.CV.Title)}}</p>
              <p class="tile_text" v-html="t(i18nKeys.CV.Text)"></p>
              <div style="padding:8px 0">
                <a href="/static/downloads/cvTemplate.zip" download="cvTemplate.zip" style="text-decoration: none">
                  <v-btn block color="primary">{{t(i18nKeys.Common.Download)}}</v-btn>
                </a>
              </div>
            </template>
            <template v-slot:right>
              <p>TODO: Image</p>
            </template>
          </ContentAndImage>

        </template>
      </AppbarContent>
    </v-layout>
  </v-app>
</template>

<script setup lang="ts">
import LogoSVG from "@/components/LogoSVG";
import AppbarContent from "@/components/AppbarContent";
import ContentAndImage from "@/components/ContentAndImage.vue";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "./i18n/keys";

const { t } = useI18n();

</script>

<style lang="scss" scoped>
@use "@/styles/tiles.scss";
@import '../../../node_modules/typeface-roboto/index.css';
</style>

<style lang="scss">
@use "@/styles/common.scss";
</style>
