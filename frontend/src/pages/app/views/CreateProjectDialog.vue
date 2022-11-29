<template>
  <v-dialog width="400" v-model="opened">
    <v-card>
      <v-card-title>{{ t(i18nKeys.Overview.CreateProject) }}</v-card-title>
      <v-card-text>

        <v-text-field
          :label="t(i18nKeys.Overview.ProjectName)"
          color="primary"
          v-model="projectName"
          :rules="nameRules"
        />

      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn
          color="primary"
          @click="emit('close')"
        >{{ t(i18nKeys.Common.Abort) }}</v-btn>
        <v-btn
          color="primary"
          :disabled="!rulesAreMet || projectName.length === 0"
          @click="Create"
        >{{ t(i18nKeys.Common.Create) }}</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">

import {computed, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";
import CreateProject from "../api/projects/CreateProject";
import getProjectNameRules from "../../../rules/projectNameRules";

const props = defineProps({
  open: Boolean
});

watch( () => props.open, () => {
  if( !props.open ) {
    projectName.value = '';
  }
})

const emit = defineEmits(['close'])

const { t } = useI18n();

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
  return getProjectNameRules([])
})

const rulesAreMet = computed(() => {
  return nameRules.value[0](projectName.value) === true
})

const projectName = ref('');

async function Create() {
  const newProject = await CreateProject(projectName.value);
  if( newProject.Name !== '' ) {
    emit('close');
  }
}

</script>

<style scoped>

</style>
