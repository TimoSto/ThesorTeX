import {describe, test, beforeEach, expect} from "vitest";
import {createPinia, setActivePinia} from "pinia";
import {useAppStateStore} from "./AppStateStore";

describe("AppStateStore", () => {
    beforeEach(() => {
        setActivePinia(createPinia())
    });
    describe("navToPage", () => {
        test("sidebar was closed", () => {
            const store = useAppStateStore();
            expect(store.history).toEqual(["main"]);
            store.navToPage("test");
            expect(store.history.length).toEqual(2);
            expect(store.currentPage).toEqual("test");
            expect(store.sidebarOpen).toBeFalsy();
        });
        test("sidebar was open", () => {
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
        test("sidebar was closed", () => {
            const store = useAppStateStore();
            store.history.push("test")
            store.goBack();
            expect(store.history.length).toEqual(1);
            expect(store.currentPage).toEqual("main");
            expect(store.sidebarOpen).toBeFalsy();
        });
        test("sidebar was open on page 3", () => {
            const store = useAppStateStore();
            store.history.push("test", "test2")
            store.sidebarOpen = true;
            store.goBack();
            expect(store.history.length).toEqual(2);
            expect(store.currentPage).toEqual("test");
            expect(store.sidebarOpen).toBeTruthy();
        });
        test("sidebar was open on page 2", () => {
            const store = useAppStateStore();
            store.history.push("test")
            store.sidebarOpen = true;
            store.goBack();
            expect(store.history.length).toEqual(1);
            expect(store.currentPage).toEqual("main");
            expect(store.sidebarOpen).toBeFalsy();
        });
    })
})