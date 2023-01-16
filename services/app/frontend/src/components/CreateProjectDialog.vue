<template>
  <v-dialog
    v-model="opened"
    width="400"
  >
    <v-card>
      <v-card-title>
        {{ t(i18nKeys.MainPage.CreateProject) }}
      </v-card-title>
      <v-card-text>
        <v-text-field
          v-model="projectName"
          :label="t(i18nKeys.MainPage.ProjectName)"
        />
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn
          color="primary"
        >
          {{ t(i18nKeys.Common.Abort) }}
        </v-btn>
        <v-btn
          color="primary"
        >
          {{ t(i18nKeys.Common.Create) }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts" setup>
import {useI18n} from "@thesortex/vue-i18n-plugin";
import {computed, ref} from "vue";
import {i18nKeys} from "../i18n/keys";

// globals
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
</script>

<style scoped>

</style>