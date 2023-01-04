import {defineStore} from "pinia";
import {ref} from "vue";
import {BibEntry} from "../api/projectData/entries/BibEntry";
import {BibCategory} from "../api/projectData/categories/BibCategory";
import {ProjectData} from "../api/projectData/ProjectData";
import GenerateModelForFields from "../api/projectData/categories/GenerateModel";
import GenerateEntry from "../api/projectData/entries/GenerateEntry";

export const useProjectDataStore = defineStore('project-data', () => {
  const entries = ref([] as BibEntry[]);

  const categories = ref([] as BibCategory[]);

  function setData(data: ProjectData) {
    data.Entries.forEach((e: BibEntry) => {
      e.Example = GenerateEntry(e, data.Categories);
    });
    entries.value = data.Entries;
    data.Categories.forEach((c: BibCategory) => {
      c.Example = GenerateModelForFields(c.Fields);
    })
    categories.value = data.Categories;
  }

  return {entries, categories, setData}
})
