import {defineConfig} from "vite";
import path, {resolve} from "path";
import dts from "vite-plugin-dts";
import fs from "fs";

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
    plugins: [
        dts({
            root: resolve(__dirname),
            outDir: resolve(__dirname, "dist"),
        }),
        {
            name: "rename-cjs",
            closeBundle: async () => {
                if (fs.existsSync(path.resolve(resolve(__dirname, "dist"), "test.js"))) {
                    fs.renameSync(path.resolve(resolve(__dirname, "dist"), "test.js"), path.resolve(resolve(__dirname, "dist"), "test.cjs"));
                }
            },
        }
    ],
});
