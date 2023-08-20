import App from "./App.vue";
import {GetVuetifyInstance} from "@thesortex/vuetify-plugin";
import {createApp} from "vue";

const app = createApp(App)
    .use(GetVuetifyInstance(false));

console.log("HAAALLLLOOO");

window.addEventListener("DOMContentLoaded", () => {
    app.mount((<any>window).feedbackTrigger);
});
