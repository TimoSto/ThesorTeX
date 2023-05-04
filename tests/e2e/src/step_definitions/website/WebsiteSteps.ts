import {Then, When} from "@cucumber/cucumber";
import {OurWorld} from "../../types";
import {expect} from "@playwright/test";
import waitForAnimations from "../../helpers/waitForAnimations";

Then("area {int} is shown in full height", async function (this: OurWorld, n: number) {
    await waitForAnimations(this.page, ["#app"]);

    const element = await this.page.locator(".fullHeightContainer").nth(n - 1);
    await expect(element).toBeVisible();
    const box = await element.boundingBox();
    const pageDimensions = await this.page.viewportSize();

    expect(box!.width).toEqual(pageDimensions!.width);
    expect(box!.height).toEqual(pageDimensions!.height);

    const scroll = await this.page.evaluate(() => {
        return window.scrollY;
    });

    expect(scroll).toEqual(pageDimensions!.height * (n - 1));
});

When("the scroll down button is clicked", async function (this: OurWorld, n: number) {
    await this.page.locator(".scroll-down").click({force: true});
});

Then("the title of area {int} is {string}", async function (this: OurWorld, n: number, title: string) {
    const foundTitle = await this.page.locator(".fullHeightContainer").nth(n - 1).locator(".text-h3").textContent();

    expect(foundTitle).toEqual(title);
});

Then("the title of area {int} is not existent", async function (this: OurWorld, n: number) {
    await expect(this.page.locator(".fullHeightContainer").nth(n - 1).locator(".text-h3")).toHaveCount(0);
});

Then("three cards are shown for the products", async function (this: OurWorld) {
    const cards = await this.page.locator(".fullHeightContainer").nth(0).locator(".v-card");
    await expect(cards).toHaveCount(3);
    await expect(cards.nth(0).locator(".text-h5")).toHaveText("Template for academic papers");
    await expect(cards.nth(1).locator(".text-h5")).toHaveText("Tool for bibliography management");
    await expect(cards.nth(2).locator(".text-h5")).toHaveText("Template for a curriculum vitae");
});

When("the download button in card {int} is clicked", async function (this: OurWorld, n: number) {
    await this.page.locator(".fullHeightContainer").nth(0).locator(".v-btn", {has: this.page!.locator(`text=Download`)}).nth(n - 1).click();
});
