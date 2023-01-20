import {beforeEach, describe, expect, it} from "vitest";
import {createPinia, setActivePinia} from "pinia";
import {useAppStateStore} from "./AppStateStore";

describe("AppStateStore", () => {
    beforeEach(() => {
        setActivePinia(createPinia())
    });
    describe("navToPage", () => {
        it("sidebar was closed", () => {
            const store = useAppStateStore();
            expect(store.history).toEqual(["main"]);
            store.navToPage("test");
            expect(store.history.length).toEqual(2);
            expect(store.currentPage).toEqual("test");
            expect(store.sidebarOpen).toBeFalsy();
        });
        it("sidebar was open", () => {
            const store = useAppStateStore();
            store.navToPage("test");
            store.sidebarOpen = true;
            store.navToPage("test2");
            expect(store.history.length).toEqual(3);
            expect(store.currentPage).toEqual("test2");
            expect(store.sidebarOpen).toBeTruthy();
        });
    });
    describe("goBack", () => {
        it("sidebar was closed", () => {
            const store = useAppStateStore();
            store.history.push("test")
            store.goBack();
            store.finishGoBack();
            expect(store.history.length).toEqual(1);
            expect(store.currentPage).toEqual("main");
            expect(store.sidebarOpen).toBeFalsy();
        });
        it("sidebar was open on page 3", () => {
            const store = useAppStateStore();
            store.history.push("test", "test2")
            store.sidebarOpen = true;
            store.goBack();
            store.finishGoBack();
            expect(store.history.length).toEqual(2);
            expect(store.currentPage).toEqual("test");
            expect(store.sidebarOpen).toBeTruthy();
        });
        it("sidebar was open on page 2", () => {
            const store = useAppStateStore();
            store.history.push("test")
            store.sidebarOpen = true;
            store.goBack();
            store.finishGoBack();
            expect(store.history.length).toEqual(1);
            expect(store.currentPage).toEqual("main");
            expect(store.sidebarOpen).toBeFalsy();
        });
        it("project was open on page 2", () => {
            const store = useAppStateStore();
            store.history.push("test")
            store.currentProject = "tesst";
            store.goBack();
            store.finishGoBack();
            expect(store.history.length).toEqual(1);
            expect(store.currentPage).toEqual("main");
            expect(store.sidebarOpen).toBeFalsy();
            expect(store.currentProject).toEqual("");
        });
        it("project was open on page 3", () => {
            const store = useAppStateStore();
            store.history.push("test", "test2")
            store.currentProject = "tesst";
            store.goBack();
            store.finishGoBack();
            expect(store.history.length).toEqual(2);
            expect(store.currentPage).toEqual("test");
            expect(store.sidebarOpen).toBeFalsy();
            expect(store.currentProject).toEqual("tesst");
        });
    })
})
