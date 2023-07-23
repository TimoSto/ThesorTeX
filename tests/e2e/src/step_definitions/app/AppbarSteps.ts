import {Then, When} from "@cucumber/cucumber";
import {OurWorld} from "../../types";
import {expect} from "@playwright/test";
import waitForAnimations from "../../helpers/waitForAnimations";

Then("the title of the app is {string}", async function (this: OurWorld, title: string) {
    await waitForAnimations(this.page, [".container", ".v-overlay-container"]);

    expect(await this.page.locator(".v-app-bar .v-toolbar-title__placeholder").textContent()).toEqual(title);
});

Then("the title of the main area is {string}", async function (this: OurWorld, title: string) {
    expect(await this.page.locator(".page.opened .v-toolbar-title__placeholder").last().textContent()).toEqual(title);
});

When("the back button is clicked", async function (this: OurWorld) {
    await this.page.locator(".v-app-bar").locator("button").nth(1).click();

    await waitForAnimations(this.page, [".container", "#page-2", ".v-overlay-container"]);
});

When("the config button is clicked", async function (this: OurWorld) {
    await this.page.locator(".v-app-bar").locator(".mdi-cog >> xpath=../..",).click();

    await waitForAnimations(this.page, [".container", ".v-overlay-container"]);
});

Then("the config dialog is shown", async function (this: OurWorld) {
    await waitForAnimations(this.page, [".container", ".v-overlay-container"]);

    expect(await this.page.locator(".v-overlay__content .v-card-title").textContent()).toEqual("Configuration");
});

When("the close button of the dialog is clicked", async function (this: OurWorld) {
    await this.page.locator(".v-overlay__content .v-btn", {hasText: "Close"}).click();
});

Then("the save button in the dialog is disabled", async function (this: OurWorld) {
    await expect(this.page.locator(".v-overlay__content .v-btn", {hasText: "Save"})).toBeDisabled();
});
