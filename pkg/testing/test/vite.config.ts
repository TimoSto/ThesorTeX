import {defineConfig} from "vite";
import {resolve} from "path";
import dts from "vite-plugin-dts";

// https://vitejs.dev/config/
export default defineConfig({
    build: {
        lib: {
            entry: resolve(__dirname, "src/main.ts"),
            formats: ["es", "cjs"],
            name: "test",
            fileName: "test"
        },
        outDir: resolve(__dirname, "dist"),
    },
    resolve: {
        alias: {
            src: resolve("src/")
        },
    },
    plugins: [dts({
        root: resolve(__dirname),
        outDir: resolve(__dirname, "dist"),
    })],
});
