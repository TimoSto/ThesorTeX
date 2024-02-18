import {Page} from "@playwright/test";
import {AxeResult} from "./types.ts";

export async function scanDefaultWCAG(page: Page): Promise<AxeResult> {

    const t = await page.title();

    console.log(t);

    const client = await page.context().newCDPSession(page);

    const {
        nodes
    } = await client.send("Accessibility.getFullAXTree");

    console.log(nodes.length);

    return {
        violations: [t]
    };
}
