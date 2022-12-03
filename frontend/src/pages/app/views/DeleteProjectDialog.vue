<template>
  <v-dialog
    v-model="opened"
    width="400"
  >
    <v-card>
      <v-card-title>{{ t(i18nKeys.Overview.DeleteProject) }}</v-card-title>
      <v-card-text>
        <i18n-t
          :keypath="i18nKeys.Overview.SureDeleteProject"
          tag="span"
        >
          <template #project>
            <i>{{ project }}</i>
          </template>
        </i18n-t>
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
          @click="CallProjectDelete"
        >
          {{ t(i18nKeys.Common.Delete) }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">

import {computed} from "vue";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "../i18n/keys";
import DeleteProject from "../api/projects/DeleteProject";
import {useErrorSuccessStore} from "../stores/errorSuccessStore";

const props = defineProps({
  open: Boolean,
  project: {
    type: String,
    required: true
  }
});

const emit = defineEmits(['close', 'success'])

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

const errorStore = useErrorSuccessStore();

function CallProjectDelete() {
  DeleteProject(props.project).then(ok => {
    errorStore.handleResponse(ok, t(i18nKeys.Success.DeleteProject), t(i18nKeys.Errors.ErrorDeleting))
    if( ok ) {
      emit('success')
    }
    emit('close');
  })
}

</script>
