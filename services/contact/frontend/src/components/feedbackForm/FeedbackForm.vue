<script setup lang="ts">
import {ref} from "vue";

const opened = ref(false);
</script>

<template>

  <v-menu :close-on-content-click="false">
    <template v-slot:activator="{ props }">
      <v-btn icon color="primary" v-bind="props">
        <v-icon>mdi-delete</v-icon>
      </v-btn>
    </template>
    <v-card width="250">
      <v-card-title>Feedback</v-card-title>
      <v-card-text>
        <v-text-field width="200"></v-text-field>
        <v-textarea />
      </v-card-text>
    </v-card>
  </v-menu>

</template>

<style scoped lang="scss">

</style>