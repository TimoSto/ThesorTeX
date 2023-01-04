<template>
  <v-dialog
    v-model="opened"
    width="450"
  >
    <v-card>
      <v-card-title>
        {{ t(i18nKeys.UnsafeClose.Title) }}
      </v-card-title>
      <v-card-text>
        {{ t(i18nKeys.UnsafeClose.Message) }}
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn
          color="primary"
          @click="abort"
        >
          {{ t(i18nKeys.Common.Abort) }}
        </v-btn>
        <v-btn
          color="primary"
          @click="accept"
        >
          {{ t(i18nKeys.Common.Continue) }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts" setup>
// globals
import {useI18n} from "vue-i18n";
import {computed} from "vue";
import {i18nKeys} from "../i18n/keys";
import {useUnsaveCloseStore} from "../stores/unsaveCloseStore";

const emit = defineEmits(['close'])

const { t } = useI18n();

const unsafeCloseStore = useUnsaveCloseStore();

// props

const props = defineProps({
  open: Boolean,
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
function abort() {
  unsafeCloseStore.promiseResolve(false);
}

function accept() {
  unsafeCloseStore.promiseResolve(true);
}
</script>

<style scoped>

</style>
