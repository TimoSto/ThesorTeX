// Plugins
import vue from "@vitejs/plugin-vue";

// Utilities
import {defineConfig} from "vite";
import vuetify from "vite-plugin-vuetify";

const root = process.env.VITE_ROOT;

export default defineConfig({
    root: root ? root : ".",
    plugins: [
        vue(),
        vuetify({
            autoImport: true,
        }),
    ],
    base: "/contact/feedbackComponent",
    build: {
        rollupOptions: {
            input: [root ? root + "/src/main.ts" : "src/main.ts"],
            output: {
                entryFileNames: `[name].js`,
                chunkFileNames: `[name].js`,
                assetFileNames: `assets/[name].[ext]`
            }
        }
    }
    // build: {
    //   lib: {
    //     // Could also be a dictionary or array of multiple entry points
    //     entry: resolve(__dirname, "src/main.ts"),
    //     name: "FeedbackComponent",
    //     // the proper extensions will be added
    //     fileName: "feedback-component",
    //   },
    // },
});
