import {expect, test} from "@playwright/test";
import getFullUri from "../../helpers/getFullUri";

test("start", async ({page}) => {
    await page.goto(getFullUri("/"));

    await expect(page).toHaveScreenshot();
});
