import {defineStore} from "pinia";
import {Entry} from "../../domain/entry/Entry";
import {Category} from "../../domain/category/category";

export const useProjectDataStore = defineStore({
    id: "project-data",

    state: () => ({
        entries: [] as Entry[],
        categories: [] as Category[]
    }),

    actions: {
        setProjectData(entries: Entry[], categories: Category[]) {
            this.entries = entries;
            this.categories = categories;
        }
    }
})
