import {beforeEach, describe, expect, it} from "vitest";
import AccessibilityDialog from "./AccessibilityDialog.vue";
import {mount} from "@vue/test-utils";
import CreateVuetifyMountingOptions from "@thesortex/vitest-vuetify";
import {Document, KeyboardEvent as FakeEvent} from "happy-dom";

let fakeDocument: Document;
describe("AccessibilityDialog.vue", () => {
    beforeEach(() => {
        expect(AccessibilityDialog).toBeTruthy();

        fakeDocument = new Document();
        localStorage.clear();
    });
    describe("opened state", () => {
        it("should have closed state", () => {
            const cmp = mount(AccessibilityDialog, CreateVuetifyMountingOptions({
                props: {
                    keydownTarget: fakeDocument
                }
            }));

            expect(cmp.vm.opened).toBe(false);
            expect(document.body.querySelector(".v-card-title")).toBeFalsy();
        });
        it("should open on 2x Control", () => {
            const cmp = mount(AccessibilityDialog, CreateVuetifyMountingOptions({
                props: {
                    keydownTarget: fakeDocument
                }
            }));

            fakeDocument.dispatchEvent(new FakeEvent("keyup", {key: "Control"}));
            fakeDocument.dispatchEvent(new FakeEvent("keyup", {key: "Control"}));
            expect(cmp.vm.opened).toBe(true);
        });
        it("should not open on 1x Control", () => {
            const cmp = mount(AccessibilityDialog, CreateVuetifyMountingOptions({
                props: {
                    keydownTarget: fakeDocument
                }
            }));

            fakeDocument.dispatchEvent(new FakeEvent("keyup", {key: "Control"}));
            expect(cmp.vm.opened).toBe(false);
        });
        it("should not open on 2x Control too far apart", async () => {
            const cmp = mount(AccessibilityDialog, CreateVuetifyMountingOptions({
                props: {
                    keydownTarget: fakeDocument
                }
            }));

            fakeDocument.dispatchEvent(new FakeEvent("keyup", {key: "Control"}));

            await new Promise(resolve => {
                setTimeout(resolve, 400);
            });

            fakeDocument.dispatchEvent(new FakeEvent("keyup", {key: "Control"}));

            expect(cmp.vm.opened).toBe(false);
        });
        it("should open on 2x Control 200ms apart", async () => {
            const cmp = mount(AccessibilityDialog, CreateVuetifyMountingOptions({
                props: {
                    keydownTarget: fakeDocument
                }
            }));

            fakeDocument.dispatchEvent(new FakeEvent("keyup", {key: "Control"}));

            await new Promise(resolve =>
                setTimeout(resolve, 200)
            );

            fakeDocument.dispatchEvent(new FakeEvent("keyup", {key: "Control"}));

            expect(cmp.vm.opened).toBe(true);
        });
    });
    describe("activating focus borders", () => {
        it("activating focus borders", () => {
            const cmp = mount(AccessibilityDialog, CreateVuetifyMountingOptions({
                props: {
                    keydownTarget: fakeDocument
                }
            }));

            cmp.vm.focusBordersActivated = true;

            cmp.vm.$nextTick(() => {
                expect(fakeDocument.getElementById("focus-borders-set")).not.toBeNull();
            });
        });
        it("focus borders initially activated", () => {
            localStorage.setItem("thesortexFocusBorders", "true");

            const cmp = mount(AccessibilityDialog, CreateVuetifyMountingOptions({
                props: {
                    keydownTarget: fakeDocument
                }
            }));

            cmp.vm.$nextTick(() => {
                expect(fakeDocument.getElementById("focus-borders-set")).not.toBeNull();
            });
        });
        it("activate and deactivating focus borders", () => {
            const cmp = mount(AccessibilityDialog, CreateVuetifyMountingOptions({
                props: {
                    keydownTarget: fakeDocument
                }
            }));

            cmp.vm.focusBordersActivated = true;

            cmp.vm.$nextTick(() => {
                expect(fakeDocument.getElementById("focus-borders-set")).not.toBeNull();
                cmp.vm.focusBordersActivated = false;
                cmp.vm.$nextTick(() => {
                    expect(fakeDocument.getElementById("focus-borders-set")).toBeNull();
                });
            });
        });
        it("focus borders initially activated, then deactivated", () => {
            localStorage.setItem("thesortexFocusBorders", "true");

            const cmp = mount(AccessibilityDialog, CreateVuetifyMountingOptions({
                props: {
                    keydownTarget: fakeDocument
                }
            }));

            cmp.vm.$nextTick(() => {
                expect(fakeDocument.getElementById("focus-borders-set")).not.toBeNull();
                cmp.vm.focusBordersActivated = false;
                cmp.vm.$nextTick(() => {
                    expect(fakeDocument.getElementById("focus-borders-set")).toBeNull();
                });
            });
        });
    });
});
