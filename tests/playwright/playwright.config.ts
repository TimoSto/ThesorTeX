import { defineConfig, devices } from '@playwright/test';
import * as process from "node:process";

/**
 * Read environment variables from file.
 * https://github.com/motdotla/dotenv
 */
// import dotenv from 'dotenv';
// dotenv.config({ path: path.resolve(__dirname, '.env') });

/**
 * See https://playwright.dev/docs/test-configuration.
 */
export default defineConfig({
    testDir: './tests',
    /* Run tests in files in parallel */
    fullyParallel: !!process.env.PARALLEL,
    /* Reporter to use. See https://playwright.dev/docs/test-reporters */
    // reporter: [
    //     ["list"],
    //     ["html", {
    //         outputFolder: "out/html",
    //         open: "on-failure",
    //     }],
    // ],
    /* Shared settings for all the projects below. See https://playwright.dev/docs/api/class-testoptions. */
    use: {
        /* Base URL to use in actions like `await page.goto('/')`. */
        // baseURL: 'http://127.0.0.1:3000',

        /* Collect trace when retrying the failed test. See https://playwright.dev/docs/trace-viewer */
        trace: 'on-first-retry',
    },

    /* Configure projects for major browsers */
    projects: [
        {
            name: 'chromium',
            use: {...devices['Desktop Chrome']},
        },
    ],

    webServer: process.env.SUT_EXECUTABLE ? {
        command: process.env.SUT_EXECUTABLE as string,
        url: process.env.BASE_URL,
        reuseExistingServer: true,
        stdout: "pipe",
        stderr: "pipe",
        timeout: 10000,
        env: {
            "E2E_PORT": process.env.E2E_PORT as string,
        }
    } : undefined,
});
