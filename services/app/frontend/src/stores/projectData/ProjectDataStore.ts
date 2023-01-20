import {defineStore} from "pinia";
import {Entry} from "../../domain/entry/Entry";

export const useProjectDataStore = defineStore({
    id: "project-data",

    state: () => ({
        entries: [] as Entry[]
    }),

    actions: {
        setProjectData(entries: Entry[]) {
            this.entries = entries;
        }
    }
})
