import { Given, When, Then } from "@cucumber/cucumber";
import assert from "assert";
import { OurWorld } from "../../support/OurWorld";
import {expect} from "@playwright/test";

Given("the page {string} was opened", async function (this: OurWorld, url: string) {
    // Use the page instance from the World instance to navigate
    await this.page.goto(url);
});

When("I press the tab key {int} times", async function (this: OurWorld, n: number) {
    while (n > 0) {
        await this.page.keyboard.press("Tab");
        n--;
    }
});

When("I press the enter key", async function (this: OurWorld) {
    await this.page.keyboard.press("Enter");
});

When("I click the tree dot button", async function (this: OurWorld) {
    // Scroll to the link...
    const element = this.page.locator(".mdi-dots-vertical >> xpath=../..");
    await expect(element).toBeVisible();
    await element.click();
});



Then("I expect to see a menu in the top right corner", async function (this: OurWorld) {
    const menuElement = this.page.locator(".v-overlay__content:visible")
    await expect(menuElement).toBeVisible();
    const box = await menuElement.boundingBox();
    await expect(box).not.toBeNull()
    if( box ) {
        const diffTop = Math.abs(box.x - 1300 + box.width)
        await expect(diffTop).toBeLessThan(15)
        await expect(diffTop).toBeGreaterThan(10)
        await expect(box.y).toBeLessThan(60)
        await expect(box.y).toBeGreaterThan(50)
    }
});

Then("I expect the three dot menu to contain the following entries {string}", async function (this: OurWorld, list: string) {
    const menuElement = this.page.locator(".v-overlay__content:visible")
    const entries = list.split(",")
    for (let e of entries) {
        const configEntry = menuElement.locator(`text=${e.trim()}`)
        await expect(configEntry).toBeVisible();
    }
})
// textContent includes whitespace, so use this method to trim
// See https://stackoverflow.com/a/42921059
const trimExcessWhiteSpace = (s: string) =>
    s.replace(/[\n\r]+|[\s]{2,}/g, " ").trim();
