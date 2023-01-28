import {When} from "@cucumber/cucumber";
import {OurWorld} from "../../../types";
import waitForAnimations from "../../helpers/waitForAnimations";

When("a new entry is created", async function (this: OurWorld) {
    await this.page.locator("#page-2").locator(".v-expansion-panel").nth(0).locator("thead button").click();

    await waitForAnimations(this.page, ["#page-2"]);
});
