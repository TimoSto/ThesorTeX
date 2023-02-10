<template>
  <div style="background-image: linear-gradient(90deg, #0c8635, #259b71, #69beaf);">
    <v-container class="bg-transparent pa-16">
      <v-row>
        <v-col cols="4">
          <v-card>
            <v-card-text>
              <span class="text-h5 text-center"
                    style="display: inline-block">Vorlage - Wissenschaftliche Arbeiten</span>
              <TemplateIcon :hide-icon="true" style="transform: scale(0.9)" />
              <v-btn variant="text" color="primary" style="width: 100%;" class="mb-2">
                Herunterladen
              </v-btn>
              <v-btn color="primary" style="width: 100%;">
                <span style="white-space: normal;">
                  Weitere Informationen
                </span>
              </v-btn>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="4">

        </v-col>
        <v-col cols="4">

        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import TemplateIcon from "../components/TemplateIcon.vue";

export default {
  name: "DownloadsPage",
  components: {TemplateIcon}
}
</script>

<style scoped>

</style>