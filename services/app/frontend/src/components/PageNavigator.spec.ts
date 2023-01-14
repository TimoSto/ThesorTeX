import {mount} from "@vue/test-utils";
import PageNavigator from "./PageNavigator.vue";
import {describe, expect, test} from "vitest";

describe("PageNavigator.vue", () => {
    test("with empty props, no page should be shown",async()=>{
        expect(PageNavigator).toBeTruthy();

        const wrapper = mount(PageNavigator,{
            props:{},
        });
        expect(wrapper.find(".container").classes()).toEqual(["container"]);
        expect(wrapper.findAll(".page")).toHaveLength(0)
    })
    test("with instantSwitch set, class should be present",async()=>{
        expect(PageNavigator).toBeTruthy();

        const wrapper = mount(PageNavigator,{
            props:{
                instantSwitch: true
            },
        });
        expect(wrapper.find(".container").classes()).toEqual(["container", "disable-animations"]);
        expect(wrapper.findAll(".page")).toHaveLength(0)
    })
    test("with one page set, only one page should be shown",async()=>{
        expect(PageNavigator).toBeTruthy();

        const wrapper = mount(PageNavigator,{
            props:{
                pages: 1,
            },
        });
        expect(wrapper.find(".container").classes()).toEqual(["container"]);
        expect(wrapper.findAll(".page")).toHaveLength(1)
    })
    test("with two pages set, only one page should be shown but two should exist",async()=>{
        expect(PageNavigator).toBeTruthy();

        const wrapper = mount(PageNavigator,{
            props:{
                pages: 2,
            },
        });
        expect(wrapper.find(".container").classes()).toEqual(["container"]);
        expect(wrapper.findAll(".page")).toHaveLength(2)
        expect(wrapper.findAll(".page")[0].classes()).toEqual(["page"]);
        expect(wrapper.findAll(".page")[1].classes()).toEqual(["page", "opened"]);
    })
    test("when navigating back, the second last page should have the class 'opened'",async()=>{
        expect(PageNavigator).toBeTruthy();

        const wrapper = mount(PageNavigator,{
            props:{
                pages: 3,
            },
        });

        await wrapper.setProps({
            pages: 3,
            navigatingBack: true
        });

        expect(wrapper.find(".container").classes()).toEqual(["container"]);
        expect(wrapper.findAll(".page")).toHaveLength(3)
        expect(wrapper.findAll(".page")[0].classes()).toEqual(["page"]);
        expect(wrapper.findAll(".page")[1].classes()).toEqual(["page", "opened"]);
        expect(wrapper.findAll(".page")[2].classes()).toEqual(["page"]);

        await wrapper.setProps({
            pages: 2,
            navigatingBack: false
        });

        expect(wrapper.find(".container").classes()).toEqual(["container"]);
        expect(wrapper.findAll(".page")).toHaveLength(2)
        expect(wrapper.findAll(".page")[0].classes()).toEqual(["page"]);
        expect(wrapper.findAll(".page")[1].classes()).toEqual(["page", "opened"]);
    })
})
