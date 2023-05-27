import {beforeEach, describe, expect, it} from "vitest";
import {createPinia, setActivePinia, storeToRefs} from "pinia";
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
                Name: "testCate",
                BibFields: [],
                CiteFields: [],
                CitaviCategory: "",
                CitaviFilter: []
            }
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
                    Name: "testCate",
                    BibFields: [],
                    CiteFields: [],
                    CitaviCategory: "",
                    CitaviFilter: []
                }
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
                    Name: "testCate",
                    BibFields: [],
                    CiteFields: [],
                    CitaviCategory: "",
                    CitaviFilter: []
                }
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
                    Name: "testCate",
                    BibFields: [],
                    CiteFields: [],
                    CitaviCategory: "",
                    CitaviFilter: []
                }
            ];
            store.setProjectData(entries, categories);

            store.actualizeCategory("", {
                Name: "testCateg",
                BibFields: [],
                CiteFields: [],
                CitaviCategory: "test",
                CitaviFilter: []
            });

            expect(store.categories[0]).toEqual({
                Name: "testCate",
                BibFields: [],
                CiteFields: [],
                CitaviCategory: "",
                CitaviFilter: []
            });
            expect(store.categories[1]).toEqual({
                Name: "testCateg",
                BibFields: [],
                CiteFields: [],
                CitaviCategory: "test",
                CitaviFilter: []
            });
        });
    });
    describe("rm category", () => {
        it("should remove", () => {
            const store = useProjectDataStore();

            const entries = [
                {
                    Key: "test",
                } as Entry
            ];

            const categories = [
                {
                    Name: "testCate",
                    BibFields: [],
                    CiteFields: [],
                    CitaviCategory: "",
                    CitaviFilter: []
                }
            ];
            store.setProjectData(entries, categories);

            store.removeCategory("testCate");

            expect(store.categories).toEqual([]);
        });
    });
    describe("categoryIsUsed", () => {
        it("should be false on unused", () => {
            const store = useProjectDataStore();

            const entries = [
                {
                    Key: "test",
                    Category: "c2"
                } as Entry
            ];

            const categories = [
                {
                    Name: "testCate",
                    BibFields: [],
                    CiteFields: [],
                    CitaviCategory: "",
                    CitaviFilter: []
                }
            ];
            store.setProjectData(entries, categories);

            const {categoryIsUsed} = storeToRefs(store);

            expect(categoryIsUsed.value("testCate")).toBe(false);
        });
        it("should be true on used", () => {
            const store = useProjectDataStore();

            const entries = [
                {
                    Key: "test",
                    Category: "testCate"
                } as Entry
            ];

            const categories = [
                {
                    Name: "testCate",
                    BibFields: [],
                    CiteFields: [],
                    CitaviCategory: "",
                    CitaviFilter: []
                }
            ];
            store.setProjectData(entries, categories);

            const {categoryIsUsed} = storeToRefs(store);

            expect(categoryIsUsed.value("testCate")).toBe(true);
        });
    });
    describe("actualize entry", () => {
        it("edit existing", () => {
            const store = useProjectDataStore();

            const entries = [
                {
                    Key: "test",
                } as Entry
            ];

            const categories = [
                {
                    Name: "testCate",
                    BibFields: [],
                    CiteFields: [],
                    CitaviCategory: "",
                    CitaviFilter: []
                }
            ];
            store.setProjectData(entries, categories);

            store.actualizeEntry("test", {Key: "test", Category: "testCate"} as Entry);

            expect(store.entries[0]).toEqual({Key: "test", Category: "testCate"});
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
                    Name: "testCate",
                    BibFields: [],
                    CiteFields: [],
                    CitaviCategory: "",
                    CitaviFilter: []
                }
            ];
            store.setProjectData(entries, categories);

            store.actualizeEntry("test", {Key: "test2"} as Entry);

            expect(store.entries[0]).toEqual({Key: "test2"});
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
                    Name: "testCate",
                    BibFields: [],
                    CiteFields: [],
                    CitaviCategory: "",
                    CitaviFilter: []
                }
            ];
            store.setProjectData(entries, categories);

            store.actualizeEntry("", {
                Key: "t2",
                Category: "ttt",
                Fields: []
            });

            expect(store.entries[1]).toEqual({
                Key: "test",
            });
            expect(store.entries[0]).toEqual({
                Key: "t2",
                Category: "ttt",
                Fields: []
            });
        });
    });
    describe("rm entry", () => {
        it("should remove", () => {
            const store = useProjectDataStore();

            const entries = [
                {
                    Key: "test",
                } as Entry
            ];

            const categories = [
                {
                    Name: "testCate",
                    BibFields: [],
                    CiteFields: [],
                    CitaviCategory: "",
                    CitaviFilter: []
                }
            ];
            store.setProjectData(entries, categories);

            store.removeEntry("test");

            expect(store.entries).toEqual([]);
        });
    });
});
