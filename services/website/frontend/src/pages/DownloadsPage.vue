<template>
  <div style="background-image: linear-gradient(90deg, #0c8635, #259b71, #69beaf);">
    <v-container class="bg-transparent pa-16">
      <v-row>
        <v-col cols="4">
          <v-card>
            <v-card-text>
              <span class="text-h5 text-center font-weight-bold"
                    style="display: inline-block; width: 100%;">Vorlage - Wissenschaftliche Arbeiten</span>
              <TemplateIcon :hide-icon="true" style="display: block; margin: 0 auto; transform: scale(0.9)" />
              <v-btn variant="text" color="primary" style="width: 100%;" class="mb-2">
                Weitere Informationen
              </v-btn>
              <v-btn color="primary" style="width: 100%;">
                <span style="white-space: normal;">
                  Herunterladen
                </span>
              </v-btn>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="4">
          <v-card>
            <v-card-text>
              <span class="text-h5 text-center font-weight-bold"
                    style="display: inline-block; width: 100%;">Tool - Literaturmanagement</span>
              <TemplateIcon :hide-icon="false" style="display: block; margin: 0 auto; transform: scale(0.9)" />
              <v-btn variant="text" color="primary" style="width: 100%;" class="mb-2">
                Weitere Informationen
              </v-btn>
              <v-btn color="primary" style="width: 100%;">
                <span style="white-space: normal;">
                  Herunterladen
                </span>
              </v-btn>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="4">
          <v-card>
            <v-card-text>
              <span class="text-h5 text-center font-weight-bold"
                    style="display: inline-block; width: 100%;">Vorlage - Lebenslauf</span>
              <CVIcon style="display: block; margin: 0 auto; transform: scale(0.7)" />
              <v-btn variant="text" color="primary" style="width: 100%;" class="mb-2">
                Weitere Informationen
              </v-btn>
              <v-btn color="primary" style="width: 100%;">
                <span style="white-space: normal;">
                  Herunterladen
                </span>
              </v-btn>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
  <DownloadsTable />
</template>

<script>
import TemplateIcon from "../components/TemplateIcon.vue";
import CVIcon from "../components/CVIcon.vue";
import DownloadsTable from "../components/DownloadsTable.vue";

export default {
  name: "DownloadsPage",
  components: {DownloadsTable, CVIcon, TemplateIcon}
}
</script>

<style scoped>

</style>