import {setup} from "@storybook/vue3";
import {vuetifyInstance} from "@thesortex/vuetify-plugin";

setup((app) => {
    // Registers your app's plugins into Storybook
    app.use(vuetifyInstance);
});
