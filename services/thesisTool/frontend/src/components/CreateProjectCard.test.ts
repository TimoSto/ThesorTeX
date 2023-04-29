import {describe, expect, it} from "vitest";
import CreateProjectCard from "./CreateProjectCard.vue";
import {mount} from "@vue/test-utils";
import CreateVuetifyMountingOptions from "@thesortex/vitest-vuetify";
import {english} from "../i18n/english";

describe("CreateProjectCard.vue", () => {
    describe("state of the save btn", () => {
        it("the save btn should be disabled, when no name was entered yet", () => {
            expect(CreateProjectCard).toBeTruthy();

            const wrapper = mount(CreateProjectCard, CreateVuetifyMountingOptions({
                props: {
                    projects: []
                }
            }, english));

            const saveBtn = wrapper.findAll("button")[1];
            expect(saveBtn.element.disabled).toBe(true);
        });
        it("the save btn should be disabled, when the rules are violated", async () => {
            expect(CreateProjectCard).toBeTruthy();

            const wrapper = mount(CreateProjectCard, CreateVuetifyMountingOptions({
                props: {
                    projects: []
                }
            }, english));

            await wrapper.find("input").setValue("hallo du");

            const saveBtn = wrapper.findAll("button")[1];
            expect(saveBtn.element.disabled).toBe(true);
        });
        it("the save btn should be disabled, when project name already exists", async () => {
            expect(CreateProjectCard).toBeTruthy();

            const wrapper = mount(CreateProjectCard, CreateVuetifyMountingOptions({
                props: {
                    projects: ["test"]
                }
            }, english));

            await wrapper.find("input").setValue("test");

            const saveBtn = wrapper.findAll("button")[1];
            expect(saveBtn.element.disabled).toBe(true);
        });
        it("the save btn should not be disabled, when project name is valid", async () => {
            expect(CreateProjectCard).toBeTruthy();

            const wrapper = mount(CreateProjectCard, CreateVuetifyMountingOptions({
                props: {
                    projects: ["test"]
                }
            }, english));

            await wrapper.find("input").setValue("test2");

            const saveBtn = wrapper.findAll("button")[1];
            expect(saveBtn.element.disabled).toBe(false);
        });
    });
    describe("emits", () => {
        it("should emit a close event, when abort is clicked", async () => {
            expect(CreateProjectCard).toBeTruthy();

            const wrapper = mount(CreateProjectCard, CreateVuetifyMountingOptions({
                props: {
                    projects: ["test"]
                }
            }, english));

            await wrapper.find("button").trigger("click");

            expect(wrapper.emitted().close).toBeTruthy();
        });
        it("should emit a confirm event, when save is clicked", async () => {
            expect(CreateProjectCard).toBeTruthy();

            const wrapper = mount(CreateProjectCard, CreateVuetifyMountingOptions({
                props: {
                    projects: ["test"]
                }
            }, english));

            await wrapper.find("input").setValue("test2");

            await wrapper.findAll("button")[1].trigger("click");

            expect(wrapper.emitted().confirm).toBeTruthy();
        });
    });
});
