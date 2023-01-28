import {Then, When} from "@cucumber/cucumber";
import {expect} from "@playwright/test";
import {OurWorld} from "../../../types";
import waitForAnimations from "../../helpers/waitForAnimations";

When("a new entry is created", async function (this: OurWorld) {
    await this.page.locator("#page-2").locator(".v-expansion-panel").nth(0).locator("thead button").click();

    await waitForAnimations(this.page, ["#page-2"]);
});

When("{string} is entered as key", async function (this: OurWorld, key: string) {
    await this.page.locator("#page-3").locator("input").nth(0).type(key);
});

When("the first category is selected", async function (this: OurWorld) {
    await this.page.locator("#page-3").locator(".v-input").nth(1).click();
    await this.page.locator(".v-overlay").locator(".v-list-item").nth(0).click();
});

Then("the fields have a length greater than 0", async function (this: OurWorld) {
    expect(await this.page.locator("#page-3").locator(".v-expansion-panel").nth(1).locator(".v-input").count()).toBeGreaterThan(0);
});
