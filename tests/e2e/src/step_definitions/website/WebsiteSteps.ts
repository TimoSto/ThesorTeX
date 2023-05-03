import {Then} from "@cucumber/cucumber";
import {OurWorld} from "../../types";
import {expect} from "@playwright/test";

Then("area {int} is shown in full height", async function (this: OurWorld, n: number) {
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
