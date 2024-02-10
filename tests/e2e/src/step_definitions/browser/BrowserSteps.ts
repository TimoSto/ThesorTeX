import {Given, Then, When} from "@cucumber/cucumber";
import {OurWorld} from "../../types";
import getFullUri from "../../helpers/getFullUri";
import {expect} from "@playwright/test";
import {getAccessibilityTree} from "../../helpers/a11yTree/a11yTree";
import {sum} from "@thesortex/a11ytree";

Given("the accessibility page was opened", async function (this: OurWorld) {
    const client = await this.page.context().newCDPSession(this.page);
    await client.send("Accessibility.enable");
    const tree = await getAccessibilityTree(client, null);
    await new Promise(r => {
        setTimeout(r, 2500);
    });
});

Given("the url {string} was opened", async function (this: OurWorld, url: string) {
    await this.page.goto(getFullUri(url));

    console.log(sum(2, 4));
});

Then("the url is {string}", async function (this: OurWorld, url: string) {
    await expect(this.page.url()).toEqual(getFullUri(url));
});

When("the button with the text {string} is clicked", async function (this: OurWorld, text: string) {
    await this.page.locator(".v-btn", {has: this.page!.locator(`text=${text}`)}).first().click();
});

When("the list item with the text {string} is clicked", async function (this: OurWorld, text: string) {
    await this.page.locator(".v-list-item", {has: this.page!.locator(`text=${text}`)}).first().click();
});

When("the {string} key is pressed", async function (this: OurWorld, key: string) {
    await this.page.keyboard.press(key);
});

When("the {string} key is pressed {int} times", async function (this: OurWorld, key: string, n: number) {
    while (n > 0) {
        await this.page.keyboard.press(key);
        n--;
    }
});

Then("the button with the text {string} is focussed", async function (this: OurWorld, text: string) {
    await expect(this.page.locator(".v-btn", {has: this.page!.locator(`text=${text}`)}).first()).toBeFocused();
});

Then("the list item with the text {string} is focussed", async function (this: OurWorld, text: string) {
    await expect(this.page.locator(".v-list-item", {has: this.page!.locator(`text=${text}`)}).first()).toBeFocused();
});

Then("the first link with the text {string} is focussed", async function (this: OurWorld, text: string) {
    await expect(this.page.locator("a", {has: this.page!.locator(`text=${text}`)}).first()).toBeFocused();
});

Then("the second link with the text {string} is focussed", async function (this: OurWorld, text: string) {
    await expect(this.page.locator("a", {has: this.page!.locator(`text=${text}`)}).nth(1)).toBeFocused();
});

Then("the expansion panel with the title {string} is focussed", async function (this: OurWorld, text: string) {
    const el = await this.page.locator(".v-expansion-panel", {has: this.page!.locator(`text=${text}`)}).first().locator("button");
    await expect(el).toBeVisible();
    await expect(el).toBeFocused();
});

Then("the expansion panel with the title {string} is expanded", async function (this: OurWorld, text: string) {
    const el = await this.page.locator(".v-expansion-panel", {has: this.page!.locator(`text=${text}`)}).first().locator(".v-expansion-panel-text");
    await expect(el).toBeVisible();
    //await expect(el).toBeFocused();
});

When("the browser is reloaded", async function (this: OurWorld) {
    await this.page.reload();
});

Then("a second tab was opened with {string}", async function (this: OurWorld, url: string) {
    await this.page.waitForTimeout(2000);
    let pages = await this.context.pages();
    expect(pages.length).toEqual(2);
    expect(await pages[1].url()).toEqual(url);
});
