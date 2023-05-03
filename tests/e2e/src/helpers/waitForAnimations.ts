import {Page} from "playwright";

export default async function waitForAnimations(page: Page, selectors: string[]) {
    for (const id of selectors) {
        const container2 = await page.$(id);
        await container2!.waitForElementState("stable");
    }
}
