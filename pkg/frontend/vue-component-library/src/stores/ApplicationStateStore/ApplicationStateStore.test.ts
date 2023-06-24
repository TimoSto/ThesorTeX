import {beforeEach, describe, expect, it} from "vitest";
import {createPinia, setActivePinia} from "pinia";
import {NavState, useApplicationStateStore} from "./ApplicationStateStore";

describe("ApplicationStateStore", () => {
    beforeEach(() => {
        setActivePinia(createPinia());
    });
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
    it("nav back one page", async () => {
        const store = useApplicationStateStore();

        store.$state.history.push("test");

        store.goBack(1);

        expect(store.$state.navState).toEqual(NavState.Back);

        await new Promise<void>((r) => {
            setTimeout(() => {
                r();
            }, 500);
        });
        expect(store.$state.navState).toEqual(NavState.None);
        expect(store.$state.history).toEqual(["main"]);
    });
});
