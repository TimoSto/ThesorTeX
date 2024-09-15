import { GetA11yTree } from "@thesortex/a11yTree";
import * as fs from "node:fs";
import { expect, test } from '@playwright/test';

test('getA11yTree', async ({page}) => {
    // Here goes the code from playwright codegen


    const client = await page.context().newCDPSession(page);

    const tree = await GetA11yTree(client);

    const exp = JSON.parse(fs.readFileSync("./generated/a11yTree.json", "utf8"));

    expect(tree).toEqual(exp);
});