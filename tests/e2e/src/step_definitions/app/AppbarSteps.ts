import {Then, When} from "@cucumber/cucumber";
import {OurWorld} from "../../types";
import {expect} from "@playwright/test";
import waitForAnimations from "../../helpers/waitForAnimations";

Then("the title of the app is {string}", async function (this: OurWorld, title: string) {
    await waitForAnimations(this.page, [".pages", ".v-overlay-container"]);

    expect(await this.page.locator(".v-app-bar .v-toolbar-title__placeholder").textContent()).toEqual(title);
});

Then("the title of the main area is {string}", async function (this: OurWorld, title: string) {
    expect(await this.page.locator(".page--container .v-toolbar-title__placeholder").last().textContent()).toEqual(title);
});

When("the back button is clicked", async function (this: OurWorld) {
    await this.page.locator(".v-app-bar").locator("button").nth(1).click();

    await waitForAnimations(this.page, [".pages", "#page-2", ".v-overlay-container"]);
});
