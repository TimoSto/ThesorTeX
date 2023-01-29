import {After, AfterAll, Before, BeforeAll} from "@cucumber/cucumber";
import {chromium, ChromiumBrowser, devices} from "playwright";
import {OurWorld} from "./types";
import {spawn} from "child_process";

let browser: ChromiumBrowser;
let sut: any;

BeforeAll(async function () {
    // Browsers are expensive in Playwright so only create 1
    browser = await chromium.launch({
        // Not headless so we can watch test runs
        headless: true,
        // Slow so we can see things happening
        slowMo: 50,
    });

    sut = spawn("../services/app/cmd/e2e/main", {
        // stdio: "ignore",
        detached: false,
    });
    sut.on("error", (err: any) => {
        throw "Could not start system under test executable";
    });
    // sut.stdout.on("data", (data: any) => console.log(data.toString()));
    sut.stderr.on("data", (data: any) => console.error(data.toString()));
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
});
// Cleanup after each scenario
After(async function (this: OurWorld) {
    await this.page.close();
    await this.context.close();
});
