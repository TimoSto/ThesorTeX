import {resolve} from "path";
import {defineConfig} from "vite";
import dts from "vite-plugin-dts";

// https://vitejs.dev/guide/build.html#library-mode
export default defineConfig({
    build: {
        lib: {
            entry: resolve(__dirname, "src/index.ts"),
            name: "a11ytree",
            fileName: "a11ytree",
        },
    },
    plugins: [dts({
        root: resolve(__dirname),
        outDir: resolve(__dirname, "dist"),
    })],
});