import {createApp} from "vue";
import App from "./App.vue";
import {vuetifyInstance} from "@thesortex/vuetify-plugin";
import {installLib} from "@thesortex/vue-component-library";

const app = createApp(App)
    .use(vuetifyInstance)
    .use(installLib);

app.mount("#app");
