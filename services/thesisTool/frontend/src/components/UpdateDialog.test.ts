import {beforeEach, describe, expect, it} from "vitest";
import UpdateDialog from "./UpdateDialog.vue";
import {mount} from "@vue/test-utils";
import {english} from "../i18n/english";
import CreateVuetifyMountingOptions from "@thesortex/vitest-vuetify";

describe("Updatedialog.vue", () => {
    beforeEach(() => {
        expect(UpdateDialog).toBeTruthy();
    });
    describe("opened state", () => {
        it("should be opened with no ignored version", () => {
            const cmp = mount(UpdateDialog, CreateVuetifyMountingOptions({
                props: {
                    version: "1.2.3"
                },
            }, english));

            expect(cmp.vm.display).toBe(true);
        });
        it("should be opened with not ignored version", () => {
            localStorage.setItem("thesortex.ignoreUpdates", JSON.stringify(["1.0.0"]));
            const cmp = mount(UpdateDialog, CreateVuetifyMountingOptions({
                props: {
                    version: "1.2.3"
                },
            }, english));

            expect(cmp.vm.display).toBe(true);
        });
        it("should not be opened with ignored version", () => {
            localStorage.setItem("thesortex.ignoreUpdates", JSON.stringify(["1.2.3"]));
            const cmp = mount(UpdateDialog, CreateVuetifyMountingOptions({
                props: {
                    version: "1.2.3"
                },
            }, english));

            expect(cmp.vm.display).toBe(false);
        });
        it("should not be opened with empty version", () => {
            const cmp = mount(UpdateDialog, CreateVuetifyMountingOptions({
                props: {
                    version: ""
                },
            }, english));

            expect(cmp.vm.display).toBeFalsy();
        });
    });
});
