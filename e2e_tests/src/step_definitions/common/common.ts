import {Then, When} from "@cucumber/cucumber";
import {expect} from "@playwright/test";
import {OurWorld} from "../../../types";

When("the deletion is confirmed", async function (this: OurWorld) {
    await this.page.locator(".v-overlay__content .v-card-actions button").nth(1).click();
});

Then("the save button in the editor is disabled", async function (this: OurWorld) {
    await expect(this.page.locator("#page-3 header").locator("button").nth(0)).toBeDisabled();
});
