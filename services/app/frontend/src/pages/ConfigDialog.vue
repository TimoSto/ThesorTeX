<template>
  <v-dialog width="450" v-model="opened">
    <v-card>
      <v-card-title>
        {{ t(i18nKeys.Config.Title) }}
      </v-card-title>
      <v-card-text>
        <v-text-field variant="underlined" :label="t(i18nKeys.Config.Port)" v-model="port" color="primary"
                      prefix="http://localhost:" />
        <v-text-field variant="underlined" :label="t(i18nKeys.Config.Dir)" v-model="dir" color="primary" />
        <v-checkbox-btn :label="t(i18nKeys.Config.Open)" v-model="openBrowser" />
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn color="primary" @click="emit('close')">
          {{ t(i18nKeys.Common.Close) }}
        </v-btn>
        <v-btn color="primary">
          {{ t(i18nKeys.Common.Save) }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts" setup>
import {useI18n} from "@thesortex/vue-i18n-plugin";

import {computed, ref} from "vue";
import {i18nKeys} from "../i18n/keys";
import GetConfig from "../api/config/GetConfig";

const emit = defineEmits(["close"]);

const {t} = useI18n();

const props = defineProps({
  open: Boolean
});

// data
const port = ref("");

const dir = ref("");

const openBrowser = ref(false);

// computed
const opened = computed({
  get(): boolean {
    return props.open;
  },
  set(v: boolean) {
    emit("close");
  }
});

// onmounted
GetConfig().then(cfg => {
  port.value = cfg.Port;
  dir.value = cfg.ProjectsDir;
  openBrowser.value = cfg.OpenBrowser;
});

</script>

<style scoped>

</style>