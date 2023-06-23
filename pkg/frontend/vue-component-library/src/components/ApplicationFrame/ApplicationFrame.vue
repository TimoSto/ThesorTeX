<script setup lang="ts">
const props = defineProps({
  hasSidebar: Boolean,
});
</script>

<template>
  <v-app>
    <v-app-bar
      color="primary"
      elevation="0"
    >
      <v-app-bar-nav-icon
        :disabled="!hasSidebar"
      />
    </v-app-bar>
  </v-app>
</template>

<style scoped lang="scss">

</style>