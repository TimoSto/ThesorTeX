<template>
  <v-dialog width="450" v-model="opened">
    <v-card>
      <v-card-title>Konfiguration</v-card-title>
      <v-card-text></v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn color="primary" @click="emit('close')">
          Schließen
        </v-btn>
        <v-btn color="primary">
          Schließen
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts" setup>

import {computed} from "vue";

const emit = defineEmits(["close"]);

const props = defineProps({
  open: Boolean
});

const opened = computed({
  get(): boolean {
    return props.open;
  },
  set(v: boolean) {
    emit("close");
  }
});

</script>

<style scoped>

</style>