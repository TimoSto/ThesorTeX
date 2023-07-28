<template>
  <v-text-field variant="underlined" :label="t(i18nKeys.Config.Port)" v-model="port" color="primary"
                prefix="http://localhost:" :rules="[portNameRules]" />
  <v-text-field variant="underlined" :label="t(i18nKeys.Config.Dir)" v-model="dir" color="primary" />
  <v-checkbox-btn :label="t(i18nKeys.Config.Open)" v-model="openBrowser" />

  <UpdateDialog :version="updateAvailable" />
</template>

<script lang="ts" setup>
import {useI18n} from "@thesortex/vue-i18n-plugin";

import {computed, ref, watch} from "vue";
import {i18nKeys} from "../i18n/keys";
import GetConfig from "../api/config/GetConfig";
import SaveConfig from "../api/config/SaveConfig";
import getPortRules from "../domain/config/PortRules";
import UpdateDialog from "../components/UpdateDialog.vue";

const emit = defineEmits(["close", "changesToSave"]);

const {t} = useI18n();

const props = defineProps({
  open: Boolean
});

// data
const port = ref("");

const dir = ref("");

const openBrowser = ref(false);

const updateAvailable = ref("");

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

const portNameRules = getPortRules(t);

watch(changesToSave, () => {
  emit("changesToSave", changesToSave.value);
});

// methods
async function saveConfig() {
  const success = await SaveConfig({
    Port: port.value,
    ProjectsDir: dir.value,
    OpenBrowser: openBrowser.value,
    UpdateAvailable: updateAvailable.value
  });

  console.log(success);

  if (success) {
    initial.value.Port = port.value;
    initial.value.ProjectsDir = dir.value;
    initial.value.OpenBrowser = openBrowser.value;
  }
}

// onmounted
GetConfig().then(cfg => {
  port.value = cfg.Port;
  dir.value = cfg.ProjectsDir;
  openBrowser.value = cfg.OpenBrowser;
  updateAvailable.value = cfg.UpdateAvailable;

  initial.value = cfg;
});

defineExpose({
  saveConfig,
});

</script>

<style scoped>

</style>