const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: [
    'vuetify'
  ],
  pages: {
    index: './src/pages/index/main.ts',
    auth: './src/pages/auth/main.ts',
    app: './src/pages/app/main.ts'
  }
})
