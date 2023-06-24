import {beforeEach, describe, expect, it} from "vitest";
import ApplicationFrame from "./ApplicationFrame.vue";
import {mount, VueWrapper} from "@vue/test-utils";
import CreateVuetifyMountingOptions from "@thesortex/vitest-vuetify";
import {Document} from "happy-dom";

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
            const cmp = mountWithProps({hasSidebar: true}, {sidebar: "<h2 id='test'>Foobar sidebar</h2>"});

            const btn = cmp.findAll(".v-app-bar-nav-icon")[0];

            expect(btn).toBeDefined();
            expect(btn.attributes("disabled")).toBeUndefined();

            const sidebarContent = cmp.find("#test");
            expect(sidebarContent).toBeDefined();
            expect(sidebarContent.text()).toEqual("Foobar sidebar");
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
    describe("link to documentation", () => {
        it("should not display a link if not set", () => {
            const cmp = mountWithProps({
                mainTitle: "Foobar"
            });

            const links = cmp.findAll(`a[href^="https://thesortex.com"]`);

            expect(links).toHaveLength(0);
        });
        it("should not display a link with an icon button if set", () => {
            const cmp = mountWithProps({
                mainTitle: "Foobar",
                documentationTarget: "test"
            });

            const links = cmp.findAll(`a[href^="https://thesortex.com"]`);

            expect(links).toHaveLength(1);

            expect(links[0].attributes("href")).toEqual("https://thesortex.com/#/tutorials?target=test");

            expect(links[0].find("button").classes()).toContain("v-btn--icon");

            expect(links[0].find("button").find("i").classes()).toContain("mdi-book-open-variant");
            expect(links[0].find("button").find("i").classes()).toContain("mdi");
            expect(links[0].find("button").find("i").classes()).toContain("v-icon");
        });
    });
    describe("integrating a11y dialog", () => {
        it("do not show if not set", () => {
            const cmp = mountWithProps({
                mainTitle: "Foobar"
            });

            const links = cmp.findAll(".mdi-human");

            expect(links).toHaveLength(0);
        });
        it("show if set", () => {
            const fakeDocument = new Document();

            const cmp = mountWithProps({
                mainTitle: "Foobar",
                showA11y: true,
                documentProp: fakeDocument
            });

            const links = cmp.findAll(".mdi-human");

            expect(links).toHaveLength(1);
        });
    });
    describe("config-btn", () => {
        it("do not show if not set", () => {
            const cmp = mountWithProps({
                mainTitle: "Foobar"
            });

            const links = cmp.findAll(".mdi-cog");

            expect(links).toHaveLength(0);
        });
        it("show if set", () => {
            const fakeDocument = new Document();

            const cmp = mountWithProps({
                mainTitle: "Foobar",
                hasConfig: true,
                documentProp: fakeDocument
            });

            const links = cmp.findAll(".mdi-cog");

            expect(links).toHaveLength(1);
        });
        //TODO: test show dialog
    });

    describe("i18n", () => {
        describe("sidebar button", () => {
            it("closed state", () => {
                const cmp = mountWithProps({hasSidebar: true});

                const btn = cmp.findAll(".v-app-bar-nav-icon")[0];

                expect(btn).toBeDefined();
                expect(btn.attributes("title")).toEqual("ApplicationFrame.OpenSidebar");
            });
            it("opened state", async () => {
                const cmp = mountWithProps({hasSidebar: true});

                const btn = cmp.findAll(".v-app-bar-nav-icon")[0];

                expect(btn).toBeDefined();

                await btn.trigger("click");

                expect(btn.attributes("title")).toEqual("ApplicationFrame.CloseSidebar");
            });
        });
        it("documentation link", () => {
            const cmp = mountWithProps({
                mainTitle: "Foobar",
                documentationTarget: "test"
            });

            const btn = cmp.findAll(`a[href^="https://thesortex.com"] button`)[0];

            expect(btn.attributes("title")).toEqual("ApplicationFrame.Documentation");
        });
    });
});

function mountWithProps(props: ApplicationFrameProps, slots?: any): VueWrapper<any> {
    return mount(ApplicationFrame, CreateVuetifyMountingOptions({
        props: props,
        slots: slots
    }));
}
