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
        <v-btn color="primary" @click="saveConfig" :disabled="!changesToSave">
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
import SaveConfig from "../api/config/SaveConfig";

const emit = defineEmits(["close"]);

const {t} = useI18n();

const props = defineProps({
  open: Boolean
});

// data
const port = ref("");

const dir = ref("");

const openBrowser = ref(false);

// importing interface gives error
const initial = ref({} as {
  Port: string,
  ProjectsDir: string,
  OpenBrowser: boolean
});

// computed
const opened = computed({
  get(): boolean {
    return props.open;
  },
  set(v: boolean) {
    emit("close");
  }
});

const changesToSave = computed(() => {
  return port.value !== initial.value.Port || dir.value !== initial.value.ProjectsDir || openBrowser.value !== initial.value.OpenBrowser;
});

// methods
async function saveConfig() {
  const success = await SaveConfig({
    Port: port.value,
    ProjectsDir: dir.value,
    OpenBrowser: openBrowser.value
  });

  if (success) {
    initial.value.Port = port.value;
    initial.value.ProjectsDir = dir.value;
    initial.value.OpenBrowser = openBrowser.value;
  }

  console.log(initial);
}

// onmounted
GetConfig().then(cfg => {
  port.value = cfg.Port;
  dir.value = cfg.ProjectsDir;
  openBrowser.value = cfg.OpenBrowser;

  initial.value = cfg;
});

</script>

<style scoped>

</style>