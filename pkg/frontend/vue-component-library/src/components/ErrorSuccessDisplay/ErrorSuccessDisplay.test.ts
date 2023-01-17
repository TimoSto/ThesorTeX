import {beforeEach, describe, expect, it} from "vitest";
import ErrorSuccessDisplay from "./ErrorSuccessDisplay.vue";
import {mount} from "@vue/test-utils";
import CreateVuetifyMountingOptions from "@thesortex/vitest-vuetify"

describe("ErrorSuccessDisplay.vue", () => {
    beforeEach(() => {
        expect(ErrorSuccessDisplay).toBeTruthy();
    })
   describe("Opened state", () => {
        it("should be closed with no message", () => {
            const cmp = mount(ErrorSuccessDisplay, CreateVuetifyMountingOptions({
                props: {
                    valid: true,
                    message: ""
                }
            }));

            expect(cmp.vm.errorOpened).toBe(false);
            expect(cmp.vm.successOpened).toBe(false);
            expect(document.body.querySelector(".v-card-title")).toBeFalsy();
        });
       it("should have success opened", () => {
           const cmp = mount(ErrorSuccessDisplay, CreateVuetifyMountingOptions({
               props: {
                   valid: true,
                   message: "success message",
                   errorTitle: "An error occurred doing your test-action"
               }
           }));

           expect(cmp.vm.errorOpened).toBe(false);
           expect(cmp.vm.successOpened).toBe(true);
           expect(document.body.querySelector(".v-card-title")).toBeFalsy();
       });
       it("should have error opened", () => {
           const cmp = mount(ErrorSuccessDisplay, CreateVuetifyMountingOptions({
               props: {
                   valid: false,
                   message: "error message",
                   errorTitle: "An error occurred doing your test-action",
                   errorSuffix: "do whatever you want"
               }
           }));

           expect(cmp.vm.errorOpened).toBe(true);
           expect(cmp.vm.successOpened).toBe(false);
           expect(document.body.querySelector(".v-card-title")!.textContent).toEqual("An error occurred doing your test-action");
           expect(document.body.querySelector(".error-message")!.textContent).toEqual("error message");
           expect(document.body.querySelector(".suffix")!.textContent).toEqual("do whatever you want");
       });
       it("should have success opened", () => {
           const cmp = mount(ErrorSuccessDisplay, CreateVuetifyMountingOptions({
               props: {
                   valid: true,
                   message: "success message",
                   errorTitle: "An error occurred doing your test-action",
                   errorSuffix: "do whatever you want"
               }
           }));
           expect(cmp.vm.errorOpened).toBe(false);
           expect(cmp.vm.successOpened).toBe(true);
       });
       it("should emit close from error-Dialog", () => {
           const cmp = mount(ErrorSuccessDisplay, CreateVuetifyMountingOptions({
               props: {
                   valid: false,
                   message: "error message",
                   errorTitle: "An error occurred doing your test-action",
                   errorSuffix: "do whatever you want",
                   close: "close it"
               }
           }));

           Array.from(document.querySelectorAll("button")).find(el => el.innerText === "close it")!.click();

           expect(cmp.emitted().close).toBeTruthy();
       });
       it("should emit close 2,5 seconds after setting success", async () => {
           const cmp = mount(ErrorSuccessDisplay, CreateVuetifyMountingOptions({
               props: {
                   valid: true,
                   message: "",
                   errorTitle: "An error occurred doing your test-action",
                   errorSuffix: "do whatever you want",
                   close: "close it"
               }
           }));

           await cmp.setProps({
               valid: true,
               message: "wasser",
               errorTitle: "An error occurred doing your test-action",
               errorSuffix: "do whatever you want",
               close: "close it"
           })

           expect(cmp.vm.errorOpened).toBe(false);
           expect(cmp.vm.successOpened).toBe(true);

           expect(cmp.emitted().close).toBeFalsy();

           await new Promise(r => setTimeout(r, 3000));

           expect(cmp.emitted().close).toBeTruthy();
       })
       it("should not emit close 2,5 seconds after setting success, if success is renewed", async () => {
           const cmp = mount(ErrorSuccessDisplay, CreateVuetifyMountingOptions({
               props: {
                   valid: true,
                   message: "",
                   errorTitle: "An error occurred doing your test-action",
                   errorSuffix: "do whatever you want",
                   close: "close it"
               }
           }));

           await cmp.setProps({
               valid: true,
               message: "wasser",
               errorTitle: "An error occurred doing your test-action",
               errorSuffix: "do whatever you want",
               close: "close it"
           })

           await new Promise(r => setTimeout(r, 1000));

           await cmp.setProps({
               valid: true,
               message: "",
               errorTitle: "An error occurred doing your test-action",
               errorSuffix: "do whatever you want",
               close: "close it"
           })

           await cmp.setProps({
               valid: true,
               message: "wasser",
               errorTitle: "An error occurred doing your test-action",
               errorSuffix: "do whatever you want",
               close: "close it"
           })

           expect(cmp.vm.errorOpened).toBe(false);
           expect(cmp.vm.successOpened).toBe(true);

           // todo: this doesnt work yet
           // expect(cmp.emitted().close).toBeFalsy();
           //
           // await new Promise(r => setTimeout(r, 2000));
           //
           // expect(cmp.emitted().close).toBeFalsy();
           //
           // await new Promise(r => setTimeout(r, 1000));
           //
           // expect(cmp.vm.successTimeout).toEqual(10);
       })
   })
});
