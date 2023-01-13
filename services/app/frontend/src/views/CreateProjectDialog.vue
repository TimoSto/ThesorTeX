<template>
  <v-dialog
    v-model="opened"
    width="400"
  >
    <v-card>
      <v-card-title>{{ t(i18nKeys.Overview.CreateProject) }}</v-card-title>
      <v-card-text>
        <v-text-field
          v-model="projectName"
          :label="t(i18nKeys.Overview.ProjectName)"
          color="primary"
          :rules="nameRules"
        />
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn
          color="primary"
          @click="emit('close')"
        >
          {{ t(i18nKeys.Common.Abort) }}
        </v-btn>
        <v-btn
          color="primary"
          :disabled="!savePossible"
          @click="Submit"
        >
          {{ t(i18nKeys.Common.Create) }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">

import {computed, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";
import CreateProject from "../api/projects/CreateProject";
import getProjectNameRules from "../rules/projectNameRules";
import {useErrorSuccessStore} from "../stores/errorSuccessStore";

// globals
const errorStore = useErrorSuccessStore();

const emit = defineEmits(['close', 'success'])

const { t } = useI18n();

// props
const props = defineProps({
  open: Boolean,
  projects: {
    required: true,
    type: Array<string>
  }
});

// data
const projectName = ref('');

// computed
const opened = computed({
  get() {
    return props.open;
  },
  set(v: boolean) {
    if( !v ) {
      emit('close');
    }
  }
});

const nameRules = computed(() => {
  return getProjectNameRules(props.projects, "", t)
})

const rulesAreMet = computed(() => {
  return nameRules.value[0](projectName.value) === true
})

const savePossible = computed(() => {
  return rulesAreMet.value === true;
})

// watchers
watch( () => props.open, () => {
  if( !props.open ) {
    projectName.value = '';
  }
})

// methods
async function Submit() {
  const resp = await CreateProject(projectName.value);
  //todo: mapping of status code to error message
  errorStore.handleResponse(resp.Status === 200, t(i18nKeys.Success.CreateProject), t(i18nKeys.Errors.ErrorCreating))
  if( resp.Status === 200 ) {
    emit('success', resp.Project);
  }
  emit('close');
}

</script>

<style scoped>

</style>
