<template>
  <v-dialog
    v-model="errorOpened"
    width="450"
  >
    <v-card>
      <v-card-title>
        {{ title }}
      </v-card-title>
      <v-card-text>
        <p>
          {{ error }}
        </p>
        <br>
        <i18n-t
          v-if="hint"
          style="color: rgba(var(--v-theme-on-background), 0.5); font-size: 14px"
          tag="p"
          :keypath="i18nKeys.Errors.ErrorHint"
        >
          {{ hint }}
        </i18n-t>
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <a
          href="https://github.com/TimoSto/ThesorTeX/issues"
          target="_blank"
          style="text-decoration: none"
        >
          <v-btn color="primary">
            {{ t(i18nKeys.Errors.CreateIssue) }}
          </v-btn>
        </a>
        <v-btn
          color="primary"
          @click="errorOpened = false"
        >
          {{ t(i18nKeys.Common.Close) }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>

  <v-snackbar v-model="successOpened">
    <span>
      {{ success }}
    </span>

    <v-progress-linear
      absolute
      bottom
      color="primary"
      :model-value="100 - Math.floor(100 * (currentTime / 5000))"
      style="margin-top: 4px"
    />
  </v-snackbar>
</template>

<script setup lang="ts">

import {i18nKeys} from "../i18n/keys";

import {computed, ref, watch} from "vue";
import {useI18n} from "vue-i18n";

const props = defineProps({
  error: {
    type: String,
    default: ''
  },
  success: {
    type: String,
    default: ''
  },
  title: {
    type: String,
    default: ''
  },
  hint: {
    type: String,
    default: ''
  }
})

const { t } = useI18n();

const emit = defineEmits(['errorClosed', 'successClosed'])

const errorOpened = computed({
  get(): boolean {
    return props.error !== '';
  },
  set(v: boolean) {
    if( !v ) {
      emit('errorClosed')
    }
  }
});

const successOpened = computed({
  get(): boolean {
    return props.success !== '';
  },
  set(v: boolean) {
    if( !v ) {
      emit('successClosed')
    }
  }
});

const currentTime = ref(0);
const counter = ref(0);

watch(() => props.success, () => {
  if( props.success !== '' ) {
    counter.value++;
    currentTime.value = 0;
    syncSnackbarTime(counter.value);
  }
})

function syncSnackbarTime(count: number) {
  currentTime.value += 100;
  if (count === counter.value) {
    if (currentTime.value < 5000) {
      setTimeout(() => {
        syncSnackbarTime(count);
      }, 100)
    } else {
      emit('successClosed')
    }
  }
}

</script>

<style scoped>

</style>
