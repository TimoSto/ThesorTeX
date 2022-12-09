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
            @nav-back="navBack"
            @open-entry="openEntry($event)"
          />
          <EntryEditor
            v-if="i === 3 && editorType === EDITOR_TYPE_ENTRY"
            :key="`page-${i}`"
            :entry-key="pages[i-1].title"
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

// globals
const { t } = useI18n();

// data
const pages = ref([{
  title: "ThesorTeX",
  disabled: false
}]);

const drawer = ref(false);

let editorType = '';
const EDITOR_TYPE_ENTRY = 'entry';

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
  pages.value.pop();
  editorType = '';
}

function openEntry(key: string) {
  pages.value.push({
    title: key,
    disabled: false
  });
  editorType = EDITOR_TYPE_ENTRY;
}
</script>
