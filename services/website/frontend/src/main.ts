import {createApp} from "vue";
import App from "./App.vue";
import {vuetifyInstance} from "@thesortex/vuetify-plugin";
import {installLib} from "@thesortex/vue-component-library";
import {installLib as svgInstallLib} from "@thesortex/svg-generator";
import router from "./router";

const app = createApp(App)
    .use(vuetifyInstance)
    .use(installLib)
    .use(svgInstallLib)
    .use(router);

app.mount("#app");
