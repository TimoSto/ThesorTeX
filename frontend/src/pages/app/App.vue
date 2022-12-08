<template>
  <v-app>
    <!-- todo: use AppbarContent.vue, but how with nav? -->

    <v-app-bar
      color="primary"
      elevation="0"
    >
      <v-app-bar-nav-icon
        :disabled="pagesCount === 1"
        @click.stop="drawer = !drawer"
      />

      <NavigationBreadcrumb
        :pages="pages"
        @go-back-to="goBackTo($event)"
      />

      <v-spacer />

      <v-btn
        icon="mdi-dots-vertical"
      />
    </v-app-bar>

    <v-navigation-drawer
      permanent
      :rail="!drawer && pagesCount > 0"
    />

    <v-main>
      <NavigationArea
        :pages="pagesCount"
      >
        <template
          v-for="i in pagesCount"
          #[i]
        >
          <Overview
            v-if="i === 1"
            :key="`page-${i}`"
            @open-project="openProject($event)"
          />
          <ProjectView
            v-if="i === 2"
            :key="`page-${i}`"
            :project-name="pages[i-1].title"
            @open-entry="openEntry($event)"
          />
          <EntryEditor
            v-if="i === 3"
            :key="`page-${i}`"
            :entry-key="pages[i-1].title"
          />
        </template>
      </NavigationArea>
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import NavigationArea from "@/components/NavigationArea";
import {computed, ref} from "vue";
import NavigationBreadcrumb from "../../components/NavigationBreadcrumb.vue";
import Overview from "./views/OverView.vue";
import ProjectView from "./views/ProjectView.vue";
import EntryEditor from "./views/EntryEditor.vue";

const drawer = ref(false);

const pages = ref([{
  title: "ThesorTeX",
  disabled: false
}]);

const pagesCount = computed(() => {
  return pages.value.length;
})

function goBackTo(index: number) {
  pages.value = pages.value.slice(0, index + 1);
}

function openProject(name: string) {
  pages.value.push({
    title: name,
    disabled: false
  });
}

function openEntry(key: string) {
  pages.value.push({
    title: key,
    disabled: false
  });
}
</script>
