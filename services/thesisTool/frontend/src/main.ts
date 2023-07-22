import {createApp} from "vue";
import {createPinia} from "pinia";
import App from "./App.vue";
import {GetVuetifyInstance} from "@thesortex/vuetify-plugin";
import {
    accessibilityDialogDe,
    accessibilityDialogEn,
    applicationFrameDe,
    applicationFrameEn,
    installLib
} from "@thesortex/vue-component-library";
import CreateI18n from "@thesortex/vue-i18n-plugin";
import {german} from "./i18n/german";
import {english} from "./i18n/english";

const combinedGerman = Object.assign(german, accessibilityDialogDe, applicationFrameDe);
console.log(combinedGerman);
const combinedEnglish = Object.assign(english, accessibilityDialogEn, applicationFrameEn);

const app = createApp(App)
    .use(GetVuetifyInstance(false))
    .use(createPinia())
    .use(CreateI18n(combinedGerman, combinedEnglish))
    .use(installLib);

app.mount("#app");
