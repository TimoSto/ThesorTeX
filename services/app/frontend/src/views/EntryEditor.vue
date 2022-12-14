<template>
  <AppbarContent
    :level="2"
    bar-color="background"
  >
    <template #bar>
      <v-app-bar-nav-icon
        @click="emit('navBack')"
      >
        <v-icon>mdi-arrow-left</v-icon>
      </v-app-bar-nav-icon>

      <v-toolbar-title>
        {{ props.entryKey }}
      </v-toolbar-title>

      <v-spacer />

      <v-btn
        icon
        :disabled="!savePossible"
        @click="CallSaveEntry"
      >
        <v-icon>mdi-content-save</v-icon>
      </v-btn>
      <v-btn
        icon
        @click="deleteOpened = true"
      >
        <v-icon>mdi-delete</v-icon>
      </v-btn>
    </template>

    <template #content>
      <div style="max-height: calc(100vh - 112px);">
        <v-expansion-panels
          multiple
          variant="accordion"
          :model-value="[0,1]"
        >
          <v-expansion-panel>
            <v-expansion-panel-title v-ripple>
              {{ t(i18nKeys.EntryEditor.General.Title) }}
            </v-expansion-panel-title>
            <v-expansion-panel-text>
              <ResponsiveTable
                :headers="generalHeaders"
                :rows="generalRows"
                :disable-ripple="true"
              >
                <template #0-1>
                  <v-text-field
                    v-model="entryKey"
                    color="primary"
                    variant="underlined"
                    :rules="keyRules"
                  />
                </template>
                <template #1-1>
                  <v-select
                    v-model="entryCategory"
                    color="primary"
                    variant="underlined"
                    :items="availableCategories"
                  />
                </template>
              </ResponsiveTable>
            </v-expansion-panel-text>
          </v-expansion-panel>
          <v-expansion-panel>
            <v-expansion-panel-title v-ripple>
              {{ t(i18nKeys.EntryEditor.Fields.Title) }}
            </v-expansion-panel-title>
            <v-expansion-panel-text>
              <p
                class="example"
              >
                <span class="prefix">Literaturverzeichnis: </span>
                <span v-html="bibEntryPreview" />
              </p>
              <p
                class="example"
              >
                <span class="prefix">Zitat: </span>
                <span v-html="bibCitePreview" />
              </p>
              <ResponsiveTable
                :headers="fieldsHeaders"
                :rows="fieldRows"
                :disable-ripple="true"
              >
                <template
                  v-for="(_,i) in fieldRows"
                  :key="`field-${i}`"
                  #[getSlotName(i)]
                >
                  <v-text-field
                    v-model="entryFields[i]"
                    color="primary"
                    variant="underlined"
                  />
                </template>
              </ResponsiveTable>
            </v-expansion-panel-text>
          </v-expansion-panel>
        </v-expansion-panels>
      </div>
    </template>
  </AppbarContent>
  <DeleteEntryDialog
    :entry-key="props.entryKey"
    :open="deleteOpened"
    :project="projectName"
    @close="deleteOpened=false"
    @success="NavBackAndRefresh"
  />
</template>

<script setup lang="ts">
import AppbarContent from "../components/AppbarContent.vue";
import {i18nKeys} from "../i18n/keys";
import {useI18n} from "vue-i18n";
import ResponsiveTable, {ResponsiveTableCell, ResponsiveTableHeaderCell} from "../components/ResponsiveTable.vue";
import {computed, ref, watch} from "vue";
import {useProjectDataStore} from "../stores/projectDataStore";
import SaveEntry from "../api/projectData/entries/SaveEntry";
import {BibEntry} from "../api/projectData/entries/BibEntry";
import CheckEntryChanged from "../api/projectData/entries/CheckEntryChanged";
import GetProjectData from "../api/projectData/GetProjectData";
import {useErrorSuccessStore} from "../stores/errorSuccessStore";
import GenerateEntry from "../api/projectData/entries/GenerateEntry";
import GenerateCite from "../api/projectData/entries/GenerateCite";
import GetEntryKeyRules from "../rules/entryKeyRules";
import DeleteEntryDialog from "./DeleteEntryDialog.vue";
import {useProjectsStore} from "../stores/projectsStore";
import {useUnsaveCloseStore} from "../stores/unsaveCloseStore";

// globals
const emit = defineEmits(['navBack', 'titleChange']);

const { t } = useI18n();

const projectDataStore = useProjectDataStore();

const errorSuccessStore = useErrorSuccessStore();

const projectsStore = useProjectsStore();

const unsafeCloseStore = useUnsaveCloseStore();

// props
const props = defineProps({
  entryKey: {
    type: String,
    required: true
  },
  projectName: {
    type: String,
    required: true
  }
});

// data
const deleteOpened = ref(false);

const generalHeaders: ResponsiveTableHeaderCell[] = [
  {
    width: '40%',
    minWidth: '',
    content: t(i18nKeys.Common.Attribute),
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: '60%',
    minWidth: '',
    content: t(i18nKeys.Common.Value),
    icon: '',
    hideUnder: -1,
    event: ''
  }
]

const generalRows: ResponsiveTableCell[][] = [
  [
    {
      content: t(i18nKeys.EntryEditor.General.Key),
      icon: '',
      event: '',
      hideUnder: -1
    },
    {
      content: '',
      icon: '',
      event: '',
      hideUnder: -1,
      slot: true
    },
  ],
  [
    {
      content: t(i18nKeys.EntryEditor.General.Category),
      icon: '',
      event: '',
      hideUnder: -1
    },
    {
      content: '',
      icon: '',
      event: '',
      hideUnder: -1,
      slot: true
    },
  ],
];

const fieldsHeaders: ResponsiveTableHeaderCell[] = [
  {
    width: '40%',
    minWidth: '',
    content: t( i18nKeys.Common.Attribute ),
    icon: '',
    hideUnder: -1,
    event: ''
  },
  {
    width: '60%',
    minWidth: '',
    content: t( i18nKeys.Common.Value ),
    icon: '',
    hideUnder: -1,
    event: ''
  }
]

const entryKey = ref('');

const entryCategory = ref('');

const entryFields = ref([] as string[])

// computed
const availableCategories = computed((): string[] => {
  return projectDataStore.categories.map(c => c.Name);
});

const initialEntry = computed(() => {
  console.log('r')
  return projectDataStore.entries.find(e => e.Key === props.entryKey)
});

const fieldRows = computed(() => {
  const rows: ResponsiveTableCell[][] = [];
  const category = projectDataStore.categories.find(c => c.Name === entryCategory.value);
  if( !category ) {
    return [];
  }
  category.Fields.forEach(f => {
    rows.push([
      {
        content: f.Field,
        icon: '',
        event: '',
        hideUnder: -1
      },
      {
        content: '',
        icon: '',
        event: '',
        hideUnder: -1,
        slot: true
      },
    ])
  });
  const bibFieldNames = category.Fields.map(f => f.Field);
  category.CiteFields.forEach(f => {
    if( bibFieldNames.indexOf(f.Field) === -1 ) {
      rows.push([
        {
          content: f.Field,
          icon: '',
          event: '',
          hideUnder: -1
        },
        {
          content: '',
          icon: '',
          event: '',
          hideUnder: -1,
          slot: true
        },
      ])
    }
  })
  return rows;
})

const rulesAreMet = computed(() => {
  return keyRules.value[0](entryKey.value) === true;
})

const savePossible = computed(() => {
  if( initialEntry.value ) {
    return CheckEntryChanged(initialEntry.value, entryKey.value, entryCategory.value, entryFields.value) && rulesAreMet.value;
  }
  return rulesAreMet.value;
})

const bibEntryPreview = computed(() => {
  const entry: BibEntry = {
    Key: entryKey.value,
    Category: entryCategory.value,
    Fields: entryFields.value,
    Comment: '',
    CiteNumber: 0,
    Example: ''
  }
  return GenerateEntry(entry, projectDataStore.categories)
})

const bibCitePreview = computed(() => {
  const entry: BibEntry = {
    Key: entryKey.value,
    Category: entryCategory.value,
    Fields: entryFields.value,
    Comment: '',
    CiteNumber: 0,
    Example: ''
  }
  return GenerateCite(entry, projectDataStore.categories)
})

const keyRules = computed(() => {
  return GetEntryKeyRules(projectDataStore.entries.map(e => e.Key), initialEntry.value ? initialEntry.value.Key : '', t);
})

const closeTried = computed(() => {
  return unsafeCloseStore.tried;
});

// watchers
watch(initialEntry, () => {
  if( initialEntry.value ) {
    entryKey.value = initialEntry.value.Key;
    entryCategory.value = initialEntry.value.Category;
    entryFields.value = initialEntry.value.Fields.map(f => f);
  }
});

watch(closeTried, (nV) => {
  if( nV ) {
    if( unsafeCloseStore.fromPage === 3 ) {
      if( !savePossible.value ) {
        unsafeCloseStore.promiseResolve(true);
      } else {
        unsafeCloseStore.triggered = true;
      }
    }
  }
})

// methods
function CallSaveEntry() {
  const entry: BibEntry = {
    Key: entryKey.value,
    Category: entryCategory.value,
    Fields: entryFields.value,
    Comment: '',
    CiteNumber: 0,
    Example: ''
  }
  SaveEntry(entry, initialEntry.value ? initialEntry.value.Key : '', props.projectName).then(data => {
    errorSuccessStore.handleResponse(data.Ok, t(i18nKeys.Success.SaveEntry), t(i18nKeys.Errors.ErrorSaving))

    if( data.Ok ) {
      GetProjectData(props.projectName).then(data => {
        projectDataStore.setData(data)
        emit('titleChange', entry.Key)
      });
      projectsStore.updateLastEditedOnProject(props.projectName, data.Data.LastModified);
    }
  })
}

function getSlotName(i: number) {
  return `${i}-1`
}

function NavBackAndRefresh() {
  deleteOpened.value = false;
  GetProjectData(props.projectName).then(data => {
    projectDataStore.setData(data);
    emit('navBack');
  })
}

// onload
if( initialEntry.value ) {
  entryKey.value = initialEntry.value.Key;
  entryCategory.value = initialEntry.value.Category;
  entryFields.value = JSON.parse(JSON.stringify(initialEntry.value.Fields));
}
</script>

<style scoped lang="scss">
.example {
  padding: 4px 8px;
  & .prefix {
    font-weight: bold;
  }
}
</style>
