import {After, AfterAll, Before, BeforeAll} from "@cucumber/cucumber";
import {APIRequestContext, chromium, ChromiumBrowser, devices} from "playwright";
import {expect, request} from "@playwright/test";
import {OurWorld} from "./types";
import {spawn} from "child_process";

let browser: ChromiumBrowser;
let sut: any;

BeforeAll(async function () {
    // Browsers are expensive in Playwright so only create 1
    browser = await chromium.launch({
        // Not headless so we can watch test runs
        headless: false,
        // Slow so we can see things happening
        slowMo: 50,
    });

    const sutExec = process.env.EXECUTABLE;
    console.log("exec", sutExec);
    if (sutExec) {
        sut = spawn(sutExec, {
            // stdio: "ignore",
            detached: false,
        });
        sut.on("exit", (code: number) => {
            if (code && code !== 0) {
                throw `System under test exited with code ${code}`;
            }
        });
        sut.on("error", (err: any) => {
            throw "Could not start system under test executable";
        });
        // sut.stdout.on("data", (data: any) => console.log(data.toString()));
        sut.stderr.on("data", (data: any) => console.error(data.toString()));

        console.log("sbu", process.env.SYSTEM_BASE_URL);
        const ctx = await request.newContext({
            baseURL: process.env.SYSTEM_BASE_URL,
            ignoreHTTPSErrors: true,
        });
        await expect.poll(() => isReachable(ctx), {timeout: 10000}).toBeTruthy();
    }
});

async function isReachable(ctx: APIRequestContext) {
    try {
        await ctx.head("/");
        return true;
    } catch {
        return false;
    }
}

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
