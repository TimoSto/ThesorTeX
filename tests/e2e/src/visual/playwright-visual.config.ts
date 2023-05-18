import type {PlaywrightTestConfig} from "@playwright/test";

const config: PlaywrightTestConfig = {
    testDir: "./",
    snapshotPathTemplate: "{testDir}/references/{testFileName}/{arg}-{platform}{ext}",
    outputDir: "../../test-results",
    globalSetup: require.resolve("./setup.ts"),
    globalTeardown: require.resolve("./teardown.ts"),
    use: {
        launchOptions: {
            executablePath: process.env.CHROME_PATH || undefined,
        },
    },
};
export default config;