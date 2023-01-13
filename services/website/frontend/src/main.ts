import { createApp } from 'vue'
import App from './App.vue'
import { vuetifyInstance } from "@thesortex/vuetify-plugin";

const app = createApp(App)
    .use(vuetifyInstance);

app.mount('#app')
