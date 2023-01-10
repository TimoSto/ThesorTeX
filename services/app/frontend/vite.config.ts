import { defineConfig } from 'vite'
import { resolve, dirname } from 'node:path'
import { fileURLToPath } from 'url'
import vue from '@vitejs/plugin-vue'
import vuetify from "vite-plugin-vuetify";
import VueI18nPlugin from '@intlify/unplugin-vue-i18n/vite'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vuetify({
      autoImport: true,
    }),
    VueI18nPlugin({
      /* options */
      // locale messages resource pre-compile option
      include: resolve(dirname(fileURLToPath(import.meta.url)), './src/i18n/**'),
    })
  ],
  server: {
    port: 3000,
    proxy: {
      "/app/createProject": "http://localhost:8448/",
      "/app/projects": "http://localhost:8448/",
      "/app/deleteProject": "http://localhost:8448/",
      "/app/projectData": "http://localhost:8448/",
      "/app/saveEntry": "http://localhost:8448/",
      "/app/saveCategory": "http://localhost:8448/",
      "/app/config": "http://localhost:8448/",
      "/app/renameProject": "http://localhost:8448/",
      "/app/deleteEntry": "http://localhost:8448/",
      "/app/deleteCategory": "http://localhost:8448/",
      "/app/uploadEntries": "http://localhost:8448/",
    }
  },
  build: {
    outDir: './assets/dist'
  }
})
