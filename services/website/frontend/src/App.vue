<template>
  <v-app>
    <v-sheet>
      <v-app-bar color="transparent" :elevation="elevation" app>
        <v-container class="pa-6 fill-height">
          <v-row align="center" class="fill-height">
            <v-toolbar-title class="text-h4 font-weight-bold">ThesorTeX</v-toolbar-title>
            <v-spacer />
            <v-btn color="primary">
              Downloads
            </v-btn>
            <v-btn color="primary">
              Tutorials
            </v-btn>
          </v-row>
        </v-container>

      </v-app-bar>
    </v-sheet>
    <v-main>
      <ToolbarAndContent :hide-bar="true" @scroll="elevation=1" @no-scroll="elevation=0">
        <template #content>
          <div class="bg-container" style="background-image: linear-gradient(90deg, #0c8635, #259b71, #69beaf);">
            <v-container class="bg-transparent pa-16">
              <v-row>
                <v-col cols="12">
                  <h2 class="text-h3 font-weight-bold text-center text-white pa-4 pt-15">LaTeX bequemer nutzen</h2>
                  <p class="text-h5 text-center text-white pa-4 pb-15">Wenn du LaTeX im akademischen oder beruflichen
                    Kontext
                    nutzt, findest
                    du hier
                    verschiedene Vorlagen und Tools, die dich dabei unterstützen können.</p>
                </v-col>
              </v-row>
            </v-container>
          </div>
          <div class="bg-container">
            <v-container class="bg-transparent mb-6">
              <v-row>
                <v-col cols="6">
                  <h2 class="text-h3 font-weight-bold pt-6 pb-6">Vorlage für wissenschaftliche Arbeiten</h2>
                  <p class="text-h6 pb-6">
                    Möchtest du eine Abschluss- oder Hausarbeit mit LaTeX schreiben? Dann schau dir doch mal <a
                    href="/sdfsdf">diese
                    Vorlage</a> an.</p>
                  <v-btn color="primary">
                    Dokumentation der Vorlage
                  </v-btn>
                </v-col>
                <v-col cols="6">
                  icon
                </v-col>
              </v-row>
            </v-container>
          </div>
          <div class="bg-container" style="background-color: #f5f5f5">
            <v-container class="bg-transparent mb-6">
              <v-row>
                <v-col cols="6">
                  icon
                </v-col>
                <v-col cols="6">
                  <h2 class="text-h3 font-weight-bold">Literaturmanagement</h2>
                  <p class="text-h6">Mit diesem Tool kannst du dein Literaturverzeichnis bequem und individuell
                    zusammenbauen. Formatiere deine Literatur nach deinen Wünschen, lege Einträge händisch an oder
                    importiere deine bibliothek aus <i>Citavi</i>.</p>
                  <v-btn>
                    Mehr erfahren
                  </v-btn>
                </v-col>
              </v-row>
            </v-container>
          </div>
          <div class="bg-container">
            <v-container class="bg-transparent mb-6">
              <v-row>
                <v-col cols="6">
                  icon
                </v-col>
                <v-col cols="6">
                  <h2 class="text-h3 font-weight-bold">Vorlage für einen Lebenslauf</h2>
                  <p class="text-h6">
                    Wenn du mithilfe der hier angebotenen Unterstützungen deinen Abschluss geschafft hast, kannst du
                    dich
                    mit <a
                    href="/sdfsdf">diese
                    Vorlage</a> für deien Lebenslauf auf deinen nächsten Job bewerben.</p>
                  <v-btn>
                    Dokumentation der Vorlage
                  </v-btn>
                </v-col>
              </v-row>
            </v-container>
          </div>
        </template>
      </ToolbarAndContent>
    </v-main>
  </v-app>
</template>

<script lang="ts" setup>
import {ref} from "vue";

// data
const elevation = ref(0);

</script>

<style scoped lang="scss">
.bg-container {
  width: 100%;
  padding: 16px 0;
}

.section {
  width: 100%;
  background-color: rgba(var(--v-theme-primary), 0.9);
  padding: 25px 0;

  & .content {
    width: calc(100% - 100px);
    max-width: 800px;
    margin: 0 auto;
    background-color: rgb(var(--v-theme-background));
  }
}
</style>

<style>
a {
  color: rgb(var(--v-theme-primary));
}
</style>
