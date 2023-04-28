<template>
  <v-card>
    <v-card-title>
      {{ t(i18nKeys.MainPage.CreateProject) }}
    </v-card-title>
    <v-card-text>
      <v-text-field
        v-model="projectName"
        :label="t(i18nKeys.MainPage.ProjectName)"
        :rules="[nameRules]"
        color="primary"
      />
    </v-card-text>
    <v-card-actions>
      <v-spacer />
      <v-btn
        color="primary"
        @click="emit('close')"
        :title="t(i18nKeys.Common.Abort)"
      >
        {{ t(i18nKeys.Common.Abort) }}
      </v-btn>
      <v-btn
        color="primary"
        :disabled="!rulesAreMet"
        :title="t(i18nKeys.Common.Create)"
        @click="emit('confirm', projectName)"
      >
        {{ t(i18nKeys.Common.Create) }}
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts" setup>
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {computed, ref} from "vue";
import {i18nKeys} from "../i18n/keys";
import getProjectNameRules from "../domain/projects/ProjectNameRules";

// globals
const emit = defineEmits(["close", "confirm"]);

const {t} = useI18n();

// props
const props = defineProps({
  projects: {
    required: true,
    type: Array<string>
  }
});

// data
const projectName = ref("");

const nameRules = computed(() => {
  return getProjectNameRules(props.projects, "", t);
});

// computed
const rulesAreMet = computed(() => {
  return nameRules.value(projectName.value) === true;
});
</script>

<style scoped>

</style>