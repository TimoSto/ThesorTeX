<template>
  <v-app>
    <!-- <v-system-bar color="deep-purple darken-3"></v-system-bar> -->

    <v-app-bar
      color="primary"
      elevation="0"
    >
      <v-app-bar-nav-icon variant="text" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title>My files</v-toolbar-title>

      <v-spacer></v-spacer>

      <v-btn variant="text" icon="mdi-dots-vertical"></v-btn>
    </v-app-bar>

    <v-navigation-drawer
      permanent
      :rail="!drawer"
    >

    </v-navigation-drawer>

    <v-main>
      <NavigationArea
        :pages="pages"
      >
        <template v-slot:page-1>
          <div style="min-width: 150px; width: 100%; height: 100%; background-color: lightblue">
            <button @click="pages++">next</button>
          </div>
        </template>
        <template v-slot:page-2 v-if="pages > 1">
          <div style="min-width: 150px; width: 100%; height: 100%; background-color: cornflowerblue">
            <button @click="pages++">next</button>
            <button @click="pages--">prev</button>
          </div>
        </template>
        <template v-slot:page-3 v-if="pages > 2">
          <div style="min-width: 150px; width: 100%; height: 100%; background-color: yellowgreen">
            <button @click="pages++">next</button>
            <button @click="pages--">prev</button>
          </div>
        </template>
        <template v-slot:page-4 v-if="pages > 3">
          <div style="min-width: 150px; width: 100%; height: 100%; background-color: #ba68c8">
            <button @click="pages++">next</button>
            <button @click="pages--">prev</button>
          </div>
        </template>
      </NavigationArea>
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import NavigationArea from "@/components/NavigationArea";
import {ref} from "vue";

const drawer = ref(false);

const pages = ref(2);
</script>
