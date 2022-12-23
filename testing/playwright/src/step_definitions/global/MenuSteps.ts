import { When, Then } from "@cucumber/cucumber";
import { OurWorld } from "../../support/OurWorld";
import {expect} from "@playwright/test";

When("I click the tree dot button", async function (this: OurWorld) {
    // Scroll to the link...
    const element = this.page.locator(".mdi-dots-vertical >> xpath=../..");
    await expect(element).toBeVisible();
    await element.click();
});

Then("I expect the first menu element to be focussed", async function (this: OurWorld) {
    const focusedElement = await this.page.evaluate(() => document.activeElement);
    expect(focusedElement).not.toBeNull();
    if( focusedElement ) {
        expect(focusedElement.classList).toContain('v-list-item')
    }
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
