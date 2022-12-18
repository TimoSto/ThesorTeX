<template>
  <v-dialog
    v-model="opened"
    width="400"
  >
    <v-card>
      <v-card-title>{{ t(i18nKeys.CategoryEditor.Delete.Title) }}</v-card-title>
      <v-card-text>
        <i18n-t
          :keypath="i18nKeys.CategoryEditor.Delete.Message"
          tag="span"
        >
          <template #name>
            <i>{{ category }}</i>
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
          @click="CallCategoryDelete"
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
import DeleteCategory from "../api/projectData/categories/DeleteCategory";

// globals
const emit = defineEmits(['close', 'success', 'error'])

const { t } = useI18n();

// props
const props = defineProps({
  open: Boolean,
  category: {
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
function CallCategoryDelete() {
  DeleteCategory(props.project, props.category).then(ok => {
    if( ok ) {
      emit('success');
    } else {
      emit('error');
    }
  })
}

</script>
