/**
 * main.ts
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Components
import App from '@/pages/auth/AuthPage.vue'

// Composables
import { createApp } from 'vue'

// Plugins
import { registerPlugins } from '@/plugins'
import {german} from "@/pages/auth/i18n/german";
import {english} from "@/pages/auth/i18n/english";
import CreateI18n from "@/plugins/i18n";

const app = createApp(App).use(CreateI18n(german, english))

registerPlugins(app)

app.mount('#app')
