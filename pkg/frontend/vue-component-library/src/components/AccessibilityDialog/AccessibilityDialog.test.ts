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
    });
    it("should have closed state", () => {
        const cmp = mount(AccessibilityDialog, CreateVuetifyMountingOptions({
            props: {
                keydownTarget: fakeDocument
            }
        }));

        expect(cmp.vm.opened).toBe(false);
        expect(document.body.querySelector(".v-card-title")).toBeFalsy();
    });
    it("should open on 2x CTRL", () => {
        const cmp = mount(AccessibilityDialog, CreateVuetifyMountingOptions({
            props: {
                keydownTarget: fakeDocument
            }
        }));

        fakeDocument.dispatchEvent(new FakeEvent("keypress", {key: "Ctrl"}));
        fakeDocument.dispatchEvent(new FakeEvent("keypress", {key: "Ctrl"}));
        expect(cmp.vm.opened).toBe(true);
    });
    it("should not open on 1x CTRL", () => {
        const cmp = mount(AccessibilityDialog, CreateVuetifyMountingOptions({
            props: {
                keydownTarget: fakeDocument
            }
        }));

        fakeDocument.dispatchEvent(new FakeEvent("keypress", {key: "Ctrl"}));
        expect(cmp.vm.opened).toBe(false);
    });
    it("should not open on 2x CTRL too far apart", async () => {
        const cmp = mount(AccessibilityDialog, CreateVuetifyMountingOptions({
            props: {
                keydownTarget: fakeDocument
            }
        }));

        fakeDocument.dispatchEvent(new FakeEvent("keypress", {key: "Ctrl"}));

        await new Promise(resolve => {
            setTimeout(resolve, 400);
        });

        fakeDocument.dispatchEvent(new FakeEvent("keypress", {key: "Ctrl"}));

        expect(cmp.vm.opened).toBe(false);
    });
    it("should open on 2x CTRL 200ms apart", async () => {
        const cmp = mount(AccessibilityDialog, CreateVuetifyMountingOptions({
            props: {
                keydownTarget: fakeDocument
            }
        }));

        fakeDocument.dispatchEvent(new FakeEvent("keypress", {key: "Ctrl"}));

        await new Promise(resolve =>
            setTimeout(resolve, 200)
        );

        fakeDocument.dispatchEvent(new FakeEvent("keypress", {key: "Ctrl"}));

        expect(cmp.vm.opened).toBe(true);
    });
});
