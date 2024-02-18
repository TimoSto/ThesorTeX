import {Then, When} from "@cucumber/cucumber";
import {OurWorld} from "../../types";
import {expect} from "@playwright/test";
import AxeBuilder from "@axe-core/playwright";
import {ScanDefaultWCAG} from "@thesortex/playwright-a11y-helper";

Then("the wcag guidelines are met", async function (this: OurWorld) {
    const page = this.page;
    const te = await page.title();
    console.log("title", te);
    const accessibilityScanResults = await ScanDefaultWCAG(page);

    expect(accessibilityScanResults.violations).toEqual([]);
});

Then("the wcag guidelines are not yet met", async function (this: OurWorld) {
    const page = this.page;
    const te = await page.title();
    console.log("title", te);
    const accessibilityScanResults = await ScanDefaultWCAG(page);

    expect(accessibilityScanResults.violations).not.toEqual([]);
});

When("the TAB key is pressed {int} times", async function (this: OurWorld, n: number) {
    for (let i = 0; i < n; i++) {
        await this.page.keyboard.press("Tab");
    }
});

When("the enter key is pressed", async function (this: OurWorld) {
    await this.page.keyboard.press("Enter");
});

When("the space key is pressed", async function (this: OurWorld) {
    await this.page.keyboard.press("Space");
});

Then("the button with the title {string} is focussed", async function (this: OurWorld, t: string) {
    await expect(this.page.locator(`.v-btn[title="${t}"]`)).toBeFocused();
});

Then("the checkbox with the aria label {string} is focussed", async function (this: OurWorld, t: string) {
    await expect(this.page.locator(`input[aria-label="${t}"]`)).toBeFocused();
});

When("{string} is typed on the keyboard", async function (this: OurWorld, text: string) {
    await this.page.keyboard.type(text);
});
