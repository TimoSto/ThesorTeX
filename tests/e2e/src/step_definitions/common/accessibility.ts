import {Then} from "@cucumber/cucumber";
import {OurWorld} from "../../types";
import {expect} from "@playwright/test";
import AxeBuilder from "@axe-core/playwright";

Then("the wcag guidelines are met", async function (this: OurWorld) {
    const page = this.page;
    const accessibilityScanResults = await new AxeBuilder({page}).analyze();

    expect(accessibilityScanResults.violations).toEqual([]);
});
