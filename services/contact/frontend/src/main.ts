import App from "./App.vue";
import {GetVuetifyInstance} from "@thesortex/vuetify-plugin";
import {createApp} from "vue";

const app = createApp(App)
  .use(GetVuetifyInstance(false));

// function to be used in a dynamic import
export function renderFeedbackComponent(parent: string) {
  app.mount(parent);
}
