import {defineStore} from "pinia";
import {ref} from "vue";
import {BibEntry} from "@/pages/app/api/projectData/entries/BibEntry";
import {BibCategory} from "@/pages/app/api/projectData/categories/BibCategory";
import {ProjectData} from "@/pages/app/api/projectData/ProjectData";
import GenerateModelForCategory from "@/pages/app/api/projectData/categories/GenerateModel";

export const useProjectDataStore = defineStore('project-data', () => {
  const entries = ref([] as BibEntry[]);

  const categories = ref([] as BibCategory[]);

  function setData(data: ProjectData) {
    entries.value = data.Entries;
    data.Categories.forEach((c: BibCategory) => {
      c.Example = GenerateModelForCategory(c);
    })
    categories.value = data.Categories;
  }

  return {entries, categories, setData}
})