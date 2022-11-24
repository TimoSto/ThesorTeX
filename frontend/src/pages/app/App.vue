<template>
  <v-app>
    <v-main>
      <HelloWorld />
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
  import HelloWorld from '../../components/HelloWorld.vue'
  import {useTheme} from "vuetify";

  const theme = useTheme();
  theme.global.name.value = 'dark';
</script>
