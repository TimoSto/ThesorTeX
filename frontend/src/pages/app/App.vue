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

      <v-menu
        persistent
        :min-width="250"
      >
        <template v-slot:activator="{ props }">
          <v-btn
            icon="mdi-dots-vertical"
            v-bind="props"
          />
        </template>
        <v-list>
          <v-list-item @click="settingsOpened = true">
            <template v-slot:prepend>
              <v-icon>mdi-cog</v-icon>
            </template>
            <v-list-item-title>
              {{ t(i18nKeys.Common.Settings) }}
            </v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>

    </v-app-bar>

    <v-navigation-drawer
      permanent
      :rail="!drawer"
      :rail-width="68"
    >
      <ProjectsSidebar
        v-if="pagesCount > 1"
        :project="pages[1].title"
        :open="drawer"
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
    <SettingsDialog
      :open="settingsOpened"
    />
  </v-app>
</template>

<script setup lang="ts">
import NavigationArea from "../../components/NavigationArea";
import {computed, ref} from "vue";
import Overview from "./views/OverView.vue";
import ProjectView from "./views/ProjectView.vue";
import EntryEditor from "./views/EntryEditor.vue";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "./i18n/keys";
import ProjectsSidebar from "./views/ProjectsSidebar.vue";
import {useProjectsStore} from "./stores/projectsStore";
import CategoryEditor from "./views/CategoryEditor.vue";
import SettingsDialog from "./views/SettingsDialog.vue";

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

const settingsOpened = ref(false);

// computed
const pagesCount = computed(() => {
  return pages.value.length;
})

const titleAppendix = computed(() => {
  let v = pagesCount.value;
  if( navigatingBack.value ) {
    v--;
  }
  switch (v) {
    case 2: {
      return ` - ${ t(i18nKeys.Project.Title) }`
    }

    case 3: {
      if( editorType === EDITOR_TYPE_ENTRY ) {
        return `- ${ t(i18nKeys.EntryEditor.Title) }`
      } else if (editorType === EDITOR_TYPE_CATEGORY ) {
        return `- ${ t(i18nKeys.CategoryEditor.Title) }`
      }
      return '';
    }

    default: {
      return ''
    }
  }
})

// methods
function openProject(name: string, always?: boolean) {
  if( !always && pagesCount.value > 1 ) {
    return
  }
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
  if( pagesCount.value > 2 ) {
    return
  }
  pages.value.push({
    title: key,
    disabled: false
  });
  editorType = EDITOR_TYPE_ENTRY;
}

function openCategory(name: string) {
  if( pagesCount.value > 2 ) {
    return
  }
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
  openProject(projectsStore.projects[v].Name, true);
  setTimeout(() => {
    instantSwitch.value = false;
  }, 750);
}

</script>

<style lang="scss">
@use "../../styles/common";
</style>
