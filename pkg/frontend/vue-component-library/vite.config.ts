/// <reference types="vitest" />
import {defineConfig} from "vite";
import vue from "@vitejs/plugin-vue";
import vuetify from "vite-plugin-vuetify";

const root = process.env.VITEST_ROOT;

export default defineConfig({
    plugins: [
        vue(),
        vuetify({
            autoImport: true,
        }),
    ],
    test: {
        root: root ? root : ".",
        globals: true,
        environment: "happy-dom",
        setupFiles: ["vuetify.test.config.ts"],
        deps: {
            inline: ["vuetify"],
        },
        testTimeout: 10000
    },
});
