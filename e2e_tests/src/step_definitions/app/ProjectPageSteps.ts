import {DataTable, Then, When} from "@cucumber/cucumber";
import {OurWorld} from "../../../types";
import {expect} from "@playwright/test";
import waitForAnimations from "../../helpers/waitForAnimations";

Then("following entries are displayed", async function (this: OurWorld, entries: DataTable) {
    for (const el of entries.hashes()) {
        const i = entries.hashes().map(h => h.key).indexOf(el.key);
        expect(await this.page.locator("#page-2").locator(".v-expansion-panel").first().locator("tbody tr").nth(i).locator("td").nth(0).locator("span").textContent()).toEqual(el.key);
    }
    await expect(this.page.locator("#page-2").locator(".v-expansion-panel").first().locator("tbody tr")).toHaveCount(entries.hashes().length);
});

Then("following categories are displayed", async function (this: OurWorld, categories: DataTable) {
    for (const el of categories.hashes()) {
        const i = categories.hashes().map(h => h.name).indexOf(el.name);
        expect(await this.page.locator("#page-2").locator(".v-expansion-panel").nth(1).locator("tbody tr").nth(i).locator("td").nth(0).locator("span").textContent()).toEqual(el.name);
    }
    await expect(this.page.locator("#page-2").locator(".v-expansion-panel").nth(1).locator("tbody tr")).toHaveCount(categories.hashes().length);
});

When("the project is deleted", async function (this: OurWorld) {
    await this.page.locator(".page--container header button").click();
});

Then("the user is asked to confirm the deletion of the project", async function (this: OurWorld) {
    await waitForAnimations(this.page);

    expect(await this.page.locator(".v-overlay__content .v-card-title").textContent()).toEqual("Delete project");
});

Then("the project-page is closed", async function (this: OurWorld) {
    await waitForAnimations(this.page, ["#page-2", "#page-1"]);

    expect(await this.page.locator("#page-2").count()).toEqual(0);
});
