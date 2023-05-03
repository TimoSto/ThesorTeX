import {DataTable, Then, When} from "@cucumber/cucumber";
import {OurWorld} from "../../types";
import {expect} from "@playwright/test";
import waitForAnimations from "../../helpers/waitForAnimations";

Then("following projects are displayed", async function (this: OurWorld, projects: DataTable) {
    for (const el of projects.hashes()) {
        const i = projects.hashes().map(h => h.project).indexOf(el.project);
        expect(await this.page.locator("#page-1").locator("tbody tr").nth(i).locator("td").nth(0).textContent()).toEqual(el.project);
    }
    await expect(this.page.locator("#page-1").locator("tbody tr")).toHaveCount(projects.hashes().length);
});

When("the project {string} is opened", async function (this: OurWorld, project: string) {
    await this.page.locator("#page-1").locator("tr", {has: this.page.locator("td", {hasText: project})}).click();
});

When("a new project is added", async function (this: OurWorld) {
    await this.page.locator("#page-1").locator("thead").locator("th").last().locator("button").click();
});

Then("the dialog for project creation is shown", async function (this: OurWorld) {
    await waitForAnimations(this.page, [".pages", ".v-overlay-container"]);

    expect(await this.page.locator(".v-overlay__content .v-card-title").textContent()).toEqual("Create new project");
});

When("the name {string} is entered into the projectname field", async function (this: OurWorld, name: string) {
    await this.page.locator(".v-overlay__content .v-card-text input").type(name);
});

When("the create project button is clicked", async function (this: OurWorld) {
    await this.page.locator(".v-overlay__content button").nth(1).click();
});
