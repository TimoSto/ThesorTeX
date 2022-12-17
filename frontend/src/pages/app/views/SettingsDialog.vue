<template>
  <v-dialog
    v-model="opened"
    width="400"
  >
    <v-card>
      <v-card-title>
        {{ t(i18nKeys.Common.Settings) }}
      </v-card-title>
      <v-card-text>
        <v-text-field
          variant="underlined"
          label="Port"
          prefix="http://localhost:"
        >
          <template #append-inner>
            <v-tooltip
              text="Die Anwendung wird auf diesem Port deines Computers erreichbar sein. Der Standard ist http://localhost:8081"
              location="top"
              :max-width="200"
            >
              <template #activator="{ props }">
                <v-icon v-bind="props">
                  mdi-information
                </v-icon>
              </template>
            </v-tooltip>

          </template>
        </v-text-field>
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">
import {i18nKeys} from '../i18n/keys'
import {computed} from "vue";
import {useErrorSuccessStore} from "../stores/errorSuccessStore";

// globals
import {useI18n} from "vue-i18n";

const emit = defineEmits(['close', 'success'])

const { t } = useI18n();

const errorStore = useErrorSuccessStore();


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

</script>

<style scoped>

</style>
