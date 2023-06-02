<script setup lang="ts">
import {onMounted, ref} from "vue";

const props = defineProps({
  title: String,
  keydownTarget: Document
});

const opened = ref(false);
let ctrlPressed = false;

onMounted(() => {
  props.keydownTarget!.addEventListener("keypress", (e: KeyboardEvent) => {
    if (e.key === "Ctrl") {
      if (!ctrlPressed) {
        ctrlPressed = true;
        setTimeout(() => {
          ctrlPressed = false;
        }, 300);
      } else {
        opened.value = true;
      }
    }
  });
});
</script>

<template>
  <v-dialog v-model="opened" width="500" height="500">
    <v-card>
      <v-card-title>{{ title }}</v-card-title>
    </v-card>
  </v-dialog>
</template>

<style scoped lang="scss">

</style>