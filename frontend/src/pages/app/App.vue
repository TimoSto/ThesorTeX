<template>
  <v-app>
    <!-- todo: use AppbarContent.vue, but how with nav? -->

    <v-app-bar
      color="primary"
      elevation="0"
    >
      <v-app-bar-nav-icon variant="text" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>

      <NavigationBreadcrumb
        :pages="pages"
        v-on:goBackTo="goBackTo($event)"
      />

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
        :pages="pagesCount"
      >
        <template v-for="i in pagesCount" v-slot:[i]>
          {{pagesCount}} {{i}}
          <button @click="pages.push({title: 'hi' + i, disabled: false})">next</button>
          <button @click="pages.pop()">prev</button>
          <div style="height: 2000px; width: 100%; border: 1px solid firebrick"></div>
        </template>
      </NavigationArea>
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import NavigationArea from "@/components/NavigationArea";
import {computed, ref} from "vue";
import NavigationBreadcrumb from "../../components/NavigationBreadcrumb.vue";

const drawer = ref(false);

const pages = ref([{
  title: "ThesorTeX",
  disabled: false
}]);

const pagesCount = computed({
  get() {
    return pages.value.length;
  },
  set(v: number) {

  }
})

function goBackTo(index: number) {
  pages.value = pages.value.slice(0, index + 1);
}
</script>
