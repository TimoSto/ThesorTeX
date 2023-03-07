import {Given} from "@cucumber/cucumber";
import {OurWorld} from "../../types";
import getFullUri from "../../helpers/getFullUri";

Given("the url {string} was opened", async function (this: OurWorld, url: string) {
    await this.page.goto(getFullUri(url));
});
