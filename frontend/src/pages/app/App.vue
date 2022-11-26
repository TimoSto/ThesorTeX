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
        :pages="1"
      >
        <template v-slot:page-1>
          <div style="width: 100%; height: 100%; background-color: lightblue">
            f1
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
</script>
