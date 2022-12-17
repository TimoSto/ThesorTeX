// Plugins
import vue from '@vitejs/plugin-vue'
import vuetify from 'vite-plugin-vuetify'

// Utilities
import { defineConfig } from 'vite'
import { fileURLToPath, URL } from 'node:url'
import { resolve } from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    // https://github.com/vuetifyjs/vuetify-loader/tree/next/packages/vite-plugin
    vuetify({
      autoImport: true,
    }),
  ],
  define: { 'process.env': {} },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
    extensions: [
      '.js',
      '.json',
      '.jsx',
      '.mjs',
      '.ts',
      '.tsx',
      '.vue',
    ],
  },
  build: {
    rollupOptions: {
      input: {
        website: resolve(__dirname, 'index.html'),
        info: resolve(__dirname, 'info/index.html'),
        auth: resolve(__dirname, 'auth/index.html'),
        app: resolve(__dirname, 'app/index.html'),
      }
    }
  },
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
    }
  },
})
