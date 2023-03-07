import {mount} from "@vue/test-utils";
import PageNavigator from "./PageNavigator.vue";
import {describe, expect, it} from "vitest";
import CreateVuetifyMountingOptions from "@thesortex/vitest-vuetify";

describe("PageNavigator.vue", () => {

    it("with empty props, no page should be shown", async () => {
        expect(PageNavigator).toBeTruthy();

        const wrapper = mount(PageNavigator, CreateVuetifyMountingOptions({
            props: {
                pages: 0,
                instantSwitch: false
            },
        }));
        expect(wrapper.find(".container").classes()).toEqual(["container"]);
        expect(wrapper.findAll(".page")).toHaveLength(0);
    });
    it("with instantSwitch set, class should be present", async () => {
        expect(PageNavigator).toBeTruthy();

        const wrapper = mount(PageNavigator, CreateVuetifyMountingOptions({
            props: {
                instantSwitch: true,
                pages: 0
            },
        }));
        expect(wrapper.find(".container").classes()).toEqual(["container", "disable-animations"]);
        expect(wrapper.findAll(".page")).toHaveLength(0);
    });
    it("with one page set, only one page should be shown", async () => {
        expect(PageNavigator).toBeTruthy();

        const wrapper = mount(PageNavigator, CreateVuetifyMountingOptions({
            props: {
                pages: 1,
                instantSwitch: false
            },
        }));
        expect(wrapper.find(".container").classes()).toEqual(["container"]);
        expect(wrapper.findAll(".page")).toHaveLength(1);
    });
    it("with two pages set, only one page should be shown but two should exist", async () => {
        expect(PageNavigator).toBeTruthy();

        const wrapper = mount(PageNavigator, CreateVuetifyMountingOptions({
            props: {
                pages: 2,
                instantSwitch: false
            },
        }));
        expect(wrapper.find(".container").classes()).toEqual(["container"]);
        expect(wrapper.findAll(".page")).toHaveLength(2);
        expect(wrapper.findAll(".page")[0].classes()).toEqual(["page"]);
        expect(wrapper.findAll(".page")[1].classes()).toEqual(["page", "opened"]);
    });
    it("when navigating back, the second last page should have the class 'opened'", async () => {
        expect(PageNavigator).toBeTruthy();

        const wrapper = mount(PageNavigator, CreateVuetifyMountingOptions({
            props: {
                pages: 3,
                instantSwitch: false
            },
        }));

        await wrapper.setProps({
            pages: 3,
            navigatingBack: true,
            instantSwitch: false
        });

        expect(wrapper.find(".container").classes()).toEqual(["container", "nav-back"]);
        expect(wrapper.findAll(".page")).toHaveLength(4);
        expect(wrapper.findAll(".page")[0].classes()).toEqual(["page"]);
        expect(wrapper.findAll(".page")[1].classes()).toEqual(["page", "opened"]);
        expect(wrapper.findAll(".page")[2].classes()).toEqual(["page"]);

        await wrapper.setProps({
            pages: 2,
            navigatingBack: false,
            instantSwitch: false
        });

        expect(wrapper.find(".container").classes()).toEqual(["container", "nav-back"]);
        expect(wrapper.findAll(".page")).toHaveLength(3);
        expect(wrapper.findAll(".page")[0].classes()).toEqual(["page"]);
        expect(wrapper.findAll(".page")[1].classes()).toEqual(["page", "opened"]);
    });
});
