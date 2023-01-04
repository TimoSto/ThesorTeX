import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import vuetify from "../plugins/vuetify";
import CreateI18n from "../plugins/i18n";
import { createPinia } from 'pinia'
import {german} from "./i18n/german";
import {english} from "./i18n/english";

const app = createApp(App)
    .use(vuetify)
    .use(CreateI18n(german, english))
    .use(createPinia());

app.mount('#app')
