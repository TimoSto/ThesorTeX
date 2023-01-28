import {Page} from "playwright";

export default async function waitForAnimations(page: Page, idsToWaitFor?: string[]) {
    const container = await page.$(".pages");
    await container!.waitForElementState("stable");

    if (idsToWaitFor) {
        for (const id of idsToWaitFor) {
            const container2 = await page.$(id);
            await container2!.waitForElementState("stable");
        }
    }

    const overlay = await page.$(".v-overlay-container");
    await overlay!.waitForElementState("stable");
}
