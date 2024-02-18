import {test} from "@playwright/test";
import getFullUri from "../../../src/helpers/getFullUri";
import {add} from "@thesortex/test";
import {ScanDefaultWCAG} from "@thesortex/playwright-a11y-helper";

test("google", async ({page}) => {
    console.log(add(3, 5));
    await page.goto(getFullUri("/"));

    const res = await ScanDefaultWCAG(page);

    console.log(res);
});
