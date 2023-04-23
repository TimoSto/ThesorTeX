import {beforeEach, describe, expect, it} from "vitest";
import ToolbarAndContent from "./ToolbarAndContent.vue";
import {mount} from "@vue/test-utils";
import CreateVuetifyMountingOptions from "@thesortex/vitest-vuetify";

describe("ToolbarAndContent.vue", () => {
    beforeEach(() => {
        expect(ToolbarAndContent).toBeTruthy();
    });

    describe("Elevate on scroll", () => {
        it("should not have elevation initially", async () => {
            Object.defineProperty(window, "innerHeight", {
                writable: true,
                configurable: true,
                value: 500
            });
            const cmp = mount(ToolbarAndContent, CreateVuetifyMountingOptions({
                props: {
                    valid: true,
                    message: ""
                },
                slots: {
                    content: "<div style=\"height: 2000px\"></div>"
                }
            }));

            expect(cmp.vm.elevation).toEqual(0);

            await cmp.vm.$nextTick();

            expect(cmp.emitted().scroll).toBeFalsy();
            expect(cmp.emitted().noScroll).toBeFalsy();
        });
        it("should have elevation when scrolled", async () => {
            Object.defineProperty(window, "innerHeight", {
                writable: true,
                configurable: true,
                value: 500
            });
            const cmp = mount(ToolbarAndContent, CreateVuetifyMountingOptions({
                props: {
                    valid: true,
                    message: ""
                },
                slots: {
                    content: "<div style=\"height: 2000px\">hallo</div>"
                }
            }));

            cmp.vm.$el.querySelector(".content").scrollTop = 300;
            cmp.vm.$el.querySelector(".content").dispatchEvent(new CustomEvent("scroll"));

            expect(cmp.vm.elevation).toEqual(1);

            await cmp.vm.$nextTick();

            expect(cmp.emitted().scroll.length).toEqual(1);
            expect(cmp.emitted().noScroll).toBeFalsy();
        });
        it("should not have elevation when scrolled back to top", async () => {
            Object.defineProperty(window, "innerHeight", {
                writable: true,
                configurable: true,
                value: 500
            });
            const cmp = mount(ToolbarAndContent, CreateVuetifyMountingOptions({
                props: {
                    valid: true,
                    message: ""
                },
                slots: {
                    content: "<div style=\"height: 2000px\">hallo</div>"
                }
            }));

            cmp.vm.$el.querySelector(".content").scrollTop = 300;
            cmp.vm.$el.querySelector(".content").dispatchEvent(new CustomEvent("scroll"));

            expect(cmp.vm.elevation).toEqual(1);

            await cmp.vm.$nextTick();

            cmp.vm.$el.querySelector(".content").scrollTop = 0;
            cmp.vm.$el.querySelector(".content").dispatchEvent(new CustomEvent("scroll"));

            expect(cmp.vm.elevation).toEqual(0);

            await cmp.vm.$nextTick();

            expect(cmp.emitted().scroll.length).toEqual(1);
            expect(cmp.emitted().noScroll.length).toEqual(1);
        });
    });
});
