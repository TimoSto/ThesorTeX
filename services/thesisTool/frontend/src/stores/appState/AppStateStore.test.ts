import {beforeEach, describe, expect, it} from "vitest";
import {createPinia, setActivePinia} from "pinia";
import {useAppStateStore} from "./AppStateStore";

describe("AppStateStore", () => {
    beforeEach(() => {
        setActivePinia(createPinia());
    });
    describe("navToPage", () => {
        it("sidebar was closed", () => {
            const store = useAppStateStore();
            expect(store.history).toEqual(["main"]);
            store.navToPage("test");
            expect(store.history.length).toEqual(2);
            expect(store.sidebarOpen).toBeFalsy();
        });
        it("sidebar was open", () => {
            const store = useAppStateStore();
            store.navToPage("test");
            store.sidebarOpen = true;
            store.navToPage("test2");
            expect(store.history.length).toEqual(3);
            expect(store.sidebarOpen).toBeTruthy();
        });
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
