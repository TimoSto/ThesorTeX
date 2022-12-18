<template>
  <v-dialog
    v-model="opened"
    width="400"
  >
    <v-card>
      <v-card-title>{{ t(i18nKeys.EntryEditor.Delete.Title) }}</v-card-title>
      <v-card-text>
        <i18n-t
          :keypath="i18nKeys.EntryEditor.Delete.Message"
          tag="span"
        >
          <template #key>
            <i>{{ entryKey }}</i>
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
          @click="CallEntryDelete"
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
import DeleteEntry from "../api/projectData/entries/DeleteEntry";

// globals
const emit = defineEmits(['close', 'success', 'error'])

const { t } = useI18n();

const errorStore = useErrorSuccessStore();

// props
const props = defineProps({
  open: Boolean,
  entryKey: {
    type: String,
    required: true
  },
  project: {
    type: String,
    required: true
  }
});

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

// methods
function CallEntryDelete() {
  DeleteEntry(props.project, props.entryKey).then(ok => {
    if( ok ) {
      emit('success');
    } else {
      emit('error');
    }
  })
}

</script>
