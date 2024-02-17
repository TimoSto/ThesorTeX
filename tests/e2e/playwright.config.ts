import type {PlaywrightTestConfig} from "@playwright/test";

const config: PlaywrightTestConfig = {
    testDir: process.env.TARGET ? `./tests/${process.env.TARGET}/` : "./tests/",
    globalSetup: require.resolve("./src/hooks/global-setup.ts"),
    globalTeardown: require.resolve("./src/hooks/global-teardown.ts"),
    use: {
        launchOptions: {
            executablePath: process.env.CHROME_PATH || undefined,
        },
    },
};
export default config;