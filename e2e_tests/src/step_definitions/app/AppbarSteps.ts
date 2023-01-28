import {Then} from "@cucumber/cucumber";
import {OurWorld} from "../../../types";
import {expect} from "@playwright/test";

Then("the title of the app is {string}", async function (this: OurWorld, title: string) {
    expect(await this.page.locator(".v-app-bar .v-toolbar-title__placeholder").textContent()).toEqual(title);
});

Then("the title of the main area is {string}", async function (this: OurWorld, title: string) {
    expect(await this.page.locator(".page--container .v-toolbar-title__placeholder").textContent()).toEqual(title);
});
