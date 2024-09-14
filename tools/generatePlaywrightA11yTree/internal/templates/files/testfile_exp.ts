import { getAccessibilityTree } from "../src/a11yTree/a11yTree";
import * as fs from "node:fs";
import { expect, test } from '@playwright/test';

test('getA11yTree', async ({page}) => {
    // Here goes the code from playwright codegen

    await page.goto('about:blank');
    await page.goto('chrome-error://chromewebdata/');
    await page.goto('http://localhost:8449/');
    await page.goto('http://localhost:8449/#/');
    await page.getByText('Vorlage für eine Abschluss- / Hausarbeit').first().click();
    await page.getByText('Vorlage für eine Abschluss- / Hausarbeit').first().click();
    await page.getByRole('link', {name: 'Zu den Downloads'}).first().click();
    await page.locator('td:nth-child(4) > .v-btn').first().click();


    const client = await page.context().newCDPSession(page);

    const tree = await getAccessibilityTree(client);

    const file = JSON.stringify(tree, null, 2);
    fs.writeFile("./generated/a11yTree.json", file, (err) => {
        expect(err).toBeNull();
    });
});