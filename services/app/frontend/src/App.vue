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
        <template #activator="{ props }">
          <v-btn
            icon="mdi-dots-vertical"
            v-bind="props"
          />
        </template>
        <v-list>
          <v-list-item
            tabindex="1"
            @click="settingsOpened = true"
          >
            <template #prepend>
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
        @switch-to="triggerSwitchToProject"
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
            @nav-back="triggerNavBack"
            @open-entry="openEntry($event)"
            @open-category="openCategory"
          />
          <EntryEditor
            v-if="i === 3 && editorType === EDITOR_TYPE_ENTRY"
            :key="`page-${i}`"
            :entry-key="pages[i-1].title"
            :project-name="pages[i-2].title"
            @nav-back="triggerNavBack"
            @title-change="pages[i-1].title = $event"
          />
          <CategoryEditor
            v-if="i === 3 && editorType === EDITOR_TYPE_CATEGORY"
            :key="`page-${i}`"
            :category-name="pages[i-1].title"
            :project-name="pages[i-2].title"
            @nav-back="triggerNavBack"
            @title-change="pages[i-1].title = $event"
          />
        </template>
      </NavigationArea>
    </v-main>
    <SettingsDialog
      :open="settingsOpened"
      @close="settingsOpened = false"
    />
    <SuccessErrorDisplay
      :error="errorStore.errorMessage"
      :success="errorStore.successMessage"
      :hint="t(i18nKeys.Errors.ErrorHint)"
      :title="t(i18nKeys.Errors.Title)"
      @error-closed="errorStore.clean"
      @success-closed="errorStore.clean"
    />
    <UnsafeCloseDialog
      :open="unsafeDialogTriggered"
      @close="unsafeDialogTriggered = false"
    />
  </v-app>
</template>

<script setup lang="ts">
import NavigationArea from "./components/NavigationArea.vue";
import {computed, ref, watch} from "vue";
import Overview from "./views/OverView.vue";
import ProjectView from "./views/ProjectView.vue";
import EntryEditor from "./views/EntryEditor.vue";
import {useI18n} from "vue-i18n";
import {i18nKeys} from "./i18n/keys";
import ProjectsSidebar from "./views/ProjectsSidebar.vue";
import {useProjectsStore} from "./stores/projectsStore";
import CategoryEditor from "./views/CategoryEditor.vue";
import SettingsDialog from "./views/SettingsDialog.vue";
import {useErrorSuccessStore} from "./stores/errorSuccessStore";
import SuccessErrorDisplay from "./components/SuccessErrorDisplay.vue";
import {useUnsaveCloseStore} from "./stores/unsaveCloseStore";
import UnsafeCloseDialog from "./views/UnsafeCloseDialog.vue";

// globals
const { t } = useI18n();

const projectsStore = useProjectsStore();

const errorStore = useErrorSuccessStore();

const unsafeCloseStore = useUnsaveCloseStore();

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

const projectNames = computed(() => {
  return projectsStore.projects.map(p => p.Name);
})

const unsafeDialogTriggered = computed({
  get(): boolean {
    return unsafeCloseStore.triggered;
  },
  set(v: boolean) {
    console.log('set', v)
    unsafeCloseStore.triggered = v;
  }
});

// watchers
watch(projectNames, (nV, oV) => {
  if( pagesCount.value > 1 ) {
    const i = oV.indexOf(pages.value[1].title);
    if( i >= 0 ) {
      pages.value[1].title = nV[i];
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

function triggerNavBack() {
  unsafeCloseStore.trigger(pagesCount.value).then(allowed => {
    // this promise is resolved either when there are no unsaved changes or when the dialog is approved
    console.log(allowed);
    if( allowed ) {
      navBack();
    }
    unsafeCloseStore.tried = false;
    unsafeCloseStore.triggered = false;
  })

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

function triggerSwitchToProject(v: number) {
  unsafeCloseStore.trigger(pagesCount.value).then(allowed => {
    // this promise is resolved either when there are no unsaved changes or when the dialog is approved
    console.log(allowed);
    if( allowed ) {
      switchToProject(v)
    }
    unsafeCloseStore.tried = false;
    unsafeCloseStore.triggered = false;
  })

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
@use "./styles/common";
</style>
