import {When} from "@cucumber/cucumber";
import {OurWorld} from "../../../types";

When("the deletion is confirmed", async function (this: OurWorld) {
    await this.page.locator(".v-overlay__content .v-card-actions button").nth(1).click();
});
