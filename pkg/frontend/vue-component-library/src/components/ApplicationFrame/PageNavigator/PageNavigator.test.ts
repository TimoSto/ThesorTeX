import {beforeEach, describe, expect, it} from "vitest";
import PageNavigator from "./PageNavigator.vue";
import {mount} from "@vue/test-utils";
import {createTestingPinia} from "@pinia/testing";
import {useApplicationStateStore} from "../../../stores/ApplicationStateStore/ApplicationStateStore";

describe("PageNavigator.vue", () => {
    beforeEach(() => {
        expect(PageNavigator).toBeTruthy();
    });
    describe("displaying pages", () => {
        it("one page", () => {
            const wrapper = mount(PageNavigator, {
                global: {
                    plugins: [createTestingPinia()]
                }
            });

            const store = useApplicationStateStore();

            const pages = wrapper.findAll(".page");

            expect(pages).toHaveLength(1);
        });
        it("nav to second page", async () => {
            const wrapper = mount(PageNavigator, {
                global: {
                    plugins: [createTestingPinia({stubActions: false})]
                }
            });

            const store = useApplicationStateStore();

            let pages = wrapper.findAll(".page");

            expect(pages).toHaveLength(1);

            store.openPage("test 2");

            await wrapper.vm.$nextTick();

            pages = wrapper.findAll(".page");

            expect(pages).toHaveLength(2);
        });
        it("nav back from second page", async () => {
            const wrapper = mount(PageNavigator, {
                global: {
                    plugins: [createTestingPinia({stubActions: false})]
                }
            });

            const store = useApplicationStateStore();

            store.history.push("test");

            await wrapper.vm.$nextTick();

            let pages = wrapper.findAll(".page");

            expect(pages).toHaveLength(2);

            store.goBack(1);

            expect(pages).toHaveLength(2);

            await new Promise(r => {
                setTimeout(r, 750);
            });

            await wrapper.vm.$nextTick();

            pages = wrapper.findAll(".page");

            expect(pages).toHaveLength(1);
        });
    });
});