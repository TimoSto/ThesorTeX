import {APIRequestContext, expect, request} from "@playwright/test";
import {ChildProcessWithoutNullStreams, spawn} from "child_process";

export let sut: ChildProcessWithoutNullStreams;

export default async function globalSetup() {
    if (process.env.SUT_EXECUTABLE) {
        sut = spawn(process.env.SUT_EXECUTABLE, {
            // stdio: "ignore",
            detached: false,
        });
        sut.on("error", (err: any) => {
            throw "Could not start system under test executable";
        });
        sut.stdout.on("data", (data: any) => console.log(data.toString()));
        sut.stderr.on("data", (data: any) => console.error(data.toString()));
    }

    const ctx = await request.newContext({
        baseURL: process.env.SYSTEM_BASE_URL,
    });
    await expect.poll(() => isReachable(ctx), {timeout: 30000}).toBeTruthy();
};

async function isReachable(ctx: APIRequestContext) {
    try {
        await ctx.head("/");
        return true;
    } catch {
        return false;
    }
}
