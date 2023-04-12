/// <reference types="vitest" />
import {defineConfig} from "vite";
import {dirname, resolve} from "node:path";
import {fileURLToPath} from "url";
import vue from "@vitejs/plugin-vue";
import vuetify from "vite-plugin-vuetify";
import VueI18nPlugin from "@intlify/unplugin-vue-i18n/vite";

const root = process.env.VITEST_ROOT;

export default defineConfig({
    root: root ? root : ".",
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
            "/getAllProjects": "http://localhost:8440/",
            "/createNewProject": "http://localhost:8440/",
            "/getProjectData": "http://localhost:8440/",
            "/deleteProject": "http://localhost:8440/",
            "/saveCategory": "http://localhost:8440/",
            "/deleteCategory": "http://localhost:8440/",
            "/saveEntry": "http://localhost:8440/",
            "/deleteEntry": "http://localhost:8440/",
            "/uploadEntries": "http://localhost:8440/",
            "/getConfig": "http://localhost:8440/",
            "/saveConfig": "http://localhost:8440/",
        }
    },
    test: {
        root: root ? root : ".",
        globals: true,
        environment: "happy-dom",
        setupFiles: ["vuetify.config.js", "src/api/mocks/testSetup.ts"],
        deps: {
            inline: ["vuetify"],
        },
    },
});
