<template>
  <v-dialog
    v-model="opened"
    width="400"
  >
    <v-card>
      <v-card-title>{{ props.initial === '' ? t(i18nKeys.Overview.CreateProject) : t(i18nKeys.Project.Rename) }}</v-card-title>
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
          {{ props.initial === '' ? t(i18nKeys.Common.Create) : t(i18nKeys.Common.Rename) }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">

import {computed, ref, toRefs, watch} from "vue";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";
import CreateProject from "../api/projects/CreateProject";
import getProjectNameRules from "../rules/projectNameRules";
import {useErrorSuccessStore} from "../stores/errorSuccessStore";
import RenameProject from "../api/projects/RenameProject";
import {useProjectsStore} from "../stores/projectsStore";
import {GetProjects} from "../api/projects/GetProjects";
import ProjectMetaData from "../api/projects/ProjectMetaData";

// globals
const errorStore = useErrorSuccessStore();

const emit = defineEmits(['close', 'success'])

const { t } = useI18n();

const projectsStore = useProjectsStore();

// props
const props = defineProps({
  open: Boolean,
  projects: {
    required: true,
    type: Array<string>
  },
  initial: {
    type: String,
    default: ''
  }
});

const initialProps = toRefs(props);

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
    projectName.value = props.initial ? props.initial : '';
  }
})

watch( () => props.initial, (nV) => {
  projectName.value = nV;
})

// methods
async function Submit() {
  if( props.initial === '' ) {
    const resp = await CreateProject(projectName.value);
    //todo: mapping of status code to error message
    errorStore.handleResponse(resp.Status === 200, t(i18nKeys.Success.CreateProject), t(i18nKeys.Errors.ErrorCreating))
    if( resp.Status === 200 ) {
      emit('success', resp.Project);
    }
    emit('close');
  } else {
    const proj = JSON.parse(JSON.stringify(projectsStore.projects.find(p => p.Name === props.initial)))
    if( proj ) {
      proj.Name = projectName.value;
      RenameProject(props.initial, proj).then(ok => {
        if( ok ) {
          GetProjects().then((p: ProjectMetaData[]) => {
            const v = !p ? [] : p;
            projectsStore.setProjects(v);
          })
        }
      })
    }
  }
}

// onload
if( props.initial ) {
  projectName.value = initialProps.initial.value;
}

</script>

<style scoped>

</style>
