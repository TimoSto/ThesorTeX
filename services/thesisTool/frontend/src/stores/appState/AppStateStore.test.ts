import {beforeEach, describe, expect, it} from "vitest";
import {createPinia, setActivePinia} from "pinia";
import {useAppStateStore} from "./AppStateStore";

describe("AppStateStore", () => {
    beforeEach(() => {
        setActivePinia(createPinia());
    });
    describe("currentItem", () => {
        it("should reset the currentItem", () => {
            const store = useAppStateStore();
            store.history.push("test", "test2");
            store.setItem("tesst");
            expect(store.currentItem).toEqual("tesst");
        });
    });
    describe("switchProject", () => {
        it("switch from project page", () => {
            const store = useAppStateStore();
            store.history.push("test");
            store.switchToProject("test3");
            expect(store.history).toEqual(["main", "test"]);
            expect(store.currentProject).toEqual("test3");
        });
        it("switch from editor page", () => {
            const store = useAppStateStore();
            store.history.push("test", "test2");
            store.switchToProject("test3");
            expect(store.history).toEqual(["main", "test"]);
            expect(store.currentProject).toEqual("test3");
        });
    });
});
