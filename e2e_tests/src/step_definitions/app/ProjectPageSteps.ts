import {DataTable, Then} from "@cucumber/cucumber";
import {OurWorld} from "../../../types";
import {expect} from "@playwright/test";

Then("following projects are displayed", async function (this: OurWorld, projects: DataTable) {
    for (const el of projects.hashes()) {
        const i = projects.hashes().indexOf(el);
        expect(await this.page.locator("#page-1").locator("tr").nth(i).locator("td").nth(0).textContent()).toEqual(el.project);
    }
    await expect(this.page.locator("#page-1").locator("tbody tr")).toHaveCount(projects.hashes().length);
});
