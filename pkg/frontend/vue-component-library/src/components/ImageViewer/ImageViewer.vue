<template>
  <v-img :src="image" @click="opened=true" style="cursor: pointer" />
  <v-dialog v-model="opened" fullscreen>
    <v-img :src="image" style="max-height: calc(100% - 50px); margin-top: 25px;" />
    <v-btn icon position="fixed" size="large" style="right: 25px; top: 25px;" @click="opened=false">
      <v-icon size="large">mdi-close</v-icon>
    </v-btn>
  </v-dialog>
</template>

<script lang="ts" setup>
import {ref} from "vue";

defineProps({
  image: {
    type: String,
    required: true
  }
});

const opened = ref(false);
</script>

<style scoped>

</style>