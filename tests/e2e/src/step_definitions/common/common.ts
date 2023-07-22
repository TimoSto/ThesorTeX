import {Then, When} from "@cucumber/cucumber";
import {expect} from "@playwright/test";
import {OurWorld} from "../../types";
import waitForAnimations from "../../helpers/waitForAnimations";

When("the deletion is confirmed", async function (this: OurWorld) {
    await this.page.locator(".v-overlay__content .v-card-actions button").nth(1).click();
});

Then("the save button in the editor is disabled", async function (this: OurWorld) {
    await expect(this.page.locator("#page-3 header").locator("button").nth(0)).toBeDisabled();
});

Then("the save button in the editor is enabled", async function (this: OurWorld) {
    await expect(this.page.locator("#page-3 header").locator("button").nth(0)).toBeEnabled();
});

When("the save button in the editor is clicked", async function (this: OurWorld) {
    await this.page.locator("#page-3 header").locator("button").nth(0).click();
});

When("the delete button in the editor is clicked", async function (this: OurWorld) {
    await this.page.locator("#page-3 header").locator("button").nth(1).click();
});

Then("the editor-page is closed", async function (this: OurWorld) {
    await waitForAnimations(this.page, [".container", "#page-2", "#page-2", ".v-overlay-container"]);

    expect(await this.page.locator("#page-3").count()).toEqual(0);
});

Then("the user is prompted that there are unsaved changes", async function (this: OurWorld) {
    await waitForAnimations(this.page, [".container", ".v-overlay-container"]);

    expect(await this.page.locator(".v-overlay__content .v-card-title").textContent()).toEqual("There are unsaved changes");
});

When("the close is confirmed", async function (this: OurWorld) {
    await this.page.locator(".v-overlay__content .v-card-actions button").nth(1).click();

    await waitForAnimations(this.page, [".container", "#page-2", ".v-overlay-container"]);
});

When("the close is aborted", async function (this: OurWorld) {
    await this.page.locator(".v-overlay__content .v-card-actions button").nth(0).click();
});
