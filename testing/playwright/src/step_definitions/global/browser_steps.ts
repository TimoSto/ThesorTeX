import { Given, When, Then } from "@cucumber/cucumber";
import assert from "assert";
import { OurWorld } from "../../support/OurWorld";
import {expect} from "@playwright/test";
// Using a cucumber expression
Given("the page {string} was opened", async function (this: OurWorld, url: string) {
    // Use the page instance from the World instance to navigate
    await this.page.goto(url);
});
// Using a regular expression
When("I click the tree dot button", async function (this: OurWorld) {
    // Scroll to the link...
    const element = this.page.locator(".mdi-dots-vertical >> xpath=../..");
    await expect(element).toBeVisible();
    await element.click();
});
Then("I expect to see a menu in the top right corner", async function (this: OurWorld) {
    const menuElement = this.page.locator(".v-overlay__content:visible")
    await expect(menuElement).toBeVisible();
    const configEntry = menuElement.locator("text=Settings")
    await expect(configEntry).toBeVisible();

});
// textContent includes whitespace, so use this method to trim
// See https://stackoverflow.com/a/42921059
const trimExcessWhiteSpace = (s: string) =>
    s.replace(/[\n\r]+|[\s]{2,}/g, " ").trim();
