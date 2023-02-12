import {defineConfig} from "vite";
import vue from "@vitejs/plugin-vue";
import vuetify from "vite-plugin-vuetify";

export default defineConfig({
    plugins: [
        vue(),
        vuetify({
            autoImport: true,
        }),
    ],
    server: {
        port: 3001,
        proxy: {
            "/templates/": "http://localhost:8449/",
            "/versions/": "http://localhost:8449/",
        }
    },
    build: {
        outDir: "./assets/dist"
    }
});
