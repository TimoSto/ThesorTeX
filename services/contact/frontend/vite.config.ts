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
});
