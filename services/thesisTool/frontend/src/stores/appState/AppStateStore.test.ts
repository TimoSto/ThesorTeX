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
    describe("goBack", () => {
        it("sidebar was closed", () => {
            const store = useAppStateStore();
            store.history.push("test");
            store.goBack();
            store.finishGoBack();
            expect(store.history.length).toEqual(1);
            expect(store.sidebarOpen).toBeFalsy();
        });
        it("sidebar was open on page 3", () => {
            const store = useAppStateStore();
            store.history.push("test", "test2");
            store.sidebarOpen = true;
            store.goBack();
            store.finishGoBack();
            expect(store.history.length).toEqual(2);
            expect(store.sidebarOpen).toBeTruthy();
        });
        it("sidebar was open on page 2", () => {
            const store = useAppStateStore();
            store.history.push("test");
            store.sidebarOpen = true;
            store.goBack();
            store.finishGoBack();
            expect(store.history.length).toEqual(1);
            expect(store.sidebarOpen).toBeFalsy();
        });
        it("project was open on page 2", () => {
            const store = useAppStateStore();
            store.history.push("test");
            store.currentProject = "tesst";
            store.goBack();
            store.finishGoBack();
            expect(store.history.length).toEqual(1);
            expect(store.sidebarOpen).toBeFalsy();
            expect(store.currentProject).toEqual("");
        });
        it("project was open on page 3", () => {
            const store = useAppStateStore();
            store.history.push("test", "test2");
            store.currentProject = "tesst";
            store.goBack();
            store.finishGoBack();
            expect(store.history.length).toEqual(2);
            expect(store.sidebarOpen).toBeFalsy();
            expect(store.currentProject).toEqual("tesst");
        });
    });
    describe("currentItem", () => {
        it("should reset the currentItem", () => {
            const store = useAppStateStore();
            store.history.push("test", "test2");
            store.setItem("tesst");
            expect(store.currentItem).toEqual("tesst");
            store.goBack();
            store.finishGoBack();
            expect(store.history.length).toEqual(2);
            expect(store.sidebarOpen).toBeFalsy();
            expect(store.currentItem).toEqual("");
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
    describe("interrupt because of unsaved changes", () => {
        describe("goBack", () => {
            it("decline", () => {
                const store = useAppStateStore();
                store.history.push("test", "test2");
                store.unsavedChanges = true;
                store.goBack();
                expect(store.unsavedDialogTriggered).toBe(true);
                store.resolveCallback(false);
                expect(store.unsavedDialogTriggered).toBe(false);
                expect(store.history).toEqual(["main", "test", "test2"]);
            });
            it("accept", () => {
                const store = useAppStateStore();
                store.history.push("test", "test2");
                store.unsavedChanges = true;
                store.goBack();
                expect(store.unsavedDialogTriggered).toBe(true);
                store.resolveCallback(true);
                expect(store.unsavedDialogTriggered).toBe(false);
            });
        });
    });
});
