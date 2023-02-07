<template>
  <v-app>
    <v-app-bar color="primary" :elevation="elevation" style="z-index: 1000">
      <v-toolbar-title>ThesorTeX</v-toolbar-title>
    </v-app-bar>
    <v-main>
      <v-sheet class="content">
        <div style="height: 1234px"></div>
      </v-sheet>
    </v-main>
  </v-app>
</template>

<script lang="ts" setup>

// data
import {onMounted, ref} from "vue";

const elevation = ref(0);
const root = ref<HTMLElement | null>(null);
onMounted(() => {
  root.value!.querySelector(".content")!.addEventListener("scroll", (e) => {
    if ((e.target! as HTMLElement).scrollTop > 0) {
      elevation.value = 1;
    } else {
      elevation.value = 0;
    }
  });
});
</script>

<style scoped>
.content {
  padding: 0;
  margin: 0;
  width: 100%;
  overflow-y: auto;
  height: calc(100vh - 64px)
}
</style>