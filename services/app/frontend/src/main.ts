import { createApp } from 'vue'
import App from './App.vue'
import vuetify from "./plugins/vuetify";

const app = createApp(App)
    .use(vuetify);

app.mount('#app')
