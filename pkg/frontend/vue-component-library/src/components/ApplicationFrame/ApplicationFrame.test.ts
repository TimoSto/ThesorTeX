import {beforeEach, describe, expect, it} from "vitest";
import ApplicationFrame from "./ApplicationFrame.vue";
import {mount, VueWrapper} from "@vue/test-utils";
import CreateVuetifyMountingOptions from "@thesortex/vitest-vuetify";

type ApplicationFrameProps = InstanceType<typeof ApplicationFrame>["$props"];

describe("ApplicationFrame.vue", () => {
    beforeEach(() => {
        expect(ApplicationFrame).toBeTruthy();
    });
    describe("sidebar", () => {
        it("the sidebar button should be disabled if no sidebar is set", () => {
            const cmp = mountWithProps({hasSidebar: false});

            const btn = cmp.findAll(".v-app-bar-nav-icon")[0];

            expect(btn).toBeDefined();
            expect(btn.attributes("disabled")).toBe("");
        });
        it("the sidebar button should not be disabled if sidebar is set", () => {
            const cmp = mountWithProps({hasSidebar: true});

            const btn = cmp.findAll(".v-app-bar-nav-icon")[0];

            expect(btn).toBeDefined();
            expect(btn.attributes("disabled")).toBeUndefined();
        });
    });
    describe("main title", () => {
        it("display main title as h1", () => {
            const cmp = mountWithProps({
                mainTitle: "Foobar"
            });

            const headings = cmp.findAll(`[role="heading"][aria-level="1"]`);

            expect(headings).toHaveLength(1);
            expect(headings[0].text()).toEqual("Foobar");
        });
        //TODO: get title suffix from the slot via refs
    });
});

function mountWithProps(props: ApplicationFrameProps): VueWrapper<any> {
    return mount(ApplicationFrame, CreateVuetifyMountingOptions({
        props: props
    }));
}
