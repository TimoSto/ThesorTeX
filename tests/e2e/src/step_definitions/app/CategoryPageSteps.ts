import {Then, When} from "@cucumber/cucumber";
import {OurWorld} from "../../types";
import waitForAnimations from "../../helpers/waitForAnimations";
import {expect} from "@playwright/test";

When("a new category is created", async function (this: OurWorld) {
    await this.page.locator("#page-2").locator(".v-expansion-panel").nth(1).locator("thead button").click();

    await waitForAnimations(this.page, ["#page-2"]);
});

When("{string} is entered as name", async function (this: OurWorld, name: string) {
    await this.page.locator("#page-3").locator("input").nth(0).type(name);
});

When("a bib field is added", async function (this: OurWorld) {
    await this.page.locator("#page-3").locator(".v-expansion-panel").nth(1).locator("thead button").click();
});

When("{string} is entered as the name of the field at index {int}", async function (this: OurWorld, name: string, n: number) {
    await this.page.locator("#page-3").locator(".v-expansion-panel").nth(1).locator("tbody tr").nth(n).locator("td").nth(0).locator("input").type(name);
});

When("{string} is entered as the prefix of the field at index {int}", async function (this: OurWorld, name: string, n: number) {
    await this.page.locator("#page-3").locator(".v-expansion-panel").nth(1).locator("tbody tr").nth(n).locator("td").nth(1).locator("input").type(name);
});

When("{string} is entered as the suffix of the field at index {int}", async function (this: OurWorld, name: string, n: number) {
    await this.page.locator("#page-3").locator(".v-expansion-panel").nth(1).locator("tbody tr").nth(n).locator("td").nth(2).locator("input").type(name);
});

Then("the displayed example entry for {string} is {string}", async function (this: OurWorld, name: string, entry: string) {
    const gotEntry = await this.page.locator("#page-2").locator(".v-expansion-panel").nth(1).locator("tbody tr", {has: this.page.locator("td", {hasText: name})}).locator("td").nth(1).textContent();

    expect(gotEntry).toEqual(entry);
});

When("the category {string} is opened", async function (this: OurWorld, name: string) {
    await this.page.locator("#page-2").locator("tr", {has: this.page.locator("td", {hasText: name})}).click();
});

Then("the user is asked to confirm the deletion of the category", async function (this: OurWorld) {
    await waitForAnimations(this.page);

    expect(await this.page.locator(".v-overlay__content .v-card-title").textContent()).toEqual("Delete category");
});
