import {Given, Then, When} from "@cucumber/cucumber";
import {OurWorld} from "../../types";
import getFullUri from "../../helpers/getFullUri";
import {expect} from "@playwright/test";

Given("the url {string} was opened", async function (this: OurWorld, url: string) {
    await this.page.goto(getFullUri(url));
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
