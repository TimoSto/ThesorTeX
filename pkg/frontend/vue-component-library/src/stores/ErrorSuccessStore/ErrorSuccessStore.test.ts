import {describe, it, beforeEach, expect} from "vitest";
import {createPinia, setActivePinia} from "pinia";
import {_useErrorSuccessStore} from "./ErrorSuccessStore";

describe("ErrorSuccessStore", () => {
    beforeEach(() => {
        setActivePinia(createPinia());
    })
    it("setting success", () => {
        const store = _useErrorSuccessStore();

        store.setMessage(true, "My success");

        expect(store.valid).toBe(true);
        expect(store.message).toEqual("My success");
    })
    it("setting error", () => {
        const store = _useErrorSuccessStore();

        store.setMessage(false, "My error");

        expect(store.valid).toBe(false);
        expect(store.message).toEqual("My error");
    })
    it("clearing data", () => {
        const store = _useErrorSuccessStore();

        store.message = "tests";
        store.valid = false;

        store.clear();

        expect(store.valid).toBe(true);
        expect(store.message).toEqual("");
    })
})