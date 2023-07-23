import {Then, When} from "@cucumber/cucumber";
import {expect} from "@playwright/test";
import {OurWorld} from "../../types";
import waitForAnimations from "../../helpers/waitForAnimations";

When("a new entry is created", async function (this: OurWorld) {
    await this.page.locator("#page-2").locator(".v-expansion-panel").nth(0).locator("thead button").click();

    await waitForAnimations(this.page, [".container", "#page-2", ".v-overlay-container"]);
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

When("{string} is entered into the input at index {int}", async function (this: OurWorld, text: string, i: number) {
    await this.page.locator("#page-3").locator(".v-expansion-panel").nth(1).locator(".v-input").nth(i).locator("input").clear();
    await this.page.locator("#page-3").locator(".v-expansion-panel").nth(1).locator(".v-input").nth(i).locator("input").type(text);
});

When("the entry {string} is opened", async function (this: OurWorld, key: string) {
    await this.page.locator("#page-2").locator("tr", {has: this.page.locator("td", {hasText: key})}).click();
});

Then("the user is asked to confirm the deletion of the entry", async function (this: OurWorld) {
    await waitForAnimations(this.page, [".container", ".v-overlay-container"]);

    expect(await this.page.locator(".v-overlay__content .v-card-title").textContent()).toEqual("Delete entry");
});

Then("the displayed entry for {string} is {string}", async function (this: OurWorld, key: string, entry: string) {
    const gotEntry = await this.page.locator("#page-2").locator(".v-expansion-panel").first().locator("tbody tr", {has: this.page.locator("td", {hasText: key})}).locator("td").nth(2).textContent();

    expect(gotEntry).toEqual(entry);
});
