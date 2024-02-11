import {Page} from "@playwright/test";

export default async function waitForAnimations(page: Page, selectors: string[]) {
    for (const id of selectors) {
        const container2 = await page.$(id);
        if (container2) {
            await container2!.waitForElementState("stable");
        } else {
            console.warn(`\nWARNING: element ${id} not found\n`);
        }
    }
}
