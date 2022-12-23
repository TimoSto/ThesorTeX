import { Given, When } from "@cucumber/cucumber";
import { OurWorld } from "../../support/OurWorld";

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
