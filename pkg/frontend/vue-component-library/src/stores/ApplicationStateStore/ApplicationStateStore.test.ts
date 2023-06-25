import {beforeEach, describe, expect, it} from "vitest";
import {createPinia, setActivePinia} from "pinia";
import {NavState, useApplicationStateStore} from "./ApplicationStateStore";

describe("ApplicationStateStore", () => {
    beforeEach(() => {
        setActivePinia(createPinia());
    });
    describe("openPage", () => {
        it("nav to new page", async () => {
            const store = useApplicationStateStore();

            store.openPage("test");

            expect(store.$state.navState).toEqual(NavState.Forth);

            await new Promise<void>((r) => {
                setTimeout(() => {
                    r();
                }, 500);
            });
            expect(store.$state.navState).toEqual(NavState.None);
            expect(store.$state.history).toEqual(["main", "test"]);
        });
    });
    describe("goBack", async () => {
        it("nav back one page", async () => {
            const store = useApplicationStateStore();

            store.$state.history.push("test");

            store.goBack(1);

            expect(store.$state.navState).toEqual(NavState.Back);

            await new Promise<void>((r) => {
                setTimeout(() => {
                    r();
                }, 750);
            });
            expect(store.$state.navState).toEqual(NavState.None);
            expect(store.$state.history).toEqual(["main"]);
        });
        describe("unsaved changes", () => {
            it("nav back one page and dont accept", async () => {
                const store = useApplicationStateStore();

                store.$state.history.push("test");

                store.$state.unsavedChanges = true;

                store.goBack(1);

                expect(store.$state.navState).toEqual(NavState.None);

                await new Promise<void>((r) => {
                    setTimeout(() => {
                        r();
                    }, 750);
                });
                expect(store.$state.navState).toEqual(NavState.None);
                expect(store.$state.history).toEqual(["main", "test"]);
            });
            it("nav back one page and accept", async () => {
                const store = useApplicationStateStore();

                store.$state.history.push("test");

                store.$state.unsavedChanges = true;

                store.goBack(1);

                expect(store.$state.navState).toEqual(NavState.None);

                await new Promise<void>((r) => {
                    setTimeout(() => {
                        r();
                    }, 750);
                });
                expect(store.$state.navState).toEqual(NavState.None);
                expect(store.$state.history).toEqual(["main", "test"]);

                store.unsavedDialogCallback();

                expect(store.$state.navState).toEqual(NavState.Back);
                expect(store.$state.history).toEqual(["main", "test"]);

                await new Promise<void>((r) => {
                    setTimeout(() => {
                        r();
                    }, 750);
                });
                expect(store.$state.navState).toEqual(NavState.None);
                expect(store.$state.history).toEqual(["main"]);
            });
        });
    });
});
