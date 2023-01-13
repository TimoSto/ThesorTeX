<template>
  <v-dialog
    v-model="opened"
    width="400"
  >
    <v-card>
      <v-card-title>{{ t(i18nKeys.Project.Rename) }}</v-card-title>
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
          {{ t(i18nKeys.Common.Rename) }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts" setup>

import {computed, ref, toRefs, watch} from "vue";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";
import getProjectNameRules from "../rules/projectNameRules";
import RenameProject from "../api/projects/RenameProject";
import {useProjectsStore} from "../stores/projectsStore";

// globals

const projectsStore = useProjectsStore();

const emit = defineEmits(['close', 'success'])

const { t } = useI18n();

// props
const props = defineProps({
  open: Boolean,
  projects: {
    required: true,
    type: Array<string>
  },
  initial: {
    type: String,
    required: true
  }
});

// data
const projectName = ref('');

const initialProps = toRefs(props);

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
  return getProjectNameRules(props.projects, props.initial, t)
})

const rulesAreMet = computed(() => {
  return nameRules.value[0](projectName.value) === true
})

const savePossible = computed(() => {
  return projectName.value !== props.initial && rulesAreMet.value === true;
})

// watchers
watch( () => props.open, () => {
  if( !props.open ) {
    projectName.value = '';
  }
})

// methods
async function Submit() {
  const proj = JSON.parse(JSON.stringify(projectsStore.projects.find(p => p.Name === props.initial)))
  if( proj ) {
    proj.Name = projectName.value;
    RenameProject(props.initial, proj).then(ok => {
      if( ok ) {
        projectsStore.renameProject(props.initial, projectName.value)
      }
    })
  }
}

// onload
projectName.value = initialProps.initial.value;
</script>

<style scoped>

</style>