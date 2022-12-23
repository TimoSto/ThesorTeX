<template>
  <v-dialog
    v-model="opened"
    width="450"
  >
    <v-card>
      <v-card-title>
        {{ t(i18nKeys.Common.Settings) }}
      </v-card-title>
      <v-card-text>
        <v-text-field
          v-model="configData.Port"
          variant="underlined"
          :label="t(i18nKeys.Settings.Port)"
          prefix="http://localhost:"
          :rules="portRules"
        >
          <template #append-inner>
            <v-tooltip
              :text="t(i18nKeys.Settings.PortHint)"
              location="top"
              :max-width="200"
            >
              <template #activator="{ props }">
                <v-icon v-bind="props">
                  mdi-information
                </v-icon>
              </template>
            </v-tooltip>
          </template>
        </v-text-field>
        <v-text-field
          v-model="configData.ProjectsDir"
          variant="underlined"
          :label="t(i18nKeys.Settings.ProjectFolder)"
        >
          <template #append-inner>
            <v-tooltip
              :text="t(i18nKeys.Settings.ProjectFolderHint)"
              location="top"
              :max-width="200"
            >
              <template #activator="{ props }">
                <v-icon v-bind="props">
                  mdi-information
                </v-icon>
              </template>
            </v-tooltip>
          </template>
        </v-text-field>
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn
          color="primary"
          @click="opened = false"
        >
          {{ t(i18nKeys.Common.Close) }}
        </v-btn>
        <v-btn
          color="primary"
          :disabled="!configChanged"
          @click="Save"
        >
          {{ t(i18nKeys.Common.Save) }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">
import {i18nKeys} from '../i18n/keys'
import {computed, ref} from "vue";
import {useErrorSuccessStore} from "../stores/errorSuccessStore";

// globals
import {useI18n} from "vue-i18n";
import {ConfigData} from "../api/config/ConfigData";
import ReadConfigData from "../api/config/ReadConfigData";
import SaveConfigData from "../api/config/SaveConfigData";
import GetPortRules from "../rules/portRules";

const emit = defineEmits(['close'])

const { t } = useI18n();

const errorStore = useErrorSuccessStore();


// props

const props = defineProps({
  open: Boolean,
});

// data

const configData = ref({} as ConfigData);

let initialConfigData = ref({} as ConfigData);

const portRules = GetPortRules(t);

// computed
const opened = computed({
  get() {
    return props.open;
  },
  set(v: boolean) {
    if( !v ) {
      configData.value = initialConfigData.value;
      emit('close');
    }
  }
});

const configChanged = computed(() => {
  return initialConfigData.value.Port != configData.value.Port ||
    initialConfigData.value.ProjectsDir != configData.value.ProjectsDir;
})

// functions
function Save() {
  SaveConfigData(configData.value).then(ok => {
    errorStore.handleResponse(ok, t(i18nKeys.Success.SaveConfig), t(i18nKeys.Errors.ErrorSaving));
    initialConfigData.value.Port = configData.value.Port;
    initialConfigData.value.ProjectsDir = configData.value.ProjectsDir;
  })
}

// onload
ReadConfigData().then(data => {
  initialConfigData.value.Port = data.Port;
  initialConfigData.value.ProjectsDir = data.ProjectsDir;
  configData.value = data;
});

</script>

<style scoped>

</style>
