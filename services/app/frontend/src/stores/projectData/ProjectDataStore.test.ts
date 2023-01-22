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
    describe("actualize category", () => {
        it("edit existing", () => {
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

            store.actualizeCategory("testCate", {Name: "testCate", CitaviCategory: "test"} as Category);

            expect(store.categories[0]).toEqual({Name: "testCate", CitaviCategory: "test"});
        });
        it("rename existing", () => {
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

            store.actualizeCategory("testCate", {Name: "testCategory"} as Category);

            expect(store.categories[0]).toEqual({Name: "testCategory"});
        });
        it("add", () => {
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

            store.actualizeCategory("", {Name: "testCateg", CitaviCategory: "test"} as Category);

            expect(store.categories[0]).toEqual({Name: "testCate"});
            expect(store.categories[1]).toEqual({Name: "testCateg", CitaviCategory: "test"});
        });
    });
});
