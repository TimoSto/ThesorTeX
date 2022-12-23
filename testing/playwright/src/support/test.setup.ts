import { Before, BeforeAll, AfterAll, After } from "@cucumber/cucumber";
import { devices, chromium } from "playwright";
import { OurWorld } from "./OurWorld";
BeforeAll(async function () {
    // Browsers are expensive in Playwright so only create 1
    (global as any).browser = await chromium.launch({
        // Not headless so we can watch test runs
        headless: false,//TODO: set dynmaically
        // Slow so we can see things happening
        slowMo: 50,
    });
});
AfterAll(async function () {
    await (global as any).browser.close();
});
// Create a new test context and page per scenario
Before(async function (this: OurWorld) {
    const chrome_desktop = devices["Desktop Chrome"];//TODO: set dynamically
    this.context = await (global as any).browser.newContext({
        viewport: { width: 1300, height: 700 },
        userAgent: chrome_desktop.userAgent,
        locale: "en-EN",
    });
    this.page = await this.context.newPage();
});
// Cleanup after each scenario
After(async function (this: OurWorld) {
    await this.page.close();
    await this.context.close();
});
