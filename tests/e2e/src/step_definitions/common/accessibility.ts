import {Then} from "@cucumber/cucumber";
import {OurWorld} from "../../types";
import {expect} from "@playwright/test";
import AxeBuilder from "@axe-core/playwright";

Then("the wcag guidelines are met", async function (this: OurWorld) {
    const page = this.page;
    const accessibilityScanResults = await new AxeBuilder({page})
        .withTags(["wcag2a", "wcag2aa", "wcag21a", "wcag21aa"])
        .analyze();

    expect(accessibilityScanResults.violations).toEqual([]);
});

Then("the wcag guidelines are not yet met", async function (this: OurWorld) {
    const page = this.page;
    const accessibilityScanResults = await new AxeBuilder({page})
        .withTags(["wcag2a", "wcag2aa", "wcag21a", "wcag21aa"])
        .analyze();

    expect(accessibilityScanResults.violations).not.toEqual([]);
});
