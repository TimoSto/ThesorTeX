<template>
  <v-app>
    <!-- todo: use AppbarContent.vue, but how with nav? -->

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
        <template v-for="i in pages" v-slot:[i]>
          <div style="min-width: 150px; min-height: 100%; background-color: cornflowerblue">
            <button @click="pages++">next</button>
            <button @click="pages--">prev</button>
            <div style="height: 2000px"></div>
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
