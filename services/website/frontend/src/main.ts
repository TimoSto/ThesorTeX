import {createApp} from "vue";
import App from "./App.vue";
import {vuetifyInstance} from "@thesortex/vuetify-plugin";
import {installLib} from "@thesortex/vue-component-library";
import router from "./router";

const app = createApp(App)
    .use(vuetifyInstance)
    .use(installLib)
    .use(router);

app.mount("#app");
