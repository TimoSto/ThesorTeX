import {Page} from "@playwright/test";
import {AxeResult} from "./types.ts";

export async function scanDefaultWCAG(page: Page): Promise<AxeResult> {

    const t = await page.title();

    console.log(t);

    return {
        violations: [t]
    };
}
