import {APIRequestContext, chromium, ChromiumBrowser, expect, request} from "@playwright/test";
import {spawn} from "child_process";

export async function StartupBrowser(): Promise<ChromiumBrowser> {
    return await chromium.launch({
        // Not headless so we can watch test runs
        headless: true,
        // Slow so we can see things happening
        slowMo: 50,
    });
}

export async function StartupSutIfExists(): Promise<any> {
    let sut: any;

    const sutExec = process.env.EXECUTABLE;
    if (sutExec) {
        console.log("start sut executable: ", sutExec);
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
    } else {
        console.log("not sut executable found in env variables");
    }

    return sut;
}

async function isReachable(ctx: APIRequestContext) {
    try {
        await ctx.head("/");
        return true;
    } catch {
        return false;
    }
}
