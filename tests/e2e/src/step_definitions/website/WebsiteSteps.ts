import {DataTable, Then, When} from "@cucumber/cucumber";
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

    expect(Math.abs(scroll - pageDimensions!.height * (n - 1))).toBeLessThan(10);
});

Then("area {int} is shown in full height slightly scrolled off", async function (this: OurWorld, n: number) {
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

    expect(Math.abs(scroll - pageDimensions!.height * (n - 1))).toBeLessThan(200);
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

When("the more info button in card {int} is clicked", async function (this: OurWorld, n: number) {
    await this.page.locator(".fullHeightContainer").nth(0).locator(".v-btn", {has: this.page!.locator(`text=Learn more`)}).nth(n - 1).click();
});

Then("the following versions are shown in area {int}", async function (this: OurWorld, n: number, versions: DataTable) {
    for (const el of versions.hashes()) {
        const i = versions.hashes().map(v => v.name).indexOf(el.name);
        let v = await this.page.locator(".fullHeightContainer").nth(n - 1).locator("tbody tr").nth(i).locator("td").nth(0).textContent();
        await expect(v).toEqual(el.name);
        let d = await this.page.locator(".fullHeightContainer").nth(n - 1).locator("tbody tr").nth(i).locator("td").nth(1).textContent();
        await expect(d).toEqual(el.date);
    }
});

When("the list item with the text {string} in area {int} is clicked", async function (this: OurWorld, text: string, n: number) {
    await this.page.locator(".fullHeightContainer").nth(n - 1).locator(".v-list-item", {has: this.page!.locator(`text=${text}`)}).first().click();
});

Then("the list item with the text {string} in area {int} is focussed", async function (this: OurWorld, text: string, n: number) {
    await expect(this.page.locator(".fullHeightContainer").nth(n - 1).locator(".v-list-item", {has: this.page!.locator(`text=${text}`)}).first()).toBeFocused();
});

Then("the button with the text {string} in area {int} and card {int} is focussed", async function (this: OurWorld, text: string, n: number, m: number) {
    await expect(this.page.locator(".fullHeightContainer").nth(n - 1).locator(".v-card").nth(m - 1).locator(".v-btn", {has: this.page!.locator(`text=${text}`)}).first()).toBeFocused();
});
