import {defineConfig} from "vite";
import vue from "@vitejs/plugin-vue";
import vuetify from "vite-plugin-vuetify";

export default defineConfig({
    define: {
        __VUE_I18N_FULL_INSTALL__: true,
        __VUE_I18N_LEGACY_API__: false,
        __INTLIFY_PROD_DEVTOOLS__: false,
    },
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
            "/documentation": "http://localhost:8449/",
        }
    },
    build: {
        outDir: "./assets/dist"
    }
});
