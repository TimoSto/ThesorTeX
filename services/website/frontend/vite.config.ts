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
            "/templates/download": "http://localhost:8449/"
        }
    },
    build: {
        outDir: "./assets/dist"
    }
});
