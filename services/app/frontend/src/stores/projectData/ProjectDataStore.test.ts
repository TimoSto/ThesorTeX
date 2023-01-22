import {beforeEach, describe, expect, it} from "vitest";
import {createPinia, setActivePinia} from "pinia";
import {useProjectDataStore} from "./ProjectDataStore";
import {Entry} from "../../domain/entry/Entry";
import {Category} from "../../domain/category/Category";

describe("ProjectDataStore", () => {
    beforeEach(() => {
        setActivePinia(createPinia());
    });
    it("setting data", () => {
        const store = useProjectDataStore();

        const entries = [
            {
                Key: "test",
            } as Entry
        ];

        const categories = [
            {
                Name: "testCate"
            } as Category
        ];
        store.setProjectData(entries, categories);

        expect(store.categories).toEqual(categories);
        expect(store.entries).toEqual(entries);
    });
});
