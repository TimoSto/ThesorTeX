import {defineConfig} from "vite";
import path, {resolve} from "path";
import dts from "vite-plugin-dts";
import fs from "fs";

// https://vitejs.dev/config/
export default defineConfig({
    build: {
        lib: {
            entry: resolve(__dirname, "src/index.ts"),
            formats: ["es", "cjs"],
            name: "index",
            fileName: "index"
        },
        outDir: resolve(__dirname, "dist"),
    },
    resolve: {
        alias: {
            src: resolve("src/")
        },
    },
    plugins: [
        dts({
            root: resolve(__dirname),
            outDir: resolve(__dirname, "dist"),
        }),
        {
            name: "rename-cjs",
            closeBundle: async () => {
                if (fs.existsSync(path.resolve(resolve(__dirname, "dist"), "index.js")) && !fs.existsSync(path.resolve(resolve(__dirname, "dist"), "index.cjs"))) {
                    fs.renameSync(path.resolve(resolve(__dirname, "dist"), "index.js"), path.resolve(resolve(__dirname, "dist"), "index.cjs"));
                }
            },
        }
    ],
});
