import {After, AfterAll, Before, BeforeAll} from "@cucumber/cucumber";
import {ChromiumBrowser, devices} from "@playwright/test";
import {OurWorld} from "./types";
import {StartupBrowser, StartupSutIfExists} from "./helpers/startup";

let browser: ChromiumBrowser;
let sut: any;

BeforeAll(async function () {
    // Browsers are expensive in Playwright so only create 1
    browser = await StartupBrowser();

    sut = await StartupSutIfExists();
});

AfterAll(async function () {
    await browser.close();
    await sut?.kill();
});
// Create a new test context and page per scenario
Before(async function (this: OurWorld) {
    const pixel2 = devices["Pixel 2"];
    this.context = await browser.newContext({
        viewport: {
            width: 1200,
            height: 800
        },
        locale: "en"
    });
    this.page = await this.context.newPage();
    this.page.context().addCookies([{
        name: "IgnoreThesorTeXUpdates",
        value: "mytokenvalue123",
        url: process.env.SYSTEM_BASE_URL
    }]);
});
// Cleanup after each scenario
After(async function (this: OurWorld) {
    await this.page.close();
    await this.context.close();
});
