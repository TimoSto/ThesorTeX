<template>
  <v-app>
    <v-layout>
      <!-- <v-system-bar color="deep-purple darken-3"></v-system-bar> -->

      <AppbarContent
        bar-color="background"
        :level="1"
      >
        <template #bar>
          <LogoSVG
            fill="rgba(0, 136, 52, 0.8)"
            style="height: 35px; margin-left: 20px;"
            :show-title="true"
            font-color="rgba(0,0,0,0.7)"
          />

          <v-spacer />

          <v-btn
            variant="text"
            color="primary"
            @click="getStarted"
          >
            {{ t(i18nKeys.Common.GetStarted) }}
          </v-btn>
        </template>
        <template #content>
          <ContentAndImage
            image-side="right"
          >
            <template #left>
              <p class="tile_heading">
                {{ t(i18nKeys.Template.Title) }}
              </p>
              <i18n-t
                :keypath="i18nKeys.Template.Text"
                tag="p"
                class="tile_text"
              >
                <template #documentation>
                  <a
                    href="/static/documentation.pdf"
                    target="_blank"
                    style="text-decoration: none"
                  >
                    {{ t(i18nKeys.Common.Documentation) }}
                  </a>
                </template>
              </i18n-t>
              <div style="padding:8px 0">
                <a
                  href="/static/downloads/projectTemplate.zip"
                  download="projectTemplate.zip"
                  style="text-decoration: none"
                >
                  <v-btn
                    block
                    color="primary"
                  >
                    {{ t(i18nKeys.Common.Download) }}
                  </v-btn>
                </a>
              </div>
            </template>
            <template #right>
              <ThesisSVG style="max-height: 250px; max-width: 250px; margin: 0 auto; display: block" />
            </template>
          </ContentAndImage>

          <ContentAndImage
            image-side="left"
          >
            <template #left>
              <!-- todo: other image, maybe turning wheels -->
              <ThesisMonitorSVG style="max-height: 250px; max-width: 250px; margin: 0 auto; display: block" />
            </template>
            <template #right>
              <p class="tile_heading">
                {{ t(i18nKeys.App.Title) }}
              </p>
              <i18n-t
                :keypath="i18nKeys.App.Text"
                tag="p"
                class="tile_text"
              >
                <template #documentation>
                  <a
                    href="/static/documentation.pdf"
                    target="_blank"
                    style="text-decoration: none"
                  >
                    {{ t(i18nKeys.Common.Documentation) }}
                  </a>
                </template>
              </i18n-t>
              <div style="padding:8px 0">
                <v-btn
                  block
                  color="primary"
                  @click="getStarted"
                >
                  {{ t(i18nKeys.Common.GetStarted) }}
                </v-btn>
              </div>
            </template>
          </ContentAndImage>

          <ContentAndImage
            image-side="right"
          >
            <template #left>
              <p class="tile_heading">
                {{ t(i18nKeys.CV.Title) }}
              </p>
              <p class="tile_text">
                {{ t(i18nKeys.CV.Text) }}
              </p>
              <div style="padding:8px 0">
                <a
                  href="/static/downloads/cvTemplate.zip"
                  download="cvTemplate.zip"
                  style="text-decoration: none"
                >
                  <v-btn
                    block
                    color="primary"
                  >
                    {{ t(i18nKeys.Common.Download) }}
                  </v-btn>
                </a>
              </div>
            </template>
            <template #right>
              <CvSVG style="max-height: 250px; max-width: 250px; margin: 0 auto; display: block" />
            </template>
          </ContentAndImage>
        </template>
      </AppbarContent>
    </v-layout>
  </v-app>
</template>

<script setup lang="ts">
import LogoSVG from "./components/LogoSVG.vue";
import AppbarContent from "./components/AppbarContent.vue";
import ContentAndImage from "./components/ContentAndImage.vue";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "./i18n/keys";
import ThesisSVG from "./components/ThesisSVG.vue";
import ThesisMonitorSVG from "./components/ThesisMonitorSVG.vue";
import CvSVG from "./components/CvSVG.vue";

const { t } = useI18n();

function getStarted() {
  window.location.href = "/getStarted/"
}

</script>
