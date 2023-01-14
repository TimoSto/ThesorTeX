import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import { vuetifyInstance } from "@thesortex/vuetify-plugin";

const app = createApp(App)
    .use(vuetifyInstance)
    .use(createPinia());

app.mount('#app')
