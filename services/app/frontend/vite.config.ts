/// <reference types="vitest" />
import {defineConfig} from "vite";
import {dirname, resolve} from "node:path";
import {fileURLToPath} from "url";
import vue from "@vitejs/plugin-vue";
import vuetify from "vite-plugin-vuetify";
import VueI18nPlugin from "@intlify/unplugin-vue-i18n/vite";

export default defineConfig({
    plugins: [
        vue(),
        vuetify({
            autoImport: true,
        }),
        VueI18nPlugin({
            /* options */
            // locale messages resource pre-compile option
            include: resolve(dirname(fileURLToPath(import.meta.url)), "./src/i18n/**"),
        }),
    ],
    server: {
        port: 3000,
        proxy: {
            "/getAllProjects": "http://localhost:8448/",
            "/createNewProject": "http://localhost:8448/",
            "/getProjectData": "http://localhost:8448/",
            "/deleteProject": "http://localhost:8448/",
            "/saveCategory": "http://localhost:8448/",
            "/deleteCategory": "http://localhost:8448/",
            "/saveEntry": "http://localhost:8448/",
            "/deleteEntry": "http://localhost:8448/",
            "/uploadEntries": "http://localhost:8448/",
            "/getConfig": "http://localhost:8448/",
            "/saveConfig": "http://localhost:8448/",
        }
    },
    build: {
        outDir: "./assets/dist"
    },
    test: {
        globals: true,
        environment: "happy-dom",
        setupFiles: ["vuetify.config.js", "src/api/mocks/testSetup.ts"],
        deps: {
            inline: ["vuetify"],
        },
    },
});
