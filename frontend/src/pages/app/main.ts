/**
 * main.ts
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Components
import App from '@/pages/app/App.vue'

// Composables
import { createApp } from 'vue'

// Plugins
import { registerPlugins } from '@/plugins'
import CreateI18n from "@/plugins/i18n";
import {german} from "@/pages/app/i18n/german";
import {english} from "@/pages/app/i18n/english";
import { createPinia } from 'pinia'

const app = createApp(App)
  .use(CreateI18n(german, english))
  .use(createPinia())

registerPlugins(app)

app.mount('#app')
