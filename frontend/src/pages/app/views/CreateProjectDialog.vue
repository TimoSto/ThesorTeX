<template>
  <v-dialog width="400" v-model="opened">
    <v-card>
      <v-card-title>{{ t(i18nKeys.Overview.CreateProject) }}</v-card-title>
      <v-card-text>

        <v-text-field
          :label="t(i18nKeys.Overview.ProjectName)+'d'"
          color="primary"
          v-model="projectName"
          :rules="nameRules"
        />

        <v-row style="padding: 8px;">
          <v-spacer />
          <v-btn variant="text" color="primary" style="margin-right: 8px;">{{ t(i18nKeys.Common.Abort) }}</v-btn>
          <v-btn color="primary"
                 :disabled="!rulesAreMet || projectName.length === 0"
                 @click="Create"
          >{{ t(i18nKeys.Common.Create) }}</v-btn>
        </v-row>

      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">

import {computed, ref} from "vue";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";
import CreateProject from "../api/projects/CreateProject";
import getProjectNameRules from "../../../rules/projectNameRules";

const props = defineProps({
  open: Boolean
});

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
