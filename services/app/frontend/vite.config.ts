/// <reference types="vitest" />
import { defineConfig } from 'vite'
import { resolve, dirname } from 'node:path'
import { fileURLToPath } from 'url'
import vue from '@vitejs/plugin-vue'
import vuetify from "vite-plugin-vuetify";
import VueI18nPlugin from '@intlify/unplugin-vue-i18n/vite'

export default defineConfig({
    plugins: [
        vue(),
        vuetify({
            autoImport: true,
        }),
        VueI18nPlugin({
            /* options */
            // locale messages resource pre-compile option
            include: resolve(dirname(fileURLToPath(import.meta.url)), './src/i18n/**'),
        }),
    ],
    server: {
        port: 3000,
        proxy: {
            "/getAllProjects": "http://localhost:8448/",
            "/createNewProject": "http://localhost:8448/",
        }
    },
    build: {
        outDir: './assets/dist'
    },
    test: {
        globals: true,
        environment: "happy-dom",
        setupFiles: ["vuetify.config.js", "src/api/mocks/testSetup.ts"],
        deps: {
            inline: ["vuetify"],
        },
    },
})
