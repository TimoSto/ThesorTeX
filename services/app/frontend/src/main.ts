import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import { vuetifyInstance } from "@thesortex/vuetify-plugin";
import lib from "@thesortex/vue-component-library"

const app = createApp(App)
    .use(vuetifyInstance)
    .use(createPinia())
    .use(lib);

app.mount('#app')
