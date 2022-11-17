const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: [
    'vuetify'
  ],
  pages: {
    index: './src/pages/index/main.ts',
    auth: './src/pages/auth/main.ts',
    app: './src/pages/app/main.ts'
  },
  devServer: {
    proxy: {
      '/': {
        target: 'http://localhost:8448/',
        ws: false
      }
    }
  },
  outputDir: './assets/dist',
})
//TODO: Logo: Graduation Head mit TeX drunter geschrieben
