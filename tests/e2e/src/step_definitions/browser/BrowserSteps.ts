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
