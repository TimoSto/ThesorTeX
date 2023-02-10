<template>
  <v-app>
    <v-sheet>
      <v-app-bar color="transparent" :elevation="elevation" app>
        <v-container class="pa-6 fill-height">
          <v-row align="center" class="fill-height">
            <v-toolbar-title class="text-h4 font-weight-bold">ThesorTeX</v-toolbar-title>
            <v-spacer />
            <v-btn color="primary" to="/downloads">
              Downloads
            </v-btn>
            <v-btn color="primary" to="/tutorials">
              Tutorials
            </v-btn>
          </v-row>
        </v-container>

      </v-app-bar>
    </v-sheet>
    <v-main>
      <ToolbarAndContent :hide-bar="true" @scroll="elevation=1" @no-scroll="elevation=0">
        <template #content>
          <router-view />
        </template>
      </ToolbarAndContent>
    </v-main>
  </v-app>
</template>

<script lang="ts" setup>
import {ref} from "vue";

// data
const elevation = ref(0);

</script>

<style scoped lang="scss">
.bg-container {
  width: 100%;
  padding: 16px 0;
}

.section {
  width: 100%;
  background-color: rgba(var(--v-theme-primary), 0.9);
  padding: 25px 0;

  & .content {
    width: calc(100% - 100px);
    max-width: 800px;
    margin: 0 auto;
    background-color: rgb(var(--v-theme-background));
  }
}
</style>

<style>
a {
  color: rgb(var(--v-theme-primary));
}
</style>
