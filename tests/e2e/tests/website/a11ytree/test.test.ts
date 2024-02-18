import {test} from "@playwright/test";
import getFullUri from "../../../src/helpers/getFullUri";
import {add} from "@thesortex/test";

test("google", async ({page}) => {
    console.log(add(3, 5));
    await page.goto(getFullUri("/"));
});
