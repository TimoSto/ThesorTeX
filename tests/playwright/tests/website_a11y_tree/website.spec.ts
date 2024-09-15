import { expect, test } from "@playwright/test";
import { GetA11yTree } from "@thesortex/a11ytree";

test('has a11y tree', async ({page}) => {
    await page.goto(process.env.BASE_URL!);

    const client = await page.context().newCDPSession(page);

    const tree = await GetA11yTree(client);

    expect(tree.PageTitle).toEqual("ThesorTeX - Use LaTeX more comfortably");

    // Expect a title "to contain" a substring.
    await expect(page).toHaveTitle("ThesorTeX - Use LaTeX more comfortably");
});