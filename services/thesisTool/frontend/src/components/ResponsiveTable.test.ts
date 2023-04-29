import {mount} from "@vue/test-utils";
import {describe, expect, it} from "vitest";
import ResponsiveTable, {SizeClasses} from "./ResponsiveTable.vue";
import CreateVuetifyMountingOptions from "@thesortex/vitest-vuetify";

describe("ResponsiveTable.vue", () => {

    it("with only headlines", async () => {
        expect(ResponsiveTable).toBeTruthy();

        const wrapper = mount(ResponsiveTable, CreateVuetifyMountingOptions({
            props: {
                disableRipple: false,
                headers: [
                    {
                        size: SizeClasses.TwentyFivePercent,
                        content: "H1"
                    },
                    {
                        size: SizeClasses.TwentyFivePercent,
                        content: "H2"
                    },
                    {
                        size: SizeClasses.TwentyFivePercent,
                        content: "H3"
                    },
                    {
                        size: SizeClasses.TenPercent,
                        content: "H4"
                    },
                    {
                        size: SizeClasses.IconBtn,
                        slot: true
                    },
                ],
                rows: []
            },
            slots: {
                "h-4": "Here is a slot"
            }
        }));
        expect(wrapper.findAll("th").length).toEqual(5);
        expect(wrapper.findAll("th")[0].classes()).toContain("size--25-percent");
        expect(wrapper.findAll("th")[1].classes()).toContain("size--25-percent");
        expect(wrapper.findAll("th")[2].classes()).toContain("size--25-percent");
        expect(wrapper.findAll("th")[3].classes()).toContain("size--15-percent");
        expect(wrapper.findAll("th")[4].classes()).toContain("size--icon-btn");
        expect(wrapper.findAll("th")[4].html()).toContain("Here is a slot");
    });
    it("with two rows", async () => {
        expect(ResponsiveTable).toBeTruthy();

        const wrapper = mount(ResponsiveTable, CreateVuetifyMountingOptions({
            props: {
                disableRipple: false,
                headers: [
                    {
                        size: SizeClasses.TwentyFivePercent,
                        content: "H1"
                    },
                    {
                        size: SizeClasses.TwentyFivePercent,
                        content: "H2"
                    },
                    {
                        size: SizeClasses.TwentyFivePercent,
                        content: "H3"
                    },
                    {
                        size: SizeClasses.TenPercent,
                        content: "H4"
                    },
                    {
                        size: SizeClasses.IconBtn,
                        slot: true
                    },
                ],
                rows: [
                    [
                        {content: "1-0"},
                        {content: "1-1"},
                        {content: "1-2"},
                        {content: "1-3"},
                        {slot: true},
                    ],
                    [
                        {content: "2-0"},
                        {content: "2-1"},
                        {content: "2-2"},
                        {content: "2-3"},
                        {slot: true},
                    ],
                ]
            },
            slots: {
                "h-4": "Here is a slot",
                "0-4": "Here is slot 1-4",
                "1-4": "Here is slot 2-4",
            }
        }));
        expect(wrapper.findAll("th").length).toEqual(5);
        expect(wrapper.findAll("th")[0].classes()).toContain("size--25-percent");
        expect(wrapper.findAll("th")[1].classes()).toContain("size--25-percent");
        expect(wrapper.findAll("th")[2].classes()).toContain("size--25-percent");
        expect(wrapper.findAll("th")[3].classes()).toContain("size--15-percent");
        expect(wrapper.findAll("th")[4].classes()).toContain("size--icon-btn");
        expect(wrapper.findAll("th")[4].html()).toContain("Here is a slot");
        expect(wrapper.find("tbody").findAll("tr").length).toEqual(2);
        expect(wrapper.find("tbody").findAll("tr")[0].findAll("td")[0].text()).toEqual("1-0");
        expect(wrapper.find("tbody").findAll("tr")[0].findAll("td")[1].text()).toEqual("1-1");
        expect(wrapper.find("tbody").findAll("tr")[0].findAll("td")[2].text()).toEqual("1-2");
        expect(wrapper.find("tbody").findAll("tr")[0].findAll("td")[3].text()).toEqual("1-3");
        expect(wrapper.find("tbody").findAll("tr")[0].findAll("td")[4].html()).toContain("Here is slot 1-4");
        expect(wrapper.find("tbody").findAll("tr")[1].findAll("td")[0].text()).toEqual("2-0");
        expect(wrapper.find("tbody").findAll("tr")[1].findAll("td")[1].text()).toEqual("2-1");
        expect(wrapper.find("tbody").findAll("tr")[1].findAll("td")[2].text()).toEqual("2-2");
        expect(wrapper.find("tbody").findAll("tr")[1].findAll("td")[3].text()).toEqual("2-3");
        expect(wrapper.find("tbody").findAll("tr")[1].findAll("td")[4].html()).toContain("Here is slot 2-4");
    });
});
