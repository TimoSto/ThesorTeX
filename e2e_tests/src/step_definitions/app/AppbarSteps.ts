import {Then} from "@cucumber/cucumber";
import {OurWorld} from "../../../types";
import {expect} from "@playwright/test";
import waitForAnimations from "../../helpers/waitForAnimations";

Then("the title of the app is {string}", async function (this: OurWorld, title: string) {
    await waitForAnimations(this.page);

    expect(await this.page.locator(".v-app-bar .v-toolbar-title__placeholder").textContent()).toEqual(title);
});

Then("the title of the main area is {string}", async function (this: OurWorld, title: string) {
    expect(await this.page.locator(".page--container .v-toolbar-title__placeholder").last().textContent()).toEqual(title);
});
