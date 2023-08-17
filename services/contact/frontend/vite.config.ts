// Plugins
import vue from "@vitejs/plugin-vue";

// Utilities
import {defineConfig} from "vite";
import vuetify from "vite-plugin-vuetify";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vuetify({
      autoImport: true,
    }),
  ],
  build: {
    rollupOptions: {
      input: ["src/main.ts"],
      output: {
        entryFileNames: `assets/[name].js`,
        chunkFileNames: `assets/[name].js`,
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
