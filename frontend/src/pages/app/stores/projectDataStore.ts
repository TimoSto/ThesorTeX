import {defineStore} from "pinia";
import {ref} from "vue";
import {BibEntry} from "@/pages/app/api/projectData/entries/BibEntry";
import {BibCategory} from "@/pages/app/api/projectData/categories/BibCategory";
import {ProjectData} from "@/pages/app/api/projectData/ProjectData";

export const useProjectDataStore = defineStore('project-data', () => {
  const entries = ref([] as BibEntry[]);

  const categories = ref([] as BibCategory[]);

  function setData(data: ProjectData) {
    entries.value = data.Entries;
    categories.value = data.Categories;
  }

  return {entries, categories, setData}
})
