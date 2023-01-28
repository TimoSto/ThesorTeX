import {Page} from "playwright";

export default async function waitForAnimations(page: Page) {
    const container = await page.$(".page--container");
    await container!.waitForElementState("stable");
}
