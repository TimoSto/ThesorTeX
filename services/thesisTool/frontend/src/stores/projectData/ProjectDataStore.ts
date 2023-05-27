import {defineStore} from "pinia";
import {Entry} from "../../domain/entry/Entry";
import {Category} from "../../domain/category/Category";

export const useProjectDataStore = defineStore({
    id: "project-data",

    state: () => ({
        entries: [] as Entry[],
        categories: [] as Category[]
    }),

    actions: {
        setProjectData(entries: Entry[], categories: Category[]) {
            this.entries = entries;
            categories.forEach(c => {
                if (!c.CitaviFilter) {
                    c.CitaviFilter = [];
                }
                c.BibFields.forEach(f => {
                    if (!f.CitaviMapping) {
                        f.CitaviMapping = [];
                    }
                });
                c.CiteFields.forEach(f => {
                    if (!f.CitaviMapping) {
                        f.CitaviMapping = [];
                    }
                });
            });
            this.categories = categories;
        },
        actualizeCategory(name: string, category: Category) {
            category = JSON.parse(JSON.stringify(category));
            const i = this.categories.map(c => c.Name).indexOf(name);
            if (i >= 0) {
                this.categories[i] = category;
            } else {
                this.categories.push(category);
                this.categories.sort((c1, c2) => {
                    return c1.Name > c2.Name ? 1 : -1;
                });
            }
        },
        removeCategory(name: string) {
            const i = this.categories.map(c => c.Name).indexOf(name);
            if (i >= 0) {
                this.categories.splice(i, 1);
            }
        },
        actualizeEntry(key: string, entry: Entry) {
            entry = JSON.parse(JSON.stringify(entry));
            const i = this.entries.map(c => c.Key).indexOf(key);
            if (i >= 0) {
                this.entries[i] = entry;
            } else {
                this.entries.push(entry);
                this.entries.sort((c1, c2) => {
                    return c1.Key > c2.Key ? 1 : -1;
                });
            }
        },
        removeEntry(key: string) {
            const i = this.entries.map(c => c.Key).indexOf(key);
            if (i >= 0) {
                this.entries.splice(i, 1);
            }
        }
    },

    getters: {
        categoryIsUsed: (state) => {
            return (category: string): string[] =>
                state.entries.filter(e => e.Category === category).map(e => e.Key);
        },
    }
});
