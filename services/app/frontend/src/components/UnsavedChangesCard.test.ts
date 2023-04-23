import {beforeEach, describe, expect, it} from "vitest";
import UnsavedChangesCard from "./UnsavedChangesCard.vue";
import {mount} from "@vue/test-utils";
import CreateVuetifyMountingOptions from "@thesortex/vitest-vuetify";
import {english} from "../i18n/english";

describe("UnsavedChangesCard.vue", () => {
    beforeEach(() => {
        expect(UnsavedChangesCard).toBeTruthy();
    });
    describe("emit resolved event", () => {
        it("resolve true", async () => {
            const cmp = mount(UnsavedChangesCard, CreateVuetifyMountingOptions({
                props: {
                    valid: true,
                    message: ""
                },
            }, english));

            cmp.vm.$el.querySelectorAll("button")[1].click();

            await cmp.vm.$nextTick();

            expect(cmp.emitted().resolve).toBeTruthy();
            expect(cmp.emitted().resolve.length).toBe(1);
            expect(cmp.emitted().resolve[0]).toEqual([true]);
        });
        it("resolve false", async () => {
            const cmp = mount(UnsavedChangesCard, CreateVuetifyMountingOptions({
                props: {
                    valid: true,
                    message: ""
                },
            }, english));

            cmp.vm.$el.querySelectorAll("button")[0].click();

            await cmp.vm.$nextTick();

            expect(cmp.emitted().resolve).toBeTruthy();
            expect(cmp.emitted().resolve.length).toBe(1);
            expect(cmp.emitted().resolve[0]).toEqual([false]);
        });
    });
});