import {createApp} from "vue";
import App from "./App.vue";
import {vuetifyInstance} from "@thesortex/vuetify-plugin";
import {installLib} from "@thesortex/vue-component-library";
import router from "./router";
import CreateI18n from "@thesortex/vue-i18n-plugin";
import {german} from "./i18n/german";
import {english} from "./i18n/english";

const app = createApp(App)
    .use(vuetifyInstance)
    .use(installLib)
    .use(CreateI18n(german, english))
    .use(router);

app.mount("#app");
