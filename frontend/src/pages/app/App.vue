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

      <v-app-bar-title>
        ThesorTeX
        {{ titleAppendix }}
      </v-app-bar-title>

      <v-spacer />

      <v-btn
        icon="mdi-dots-vertical"
      />
    </v-app-bar>

    <v-navigation-drawer
      permanent
      :rail="!drawer"
    >
      <ProjectsSidebar
        v-if="pagesCount > 1"
        :project="pages[1].title"
        @switch-to="switchToProject"
      />
    </v-navigation-drawer>

    <v-main>
      <NavigationArea
        :pages="pagesCount"
        :instant-switch="instantSwitch"
        :navigating-back="navigatingBack"
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
            @nav-back="navBack"
            @open-entry="openEntry($event)"
            @open-category="openCategory"
          />
          <EntryEditor
            v-if="i === 3 && editorType === EDITOR_TYPE_ENTRY"
            :key="`page-${i}`"
            :entry-key="pages[i-1].title"
            :project-name="pages[i-2].title"
            @nav-back="navBack"
          />
          <CategoryEditor
            v-if="i === 3 && editorType === EDITOR_TYPE_CATEGORY"
            :key="`page-${i}`"
            :category-name="pages[i-1].title"
            :project-name="pages[i-2].title"
            @nav-back="navBack"
          />
        </template>
      </NavigationArea>
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import NavigationArea from "@/components/NavigationArea";
import {computed, ref} from "vue";
import Overview from "./views/OverView.vue";
import ProjectView from "./views/ProjectView.vue";
import EntryEditor from "./views/EntryEditor.vue";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "./i18n/keys";
import ProjectsSidebar from "./views/ProjectsSidebar.vue";
import {useProjectsStore} from "./stores/projectsStore";
import CategoryEditor from "./views/CategoryEditor.vue";

// globals
const { t } = useI18n();

const projectsStore = useProjectsStore();

// data
const pages = ref([{
  title: "ThesorTeX",
  disabled: false
}]);

const drawer = ref(false);

let editorType = '';
const EDITOR_TYPE_ENTRY = 'entry';
const EDITOR_TYPE_CATEGORY = 'category';

const instantSwitch = ref(false);

const navigatingBack = ref(false);

// computed
const pagesCount = computed(() => {
  return pages.value.length;
})

const titleAppendix = computed(() => {
  switch (pagesCount.value) {
    case 2: {
      return ` - ${ t(i18nKeys.Project.Title) }`
    }

    case 3: {
      return `- ${ t(i18nKeys.EntryEditor.Title) }`
    }

    default: {
      return ''
    }
  }
})

// methods
function openProject(name: string) {
  pages.value.push({
    title: name,
    disabled: false
  });
}

function navBack() {
  navigatingBack.value = true;
  setTimeout(() => {
    pages.value.pop();
    editorType = '';
    if( pagesCount.value === 1 ) {
      drawer.value = false;
    }
    navigatingBack.value = false;
  }, 750);
}

function openEntry(key: string) {
  pages.value.push({
    title: key,
    disabled: false
  });
  editorType = EDITOR_TYPE_ENTRY;
}

function openCategory(name: string) {
  pages.value.push({
    title: name,
    disabled: false
  });
  editorType = EDITOR_TYPE_CATEGORY;
}

function switchToProject(v: number) {
  instantSwitch.value = true;
  pages.value = pages.value.slice(0, 1);
  editorType = '';
  openProject(projectsStore.projects[v].Name);
  setTimeout(() => {
    instantSwitch.value = false;
  }, 750);
}
</script>
