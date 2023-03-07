import {OurWorld} from "../../types";
import {Then} from "@cucumber/cucumber";
import {expect} from "@playwright/test";

Then("the sidebar is closed", async function (this: OurWorld) {
    const classList = await this.page.locator("nav").evaluate((el: Element) => {
        return Array.from(el.classList);
    });

    expect(classList.indexOf("v-navigation-drawer--rail")).toBeGreaterThan(-1);
});

Then("the sidebar is disabled", async function (this: OurWorld) {
    const iconBtn = this.page.locator("header .v-app-bar-nav-icon");
    await expect(iconBtn).toBeDisabled();
});

Then("the sidebar is empty", async function (this: OurWorld) {
    expect(await this.page.locator(".v-navigation-drawer__content").evaluate(el => el.children.length)).toEqual(0);
});
