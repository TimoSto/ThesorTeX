import {beforeEach, describe, expect, it} from "vitest";
import {createPinia, setActivePinia} from "pinia";
import {useProjectsListStore} from "./ProjectsListStore";

describe("ProjectsListStore", () => {
    beforeEach(() => {
        setActivePinia(createPinia())
    });
    describe("addProject()", () => {
        it("should add on empty", () => {
            const store = useProjectsListStore();
            const p = {Name: "test", Created: "now", LastEdited: "now", NumberOfEntries: "2"};
            store.addProject(p);
            expect(store.projects).toEqual([p]);
        })
        it("should sort correctly", () => {
            const store = useProjectsListStore();
            store.projects = [
                {Name: "a", Created: "now", LastEdited: "now", NumberOfEntries: "2"},
                {Name: "zen", Created: "now", LastEdited: "now", NumberOfEntries: "2"}
            ]
            const p = {Name: "test", Created: "now", LastEdited: "now", NumberOfEntries: "2"};
            store.addProject(p);
            expect(store.projects).toEqual([
                {Name: "a", Created: "now", LastEdited: "now", NumberOfEntries: "2"},
                p,
                {Name: "zen", Created: "now", LastEdited: "now", NumberOfEntries: "2"}
            ]);
        });
    });
    describe("removeProject", () => {
        it("should delete last", () => {
            const store = useProjectsListStore();
            store.projects.push({Name: "test", Created: "now", LastEdited: "now", NumberOfEntries: "2"});
            store.removeProject("test");
            expect(store.projects).toEqual([]);
        })
    });
});
