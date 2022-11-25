/**
 * main.ts
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Components
import App from '@/pages/website/Website.vue'

// Composables
import { createApp } from 'vue'

// Plugins
import { registerPlugins } from '@/plugins'
import CreateI18n from "@/plugins/i18n";
import {german} from "@/pages/website/i18n/german";
import {english} from "@/pages/website/i18n/english";

const app = createApp(App).use(CreateI18n(german, english))

registerPlugins(app)

app.mount('#app')
