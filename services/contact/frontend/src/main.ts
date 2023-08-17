/**
 * main.ts
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Components
import App from "./App.vue";
import {GetVuetifyInstance} from "@thesortex/vuetify-plugin";

// Composables
import {createApp} from "vue";

const app = createApp(App)
  .use(GetVuetifyInstance(false));

export function renderFeedbackComponent(parent: string) {
  app.mount(parent);
}
